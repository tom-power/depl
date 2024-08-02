#!/bin/bash
deplBinPath=./go/build/_depl
deplShPath=./config/depl.sh
deplCompletionPath=./config/_depl

function echoWithStar() {
    echo "*$1"
}

# build
echoWithStar "start"
sh/build.sh
echoWithStar "built"

# copy to local
cp -p $deplBinPath $deplShPath ~/bin/
echoWithStar "copied to local"

# copy to local completions
if [ "$1" == "--completions" ]; then
    cp $deplCompletionPath ~/.oh-my-zsh/custom/completions/
fi
echoWithStar "copied to local completions"

# copy to remote
source sh/.env &&
scp -p $deplBinPath $deplShPath $remoteUser@$remoteHost:$deployDir
echoWithStar "copied to remote"
echoWithStar "end"
