package tags

import (
	"github.com/go-logr/logr"
	log "github.com/sirupsen/logrus"
	"github.com/spectrocloud-labs/valid8or-plugin-vsphere/api/v1alpha1"
	tags "github.com/spectrocloud-labs/valid8or-plugin-vsphere/internal/validators/tags"
	"github.com/spectrocloud-labs/valid8or-plugin-vsphere/internal/vcsim"
	"github.com/spectrocloud-labs/valid8or-plugin-vsphere/internal/vsphere"
	"github.com/spectrocloud-labs/valid8or-plugin-vsphere/tests/utils/test"
	"github.com/spectrocloud-labs/valid8or/pkg/types"
	"github.com/vmware/govmomi/find"
	_ "github.com/vmware/govmomi/vapi/simulator"
	vtags "github.com/vmware/govmomi/vapi/tags"
	"github.com/vmware/govmomi/vim25/mo"
)

var fakeThumbprint = "A3:B5:9E:5F:E8:84:EE:84:34:D9:8E:EF:85:8E:3F:B6:62:AC:10:85"
var categories = []vtags.Category{
	{
		ID:              "urn:vmomi:InventoryServiceCategory:552dfe88-38ab-4c76-8791-14a2156a5f3f:GLOBAL",
		Name:            "k8s-region",
		Description:     "",
		Cardinality:     "SINGLE",
		AssociableTypes: []string{"Datacenter", "Folder"},
		UsedBy:          []string{},
	},
	{
		ID:              "urn:vmomi:InventoryServiceCategory:167242af-7e93-41ed-8704-52791115e1a8:GLOBAL",
		Name:            "k8s-zone",
		Description:     "",
		Cardinality:     "SINGLE",
		AssociableTypes: []string{"Datacenter", "ClusterComputeResource", "HostSystem", "Folder"},
		UsedBy:          []string{},
	},
	{
		ID:              "urn:vmomi:InventoryServiceCategory:4adb4e4b-8aee-4beb-8f6c-66d22d768cbc:GLOBAL",
		Name:            "AVICLUSTER_UUID",
		Description:     "",
		Cardinality:     "SINGLE",
		AssociableTypes: []string{"com.vmware.content.library.Item"},
		UsedBy:          []string{},
	},
}
var attachedTags = []vtags.AttachedTags{
	{
		ObjectID: nil,
		TagIDs:   []string{"urn:vmomi:InventoryServiceTag:b4f0bd2c-1d62-4af6-ae41-bef79caf5f21:GLOBAL"},
		Tags: []vtags.Tag{
			{
				ID:          "urn:vmomi:InventoryServiceTag:b4f0bd2c-1d62-4af6-ae41-bef79caf5f21:GLOBAL",
				Description: "",
				Name:        "usdc",
				CategoryID:  "urn:vmomi:InventoryServiceCategory:552dfe88-38ab-4c76-8791-14a2156a5f3f:GLOBAL",
				UsedBy:      nil,
			},
		},
	},
	{
		ObjectID: nil,
		TagIDs:   []string{"urn:vmomi:InventoryServiceTag:e886a5b2-73cd-488e-85be-9c8b1bc740eb:GLOBAL"},
		Tags: []vtags.Tag{
			{
				ID:          "urn:vmomi:InventoryServiceTag:e886a5b2-73cd-488e-85be-9c8b1bc740eb:GLOBAL",
				Description: "",
				Name:        "zone1",
				CategoryID:  "urn:vmomi:InventoryServiceCategory:167242af-7e93-41ed-8704-52791115e1a8:GLOBAL",
				UsedBy:      nil,
			},
		},
	},
}

func Execute() error {
	testCtx := test.NewTestContext()
	return test.Flow(testCtx).
		Test(NewVMMigrationTest("vm-migration-integration-test")).
		TearDown().Audit()
}

type VMMigrationTest struct {
	*test.BaseTest
	log *log.Entry
}

func NewVMMigrationTest(description string) *VMMigrationTest {
	return &VMMigrationTest{
		log:      log.WithField("test", "role-privilege-integration-test"),
		BaseTest: test.NewBaseTest("vsphere-plugin", description, nil),
	}
}

