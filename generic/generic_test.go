package generic_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/mike-carey/change-all-stacks/query"

	fakes "github.com/mike-carey/change-all-stacks/query/fakes"

	"github.com/cloudfoundry-community/go-cfclient"

	"github.com/cheekybits/genny/generic"
)

type Item generic.Type

var _ = Describe(fmt.Sprintf("%sService", getItemString()), func() {

	var (
		storage *Storage
		fakeClient *fakes.FakeCFClient
		service *ItemService
	)

	BeforeEach(func() {
		storage = GetStorage()
		fakeClient = GetFakeClient(storage)
		service = NewItemService(fakeClient, GinkgoWriter)
	})

	It(fmt.Sprintf("Should get all %s", getItemString()), func() {
		items, err := service.GetAllItems()
		Expect(err).To(BeNil())
		Expect(items).Should(ConsistOf(storage.Items))
	})

	It(fmt.Sprintf("Should get %s by guid", getItemString()), func() {
		item, err := service.GetItemByGuid(fmt.Sprintf("%s-1", getItemString()))
		Expect(err).To(BeNil())
		Expect(item).Should(Equal(storage.Items[0]))
	})

	It(fmt.Sprintf("Should get %s by name", getItemString()), func() {
		item, err := service.GetItemByName(fmt.Sprintf("%s-1", getItemString()))
		Expect(err).To(BeNil())
		Expect(item).To(Equal(storage.Items[0]))
	})
})
