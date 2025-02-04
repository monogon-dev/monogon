# Anything other than family 6 is not interesting to us
filegroup(name = "fam6h", srcs = glob(["intel-ucode/06-*"]), visibility = ["//visibility:public"])
