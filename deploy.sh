#!/bin/bash
root="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null 2>&1 && pwd)"

$root/build.sh

# local
cp -p $root/go/build/depl ~/bin/

# remote
. $root/.env &&
scp ./go/build/depl $remoteUser@$remoteHost:$deployDir