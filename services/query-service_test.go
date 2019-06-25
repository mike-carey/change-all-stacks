package services_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/mike-carey/change-all-stacks/services"

	queryFakes "github.com/mike-carey/change-all-stacks/query/fakes"

	cfclient "github.com/cloudfoundry-community/go-cfclient"

)

var _ = Describe("QueryService", func() {

	var (
		fakeInquisitor *queryFakes.FakeInquisitor
		apps []cfclient.App
		orgs []cfclient.Org
		spaces []cfclient.Space

		queryService QueryService
	)

	BeforeEach(func() {
		fakeInquisitor = new(queryFakes.FakeInquisitor)
		queryService = NewQueryService(fakeInquisitor)
		apps = []cfclient.App{
			cfclient.App{
				Guid: "app-1",
				Name: "app-1",
				SpaceGuid: "space-1",
			},
			cfclient.App{
				Guid: "app-2",
				Name: "app-2",
				SpaceGuid: "space-1",
			},
			cfclient.App{
				Guid: "app-3",
				Name: "app-3",
				SpaceGuid: "space-2",
			},
			cfclient.App{
				Guid: "app-4",
				Name: "app-4",
				SpaceGuid: "space-3",
			},
			cfclient.App{
				Guid: "app-5",
				Name: "app-5",
				SpaceGuid: "space-4",
			},
		}

		spaces = []cfclient.Space{
			cfclient.Space{
				Guid: "space-1",
				Name: "space-1",
				OrganizationGuid: "org-1",
			},
			cfclient.Space{
				Guid: "space-2",
				Name: "space-2",
				OrganizationGuid: "org-1",
			},
			cfclient.Space{
				Guid: "space-3",
				Name: "space-3",
				OrganizationGuid: "org-2",
			},
			cfclient.Space{
				Guid: "space-4",
				Name: "space-4",
				OrganizationGuid: "org-3",
			},
		}

		orgs = []cfclient.Org{
			cfclient.Org{
				Guid: "org-1",
				Name: "org-1",
			},
			cfclient.Org{
				Guid: "org-2",
				Name: "org-2",
			},
			cfclient.Org{
				Guid: "org-3",
				Name: "org-3",
			},
		}

		fakeInquisitor.GetAllAppsReturns(apps, nil)
		fakeInquisitor.GetAllOrgsReturns(orgs, nil)
		fakeInquisitor.GetAllSpacesReturns(spaces, nil)

		fakeInquisitor.GetAppByGuidStub = func(appGuid string) (cfclient.App, error) {
			for _, app := range apps {
				if app.Guid == appGuid {
					return app, nil
				}
			}

			return cfclient.App{}, fmt.Errorf("Could not find app with guid: %s", appGuid)
		}
		fakeInquisitor.GetOrgByGuidStub = func(orgGuid string) (cfclient.Org, error) {
			for _, org := range orgs {
				if org.Guid == orgGuid {
					return org, nil
				}
			}

			return cfclient.Org{}, fmt.Errorf("Could not find org with guid: %s", orgGuid)
		}
		fakeInquisitor.GetSpaceByGuidStub = func(spaceGuid string) (cfclient.Space, error) {
			for _, space := range spaces {
				if space.Guid == spaceGuid {
					return space, nil
				}
			}

			return cfclient.Space{}, fmt.Errorf("Could not find space with guid: %s", spaceGuid)
		}

		fakeInquisitor.GetAppByNameStub = func(appName string) (cfclient.App, error) {
			for _, app := range apps {
				if app.Name == appName {
					return app, nil
				}
			}

			return cfclient.App{}, fmt.Errorf("Could not find app with name: %s", appName)
		}
		fakeInquisitor.GetOrgByNameStub = func(orgName string) (cfclient.Org, error) {
			for _, org := range orgs {
				if org.Name == orgName {
					return org, nil
				}
			}

			return cfclient.Org{}, fmt.Errorf("Could not find org with guid: %s", orgName)
		}
		fakeInquisitor.GetSpaceByNameStub = func(spaceName string) (cfclient.Space, error) {
			for _, space := range spaces {
				if space.Name == spaceName {
					return space, nil
				}
			}

			return cfclient.Space{}, fmt.Errorf("Could not find space with name: %s", spaceName)
		}
	})

	Describe("GetAllAppsWithinOrgs", func() {
		It("Should find all apps", func() {
			a, e := queryService.GetAllAppsWithinOrgs()

			Expect(e).To(BeNil())
			Expect(a).Should(ConsistOf(apps))
		})

		It("Should filter apps by orgs", func() {
			a, e := queryService.GetAllAppsWithinOrgs(orgs[0].Name, orgs[1].Name)

			Expect(e).To(BeNil())
			Expect(a).Should(ConsistOf(apps[0:4]))
		})
	})

})
