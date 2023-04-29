# RPS

`rps` (or Repository Selector) is a command-line tool that enables
quick installation of GitHub repositories that are owned or starred by a user.
If you prefer your directories to be clean, like me, then this tool removes the
irritation of re-downloading your projects!

![](./docs/splash.png)

## Table of Contents

 * [Usage](#usage)
 * [Documentation](#documentation)
 * [Installation](#installation)
 * [Upcoming Features](#upcoming-features)

## Usage

It is a simple command-line tool that allows you to pick which repository you
want to clone. You can either use Vim movement or arrow keys to move. `/` can
be used to search through either the name or owner of the repository. Finally,
`Enter` selects and clones the repository in your current directory.

## Documentation

TBD

## Installation

The installation steps assume that you have filled out the config.yaml with
your specific information. Please note that `rps` will **NOT** work without the
config placed in this directory and filled out properly.

```bash
mkdir -p ~/.config/rps
make all
cp ./dist/rps ./config.yaml ~/.config/rps/
```

## Upcoming Features

I will probbaly re-factor this in the near future so that it is somewhat
maintainable. Documentation will be necessary to maintain this project
longterm, and thorough testing still needs to be implemented. Don't sue me, I
did this in one (1) day.
