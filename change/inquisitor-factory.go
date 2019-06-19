package change

import (
	"github.com/mike-carey/cfquery/query"

	"github.com/cloudfoundry-community/go-cfclient"
)

type InquisitorFactory interface {
	CreateInquisitor(config cfclient.Config) (query.Inquisitor, error)
}

func NewInquisitorFactory() InquisitorFactory {
	return &inquisitorFactory{}
}

type inquisitorFactory struct {}

func (i *inquisitorFactory) CreateInquisitor(config cfclient.Config) (query.Inquisitor, error) {
	cli, err := cfclient.NewClient(&config)
	if err != nil {
		return nil, err
	}

	return query.NewInquisitor(cli), nil
}
