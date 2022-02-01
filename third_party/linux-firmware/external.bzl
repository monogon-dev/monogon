load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

def linux_firmware_external(name, version):
    sums = {
        "20211216": "c0f735dd232c22d41ce4d23a050a8d6efe3b6b8cbf9d0a636af5f9df66a619a3",
    }
    all_content = """filegroup(name = "all_files", srcs = glob(["**"]), visibility = ["//visibility:public"])"""

    http_archive(
        name = name,
        build_file_content = all_content,
        sha256 = sums[version],
        strip_prefix = "linux-firmware-" + version,
        urls = ["https://git.kernel.org/pub/scm/linux/kernel/git/firmware/linux-firmware.git/snapshot/linux-firmware-%s.tar.gz" % version],
    )
