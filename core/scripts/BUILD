sh_binary(
    name = "launch",
    srcs = ["launch.sh"],
    data = [
        "@//core:image",
        "@//core:swtpm_data",
        "@edk2//:firmware",
    ],
    tags = ["local"],
)

sh_test(
    name = "test_boot",
    size = "small",
    srcs = ["test_boot.sh"],
    tags = ["local"],
    deps = [":launch"],
)
