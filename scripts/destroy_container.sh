#!/bin/bash

# TODO(by 2021/02/01): remove this (backward compatibility for dev envs)
! podman pod stop nexantic
! podman pod rm nexantic --force

podman pod stop monogon
podman pod rm monogon --force
