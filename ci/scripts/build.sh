#!/bin/sh
set -e -u -x

#LIST RESOURCES
ls -lah

#GET VERSION AND PATH
VERSION=$(cat resource-version/version)
BUILDPATH=$(pwd)

#CREATE GO DIRECTORY STRUCTURE
mkdir -p $GOPATH/src/github.com/praqma
cp -R git-phlow/ $GOPATH/src/github.com/praqma

#NAVIGATE TO FOLDER AND GET DEPS
cd $GOPATH/src/github.com/praqma/git-phlow
go get -d -t -v ./...

#BUILD BINARY
go build -ldflags "-X main.version=`echo $VERSION` -X main.sha1=`git rev-parse HEAD` -X main.date=`date +'%d-%m-%Y'`"


tar -cvzf $BUILDPATH/build-artifacts/git-phlow-$VERSION-darwin.tar.gz git-phlow
tar -cvzf $BUILDPATH/build-artifacts/git-phlow-$VERSION-linux.tar.gz git-phlow

cd $BUILDPATH/build-artifacts
ls
pwd