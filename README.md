# acli
A command line utility to scaffold projects quickly.

## Requirements

Depending on the template you will need:

- Go, Rust, C#, Node
- make

## Usage

- Scaffold a new project
  - `aclid init project-name --template vite-go --git-init`
- List all templates
  - `aclid ls`

## Installing, running, and building

### Install a new project

> Requirements: node

- `cd project-name`
- `make install` 

### Run in frontend

- `make runui` 

### Run from server

- `make run`

### Make a docker container

> Required: Docker

- `make docker`

## Templates

### Vite-Golang
