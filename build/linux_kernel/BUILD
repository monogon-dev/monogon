genrule(
    name = "image",
    srcs = [
        "@linux_kernel//:all",
        "@//cmd/init",
        "@//build/utils",
        "initramfs.list",
        "linux-smalltown.config",
    ],
    outs = [
        "bzImage",
    ],
    cmd = """
    DIR=external/linux_kernel

    mkdir $$DIR/.bin

    cp $(location linux-smalltown.config) $$DIR/.config
    cp $(location @//cmd/init) $$DIR/.bin/init
    cp $(locations @//build/utils) $$DIR/.bin/
    cp $(location initramfs.list) $$DIR/initramfs.list

    (cd $$DIR && make -j 16) >/dev/null

    cp $$DIR/arch/x86/boot/bzImage $(RULEDIR)
    """,
    visibility = ["//visibility:public"],
)
