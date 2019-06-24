// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package query_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/mike-carey/change-all-stacks/query"
	"fmt"
	"strings"

	cfclient "github.com/cloudfoundry-community/go-cfclient"
)

var _ = Describe(fmt.Sprintf("%sGroupBy", getSpaceString()), func() {

	var (
		items []cfclient.Space
	)

	BeforeEach(func() {
		items = []cfclient.Space{
			cfclient.Space{
				Guid: fmt.Sprintf("%s-%d-%d", getSpaceString(), 1, 1),
			},
			cfclient.Space{
				Guid: fmt.Sprintf("%s-%d-%d", getSpaceString(), 1, 2),
			},
			cfclient.Space{
				Guid: fmt.Sprintf("%s-%d-%d", getSpaceString(), 2, 1),
			},
			cfclient.Space{
				Guid: fmt.Sprintf("%s-%d-%d", getSpaceString(), 2, 2),
			},
		}
	})

	It("Should group by first number", func() {
		i, e := SpaceGroupBy(items, func(item cfclient.Space) (string, error) {
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
		i, e := SpaceGroupBy(items, func(item cfclient.Space) (string, error) {
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
