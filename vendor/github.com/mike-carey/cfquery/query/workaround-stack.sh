#!/usr/bin/env bash

set -euo pipefail

FILE=$1
# Replace the c.Client.GetStackByGuid call
sed -i.bak 's/\(.*\)s.Client.GetStackByGuid\(.*\)/\1s.getStackByGuid\2/g' $FILE
rm $FILE.bak

cat >> $FILE <<HEREDOC

func (s *StackService) getStackByGuid(guid string) (cfclient.Stack, error) {
	// Grab all stacks and find the one that has the desired guid
	stacks, err := s.GetAllStacks()
	if err != nil {
		return cfclient.Stack{}, nil
	}

	for _, stack := range stacks {
		if stack.Guid == guid {
			return stack, nil
		}
	}

	return cfclient.Stack{}, fmt.Errorf("Could not find stack with guid: '%s'", guid)
}
HEREDOC
