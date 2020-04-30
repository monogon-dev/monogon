genrule(
    name = "initramfs",
    srcs = [
        "//core/cmd/init",
        "//core/cmd/kube-controlplane",
        "//third_party/xfsprogs:mkfs.xfs",
        "@io_k8s_kubernetes//cmd/kubelet:_kubelet-pure",
        "@com_github_containerd_containerd//cmd/containerd",
        "@com_github_containerd_containerd//cmd/containerd-shim",
        "@com_github_containerd_containerd//cmd/containerd-shim-runc-v1",
        "@com_github_containerd_containerd//cmd/containerd-shim-runc-v2",
        "@com_github_containernetworking_plugins//plugins/main/loopback",
        "@com_github_containernetworking_plugins//plugins/main/ptp",
        "@com_github_containernetworking_plugins//plugins/ipam/host-local",
        "@com_github_opencontainers_runc//:runc",
        "@com_github_google_gvisor//runsc",
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
file /bin/mkfs.xfs $(location //third_party/xfsprogs:mkfs.xfs) 0755 0 0
file /bin/kube-controlplane $(location //core/cmd/kube-controlplane) 0755 0 0
file /bin/kubelet $(location @io_k8s_kubernetes//cmd/kubelet:_kubelet-pure) 0755 0 0
dir /containerd 0755 0 0
file /containerd/containerd $(location @com_github_containerd_containerd//cmd/containerd) 0755 0 0
file /containerd/containerd-shim $(location @com_github_containerd_containerd//cmd/containerd-shim) 0755 0 0
file /containerd/containerd-shim-runc-v1 $(location @com_github_containerd_containerd//cmd/containerd-shim-runc-v1) 0755 0 0
file /containerd/containerd-shim-runc-v2 $(location @com_github_containerd_containerd//cmd/containerd-shim-runc-v2) 0755 0 0
file /containerd/runsc $(location @com_github_google_gvisor//runsc) 0755 0 0
file /containerd/runc $(location @com_github_opencontainers_runc//:runc) 0755 0 0
dir /containerd/cni-plugins 0755 0 0
file /containerd/cni-plugins/loopback $(location @com_github_containernetworking_plugins//plugins/main/loopback) 0755 0 0
file /containerd/cni-plugins/ptp $(location @com_github_containernetworking_plugins//plugins/main/ptp) 0755 0 0
file /containerd/cni-plugins/host-local $(location @com_github_containernetworking_plugins//plugins/ipam/host-local) 0755 0 0
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
