load(":musl.bzl", "musl_headers")
load(":linux.bzl", "linux_headers")
load(":tarball.bzl", "musl_gcc_tarball")

linux_headers(
    name = "linux_headers",
    src = "@linux//:all",
    arch = "x86_64",
    visibility = ["//visibility:public"],
)

musl_headers(
    name = "musl_headers",
    src = "@musl//:all",
    arch = "x86_64",
    visibility = ["//visibility:public"],
)

musl_gcc_tarball(
    name = "sysroot",
    linux_headers = ":linux_headers",
    musl = "//third_party/musl",
    musl_headers = ":musl_headers",
)