func (t *VMMigrationTest) Execute(ctx *test.TestContext) (tr *test.TestResult) {
	t.log.Printf("Executing %s and %s", t.GetName(), t.GetDescription())
	//if tr := t.PreRequisite(ctx); tr.IsFailed() {
	//	return tr
	//}

	if result := t.testGenerateManifestsInteractive(ctx); result.IsFailed() {
		return result
	}

	return test.Success()
}

func (t *VMMigrationTest) testGenerateManifestsInteractive(ctx *test.TestContext) (tr *test.TestResult) {
	vcSim := vcsim.NewVCSim("admin@vsphere.local")
	vcSim.Start()
	vsphereCloudAccount := vcSim.GetTestVsphereAccount()

	vsphereCloudDriver, err := vsphere.NewVSphereDriver(vsphereCloudAccount.VcenterServer, vsphereCloudAccount.Username, vsphereCloudAccount.Password, "DC0")
	if err != nil {
		return tr
	}

	tm := vtags.NewManager(vsphereCloudDriver.RestClient)
	finder := find.NewFinder(vsphereCloudDriver.Client.Client)

	var log logr.Logger
	tagService := tags.NewTagsValidationService(log)

	rule := v1alpha1.RegionZoneValidationRule{
		RegionCategoryName: "k8s-region",
		ZoneCategoryName:   "k8s-zone",
		Datacenter:         "DC0",
		Clusters:           []string{"DC0_C0"},
	}

	testCases := []struct {
		name             string
		expectedErr      bool
		validationResult types.ValidationResult
		categories       []vtags.Category
		attachedTags     []vtags.AttachedTags
	}{
		{
			name:             "DataCenter and Cluster tags Exist",
			expectedErr:      false,
			validationResult: types.ValidationResult{},
			categories:       categories,
			attachedTags:     attachedTags,
		},
		{
			name:             "Empty categories and attachedTags",
			expectedErr:      true,
			validationResult: types.ValidationResult{},
			categories:       []vtags.Category{},
			attachedTags:     []vtags.AttachedTags{},
		},
	}
	for _, tc := range testCases {
		tags.GetCategories = func(manager *vtags.Manager) ([]vtags.Category, error) {
			return tc.categories, nil
		}
		tags.GetAttachedTagsOnObjects = func(tagsManager *vtags.Manager, refs []mo.Reference) ([]vtags.AttachedTags, error) {
			return tc.attachedTags, nil
		}
		vr, err := tagService.ReconcileRegionZoneTagRules(tm, finder, rule)
		if tc.expectedErr && err == nil {
			tr.Failed()
		} else if vr.Condition.Status != "True" {
			tr.Failed()
		}
	}

	return tr
}

//func (t *VMMigrationTest) testDeployFromConfigFile() (tr *test.TestResult) {
//	confFileCmd, buffer := common.InitCmd([]string{
//		"vmo", "migrate-vm", "-f", t.filePath("vm.yaml"),
//	})
//	return common.ExecCLI(confFileCmd, buffer, t.log)
//}

//func (t *VMMigrationTest) PreRequisite(ctx *test.TestContext) (tr *test.TestResult) {
//	t.log.Printf("Executing ExecuteRequisite for %s and %s", t.GetName(), t.GetDescription())
//
//
//
//	// setup vCenter simulator
//	vcSimulator := vcsim.NewVCSim("admin2@vsphere.local")
//	vcSimulator.Start()
//	ctx.Put("vcsim", vcSimulator)
//
//	return test.Success()
//}

func (t *VMMigrationTest) TearDown(ctx *test.TestContext) {
	t.log.Printf("Executing TearDown for %s and %s ", t.GetName(), t.GetDescription())

	//if err := common.TearDownFun()(ctx); err != nil {
	//	t.log.Errorf("Failed to run teardown fun: %v", err)
	//}

	// shut down vCenter simulator
	vcSimulator := ctx.Get("vcsim")
	vcSimulator.(*vcsim.VCSimulator).Shutdown()
}

//func (t *VMMigrationTest) filePath(file string) string {
//	return fmt.Sprintf("%s/%s/%s", file_utils.VMTestCasesPath(), "data", file)
//}
