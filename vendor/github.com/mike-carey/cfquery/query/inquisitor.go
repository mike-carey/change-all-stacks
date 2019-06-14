package query

import (
	// "fmt"
	"sync"
	// "reflect"

	"github.com/mike-carey/cfquery/cf"

	// "github.com/iancoleman/strcase"
)

type Service interface {
	ServiceName() string
}

type inquisitor struct {
	CFClient cf.CFClient
	services map[string]Service
	mutex       *sync.Mutex
}

func NewInquisitor(cfClient cf.CFClient) Inquisitor {
	return &inquisitor{
		CFClient: cfClient,
		services: make(map[string]Service, 0),
		mutex: &sync.Mutex{},
	}
}

func (i *inquisitor) lock() {
	i.mutex.Lock()
}

func (i *inquisitor) unlock() {
	i.mutex.Unlock()
}

// func (i *inquisitor) getService(name string) Service {
// 	if service, ok := i.services[name]; ok {
// 		return service
// 	}
//
// 	className := strcase.ToCamel(name)
// 	serviceName := fmt.Sprintf("%sService", className)
// 	funcName := fmt.Sprintf("new%s", serviceName)
//
// 	fnZeroValue := reflect.ValueOf(nil)
//
// 	fn := reflect.ValueOf(i).MethodByName(funcName)
// 	if fn == fnZeroValue {
// 		panic(fmt.Sprintf("Unknown service %s", serviceName))
// 	}
//
// 	serviceValue := fn.Call([]reflect.Value{})[0]
//
// 	if serviceValue == fnZeroValue {
// 		panic(fmt.Sprintf("Error creating service %s", serviceName))
// 	}
//
// 	service := serviceValue.Interface().(Service)
//
// 	i.lock()
// 	i.services[name] = service
// 	i.unlock()
//
// 	return service
// }
