#!/bin/sh
set -e -u -x

#LIST RESOURCES IN ENTRY FOLDER
#FOR DEBUGGING PURPOSES
ls -lah

#GET VERSION AND PATH
#STORE VERSION AND PATH FOR BUILD
VERSION=$(cat resource-version/version)
BUILDPATH=$(pwd)

#CREATE GO DIRECTORY STRUCTURE
#THE STRUCTURE IS NECESSARY FOR GO TOOLS OTHERWISE 
# 'build' AND 'get' WILL FAIL
mkdir -p $GOPATH/src/github.com/praqma
cp -R git-phlow/ $GOPATH/src/github.com/praqma

#NAVIGATE TO FOLDER AND GET DEPS
cd $GOPATH/src/github.com/praqma/git-phlow
go get -d -t -v ./...

#BUILD BINARY
#GENEREATE ONE BINARY FOR DARWIN AND LINUX
go build -ldflags "-X   github.com/praqma/git-phlow/options.Version=`echo $VERSION` -X  github.com/praqma/git-phlow/options.Sha1=`git rev-parse HEAD` -X  github.com/praqma/git-phlow/options.Date=`date +'%d-%m-%Y'`"

#GENERATE A COMPRESSED TAR FILE FOR PLATFORMS LINUX AND DARWIN
tar -cvzf $BUILDPATH/build-artifacts/git-phlow-$VERSION-darwin.tar.gz git-phlow
tar -cvzf $BUILDPATH/build-artifacts/git-phlow-$VERSION-linux.tar.gz git-phlow

cd $BUILDPATH/build-artifacts



