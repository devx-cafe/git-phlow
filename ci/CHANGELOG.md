## Changelog


#### Features

- deliver is updated to support target configuration, and Jira #230 @groenborg
you can now deliver work, with settings from the target configuration.. Hurray

- wrapup is now support on windows as well #188 @groenborg
wrapup will now format messages based on GitHub and Jira branches

- workon is now updated to support target configuration, and Jira #224 @groenborg
you can now work on issues located on GitHub and Jira, with settings from the target configuration


#### Bug Fixes
- fix configuration for windows #236 @groenborg
Turns out the config before did not work on windows, however, it does now.

- auth errors are handled when commands using service API are called #240 @groenborg
a successful bug-hunt resulted in one less error, where the tool basically would explode with errors if commands were executed without authorization


#### improvements
- We removed the beautiful spinner when running `workon` and `deliver` #241 @groenborg
it might return in future versions, when it no longer messes with the formatting. Never mess with the formatting!

- Error messages have also been improved, from "something happened" to "your hyper drive core melted a super size hole in you spaceship.. Remember to cool it!"