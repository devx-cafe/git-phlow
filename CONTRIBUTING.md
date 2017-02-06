
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

### Using the Phlow
Since this project is about creating a git extension to replace the current alias based solution, we are in a bootstrap situation. 

As long as the extension is not feature complete and more or less ready for production, you should use the alias based version. All the needed aliases, in the most up to date version, can be found in the following [.gitconfig](https://github.com/Praqma/the-phlow/blob/master/.gitconfig).

### Example workflow   

- I go to github and look at the issues
- I find an issue I want to work on, and note it's id eg. `42`
- Then i go to my local git repository and type: 

```git
git work-on 42

#creates branch: 42-some-issue-title
```
- On that branch I start making my changes
- When I'm done I type: 

```git
git wrapup

#commits the changes
```
- now that the changes are ready I type: 

```git
git deliver

# renames branch to: ready/42-some-issue-title
# pushes branch to upstream
```

Then the branch should be automatically tested and eventually merged into master.

## Finding a task to get you started
When you have the project up and running, check our [github issues](https://github.com/Praqma/git-phlow/issues) and find one you want to work on or fix.

Even though this is still a small fledgeling project, we want to make it easily accessibly to contribute, whether you are an experienced go-lang and git expert or not. 

Some issues are labelled with `Difficulty - easy`. This label indicates that we believe the issue is of a nature where people unfamiliar with the project, unexperienced in go-lang or programming in general or just new to contributing to open source repositories, can easily get started. We might not be there quite yet, be will will try to go that extra mile in providing better descriptions and clear scopes for these issues.

Please note: If you do consider yourself the target audience for these *Easy* issues, be awesome and pick some of the slightly harder tasks. It is tempting to quickly grab and complete the easy ones, but this quickly leaves no obvious tasks for newcomers to tackle. 

Another way that you can help, if you are more experienced in specific areas, is to help **groom** the *Easy* tasks, to make them even easier. Suggest how a given problem might be solved, where in the code base a given bug might probably be found, or mention in the issue that you are willing to provide specific guidance. Any help in grooming the larger issues and splitting them up into beginner friendly chunks is also greatly appreciated.

Most of the people so far involved with the project are all pretty new to Go. If you know little of git or extensions, but are well versed in Go, please keep your eyes open on ongoing issues and offer help where you can. Contributions to an open source project are absolutely not only about getting a code contribution merged. Helping others contribute, writing documentation, grooming tasks, triaging reported bugs and so on are all very valuable contributions and will be greatly appreciated.


