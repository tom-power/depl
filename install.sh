#!/bin/bash
root="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null 2>&1 && pwd)"

$root/build.sh

cp -p $root/go/build/flanalystsh $root/../
cp -p $root/sh/flanalyst.sh $root/../