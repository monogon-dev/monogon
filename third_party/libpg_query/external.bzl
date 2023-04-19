load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

def libpg_query_external(name, version):
    sums = {
        "13-2.2.0": "07916be1a2b780dee6feed936aaa04ccee2a3afde8570a6920c3a839c87539c6",
    }
    http_archive(
        name = name,
        build_file = "//third_party/libpg_query/external:BUILD.repo",
        sha256 = sums[version],
        strip_prefix = "libpg_query-" + version,
        urls = ["https://github.com/pganalyze/libpg_query/archive/refs/tags/%s.tar.gz" % version],
    )
