## Install pre-release binary from AWS bucket

Start by installing git-phlow via scoop, so the scoop folder structure has been created and the PATHS are set correctly. This can be skipped if you already have git-phlow installed via scoop or if you know how to access `exe` files within the windows path from powershell.

1. Download the binary from AWS and unzip it. It should contain a `git-phlow.exe` file.

2. Navigate to your scoop folder which is commonly located at `C:\Users\$USER\scoop\apps\git-phlow`.

3. You will have folder called current, which contains a `git-phlow.exe` and two JSON files.  Replace the `git-phlow.exe` with the `git-phlow.exe` you dowloaded from AWS. This only works if you used the scoop installation.


This is a workaround until we have beta-releases on scoop available. If you wan't to see that scoop have registered the changes you can run `scoop status` which will tell you that the version is outdated and that the manifest is wrong because of the new `git-phlow.exe` file. When there is a new actual release, scoop should just override the beta and install the new version. 


If the steps have been followed, you should have a successfully installed pre-release of git-phlow.




