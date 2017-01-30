#!/bin/bash
set -e -u -x

go get -d
go test -v ../../...
