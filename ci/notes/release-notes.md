This version bumps up the required version of git. Git v1.7.0 or later is required.
clean command have been added, so now you can cleanup the workspace by deleting all local and remote branches
prefixed with 'delivered'.

Agent commands have now been added. So the extension can be used by integration services.

#### Features
- agent commands @groenborg
- updated documentation @groenborg
- clean #60 (remove delivered branches from workspace) @groenborg
- colored output #41 (messages are now more informative and pretty) @groenborg
- deliver with tests #54 @groenborg
    - flag (s) for showing test output
- wrapup soft #61 (wrapup does not add all by default anymore) @groenborg

#### Agent Features
 - upnext #55 (get next branch for integration) @groenborg
 - deliver #55 (deliver in an agent version) @groenborg