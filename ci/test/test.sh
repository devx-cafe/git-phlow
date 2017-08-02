#!/bin/sh
set -e -u -x

#Set new GOPATH to current dir
export GOPATH=$PWD
export PATH=$PATH:$GOPATH/bin

# RESOLVE DEPENDENCIES - TEST AND PRODUCTION
cd $GOPATH/src/github.com/praqma/git-phlow
go get -t -d -v ./...

#install ginkgo
go install github.com/onsi/ginkgo/ginkgo


# run tests
go test -p 1 -v ./...



