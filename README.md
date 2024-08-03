# goball of bash scripts for deploying gradle projects

## Installation

Install [go](https://golang.org/) then:

```bash
git clone https://github.com/tom-power/depl.git &&
cd ./depl &&
cp .config/example.install.env sh/.env &&
nano sh/.env &&
sh/install.sh [--remote|--completions]
```

## Setup

Given a [project](https://github.com/tom-power/depl-example-project) you'd like to deploy, and a remote host you have [ssh](https://www.openssh.com/) access to:

```bash
cp ./config/example.env /path/to/project/sh/.env &&
cd /path/to/project &&
nano .env
# setup git ..
# setup git on remote
# push to remote ..
# clone remote to $deployDir ..
```

## Usage

`depl.sh list` to see commands and `depl.sh <command> --explain` to see what a command does.

an example to deploy a project:

```bash
depl.sh gitPush &&
depl.sh remote gitPull &&
depl.sh remote gradleBuild &&
depl.sh remote dockerUp
```

### Shell integration

- [omz completion](https://github.com/tom-power/depl/blob/master/config/.oh-my-zsh/custom/completions/_depl)
