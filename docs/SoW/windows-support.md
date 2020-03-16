# Git Phlow Windows support - statement of work

## Problem
[git-phlow](https://github.com/code-cafe/git-phlow) currently doesn’t work on Windows due to the way it executes commands, thus a very large user base on Windows can not benefit from the tool, as well as the tool can not benefit from having many users adding improvements..

When adding Windows support, we also have to address the issue of installation and updates.

## Solution
Add support for Windows by implementing command execution in a Windows compliant manner.

We will provide installation and update support by providing either [MSI](https://technet.microsoft.com/en-us/library/cc978328.aspx) or [WiX](http://wixtoolset.org/) installer package, and optionally [Chocolatey](https://chocolatey.org/) package manager installer. Chocolatey would be aligned with the way we install on other systems (apt-get, brew) but the user base of Chocolatey is yet small so this is only an option if estimate allows it.

## Implementation
Git Phlow does git operations by calling the locally installed git on command line instead of git libraries internally. In Go language we found no viable git libraries, and the Git extension with the largest user base called [Git LFS](https://git-lfs.github.com/) is working the same way. We will make our Git extension Windows compliant the same way Git LFS works as this is a proven concept.

## Deliveries
A new public release of Git Phlow that supports Windows 7 and 10 versions and distributed like our existing git-phlow package today [through file published on Github](https://github.com/code-cafe/git-phlow#alternative). Installers will be unsigned for now.

Further updated documentation covering installation on Windows and any needed notes on usage.

## Work load
Estimate is about 50-60 hours. Chocolatey installer package support, will be included in the 60 hours if possible, else it will be considered extra 8-16 hours that are optionally agreed upon.

There are some uncertainty related to Windows environment for testing, so after 24 hours we will evaluate the estimate and agree on continued effort with stakeholders if there are changes in the overall estimate.
