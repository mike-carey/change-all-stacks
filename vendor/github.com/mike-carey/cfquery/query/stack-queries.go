package query

import (
	"fmt"
	"net/url"

	"github.com/cloudfoundry-community/go-cfclient"
)

func (s *StackService) GetStackByName(name string) (cfclient.Stack, error) {
	v := url.Values{}
	v["name"] = []string{name}

	stacks, err := s.Client.ListStacksByQuery(v)
	if err != nil {
		return cfclient.Stack{}, nil
	}

	if len(stacks) < 1 {
		return cfclient.Stack{}, fmt.Errorf("Could not find stack with name of '%s'", name)
	}

	if len(stacks) > 1 {
		return cfclient.Stack{}, fmt.Errorf("Multiple stacks with name of '%s'", name)
	}

	return stacks[0], nil
}

func (i *inquisitor) GetStackByName(name string) (cfclient.Stack, error) {
	return i.getStackService().GetStackByName(name)
}
