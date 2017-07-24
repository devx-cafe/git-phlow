#!/bin/bash

mkdir -p $GOPATH/src/github.com/praqma
cp -R git-phlow/ $GOPATH/src/github.com/praqma
cd $GOPATH/src/github.com/praqma/git-phlow

go build ci/cover/cover.go


ls

go get -d -t -v ./...



#Run the coverage
ci/cover/gencover.sh $(go list ./...)

cat coverfiles/percentage
./cover

