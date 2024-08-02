#!/bin/bash
deplBinPath=./go/build/_depl
deplShPath=./config/depl.sh
deplCompletionPath=./config/_depl

function echoWithStar() {
    echo "*$1"
}

source sh/.env

# build
echoWithStar "start"
sh/build.sh
echoWithStar "built"

# copy to local
cp $deplBinPath $deplShPath $binPath
echoWithStar "copied to local"

# copy to local completions
if [ "$1" == "--completions" ]; then
    cp $deplCompletionPath ~/.oh-my-zsh/custom/completions/
    echoWithStar "copied to local completions"
fi

# copy to remote
if [ "$1" == "--remote" ]; then
    scp -$deplBinPath $deplShPath $remoteUser@$remoteHost:$deployDir
    echoWithStar "copied to remote"    
fi

echoWithStar "end"
