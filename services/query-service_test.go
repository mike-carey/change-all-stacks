package services_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/mike-carey/change-all-stacks/services"

	"github.com/mike-carey/change-all-stacks/query"

	queryFakes "github.com/mike-carey/change-all-stacks/query/fakes"

	cfclient "github.com/cloudfoundry-community/go-cfclient"

)

var _ = Describe("QueryService", func() {

	var (
		fakeInquisitor *queryFakes.FakeInquisitor
		apps []cfclient.App
		orgs []cfclient.Org
		spaces []cfclient.Space
		stacks []cfclient.Stack
		buildpacks []cfclient.Buildpack

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
				StackGuid: "stack-1",
				Buildpack: "buildpack-1",
			},
			cfclient.App{
				Guid: "app-2",
				Name: "app-2",
				SpaceGuid: "space-1",
				StackGuid: "stack-2",
				Buildpack: "buildpack-1",
			},
			cfclient.App{
				Guid: "app-3",
				Name: "app-3",
				SpaceGuid: "space-2",
				StackGuid: "stack-2",
				DetectedBuildpack: "buildpack-1",
			},
			cfclient.App{
				Guid: "app-4",
				Name: "app-4",
				SpaceGuid: "space-3",
				StackGuid: "stack-3",
				DetectedBuildpack: "buildpack-2",
			},
			cfclient.App{
				Guid: "app-5",
				Name: "app-5",
				SpaceGuid: "space-4",
				StackGuid: "stack-1",
				DetectedBuildpackGuid: "buildpack-3",
			},
			cfclient.App{
				Guid: "app-6",
				Name: "app-6",
				SpaceGuid: "space-4",
				StackGuid: "stack-2",
				Buildpack: "buildpack-6",
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

		stacks = []cfclient.Stack{
			cfclient.Stack{
				Guid: "stack-1",
				Name: "stack-1",
			},
			cfclient.Stack{
				Guid: "stack-2",
				Name: "stack-2",
			},
			cfclient.Stack{
				Guid: "stack-3",
				Name: "stack-3",
			},
		}

		buildpacks = []cfclient.Buildpack{
				cfclient.Buildpack{
					Guid: "buildpack-1",
					Name: "buildpack-1",
					Stack: "stack-1",
				},
				cfclient.Buildpack{
					Guid: "buildpack-2",
					Name: "buildpack-2",
					Stack: "stack-1",
				},
				cfclient.Buildpack{
					Guid: "buildpack-3",
					Name: "buildpack-3",
					Stack: "stack-2",
				},
				cfclient.Buildpack{
					Guid: "buildpack-4",
					Name: "buildpack-4",
					Stack: "stack-2",
				},
				cfclient.Buildpack{
					Guid: "buildpack-5",
					Name: "buildpack-5",
					Stack: "stack-3",
				},
		}

		fakeInquisitor.GetAllAppsReturns(apps, nil)
		fakeInquisitor.GetAllOrgsReturns(orgs, nil)
		fakeInquisitor.GetAllSpacesReturns(spaces, nil)
		fakeInquisitor.GetAllStacksReturns(stacks, nil)
		fakeInquisitor.GetAllBuildpacksReturns(buildpacks, nil)

		fakeInquisitor.GetAppByGuidStub = func(appGuid string) (cfclient.App, error) {
			for _, app := range apps {
				if app.Guid == appGuid {
					return app, nil
				}
			}

			item := cfclient.App{}
			return item, query.NewNotFoundError(item, "guid", appGuid)
		}
		fakeInquisitor.GetOrgByGuidStub = func(orgGuid string) (cfclient.Org, error) {
			for _, org := range orgs {
				if org.Guid == orgGuid {
					return org, nil
				}
			}

			item := cfclient.Org{}
			return item, query.NewNotFoundError(item, "guid", orgGuid)
		}
		fakeInquisitor.GetSpaceByGuidStub = func(spaceGuid string) (cfclient.Space, error) {
			for _, space := range spaces {
				if space.Guid == spaceGuid {
					return space, nil
				}
			}

			item := cfclient.Space{}
			return item, query.NewNotFoundError(item, "guid", spaceGuid)
		}
		fakeInquisitor.GetStackByGuidStub = func(stackGuid string) (cfclient.Stack, error) {
			for _, stack := range stacks {
				if stack.Guid == stackGuid {
					return stack, nil
				}
			}

			item := cfclient.Stack{}
			return item, query.NewNotFoundError(item, "guid", stackGuid)
		}
		fakeInquisitor.GetBuildpackByGuidStub = func(buildpackGuid string) (cfclient.Buildpack, error) {
			for _, buildpack := range buildpacks {
				if buildpack.Guid == buildpackGuid {
					return buildpack, nil
				}
			}

			item := cfclient.Buildpack{}
			return item, query.NewNotFoundError(item, "guid", buildpackGuid)
		}

		fakeInquisitor.GetAppByNameStub = func(appName string) (cfclient.App, error) {
			for _, app := range apps {
				if app.Name == appName {
					return app, nil
				}
			}

			item := cfclient.App{}
			return item, query.NewNotFoundError(item, "name", appName)
		}
		fakeInquisitor.GetOrgByNameStub = func(orgName string) (cfclient.Org, error) {
			for _, org := range orgs {
				if org.Name == orgName {
					return org, nil
				}
			}

			item := cfclient.Org{}
			return item, query.NewNotFoundError(item, "guid", orgName)
		}
		fakeInquisitor.GetSpaceByNameStub = func(spaceName string) (cfclient.Space, error) {
			for _, space := range spaces {
				if space.Name == spaceName {
					return space, nil
				}
			}

			item := cfclient.Space{}
			return item, query.NewNotFoundError(item, "name", spaceName)
		}
		fakeInquisitor.GetStackByNameStub = func(stackName string) (cfclient.Stack, error) {
			for _, stack := range stacks {
				if stack.Name == stackName {
					return stack, nil
				}
			}

			item := cfclient.Stack{}
			return item, query.NewNotFoundError(item, "name", stackName)
		}
		fakeInquisitor.GetBuildpackByNameStub = func(buildpackName string) (cfclient.Buildpack, error) {
			for _, buildpack := range buildpacks {
				if buildpack.Name == buildpackName {
					return buildpack, nil
				}
			}

			item := cfclient.Buildpack{}
			return item, query.NewNotFoundError(item, "name", buildpackName)
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

	Describe("FilterAppsByStackName", func() {

		It("Should filter out apps without a certain stack", func() {
			a, e := queryService.FilterAppsByStackName(apps, stacks[0].Name)

			Expect(e).To(BeNil())
			Expect(a).Should(ConsistOf(apps[0], apps[4]))
		})

	})

	Describe("FilterAppsByStackName", func() {

		It("Should filter out apps without a certain buildpack", func() {
			a, e := queryService.FilterAppsByBuildpackName(apps, buildpacks[0].Name)

			Expect(e).To(BeNil())
			Expect(a).Should(ConsistOf(apps[0:3]))
		})

		It("Should still find apps with a buildpack that no longer exists", func() {
			a, e := queryService.FilterAppsByBuildpackName(apps, apps[5].Buildpack)

			Expect(e).To(BeNil())
			Expect(a).Should(ConsistOf(apps[5]))
		})

		It("Should throw an error if an app's buildpack cannot be determined", func() {
			apps = append(apps, cfclient.App{
				Guid: "app-injected",
				Name: "app-injected",
			})
			_, e := queryService.FilterAppsByBuildpackName(apps, apps[5].Buildpack)

			Expect(e).NotTo(BeNil())
		})

	})

})
