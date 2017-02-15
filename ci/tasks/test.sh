#!/bin/bash
set -e -u -x

#SETTING UP BUILD PATH AND FOLDER STRUCTURE
buildpath=$(pwd)

mkdir -p $GOPATH/src/github.com/praqma
git clone https://praqma:${github-token}@github.com/Praqma/git-phlow.git $GOPATH/src/github.com/praqma

#CHANGE TO PROJECT DIRECTORY
cd $GOPATH/src/github.com/praqma/git-phlow

#FINDING READY BRANCHES
readybranch=$(git branch -r --list "origin/ready/*" | tail -1 | sed "s/^[ \t]*//")
readybranch2=${readybranch#*/}

echo -----readyBranch------
echo $readybranch2

if [ "$readybranch2" ]; then
    echo "exit 0 - Ready Branch found!"

else
    echo "exit 1 - No Ready Branches found" > ../results.txt
    exit 1
fi

#Starting PIP
git checkout master
git checkout $readybranch2

echo -----------
echo PULL
git fetch
git branch -a
git pull --rebase origin master

echo -----------
echo CHECKOUT_MASTER
git checkout master

echo -----------
echo SQUASH
git merge --squash $readybranch

# RESOLVE DEPENDENCIES - TEST AND PRODUCTION
go get -d -v ./...
go get github.com/stretchr/testify
go get github.com/smartystreets/goconvey
go get github.com/spf13/cobra

#RUN WHOLE TEST SUITE -VERBOSE
go test -v ./...

#If no errors so far then Git push master to origin
if [ $? -eq 0 ]; then

    #On test success
    #Push master to origin
    git push origin master

fi

#Do this no matter outcome
#Delete the remote ready branch
git push origin :$readybranch2
git fetch --prune
