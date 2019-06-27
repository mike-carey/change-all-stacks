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

var _ = Describe(fmt.Sprintf("%sService", getOrgString()), func() {

	var (
		storage    *Storage
		fakeClient *fakes.FakeCFClient
		service    *OrgService
	)

	BeforeEach(func() {
		storage = GetStorage()
		fakeClient = GetFakeClient(storage)
		service = NewOrgService(fakeClient)
	})

	It(fmt.Sprintf("Should get all %s", getOrgString()), func() {
		items, err := service.GetAllOrgs()
		Expect(err).To(BeNil())
		Expect(items).Should(ConsistOf(storage.Orgs))
	})

	It(fmt.Sprintf("Should get %s by guid", getOrgString()), func() {
		item, err := service.GetOrgByGuid(fmt.Sprintf("%s-1", getOrgString()))
		Expect(err).To(BeNil())
		Expect(item).Should(Equal(storage.Orgs[0]))
	})

	It(fmt.Sprintf("Should get %s by name", getOrgString()), func() {
		item, err := service.GetOrgByName(fmt.Sprintf("%s-1", getOrgString()))
		Expect(err).To(BeNil())
		Expect(item).To(Equal(storage.Orgs[0]))
	})
})
