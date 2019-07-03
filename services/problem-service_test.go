package services_test

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"bytes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/mike-carey/change-all-stacks/services"

	data "github.com/mike-carey/change-all-stacks/data"
	cfclient "github.com/cloudfoundry-community/go-cfclient"
	fakes "github.com/mike-carey/change-all-stacks/query/fakes"
)

const (
	CFLinuxFS2 = "cfclinuxfs2"
	CFLinuxFS3 = "cfclinuxfs3"
	Windows2012 = "windows2012R2"
	Windows2016 = "windows2016"

	JavaBuildpack = "java_buildpack"
	JavaBuildpackOffline = "java_buildpack_offline"
	DotnetCoreBuildpack = "dotnet_core_buildpack"
	DotnetCoreBuildpackOldVersion = "dotnet_core_buildpack_old_version"

	DropletNotFoundJson = `{"errors": [{"code": 10010, "title": "CFDropletNotFound", "Detail": "Could not find droplet"}]}`
)

var _ = Describe("ProblemService", func() {

	var (
		cflinuxfs2 cfclient.Stack
		// cflinuxfs3 cfclient.Stack
		// windows2012 cfclient.Stack
		// windows2016 cfclient.Stack

		java_buildpack_2 cfclient.Buildpack
		java_buildpack_3 cfclient.Buildpack
		java_buildpack_offline_2 cfclient.Buildpack
		java_buildpack_offline_3 cfclient.Buildpack
		dotnet_core_buildpack_2 cfclient.Buildpack
		dotnet_core_buildpack_3 cfclient.Buildpack
		dotnet_core_buildpack_old_version cfclient.Buildpack

		fakeInquisitor *fakes.FakeInquisitor
		fakeClient *fakes.FakeCFClient

		problemService ProblemService
	)

	BeforeEach(func() {
		cflinuxfs2 = cfclient.Stack{
			Guid: CFLinuxFS2,
			Name: CFLinuxFS2,
		}

		// cflinuxfs3 = cfclient.Stack{
		// 	Guid: CFLinuxFS3,
		// 	Name: CFLinuxFS3,
		// }
		//
		// windows2012 = cfclient.Stack{
		// 	Guid: Windows2012,
		// 	Name: Windows2012,
		// }
		//
		// windows2016 = cfclient.Stack{
		// 	Guid: Windows2016,
		// 	Name: Windows2016,
		// }

		java_buildpack_2 = cfclient.Buildpack{
			Guid: JavaBuildpack,
			Name: JavaBuildpack,
			Stack: CFLinuxFS2,
		}

		java_buildpack_3 = cfclient.Buildpack{
			Guid: JavaBuildpack,
			Name: JavaBuildpack,
			Stack: CFLinuxFS3,
		}

		java_buildpack_offline_2 = cfclient.Buildpack{
			Guid: JavaBuildpackOffline,
			Name: JavaBuildpackOffline,
			Stack: CFLinuxFS2,
		}

		java_buildpack_offline_3 = cfclient.Buildpack{
			Guid: JavaBuildpackOffline,
			Name: JavaBuildpackOffline,
			Stack: CFLinuxFS3,
		}

		dotnet_core_buildpack_2 = cfclient.Buildpack{
			Guid: DotnetCoreBuildpack,
			Name: DotnetCoreBuildpack,
			Stack: CFLinuxFS2,
		}

		dotnet_core_buildpack_3 = cfclient.Buildpack{
			Guid: DotnetCoreBuildpack,
			Name: DotnetCoreBuildpack,
			Stack: CFLinuxFS3,
		}

		dotnet_core_buildpack_old_version = cfclient.Buildpack{
			Guid: DotnetCoreBuildpackOldVersion,
			Name: DotnetCoreBuildpackOldVersion,
			Stack: CFLinuxFS2,
		}

		buildpacks := []cfclient.Buildpack{
			java_buildpack_2,
			java_buildpack_3,
			java_buildpack_offline_2,
			java_buildpack_offline_3,
			dotnet_core_buildpack_2,
			dotnet_core_buildpack_3,
			dotnet_core_buildpack_old_version,
		}

		fakeClient = new(fakes.FakeCFClient)
		fakeInquisitor = new(fakes.FakeInquisitor)
		fakeInquisitor.GetAllBuildpacksReturns(buildpacks, nil)
		fakeInquisitor.GetBuildpackByGuidStub = func(guid string) (cfclient.Buildpack, error) {
			for _, b := range buildpacks {
				if b.Guid == guid {
					return b, nil
				}
			}

			return cfclient.Buildpack{}, fmt.Errorf("Cannot find buildpack by guid: %s", guid)
		}
		fakeInquisitor.ClientReturns(fakeClient)

		problemService = NewProblemService(fakeInquisitor)
	})

	It("No problems", func () {
		apps := []cfclient.App{
			cfclient.App{
				Guid: "app",
				Name: "app",
				Buildpack: JavaBuildpack,
				StackGuid: cflinuxfs2.Guid,
			},
		}

		req := &cfclient.Request{}
		res := makeResponse(`{"data": {"guid": "droplet-guid"}}`)

		fakeClient.NewRequestReturns(req)
		fakeClient.DoRequestReturns(res, nil)

		pbs, err := problemService.FindProblems("", apps, CFLinuxFS2, CFLinuxFS3)

		Expect(err).To(BeNil())
		Expect(pbs).To(BeEmpty())
	})

	It("Buildpack does not have cfclinuxfs3", func () {
		app := cfclient.App{
			Guid: "app",
			Name: "app",
			Buildpack: DotnetCoreBuildpackOldVersion,
		}
		apps := []cfclient.App{app,}

		expectedReason := data.InvalidBuildpack(DotnetCoreBuildpackOldVersion).GetReason(app)

		req := &cfclient.Request{}
		res := makeResponse(`{"data": {"guid": "droplet-guid"}}`)

		fakeClient.NewRequestReturns(req)
		fakeClient.DoRequestReturns(res, nil)

		pbs, err := problemService.FindProblems("", apps, CFLinuxFS2, CFLinuxFS3)

		Expect(err).To(BeNil())
		Expect(pbs).NotTo(BeEmpty())
		Expect(pbs).To(HaveLen(1))
		Expect(pbs[0].App).To(Equal(app))
		Expect(pbs[0].Reason.GetReason(app)).To(Equal(expectedReason))
		Expect(pbs[0].GetReason()).To(Equal(expectedReason))
	})

	It("current_droplet does not exist", func () {
		app := cfclient.App{
			Guid: "app",
			Name: "app",
			Buildpack: JavaBuildpack,
		}
		apps := []cfclient.App{app,}

		expectedReason := data.MissingDroplet().GetReason(app)

		req := &cfclient.Request{}
		res := makeResponse(DropletNotFoundJson)

		fakeClient.NewRequestReturns(req)
		fakeClient.DoRequestReturns(res, nil)

		pbs, err := problemService.FindProblems("", apps, CFLinuxFS2, CFLinuxFS3)

		Expect(err).To(BeNil())
		Expect(pbs).NotTo(BeEmpty())
		Expect(pbs).To(HaveLen(1))
		Expect(pbs[0].App).To(Equal(app))
		Expect(pbs[0].Reason.GetReason(app)).To(Equal(expectedReason))
		Expect(pbs[0].GetReason()).To(Equal(expectedReason))
	})

})

func makeResponse(str string) *http.Response {
	body := ioutil.NopCloser(bytes.NewReader([]byte(str)))
	res := http.Response {
		Body: body,
	}

	return &res
}
