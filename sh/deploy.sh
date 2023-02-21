#!/bin/bash
source sh/.env &&
scp -p ./go/build/_depl ./config/depl.sh $remoteUser@$remoteHost:$deployDir
