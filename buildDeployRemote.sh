#!/bin/bash
root="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null 2>&1 && pwd)"

$root/buildDeploy.sh

# remote
git push origin master -f
. $root/.env &&
ssh -t $remoteUser@$remoteHost "\
cd $deployDir && \
git fetch --all && git reset --hard origin/master && \
cd ./go && ./build.sh && cd ../ && \
cp -p ./go/build/flanalystsh ../ && \
cp -p ./sh/flanalyst.sh ../"