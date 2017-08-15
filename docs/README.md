### Configuration
git phlow is configured via `.phlow`, or `.gitignore`  files., in your local workspace or in the home folder of your computer. The configuration is standard INI files, just like gitconfig.

#### Config Block
Each configuration is has it's own block which can be used either by setting the block name to `[default]` or by specifying which configuration you want to use by setting it in the git phlow command.

A block consists of following fields.

- `integration_branch` The name of your default integration branch. normally it's `master`
- `remote` The name of your remote. By default git sets this to `origin`
- `service` The service you use for issues.
- `issue_url` The url to your service where issues are hosted and managed
- `delivery_branch_prefix` The prefix you want your pretested integration system to look for

#### Supported services
-  `github`
-  ~~`jira`~~ -  coming soon


#### Example Configuration
```ini
[name]
    integration_branch = master
    remote = origin
    service = github
    issue_url = https://api.github.com
    pipeline_url = http://concourse.bosh.praqma.cloud/teams/main/pipelines/git-phlow #optional
    delivery_branch_prefix = ready
```

#### Internal Default
If no  `.phlow` or `.gitconfig` files are located, git phlow will use an interal default configuration.  The internal configuration is the last resort, and it is strongly recommended to use `.phlow` to keep your projects configuration, even though it mimiks the default.

```ini
[default]
    integration_branch = master
    remote = origin
    service = github
    issue_url = https://api.github.com
    delivery_branch_prefix = ready
``` 