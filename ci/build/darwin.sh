#!/bin/sh
set -e -u -x

#GET VERSION AND PATH
#STORE VERSION AND PATH FOR BUILD
VERSION=$(cat gp-version/version)
BUILDPATH=$(pwd)

#CREATE GO DIRECTORY STRUCTURE
#THE STRUCTURE IS NECESSARY FOR GO TOOLS OTHERWISE
# 'build' AND 'get' WILL FAIL
mkdir -p $GOPATH/src/github.com/praqma
cp -R git-phlow/ $GOPATH/src/github.com/praqma

#NAVIGATE TO FOLDER AND GET DEPS
cd $GOPATH/src/github.com/praqma/git-phlow
go get -d -t -v ./...

#BUILD WE ONLY BUILD FOR amd64
export GOARCH=amd64

#DARWIN
#BUILD AND COMPRESS
export GOOS=darwin
go build -ldflags "-X   github.com/praqma/git-phlow/options.Version=`echo $VERSION` -X  github.com/praqma/git-phlow/options.Sha1=`git rev-parse HEAD` -X  github.com/praqma/git-phlow/options.Date=`date +'%d-%m-%Y'`"
tar -cvzf $BUILDPATH/build-artifacts/git-phlow-$VERSION-darwin-$GOARCH.tar.gz git-phlow