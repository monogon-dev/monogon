filegroup(name = "all_files", srcs = glob(["**"]), visibility = ["//visibility:public"])
filegroup(name = "metadata", srcs = ["WHENCE"], visibility = ["//visibility:public"])
filegroup(name = "amd_ucode", srcs = glob(["amd-ucode/*.bin"]), visibility = ["//visibility:public"])
