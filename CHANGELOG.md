# Changelog

It took a long time, but now it's here; Documentation! This update includes a lot of documentation in the repository, as well as in git phlow. All the commands have been updated to show the same documentation as the one found on our github repository.
We also updated our pipeline setup and tasks, but that stuff is mostly for our own sake.

We always welcome updates to our documentation, so if you feel like contributing, you should.

### Added

- `phlow --init` is not used instead of `phlow config bootstrap` #328 @groenborg

### Changed

- `git phlow --version` now prints versioning details in json format #328 @groenborg
- `phlow cleanup` now starts a guided flow for deleting branches #329 @groenborg
- `phlow cleanup --delivered` works as the base command in previous versions #329 @groenborg
- `phlow cleanup --tidy` prunes remote branches #329 @groenborg
- `phlow cleanup --dry-run` will let you run the cleanup commands safely without removing any branches #329 @groenborg

### Removed

- Jira support have been removed completely
- `phlow config` command has been removed
