# RPS

`rps` (or Repository Selector) is a command-line tool that enables quick
installation of GitHub repositories that are owned or starred by a user.

![](./docs/splash.gif)

## Usage

`rps` is a simple command-line tool that allows you to simply download
repositories that you either own or have starred. This tool uses `fzf` as the
prompt to select the repository. You can select by using the arrow keys or
typing in the name of the specific repository. Repository Selector then invokes
`git clone <repository>` in the current directory using SSH. 

## Installation

This project assumes that you have a few things installed:

 - [Git](https://git-scm.com)
 - [Golang 1.16+](https://go.dev/doc/install)
 - [FZF](https://github.com/junegunn/fzf)

After the proper configuration, it can be built and installed using the
following:

```bash
make build && make install
```

The zsh function inside `rps.zsh` can then be copied and sourced by zsh to be
used as a standalone function. 

`rps` requires a GitHub token with the appropriate permissions.
Specifically, the ability to read the user's starred
repositories, and all repositories owned by the user. This token
has to be entered into the `GITHUB_TOKEN` variable in [`rps.zsh`](./rps.zsh).

## Generating a GitHub Token

A guide to generate an appropriate GitHub token can be found [here](https://docs.github.com/en/enterprise-server@3.6/authentication/keeping-your-account-and-data-secure/managing-your-personal-access-tokens).

## Inspiration

This is essentially copied this tutorial [here](https://mattorb.com/fuzzy-find-github-repository/).
