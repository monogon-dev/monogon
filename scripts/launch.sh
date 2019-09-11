#!/bin/sh

swtpm socket --tpmstate dir=$PWD/vm/tpm --ctrl type=unixio,path=$PWD/vm/tpm/swtpm-sock --tpm2 &

qemu-system-x86_64 -cpu host -smp sockets=1,cpus=1,cores=2,threads=2,maxcpus=4 -m 1024 -machine q35 -enable-kvm -nographic -nodefaults \
    -drive if=pflash,format=raw,readonly,file=$PWD/third_party/edk2/Build/OvmfX64/RELEASE_GCC5/FV/OVMF_CODE.fd \
    -drive if=pflash,format=raw,snapshot=on,file=$PWD/third_party/edk2/Build/OvmfX64/RELEASE_GCC5/FV/OVMF_VARS.fd \
    -drive if=virtio,format=raw,cache=unsafe,file=$PWD/vm/smalltown.img \
    -netdev user,id=net0,hostfwd=tcp::7833-:7833,hostfwd=tcp::7834-:7834 \
    -device virtio-net-pci,netdev=net0 \
    -chardev socket,id=chrtpm,path=$PWD/vm/tpm/swtpm-sock \
    -tpmdev emulator,id=tpm0,chardev=chrtpm \
    -device tpm-tis,tpmdev=tpm0 \
    -debugcon file:vm/debug.log \
    -global isa-debugcon.iobase=0x402 \
    -device ipmi-bmc-sim,id=ipmi0 \
    -device virtio-rng-pci \
    -serial mon:stdio