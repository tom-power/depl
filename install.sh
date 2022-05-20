#!/bin/bash
root="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null 2>&1 && pwd)"

$root/build.sh

cp $root/go/build/flanalystsh ~/proj/os/flanalyst.ovh/
cp $root/sh/flanalyst.sh ~/proj/os/flanalyst.ovh/

echo "installed"