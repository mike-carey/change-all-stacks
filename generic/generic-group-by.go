package generic

import (
	"errors"

	"github.com/cloudfoundry-community/go-cfclient"

	"github.com/cheekybits/genny/generic"
)

type Item generic.Type

func ItemGroupBy(items []cfclient.Item, getKey func(cfclient.Item) (string, error)) (map[string][]cfclient.Item, error) {
	pool := make(map[string][]cfclient.Item)

	for _, item := range items {
		key, err := getKey(item)
		if err != nil {
			logger.Errorf("Could not get key from item: %v", item)
			return pool, errors.Wrap(err, "Could not get key for item")
		}

		if _, ok := pool[key]; !ok {
			pool[key] = make([]cfclient.Item, 0)
		}

		pool[key] = append(pool[key], item)
	}

	return pool, nil
}
