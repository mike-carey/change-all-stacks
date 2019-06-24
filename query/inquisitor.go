package query

import (
	"io"

	"github.com/cloudfoundry-community/go-cfclient"
)

type Inquisitor interface {
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
	logger io.Writer

	appService *AppService

	spaceService *SpaceService

	stackService *StackService

	orgService *OrgService

	buildpackService *BuildpackService
}

func NewInquisitor(client CFClient, logger io.Writer) Inquisitor {
	return &inquisitor{
		client: client,
		logger: logger,

		appService: NewAppService(client, logger),

		spaceService: NewSpaceService(client, logger),

		stackService: NewStackService(client, logger),

		orgService: NewOrgService(client, logger),

		buildpackService: NewBuildpackService(client, logger),
	}
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
