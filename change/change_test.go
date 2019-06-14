package change_all_stacks_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	cfclient "github.com/cloudfoundry-community/go-cfclient"

	. "github.com/mike-carey/change-all-stacks/change"

	fakes "github.com/mike-carey/change-all-stacks/change/fakes"
)

var _ = Describe("Change", func() {

	var (
		fakeChanger *fakes.FakeChanger
	)

	BeforeEach(func() {
		fakeChanger = new(fakes.FakeChanger)
	})

	It("Should change the stack", func() {
		runner := &Runner{
			Logger: &Logger{Verbose: false,},
		}

		app := cfclient.App{
			Name: "app-1",
		}

		runner.ChangeStackInApp(fakeChanger, app)

		Expect(fakeChanger.ChangeStackCallCount()).To(Equal(1))
	})

})
