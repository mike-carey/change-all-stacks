# Change All Stacks
A script to change all stacks over all apps over all foundations.

## Usage
Create a `cf.json` file to hold configurations.  The file should be an object/map where every key is a unique name for a foundation and the cfclient configuration associated with that foundation as the value.  Config options defined in (Cloudfoundry's Go CFClient)[https://github.com/cloudfoundry-community/go-cfclient].

Running `change-all-stacks X Y` will change all apps with stack `X` to stack `Y`.
