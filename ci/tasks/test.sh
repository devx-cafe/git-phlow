#!/bin/bash
set -e -u -x

cd git-phlow-repo/

go get -d -t
go test -v ./...
