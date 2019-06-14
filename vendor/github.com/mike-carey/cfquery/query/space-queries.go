package query

import (
	"github.com/cloudfoundry-community/go-cfclient"
)

func (g Spaces) GroupByOrg(_ Inquisitor) (SpaceGroup, error) {
	return SpaceGroupBy(g, func(space cfclient.Space) (string, error) {
		return space.OrganizationGuid, nil
	})
}
