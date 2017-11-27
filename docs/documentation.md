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

## requirements
- Activate issues on the designated GitHub repository
- Create the labels used by the phlow. Find the labels [here](https://www.praqma.com/stories/a-pragmatic-workflow/)
- Download and install git-phlow
- Start creating tasks in GitHub issues


### workon
```
git phlow workon [issue]
```
**flags**

`-t`, `--target` - INI block in .gitconfig

**desc**

Creates a new local branch from a remote integration branch, based on an issue from the chosen issue management system [See configuration]. 

### wrapup
```
git phlow wrapup
```
**flags**

`-t`, `--target` - INI block in .gitconfig

`--force` - use a custom commit message instead of the auto-generated

**desc**

Add changes from the workspace to the index/staging area and commits with `smart commit`syntax that will close the mentioned issue when it is integrated on the integration branch.

### deliver
```
git phlow deliver
```
**flags**

`-l`, `--local` - integrates locally and pushes changes to remote integration branch

`-t` , `--target`  - INI block in .gitconfig

**desc**

deliver pushes a branch defaultly prefixed with `ready/` to the repository, where it is ready to be integrated automatically into the integration branch, if a CI server is configured [see workflow]. 
local branch will be prefixed with `delivered/`

### auth
```
git phlow auth
```
**flags**

`-t` , `--target`  - INI block in .gitconfig

**desc**

To connect with the issue management systems, it is required to authenticate with a valid account. GitHub requires a personal user.  Jira requires an account, probably created by the administrator. 

### issues
```
git phlow issues
```
 **flags**

`-m` , `--mine` - only list issues assigned to you

`-t` , `--target` - INI block in .gitconfig

**desc**

list the first 30 issues connected to the repository. This is only valid for GitHub. Issues on Jira will be taken from every project. 

### mkalias
```
git phlow mkalias
```

**desc**

creates git alias' for phlow commands, so it is no longer necessary to write `phlow` before every command. e.g. `git workon 42`

### web
```
git phlow web
```
**flags**

`-t` , `--target` - INI block in .gitconfig

**desc**

extracts the issue id from the current git branch, and opens the default browser on that issue. If no issue id or Jira key can be extracted, another behaviour will be triggered. On Jira the dashboard will be opened. On Github, the repository main page will be opened.

### cleanup
```
git phlow cleanup 
```

**flags**

`-f` , `--force` - force remove delivered branches

`-l` , `--local` - only remove local delivered branches

`-t` , `--target` - INI block in .gitconfig

**desc**

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
**desc**

shows the configuration of the specified INI block. If no arguments are passed it will show the default configuration git phlow will use if no `-t` flag is set when a command is executed. This is good to use when debugging. Look for the `scope: internal`, this means that git phlow is using the internal default block, and not one found in a custom .gitconfig file. 

remember to add this `git config --local include.path ../.gitconfig` if you are using a config file in your repo. 

#### boostrap
```
git phlow config bootstrap
```

**desc**

create a new .gitconfig file in the current directory, with the default values of a git phlow configuration. A good way to get started with the configuration file. 

## Global flags

`-h`,`--help` - get the help text of a command

`-v`,`--verbose` - prints all the executed git commands

`--no-color` - suppresses the colors of the output. Meant to be used on windows in terminals with no support for ANSI colors. 