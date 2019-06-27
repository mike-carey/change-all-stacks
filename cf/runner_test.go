package cf_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/mike-carey/change-all-stacks/cf"

	cfclient "github.com/cloudfoundry-community/go-cfclient"

	fakes "github.com/mike-carey/change-all-stacks/cf/fakes"
)

var _ = Describe("Runner", func() {

	var (
		apiAddress = "ApiAddress"
		username = "Username"
		password = "Password"
		skipSslValidation = true
		pluginPath = "PluginPath"
		org = "org"
		space = "space"
		appName = "app"
		stackName = "stack"

		config *cfclient.Config
		fakeExecutor *fakes.FakeExecutor
		runner Runner
	)

	BeforeEach(func() {
		config = &cfclient.Config{
			ApiAddress: apiAddress,
			Username: username,
			Password: password,
			SkipSslValidation: skipSslValidation,
		}

		fakeExecutor = new(fakes.FakeExecutor)

		runner = NewRunner(fakeExecutor)
	})

	It("Should Setup the suite", func() {
		err := runner.Setup(config, pluginPath, org, space)

		i := fakeExecutor.Invocations()

		Expect(err).To(BeNil())

		Expect(i).To(HaveKey("Api"))
		Expect(i["Api"]).To(HaveLen(1))
		Expect(i["Api"][0][0]).To(Equal(apiAddress))
		Expect(i["Api"][0][1]).To(Equal(skipSslValidation))

		Expect(i).To(HaveKey("Auth"))
		Expect(i["Auth"]).To(HaveLen(1))
		Expect(i["Auth"][0][0]).To(Equal(username))
		Expect(i["Auth"][0][1]).To(Equal(password))

		Expect(i).To(HaveKey("Target"))
		Expect(i["Target"]).To(HaveLen(1))
		Expect(i["Target"][0][0]).To(Equal(org))
		Expect(i["Target"][0][1]).To(Equal(space))

		Expect(i).To(HaveKey("InstallPlugin"))
		Expect(i["InstallPlugin"]).To(HaveLen(1))
		Expect(i["InstallPlugin"][0][0]).To(Equal(pluginPath))
	})

	It("Should Run the suite", func() {
		// We need to setup first
		err := runner.Setup(config, pluginPath, org, space)

		Expect(err).To(BeNil())
		Expect(fakeExecutor.Invocations()).NotTo(HaveKey("ChangeStack"))

		err = runner.Run(appName, stackName)

		i := fakeExecutor.Invocations()

		Expect(err).To(BeNil())

		Expect(i).To(HaveKey("ChangeStack"))
		Expect(i["ChangeStack"]).To(HaveLen(1))
		Expect(i["ChangeStack"][0][0]).To(Equal(appName))
		Expect(i["ChangeStack"][0][1]).To(Equal(stackName))
	})

	It("Should throw an error if not setup and run is called", func() {
		err := runner.Run(appName, stackName)

		Expect(err).NotTo(BeNil())
	})

})
