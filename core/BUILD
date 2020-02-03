genrule(
    name = "initramfs",
    srcs = [
        "//core/cmd/init",
        "//core/build/utils:mkfs.xfs",
        "//core/cmd/kube-controlplane",
    ],
    outs = [
        "initramfs.cpio.lz4",
    ],
    cmd = """
    $(location @linux//:gen_init_cpio) - <<- 'EOF' | lz4 -l > \"$@\" 
dir /dev 0755 0 0
nod /dev/console 0600 0 0 c 5 1
nod /dev/null 0644 0 0 c 1 3
file /init $(location //core/cmd/init) 0755 0 0
dir /bin 0755 0 0
file /bin/mkfs.xfs $(location //core/build/utils:mkfs.xfs) 0755 0 0
file /bin/kube-controlplane $(location //core/cmd/kube-controlplane) 0755 0 0
EOF
    """,
    tools = [
        "@linux//:gen_init_cpio",
    ],
)

genrule(
    name = "image",
    srcs = [
        "//third_party/linux:bzImage",
        ":initramfs",
    ],
    outs = [
        "smalltown.img",
    ],
    cmd = """
    $(location //core/cmd/mkimage) \
        -efi $(location //third_party/linux:bzImage) \
        -initramfs $(location :initramfs) \
        -out $@
    """,
    tools = [
        "//core/cmd/mkimage",
    ],
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
