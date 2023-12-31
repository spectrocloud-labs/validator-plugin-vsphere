package vsphere

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/govc/host/service"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/session"
	"github.com/vmware/govmomi/session/keepalive"
	ssoadmintypes "github.com/vmware/govmomi/ssoadmin/types"
	"github.com/vmware/govmomi/vapi/rest"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/soap"
	"github.com/vmware/govmomi/vim25/types"
	"golang.org/x/exp/slices"
)

const KeepAliveIntervalInMinute = 10

var sessionCache = map[string]Session{}
var sessionMU sync.Mutex
var restClientLoggedOut = false

type VsphereCloudAccount struct {
	// Insecure is a flag that controls whether to validate the vSphere server's certificate.
	Insecure bool `json:"insecure"`

	// password
	// Required: true
	Password string `json:"password"`

	// username
	// Required: true
	Username string `json:"username"`

	// VcenterServer is the address of the vSphere endpoint
	// Required: true
	VcenterServer string `json:"vcenterServer"`
}

type Session struct {
	GovmomiClient *govmomi.Client
	RestClient    *rest.Client
}

type VSphereCloudDriver struct {
	VCenterServer   string
	VCenterUsername string
	VCenterPassword string
	Datacenter      string
	Client          *govmomi.Client
	RestClient      *rest.Client
}

func NewVSphereDriver(VCenterServer, VCenterUsername, VCenterPassword, datacenter string) (*VSphereCloudDriver, error) {
	session, err := GetOrCreateSession(context.TODO(), VCenterServer, VCenterUsername, VCenterPassword, true)
	if err != nil {
		return nil, err
	}

	return &VSphereCloudDriver{
		VCenterServer:   VCenterServer,
		VCenterUsername: VCenterUsername,
		VCenterPassword: VCenterPassword,
		Datacenter:      datacenter,
		Client:          session.GovmomiClient,
		RestClient:      session.RestClient,
	}, nil
}

func (v *VSphereCloudDriver) GetCurrentVmwareUser(ctx context.Context) (string, error) {
	userSession, err := v.Client.SessionManager.UserSession(ctx)
	if err != nil {
		return "", err
	}

	return userSession.UserName, nil
}

func (v *VSphereCloudDriver) ValidateUserPrivilegeOnEntities(ctx context.Context, authManager *object.AuthorizationManager, datacenter string, finder *find.Finder, entityName, entityType string, privileges []string, userName, clusterName string) (isValid bool, failures []string, err error) {
	var folder *object.Folder
	var cluster *object.ClusterComputeResource
	var host *object.HostSystem
	var vapp *object.VirtualApp
	var resourcePool *object.ResourcePool
	var vm *object.VirtualMachine

	var moID types.ManagedObjectReference

	switch entityType {
	case "folder":
		_, folder, err = v.GetFolderIfExists(ctx, finder, datacenter, entityName)
		if err != nil {
			return false, failures, err
		}
		moID = folder.Reference()
	case "resourcepool":
		_, resourcePool, err = v.GetResourcePoolIfExists(ctx, finder, datacenter, clusterName, entityName)
		if err != nil {
			return false, failures, err
		}
		moID = resourcePool.Reference()
	case "vapp":
		_, vapp, err = v.GetVAppIfExists(ctx, finder, datacenter, entityName)
		if err != nil {
			return false, failures, err
		}
		moID = vapp.Reference()
	case "vm":
		_, vm, err = v.GetVMIfExists(ctx, finder, datacenter, clusterName, entityName)
		if err != nil {
			return false, failures, err
		}
		moID = vm.Reference()
	case "host":
		_, host, err = v.GetHostIfExists(ctx, finder, datacenter, clusterName, entityName)
		if err != nil {
			return false, failures, err
		}
		moID = host.Reference()
	case "cluster":
		_, cluster, err = v.GetClusterIfExists(ctx, finder, datacenter, entityName)
		if err != nil {
			return false, failures, err
		}
		moID = cluster.Reference()
	}

	userPrincipal := getUserPrincipalFromUsername(userName)
	privilegeResult, err := authManager.FetchUserPrivilegeOnEntities(ctx, []types.ManagedObjectReference{moID}, userPrincipal)
	if err != nil {
		return false, failures, err
	}

	privilegesMap := make(map[string]bool)
	for _, result := range privilegeResult {
		for _, privilege := range result.Privileges {
			privilegesMap[privilege] = true
		}
	}

	for _, privilege := range privileges {
		if _, ok := privilegesMap[privilege]; !ok {
			err = fmt.Errorf("some entity privileges were not found for user: %s", userName)
			failures = append(failures, fmt.Sprintf("user: %s does not have privilege: %s on entity type: %s with name: %s", userName, privilege, entityType, entityName))
		}
	}

	if len(failures) == 0 {
		isValid = true
	}

	return isValid, failures, nil
}

