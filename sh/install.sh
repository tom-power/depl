#!/bin/bash
deplBinPath=./go/build/_depl
deplShPath=./config/depl.sh
deplCompletionPath=./config/_depl

function echoWithHyphens() {
    echo "--------------$1--------------"
}

# build
echoWithHyphens "start"
sh/build.sh
echoWithHyphens "built"

# copy to local
cp -p $deplBinPath $deplShPath ~/bin/
#cp -p $deplCompletionPath ~/config/.oh-my-zsh/custom/completions/
echoWithHyphens "copied to local"

# copy to remote
source sh/.env &&
scp -p $deplBinPath $deplShPath $remoteUser@$remoteHost:$deployDir
echoWithHyphens "copied to remote"
#scp -p $deplCompletionPath $remoteUser@$remoteHost:~/config/.oh-my-zsh/custom/completions/
echoWithHyphens "end"
