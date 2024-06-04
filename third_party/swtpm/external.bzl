load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

def swtpm_external(name, version):
    sums = {
        # master at 2024/06/04
        "0c9a6c4a12a63b86ab472e69e95bd75853d4fa96": "169ddb139597fa808e112d452457445c79ed521bb34f999066d20de9214056ce",
    }

    http_archive(
        name = name,
        patch_args = ["-p1"],
        patches = [
            "//third_party/swtpm/patches:0001-bazel-compat-glib.h-glib-glib.h.patch",
            "//third_party/swtpm/patches:0002-swtpm_localca-replace-gmp-mpz-dependency-with-boring.patch",
            "//third_party/swtpm/patches:0003-swtpm_setup-replace-dep-on-JSON-GLib-with-sheredom-j.patch",
            "//third_party/swtpm/patches:0004-bazel-support-implement.patch",
        ],
        sha256 = sums[version],
        strip_prefix = "swtpm-" + version,
        urls = ["https://github.com/stefanberger/swtpm/archive/%s.tar.gz" % version],
    )
