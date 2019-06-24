package generic_test

import (
	"fmt"
	"strings"

	"github.com/cloudfoundry-community/go-cfclient"

	"github.com/cheekybits/genny/generic"
)

type Item generic.Type

func getItemString() string {
	s := strings.Split(fmt.Sprintf("%T", &cfclient.Item{}), ".")
	return s[len(s)-1]
}
