load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
load("//core/tools/kconfig-patcher:kconfig-patcher.bzl", "kconfig_patch")

go_library(
    name = "go_default_library",
    srcs = ["main.go"],
    importpath = "git.monogon.dev/source/nexantic.git/core/tools/ktest",
    visibility = ["//visibility:private"],
)

go_binary(
    name = "ktest",
    embed = [":go_default_library"],
    pure = "on",
    visibility = ["//visibility:public"],
)

kconfig_patch(
    name = "testing-config",
    src = "//third_party/linux:kernel-config",
    out = "testing.config",
    override_configs = {
        # Unlock command line
        "CONFIG_CMDLINE_OVERRIDE": "n",
        "CONFIG_CMDLINE_BOOL": "n",
        # Shave off 1 second from boot time
        "CONFIG_SERIO_I8042": "",
        "CONFIG_KEYBOARD_ATKBD": "",
        "CONFIG_RTC_DRV_CMOS": "",
        # Shave off an additional 18ms (half of the boot time)
        "CONFIG_DEBUG_WX": "",
    },
)

genrule(
    name = "linux-testing",
    srcs = [
        "@linux//:all",
        ":testing-config",
    ],
    outs = [
        "linux-testing.elf",
    ],
    cmd = """
    DIR=external/linux

    mkdir $$DIR/.bin

    cp $(location :testing-config) $$DIR/.config

    (cd $$DIR && make -j $$(nproc) vmlinux >/dev/null)

    cp $$DIR/vmlinux $@
    """,
    visibility = ["//visibility:public"],
)

filegroup(
    name = "test-script",
    srcs = ["run_ktest.sh"],
    visibility = ["//visibility:public"],
)
