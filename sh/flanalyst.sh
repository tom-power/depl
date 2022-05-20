#!/bin/bash
dir="$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )"
eval $(eval $dir/sh/flanalystsh $1)