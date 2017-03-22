Last release brought some unexpected problems with our authentication og permissions to github, whoops.
We have made a fix for theses problems, so just go ahead and get the update and these problems should be a part of the past.
To re-authenticate you need to delete the old 'git phlow' token on github.com and delete your phlow settings in your global .gitconfig file.

#### Bug fixes
- phlow can be authenticated on multiple machines #72 @groenborg
- list issues should now work on private repositories #71 @groenborg