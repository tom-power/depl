#!/bin/bash
command=$(_depl "$1" "$2" "$3")
if [[ "$2" == "--explain" || "$3" == "--explain" ]]; then
  echo "$command"
else
  eval "$command"
fi
