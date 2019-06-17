package change_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/mike-carey/change-all-stacks/change"

	fakes "github.com/mike-carey/change-all-stacks/change/fakes"

	"code.cloudfoundry.org/cli/plugin"
)

var _ = Describe("Plugin", func() {

	var (
		count int
		params []PluginParams
		errCh chan error
		casPlugin *ChangeAllStacksPlugin
		fakeCliConnection *fakes.FakeCliConnection
	)

	BeforeEach(func() {
		count = 0
		errCh = make(chan error, 0)

		casPlugin = NewChangeAllStacksPlugin(func(_ plugin.CliConnection, p PluginParams) error {
			count += 1
			params = append(params, p)
			return nil
		}, PluginParams{}, errCh)
		fakeCliConnection = new(fakes.FakeCliConnection)
	})

	It("Should call the run func", func() {
		casPlugin.Run(fakeCliConnection, []string{})
		e := <-errCh

		Expect(count).To(Equal(1))
		Expect(e).To(BeNil())
	})

})
