# Smalltown Operating System

## Run build

The build uses a Fedora 30 base image with a set of dependencies.
Guide has been tested on a Fedora 30 host, with latest rW deployed.

Launch the VM:

```
scripts/bin/bazel run //core/scripts:launch
```

Exit qemu using the monitor console: `Ctrl-A c quit`.
