load("//core/build:def.bzl", "smalltown_initramfs")

smalltown_initramfs(
    name = "initramfs",
    extra_dirs = [
        "/kubernetes/conf/flexvolume-plugins",
        "/containerd/run",
    ],
    files = {
        "//core/cmd/init": "/init",
        "//third_party/xfsprogs:mkfs.xfs": "/bin/mkfs.xfs",

        # CA Certificate bundle & os-release
        "@cacerts//file": "/etc/ssl/cert.pem",
        ":os-release-info": "/etc/os-release",

        # Hyperkube
        "//core/cmd/kube": "/kubernetes/bin/kube",

        # runsc/gVisor
        "@com_github_google_gvisor//runsc": "/containerd/bin/runsc",
        "@com_github_google_gvisor_containerd_shim//cmd/containerd-shim-runsc-v1": "/containerd/bin/containerd-shim-runsc-v1",

        # Containerd
        "@com_github_containerd_containerd//cmd/containerd": "/containerd/bin/containerd",

        # Containerd config files
        "//core/internal/containerd:runsc.toml": "/containerd/conf/runsc.toml",
        "//core/internal/containerd:config.toml": "/containerd/conf/config.toml",
        "//core/internal/containerd:loopback.json": "/containerd/conf/cni/99-loopback.conf",
        "//core/internal/containerd:ptp.json": "/containerd/conf/cni/10-ptp.conf",

        # CNI Plugins
        "@com_github_containernetworking_plugins//plugins/main/loopback": "/containerd/bin/cni/loopback",
        "@com_github_containernetworking_plugins//plugins/main/ptp": "/containerd/bin/cni/ptp",
        "@com_github_containernetworking_plugins//plugins/ipam/host-local": "/containerd/bin/cni/host-local",
    },
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
