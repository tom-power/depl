#!/bin/bash
root="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null 2>&1 && pwd)"

# copy to remote
. $root/.env &&
scp -p $root/go/build/_depl $root/sh/depl.sh $remoteUser@$remoteHost:$deployDir