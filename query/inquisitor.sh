#!/usr/bin/env bash

set -euo pipefail

function inquisitor() {
  keep_backup=false
  args=()
  while [[ -n "${1:-}" ]]; do
    case $1 in
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
  go fmt $file_out
}

if [[ ${BASH_SOURCE[0]} != $0 ]]; then
  export -f inquisitor
else
  set -euo pipefail

  inquisitor "${@:-}"
  exit $?
fi
