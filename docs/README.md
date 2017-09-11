# Documentation


### phlow auth
```
git phlow auth -t <target>
```
Git phlow auth, authorizes towards a chosen service which can be either `github` or `jira`. Auth is used by git phlow to access and manage issues in your chosen provider.

Auth is dependent on a configuration file unless you are using GitHub.com for your issues and Repositories, in that case you can complete omit using the `--target` flag

#### Self hosted services or Jira
To authenticate with your own hosted service or Jira you _must_ provide a valid configuration (see **Configuration** section). Your configuration must contain a base url to where your service is hosted and the name of the service. When authenticating set your `--target` to the name of the configuration e.g.

```ini
[myjiraconfig]
    integration_branch = master
    remote = origin
    service = jira
    issue_url = https://my.jira.instance.com
    delivery_branch_prefix = ready
```
```bash
git phlow auth --target myjiraconfig
```

### Configuration
git phlow is configured via `.phlow`  files, in your local workspace or in the home folder of your computer. The configuration is standard INI files, just like `.gitconfig`.

#### Config Block
Each configuration is has its own block which can be used by setting the block name to `[default]` or by specifying which configuration you want to use by setting it in the git phlow command.

A block consists of following fields.

- `integration_branch` The name of your default integration branch. normally it's `master`
- `remote` The name of your remote. By default git sets this to `origin`
- `service` The service you use for issues.
- `issue_url` The url to your service where issues are hosted and managed
- `delivery_branch_prefix` The prefix you want your pretested integration system to look for

#### Supported services
-  `github`
-  `jira` - `auth` command only


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
If no  `.phlow`  files are located, git phlow will use an interal default configuration.  The internal configuration is the last resort, and it is strongly recommended to use `.phlow` to keep your projects configuration, even though it mimics the default.

```ini
[default]
    integration_branch = master
    remote = origin
    service = github
    issue_url = https://api.github.com
    delivery_branch_prefix = ready
``` 