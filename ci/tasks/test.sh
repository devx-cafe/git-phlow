#!/bin/bash
set -e -u -x

cd git-phlow-repo/

tree $GOPATH -d -x

go get -d
go get github.com/stretchr/testify
go get github.com/smartystreets/goconvey

pwd
ls

go test -v ./...
