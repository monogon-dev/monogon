sh_library(
    name = "vm_deps",
    data = [
        "//core:image",
        "//core:swtpm_data",
        "//third_party/edk2:firmware",
    ],
)

sh_binary(
    name = "launch",
    srcs = ["launch.sh"],
    deps = [":vm_deps"],
)

sh_library(
    name = "test_deps",
    data = [
        ":launch",
        "//third_party/kubernetes:kubectl",
    ],
)

sh_test(
    name = "test_boot",
    size = "medium",
    srcs = ["test_boot.sh"],
    # expects wants a pty, which do not exist in the sandbox
    tags = ["local"],
    deps = [
        ":test_deps",
        ":vm_deps",
    ],
)
