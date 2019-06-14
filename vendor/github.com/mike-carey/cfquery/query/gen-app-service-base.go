// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package query

import (
	"fmt"
	"reflect"
	"strings"
	"sync"

	"github.com/cloudfoundry-community/go-cfclient"
	"github.com/mike-carey/cfquery/cf"
	"github.com/mike-carey/cfquery/logger"
	"github.com/mike-carey/cfquery/util"
)

type AppService struct {
	Client      cf.CFClient
	storage     AppMap
	filled      bool
	mutex       *sync.Mutex
	serviceName string
	key         string
}

func NewAppService(client cf.CFClient) *AppService {
	return &AppService{
		Client:  client,
		storage: make(map[string]cfclient.App, 0),
		filled:  false,
		mutex:   &sync.Mutex{},
	}
}

func (s *AppService) ServiceName() string {
	if s.serviceName == "" {
		name := fmt.Sprintf("%T", s)

		_name := strings.Split(name, ".")
		name = _name[len(_name)-1]

		s.serviceName = fmt.Sprintf("%s", name)
	}

	return s.serviceName
}

func (s *AppService) Key() string {
	if s.key == "" {
		key := s.ServiceName()
		s.key = key[:len(key)-len("Service")]
		logger.Info(s.key)
	}

	return s.key
}

func (s *AppService) lock() {
	s.mutex.Lock()
	logger.Infof("Locked %v", reflect.TypeOf(s))
}

func (s *AppService) unlock() {
	s.mutex.Unlock()
	logger.Infof("Unlocked %v", reflect.TypeOf(s))
}

func (s *AppService) GetAppMap() (AppMap, error) {
	_, err := s.GetAllApps()
	if err != nil {
		return nil, err
	}

	return s.storage, nil
}

func (i *inquisitor) GetAppMap() (AppMap, error) {
	return i.getAppService().GetAppMap()
}

func (s *AppService) GetAppByGuid(guid string) (*cfclient.App, error) {
	s.lock()

	defer s.unlock()

	if s.filled {
		if val, ok := s.storage[guid]; ok {
			return &val, nil
		}
	}

	logger.Infof("Did not find %s in storage", guid)
	item, err := s.Client.GetAppByGuid(guid)
	if err != nil {
		return nil, err
	}

	// Save off in storage
	s.storage[guid] = item

	return &item, nil
}

func (i *inquisitor) GetAppByGuid(guid string) (*cfclient.App, error) {
	return i.getAppService().GetAppByGuid(guid)
}

func (s *AppService) GetManyAppsByGuid(guids ...string) (AppMap, error) {
	pool := make(AppMap, len(guids))

	type Result struct {
		Guid   string
		Object *cfclient.App
		Error  error
	}

	resCh := make(chan Result, len(guids))

	for _, guid := range guids {
		go func(guid string) {
			obj, err := s.GetAppByGuid(guid)
			res := Result{
				Guid:   guid,
				Error:  err,
				Object: obj,
			}

			resCh <- res
		}(guid)
	}

	errs := make([]error, 0)

	for _, _ = range guids {
		select {
		case res := <-resCh:
			if res.Error != nil {
				errs = append(errs, res.Error)
			}

			pool[res.Guid] = *res.Object
		}
	}

	if len(errs) > 0 {
		return nil, util.StackErrors(errs)
	}

	return pool, nil
}

func (i *inquisitor) GetManyAppsByGuid(guids ...string) (AppMap, error) {
	return i.getAppService().GetManyAppsByGuid(guids...)
}

func (s *AppService) GetAllApps() (Apps, error) {
	s.lock()

	if s.filled {
		logger.Infof("Reusing storage")
		siSlice := make(Apps, 0, len(s.storage))
		for _, si := range s.storage {
			siSlice = append(siSlice, si)
		}

		s.unlock()

		return siSlice, nil
	}

	logger.Infof("Calling out to CFClient")
	sis, err := s.Client.ListApps()
	if err != nil {
		return nil, err
	}

	go func(s *AppService, sis Apps) {
		logger.Infof("Storing contents to storage")
		for _, si := range sis {
			s.storage[si.Guid] = si
		}

		logger.Infof("Done storing contents to storage")
		s.filled = true

		s.unlock()
	}(s, sis)

	logger.Infof("Returning list while populating happens")
	return sis, nil
}

func (i *inquisitor) GetAllApps() (Apps, error) {
	return i.getAppService().GetAllApps()
}

func (i *inquisitor) newAppService() *AppService {
	return NewAppService(i.CFClient)
}

func (i *inquisitor) getAppService() *AppService {
	class := &AppService{}
	key := class.Key()

	if service, ok := i.services[key]; ok {
		return service.(*AppService)
	}

	service := i.newAppService()

	i.lock()
	i.services[key] = service
	i.unlock()

	return service
}