func GetOrCreateSession(
	ctx context.Context,
	server, username, password string, refreshRestClient bool) (Session, error) {

	sessionMU.Lock()
	defer sessionMU.Unlock()

	sessionKey := server + username
	currentSession, ok := sessionCache[sessionKey]

	if ok {
		if refreshRestClient && restClientLoggedOut {
			//Rest Client
			restClient, err := createRestClientWithKeepAlive(ctx, sessionKey, username, password, currentSession.GovmomiClient)
			if err != nil {
				return currentSession, err
			}
			currentSession.RestClient = restClient
			restClientLoggedOut = false
		}
		return currentSession, nil
	}

	// govmomi Client
	govClient, err := createGovmomiClientWithKeepAlive(ctx, sessionKey, server, username, password)
	if err != nil {
		return currentSession, err
	}

	//Rest Client
	restClient, err := createRestClientWithKeepAlive(ctx, sessionKey, username, password, govClient)
	if err != nil {
		return currentSession, err
	}

	currentSession.GovmomiClient = govClient
	currentSession.RestClient = restClient

	// Cache the currentSession.
	sessionCache[sessionKey] = currentSession
	return currentSession, nil
}

func createGovmomiClientWithKeepAlive(ctx context.Context, sessionKey, server, username, password string) (*govmomi.Client, error) {
	//get vcenter URL
	vCenterURL, err := getVCenterUrl(server, username, password)
	if err != nil {
		return nil, err
	}

	insecure := true

	soapClient := soap.NewClient(vCenterURL, insecure)
	vimClient, err := vim25.NewClient(ctx, soapClient)
	if err != nil {
		return nil, err
	}

	vimClient.UserAgent = "spectro-palette"

	c := &govmomi.Client{
		Client:         vimClient,
		SessionManager: session.NewManager(vimClient),
	}

	send := func() error {
		ctx := context.Background()
		_, err := methods.GetCurrentTime(ctx, vimClient.RoundTripper)
		if err != nil {
			ClearCache(sessionKey)
		}
		return err
	}

	// this starts the keep alive handler when Login is called, and stops the handler when Logout is called
	// it'll also stop the handler when send() returns error, so we wrap around the default send()
	// with err check to clear cache in case of error
	vimClient.RoundTripper = keepalive.NewHandlerSOAP(vimClient.RoundTripper, KeepAliveIntervalInMinute*time.Minute, send)

	// Only login if the URL contains user information.
	if vCenterURL.User != nil {
		err = c.Login(ctx, vCenterURL.User)
		if err != nil {
			return nil, err
		}
	}

	return c, nil
}

func getVCenterUrl(vCenterServer string, vCenterUsername string, vCenterPassword string) (*url.URL, error) {
	// parse vcenter URL
	for _, scheme := range []string{"http://", "https://"} {
		vCenterServer = strings.TrimPrefix(vCenterServer, scheme)
	}
	vCenterServer = fmt.Sprintf("https://%s/sdk", strings.TrimSuffix(vCenterServer, "/"))

	vCenterURL, err := url.Parse(vCenterServer)
	if err != nil {
		return nil, errors.Errorf("invalid vCenter server")

	}
	vCenterURL.User = url.UserPassword(vCenterUsername, vCenterPassword)

	return vCenterURL, nil
}

