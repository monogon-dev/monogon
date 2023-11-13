load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

def libpg_query_external(name, version):
    sums = {
        "15-4.2.3": "8b820d63442b1677ce4f0df2a95b3fafdbc520a82901def81217559ec4df9e6b",
    }
    http_archive(
        name = name,
        build_file = "//third_party/libpg_query/external:BUILD.repo",
        sha256 = sums[version],
        strip_prefix = "libpg_query-" + version,
        urls = ["https://github.com/pganalyze/libpg_query/archive/refs/tags/%s.tar.gz" % version],
    )
