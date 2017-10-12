#!/bin/bash

export GOPATH=$PWD
export PATH=$PATH:$GOPATH/bin

cd $GOPATH/src/github.com/praqma/git-phlow

go build ci/static/cover.go

go get -d -t ./...
go get github.com/mattn/goveralls

#Run the coverage
ci/static/gencover.sh $(go list ./...)

cat coverfiles/percentage
./cover

