$currentPath = Get-Location

$GOPATH = $currentPath.Path
[Environment]::SetEnvironmentVariable("GOPATH", $GOPATH, "Machine")

$GOBIN = "$($currentPath.Path)/bin"
[System.Environment]::SetEnvironmentVariable('GOBIN', $GOBIN, "Machine")

$GITHUB_FAKE = "$($currentPath.Path)/src/github.com/code-cafe/git-phlow/build"
[System.Environment]::SetEnvironmentVariable('GITHUB_FAKE', $GITHUB_FAKE, "Machine")

$TEST_REPO = "$GOPATH/src/github.com/code-cafe/git-phlow/build/github"

New-Item -ItemType directory -Path $TEST_REPO
Set-Location $TEST_REPO

ls

function Create-Inital-Commit () {
    git init
    git config user.email "john@doe.com"
    git config user.name "johndoe"
    git config core.autocrlf true
    git add .
    git commit --allow-empty -m "initial commit"
}

function Two-Commits-On-Master () {
    "hello" > README.md
    Add-All-And-Commit "1 commit - changes to readme"
     "file1" > file1.txt
    Add-All-And-Commit "2 commit - changes to file1"
     "#!/bin/bash  \n exit 0" > test.sh
     "#!/bin/bash  \n exit 1" > testerr.sh
    Add-All-And-Commit "3 commit - test script added"
}

function Branch-Foo-Additions (){
    git checkout -b  foo
    "hello world to README" > README.md
    "on branch foo" > file1.txt
    Add-All-And-Commit "changes on branch foo"
}


function Branch-Master-Additions () {
    git checkout master
    "hello world to README from master" > README.md
    Add-All-And-Commit "changes on branch master"
}


function Branch-Issue-Additions () {
    git checkout -b 11-issue-bar
    "changes in file 2" > file2.txt
    "changes in file 1" > file1.txt
    Add-All-And-Commit "changes on branch 11-issue"
}

function Return-To-Master () {
    git checkout master
}


function Create-Delivered-Branches () {
    git checkout -b delivered/42-issue-branch
    "on branch delivered/42-issue-branch" > README.md
    Add-All-And-Commit "delivered/42 branch commit"
    Return-To-Master
    git checkout -b delivered/24-issue-branch
    "on branch delivered/24-issue-branch" > README.md
    Add-All-And-Commit "delivered/24 branch commit"
}

function Create-Ready-Branches () {
    git checkout -b ready/15-issue-branch
    "on branch ready/15-issue-branch" > README.md
    Add-All-And-Commit "ready/16 branch commit"
    Return-To-master
    git checkout -b ready/99-issue-branch
    "on branch ready/99-issue-branch" > README.md
    Add-All-And-Commit "ready/99 branch commit"
}

function Add-All-And-Commit ($message){
    git add .
    git commit -m $message
}

function Create-Origin () {
    Write-Host $GITHUB_FAKE
    git clone $GITHUB_FAKE/github $GITHUB_FAKE/phlow-test-pkg
    Set-Location $GITHUB_FAKE/phlow-test-pkg
    git branch bar
    git branch delivered/1-issue-branch
}


#ACTUAL SCRIPT
Write-Host "CREATING TEST REPOSITORY"
Create-Inital-Commit
Two-Commits-On-master
Branch-Foo-Additions
Branch-Master-Additions
Branch-Issue-Additions
Return-To-Master
Create-Delivered-Branches
Return-To-Master
Create-Ready-Branches
Return-To-Master
Create-Origin

Write-Host "WRAPPING UP JOB"
