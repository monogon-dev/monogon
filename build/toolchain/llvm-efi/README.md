llvm-efi
========

llvm-efi is a Bazel cc toolchain that uses the sandbox sysroot LLVM/clang with flags targeting freestanding EFI.
EFI headers are not shipped as part of the toolchain, but are available as a cc_library from `@gnuefi//:gnuefi`.

Usage
-----

To use this toolchain explicitly while building a `cc_binary`, do:

    bazel build --platforms=//build/platforms:efi_amd64 //foo/bar

During an actual build however, the right toolchain should be selected using transitions
or other configuration mechanisms.

fltused
-------

This is a special symbol emitted by MSVC-compatible compilers. In an EFI environment it can be ignored, but it needs to
be defined. See fltused.c for more information on the symbol. Since we cannot build an object file with Bazel and
building things for toolchains isn't a thing anyways, this file is prebuilt. If this ever needs to be rebuilt (which
will probably never happen since there is only one static symbol in there) this can be done with the following clang
invocation:

    clang -target x86_64-unknown-windows -fno-ms-compatibility -fno-ms-extensions -ffreestanding -o fltused.o .o -c fltused.c
   
