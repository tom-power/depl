#!/bin/bash
# build
sh/build.sh

# copy to local
cp -p ./go/build/_depl ./config/depl.sh ~/bin/
# cp -p ./config/.oh-my-zsh/custom/completions/_depl ~/config/.oh-my-zsh/custom/completions/

# copy to remote
source sh/.env &&
scp -p ./go/build/_depl ./config/depl.sh $remoteUser@$remoteHost:$deployDir
# scp -p ./config/.oh-my-zsh/custom/completions/_depl $remoteUser@$remoteHost:~/config/.oh-my-zsh/custom/completions/
