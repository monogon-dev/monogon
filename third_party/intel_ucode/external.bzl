load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

def intel_ucode_external(name, version):
    sums = {
        "20220207": "532527bd17f3ea6664452b536699818a3bf896e4ace689a43a73624711b7c921",
        "20230214": "3a3cfe2c7642339af9f4c2ad69f5f367dfa4cd1f7f9fd4124dedefb7803591d4",
        "20230808": "fe49bb719441f20335ed6004090ab38cdc374134d36d4f5d30be7ed93b820313",
        "20231114": "cee26f311f7e2c039dd48cd30f995183bde9b98fb4c3039800e2ddaf5c090e55",
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
