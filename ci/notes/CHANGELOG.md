Last release brought some unexpected problems with our authentication og permissions to github, whoops.
We have made a fix for theses problems, so just go ahead and get the update and these problems should be a part of the past.
To re-authenticate you need to delete the old 'git phlow' token on github.com and delete your phlow settings in your global .gitconfig file.


#### Features
- web command #63 (Browser can be started from phlow) @groenborg
- list issues #69 (Lists opens issues in the terminal) @groenborg
    - List issues only assigned to you
- improved user experience #67 (certain commands can be run outside a repo) @groenborg

#### Bug fixes
- workon should fetch issues #68 (github issues should work) @groenborg