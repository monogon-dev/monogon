#!/bin/sh

swtpm socket --tpmstate dir=tpm --ctrl type=unixio,path=tpm-socket --tpm2 &

qemu-system-x86_64 \
    -cpu host -smp sockets=1,cpus=1,cores=2,threads=2,maxcpus=4 -m 1024 -machine q35 -enable-kvm -nographic -nodefaults \
    -drive if=pflash,format=raw,readonly,file=external/edk2/OVMF_CODE.fd \
    -drive if=pflash,format=raw,snapshot=on,file=external/edk2/OVMF_VARS.fd \
    -drive if=virtio,format=raw,snapshot=on,cache=unsafe,file=smalltown.img \
    -netdev user,id=net0,hostfwd=tcp::7833-:7833,hostfwd=tcp::7834-:7834 \
    -device virtio-net-pci,netdev=net0 \
    -chardev socket,id=chrtpm,path=tpm-socket \
    -tpmdev emulator,id=tpm0,chardev=chrtpm \
    -device tpm-tis,tpmdev=tpm0 \
    -debugcon file:debug.log \
    -global isa-debugcon.iobase=0x402 \
    -device ipmi-bmc-sim,id=ipmi0 \
    -device virtio-rng-pci \
    -serial mon:stdio
