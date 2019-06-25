package generic

import (
	"io"
	"fmt"
	"sync"

	"github.com/cloudfoundry-community/go-cfclient"

	"github.com/cheekybits/genny/generic"
)

type Item generic.Type

type ItemService struct {
	client CFClient
	logger io.Writer
	cacheList []cfclient.Item
	cacheMap map[string]cfclient.Item
	mutex *sync.Mutex
	fullyLoaded bool
}

func NewItemService(client CFClient, logger io.Writer) *ItemService {
	return &ItemService{
		client: client,
		logger: logger,
		cacheList: make([]cfclient.Item, 0),
		cacheMap: make(map[string]cfclient.Item, 0),
		mutex: &sync.Mutex{},
		fullyLoaded: false,
	}
}

func (s *ItemService) logf(msg string, args ...interface{}) {
	s.logger.Write([]byte(fmt.Sprintf(msg + "\n", args...)))
}

func (s *ItemService) lock() {
	s.logf("Locking")
	s.mutex.Lock()
}

func (s *ItemService) unlock() {
	s.logf("Unlocking")
	s.mutex.Unlock()
}

func (s *ItemService) GetAllItems() ([]cfclient.Item, error) {
	s.lock()

	if !s.fullyLoaded {
		s.logf("%T is not already fully loaded", s)
		items, err := s.client.ListItems()
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

func (s *ItemService) GetItemByGuid(guid string) (cfclient.Item, error) {
	s.lock()
	defer s.unlock()

	if item, ok := s.cacheMap[guid]; ok {
		s.logf("Found a cached %T with a guid of %s", item, guid)
		return item, nil
	}

	if s.fullyLoaded {
		s.logf("%T is already fully loaded but did not find %s in cacheMap", s, guid)
		item := cfclient.Item{}
		return item, fmt.Errorf("Could not find %T by guid: %s", item, guid)
	}

	s.logf("Did not find cached %T and %T is not fully loaded, querying by guid: %s", cfclient.Item{}, s, guid)
	i, err := s.client.GetItemByGuid(guid)
	if err != nil {
		return cfclient.Item{}, nil
	}

	s.logf("Saving off a single %T to the cacheMap with guid: %s", i, i.Guid)
	s.cacheMap[i.Guid] = i
	return i, nil
}

func (s *ItemService) GetItemByName(name string) (cfclient.Item, error) {
	items, err := s.GetAllItems()
	if err != nil {
		return cfclient.Item{}, err
	}

	s.lock()
	defer s.unlock()

	for _, item := range items {
		if item.Name == name {
			return item, nil
		}
	}

	item := cfclient.Item{}
	return item, fmt.Errorf("Could not find %T by name: %s", item, name)
}

// Proxy all functions onto the inquisitor struct
func (i *inquisitor) GetAllItems() ([]cfclient.Item, error) {
	return i.getItemService().GetAllItems()
}

func (i *inquisitor) GetItemByGuid(guid string) (cfclient.Item, error) {
	return i.getItemService().GetItemByGuid(guid)
}

func (i *inquisitor) GetItemByName(name string) (cfclient.Item, error) {
	return i.getItemService().GetItemByName(name)
}
