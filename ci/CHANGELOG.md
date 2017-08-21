## Changelog
An apology to the users who was affected by the recent major release, that was not cool. We have fixed a few things that should make life easier.
We also cleaned up the code a bit and made it look nicer.

#### Improvements
- bootstrap .phlow if it does not exist #203 @groenborg
git phlow will now create a .phlow file if none exists. It will also see if it can find the default branch like it did before.

- workon should not do a pull rebase #200 @groenborg
workon will not do a pull rebase, when setting up a workspace.
