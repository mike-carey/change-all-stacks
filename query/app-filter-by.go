// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package query

import (
	cfclient "github.com/cloudfoundry-community/go-cfclient"
	"github.com/pkg/errors"
)

func AppFilterBy(items []cfclient.App, check func(cfclient.App) (bool, error)) ([]cfclient.App, error) {
	pool := make([]cfclient.App, 0)

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
