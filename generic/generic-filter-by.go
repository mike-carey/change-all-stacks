package generic

import (
	"errors"

	"github.com/cloudfoundry-community/go-cfclient"

	"github.com/cheekybits/genny/generic"
)

type Item generic.Type

func ItemFilterBy(items []cfclient.Item, check func(cfclient.Item) (bool, error)) ([]cfclient.Item, error) {
	pool := make([]cfclient.Item, 0)

	for _, item := range items {
		ok, err := check(item)
		if err != nil {
			return pool, errors.Wrap(err, "Could not check filter for item")
		}

		if ok {
			pool = append(pool, item)
		}
	}

	return pool, nil
}
