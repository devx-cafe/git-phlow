This version bumps up the required version of git. Git v1.7.0 or later is required.
clean command have been added, so now you can cleanup the workspace by deleting all local and remote branches
prefixed with 'delivered'.

Agent commands have now been added. So the extension can be used by integration services.


Features
- agent commands
- updated documentation
- clean #60 (remove delivered branches from workspace) @groenborg

Agent Features
 - upnext #55 (get next branch for integration) @groenborg
 - integrate #55 (deliver in an agent version) @groenborg