func createRestClientWithKeepAlive(ctx context.Context, sessionKey, username, password string, govClient *govmomi.Client) (*rest.Client, error) {
	// create RestClient for operations like get tags
	restClient := rest.NewClient(govClient.Client)

	err := restClient.Login(ctx, url.UserPassword(username, password))
	if err != nil {
		return nil, err
	}

	return restClient, nil
}

func ClearCache(sessionKey string) {
	sessionMU.Lock()
	defer sessionMU.Unlock()
	delete(sessionCache, sessionKey)
}

func (v *VSphereCloudDriver) CreateVSphereVMFolder(ctx context.Context, datacenter string, folders []string) error {
	finder, _, err := v.getFinderWithDatacenter(ctx, datacenter)
	if err != nil {
		return err
	}

	for _, folder := range folders {
		folderExists, _, err := v.GetFolderIfExists(ctx, finder, datacenter, folder)
		if folderExists {
			continue
		}

		dir := path.Dir(folder)
		name := path.Base(folder)

		if dir == "" {
			dir = "/"
		}

		folder, err := finder.Folder(ctx, dir)
		if err != nil {
			return fmt.Errorf("error fetching folder: %s. Code:%d", err.Error(), http.StatusInternalServerError)
		}

		if _, err := folder.CreateFolder(ctx, name); err != nil {
			return fmt.Errorf("error creating folder: %s. Code:%d", err.Error(), http.StatusInternalServerError)
		}
	}

	return nil
}

func (v *VSphereCloudDriver) getFinderWithDatacenter(ctx context.Context, datacenter string) (*find.Finder, string, error) {
	finder, err := v.getFinder(ctx)
	if err != nil {
		return nil, "", err
	}
	dc, govErr := finder.DatacenterOrDefault(ctx, datacenter)
	if govErr != nil {
		return nil, "", fmt.Errorf("failed to fetch datacenter: %s. code: %d", govErr.Error(), http.StatusBadRequest)
	}
	//set the datacenter
	finder.SetDatacenter(dc)

	return finder, dc.Name(), nil
}

func (v *VSphereCloudDriver) getFinder(ctx context.Context) (*find.Finder, error) {
	if v.Client == nil {
		return nil, fmt.Errorf("failed to fetch govmomi client: %d", http.StatusBadRequest)
	}

	finder := find.NewFinder(v.Client.Client, true)
	return finder, nil
}

func (v *VSphereCloudDriver) FolderExists(ctx context.Context, finder *find.Finder, datacenter, folderName string) (bool, error) {

	if _, err := finder.Folder(ctx, folderName); err != nil {
		return false, nil
	}
	return true, nil
}

func (v *VSphereCloudDriver) GetFolderNameByID(ctx context.Context, datacenter, id string) (string, error) {
	finder, dc, err := v.getFinderWithDatacenter(ctx, datacenter)
	if err != nil {
		return "", err
	}

	fos, govErr := finder.FolderList(ctx, "*")
	if govErr != nil {
		return "", fmt.Errorf("failed to fetch vSphere folders. Datacenter: %s, Error: %s", datacenter, govErr.Error())
	}

	prefix := fmt.Sprintf("/%s/vm/", dc)
	for _, fo := range fos {
		inventoryPath := fo.InventoryPath
		//get vm folders, items with path prefix '/{Datacenter}/vm'
		if strings.HasPrefix(inventoryPath, prefix) {
			folderName := strings.TrimPrefix(inventoryPath, prefix)
			//skip spectro folders & sub-folders
			if !strings.HasPrefix(folderName, "spc-") && !strings.Contains(folderName, "/spc-") {
				if fo.Reference().Value == id {
					return folderName, nil
				}
			}
		}
	}

	return "", fmt.Errorf("unable to find folder with id: %s", id)
}

func (v *VSphereCloudDriver) GetFinderWithDatacenter(ctx context.Context, datacenter string) (*find.Finder, string, error) {
	finder, err := v.getFinder(ctx)
	if err != nil {
		return nil, "", err
	}
	dc, govErr := finder.DatacenterOrDefault(ctx, datacenter)
	if govErr != nil {
		return nil, "", fmt.Errorf("failed to fetch datacenter: %s. code: %s"+govErr.Error(), http.StatusBadRequest)
	}
	//set the datacenter
	finder.SetDatacenter(dc)

	return finder, dc.Name(), nil
}

