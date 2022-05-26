# goball of bash scripts for deploying gradle projects

## Installation

Clone this project:

```bash
git clone https://github.com/tom-power/depl.git
cd ./depl
```

add .env file

```bash
cp ./example.evn ./.env
```
install [go](https://golang.org/) and run `./build.sh`.

## Setup

Given a [project](https://github.com/tom-power/depl-example-project) you'd like to deploy, and a remote host you have ssh access to:

```bash
cp ./example.project.env /path/to/project/sh/.env
cd /path/to/project
depl.sh init // sets up a remote git repository and a directory to deploy the project from
```

## Usage

`depl.sh list` to see commands and `depl.sh explain <command>` to see what a command does.

an example to deploy a project:

```bash
depl.sh copyLibs &&
depl.sh test &&
depl.sh push &&
depl.sh remote deploy
```

### Shell integration

- [omz completion](https://github.com/tom-power/depl/blob/master/config/.oh-my-zsh/custom/completions/_depl)