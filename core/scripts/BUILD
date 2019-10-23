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
