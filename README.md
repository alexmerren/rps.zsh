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

The only option to install this project is building from source.
To run all tests, vendor dependencies, and build the executable:
run `make all`. 

`rps` requires a GitHub token with the appropriate permissions.
Specifically, the ability to read the user's starred
repositories, and all repositories owned by the user. This token
has to be entered into the `GITHUB_TOKEN` section of the
`config.yaml` file.

The binary can then be installed gloablly via `make install`.  This will
install the executable in `usr/local/bin`, and the config file in
`/usr/local/share/rps`. If you have re-built the binary, `make
install_no_config` only moves the binary and does not overwrite the
configuration file.

## Generating a GitHub Token

[Here](https://docs.github.com/en/enterprise-server@3.6/authentication/keeping-your-account-and-data-secure/managing-your-personal-access-tokens) you can find a guide to generating a GitHub token.

## Upcoming Features

- [ ] Create script to call rps
- [ ] Create rps to list repositories and print them to stdout
- [ ] De-duplicate the incoming repositories
- [ ] 
