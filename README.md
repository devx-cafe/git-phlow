---
maintainer: groenborg
---

# git phlow 

| tollgate | build | goreport | coveralls | issues | 
| ------------- | --- | ----------------- | ----- | ----- |
| ![integration status](https://concourse.bosh.praqma.cloud/api/v1/teams/main/pipelines/git-phlow/jobs/integration-test/badge) | ![build status](https://concourse.bosh.praqma.cloud/api/v1/teams/main/pipelines/git-phlow/jobs/shipit/badge)  |[![Go Report Card](https://goreportcard.com/badge/github.com/Praqma/git-phlow)](https://goreportcard.com/report/github.com/Praqma/git-phlow) | [![Coverage Status](https://coveralls.io/repos/github/Praqma/git-phlow/badge.svg?branch=master)](https://coveralls.io/github/Praqma/git-phlow?branch=master) | [![Stories in Ready](https://badge.waffle.io/Praqma/git-phlow.svg?label=ready&title=Ready)](http://waffle.io/Praqma/git-phlow) |


git-phlow (pronounced _"git flow"_), is a CLI extension for git, which provides an extra set of commands to easily use our pragmatic workflow called **the phlow**. It provides a branching model, which makes collaboration. It also provides automatic issue tracking using [github](https://github.com) issues with [waffle](https://waffle.io/). The branching models uses branches prefixed with ready(can be configured), to mark delivered task to the remote repository from where automation services can pick up and integrate the changes into the stable branch. 

Read about the pragmatic workflow, **the phlow**, [here](http://www.praqma.com/stories/a-pragmatic-workflow/).

git-phlows core features will include:

- *works with GitHub and waffle*: git-phlow assigns labels to move around user stories on waffle boards, and managing your github issues.

- *branches based on issues*:  work-on will create a branch from a GitHub issues and move it to "wip"

- *issue tracking*: having all your user stories as waffle cards or GitHub issues, git-phlow will mark them appropriately to their status.

- *build for pretest integration*: deliver will push your branch to GitHub, ready for your pipeline to pull, integrate, test and merge back in.

## Using the phlow (workflow)
To get started using the-phlow, read [praqmas](http://www.praqma.com/stories/a-pragmatic-workflow/) article about how the phlow works, and how to configure waffle for your projects.
A few things nees to be configured before hand: 

- Activate __issues__ on the designated GitHub repository
- Create the __labels__ used by the phlow. Read the article above to find them
- Download and install _git-phlow_ using _brew_ or _wget_
- Start creating tasks in GitHub issues

## installing git-phlow (tool)
### Using Homebrew  
If you are a Mac user and have homebrew installed, you can tap our official repository.

```shell
#Tap our repository
brew tap praqma/praqma-tap

#Install git-phlow
brew install git-phlow
``` 
### Alternative
You can also install git-phlow running the following commands. That will download and install the project.
Remember to use the latest version.  

```shell
#For linux
wget https://github.com/Praqma/git-phlow/releases/download/v1.1.5/git-phlow-1.1.5-linux-amd64.tar.gz && tar -xzf git-phlow-1.1.5-linux-amd64.tar.gz && mv git-phlow /usr/local/bin
```

### Using git Phlow
- Go to GitHub or Waffle and look at the issues
- Find an issue to work on, and note it's id eg. `42`
- Then go to the local git repository and type: 

```git
git phlow workon 42

#creates branch: 42-some-issue-title
```
- Make the changes and complete the task
- When it's done, type: 

```git
git phlow wrapup --hard

#adds everything and commits the changes
#alternately you can add your changes manually and use
git phlow wrapup
```
- Now that the changes are ready, type: 

```git
git phlow deliver
# renames local branch to: delivered/42-some-issue-title
# pushes remote branch named: ready/42-some-issue-title
```

If the automation tools have been configured to listen for __ready/__ branches, they will pick up the branch and integrate them with the stable branch. 
if you do not have an automation system, you can deliver the work to your default branch by doing a local deliver instead. 

```git
git phlow deliver --local

#changes branch to master
#merges changes from 42-some-issue-branch
#pushes changes to github
```
### Automation systems
For automation systems we use, travis CI, Concourse CI and Jenkins. They can all be configured to follow the phlow and listen for ready branches. 
You can see the concourse pipeline configured for this project [here](https://concourse.kubernetes.praqma.cloud/teams/main/pipelines/git-phlow) 

### git-phlow on Windows

We recommend running git-phlow using _PowerShell_.
There are known issues with _cmd_ rendering formatted text incorrectly, and _git bash_'s input being handled incorrectly.

## Project status
This is the official repository for the git-phlow extension. The project is currently under development and many additions are still to come. but version 0.1.5 is available and stable for both linux and Mac. 


## Getting started
**Know the phlow, know go**
To understand the-phlow read praqmas great story about what is is and how it works. [read me, please!](http://www.praqma.com/stories/a-pragmatic-workflow/)

### Contribution
To contribute to the project, I refer to reading our [contribution guide](https://github.com/Praqma/git-phlow/blob/master/CONTRIBUTING.md)

### CodeScene Analysis
[![](https://codescene.io/projects/1345/status.svg) Get more details at **codescene.io**.](https://codescene.io/projects/1345/jobs/latest-successful/results)
