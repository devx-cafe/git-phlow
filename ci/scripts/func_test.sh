#!/bin/bash
set -e -u -x

#MOVE PROJECT TO GOPATH
buildpath=$(pwd)

mkdir -p $GOPATH/src/github.com/praqma
cp -R git-phlow/ $GOPATH/src/github.com/praqma

# RESOLVE DEPENDENCIES - TEST AND PRODUCTION
cd $GOPATH/src/github.com/praqma/git-phlow
go get -d -t -v ./...

#RUN WHOLE TEST SUITE -VERBOSE
go test -v -p 1 ./...