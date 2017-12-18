# Trouble shooting

git-phlow can be difficult to use if you are not very familiar with the workings of git, therefore have we created a section that will help you troubleshoot some of the more difficult situations to get in. 

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
git deliver # <- MISTAKE

git status
# On branch master
# Your branch is up to date with 'origin/master'.
#
# nothing to commit, working tree clean

git branch
#   delivered/18-ai-is-now-conscious            37b0800 [ahead 1] close #18 ai is now conscious
# * master                                      e5c808d close #12 added some cleverness
#   remotes/origin/HEAD                         -> origin/master
#   remotes/origin/master                       e5c808d close #12 added some cleverness
#   remotes/origin/ready/18-ai-is-now-conscious 37b0800 close #18 ai is now conscious


git branch -m delivered/18-ai-is-now-conscious 18-ai-is-now-conscious #1
git checkout 18-ai-is-now-conscious    #2
git phlow deliver --local   #3
git push -d origin ready/18-ai-is-now-conscious #4

git status
# On branch master
# Your branch is up to date with 'origin/master'.

# nothing to commit, working tree clean


git branch
#   delivered/18-ai-is-now-conscious          37b0800 close #18 ai is now conscious
# * master                                    37b0800 close #18 ai is now conscious
#   remotes/origin/HEAD                       -> origin/master
#   remotes/origin/master                     37b0800 close #18 ai is now conscious
```

### Did `deliver` before `wrapup`
read [deliver]() docs to understand the command
##### Experience 
The workspace is still dirty and your changes have not been added. 
##### Issue
`deliver`pushes a branch to GitHub but it is empty because you work was not comitted. phlow will not deliver branches prefixed with `delivered` so you cannot re-deliver it. 

##### Solution
1. rename delivered branch
2. checkout the branch
3. wrapup the work
4. delete the remote `ready` branch
5. deliver again

```sh
git status
# On branch master
# Your branch is up to date with 'origin/master'.
#
# Untracked files:
#   (use "git add <file>..." to include in what will be committed)
#
# 	factorybuilder.go
#
# nothing added to commit but untracked files present (use "git add" to track)

git branch
#   delivered/22-added-nodejs-api-node            37b0800 close #18 ai is now conscious
# * master                                        37b0800 close #18 ai is now conscious
#   remotes/origin/HEAD                           -> origin/master
#   remotes/origin/master                         37b0800 close #18 ai is now conscious
#   remotes/origin/ready/22-added-nodejs-api-node 37b0800 close #18 ai is now conscious


git branch -m delivered/22-added-nodejs-api-node #1
git checkout 22-added-nodejs-api-node    #2
git phlow wrapup            #3
git push -d origin ready/22-added-nodejs-api-node #4
git phlow deliver


git status
# On master
# Your branch is up to date with 'origin/master'.
#
# nothing to commit, working tree clean

```

### Work disappeared after deliver
Have no fear, the work is not gone. the `deliver` command switches to the integration branch e.g. `master` after it is done pushing the `ready` branch to the remote repository.
`master` does not contain the changes that was made on the other branch. If you try to `fetch` or `pull`and nothing happens then it is probably because the CI server has not integrated the `ready` branch 



```sh
git branch
# * 11-improved-api-performance-kpi               e5c808d [ahead 1] close #11 improved api performance kpi    <- YOU ARE HERE
#   master                                        37b0800 close #18 ai is now conscious
#   remotes/origin/HEAD                           -> origin/master
#   remotes/origin/master                         37b0800 close #18 ai is now conscious

git deliver
# delivering...
# Delivered branch 11-improved-api-performance-kpi 

git branch
#   delivered/11-improved-api-performance-kpi     e5c808d [ahead 1] close #11 improved api performance kpi    <- THE WORK IS HERE
# * master                                        37b0800 close #18 ai is now conscious                       <- YOU ARE HERE
#   remotes/origin/HEAD                           -> origin/master
#   remotes/origin/master                         37b0800 close #18 ai is now conscious
#   remotes/origin/ready/22-added-nodejs-api-node 37b0800 close #11 improved api performance kpi


```

##### Getting back to your work
1. A workon _after_ the CI have done the integration should give back your work.
2. Check out the branch locally will also show your changes, but your are breaking the workflow if you continue the work on the same branch. 


### Failed build

##### Experience 
The issue branch was delivered to the remote repository, but it the CI server failed the integration and now it needs to be fixed manually.
##### Issue
If the build fails there is no automatic way to recover, so manual steps are needed. The build can fail for many reasons, unmergable changes, failing tests, etc. 

##### Solution
1. rename delivered branch
2. checkout the branch
3. pull and merge, or pull rebase
4. delete the remote `ready` branch
5. deliver again

```sh
git branch -m delivered/18-ai-is-now-conscious 18-ai-is-now-conscious  #1
git checkout 18-ai-is-now-conscious  #2
git pull --rebase  #3
git push -d origin ready/18-ai-is-now-conscious  #4
git phlow deliver  #5
```

### build fails due to outside error, and branch will not re-trigger

##### Experience
If the CI service pick up the `ready/` branch and starts integration, but an outside error makes the build fail (CI server crash, network issues, etc.). In some cases the branch will not build again because the CI service already picked up the branch once, even though it is not integrated.

##### Solution
Some CI services has the commit sha stored af the ID for the build, so to make the build re-trigger we need to fake a new ID by changing sha of the commit. Follow the same steps as in [Failed build](#failed-build), but in stead of doing step `#3`, just do a `git commit --amend` and use the same messages. 