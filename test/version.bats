#!/usr/bin/env bats

source $BATS_TEST_DIRNAME/.envrc

load helpers/print/bprint

@test "It should print version" {
  run $SCRIPT --version

  [ $status -eq 0 ]
  [ "$output" = "$VERSION" ]
}
