package change_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/mike-carey/change-all-stacks/change"

	fakes "github.com/mike-carey/change-all-stacks/change/fakes"
)

var _ = Describe("Runner", func() {

	var (
		fakeCommand *fakes.FakeCFCommand
		fakeLogger *fakes.FakeLogger
		runner Runner
	)

	BeforeEach(func() {
		fakeCommand = new(fakes.FakeCFCommand)
		fakeLogger = new(fakes.FakeLogger)
	})

	for active, cmd := range map[string]string{"enabled": "String", "disabled": "Execute"} {
		Context(fmt.Sprintf("Dry Run %s", active), func() {
				BeforeEach(func() {
					runner = NewRunner(fakeCommand, fakeLogger, active == "enabled")
				})

				It(fmt.Sprintf("Should run the %s method for the api command", cmd), func() {
					By("With ssl validation")
					err := runner.Api("api-address", false)
					Expect(err).To(BeNil())

					inv := fakeCommand.Invocations()[cmd]
					Expect(len(inv)).To(Equal(1))
					Expect(inv[0][0]).To(Equal([]string{"api", "api-address", ""}))

					By("Without ssl validation")
					err = runner.Api("api-address", true)
					Expect(err).To(BeNil())

					inv = fakeCommand.Invocations()[cmd]
					Expect(len(inv)).To(Equal(2))
					Expect(inv[1][0]).To(Equal([]string{"api", "api-address", "--skip-ssl-validation"}))
				})

				It(fmt.Sprintf("Should run the %s method for the auth command", cmd), func() {
					err := runner.Auth("username", "password")
					Expect(err).To(BeNil())

					inv := fakeCommand.Invocations()[cmd]
					Expect(len(inv)).To(Equal(1))

					expectedPassword := "password"
					if cmd == "String" {
						expectedPassword = RedactedString
					}

					Expect(inv[0][0]).To(Equal([]string{"auth", "username", expectedPassword}))
				})

				It(fmt.Sprintf("Should run the %s method for the target command", cmd), func() {
					err := runner.Target("org", "space")
					Expect(err).To(BeNil())

					inv := fakeCommand.Invocations()[cmd]
					Expect(len(inv)).To(Equal(1))

					Expect(inv[0][0]).To(Equal([]string{"target", "-o", "org", "-s", "space"}))
				})

				It(fmt.Sprintf("Should run the %s method for the install-plugin command", cmd), func() {
					err := runner.InstallPlugin("plugin")
					Expect(err).To(BeNil())

					inv := fakeCommand.Invocations()[cmd]
					Expect(len(inv)).To(Equal(1))

					Expect(inv[0][0]).To(Equal([]string{"install-plugin", "plugin", "-f"}))
				})

				It(fmt.Sprintf("Should run the %s method for the change-stack command", cmd), func() {
					err := runner.ChangeStack("app", "stack")
					Expect(err).To(BeNil())

					inv := fakeCommand.Invocations()[cmd]
					Expect(len(inv)).To(Equal(1))

					Expect(inv[0][0]).To(Equal([]string{"change-stack", "app", "stack"}))
				})
		})
	}

})
