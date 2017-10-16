# Documentation
### The workflow
The Praqma workflow - often called The Phlow - is a combination of the Git extension and CI server for automation. In this the CI server is concourse with the git-phlow resource installed.

![image](/docs/images/phlowmodel.svg)

The top diagram Shows the steps the entire workflow entails.
1. You do `git phlow workon` on an issue which has been groomed and is ready to become a feature in your project.
2. You do `git phlow wrapup` and `git phlow deliver`, to deliver your changes to the remote repository. The branches will be prefixed with `delivery-branch-prefix` from the .gitconfig file.
3. In the concourse Pipeline we are using the git-phlow tollgate resource, that automate the integration if it passes our toll-gate. It integrate branches prefixed with ready, meaning the job will trigger on such branches and merge the changes with the head of the integration.
4. Now the tollgate tests will run in the job and if they pass, the already merged changes will be pushed to the integration branch. The ready branch will be deleted if everything goes well so it leaves things nice and clean. 


### phlow workon
```
git phlow workon <issue>
```
First step in the workflow is to start work on an issue. Run the command and it will create a local workspace named after the title of the issue.

### phlow wrapup
```
git phlow wrapup
```
Run wrapup to add your changes and make a commit with the `smart commit`syntax that will close an issue when it lands on the master branch.

### phlow deliver
```
git phlow deliver
```
Deliver rebases the latests changes from it's origin and pushes a `ready/name` branch to the remote repository and renames the local branch to `delivered/name`. Then it checks out the integration branch e.g. `master`. 

the `ready` branch will then be availabble for the CI server to integrate it into the integration branch. 

If you have no CI server configured to do automatic integration of ready branches you can run a local deliver in stead, and let your local git do the integration
```
git phlow deliver --local
```

### phlow auth
```
git phlow auth
```
Git phlow auth, authorizes towards a chosen service which can be either `github` or `jira`. Auth is used by git phlow to access and manage issues in your chosen provider.

#### Self hosted services or Jira
To authenticate with your own hosted service or Jira you _must_ provide a valid configuration (see **Configuration** section). Your configuration must contain a base url to where your service is hosted and the name of the service. When authenticating set your `--target` to the name of the configuration e.g.

### Configuration
git phlow is configured via `.gitconfig`, in your local workspace, this means that git phlow usese git own internal config and can therefore be placed anywhere git can read it e.g. global .gitconfig in you home folder.

To get started with a a `.gitconfig` in your local repository you must first let git know that the filepath must be included, run this command to do so:
```sh
git config --local include.path ../.gitconfig
```

##### Config Block
Each configuration is its own block. You can have as many configuration blocks as you want, just as long as they all have the necessary fields. To use a specific block you run a command with the `-t or --target` option with the name of the block. git phlow by default looks for a block named `[phlow]`. So if you name your preferred blovk `phlow` you do not need to use the `-t` flag.

A block consists of following fields.

- `integration-branch` The name of your default integration branch. normally it's master
- `remote` The name of your remote. By default git sets this to `origin`
- `service` The service you use for issues.
- `issue-api` The url to your api where issues are hosted and managed
- `issue-web` The url to your webpage where issues are hosted and managed
- `delivery-branch-prefix` The prefix you want your pretested integration system to look for

##### Supported services
-  `github`
-  `jira`

##### Example Configuration
```ini
[phlow]
    remote                 = myfork
    service                = github
    integration-branch     = superintegration
    issue-api              = https://api.github.com
    issue-web              = https://github.com
    delivery-branch-prefix = integrate-me-please
```

**important**: When entering the url of your issue service providers, it is important that only the base url is provided, which means no `www.mygithubinstance.com:8080/` but `www.mygithubinstance.com:8080`

##### Internal Default
If no  `[phlow]`  blocks are located within the git config, git phlow will use an interal default configuration. The internal configuration is the last resort, and it is strongly recommended to use `.gitconfig` to keep your projects configuration, even though it mimics the default.

```ini
[phlow]
    remote                 = origin
    service                = github
    integration-branch     = master
    issue-api              = https://api.github.com
    issue-web              = https://github.com
    delivery-branch-prefix = ready
```

## Trouble shooting

This section is meant to provide guidance to recover from the most known issues you might run into. 

### Did `deliver` but meant `deliver --local`
read [deliver]() docs to understand the command
##### Experience
the work is gone and you cannot deliver the work again.

##### Issue
Deliver pushes the work to the remote repository on a `ready/name` branch, so it can be integrated if a server have been configured. 
phlow will not deliver branches prefixed with `delivered` so you cannot re-deliver it. 

##### Solution
1. rename the delivered branch
2. checkout the branch
3. deliver your work again, nothing should stop you. 
4. delete the remote branch

```sh
git branch -m delivered/10-my-issue 10-my-issue #1
git checkout 10-my-issue    #2
git phlow deliver --local   #3
git push -d origin ready/10-my-issue #4
```

### Did `deliver` before `wrapup`
read [deliver]() docs to understand the command
##### Experience 
The workspace is still dirty and your changes have not been added. 
##### Issue
`deliver`pushes a branch to GitHub but it is empty because you work was not comitted. phlow will not deliver branches prefixed with `delivered` so you cannot re-deliver it. 

##### Solution
1. rename delivred branch
2. checkout the branch
3. wrapup the work
4. delete the remote `ready` branch
5. deliver again
```
git branch -m delivered/10-my-issue 10-my-issue #1
git checkout 10-my-issue    #2
git phlow wrapup            #3
git push -d origin ready/10-my-issue #4
git phlow deliver
```

### Work disappeared after deliver
Have no fear, the work is not gone. the `deliver` command switches to the integration branch e.g. `master` after it is done pushing the `ready` branch to the remote repository.
`master` does not contain the changes that was made on the oter branch. If you try to `fetch` or `pull`and nothing happens then it is probably because the CI server has not integrated the `ready` branch 

##### Getting back to your work
1. A workon _after_ the CI have done the integration should give back your work.
2. Check out the branch locally will also show your changes, but your are breaking the workflow if you continue the work on the same branch. 
