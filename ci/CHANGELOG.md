## Changelog
git phlow have had a _huge_ makeover in this release. Configuration have been introduced, and git phlow now works by reading
configuration from a .phlow file, instead of guessing your remote and integration branch. This also brings support for new services like Jira and bitbucket.
The underlying git implementation have been stabilized and is now up and running on windows, so do not feel cheated anymore.

#### Features
- Auth for jira #185 @groenborg
git phlow can now authorize against Jira using the right configuration. See [docs](docs/README.md)

- phlow as configuration #192 @groenborg
git phlow now works based on a .phlow config file

#### Improvements
- git for windows is stabilized

#### Bug fixes
- pull when setting up workspace #193 @lakruzz

