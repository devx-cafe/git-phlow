
## How to Contribute

Hi we are thrilled that you are interested in Contributing to this project. Good contributions keeps this project great!
This document outlines how you should contribute and helps make it easier to get your contributions accepted. 


coding guidelines [effective go](https://golang.org/doc/effective_go.html) 
contributing guidelines: [go official contribute guide](https://golang.org/doc/contribute.html) 

### Getting started
- clone the repository
- read godocs and guidelines for writing gocode 
- run the project and get to know it
- `work on`an issue and use `deliver` to deliver your branch and make it available for review

### Using the phlow
When working on git-phlow you should use our own git alias's until git-phlow is ready for production.
When you have the project up and runnig, check our github issues and find the one you want to work on or fix

Some issues are labelled with `Difficulty - easy` those issues are marked like that, so people unfamiliar with the project, writing go code or 
contributing to open source repositories, have something they can work on to begin with. 
Please note: If you do not qualify for these criterias, be awesome and pick something else, so others can take them. 


### Example workflow   

- I go to github and look at the issues
- I find the issue i want to work on, and note it's id eg. `42`
- Then i go to my local git repository and type: 

```git
git work-on 42

#creates branch: 42-some-issue-title
```
- On that branch i start making my changes
- When i'm done i type: 

```git
git wrapup

#commits the changes
```
- now that the changes are ready i type: 

```git
git deliver

# renames branch to: ready/42-some-issue-title
# pushes branch to upstream
```

Then the branch will be reviewd and eventually merged into master.
