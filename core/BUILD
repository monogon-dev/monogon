genrule(
    name = "initramfs",
    srcs = [
        "//core/cmd/init",
        "//core/cmd/kube",
        "//third_party/xfsprogs:mkfs.xfs",
        "@io_k8s_kubernetes//cmd/kubelet:_kubelet-pure",
        "@com_github_containerd_containerd//cmd/containerd",
        "@com_github_containerd_containerd//cmd/containerd-shim-runc-v2",
        "@com_github_containernetworking_plugins//plugins/main/loopback",
        "@com_github_containernetworking_plugins//plugins/main/ptp",
        "@com_github_containernetworking_plugins//plugins/ipam/host-local",
        "@com_github_opencontainers_runc//:runc",
        "@com_github_google_gvisor//runsc",
        "@com_github_google_gvisor_containerd_shim//cmd/containerd-shim-runsc-v1",
        "//core/internal/containerd:ptp.json",
        "//core/internal/containerd:loopback.json",
        "//core/internal/containerd:config.toml",
        "//core/internal/containerd:runsc.toml",
        "@cacerts//file",
        ":os-release-info",
    ],
    outs = [
        "initramfs.cpio.lz4",
    ],
    cmd = """
    $(location @linux//:gen_init_cpio) - <<- 'EOF' | lz4 -l > \"$@\" 
dir /dev 0755 0 0
nod /dev/console 0600 0 0 c 5 1
nod /dev/null 0644 0 0 c 1 3
nod /dev/kmsg 0644 0 0 c 1 11
nod /dev/ptmx 0644 0 0 c 5 2
file /init $(location //core/cmd/init) 0755 0 0
dir /etc 0755 0 0
file /etc/os-release $(location :os-release-info) 0644 0 0
dir /etc/ssl 0755 0 0
file /etc/ssl/cert.pem $(location @cacerts//file) 0444 0 0
dir /bin 0755 0 0
file /bin/mkfs.xfs $(location //third_party/xfsprogs:mkfs.xfs) 0755 0 0
dir /kubernetes 0755 0 0
dir /kubernetes/bin 0755 0 0
file /kubernetes/bin/kube $(location //core/cmd/kube) 0755 0 0
dir /kubernetes/conf 0755 0 0
dir /kubernetes/conf/flexvolume-plugins 0755 0 0
dir /containerd 0755 0 0
dir /containerd/bin 0755 0 0
file /containerd/bin/containerd $(location @com_github_containerd_containerd//cmd/containerd) 0755 0 0
file /containerd/bin/containerd-shim-runsc-v1 $(location @com_github_google_gvisor_containerd_shim//cmd/containerd-shim-runsc-v1) 0755 0 0
file /containerd/bin/runsc $(location @com_github_google_gvisor//runsc) 0755 0 0
dir /containerd/bin/cni 0755 0 0
file /containerd/bin/cni/loopback $(location @com_github_containernetworking_plugins//plugins/main/loopback) 0755 0 0
file /containerd/bin/cni/ptp $(location @com_github_containernetworking_plugins//plugins/main/ptp) 0755 0 0
file /containerd/bin/cni/host-local $(location @com_github_containernetworking_plugins//plugins/ipam/host-local) 0755 0 0
dir /containerd/run 0755 0 0
dir /containerd/conf 0755 0 0
dir /containerd/conf/cni 0755 0 0
file /containerd/conf/cni/10-ptp.conf $(location //core/internal/containerd:ptp.json) 0444 0 0
file /containerd/conf/cni/99-loopback.conf $(location //core/internal/containerd:loopback.json) 0444 0 0
file /containerd/conf/config.toml $(location //core/internal/containerd:config.toml) 0444 0 0
file /containerd/conf/runsc.toml $(location //core/internal/containerd:runsc.toml) 0444 0 0
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

load("//core/build/genosrelease:defs.bzl", "os_release")

os_release(
    name = "os-release-info",
    os_id = "smalltown",
    os_name = "Smalltown",
    stamp_var = "STABLE_SIGNOS_version",
)
