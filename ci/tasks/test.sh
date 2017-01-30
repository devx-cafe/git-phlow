#!/bin/bash
set -e -u -x

cd git-phlow-repo/

go get -d
go test -v /...
