package generic_test

import (
	"fmt"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/mike-carey/change-all-stacks/query/generic"

	"github.com/cloudfoundry-community/go-cfclient"

	"github.com/cheekybits/genny/generic"
)

type Item generic.Type

var _ = Describe(fmt.Sprintf("%sGroupBy", getItemString()), func() {

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

	It("Should group by first number", func() {
		i, e := ItemGroupBy(items, func(item cfclient.Item) (string, error){
			s := strings.Split(item.Guid, "-")
			return s[1], nil
		})

		Expect(e).To(BeNil())
		Expect(i).Should(HaveLen(2))
		Expect(i).Should(HaveKey("1"))
		Expect(i["1"]).Should(HaveLen(2))
		Expect(i["1"]).Should(ConsistOf(items[0], items[1]))
		Expect(i).Should(HaveKey("2"))
		Expect(i["2"]).Should(HaveLen(2))
		Expect(i["2"]).Should(ConsistOf(items[2], items[3]))
	})

	It("Should group by second number", func() {
		i, e := ItemGroupBy(items, func(item cfclient.Item) (string, error){
			s := strings.Split(item.Guid, "-")
			return s[2], nil
		})

		Expect(e).To(BeNil())
		Expect(i).Should(HaveLen(2))
		Expect(i).Should(HaveKey("1"))
		Expect(i["1"]).Should(HaveLen(2))
		Expect(i["1"]).Should(ConsistOf(items[0], items[2]))
		Expect(i).Should(HaveKey("2"))
		Expect(i["2"]).Should(HaveLen(2))
		Expect(i["2"]).Should(ConsistOf(items[1], items[3]))
	})

})
