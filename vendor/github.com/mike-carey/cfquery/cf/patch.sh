#!/usr/bin/env bash

###
# This file modifies the generated file from ifacemaker to import cfclient and modify the classes that needed to be referenced  from that package.
##

function patch() {
  keep_backup=false
  args=()
  while [[ -n "$1" ]]; do
    case $1 in
      --keep-backup )
        keep_backup=true
        ;;
      -- )
        args+=($@)
        break
        ;;
      * )
        args+=("$1")
        ;;
    esac
    shift
  done

  file=client.go

  echo "Modifying $file"
  sed -i.bak 's|\(import (\)|\1\
'$'\tcfclient "github.com/cloudfoundry-community/go-cfclient"|g' $file
  sed -i.bak '/^[type|/\/\/]/! s/\([][( \*]\)\([A-Z]\)/\1cfclient.\2/g' $file
  if [[ $keep_backup != true ]]; then
    rm -f $file.bak
  fi
}

if [[ ${BASH_SOURCE[0]} != $0 ]]; then
export -f patch
else
set -euo pipefail

patch "${@:-}"
exit $?
fi
