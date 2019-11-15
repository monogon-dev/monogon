sh_binary(
    name = "launch",
    srcs = ["launch.sh"],
    data = [
        "@//core:image",
        "@//core:swtpm_data",
        "@edk2//:firmware",
    ],
)

sh_test(
    name = "test_boot",
    size = "small",
    srcs = ["test_boot.sh"],
    # expects wants a pty, which do not exist in the sandbox
    tags = ["local"],
    deps = [":launch"],
)
