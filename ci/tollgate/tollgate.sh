#!/bin/sh
set -e -u -x

#CREATE GO DIRECTORY STRUCTURE
#THE STRUCTURE IS NECESSARY FOR GO TOOLS OTHERWISE
# 'build' AND 'get' WILL FAIL

mkdir -p $GOPATH/src/github.com/praqma
cp -R tollgate/ $GOPATH/src/github.com/praqma/git-phlow

# RESOLVE DEPENDENCIES - TEST AND PRODUCTION
cd $GOPATH/src/github.com/praqma/git-phlow
go get -d -v ./...

# MINIMUM REQUREMENT FOR CODE SHARING IS BUILDABLE
go build