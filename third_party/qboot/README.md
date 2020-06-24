# qboot firmware
This is a firmware used for initializing QEMU MicroVM-based virtual machines. It initializes the virtual CPU, and
relocates the Kernel and initramfs to the correct locations and jumps into it. It is the analogue to EDK II on the
normal systems, but orders of magnitude faster and lighter.

This firmware is usually shipped as a precompiled binary by QEMU, but the version they currently ship has a critical
bug (https://github.com/bonzini/qboot/pull/28) preventing our VMs from starting which has been fixed upstream,
but QEMU needs to rebuild their firwmare and Fedora needs to ship an updated QEMU. Since it is not a lot of code, this
just builds qboot in Bazel, getting us that critical fix immediately.
