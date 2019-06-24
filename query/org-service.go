// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package query

import (
	"fmt"
	"io"
	"sync"

	cfclient "github.com/cloudfoundry-community/go-cfclient"
)

type OrgService struct {
	client      CFClient
	logger      io.Writer
	cacheList   []cfclient.Org
	cacheMap    map[string]cfclient.Org
	mutex       *sync.Mutex
	fullyLoaded bool
}

func NewOrgService(client CFClient, logger io.Writer) *OrgService {
	return &OrgService{
		client:      client,
		logger:      logger,
		cacheList:   make([]cfclient.Org, 0),
		cacheMap:    make(map[string]cfclient.Org, 0),
		mutex:       &sync.Mutex{},
		fullyLoaded: false,
	}
}

func (s *OrgService) logf(msg string, args ...interface{}) {
	s.logger.Write([]byte(fmt.Sprintf(msg+"\n", args...)))
}

func (s *OrgService) lock() {
	s.logf("Locking")
	s.mutex.Lock()
}

func (s *OrgService) unlock() {
	s.logf("Unlocking")
	s.mutex.Unlock()
}

func (s *OrgService) GetAllOrgs() ([]cfclient.Org, error) {
	s.lock()

	if !s.fullyLoaded {
		s.logf("%T is not already fully loaded", s)
		items, err := s.client.ListOrgs()
		if err != nil {
			return nil, err
		}

		s.logf("Writing cache list for %T", s)
		s.cacheList = items

		go func() {
			defer s.unlock()

			s.logf("Writing cache map for %T", s)
			for _, item := range items {
				s.cacheMap[item.Guid] = item
			}
		}()
	} else {
		defer s.unlock()
	}

	return s.cacheList, nil
}

func (s *OrgService) GetOrgByGuid(guid string) (cfclient.Org, error) {
	s.lock()
	defer s.unlock()

	if item, ok := s.cacheMap[guid]; ok {
		s.logf("Found a cached %T with a guid of %s", item, guid)
		return item, nil
	}

	if s.fullyLoaded {
		s.logf("%T is already fully loaded but did not find %s in cacheMap", s, guid)
		item := cfclient.Org{}
		return item, fmt.Errorf("Could not find %T by guid: %s", item, guid)
	}

	s.logf("Did not find cached %T and %T is not fully loaded, querying by guid: %s", cfclient.Org{}, s, guid)
	i, err := s.client.GetOrgByGuid(guid)
	if err != nil {
		return cfclient.Org{}, nil
	}

	s.logf("Saving off a single %T to the cacheMap with guid: %s", i, i.Guid)
	s.cacheMap[i.Guid] = i
	return i, nil
}

func (s *OrgService) GetOrgByName(name string) (cfclient.Org, error) {
	items, err := s.GetAllOrgs()
	if err != nil {
		return cfclient.Org{}, err
	}

	s.lock()
	defer s.unlock()

	for _, item := range items {
		if item.Name == name {
			return item, nil
		}
	}

	item := cfclient.Org{}
	return item, fmt.Errorf("Could not find %T by name: %s", item, name)
}

// Proxy all functions onto the inquisitor struct
func (i *inquisitor) GetAllOrgs() ([]cfclient.Org, error) {
	return i.getOrgService().GetAllOrgs()
}

func (i *inquisitor) GetOrgByGuid(guid string) (cfclient.Org, error) {
	return i.getOrgService().GetOrgByGuid(guid)
}

func (i *inquisitor) GetOrgByName(name string) (cfclient.Org, error) {
	return i.getOrgService().GetOrgByName(name)
}