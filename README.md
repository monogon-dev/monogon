# Smalltown Operating System

## Run build

The build uses a Fedora 30 base image with a set of dependencies.

Start container shell:

```
modprobe kvm

podman build -t smalltown-builder .

podman run -it --rm \
    -v $(pwd):/work \
    -v /dev/null:/work/.git \
    -v /dev/null:/work/.idea \
    -v /dev/null:/work/.arcconfig \
    --device /dev/kvm \
    --net=host \
    smalltown-builder bash
```

Launch the VM:

    bazelisk run scripts:launch

Exit qemu using the monitor console: `Ctrl-A c quit`.

If your host is low on entropy, consider running rngd from rng-tools for development.
