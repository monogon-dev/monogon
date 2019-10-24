load("@bazel_gazelle//:def.bzl", "gazelle")

# gazelle:prefix git.monogon.dev/source/smalltown.git
# gazelle:exclude generated
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
        "tpm/signkey.pem",
        "tpm/issuercert.pem",
    ],
    cmd = """
    mkdir -p tpm/ca

    cat <<EOF > tpm/swtpm.conf
create_certs_tool= /usr/share/swtpm/swtpm-localca
create_certs_tool_config = tpm/swtpm-localca.conf
create_certs_tool_options = /etc/swtpm-localca.options
EOF

    cat <<EOF > tpm/swtpm-localca.conf
statedir = tpm/ca
signingkey = tpm/ca/signkey.pem
issuercert = tpm/ca/issuercert.pem
certserial = tpm/ca/certserial
EOF

    swtpm_setup \
        --tpmstate tpm \
        --create-ek-cert \
        --create-platform-cert \
        --allow-signing \
        --tpm2 \
        --display \
        --pcr-banks sha1,sha256,sha384,sha512 \
        --config tpm/swtpm.conf

    cp tpm/tpm2-00.permall $(location tpm/tpm2-00.permall)
    cp tpm/ca/issuercert.pem $(location tpm/issuercert.pem)
    cp tpm/ca/signkey.pem $(location tpm/signkey.pem)
    """,
    visibility = ["//visibility:public"],
)
