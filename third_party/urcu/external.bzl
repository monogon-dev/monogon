load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

def urcu_external(name, version):
    sums = {
        "0.14.0": "sha256-QvtRKaP//lpLeQ3+HqOnNMae4JX++/ZJMmJpu6lMJi0=",
    }

    http_archive(
        name = name,
        integrity = sums[version],
        strip_prefix = "userspace-rcu-" + version,
        build_file = "@//third_party/urcu:urcu.bzl",
        patch_args = ["-p1"],
        patches = ["//third_party/urcu/patches:generated-files.patch"],
        urls = ["https://github.com/urcu/userspace-rcu/archive/refs/tags/v%s.tar.gz" % version],
    )
