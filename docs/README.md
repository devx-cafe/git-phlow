# Documentation

Table of contents
- [Installation](/docs/installation.md)
- [Workflow](/docs/workflow.md)
- [Documentation](/docs/documentation.md)
- [Configuration](/docs/configuration.md)
- [Contribution guide](/CONTRIBUTING.md)
- [Trouble shooting](/docs/troubleshooting.md)


### Getting started
To get started using this git extension we recommend to read the [blogpost](https://www.praqma.com/stories/git-phlow/) about Git-phlow and the entire automated workflow.

The tool is meant to help developers and non-developers abstract away the complexity of git and all the manual work of integrating changes into other branches. 

1. Get the tool
```sh
#tap our tool repository
brew tap praqma/praqma-tap

# Install git-phlow
brew install git-phlow
```

2. Go to your project and create a configuration. You can use the default generated by the git-phlow

```sh
git phlow config bootstrap # creates .gitconfig
``` 

3. Ensure that git knows you are using a local .gitconfig file
```sh
git config --local include.path ../.gitconfig
```
4. start using the workflow

```sh
git phlow workon <issue>

#Add changes

git phlow wrapup

git phlow deliver

```



### Compatibility Matrix

|                | support | version | 
| -------------- | ------- | ------- |
| GitHub         | &#x2714;| v3
| GitHub Premise | &#x2715;| -
| Jira Cloud     | &#x2715;| -
| Jira Premise   | &#x2714;| 7.3.0
|**macOS**       |
| bash           | &#x2714;| 3.2.57
| zsh            | &#x2714;| 5.4.2
| **windows**    |
| ([git bash](https://git-for-windows.github.io))**| &#x2714;| 4.4.12 |
| powershell     | &#x2714;| 5.1.14393  |
| cmd            | &#x2714;| 10.0.14393 |
| [cmder](http://cmder.net)| &#x2715;  | - |
_** reading from stdin is faulty_