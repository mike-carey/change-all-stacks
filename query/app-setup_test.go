// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package query_test

import (
	"fmt"
	"strings"

	cfclient "github.com/cloudfoundry-community/go-cfclient"
)

func getAppString() string {
	s := strings.Split(fmt.Sprintf("%T", &cfclient.App{}), ".")
	return s[len(s)-1]
}
