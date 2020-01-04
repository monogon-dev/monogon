#!/bin/bash

create_symlinks() {
  local package=$1

  rm -rf "$package"
  ln -r -s "bazel-bin/gopath/src/git.monogon.dev/source/nexantic.git/$package" "$package"
}

create_symlinks core/generated
