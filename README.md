# git-phlow - the release train extension



![build status](https://elementsx.me/api/v1/teams/main/pipelines/git-phlow/jobs/integration/badge)
 
git-phlow (pronounced _"git flow"_), is a CLI extension for git, which is designed to provide an extra set of commands to easily create, wrapup, and deliver tasks/issues created on [github](https://github.com) and [waffle](https://waffle.io/). 

git-phlow is the extension build for [the phlow](http://www.praqma.com/stories/a-pragmatic-workflow/), a praqmatic workflow, simplifying your software delivery. 

Some of git-phlow' core features will include: 

- *works with GitHub and waffle*: git-phlow assigns labels to move around user stories on waffle boards, and managing your github issues.

- *branches based on issues*:  work-on will create a branch from a GitHub issues and move it to "wip"

- *issue tracking*: having all your user stories as waffle cards or GitHub issues, git-phlow will mark them appropriately to their status. 

- *build for pretest integration*: deliver will push your branch to GitHub, ready for your pipeline to pull, integrate, test and merge back in. 


### Project status
This is the official repository for the git-phlow extension. As of now it is currently under heavy development. The foundation of the application still changes daily and test procedures are still a work on. 

### Get started with the-phlow and waffle
To follow the the-phlow, please read this great article about how the-phlow works and how to get going

### Contribution
To contribute to the project, fork it on GitHub. The project is written in go and follow go's own official style and commenting guidelines.  

coding guidelines [effective go](https://golang.org/doc/effective_go.html) 

contributing guidelines: [go official contribute guide](https://golang.org/doc/contribute.html) 

---
Maintainer: groenborg
---