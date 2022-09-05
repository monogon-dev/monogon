load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

def libpg_query_external(name, version):
    sums = {
        "13-2.1.2": "101a7851ee065d824fe06e300b78355a79bd5411864de707761282a0c57a0a97",
    }
    http_archive(
        name = name,
        build_file = "//third_party/libpg_query/external:BUILD.repo",
        sha256 = sums[version],
        strip_prefix = "libpg_query-" + version,
        urls = ["https://github.com/pganalyze/libpg_query/archive/refs/tags/%s.tar.gz" % version],
    )
