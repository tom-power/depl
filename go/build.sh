#!/bin/bash
go test -count=1 ./depl/ &&
go test -count=1 ./depl_test/ &&
if [[ ! -d ./build ]]; then
  mkdir -p ./build
fi
go build -o ./build/depl main.go