package cf_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/mike-carey/change-all-stacks/cf"

	cfclient "github.com/cloudfoundry-community/go-cfclient"

	fakes "github.com/mike-carey/change-all-stacks/cf/fakes"
)

var _ = Describe("Worker", func() {

	var (
		apiAddress = "ApiAddress"
		username = "Username"
		password = "Password"
		skipSslValidation = true
		pluginPath = "PluginPath"
		org = "org"
		space = "space"
		stackName = "stack"

		config *cfclient.Config
		apps []cfclient.App
		fakeRunner *fakes.FakeRunner
		worker Worker
	)

	BeforeEach(func() {
		config = &cfclient.Config{
			ApiAddress: apiAddress,
			Username: username,
			Password: password,
			SkipSslValidation: skipSslValidation,
		}

		fakeRunner = new(fakes.FakeRunner)

		apps = []cfclient.App{
			cfclient.App{
				Guid: "app-1",
				Name: "app-1",
			},
			cfclient.App{
				Guid: "app-2",
				Name: "app-2",
			},
		}

		worker = NewWorker(fakeRunner, config, pluginPath, org, space)
	})

	It("Should Setup the suite, and then run on each app", func() {
		err := worker.Work(apps, stackName)

		i := fakeRunner.Invocations()

		Expect(err).To(BeNil())
		Expect(i).To(HaveKey("Setup"))
		Expect(i["Setup"]).To(HaveLen(1))
		Expect(i["Setup"][0][0]).To(Equal(config))
		Expect(i["Setup"][0][1]).To(Equal(pluginPath))
		Expect(i["Setup"][0][2]).To(Equal(org))
		Expect(i["Setup"][0][3]).To(Equal(space))

		Expect(i).To(HaveKey("Run"))
		Expect(i["Run"]).To(HaveLen(2))

		first := 0
		second := 1
		if i["Run"][0][0] == apps[1].Name {
			first = 1
			second = 0
		}
		// Async, any order
		Expect(i["Run"][first][0]).To(Equal(apps[0].Name))
		Expect(i["Run"][first][1]).To(Equal(stackName))
		Expect(i["Run"][second][0]).To(Equal(apps[1].Name))
		Expect(i["Run"][second][1]).To(Equal(stackName))
	})

})
