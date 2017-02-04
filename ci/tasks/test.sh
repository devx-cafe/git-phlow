#!/bin/bash
set -e -u -x

#MOVE PROJECT TO GOPATH
buildpath=$(pwd)

mkdir -p $GOPATH/src/github.com/praqma
cp -R git-phlow/ $GOPATH/src/github.com/praqma

# RESOLVE DEPENDENCIES - TEST AND PRODUCTION
cd $GOPATH/src/github.com/praqma/git-phlow
go get -d -v ./...
go get github.com/stretchr/testify
go get github.com/smartystreets/goconvey
go get github.com/spf13/cobra

#RUN WHOLE TEST SUITE -VERBOSE
go test -v ./...
