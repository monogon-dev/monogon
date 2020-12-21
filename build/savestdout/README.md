savestdout
==========

`savestdout` is a small tool to save the stdout of a command to a file, without using
a shell.

It was made to be used in Bazel rule definitions that want to run a command and save
its output to stdout without going through `ctx.actions.run\_shell`.

Once [bazelbuild/bazel/issues/5511](https://github.com/bazelbuild/bazel/issues/5511)
gets fixed, rules that need this behaviour can start using native Bazel functionality
instead, and this tool should be deleted.

Usage
-----

Command line usage:

    bazel build //build/savestdout
    bazel run bazel-bin/build/savestdout/*/savestdout /tmp/foo ps aux

For an example of use in rules, see `node_initramfs` in `//code/def.bzl`.
