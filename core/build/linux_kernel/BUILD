genrule(
    name = "kernel",
    srcs = [
        "@linux_kernel//:all",
        "linux-smalltown.config",
    ],
    outs = [
        "bzImage",
    ],
    cmd = """
    DIR=external/linux_kernel

    mkdir $$DIR/.bin

    cp $(location linux-smalltown.config) $$DIR/.config

    (cd $$DIR && make -j $$(nproc)) >/dev/null

    cp $$DIR/arch/x86/boot/bzImage $(RULEDIR)
    """,
    visibility = ["//visibility:public"],
)

genrule(
    name = "initramfs",
    srcs = [
        "@//core/cmd/init",
        "@//core/build/utils:mkfs.xfs",
        "@kubernetes//cmd/kube-apiserver",
    ],
    outs = [
        "initramfs.cpio.lz4",
    ],
    cmd = """
    $(location @linux_kernel//:gen_init_cpio) - <<- 'EOF' | lz4 -l > \"$@\" 
dir /dev 0755 0 0
nod /dev/console 0600 0 0 c 5 1
nod /dev/null 0644 0 0 c 1 3
file /init $(location @//core/cmd/init) 0755 0 0
dir /bin 0755 0 0
file /bin/mkfs.xfs $(location @//core/build/utils:mkfs.xfs) 0755 0 0
file /bin/kube-apiserver $(location @kubernetes//cmd/kube-apiserver) 0755 0 0
EOF
    """,
    tools = [
        "@linux_kernel//:gen_init_cpio",
    ],
    visibility = ["//visibility:public"],
)
