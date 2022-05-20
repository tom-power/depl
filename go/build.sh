#!/bin/bash
go test -count=1 ./flanalystsh/ &&
go test -count=1 ./flanalystsh_test/ &&
if [[ ! -d ./build ]]; then
  mkdir -p ./build
fi
go build -o ./build/flanalystsh main.go