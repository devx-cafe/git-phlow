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

Creates a new local branch from a remote integration branch, based on an issue from the chosen issue management system [[See configuration](/docs/configuration.md)]. 

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

Add changes from the workspace to the index/staging area and commits with `smart commit`syntax that will close the mentioned issue when it is integrated on the integration branch.
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

deliver pushes a branch prefixed with `ready/` to the repository, where it is ready to be integrated automatically into the integration branch, if a CI server is configured [see workflow]. 
local branch will be prefixed with `delivered/`
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

To connect with the issue management systems, it is required to authenticate with a valid account. GitHub requires a personal user.  Jira requires an account, probably created by the administrator. 

### issues
```
git phlow issues
```
 **flags**

`-m` , `--mine` - only list issues assigned to you

`-t` , `--target` - INI block in .gitconfig

**description**

list the first 30 issues connected to the repository. This is only valid for GitHub. Issues on Jira will be taken from every project. 
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

creates git alias' for phlow commands, so it is no longer necessary to write `phlow` before every command. e.g. `git workon 42`

### web
```
git phlow web
```
**flags**

`-t` , `--target` - INI block in .gitconfig

**description**

extracts the issue id from the current git branch, and opens the default browser on that issue. If no issue id or Jira key can be extracted, another behaviour will be triggered. On Jira the dashboard will be opened. On Github, the repository main page will be opened.

### cleanup
```
git phlow cleanup 
```

**flags**

`-f` , `--force` - force remove delivered branches

`-l` , `--local` - only remove local delivered branches

`-t` , `--target` - INI block in .gitconfig

**description**

cleanup will remove the branches prefixed with `delivered/` locally and remove the remote branches as well. A local version is available, and will just remove the local branches.
Some branches can not be deleted, because git cannot detect if they have been integrated into the integration branch. Those can be deleted with the `force` flag. 

### config
```
git phlow config
``` 

####  show
```
git phlow config show <ini name>
```
**description**

shows the configuration of the specified INI block. If no arguments are passed it will show the default configuration git phlow will use if no `-t` flag is set when a command is executed. This is good to use when debugging. Look for the `scope: internal`, this means that git phlow is using the internal default block, and not one found in a custom .gitconfig file. 

remember to add this `git config --local include.path ../.gitconfig` if you are using a config file in your repo. 

#### boostrap
```
git phlow config bootstrap
```

**description**

create a new .gitconfig file in the current directory, with the default values of a git phlow configuration. A good way to get started with the configuration file. 

## Global flags

`-h`,`--help` - get the help text of a command

`-v`,`--verbose` - prints all the executed git commands

`--no-color` - suppresses the colors of the output. Meant to be used on windows in terminals with no support for ANSI colors. 