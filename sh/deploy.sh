#!/bin/bash
# build
sh/build.sh

deplBinPath=./go/build/_depl
deplShPath=./config/depl.sh
deplCompletionPath=./config/_depl

# copy to local
cp -p $deplBinPath $deplShPath ~/bin/
#cp -p $deplCompletionPath ~/config/.oh-my-zsh/custom/completions/

# copy to remote
source sh/.env &&
scp -p $deplBinPath $deplShPath $remoteUser@$remoteHost:$deployDir
#scp -p $deplCompletionPath $remoteUser@$remoteHost:~/config/.oh-my-zsh/custom/completions/
