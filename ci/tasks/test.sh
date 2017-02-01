#!/bin/bash
set -e -u -x

#RESOLVE OWN PROJECT AS DEPENDENCY
buildpath=$(pwd)

mkdir -p $GOPATH/src/github.com/praqma
cp -R git-phlow/ $GOPATH/src/github.com/praqma

# RESOLVE EXTERNAL DEPENDENCIES
cd $GOPATH/src/github.com/praqma/git-phlow
go get -d -v
go get github.com/stretchr/testify
go get github.com/smartystreets/goconvey

#cp binary to buildpath
go test -v ./...
