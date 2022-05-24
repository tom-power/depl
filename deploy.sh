#!/bin/bash
root="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null 2>&1 && pwd)"

$root/build.sh

# remote
. $root/.env &&
scp -p \
./go/build/_depl $root/sh/depl.sh \
$remoteUser@$remoteHost:$deployDir