# Smalltown Operating System

## Run build

The build uses a Fedora 30 base image with a set of dependencies.
Guide has been tested on a Fedora 30 host, with latest rW deployed.

Build the base image:

```
podman build -t smalltown-builder .
```

Launch the VM:

```
scripts/bin/bazel run scripts:launch
```

Exit qemu using the monitor console: `Ctrl-A c quit`.
