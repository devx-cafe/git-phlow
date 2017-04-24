
## How to Contribute

Hi. We are thrilled that you are interested in Contributing to this project. Good contributions keep this project great!
This document outlines how you should contribute and helps make it easier to get your contributions accepted. 


Coding guidelines [effective go](https://golang.org/doc/effective_go.html) 

### Getting started
- clone the repository
- read godocs and guidelines for writing gocode 
- run the project and get to know it
- `work on` an issue in a feature branch, then `deliver` your branch and make it available for review.

## Workflow options
In general, this project runs according to the workflow described in Praqma's blogpost [A pragmatic workflow](http://www.praqma.com/stories/a-pragmatic-workflow/). We call this workflow and branching model "The Phlow".

The super short summary of this workflow is:
- create a feature branch for your issue. E.g. feature/42-some-issue-title
- work on this until done
- push the branch as ready/42-some-issue-title
- The build server will then automatically see this branch, and integrate it into master if the build passes.

This means that there are two ways of submitting contributions, depending on whether you have been added as a collaborator to the GitHub repository or not.

As a collaborator you can just follow the phlow as described. We recommend frequent discussions in the comments section of the issue in question, as this documents thoughts and decisions. We also strongly encourage pair programming, but ask that you at least summarize overall decisions in the issue.

If you are not a collaborator (yet), we suggest that you follow the typical GitHub suggested workflow summarized as follows:
- Fork the project to your own github account
- Create a feature branch as above
- When done, push the branch to your own repo
- Send a pull request against the original repository.
- One of the maintainers will then review your contribution and if it is accepted we will simply use the described automated Phlow to integrate it.

## Finding a task to get you started
When you have the project up and running, check our [github issues](https://github.com/Praqma/git-phlow/issues) and find one you want to work on or fix.

Even though this is still a small fledgeling project, we want to make it easily accessibly to contribute, whether you are an experienced go-lang and git expert or not. 

Some issues are labelled with `Difficulty - easy`. This label indicates that we believe the issue is of a nature where people unfamiliar with the project, unexperienced in go-lang or programming in general or just new to contributing to open source repositories, can easily get started. We might not be there quite yet, be will will try to go that extra mile in providing better descriptions and clear scopes for these issues.

Please note: If you do consider yourself the target audience for these *Easy* issues, be awesome and pick some of the slightly harder tasks. It is tempting to quickly grab and complete the easy ones, but this quickly leaves no obvious tasks for newcomers to tackle. 

Another way that you can help, if you are more experienced in specific areas, is to help **groom** the *Easy* tasks, to make them even easier. Suggest how a given problem might be solved, where in the code base a given bug might probably be found, or mention in the issue that you are willing to provide specific guidance. Any help in grooming the larger issues and splitting them up into beginner friendly chunks is also greatly appreciated.

Most of the people so far involved with the project are all pretty new to Go. If you know little of git or extensions, but are well versed in Go, please keep your eyes open on ongoing issues and offer help where you can. Contributions to an open source project are absolutely not only about getting a code contribution merged. Helping others contribute, writing documentation, grooming tasks, triaging reported bugs and so on are all very valuable contributions and will be greatly appreciated.


To understand the-phlow read praqmas great story about what is is and how it works.
it will you understand what git-phlow is and what we are building
[read me, please!](http://www.praqma.com/stories/a-pragmatic-workflow/)

If you haven't written go code before, I recommend to take the [go tour](https://tour.golang.org/welcome/1); it will provide great insight in the go programming language. As a reference i recommend [gobyexample](https://gobyexample.com/)

**Project dependencies**

The project uses other frameworks and libraries, it is recommend to read about them to get an understanding of what they do, provide and how to use them.

[Cobra](https://github.com/spf13/cobra) is both a library and a CLI, and is the library we use in our CLI. Cobra itself has a CLI, which sets up boiler plate cmd's in the cmd folder. use it by typing `cobra add <name of cmd>`

[GoConvey](https://github.com/smartystreets/goconvey) is the testing framework we use for unit tests, it provides and BDD expressions in your tests, like the nodejs [mocha](https://mochajs.org/) test framework

### Get you shell configured for the-phlow

**Get the latest release**
Download the latest release for your platform and install it. You can follow the guidelines on the README page.

### Get the project set up on your machine

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

#go to the project folder /src/github.com/praqma/git-phlow and enter following command

go get -d -t -v ./...		#resolves all dependencies                               
```
### Running tests
Go has a great toolset for running rests. As a standard go runs the test from the current directory you are in, if you want to run tests from subdirectories or just all tests, run:

`go test -v -p 1 ./...  # -p 1 makes sure to run test in one thread. Important for the project test fixture` 

For specific directories, run:

`go test /folder`


### Building
To build binaries go also comes with a pre installed tool, `go build`

To build and drop a binary in the root directory of your project, run:

`go build  #Build and adds a binary to the root of the project`

To install the binary in the bin folder in your GOPATH, you can run a different command.

`go install`

if you have exported GOBIN you can run the binary directly from the terminal

`git phlow --help`


### Tools
IntelliJ has a [plugin](http://go-ide.com) for go development - I personally use this (groenborg)

[Atom](https://atom.io) has a ton of packages for go development as well very useful

[gogland](https://www.jetbrains.com/go/) as of february jetbrains just announced an official go ide, it is still in pre release
