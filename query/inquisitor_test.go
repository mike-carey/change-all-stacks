package query_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/mike-carey/change-all-stacks/query"

	fakes "github.com/mike-carey/change-all-stacks/query/fakes"
)

var _ = Describe("Inquisitor", func() {

	var (
		storage    *Storage
		fakeClient *fakes.FakeCFClient
		inquisitor Inquisitor
	)

	BeforeEach(func() {
		storage = GetStorage()
		fakeClient = GetFakeClient(storage)
		inquisitor = NewInquisitor(fakeClient, GinkgoWriter)
	})

	It("Should get all apps", func() {
		items, err := inquisitor.GetAllApps()
		Expect(err).To(BeNil())
		Expect(items).Should(ConsistOf(storage.Apps))
	})

	It("Should get app by guid", func() {
		item, err := inquisitor.GetAppByGuid("App-1")
		Expect(err).To(BeNil())
		Expect(item).Should(Equal(storage.Apps[0]))
	})

	It("Should get app by name", func() {
		item, err := inquisitor.GetAppByName("App-1")
		Expect(err).To(BeNil())
		Expect(item).To(Equal(storage.Apps[0]))
	})

	It("Should get all spaces", func() {
		items, err := inquisitor.GetAllSpaces()
		Expect(err).To(BeNil())
		Expect(items).Should(ConsistOf(storage.Spaces))
	})

	It("Should get space by guid", func() {
		item, err := inquisitor.GetSpaceByGuid("Space-1")
		Expect(err).To(BeNil())
		Expect(item).Should(Equal(storage.Spaces[0]))
	})

	It("Should get space by name", func() {
		item, err := inquisitor.GetSpaceByName("Space-1")
		Expect(err).To(BeNil())
		Expect(item).To(Equal(storage.Spaces[0]))
	})

	It("Should get all stacks", func() {
		items, err := inquisitor.GetAllStacks()
		Expect(err).To(BeNil())
		Expect(items).Should(ConsistOf(storage.Stacks))
	})

	It("Should get stack by guid", func() {
		item, err := inquisitor.GetStackByGuid("Stack-1")
		Expect(err).To(BeNil())
		Expect(item).Should(Equal(storage.Stacks[0]))
	})

	It("Should get stack by name", func() {
		item, err := inquisitor.GetStackByName("Stack-1")
		Expect(err).To(BeNil())
		Expect(item).To(Equal(storage.Stacks[0]))
	})

	It("Should get all orgs", func() {
		items, err := inquisitor.GetAllOrgs()
		Expect(err).To(BeNil())
		Expect(items).Should(ConsistOf(storage.Orgs))
	})

	It("Should get org by guid", func() {
		item, err := inquisitor.GetOrgByGuid("Org-1")
		Expect(err).To(BeNil())
		Expect(item).Should(Equal(storage.Orgs[0]))
	})

	It("Should get org by name", func() {
		item, err := inquisitor.GetOrgByName("Org-1")
		Expect(err).To(BeNil())
		Expect(item).To(Equal(storage.Orgs[0]))
	})

	It("Should get all buildpacks", func() {
		items, err := inquisitor.GetAllBuildpacks()
		Expect(err).To(BeNil())
		Expect(items).Should(ConsistOf(storage.Buildpacks))
	})

	It("Should get buildpack by guid", func() {
		item, err := inquisitor.GetBuildpackByGuid("Buildpack-1")
		Expect(err).To(BeNil())
		Expect(item).Should(Equal(storage.Buildpacks[0]))
	})

	It("Should get buildpack by name", func() {
		item, err := inquisitor.GetBuildpackByName("Buildpack-1")
		Expect(err).To(BeNil())
		Expect(item).To(Equal(storage.Buildpacks[0]))
	})

})
