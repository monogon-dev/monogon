load("@bazel_gazelle//:def.bzl", "gazelle")

# gazelle:prefix git.monogon.dev/source/smalltown.git
gazelle(name = "gazelle")

genrule(
    name = "image",
    srcs = [
        "@//cmd/mkimage",
        "@//build/linux_kernel:image",
    ],
    outs = [
        "smalltown.img",
    ],
    cmd = """
    $(location @//cmd/mkimage) $(location @//build/linux_kernel:image) $@
    """,
    visibility = ["//visibility:public"],
)

genrule(
    name = "swtpm_data",
    outs = [
        "tpm/tpm2-00.permall",
    ],
    tags = ["local"],
    cmd = """
    mkdir tpm

    swtpm_setup \
        --tpmstate tpm \
        --create-ek-cert \
        --create-platform-cert \
        --allow-signing \
        --tpm2 \
        --display \
        --pcr-banks sha1,sha256,sha384,sha512

    cp tpm/tpm2-00.permall $@
    """,
    visibility = ["//visibility:public"],
)
