# Workflow 

The Praqma workflow - called git Phlow - is a combination of the Git extension and CI service for automation. In this the CI server is concourse with the git-phlow resource installed.

![image](/docs/images/phlowmodel.svg)

The top diagram shows the steps the entire workflow entails.
1. You do `git phlow workon` on an issue which has been groomed and is ready to become a feature in your project.
2. You do `git phlow wrapup` and `git phlow deliver`, to deliver your changes to the remote repository. The branches will be prefixed with `delivery-branch-prefix` from the .gitconfig file.
3. In the Concourse pipeline we are using the git-phlow tollgate resource. It will integrate branches prefixed with `ready/`. This means the pipeline triggers on such branches and if success merges the changes with the integration branch.
4. Now the tollgate tests will run in the job and if they pass, the already merged changes will be pushed to the integration branch. The ready branch will be deleted if everything goes well so it leaves things nice and clean. 

## git-phlow usage
- First time, authenticate to GitHub:

```git
git phlow auth
```

- Look at the issues by using `ghi` or going to GitHub or Waffle
- Find an issue to work on, and note its id e.g. `42`
- Go to the local git repository and type:

```git
git phlow workon 42
```

- This will create a branch: `42-some-issue-title`
- Make the changes and complete the task
- When done, add everything and commit with a single command:

```git
git phlow wrapup
```

- Now that the changes are ready, type:

```git
git phlow deliver
```

- It renames the local branch to: `delivered/42-some-issue-title`
- It pushes remote branch named: `ready/42-some-issue-title`

If the automation tools have been configured to listen for __ready/__ branches, they will pick up the branch and integrate them with the stable branch.
If you do not have an automation system, you can deliver the work to your default branch by doing a local deliver instead.

```git
git phlow deliver --local
```

- Changes branch to `master`
- Merges changes from `42-some-issue-title`
- Pushes changes to GitHub
