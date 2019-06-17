package change_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/mike-carey/change-all-stacks/change"

	fakes "github.com/mike-carey/change-all-stacks/change/fakes"
)

var _ = Describe("Handle", func() {

	var (
		fakeChanger *fakes.FakeChanger
		fakeWriter *fakes.FakeWriter
		handler Handler
	)

	BeforeEach(func() {
		fakeChanger = new(fakes.FakeChanger)
		fakeWriter = new(fakes.FakeWriter)
		handler = NewHandler(fakeChanger, fakeWriter)
	})

	It("Should handle changing the stack", func() {
		org := "org"
		space := "space"
		app := "app"
		stack := "stack"

		handler.Handle(org, space, app, stack)

		i := fakeChanger.Invocations()

		Expect(len(i["ChangeStack"])).To(Equal(1))
		Expect(i["ChangeStack"][0][0]).To(Equal(app))
		Expect(i["ChangeStack"][0][1]).To(Equal(stack))
	})

	It("Should handle a dry run", func() {
		org := "org"
		space := "space"
		app := "app"
		stack := "stack"

		handler.HandleDryRun(org, space, app, stack)

		i := fakeWriter.Invocations()

		Expect(len(i["Write"])).To(Equal(1))
	})


})
