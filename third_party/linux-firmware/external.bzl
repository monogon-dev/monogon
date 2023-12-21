load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

def linux_firmware_external(name, version):
    sums = {
        "20211216": "c0f735dd232c22d41ce4d23a050a8d6efe3b6b8cbf9d0a636af5f9df66a619a3",
        "20230310": "14c472af10f9b566c4f575aeb30d8a274d54b1660007e7426b7e4ea21dff81aa",
        # 2023-07-25 master for Zenbleed (CVE-2023-20593)
        "b6ea35ff6b9869470a0c68813f1668acb3d356a8": "67e58b74fb0eebb17fdf95c58a24c6244f93bb0ae8e880f1814ad80463f3a935",
        # 2023-08-09 master for Inception (CVE-2023-20569) and
        # Phantom (CVE-2022-23825)
        "f2eb058afc57348cde66852272d6bf11da1eef8f": "fcd570b8b259049dd84a0326f17a313271962f806ca32dbd9e40cdd9079857d0",
        "20230919": "1dac602218f83f2c81dd72e599ae6c926901b3d36babccce46cd84293a37e473",
        "20231211": "d0ba54f05f5dd34b0fc5a1e1970cd9cbc48491d2da97f3798a9e13530dc18298",
    }
    all_content = """
filegroup(name = "all_files", srcs = glob(["**"]), visibility = ["//visibility:public"])
filegroup(name = "metadata", srcs = ["WHENCE"], visibility = ["//visibility:public"])
filegroup(name = "amd_ucode", srcs = glob(["amd-ucode/*.bin"]), visibility = ["//visibility:public"])
    """

    http_archive(
        name = name,
        build_file_content = all_content,
        sha256 = sums[version],
        strip_prefix = "linux-firmware-" + version,
        urls = ["https://git.kernel.org/pub/scm/linux/kernel/git/firmware/linux-firmware.git/snapshot/linux-firmware-%s.tar.gz" % version],
    )
