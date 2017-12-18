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