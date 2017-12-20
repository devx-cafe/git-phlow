# Documentation

Available commands: 

- [`workon`](#workon)
- [`wrapup`](#wrapup)
- [`deliver`](#deliver)
- [`auth`](#auth)
- [`issues`](#issues)
- [`web`](#web)
- [`mkalias`](#mkalias)
- [`cleanup`](#cleanup)
- [`config` ](#config)
- [`agent`](#agent)
- [global flags](#global-flags)

## Requirements
- Activate issues on the designated GitHub repository
- Create the labels used by git-phlow. Find the labels [here](https://www.praqma.com/stories/a-pragmatic-workflow/)
- Download and [install](/docs/installation.md) git-phlow
- Start creating tasks in GitHub issues


### workon
```
git phlow workon [issue]
```
**flags**

`-t`, `--target` - INI block in .gitconfig

**description**

workon is the most essential command of git phlow. The command will create a new local branch from a remote integration branch, and base the name on an issue from the chosen issue management system [[See configuration](/docs/configuration.md)].
If you run workon with unstaged changes in your workspace those changes will follow to the new branch. If you have just committed changes you will experience that those changes will disappear. This is due to the fact that when the branch is based on the remote integration branch, only changes integrated into the the remote integration branch will be available. Deliver the commits with the `deliver` to have them availbe when you do workon. 

```sh
git status
# On branch master
# Your branch is up to date with 'origin/master'.

git phlow workon 18
# Preparing workspace...
# 
# -------- Issue #18 updated --------  
# Label    => Status - in progress 
# Assignee => @groenborg 
# ----------------------------------
# Created workspace:  18-ai-is-now-conscious 

git status
# On branch 18-ai-is-now-conscious
# Your branch is up to date with 'origin/master'.

```


### wrapup
```
git phlow wrapup
```
**flags**

`-t`, `--target` - INI block in .gitconfig

`--force` - use a custom commit message instead of the auto-generated

**description**

wrapup is used to collect all the changes just made and put them into a commit, that is ready to be delivered into the integration branch. The command will add changes from the workspace to the index/staging area of git, and commit with `smart commit`syntax that will close the mentioned issue when it is integrated on the integration branch. The smart commits only works on systems that support it.
```sh
git status
# On branch 18-ai-is-now-conscious
# Your branch is up to date with 'origin/master'.
# 
# Changes not staged for commit:
#   (use "git add <file>..." to update what will be committed)
#   (use "git checkout -- <file>..." to discard changes in working directory)
# 
# 	modified:   robo-mind.go
# 

git phlow wrapup
# close #18 ai is now conscious

git status
# On branch 18-ai-is-now-conscious
# Your branch is ahead of 'origin/master' by 1 commit.
#   (use "git push" to publish your local commits)
# 
# nothing to commit, working tree clean
```


### deliver
```
git phlow deliver
```
**flags**

`-l`, `--local` - integrates locally and pushes changes to remote integration branch

`-t` , `--target`  - INI block in .gitconfig

**description**

deliver is the command used to hand over the work you just commited with `wrapup`. Checkout the issue branch you need to deliver and run the command. This will push local branch to the repository and prefixed it with `ready/`. On the repository it is ready to be integrated automatically into the integration branch, if a CI server is configured [see workflow](/docs/workflow.md). 
When delivered, local branches will be prefixed with `delivered/`

using the `--local` option will integrate the branch locally and push the changes on the integration branch
```sh
git status
# On branch 18-ai-is-now-conscious
# Your branch is ahead of 'origin/master' by 1 commit.
#   (use "git push" to publish your local commits)
# 
# nothing to commit, working tree clean

git phlow deliver
# delivering...
# Delivered branch 18-ai-is-now-conscious 
```


### auth
```
git phlow auth
```
**flags**

`-t` , `--target`  - INI block in .gitconfig

**description**

auth authorizes git-phlow towards a chosen issue management system like GitHub Issues and Jira. The auth command uses the .gitconfig file to figure out which service it tries to authorizes against. Support issue management systems right now are: Jira, GitHub

### issues
```
git phlow issues
```
 **flags**

`-m` , `--mine` - only list issues assigned to you

`-t` , `--target` - INI block in .gitconfig

**description**

issues lists the first 30 issues from the connected issue management system. The are not picked in any specific order, and for Jira they are not picked by any specific project. The command is helpful for getting an overview of the next issues without leaving your terminal. The issue command uses the configuration `issue-api` key-value pair to fetch the issues.

```sh
git phlow issues
#38: testing windows Status - in progress @kryptag 
#23: add impl for postgres Prio 1 - must have Size 3 - large Status - in progress @groenborg Backlog
#22: added nodejs api node Prio 1 - must have Size 1 - small Backlog
#17: AI now has bank access Action - awaiting feed-back Prio 1 - must have Size 1 - small Status - up next Backlog
#16: limit AI access skrzz buzz Prio 1 - must have Size 1 - small Status - up next Backlog
#14: achieve market goals Prio 1 - must have Size 1 - small Status - in progress Status - up next @groenborg Backlog
#13: added pipeline automation Prio 1 - must have Size 3 - large Backlog

```

### mkalias
```
git phlow mkalias
```

**description**

mkalias creates aliases for the git phlow commands. This allows you to use 'git workon', rather than 'git phlow workon'. The aliases are added to your global .gitconfig file.

### web
```
git phlow web
```
**flags**

`-t` , `--target` - INI block in .gitconfig

**description**

web will open the default browser on the issue webpage. It does so extracts the issue id from the current git branch, and uses the issue-web url of the configuration. If no issue id or Jira key can be extracted, another behavior will be triggered. On Jira the dashboard will be opened. On Github the repository main page will be opened.

### cleanup
```
git phlow cleanup 
```

**flags**

`-f` , `--force` - force remove delivered branches

`-l` , `--local` - only remove local delivered branches

`-t` , `--target` - INI block in .gitconfig

**description**

cleanup is for tidying up the git workspace. As you follow the workflow a lot of branches prefixed with `delivered/` will be in the workspace, and should just be deleted if they have been successfully integrated. Running the command will remove these branches locally and remote as well. A local version of this command is available, and will just remove the local branches.
Some branches can not be deleted, because git cannot detect if they have been integrated into the integration branch. That can be due to a rebase or squash. Those can be deleted with the `force` flag. 

git cannot always detect if a branch have been integrated,so for those you have to run delete with `--force`

### config
```
git phlow config
``` 

**description**

config is the top-level command for showing and creating git-phlow .gitconfig files. Manipulating the individual key-value pairs is done by using git's build in `git config`. See more in [configuration documentation](/docs/configuration.md). 

When using a local `.gitconfig` file, you need to tell git that it should also look for configuration in the local workspace folder by adding: 
`git config --local include.path ../.gitconfig`

#####  show
```
git phlow config show <ini name>
```
**description**

show shows the configuration of the specified INI block. If no arguments are passed it will show the default configuration git phlow will use if no `-t` flag is set when a command is executed. This is good to use when debugging. Look for the `scope: internal`, this means that git phlow is using the internal default block, and not one found in a custom .gitconfig file.  

##### bootstrap
```
git phlow config bootstrap
```

**description**

bootstrap creates a new .gitconfig file in the current directory, with the default values of a git-phlow configuration. This is a good way to get started with the configuration file. 


### agent
```
git phlow agent [cmd] 
```

**flags**

_none_

**description**

The agent commands are not well supported in these later versions, but will be back on the roadmap. The agent commands are commands which makes the tool able to act as a pretested integration setup. 

They are not recommemded to use right now


## Global flags

`-h`,`--help` - get the help text of a command

`-v`,`--verbose` - prints all the executed git commands

`--no-color` - suppresses the colors of the output. Meant to be used on windows in terminals with no support for ANSI colors. 
