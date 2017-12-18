---
maintainer: groenborg
---

# git phlow

| tollgate | build | goreport | coveralls |  
| ------------- | --- | ----------------- | ----- |
| ![integration status](https://concourse.bosh.praqma.cloud/api/v1/teams/main/pipelines/git-phlow/jobs/checkin/badge) | ![build status](https://concourse.bosh.praqma.cloud/api/v1/teams/main/pipelines/git-phlow/jobs/takeoff/badge)  |[![Go Report Card](https://goreportcard.com/badge/github.com/Praqma/git-phlow)](https://goreportcard.com/report/github.com/Praqma/git-phlow) | [![Coverage Status](https://coveralls.io/repos/github/Praqma/git-phlow/badge.svg?branch=master)](https://coveralls.io/github/Praqma/git-phlow?branch=master) 

Git-phlow (pronounced _"git flow"_), is a CLI extension for git, which provides an extra set of commands to easily use our pragmatic workflow by the same name, **Git phlow**. It provides a branching model which makes collaboration easy. It also provides automatic issue tracking using [GitHub](https://github.com) issues or Jira.

git-phlows core features include:

- *works with GitHub and waffle*: git-phlow assigns labels to move around user stories on waffle boards, and managing your github issues.

- *works with jira*: sets assignee and does transition

- *branches based on issues*: create a workspace from an issue to ensure traceability

- *fully automatable*: `deliver` your branches, ready for your pipeline to integrate, test and merge

## Getting started
To read the concepts behind git-phlow read [praqma's article](https://www.praqma.com/stories/git-phlow/) about the workflow.
Consult the [official documentation](/docs/README.md) for information on how to get started using git-phlow

### Automation systems
We use Travis CI, Concourse CI and Jenkins. They can all be configured to follow git phlow. You can see the Concourse pipeline configured for Git-phlow [here](https://concourse.bosh.praqma.cloud/teams/main/pipelines/git-phlow)

### git-phlow on Windows

git-phlow works in _PowerShell_ on windows.
There are known issues with _cmd_ rendering formatted text incorrect, and _git bash_'s input being handled incorrectly.

### Project status
This is the official repository for the git-phlow extension. The project is currently stable at version 3.8.2 for both windows, macOS and linux. 

### Contribution
Contributions are welcome! Please read the [contribution guide](https://github.com/Praqma/git-phlow/blob/master/CONTRIBUTING.md)


### Core Team
| [@groenborg](https://github.com/groenborg) | [@MichaelMadsen](https://github.com/MichaelMadsen)
|---|---|
| [![](https://avatars1.githubusercontent.com/u/5576954?s=100&v=4)](https://github.com/groenborg) | [![](https://avatars3.githubusercontent.com/u/657289?s=100&v=4)](https://github.com/MichaelMadsen)


### CodeScene Analysis
[![](https://codescene.io/projects/1345/status.svg) Get more details at **codescene.io**.](https://codescene.io/projects/1345/jobs/latest-successful/results)

