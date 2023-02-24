#!/bin/bash
# build
cd ./go && ./build.sh

# copy to remote
cd ../
source sh/.env &&
scp -p ./go/build/_depl ./config/depl.sh $remoteUser@$remoteHost:$deployDir &&
scp -p ./config/.oh-my-zsh/custom/completions/_depl $remoteUser@$remoteHost:~/config/.oh-my-zsh/custom/completions/
