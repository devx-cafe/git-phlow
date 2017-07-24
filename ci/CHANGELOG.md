This new version contains a **Major** semantic change in the **wrapup** command. It has been returned to the default behaviour, so wrapup no longer requires `--hard` to add you changes, it does that automatically. I apologize for the inconvenience and confusion these changes to `wrapup` have introduced.
The output and error handling is under improvement to make the tool more userfriendly to use. 



#### Improvements
- upnext now has a prefix option #79 @groenborg
- deliver now always change back to your default branch #75 @groenborg 
- made workon less verbose #115 @groenborg
- return wrapup to old behaviour #89 @groenborg
- wrapup now has a `--force` flag for a custom commit message #90 @groenborg

#### Bug fixes
- private repo access for web api should be fixed #81 #83 #84@groenborg 
- fixed typo in workon #95 @randomsort
- removed squash message #85 @groenborg


