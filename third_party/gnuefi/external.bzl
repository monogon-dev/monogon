load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

def gnuefi_external(name, version):
    sums = {
        "3.0.14": "b73b643a0d5697d1f396d7431448e886dd805668789578e3e1a28277c9528435",
    }

    http_archive(
        name = name,
        sha256 = sums[version],
        build_file = "@//third_party/gnuefi:gnuefi.bzl",
        strip_prefix = "gnu-efi-%s" % version,
        urls = ["https://netix.dl.sourceforge.net/project/gnu-efi/gnu-efi-%s.tar.bz2" % version],
    )
