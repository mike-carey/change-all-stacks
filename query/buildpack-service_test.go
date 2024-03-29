// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package query_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/mike-carey/change-all-stacks/query"
	"fmt"

	fakes "github.com/mike-carey/change-all-stacks/query/fakes"
)

var _ = Describe(fmt.Sprintf("%sService", getBuildpackString()), func() {

	var (
		storage    *Storage
		fakeClient *fakes.FakeCFClient
		service    *BuildpackService
	)

	BeforeEach(func() {
		storage = GetStorage()
		fakeClient = GetFakeClient(storage)
		service = NewBuildpackService(fakeClient)
	})

	It(fmt.Sprintf("Should get all %s", getBuildpackString()), func() {
		items, err := service.GetAllBuildpacks()
		Expect(err).To(BeNil())
		Expect(items).Should(ConsistOf(storage.Buildpacks))
	})

	It(fmt.Sprintf("Should get %s by guid", getBuildpackString()), func() {
		item, err := service.GetBuildpackByGuid(fmt.Sprintf("%s-1", getBuildpackString()))
		Expect(err).To(BeNil())
		Expect(item).Should(Equal(storage.Buildpacks[0]))
	})

	It(fmt.Sprintf("Should get %s by name", getBuildpackString()), func() {
		item, err := service.GetBuildpackByName(fmt.Sprintf("%s-1", getBuildpackString()))
		Expect(err).To(BeNil())
		Expect(item).To(Equal(storage.Buildpacks[0]))
	})
})
