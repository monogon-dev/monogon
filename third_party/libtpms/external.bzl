load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

def libtpms_external(name, version):
    sums = {
        # master at 2024/01/09 (0.10.0 prerelease).
        "93a827aeccd3ab2178281571b1545dcfffa2991b": "e509e0ba109f77da517b5e58a9f093beb040525e6be51de06d1153c8278c70d1",
    }

    http_archive(
        name = name,
        patch_args = ["-p1"],
        patches = [
            "//third_party/libtpms/patches:0001-boringssl-compat-new-SHA-types.patch",
            "//third_party/libtpms/patches:0002-boringssl-compat-removed-const_DES_cblock.patch",
            "//third_party/libtpms/patches:0003-boringssl-compat-removed-EC_POINTs_mul.patch",
            "//third_party/libtpms/patches:0004-boringssl-compat-removed-camellia-support.patch",
            "//third_party/libtpms/patches:0005-boringssl-compat-remove-constant-time-flags-UNSAFE.patch",
            "//third_party/libtpms/patches:0006-bazel-support-implement.patch",
        ],
        sha256 = sums[version],
        strip_prefix = "libtpms-" + version,
        urls = ["https://github.com/stefanberger/libtpms/archive/%s.tar.gz" % version],
    )