func GetVmwareUserPrivileges(userPrincipal string, groupPrincipals []string, authManager *object.AuthorizationManager) (map[string]bool, error) {
	groupPrincipalMap := make(map[string]bool)
	for _, principal := range groupPrincipals {
		groupPrincipalMap[principal] = true
	}

	// Get the current user's roles
	authRoles, err := authManager.RoleList(context.TODO())
	if err != nil {
		return nil, err
	}

	// create a map to store privileges for current user
	privileges := make(map[string]bool)

	// Print the roles
	for _, authRole := range authRoles {
		// print permissions for every role
		permissions, err := authManager.RetrieveRolePermissions(context.TODO(), authRole.RoleId)
		if err != nil {
			return nil, err
		}
		for _, perm := range permissions {
			if perm.Principal == userPrincipal || groupPrincipalMap[perm.Principal] {
				for _, priv := range authRole.Privilege {
					privileges[priv] = true
				}
			}
		}
	}
	return privileges, nil
}

func (v *VSphereCloudDriver) GetVSphereResourcePools(ctx context.Context, datacenter string, cluster string) (resourcePools []string, err error) {
	finder, dc, err := v.getFinderWithDatacenter(ctx, datacenter)
	if err != nil {
		return nil, err
	}

	searchPath := fmt.Sprintf("/%s/host/%s/Resources/*", dc, cluster)
	pools, govErr := finder.ResourcePoolList(ctx, searchPath)
	if govErr != nil {
		//ignore NotFoundError, to allow selection of "Resources" as the default option for rs pool
		if _, ok := govErr.(*find.NotFoundError); !ok {
			return nil, fmt.Errorf("failed to fetch vSphere resource pools. datacenter: %s, code: %d", datacenter, http.StatusBadRequest)
		}
	}

	for i := 0; i < len(pools); i++ {
		pool := pools[i]
		prefix := fmt.Sprintf("/%s/host/%s/Resources/", dc, cluster)
		poolPath := strings.TrimPrefix(pool.InventoryPath, prefix)
		resourcePools = append(resourcePools, poolPath)
		childPoolSearchPath := fmt.Sprintf("/%s/host/%s/Resources/%s/*", dc, cluster, poolPath)
		childPools, err := finder.ResourcePoolList(ctx, childPoolSearchPath)
		if err == nil {
			pools = append(pools, childPools...)
		}
	}

	sort.Strings(resourcePools)
	return resourcePools, nil
}

func (v *VSphereCloudDriver) getClusterDatastores(ctx context.Context, finder *find.Finder, datacenter string, cluster mo.ClusterComputeResource) (datastores []string, err error) {
	dsMobjRefs := cluster.Datastore

	for i := range dsMobjRefs {
		inventoryPath := ""
		dsObjRef, err := finder.ObjectReference(ctx, dsMobjRefs[i])
		if err != nil {
			return nil, fmt.Errorf("error: %s, code: %d", err.Error(), http.StatusBadRequest)
		}
		if dsObjRef != nil {
			ref := dsObjRef
			switch ref.(type) {
			case *object.Datastore:
				n := dsObjRef.(*object.Datastore)
				inventoryPath = n.InventoryPath
			default:
				continue
			}

			if inventoryPath != "" {
				prefix := fmt.Sprintf("/%s/datastore/", datacenter)
				datastore := strings.TrimPrefix(inventoryPath, prefix)
				datastores = append(datastores, datastore)
			}
		}
	}

	sort.Strings(datastores)
	return datastores, nil
}

func (v *VSphereCloudDriver) getClusterComputeResources(ctx context.Context, finder *find.Finder) ([]*object.ClusterComputeResource, error) {
	ccrs, err := finder.ClusterComputeResourceList(ctx, "*")
	if err != nil {
		return nil, fmt.Errorf("failed to get compute cluster resources: %s", err.Error())
	}
	return ccrs, nil
}

