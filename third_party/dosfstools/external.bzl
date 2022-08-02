load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

def dosfstools_external(name, version):
    sums = {
        "c888797b1d84ffbb949f147e3116e8bfb2e145a7": "4a40b488c0c259c11fb54783fc6f01e5ee912582bb49d33d0d11b11f85a42e8d",
    }

    http_archive(
        name = name,
        sha256 = sums[version],
        strip_prefix = "dosfstools-" + version,
        build_file = "@//third_party/dosfstools:dosfstools.bzl",
        urls = ["https://github.com/dosfstools/dosfstools/archive/%s.zip" % (version)],
    )
