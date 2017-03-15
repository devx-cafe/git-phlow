We are happy to announce this new release of git-phlow!
We keep adding features and remove bugs. This time around we have added a bunch of new features, and drastically improved the UI and the overall phlow experience.
Lets face it! The UI was a bit tedious before. This version bumps up the required version of git. Git v1.7.0 or later is required.
Some fancy new commands have been added and previous ones have been improved.

Agent commands have also now been added. So the extension can be used by integration services, who would like to use the pragmatic workflow

#### Features
- agent commands @groenborg
- updated documentation @groenborg
- clean #60 (remove delivered branches from workspace) @groenborg
- colored output #41 (messages are now more informative and pretty) @groenborg
- deliver with tests #54 @groenborg
    - flag (s) for showing test output
- wrapup soft #61 (wrapup does not add all by default anymore) @groenborg
- loading spinners #62 (spinners appear when doing time consuming tasks) @groenborg

#### Agent Features
 - upnext #55 (get next branch for integration) @groenborg
 - deliver #55 (deliver in an agent version) @groenborg