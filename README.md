# Nexantic monorepo

This is the monorepo storing all of nexantic's internal projects and libraries.

## Environment

All builds should be executed using the shipped `nexantic-dev` container which is automatically built by the create
script.

The container contains all necessary dependencies and env configurations necessary to get started right away.

#### Usage

Spinning up: `scripts/create_container.sh` 

Spinning down: `scripts/destroy_container.sh` 

Running commands: `scripts/run_in_container.sh @`

Using bazel: `scripts/bin/bazel @`
 
