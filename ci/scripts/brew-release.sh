#!/bin/sh


#GET VERSION FROM SEMVER
export VERSION=$(cat gp-version/version)

#GENERIC GH-RELEASE URL WITH A STRING INTERPOLATION
export URL=https://github.com/Praqma/git-phlow/releases/download/v$VERSION/git-phlow-$VERSION-darwin-amd64.tar.gz

#HASH OF BINARY ARTIFACT
export HASH=$(shasum -a 256 phlow-artifact-darwin-s3/git-phlow-$VERSION-darwin-amd64.tar.gz | awk '{ print $1 }')    

#PRINT BUILD INFO
echo $HASH
echo $VERSION
echo $URL

go run template.go > praqma-tap/Formula/git-phlow.rb

#FOR debugging purposes
cat praqma-tap/Formula/git-phlow.rb








