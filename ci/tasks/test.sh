#!/bin/bash
set -e -u -x

cd git-phlow-repo/

go get -d ./...
go get github.com/stretchr/testify

go test -v ./...
