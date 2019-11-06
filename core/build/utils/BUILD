# TODO(leo): I have not been able to figure out a clever way of separating this
# into multiple rules, particularly musl, which hardcodes sandbox paths into its
# toolchain such that a different rule cannot consume it.
#
# For now, using a single massive genrule is the least annoying way to do this.
# As soon as we build more than just mkfs.xfs, we should re-visit this.
#
# Some possibilities:
#
#  - Build the musl toolchain in the build container and use native rules
#    for headers_install and util_linux (they should, in theory, generate
#    well-defined artifacts that we can use in the build).
#
#    This would use Bazel's toolchain definition mechanism to consume the
#    external toolchain, and would be compatible with the native C rules.
#
#    Maybe we can even build the external toolchain inside Bazel somehow?
#
#  - Write a custom rule that handles the toolchain.
#
#  - Converting *everything* to native rules is probably not an option due
#    to how complex the third party build systems we touch are.

genrule(
    name = "utils",
    srcs = [
        "@xfsprogs_dev//:all",
        "@musl//:all",
        "@util_linux//:all",
        "@linux_kernel//:all",
    ],
    outs = [
        "mkfs.xfs",
    ],
    cmd = """
    MUSL=$$PWD/$(RULEDIR)/musl_prefix

    echo "Compiling and installing musl..."
    (
        cd external/musl
        ./configure --prefix=$$MUSL --syslibdir=$$MUSL/lib
        make -j $$(nproc) install
    ) > /dev/null

    echo "Installing Linux kernel headers..."
    (
        cd external/linux_kernel
        make headers_install ARCH=x86_64 INSTALL_HDR_PATH=$$MUSL
    ) > /dev/null

    echo "Compiling util_linux..."
    (
        cd external/util_linux
        ./autogen.sh
        ./configure \
        CC="$$MUSL/bin/musl-gcc" \
        --without-systemd \
        --without-udev \
        --without-btrfs \
        --disable-pylibmount \
        --without-tinfo \
        --prefix=$$MUSL \
        --disable-makeinstall-chown \
        --disable-makeinstall-setuid \
        --with-bashcompletiondir=$$MUSL/usr/share/bash-completion
        make -j $$(nproc) libuuid.la libblkid.la
        echo "Installing util_linux..."
        cp -v .libs/* $$MUSL/lib/
        mkdir -p $$MUSL/include/{uuid,blkid}
        cp libuuid/src/uuid.h $$MUSL/include/uuid/
        cp libblkid/src/blkid.h $$MUSL/include/blkid/
    ) > /dev/null

    echo "Compiling mkfs.xfs..."
    (
        cd external/xfsprogs_dev
        make configure
        ./configure CC="$$MUSL/bin/musl-gcc" CFLAGS="-static" --prefix=$$MUSL
        echo COMPILERING
        make mkfs -j8
    ) > /dev/null

    cp external/xfsprogs_dev/mkfs/mkfs.xfs $(RULEDIR)
    """,
    visibility = ["//visibility:public"],
)
