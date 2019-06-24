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

type SpaceService struct {
	client      CFClient
	logger      io.Writer
	cacheList   []cfclient.Space
	cacheMap    map[string]cfclient.Space
	mutex       *sync.Mutex
	fullyLoaded bool
}

func NewSpaceService(client CFClient, logger io.Writer) *SpaceService {
	return &SpaceService{
		client:      client,
		logger:      logger,
		cacheList:   make([]cfclient.Space, 0),
		cacheMap:    make(map[string]cfclient.Space, 0),
		mutex:       &sync.Mutex{},
		fullyLoaded: false,
	}
}

func (s *SpaceService) logf(msg string, args ...interface{}) {
	s.logger.Write([]byte(fmt.Sprintf(msg+"\n", args...)))
}

func (s *SpaceService) lock() {
	s.logf("Locking")
	s.mutex.Lock()
}

func (s *SpaceService) unlock() {
	s.logf("Unlocking")
	s.mutex.Unlock()
}

func (s *SpaceService) GetAllSpaces() ([]cfclient.Space, error) {
	s.lock()

	if !s.fullyLoaded {
		s.logf("%T is not already fully loaded", s)
		items, err := s.client.ListSpaces()
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

func (s *SpaceService) GetSpaceByGuid(guid string) (cfclient.Space, error) {
	s.lock()
	defer s.unlock()

	if item, ok := s.cacheMap[guid]; ok {
		s.logf("Found a cached %T with a guid of %s", item, guid)
		return item, nil
	}

	if s.fullyLoaded {
		s.logf("%T is already fully loaded but did not find %s in cacheMap", s, guid)
		item := cfclient.Space{}
		return item, fmt.Errorf("Could not find %T by guid: %s", item, guid)
	}

	s.logf("Did not find cached %T and %T is not fully loaded, querying by guid: %s", cfclient.Space{}, s, guid)
	i, err := s.client.GetSpaceByGuid(guid)
	if err != nil {
		return cfclient.Space{}, nil
	}

	s.logf("Saving off a single %T to the cacheMap with guid: %s", i, i.Guid)
	s.cacheMap[i.Guid] = i
	return i, nil
}

func (s *SpaceService) GetSpaceByName(name string) (cfclient.Space, error) {
	items, err := s.GetAllSpaces()
	if err != nil {
		return cfclient.Space{}, err
	}

	s.lock()
	defer s.unlock()

	for _, item := range items {
		if item.Name == name {
			return item, nil
		}
	}

	item := cfclient.Space{}
	return item, fmt.Errorf("Could not find %T by name: %s", item, name)
}

// Proxy all functions onto the inquisitor struct
func (i *inquisitor) GetAllSpaces() ([]cfclient.Space, error) {
	return i.getSpaceService().GetAllSpaces()
}

func (i *inquisitor) GetSpaceByGuid(guid string) (cfclient.Space, error) {
	return i.getSpaceService().GetSpaceByGuid(guid)
}

func (i *inquisitor) GetSpaceByName(name string) (cfclient.Space, error) {
	return i.getSpaceService().GetSpaceByName(name)
}