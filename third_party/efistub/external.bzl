load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

def efistub_external(name, version):
    sums = {
        "3542da2442d8b29661b47c42ad7e5fa9bc8562ec": "021c135bee39ca7346d1f09923be7c044a3d35866ff411a7c9626702ff4c9523",
    }

    http_archive(
        name = name,
        build_file = "@//third_party/efistub:efistub.bzl",
        sha256 = sums[version],
        strip_prefix = "systemd-%s" % version,
        patch_args = ["-p1"],
        patches = [
            "//third_party/efistub/patches:use-sysv-for-kernel.patch",
            "//third_party/efistub/patches:remove-wrong-cmdline-assertion.patch",
            "//third_party/efistub/patches:ab-slot-handling.patch",
        ],
        urls = ["https://github.com/systemd/systemd/archive/%s.zip" % version],
    )