type HostDateInfo struct {
	HostName   string
	NtpServers []string
	types.HostDateTimeInfo
	Service       *types.HostService
	Current       *time.Time
	ClientStatus  string
	ServiceStatus string
}

func (info *HostDateInfo) servers() []string {
	return info.NtpConfig.Server
}

func (v *VSphereCloudDriver) ValidateHostNTPSettings(ctx context.Context, finder *find.Finder, datacenter, clusterName string, hosts []string) (bool, []string, error) {
	var hostsDateInfo []HostDateInfo
	var failures []string

	for _, host := range hosts {
		_, hostObj, err := v.GetHostIfExists(ctx, finder, datacenter, clusterName, host)
		if err != nil {
			return false, nil, err
		}

		s, err := hostObj.ConfigManager().DateTimeSystem(ctx)
		if err != nil {
			return false, nil, err
		}

		var hs mo.HostDateTimeSystem
		if err = s.Properties(ctx, s.Reference(), nil, &hs); err != nil {
			return false, nil, err
		}

		ss, err := hostObj.ConfigManager().ServiceSystem(ctx)
		if err != nil {
			return false, nil, err
		}

		services, err := ss.Service(ctx)
		if err != nil {
			return false, nil, err
		}

		res := &HostDateInfo{HostDateTimeInfo: hs.DateTimeInfo}

		for i, service := range services {
			if service.Key == "ntpd" {
				res.Service = &services[i]
				break
			}
		}

		if res.Service == nil {
			failures = append(failures, fmt.Sprintf("Host: %s has no NTP service operating on it", host))
			return false, failures, fmt.Errorf("host: %s has no NTP service operating on it", host)
		}

		res.Current, err = s.Query(ctx)
		if err != nil {
			return false, nil, err
		}

		res.ClientStatus = service.Policy(*res.Service)
		res.ServiceStatus = service.Status(*res.Service)
		res.HostName = host
		res.NtpServers = res.servers()

		hostsDateInfo = append(hostsDateInfo, *res)
	}

	for _, dateInfo := range hostsDateInfo {
		if dateInfo.ClientStatus != "Enabled" {
			failureMsg := fmt.Sprintf("NTP client status is disabled or unknown for host: %s", dateInfo.HostName)
			failures = append(failures, failureMsg)
		}

		if dateInfo.ServiceStatus != "Running" {
			failureMsg := fmt.Sprintf("NTP service status is stopped or unknown for host: %s", dateInfo.HostName)
			failures = append(failures, failureMsg)
		}
	}

	err := validateHostNTPServers(hostsDateInfo)
	if err != nil {
		failures = append(failures, fmt.Sprintf("%s", err.Error()))
	}

	if len(failures) > 0 {
		return false, failures, err
	}

	return true, failures, nil
}

func validateHostNTPServers(hostsDateInfo []HostDateInfo) error {
	var intersectionList []string
	for i := 0; i < len(hostsDateInfo)-1; i++ {
		if intersectionList == nil {
			intersectionList = intersection(hostsDateInfo[i].NtpServers, hostsDateInfo[i+1].NtpServers)
		} else {
			intersectionList = intersection(intersectionList, hostsDateInfo[i+1].NtpServers)
		}

		if intersectionList == nil {
			return fmt.Errorf("some of the hosts has differently configured NTP servers")
		}
	}

	return nil
}

func intersection(listA []string, listB []string) []string {
	var intersect []string
	for _, element := range listA {
		if slices.Contains(listB, element) {
			intersect = append(intersect, element)
		}
	}

	if len(intersect) == 0 {
		return nil
	}
	return intersect
}

func getUserPrincipalFromPrincipalID(id ssoadmintypes.PrincipalId) string {
	return fmt.Sprintf("%s\\%s", strings.ToUpper(id.Domain), id.Name)
}

func getUserPrincipalFromUsername(username string) string {
	splitStr := strings.Split(username, "@")
	return fmt.Sprintf("%s\\%s", strings.ToUpper(splitStr[1]), splitStr[0])
}
