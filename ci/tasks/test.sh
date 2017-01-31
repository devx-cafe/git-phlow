#!/bin/bash
set -e -u -x

cd git-phlow-repo/

ls
go get -d
go get github.com/stretchr/testify
go get github.com/smartystreets/goconvey

pwd
ls

go test -v ./...
