---
maintainer: groenborg
---

# git phlow 

| integration | build | goreport |
| ------------- | --- | ----------------- |
| ![integration status](https://concourse.kubernetes.praqma.cloud/api/v1/teams/main/pipelines/git-phlow/jobs/integration-tests/badge) | ![build status](https://concourse.kubernetes.praqma.cloud/api/v1/teams/main/pipelines/git-phlow/jobs/create-patch/badge)  |[![Go Report Card](https://goreportcard.com/badge/github.com/Praqma/git-phlow)](https://goreportcard.com/report/github.com/Praqma/git-phlow)       |


git-phlow (pronounced _"git flow"_), is a CLI extension for git, which provides an extra set of commands to easily use our pragmatic workflow called **the phlow**.  It has an automatic branching model, which compliant CI/CD services can use for full automation.  It also provides automatic issue tracking using [github](https://github.com) issues with [waffle](https://waffle.io/).

Read about the pragmatic workflow, **the phlow**, [here](http://www.praqma.com/stories/a-pragmatic-workflow/).

git-phlows core features will include:

- *works with GitHub and waffle*: git-phlow assigns labels to move around user stories on waffle boards, and managing your github issues.

- *branches based on issues*:  work-on will create a branch from a GitHub issues and move it to "wip"

- *issue tracking*: having all your user stories as waffle cards or GitHub issues, git-phlow will mark them appropriately to their status.

- *build for pretest integration*: deliver will push your branch to GitHub, ready for your pipeline to pull, integrate, test and merge back in.

## the-phlow and waffle
To get started using the-phlow, read [praqmas](http://www.praqma.com/stories/a-pragmatic-workflow/) article about how the-phlow and the pragmatic workflow works, and how to configure waffle for your projects.

## installing the phlow
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
#For Mac
wget https://github.com/Praqma/git-phlow/releases/download/{$VERSION}/git-phlow-{$VERSION}-darwin-amd64.tar.gz && tar -xzf git-phlow-{$VERSION}-darwin-amd64.tar.gz && mv git-phlow /usr/local/bin

#For linux
wget https://github.com/Praqma/git-phlow/releases/download/v{$VERSION}/git-phlow-{$VERSION}-linux-amd64.tar.gz && tar -xzf git-phlow-{$VERSION}-linux-amd64.tar.gz && mv git-phlow /usr/local/bin
```

### Using the Phlow
- I go to github and look at the issues or waffle
- I find an issue I want to work on, and note it's id eg. `42`
- Then i go to my local git repository and type: 

```git
git phlow workon 42

#creates branch: 42-some-issue-title
```
- On that branch I start making my changes
- When I'm done I type: 

```git
git phlow wrapup

#commits the changes
```
- now that the changes are ready I type: 

```git
git phlow deliver

# renames branch to: ready/42-some-issue-title
# pushes branch to upstream
```

Then the branch should be automatically tested and eventually merged into master. 
if you do not have an automation system, you can deliver the work to your default branch by doing a local deliver instead. 

```git
git phlow deliver --local

#changes branch to master
#merges changes from 42-some-issue-branch
#pushes changes to github
```

## Project status
This is the official repository for the git-phlow extension. The project is currently under development and many additions are still to come. but version 0.1.5 is available and stable for both linux and Mac. 


## Getting started
**Know the phlow, know go**
To understand the-phlow read praqmas great story about what is is and how it works. [read me, please!](http://www.praqma.com/stories/a-pragmatic-workflow/)

### Contribution
To contribute to the project, I refer to reading our [contribution guide](https://github.com/Praqma/git-phlow/blob/master/CONTRIBUTING.md)
