# Installation

git-phlow is available on:
- Linux
- Windows
- macOS

### Mac
If you have homebrew installed, you can tap our official repository. Otherwise you can download the official release on the releases page on GitHub.

```shell
# Tap our repository
brew tap praqma/praqma-tap

# Install git-phlow
brew install git-phlow

# Upgrade to a new version
brew upgrade git-phlow
```


### Linux
Install git-phlow by running the following commands. This will download git-phlow from the GitHub releases, extract it and move in to the /usr/local/bin directory, and enable you to run `git phlow` in the terminal. Remember to use the latest release.

```shell
wget https://github.com/code-cafe/git-phlow/releases/download/v3.8.2/git-phlow-3.8.2-linux-amd64.tar.gz && tar -xzf git-phlow-3.8.2-linux-amd64.tar.gz

mv git-phlow /usr/local/bin
```

You may need to run the last command with sudo, depending on your setup:

```
sudo mv git-phlow /usr/local/bin
```

### Windows

On Windows you have the option of using either using [scoop](http://scoop.sh/) or [Chocolatey](https://chocolatey.org/)

**Install using Scoop**

_First install Scoop if you don't have it and add our scoop-bucket_:

Open PowerShell and get the latest release of scoop
```
iex (new-object net.webclient).downloadstring('https://get.scoop.sh')
```
Add Praqmas scoop-bucket
```
scoop bucket add praqma-tools https://github.com/code-cafe/scoop-praqma-tools.git
```

_Then install using scoop_:
```
scoop install git-phlow
```

**Install using Chocolatey**

Package is available here, also including installation command: https://chocolatey.org/packages/git-phlow/

You need to have Chocolatey installed, and then just run the following command in PowerShell:

    choco install git-phlow


##### Compatibility
- on windows there can be cases where the colors will not show. Therefore we recommend to run every command with the `--no-color` option. 

## Updating git-phlow
### Mac
On Mac using homebrew, run the upgrade command.

```shell
brew upgrade git-phlow
```
### Linux
In Linux you run the installation again, but this time override the old `git-phlow` binary.

### Windows
Using either Scoop or Chocolatey you installed from originally run:
```
scoop update git-phlow
```
or
```
choco upgrade git-phlow
```
_Please notice Chocolatey packages has an approval time on new releases, so it might not be available until a few days after the other releases_
