load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

def gnuefi_external(name, version):
    sums = {
        "3.0.14": "5785e77825fec5e666e4c20d7aaa9af4cd952351b2c09593972744fe8436f957",
    }

    http_archive(
        name = name,
        sha256 = sums[version],
        build_file = "@//third_party/gnuefi:gnuefi.bzl",
        strip_prefix = "gnu-efi-%s" % version,
        urls = ["https://github.com/ncroxon/gnu-efi/archive/refs/tags/%s.tar.gz" % version],
    )

