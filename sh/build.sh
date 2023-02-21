#!/bin/bash
cd ./go && ./build.sh

# copy to local
cd ../ && cp -p ./go/build/_depl ./config/depl.sh ~/bin/
