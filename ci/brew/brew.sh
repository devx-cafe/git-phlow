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

#CLONE REPO TO CC OUTPUT
git clone praqma-tap updated-praqma-tap

#GO SCRIPT THAT GENERATES BREW FORMULA
go run git-phlow/ci/brew/template.go > updated-praqma-tap/Formula/git-phlow.rb

#OUTPUT OF FORMULA FOR LOGGING
cat updated-praqma-tap/Formula/git-phlow.rb

cd updated-praqma-tap
git status

#ADD GIT USER
git config --global user.email "concourse@praqma.net"
git config --global user.name "concourse"

#ADD NEW FORMULA AND COMMIT WITH RELEASE
git add --all
git commit -m "$VERSION released"