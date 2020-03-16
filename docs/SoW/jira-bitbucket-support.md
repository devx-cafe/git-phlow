# Git Phlow Jira support - statement of work

## Problem
[git-phlow](https://github.com/code-cafe/git-phlow) only supports GitHub issues and Github repositories while a wast number of companies uses Jira for issue tracking in combination with Bitbucket for repositories. Support for Jira combined with Bitbucket would substantially expand the user base of the git-phlow.

## Solution
The solution will require you to be compliant in Jira with the [Pragmatic workflow](http://www.praqma.com/stories/a-pragmatic-workflow/) which means that an Jira issues should be able to have the [Jira workflow statuses](https://confluence.atlassian.com/adminjiracloud/working-with-workflows-776636540.html) ‘IN PROGRESS’ and ‘ CLOSED’.

The Jira workflow transitions will happen to ‘IN PROGRESS’ as with Github issues using the git-phlow ‘workon’ command. Transition to that state should be allowed from any other state you expect an Jira issue can be in when someone tries to do ‘workon’

The Jira workflow transition to ‘CLOSED’ will be based on [Smart Commit](https://marketplace.atlassian.com/plugins/com.lb.software.stash.smart.commit.lb-software-stash-smart-commit/server/overview) configuration in Bitbucket matching the word ‘close’ like used today on Github.

_Please note the solution requires you to follow at least parts of the Pragmatic workflow._

## Implementation
We will extend the current git-phlow tool to support Jira in combination with Bitbucket, and we expect to have feature parity with the existing GitHub support.

When the git-phlow workon command are not able to transition the Jira issue to ‘IN PROGRESS’ status it will report a warning but complete the git changes like creating the branch.

## Deliveries
A new public release of Git Phlow that supports Jira combined with Bitbucket, and updated documentation covering how to use git-phlow against Jira.

We will supply scripted configuration snippets to configure a Jira workflow compliant with the  Pragmatic workflow, together with documentation. We will also supply documentation on the needed Bitbucket Smart Commit configuration.

## Work load
2 weeks of work.
