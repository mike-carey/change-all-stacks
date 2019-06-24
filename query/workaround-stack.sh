#!/usr/bin/env bash

set -euo pipefail

FILE=$1
# Replace the c.Client.GetStackByGuid call
sed -i.bak 's/\('$'\t''*\)\(.*\)s.client.GetStackByGuid\(.*\)/\1s.unlock()\
\1\2s.getStackByGuid\3/g' $FILE
sed -i.bak '/func (s \*.*) GetStackByGuid/,/s.unlock()/ s/\(.*\)\(defer s.unlock()\)/\1\/\/\2/' $FILE
rm $FILE.bak

cat >> $FILE <<HEREDOC

// Worksaround until merge: https://github.com/cloudfoundry-community/go-cfclient/pull/234
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
