workspace(name = "dev_source_monogon")

# bazeldnf is used to generate our sandbox root.
load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "bazeldnf",
    sha256 = "cd75fbbad6f191c26b036132d57ca731cce067e9330306a8a2beb3e51af991a8",
    urls = [
        "https://github.com/rmohr/bazeldnf/releases/download/v0.5.8/bazeldnf-v0.5.8.tar.gz",
    ],
)

load("@bazeldnf//:deps.bzl", "bazeldnf_dependencies")

bazeldnf_dependencies()

load("//third_party/sandboxroot:repositories.bzl", "sandbox_dependencies")

sandbox_dependencies()
