//

package query

import (
	"github.com/cloudfoundry-community/go-cfclient"
)

// Inquisitor ...
type Inquisitor interface {
	GetAppMap() (AppMap, error)
	GetAppByGuid(guid string) (*cfclient.App, error)
	GetManyAppsByGuid(guids ...string) (AppMap, error)
	GetAllApps() (Apps, error)
	GetOrgMap() (OrgMap, error)
	GetOrgByGuid(guid string) (*cfclient.Org, error)
	GetManyOrgsByGuid(guids ...string) (OrgMap, error)
	GetAllOrgs() (Orgs, error)
	GetServiceBindingMap() (ServiceBindingMap, error)
	GetServiceBindingByGuid(guid string) (*cfclient.ServiceBinding, error)
	GetManyServiceBindingsByGuid(guids ...string) (ServiceBindingMap, error)
	GetAllServiceBindings() (ServiceBindings, error)
	GetServiceInstanceMap() (ServiceInstanceMap, error)
	GetServiceInstanceByGuid(guid string) (*cfclient.ServiceInstance, error)
	GetManyServiceInstancesByGuid(guids ...string) (ServiceInstanceMap, error)
	GetAllServiceInstances() (ServiceInstances, error)
	GetSpaceMap() (SpaceMap, error)
	GetSpaceByGuid(guid string) (*cfclient.Space, error)
	GetManySpacesByGuid(guids ...string) (SpaceMap, error)
	GetAllSpaces() (Spaces, error)
	GetStackMap() (StackMap, error)
	GetStackByGuid(guid string) (*cfclient.Stack, error)
	GetManyStacksByGuid(guids ...string) (StackMap, error)
	GetAllStacks() (Stacks, error)
	GetStackByName(name string) (cfclient.Stack, error)
}
