# git-phlow - the release train extension

| Job | Status  |
| ------------- | ----------------- |
| `integration` | ![build status](https://elementsx.me/api/v1/teams/main/pipelines/git-phlow/jobs/integration/badge)       |

git-phlow (pronounced _"git flow"_), is a CLI extension for git, which is designed to provide an extra set of commands to easily create, wrapup, and deliver tasks/issues created on [github](https://github.com) and [waffle](https://waffle.io/). 

git-phlow is the extension build for [the phlow](http://www.praqma.com/stories/a-pragmatic-workflow/), a praqmatic workflow, simplifying your software delivery. 

Some of git-phlow' core features will include: 

- *works with GitHub and waffle*: git-phlow assigns labels to move around user stories on waffle boards, and managing your github issues.

- *branches based on issues*:  work-on will create a branch from a GitHub issues and move it to "wip"

- *issue tracking*: having all your user stories as waffle cards or GitHub issues, git-phlow will mark them appropriately to their status. 

- *build for pretest integration*: deliver will push your branch to GitHub, ready for your pipeline to pull, integrate, test and merge back in. 


## Project status
This is the official repository for the git-phlow extension. As of now it is currently under heavy development. The foundation of the application still changes daily and test procedures are still a work on. 

## the-phlow and waffle
To get started using the-phlow, read [praqmas](http://www.praqma.com/stories/a-pragmatic-workflow/) article about how the-phlow and the pragmatic workflow works, and how to configure waffle for your projects. 

##Getting started
**Know the phlow, know go**

To understand the-phlow read praqmas great story about what is is and how it works.
it will you understand what git-phlow is and what we are building
[read me, please!](http://www.praqma.com/stories/a-pragmatic-workflow/)

If you haven't written go code before, I recommend to take the [go tour](https://tour.golang.org/welcome/1); it will provide great insight in the go programming language. As a reference i recommend [gobyexample](https://gobyexample.com/)

**Project dependencies**

The project uses other frameworks and libraries, it is recommend to read about them to get an understanding of what they do, provide and how to use them. 

[Cobra](https://github.com/spf13/cobra) is both a library and a CLI, and is the library we use in our CLI. Cobra itself has a CLI, which sets up boiler plate cmd's in the cmd folder. use it by typing `cobra add <name of cmd>`

[GoConvey](https://github.com/smartystreets/goconvey) is the testing framework we use for unit tests, it provides and BDD expressions in your tests, like the nodejs [mocha](https://mochajs.org/) test framework

###Get you shell configured for the-phlow

**Get the alias's**

Clone this repository or just copy the alias' in the .gitconfig file and add it to your own. This will give you 
the necessary git commands to use the-phlow

`git clone git@github.com:Praqma/the-phlow.git`

The needed alias's from the-phlow
```git
git work-on 
git wrapup
git deliver
git local-deliver
```
**get `ghi`**

You also need to have `ghi` installed on your machine, it is a CLI for managing github issues, and we use it in the-phlow alias's. Get it [here](https://github.com/stephencelis/ghi)

###Get the project set up on your machine

Go to [get golang](https://golang.org/doc/install) and get the latest version and install it on your machine. 
it should be available on `brew` aswell.

Next you need to setup your shell for go development, this is also documented on go' website. Export GOPATH and GOBIN, in your shell config, 
eg. `zshrc`if you are using zsh.

```bash
export GOPATH=$HOME/<chooseyourfolder> 
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOBIN
```

Next up is getting our project. Use go's own tool `go get` or clone the repository into your GOPATH under, src/github.com/praqma/
```bash
go get -v github.com/praqma/git-phlow   #downloads the project to you gopath
go get -d                               #resolve dependencies
go get -t                               #should resolve test dependencies
```
###Running tests
Go has a great toolset for running rests. As a standard go runs the test from the current directory you are in, if you want to run tests from subdirectories or just all tests, run:

`go test ./...`

For specific directories, run: 

`go test /folder`


###Building
To build binaries go also comes with a pre installed tool, `go build`

To build and drop a binary in the root directory of your project, run:

`go build  #Build and adds a binary to the root of the project`

To install the binary in the bin folder in your GOPATH, you can run a different command.

`go install`

if you have exported GOBIN you can run the binary directly from the terminal

`git-phlow --help`


###Tools
IntelliJ has a [plugin](http://go-ide.com) for go development - I personally use this (groenborg)

[Atom](https://atom.io) has a ton of packages for go development as well very useful

[gogland](https://www.jetbrains.com/go/) as of february jetbrains just announced an official go ide, it is still in pre release



### Contribution
To contribute to the project, fork it on GitHub. The project is written in go and follow go's own official style and commenting guidelines.  

coding guidelines [effective go](https://golang.org/doc/effective_go.html) 
contributing guidelines: [go official contribute guide](https://golang.org/doc/contribute.html) 

---
Maintainer: groenborg
---
