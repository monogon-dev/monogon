load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

def intel_ucode_external(name, version):
    sums = {
        "20220207": "532527bd17f3ea6664452b536699818a3bf896e4ace689a43a73624711b7c921",
    }
    all_content = """
# Anything other than family 6 is not interesting to us
filegroup(name = "fam6h", srcs = glob(["intel-ucode/06-*"]), visibility = ["//visibility:public"])
    """

    http_archive(
        name = name,
        build_file_content = all_content,
        sha256 = sums[version],
        strip_prefix = "Intel-Linux-Processor-Microcode-Data-Files-microcode-" + version,
        urls = ["https://github.com/intel/Intel-Linux-Processor-Microcode-Data-Files/archive/refs/tags/microcode-%s.tar.gz" % version],
    )
