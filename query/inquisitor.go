package query

import (
	"github.com/cloudfoundry-community/go-cfclient"
)

//go:generate counterfeiter -o fakes/fake_inquisitor.go Inquisitor
type Inquisitor interface {
	Client() CFClient

	GetAllApps() ([]cfclient.App, error)
	GetAppByGuid(appGuid string) (cfclient.App, error)
	GetAppByName(appName string) (cfclient.App, error)

	GetAllSpaces() ([]cfclient.Space, error)
	GetSpaceByGuid(spaceGuid string) (cfclient.Space, error)
	GetSpaceByName(spaceName string) (cfclient.Space, error)

	GetAllStacks() ([]cfclient.Stack, error)
	GetStackByGuid(stackGuid string) (cfclient.Stack, error)
	GetStackByName(stackName string) (cfclient.Stack, error)

	GetAllOrgs() ([]cfclient.Org, error)
	GetOrgByGuid(orgGuid string) (cfclient.Org, error)
	GetOrgByName(orgName string) (cfclient.Org, error)

	GetAllBuildpacks() ([]cfclient.Buildpack, error)
	GetBuildpackByGuid(buildpackGuid string) (cfclient.Buildpack, error)
	GetBuildpackByName(buildpackName string) (cfclient.Buildpack, error)
}

type inquisitor struct {
	client CFClient

	appService *AppService

	spaceService *SpaceService

	stackService *StackService

	orgService *OrgService

	buildpackService *BuildpackService
}

func NewInquisitor(client CFClient) Inquisitor {
	return &inquisitor{
		client: client,

		appService: NewAppService(client),

		spaceService: NewSpaceService(client),

		stackService: NewStackService(client),

		orgService: NewOrgService(client),

		buildpackService: NewBuildpackService(client),
	}
}

func (i *inquisitor) Client() CFClient {
	return i.client
}

func (i *inquisitor) getAppService() *AppService {
	return i.appService
}

func (i *inquisitor) getSpaceService() *SpaceService {
	return i.spaceService
}

func (i *inquisitor) getStackService() *StackService {
	return i.stackService
}

func (i *inquisitor) getOrgService() *OrgService {
	return i.orgService
}

func (i *inquisitor) getBuildpackService() *BuildpackService {
	return i.buildpackService
}
