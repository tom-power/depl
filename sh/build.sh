#!/bin/bash
# build
cd ./go && ./build.sh

# copy to local
cd ../ &&
cp -p ./go/build/_depl ./config/depl.sh ~/bin/
# cp -p ./config/.oh-my-zsh/custom/completions/_depl ~/config/.oh-my-zsh/custom/completions/
