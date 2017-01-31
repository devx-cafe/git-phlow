#!/bin/bash
set -e -u -x

#RESOLVE OWN PROJECT AS DEPENDENCY
buildpath=$pwd
deppath= $GOPATH/src/github.com/praqma

mkdir -p $deppath
cp -R git-phlow/ $deppath

# RESOLVE EXTERNAL DEPENDENCIES
cd $deppath/git-phlow
go get -d
go get github.com/stretchr/testify
go get github.com/smartystreets/goconvey

#cp binary to buildpath
go test -v ./...
