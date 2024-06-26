package vsphere

import (
	"context"

	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/view"
	"github.com/vmware/govmomi/vim25/mo"
)

func (v *VSphereCloudDriver) GetVAppIfExists(ctx context.Context, finder *find.Finder, datacenter, vAppName string) (bool, *object.VirtualApp, error) {
	vapp, err := finder.VirtualApp(ctx, vAppName)
	if err != nil {
		return false, nil, err
	}
	return true, vapp, nil
}

func (v *VSphereCloudDriver) GetVapps(ctx context.Context) ([]mo.VirtualApp, error) {
	m := view.NewManager(v.Client.Client)

	containerView, err := m.CreateContainerView(ctx, v.Client.Client.ServiceContent.RootFolder, []string{"VirtualApp"}, true)
	if err != nil {
		return nil, err
	}
	var vApps []mo.VirtualApp
	err = containerView.Retrieve(ctx, []string{"VirtualApp"}, nil, &vApps)
	if err != nil {
		return nil, err
	}

	return vApps, nil
}
