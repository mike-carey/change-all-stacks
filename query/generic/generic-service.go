package generic

import (
	"io"
	"sync"

	"github.com/mike-carey/change-all-stacks/logger"

	"github.com/cloudfoundry-community/go-cfclient"

	"github.com/cheekybits/genny/generic"
)

type Item generic.Type

type ItemService struct {
	client CFClient
	cacheList []cfclient.Item
	cacheMap map[string]cfclient.Item
	mutex *sync.Mutex
	fullyLoaded bool
}

func NewItemService(client CFClient) *ItemService {
	return &ItemService{
		client: client,
		cacheList: make([]cfclient.Item, 0),
		cacheMap: make(map[string]cfclient.Item, 0),
		mutex: &sync.Mutex{},
		fullyLoaded: false,
	}
}

func (s *ItemService) lock() {
	logger.Debugf("Locking %T", s)
	s.mutex.Lock()
	logger.Debugf("Locked %T", s)
}

func (s *ItemService) unlock() {
	logger.Debugf("Unlocking %T", s)
	s.mutex.Unlock()
	logger.Debugf("Unlocked %T", s)
}

func (s *ItemService) GetAllItems() ([]cfclient.Item, error) {
	s.lock()

	if !s.fullyLoaded {
		logger.Debugf("%T is not already fully loaded", s)
		items, err := s.client.ListItems()
		if err != nil {
			return nil, err
		}

		logger.Debugf("Writing cache list for %T", s)
		s.cacheList = items

		go func(s *ItemService, items []cfclient.Item) {
			defer s.unlock()

			logger.Debugf("Writing cache map for %T", s)
			for _, item := range items {
				s.cacheMap[item.Guid] = item
			}
		}(s, items)
	} else {
		defer s.unlock()
	}

	return s.cacheList, nil
}

func (s *ItemService) GetItemByGuid(guid string) (cfclient.Item, error) {
	s.lock()
	defer s.unlock()

	if item, ok := s.cacheMap[guid]; ok {
		logger.Debugf("Found a cached %T with a guid of %s", item, guid)
		return item, nil
	}

	if s.fullyLoaded {
		logger.Debugf("%T is already fully loaded but did not find %s in cacheMap", s, guid)
		item := cfclient.Item{}
		return item, NewNotFoundError(item, "guid", guid)
	}

	logger.Debugf("Did not find cached %T and %T is not fully loaded, querying by guid: %s", cfclient.Item{}, s, guid)
	i, err := s.client.GetItemByGuid(guid)
	if err != nil {
		return cfclient.Item{}, nil
	}

	logger.Debugf("Saving off a single %T to the cacheMap with guid: %s", i, i.Guid)
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
	return item, NewNotFoundError(item, "name", name)
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
