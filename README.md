<p align="center"><img src="/docs/images/gitphlow.svg" width="360"></p>

![Test-integration - Build - Test - Deliver](https://github.com/code-cafes/git-phlow/workflows/Test-integration%20-%20Build%20-%20Test%20-%20Deliver/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/code-cafes/git-phlow)](https://goreportcard.com/report/github.com/code-cafes/git-phlow)

# git phlow
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
Contributions are welcome! Please read the [contribution guide](https://github.com/code-cafe/git-phlow/blob/master/CONTRIBUTING.md)


