#!/usr/bin/env bash

function generate-interface() {
  local pkg=''
  local struct=''
  local interface=''
  local files_in=()
  local file_out=''

  local keep_backup=false
  local args=()
  local ignore=()
  while [[ -n "${1:-}" ]]; do
    case $1 in
      --pkg )
        pkg="$2"
        shift
        ;;
      --out )
        file_out="$2"
        shift
        ;;
      --interface )
        interface="$2"
        shift
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

  struct="$1"
  shift

  if [[ -z "$pkg" ]]; then
    pkg=$(basename $PWD)
  fi

  if [[ -z "$interface" ]]; then
    interface="${struct}Interface"
  fi

  if [[ -z "$file_out" ]]; then
    file_out=$(echo $struct | tr A-Z a-z)-interface.go
  fi

  for file in ${@}; do
    if [[ "$file" == "$file_out" ]]; then
      continue
    fi

    if [[ " ${ignore[@]} " =~ " $file " ]]; then
      continue
    fi

    files_in+=("--file $file")
  done

  echo "Generating interface: $interface to $file_out"
  ifacemaker --pkg $pkg --struct $struct --iface $interface --output $file_out ${files_in[@]}
}

if [[ ${BASH_SOURCE[0]} != $0 ]]; then
  export -f generate-interface
else
  set -euo pipefail

  generate-interface "${@:-}"
  exit $?
fi
