package cf_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/mike-carey/change-all-stacks/cf"

	fakes "github.com/mike-carey/change-all-stacks/cf/fakes"
)

var _ = Describe("Executor", func() {

	var (
		fakeCommand *fakes.FakeCFCommand
		executor Executor
	)

	BeforeEach(func() {
		fakeCommand = new(fakes.FakeCFCommand)
	})

	for active, cmd := range map[string]string{"enabled": "String", "disabled": "Execute"} {
		Context(fmt.Sprintf("Dry Run %s", active), func() {
				BeforeEach(func() {
					executor = NewExecutor(fakeCommand, active == "enabled")
				})

				It(fmt.Sprintf("Should run the %s method for the api command", cmd), func() {
					By("With ssl validation")
					err := executor.Api("api-address", false)
					Expect(err).To(BeNil())

					inv := fakeCommand.Invocations()[cmd]
					Expect(len(inv)).To(Equal(1))
					Expect(inv[0][0]).To(Equal([]string{"api", "api-address", ""}))

					By("Without ssl validation")
					err = executor.Api("api-address", true)
					Expect(err).To(BeNil())

					inv = fakeCommand.Invocations()[cmd]
					Expect(len(inv)).To(Equal(2))
					Expect(inv[1][0]).To(Equal([]string{"api", "api-address", "--skip-ssl-validation"}))
				})

				It(fmt.Sprintf("Should run the %s method for the auth command", cmd), func() {
					err := executor.Auth("username", "password")
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
					err := executor.Target("org", "space")
					Expect(err).To(BeNil())

					inv := fakeCommand.Invocations()[cmd]
					Expect(len(inv)).To(Equal(1))

					Expect(inv[0][0]).To(Equal([]string{"target", "-o", "org", "-s", "space"}))
				})

				It(fmt.Sprintf("Should run the %s method for the install-plugin command", cmd), func() {
					err := executor.InstallPlugin("plugin")
					Expect(err).To(BeNil())

					inv := fakeCommand.Invocations()[cmd]
					Expect(len(inv)).To(Equal(1))

					Expect(inv[0][0]).To(Equal([]string{"install-plugin", "plugin", "-f"}))
				})

				It(fmt.Sprintf("Should run the %s method for the change-stack command", cmd), func() {
					err := executor.ChangeStack("app", "stack")
					Expect(err).To(BeNil())

					inv := fakeCommand.Invocations()[cmd]
					Expect(len(inv)).To(Equal(1))

					Expect(inv[0][0]).To(Equal([]string{"change-stack", "app", "stack"}))
				})
		})
	}

})
