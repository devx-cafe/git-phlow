# Configuration

git-phlow is configured in `.gitconfig`, in your local workspace, this means that git-phlow uses git's own internal config and can therefore be placed anywhere git can read it e.g. global .gitconfig in you home folder.

To get started with a `.gitconfig` in your local repository you must first let git know that the filepath must be included, run this command to do so:
```sh
git config --local include.path ../.gitconfig
```

##### Config Block
Each configuration is its own block. You can have as many configuration blocks as you want, just as long as they all have the necessary fields. To use a specific block you run a command with the `-t or --target` option with the name of the block. git-phlow by default looks for a block named `[phlow]`. So if you name your preferred block `phlow` you do not need to use the `-t` flag.

A block consists of following fields.

- `integration-branch` The name of your default integration branch. `master` is a common to use as integration branch
- `remote` The name of your remote. By default git sets this to `origin`
- `service` The service you use for issues
- `issue-api` The url to your api where issues are hosted and managed
- `issue-web` The url to your webpage where issues are hosted and managed
- `delivery-branch-prefix` The prefix for the branches pushed with `git phlow deliver`

##### Supported services
git-phlow can be used with **jira on-premise** and **github cloud**. Only one service can be used per configuration block.
To configure git-phlow, set the `service` property to either **jira** or **github**
- `service = jira`
- `service = github`

##### Example Configuration with jira
```ini
[phlow]
    remote                 = myfork
    service                = jira
    integration-branch     = master
    issue-api              = https://jira.com
    issue-web              = https://jira.com
    delivery-branch-prefix = ready
```

**important**: When entering the url of your issue service providers, it is important that only the base url is provided, which means no `www.mygithubinstance.com:8080/` but `www.mygithubinstance.com:8080`

##### Internal Default
If no  `[phlow]`  blocks are located within the git config, git-phlow will use an internal default configuration. The internal configuration is the last resort, and it is strongly recommended to use `.gitconfig` to keep your projects configuration, even though it mimics the default.

```ini
[phlow]
    remote                 = origin
    service                = github
    integration-branch     = master
    issue-api              = https://api.github.com
    issue-web              = https://github.com
    delivery-branch-prefix = ready
```