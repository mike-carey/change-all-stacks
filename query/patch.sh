#!/usr/bin/env bash

###
# This file modifies the generated file from ifacemaker to import cfclient and modify the classes that needed to be referenced  from that package.
##

function patch() {
  local keep_backup=false
  local args=()
  local ignore=()
  while [[ -n "${1:-}" ]]; do
    case $1 in
      --keep-backup )
        keep_backup=true
        ;;
      --ignore )
        ignore+=($2)
        shift
        ;;
      -- )
        shift
        args+=($@)
        break
        ;;
      * )
        args+=("$1")
        ;;
    esac
    shift
  done

  set -- ${args[@]}

  for file in ${@}; do
    if [[ " ${ignore[@]} " =~ " $file " ]]; then
      echo "Skipping $file"
      continue
    fi

    echo "Modifying $file"

    if [[ "$file" =~ .*_test.go ]]; then
      sed -i.bak 's|\(import (\)|\1\
'$'\t. "github.com/onsi/ginkgo"\
'$'\t. "github.com/onsi/gomega"\
'$'\t. "github.com/mike-carey/change-all-stacks/query"|g' $file
    else
      sed -i.bak 's|\(import (\)|\1\
'$'\tcfclient "github.com/cloudfoundry-community/go-cfclient"|g' $file
      sed -i.bak '/^[type|/\/\/]/! s/\([][( \*]\)\([A-Z]\)/\1cfclient.\2/g' $file
    fi
    if [[ $keep_backup != true ]]; then
      rm -f $file.bak
    fi
  done
}

if [[ ${BASH_SOURCE[0]} != $0 ]]; then
  export -f patch
else
  set -eo pipefail

  patch "${@:-}"
  exit $?
fi
