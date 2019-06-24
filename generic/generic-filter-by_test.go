package generic_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/mike-carey/change-all-stacks/query/generic"

	"github.com/cheekybits/genny/generic"
)

type Item generic.Type

var _ = Describe(fmt.Sprintf("%sFilterBy", getItemString()), func() {

	var (
		items []cfclient.Item
	)

	BeforeEach(func() {
		items = []cfclient.Item{
			cfclient.Item{
				Guid: fmt.Sprintf("%s-%d-%d", getItemString(), 1, 1),
			},
			cfclient.Item{
				Guid: fmt.Sprintf("%s-%d-%d", getItemString(), 1, 2),
			},
			cfclient.Item{
				Guid: fmt.Sprintf("%s-%d-%d", getItemString(), 2, 1),
			},
			cfclient.Item{
				Guid: fmt.Sprintf("%s-%d-%d", getItemString(), 2, 2),
			},
		}
	})

	It("Should filter by first number", func() {
		By("Equaling 1")
		i1, e1 := ItemFilterBy(items, func(item cfclient.Item) (bool, error){
			s := strings.Split(item.Guid, "-")
			return s[1] == "1", nil
		})

		Expect(e1).To(BeNil())
		Expect(i1).Should(HaveLen(2))
		Expect(i1).Should(ConsistOf(items[0], items[1]))

		By("Equaling 2")
		i2, e2 := ItemFilterBy(items, func(item cfclient.Item) (bool, error){
			s := strings.Split(item.Guid, "-")
			return s[1] == "2", nil
		})
		Expect(e2).To(BeNil())
		Expect(i2).Should(HaveLen(2))
		Expect(i2).Should(ConsistOf(items[2], items[3]))
	})

	It("Should filter by second number", func() {
		By("Equaling 1")
		i1, e1 := ItemFilterBy(items, func(item cfclient.Item) (bool, error){
			s := strings.Split(item.Guid, "-")
			return s[2] == "1", nil
		})

		Expect(e1).To(BeNil())
		Expect(i1).Should(HaveLen(2))
		Expect(i1).Should(ConsistOf(items[0], items[2]))

		By("Equaling 2")
		i2, e2 := ItemFilterBy(items, func(item cfclient.Item) (bool, error){
			s := strings.Split(item.Guid, "-")
			return s[2] == "2", nil
		})
		Expect(e2).To(BeNil())
		Expect(i2).Should(HaveLen(2))
		Expect(i2).Should(ConsistOf(items[1], items[3]))
	})
})
