## Changelog
Windows and Jira support is here! Feast on the godness and the docs

#### Features

- `phlow web` is updated to support target configuration, and Jira #247 @groenborg
typing web will open your browser on your issue page or repository.

- `phlow issues` you can now list jira issues and github issues #246 @groenborg
phlow issues fetches the 30 most recent issues and prints them in your terminal. Nice

- `wrapup` wrapup now support bitbucket smart commits as well #257 @groenborg

- `config` git phlow now builds upon the gitconfig of git #209 @groenborg
The .phlow file have been replaced by .gitconfig and the internal of git.
You can also bootstrap and show your configuration with `git phlow config`

#### Bug Fixes
 - git phlow web no longer opens a malformed url #252 @groenborg

#### improvements
- we did some spring cleaning this fall and removed a lot of bad code #258 @groenborg
such deletion, much improvement!

- we also did some further improvements in the messages that are suppose to help trouble shoot issues!