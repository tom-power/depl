#!/bin/bash
command=$(_depl "$1" "$2" "$3")
if [[ "$1" == "list" || "$2" == "--explain" || "$3" == "--explain" ]]; then
  echo "$command"
else
  eval "$command"
fi
