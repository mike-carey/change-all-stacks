#!/usr/bin/env bash

set -euo pipefail

function inquisitor() {
  pkg=''
  keep_backup=false
  args=()
  while [[ -n "${1:-}" ]]; do
    case $1 in
      -pkg )
        pkg="$2"
        shift
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

  set -- ${args[@]}

  local file_in=$1
  local file_out=$(basename $1 .erb)

  erb $file_in > $file_out
  if [[ -n "$pkg" ]]; then
    sed -i.bak "s/package .*/package $pkg/g" $file_out
    rm $file_out.bak
  fi
  go fmt $file_out
}

if [[ ${BASH_SOURCE[0]} != $0 ]]; then
  export -f inquisitor
else
  set -euo pipefail

  inquisitor "${@:-}"
  exit $?
fi
