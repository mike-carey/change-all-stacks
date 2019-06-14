package query

import (
	"github.com/cloudfoundry-community/go-cfclient"
)

func (g ServiceInstances) GroupBySpace(_ Inquisitor) (ServiceInstanceGroup, error) {
	return ServiceInstanceGroupBy(g, func(si cfclient.ServiceInstance) (string, error) {
		return si.SpaceGuid, nil
	})
}

func (g ServiceInstances) GroupByOrg(i Inquisitor) (ServiceInstanceGroup, error) {
	return ServiceInstanceGroupBy(g, func(si cfclient.ServiceInstance) (string, error) {
		s, e := i.GetSpaceByGuid(si.SpaceGuid)
		if e != nil {
			return "", e
		}

		return s.OrganizationGuid, nil
	})
}

// GroupBySpaceAndOrg ...
func (g ServiceInstances) GroupBySpaceAndOrg(i Inquisitor) (MappedServiceInstanceGroup, error) {
	ag, err := g.GroupBySpace(i)
	if err != nil {
		return nil, err
	}

	return ag.GroupByOrg(i)
}

func (g ServiceInstanceGroup) GroupByOrg(i Inquisitor) (MappedServiceInstanceGroup, error) {
	return ServiceInstanceGroupMappedSliceBy(g, func(_ string, apps ServiceInstances) (string, error) {
		s, e := i.GetSpaceByGuid(apps[0].SpaceGuid)
		if e != nil {
			return "", e
		}

		return s.OrganizationGuid, nil
	})
}
