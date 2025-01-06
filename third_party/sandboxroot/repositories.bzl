load("@bazeldnf//:deps.bzl", "rpm")

def sandbox_dependencies():
    rpm(
        name = "acpica-tools-0__20220331-8.fc40.aarch64",
        sha256 = "78532f8e3494df416221aacf699da51857ba242018e9fa8dc96e4d5527ce77a9",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/a/acpica-tools-20220331-8.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/a/acpica-tools-20220331-8.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/a/acpica-tools-20220331-8.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/a/acpica-tools-20220331-8.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/a/acpica-tools-20220331-8.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "alternatives-0__1.27-1.fc40.aarch64",
        sha256 = "f7178972cc3f7b68a6f024b4a5673ad81ba284f7f7d9afe3a875487704b9b200",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/a/alternatives-1.27-1.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/a/alternatives-1.27-1.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/a/alternatives-1.27-1.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/a/alternatives-1.27-1.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/a/alternatives-1.27-1.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "ansible-srpm-macros-0__1-16.fc40.aarch64",
        sha256 = "a221968063ee17b8d4ee3e7013d40b2789638a76ce8e94ebd15d694f6b48b4bd",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/a/ansible-srpm-macros-1-16.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/a/ansible-srpm-macros-1-16.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/a/ansible-srpm-macros-1-16.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/a/ansible-srpm-macros-1-16.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/a/ansible-srpm-macros-1-16.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "audit-libs-0__4.0.2-1.fc40.aarch64",
        sha256 = "ce61feca5ff38e9d9a0fd4607dc363e1a513d8661cea77b76fa7005bf73c7a21",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/a/audit-libs-4.0.2-1.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/a/audit-libs-4.0.2-1.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/a/audit-libs-4.0.2-1.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/a/audit-libs-4.0.2-1.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/a/audit-libs-4.0.2-1.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "authselect-0__1.5.0-6.fc40.aarch64",
        sha256 = "83af73051df98a25a8727c5b12445e570eb9c748a0309c16fe187e702b3404a1",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/a/authselect-1.5.0-6.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/a/authselect-1.5.0-6.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/a/authselect-1.5.0-6.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/a/authselect-1.5.0-6.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/a/authselect-1.5.0-6.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "authselect-libs-0__1.5.0-6.fc40.aarch64",
        sha256 = "a2b8cab9da7670c4cbc0064553d15fb8c289f2f2eeedc5e90ccaf8bb1e0b3c0e",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/a/authselect-libs-1.5.0-6.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/a/authselect-libs-1.5.0-6.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/a/authselect-libs-1.5.0-6.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/a/authselect-libs-1.5.0-6.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/a/authselect-libs-1.5.0-6.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "basesystem-0__11-20.fc40.aarch64",
        sha256 = "6404b1028262aeaf3e083f08959969abea1301f7f5e8610492cf900b3d13d5db",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/b/basesystem-11-20.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/b/basesystem-11-20.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/b/basesystem-11-20.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/b/basesystem-11-20.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/b/basesystem-11-20.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "bash-0__5.2.26-3.fc40.aarch64",
        sha256 = "c514625af42909a46825315f5aa6b57f61a93ac80103401ae810dbb4576ed51b",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/b/bash-5.2.26-3.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/b/bash-5.2.26-3.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/b/bash-5.2.26-3.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/b/bash-5.2.26-3.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/b/bash-5.2.26-3.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "bc-0__1.07.1-21.fc40.aarch64",
        sha256 = "ca9fba2eca165ee988e40e4661f8518eb41728305d1635df943832ce6d251511",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/b/bc-1.07.1-21.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/b/bc-1.07.1-21.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/b/bc-1.07.1-21.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/b/bc-1.07.1-21.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/b/bc-1.07.1-21.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "binutils-0__2.41-38.fc40.aarch64",
        sha256 = "c571abac49e12ab6af9494b72f0adf54f66c26f5d8d29a727bcdb6ed2e20d8bb",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/b/binutils-2.41-38.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/b/binutils-2.41-38.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/b/binutils-2.41-38.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/b/binutils-2.41-38.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/b/binutils-2.41-38.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "binutils-gold-0__2.41-38.fc40.aarch64",
        sha256 = "4c050a2a4ddd69e5e16401f37f37084f6fd5d2a2ab847b20eb87d7b2e888e76f",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/b/binutils-gold-2.41-38.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/b/binutils-gold-2.41-38.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/b/binutils-gold-2.41-38.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/b/binutils-gold-2.41-38.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/b/binutils-gold-2.41-38.fc40.aarch64.rpm",
        ],
    )

    rpm(
        name = "bison-0__3.8.2-7.fc40.aarch64",
        sha256 = "fcf8266cfd6a82bbdc8b0b96ecce4f14f8d0115d6af00e79a60bfded03059b04",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/b/bison-3.8.2-7.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/b/bison-3.8.2-7.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/b/bison-3.8.2-7.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/b/bison-3.8.2-7.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/b/bison-3.8.2-7.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "bzip2-libs-0__1.0.8-18.fc40.aarch64",
        sha256 = "914d171e310ce5f59f0d3017c2e0ec67a91e08e1b940fab1e595cc15584c5dd2",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/b/bzip2-libs-1.0.8-18.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/b/bzip2-libs-1.0.8-18.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/b/bzip2-libs-1.0.8-18.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/b/bzip2-libs-1.0.8-18.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/b/bzip2-libs-1.0.8-18.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "ca-certificates-0__2024.2.69_v8.0.401-1.0.fc40.aarch64",
        sha256 = "1afcf80d5e7b22ee512ec9f24b4f2b148888ef95af3486cf48f2204c3406b12d",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/c/ca-certificates-2024.2.69_v8.0.401-1.0.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/c/ca-certificates-2024.2.69_v8.0.401-1.0.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/c/ca-certificates-2024.2.69_v8.0.401-1.0.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/c/ca-certificates-2024.2.69_v8.0.401-1.0.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/c/ca-certificates-2024.2.69_v8.0.401-1.0.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "capstone-0__5.0.1-3.fc40.aarch64",
        sha256 = "5adb30db6b49d69b168e0a166ecba504693aaef09c5cb60a2cafafe0a465ac3d",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/c/capstone-5.0.1-3.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/c/capstone-5.0.1-3.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/c/capstone-5.0.1-3.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/c/capstone-5.0.1-3.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/c/capstone-5.0.1-3.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "checkpolicy-0__3.7-2.fc40.aarch64",
        sha256 = "1f5ffd99939e10a408949bb5106acc6170eb1b5813eca8f59689abd7c5f9b78c",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/c/checkpolicy-3.7-2.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/c/checkpolicy-3.7-2.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/c/checkpolicy-3.7-2.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/c/checkpolicy-3.7-2.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/c/checkpolicy-3.7-2.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "clang-0__18.1.8-1.fc40.aarch64",
        sha256 = "d3ff619048e64d7a4f5be370b7de74cdd1bbae30b2a7a7bbd0d7bd2c85bd4c03",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/c/clang-18.1.8-1.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/c/clang-18.1.8-1.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/c/clang-18.1.8-1.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/c/clang-18.1.8-1.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/c/clang-18.1.8-1.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "clang-libs-0__18.1.8-1.fc40.aarch64",
        sha256 = "967ee7eaf486542746900f8408beb0111c9c7a82f6d61fc2a44fc144dd4496e7",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/c/clang-libs-18.1.8-1.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/c/clang-libs-18.1.8-1.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/c/clang-libs-18.1.8-1.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/c/clang-libs-18.1.8-1.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/c/clang-libs-18.1.8-1.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "clang-resource-filesystem-0__18.1.8-1.fc40.aarch64",
        sha256 = "eed9d12d800d7e2d8d5d81fa9f90e11095e21f8f29e1be219a95b50558295db9",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/c/clang-resource-filesystem-18.1.8-1.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/c/clang-resource-filesystem-18.1.8-1.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/c/clang-resource-filesystem-18.1.8-1.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/c/clang-resource-filesystem-18.1.8-1.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/c/clang-resource-filesystem-18.1.8-1.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "cmake-filesystem-0__3.30.5-1.fc40.aarch64",
        sha256 = "8b33d0c73235f8837f0ed8034570d17cca89bd560b4a26e135349f4a409ba316",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/c/cmake-filesystem-3.30.5-1.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/c/cmake-filesystem-3.30.5-1.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/c/cmake-filesystem-3.30.5-1.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/c/cmake-filesystem-3.30.5-1.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/c/cmake-filesystem-3.30.5-1.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "coreutils-single-0__9.4-9.fc40.aarch64",
        sha256 = "bf7c0c6667c49c305f1b0e4704386928566fdb12963a2382322db9aec5107eac",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/c/coreutils-single-9.4-9.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/c/coreutils-single-9.4-9.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/c/coreutils-single-9.4-9.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/c/coreutils-single-9.4-9.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/c/coreutils-single-9.4-9.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "cpp-0__14.2.1-3.fc40.aarch64",
        sha256 = "089dd79d4551d691f1ef916122f6cde79b98a1c649d2a22bec10980bc8890f26",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/c/cpp-14.2.1-3.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/c/cpp-14.2.1-3.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/c/cpp-14.2.1-3.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/c/cpp-14.2.1-3.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/c/cpp-14.2.1-3.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "cracklib-0__2.9.11-5.fc40.aarch64",
        sha256 = "0d843063ab65f3449abf0b3e6b52c5ee19dcb3a94e6d32bceb4feca9bfc8ca5c",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/c/cracklib-2.9.11-5.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/c/cracklib-2.9.11-5.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/c/cracklib-2.9.11-5.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/c/cracklib-2.9.11-5.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/c/cracklib-2.9.11-5.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "crypto-policies-0__20241011-1.git5930b9a.fc40.aarch64",
        sha256 = "d7a62ff0193375607d28d8fe7eedf3ff5b6ddac154e1474d79787b9f32ae298d",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/c/crypto-policies-20241011-1.git5930b9a.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/c/crypto-policies-20241011-1.git5930b9a.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/c/crypto-policies-20241011-1.git5930b9a.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/c/crypto-policies-20241011-1.git5930b9a.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/c/crypto-policies-20241011-1.git5930b9a.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "curl-0__8.6.0-10.fc40.aarch64",
        sha256 = "188adeed41a595a89749fbe86d5f608c466d3cdb017cd0a45b6f3a3d5abd93b2",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/c/curl-8.6.0-10.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/c/curl-8.6.0-10.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/c/curl-8.6.0-10.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/c/curl-8.6.0-10.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/c/curl-8.6.0-10.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "cyrus-sasl-lib-0__2.1.28-19.fc40.aarch64",
        sha256 = "0c42d638ae119b1c3f50d785b8e7b14641db8665f4367a57e6299da58e2bb2b7",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/c/cyrus-sasl-lib-2.1.28-19.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/c/cyrus-sasl-lib-2.1.28-19.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/c/cyrus-sasl-lib-2.1.28-19.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/c/cyrus-sasl-lib-2.1.28-19.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/c/cyrus-sasl-lib-2.1.28-19.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "daxctl-libs-0__80-1.fc40.aarch64",
        sha256 = "735def87639016ee72e94020e78313c97ede16b82abf5165d53145dae9df2179",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/d/daxctl-libs-80-1.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/d/daxctl-libs-80-1.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/d/daxctl-libs-80-1.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/d/daxctl-libs-80-1.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/d/daxctl-libs-80-1.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "dbus-1__1.14.10-3.fc40.aarch64",
        sha256 = "02b5cfeeba2a7a1eae5be0df735a966ab1cc235d04b3c0b834ccfd77409e681a",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/d/dbus-1.14.10-3.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/d/dbus-1.14.10-3.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/d/dbus-1.14.10-3.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/d/dbus-1.14.10-3.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/d/dbus-1.14.10-3.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "dbus-broker-0__36-2.fc40.aarch64",
        sha256 = "e105b73c57be1dc1d6453e2c6da8dc2c4bbc7a955823fdb0ae7dbdc47a4df635",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/d/dbus-broker-36-2.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/d/dbus-broker-36-2.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/d/dbus-broker-36-2.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/d/dbus-broker-36-2.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/d/dbus-broker-36-2.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "dbus-common-1__1.14.10-3.fc40.aarch64",
        sha256 = "81bade4072aca4f5d22be29a916d9d0cfc9262a6c5d92ddfe750f7b8bf03f7c9",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/d/dbus-common-1.14.10-3.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/d/dbus-common-1.14.10-3.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/d/dbus-common-1.14.10-3.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/d/dbus-common-1.14.10-3.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/d/dbus-common-1.14.10-3.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "diffutils-0__3.10-5.fc40.aarch64",
        sha256 = "7ddb22d4bf02a59817f978a2f7a0eb17e45876904f32e2968dcf142d44927ad3",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/d/diffutils-3.10-5.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/d/diffutils-3.10-5.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/d/diffutils-3.10-5.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/d/diffutils-3.10-5.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/d/diffutils-3.10-5.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "dwz-0__0.15-8.fc40.aarch64",
        sha256 = "c4abb41fe58359bd7fd2231e04537ed0f50adc5a12645d36402dec5076eeaa09",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/d/dwz-0.15-8.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/d/dwz-0.15-8.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/d/dwz-0.15-8.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/d/dwz-0.15-8.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/d/dwz-0.15-8.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "e2fsprogs-libs-0__1.47.0-5.fc40.aarch64",
        sha256 = "4d3db8e501c3c4f2fe2ae46bfe1bac88f8908e60d263846f1900105e56e10fe4",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/e/e2fsprogs-libs-1.47.0-5.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/e/e2fsprogs-libs-1.47.0-5.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/e/e2fsprogs-libs-1.47.0-5.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/e/e2fsprogs-libs-1.47.0-5.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/e/e2fsprogs-libs-1.47.0-5.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "ed-0__1.20.2-1.fc40.aarch64",
        sha256 = "28a2b81ee22b8bc1ba882476e4041f78aa0160e9d03d375c4af582abfdd46c3c",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/e/ed-1.20.2-1.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/e/ed-1.20.2-1.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/e/ed-1.20.2-1.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/e/ed-1.20.2-1.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/e/ed-1.20.2-1.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "edk2-aarch64-0__20241117-5.fc40.aarch64",
        sha256 = "327bc420740c6c0ed15652727c40c72898017a2440b9fd2c7ad4b0bbab0c2b06",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/e/edk2-aarch64-20241117-5.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/e/edk2-aarch64-20241117-5.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/e/edk2-aarch64-20241117-5.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/e/edk2-aarch64-20241117-5.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/e/edk2-aarch64-20241117-5.fc40.noarch.rpm",
        ],
    )

    rpm(
        name = "edk2-ovmf-0__20241117-5.fc40.aarch64",
        sha256 = "c367fb9da8630a0d227207d8f9b72eb48a687a35c9d2921605b001d5ff7df7ba",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/e/edk2-ovmf-20241117-5.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/e/edk2-ovmf-20241117-5.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/e/edk2-ovmf-20241117-5.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/e/edk2-ovmf-20241117-5.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/e/edk2-ovmf-20241117-5.fc40.noarch.rpm",
        ],
    )

    rpm(
        name = "efi-srpm-macros-0__5-11.fc40.aarch64",
        sha256 = "34ed8bd59f9b299975871ebce1d15208cd66a4383f46a4f0d75e01303bacac2c",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/e/efi-srpm-macros-5-11.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/e/efi-srpm-macros-5-11.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/e/efi-srpm-macros-5-11.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/e/efi-srpm-macros-5-11.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/e/efi-srpm-macros-5-11.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "elfutils-debuginfod-client-0__0.192-7.fc40.aarch64",
        sha256 = "71b88a7d90e259185fca766882e991778d8f14124e7a72813fffa84b9540d567",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/e/elfutils-debuginfod-client-0.192-7.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/e/elfutils-debuginfod-client-0.192-7.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/e/elfutils-debuginfod-client-0.192-7.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/e/elfutils-debuginfod-client-0.192-7.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/e/elfutils-debuginfod-client-0.192-7.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "elfutils-default-yama-scope-0__0.192-7.fc40.aarch64",
        sha256 = "0df162d8f4bddb1b8f9e8f4b94777f06c31024fa1db83a956817e6e1f7b6b6f1",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/e/elfutils-default-yama-scope-0.192-7.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/e/elfutils-default-yama-scope-0.192-7.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/e/elfutils-default-yama-scope-0.192-7.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/e/elfutils-default-yama-scope-0.192-7.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/e/elfutils-default-yama-scope-0.192-7.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "elfutils-libelf-0__0.192-7.fc40.aarch64",
        sha256 = "0ea9fab59239e1ca715da9f0181bd477fab854a4a12918cc792fdca8497ffd3c",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/e/elfutils-libelf-0.192-7.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/e/elfutils-libelf-0.192-7.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/e/elfutils-libelf-0.192-7.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/e/elfutils-libelf-0.192-7.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/e/elfutils-libelf-0.192-7.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "elfutils-libelf-devel-0__0.192-7.fc40.aarch64",
        sha256 = "b86565050362a9b76ee57995d40e16ac3fb3f53cd97ab61bfdd594d385a348d4",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/e/elfutils-libelf-devel-0.192-7.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/e/elfutils-libelf-devel-0.192-7.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/e/elfutils-libelf-devel-0.192-7.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/e/elfutils-libelf-devel-0.192-7.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/e/elfutils-libelf-devel-0.192-7.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "elfutils-libs-0__0.192-7.fc40.aarch64",
        sha256 = "817bf6df648972d603dac2c2d96b0d3df54c5ae2feed1c2b47aa018bc7213bdd",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/e/elfutils-libs-0.192-7.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/e/elfutils-libs-0.192-7.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/e/elfutils-libs-0.192-7.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/e/elfutils-libs-0.192-7.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/e/elfutils-libs-0.192-7.fc40.aarch64.rpm",
        ],
    )

    rpm(
        name = "expat-0__2.6.3-1.fc40.aarch64",
        sha256 = "d36ce90295fb55d4fc19e324f501060351e08fa1d759695b8ff022b4f3d5d2e0",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/e/expat-2.6.3-1.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/e/expat-2.6.3-1.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/e/expat-2.6.3-1.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/e/expat-2.6.3-1.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/e/expat-2.6.3-1.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "fedora-gpg-keys-0__40-2.aarch64",
        sha256 = "849feb04544096f9bbe16bc78c2198708fe658bdafa08575c911e538a7d31c18",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/f/fedora-gpg-keys-40-2.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/f/fedora-gpg-keys-40-2.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/f/fedora-gpg-keys-40-2.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/f/fedora-gpg-keys-40-2.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/f/fedora-gpg-keys-40-2.noarch.rpm",
        ],
    )
    rpm(
        name = "fedora-release-common-0__40-40.aarch64",
        sha256 = "dde6f4b5d4415ce20df40cf1cb9ff3015aa5b1896c5b5625e49aa686cdce1d1d",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/f/fedora-release-common-40-40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/f/fedora-release-common-40-40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/f/fedora-release-common-40-40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/f/fedora-release-common-40-40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/f/fedora-release-common-40-40.noarch.rpm",
        ],
    )
    rpm(
        name = "fedora-release-container-0__40-40.aarch64",
        sha256 = "4d071c5ec7069a0c92689ea42aeb481e682cc0e2df61260d703c2a351bbc43cf",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/f/fedora-release-container-40-40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/f/fedora-release-container-40-40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/f/fedora-release-container-40-40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/f/fedora-release-container-40-40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/f/fedora-release-container-40-40.noarch.rpm",
        ],
    )
    rpm(
        name = "fedora-release-identity-container-0__40-40.aarch64",
        sha256 = "4467c98d42b1090af85b670dd31f8f0c3ce4a750e1c7ccd73f3a2b04afda381a",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/f/fedora-release-identity-container-40-40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/f/fedora-release-identity-container-40-40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/f/fedora-release-identity-container-40-40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/f/fedora-release-identity-container-40-40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/f/fedora-release-identity-container-40-40.noarch.rpm",
        ],
    )
    rpm(
        name = "fedora-repos-0__40-2.aarch64",
        sha256 = "e85d69eeea62f4f5a7c6584bc8bae3cb559c1c381838ca89f7d63b28d2368c4b",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/f/fedora-repos-40-2.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/f/fedora-repos-40-2.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/f/fedora-repos-40-2.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/f/fedora-repos-40-2.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/f/fedora-repos-40-2.noarch.rpm",
        ],
    )
    rpm(
        name = "file-0__5.45-4.fc40.aarch64",
        sha256 = "7dd6dd85ee6c8caa0890eb5e8d8fec1dcf6c246e88c02d93f8e36fc8d8785729",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/f/file-5.45-4.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/f/file-5.45-4.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/f/file-5.45-4.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/f/file-5.45-4.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/f/file-5.45-4.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "file-libs-0__5.45-4.fc40.aarch64",
        sha256 = "0486187bf43bb560234917b4dfc0ce2c3b422f1a0a87e031419ca16f41abc1c1",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/f/file-libs-5.45-4.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/f/file-libs-5.45-4.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/f/file-libs-5.45-4.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/f/file-libs-5.45-4.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/f/file-libs-5.45-4.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "filesystem-0__3.18-8.fc40.aarch64",
        sha256 = "b133dafefe7916b576aac366509ad2d78682df60802753792f9b084909031446",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/f/filesystem-3.18-8.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/f/filesystem-3.18-8.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/f/filesystem-3.18-8.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/f/filesystem-3.18-8.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/f/filesystem-3.18-8.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "findutils-1__4.9.0-9.fc40.aarch64",
        sha256 = "1d1e1c78d3cdb61c5440f324524f56901c1eb0710c028078b0869e9477eecb41",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/f/findutils-4.9.0-9.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/f/findutils-4.9.0-9.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/f/findutils-4.9.0-9.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/f/findutils-4.9.0-9.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/f/findutils-4.9.0-9.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "flex-0__2.6.4-16.fc40.aarch64",
        sha256 = "ac58b7e80c27b8e08c871537a7987b3b5047cd15e3bf04e40db36c160f623b7b",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/f/flex-2.6.4-16.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/f/flex-2.6.4-16.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/f/flex-2.6.4-16.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/f/flex-2.6.4-16.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/f/flex-2.6.4-16.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "fonts-srpm-macros-1__2.0.5-14.fc40.aarch64",
        sha256 = "ebf245973cea76d51b22de0e587fc77b3d6a776fb629c4130971182536afd9d7",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/f/fonts-srpm-macros-2.0.5-14.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/f/fonts-srpm-macros-2.0.5-14.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/f/fonts-srpm-macros-2.0.5-14.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/f/fonts-srpm-macros-2.0.5-14.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/f/fonts-srpm-macros-2.0.5-14.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "forge-srpm-macros-0__0.4.0-1.fc40.aarch64",
        sha256 = "05d2dd0de0c5efe676e4f4f5a7b4c184ee6a018e4ea2ce04507f565c6f893377",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/f/forge-srpm-macros-0.4.0-1.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/f/forge-srpm-macros-0.4.0-1.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/f/forge-srpm-macros-0.4.0-1.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/f/forge-srpm-macros-0.4.0-1.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/f/forge-srpm-macros-0.4.0-1.fc40.noarch.rpm",
        ],
    )

    rpm(
        name = "fpc-srpm-macros-0__1.3-12.fc40.aarch64",
        sha256 = "7df65ab4ab462818320c8391aa8b08e63fddba2c60944e40f0b207118effbae5",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/f/fpc-srpm-macros-1.3-12.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/f/fpc-srpm-macros-1.3-12.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/f/fpc-srpm-macros-1.3-12.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/f/fpc-srpm-macros-1.3-12.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/f/fpc-srpm-macros-1.3-12.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "fuse3-libs-0__3.16.2-3.fc40.aarch64",
        sha256 = "f89d257984f063b798e2fa5ac6d1fe780fcc25cb2432a3ab05f8d973e110cebf",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/f/fuse3-libs-3.16.2-3.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/f/fuse3-libs-3.16.2-3.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/f/fuse3-libs-3.16.2-3.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/f/fuse3-libs-3.16.2-3.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/f/fuse3-libs-3.16.2-3.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "gawk-0__5.3.0-3.fc40.aarch64",
        sha256 = "db53a17eda9d5d7515a38300b83ec8c7e16dfedd75f3b63ad64082622ad3f86a",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/g/gawk-5.3.0-3.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/g/gawk-5.3.0-3.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/g/gawk-5.3.0-3.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/g/gawk-5.3.0-3.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/g/gawk-5.3.0-3.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "gc-0__8.2.2-6.fc40.aarch64",
        sha256 = "b83f473cc569a8cc0f0323d4209ae8153aefd57daaf51bab4b31c7a5dfeac1d4",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/g/gc-8.2.2-6.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/g/gc-8.2.2-6.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/g/gc-8.2.2-6.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/g/gc-8.2.2-6.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/g/gc-8.2.2-6.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "gcc-0__14.2.1-3.fc40.aarch64",
        sha256 = "a41dd464e6d066c99140bec107cc8d65c04c8d0fa74bc12d0e5cb604a9ae662c",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/g/gcc-14.2.1-3.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/g/gcc-14.2.1-3.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/g/gcc-14.2.1-3.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/g/gcc-14.2.1-3.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/g/gcc-14.2.1-3.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "gcc-c__plus____plus__-0__14.2.1-3.fc40.aarch64",
        sha256 = "f4d58d7f65ac51b14bc7c20ffcfaafd12bc53714b4bf5d18fd6cf5f8c79c7ad8",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/g/gcc-c++-14.2.1-3.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/g/gcc-c++-14.2.1-3.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/g/gcc-c++-14.2.1-3.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/g/gcc-c++-14.2.1-3.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/g/gcc-c++-14.2.1-3.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "gdbm-1__1.23-6.fc40.aarch64",
        sha256 = "ff17ebd4b5f9b34dc9ef24d9cecf9deaac0dd4eb7c1100161f00dcd799ac1dde",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/g/gdbm-1.23-6.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/g/gdbm-1.23-6.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/g/gdbm-1.23-6.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/g/gdbm-1.23-6.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/g/gdbm-1.23-6.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "gdbm-libs-1__1.23-6.fc40.aarch64",
        sha256 = "f72db6cba1a3dcc921f49e230382b8be5b76ee48313535f93a1ad9959ad322a2",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/g/gdbm-libs-1.23-6.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/g/gdbm-libs-1.23-6.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/g/gdbm-libs-1.23-6.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/g/gdbm-libs-1.23-6.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/g/gdbm-libs-1.23-6.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "ghc-srpm-macros-0__1.9.1-1.fc40.aarch64",
        sha256 = "1509ca46a18243b3f181aac3d77639b805c470816f892fbf62acd0ae96f01f9a",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/g/ghc-srpm-macros-1.9.1-1.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/g/ghc-srpm-macros-1.9.1-1.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/g/ghc-srpm-macros-1.9.1-1.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/g/ghc-srpm-macros-1.9.1-1.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/g/ghc-srpm-macros-1.9.1-1.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "glib2-0__2.80.3-1.fc40.aarch64",
        sha256 = "50abb096ccb73dc2bafdafe48246e60173ecafd6fa2806ca306d9c5f78077034",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/g/glib2-2.80.3-1.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/g/glib2-2.80.3-1.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/g/glib2-2.80.3-1.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/g/glib2-2.80.3-1.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/g/glib2-2.80.3-1.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "glibc-0__2.39-33.fc40.aarch64",
        sha256 = "e6c090eb68b323df96ae9433fa3954cfacc62dd4be9e8bc4beb2fbb53e8b1cb7",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/g/glibc-2.39-33.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/g/glibc-2.39-33.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/g/glibc-2.39-33.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/g/glibc-2.39-33.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/g/glibc-2.39-33.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "glibc-common-0__2.39-33.fc40.aarch64",
        sha256 = "8e40d90ac3e2c536b0aa94ea4a6e3895ac82a4190c4b4530efe68693c3da4cc6",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/g/glibc-common-2.39-33.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/g/glibc-common-2.39-33.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/g/glibc-common-2.39-33.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/g/glibc-common-2.39-33.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/g/glibc-common-2.39-33.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "glibc-devel-0__2.39-33.fc40.aarch64",
        sha256 = "a7d1e5b962ad0b28c95559edaaab5dcc9ef1b77e6d850a0c167b71770042dffd",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/g/glibc-devel-2.39-33.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/g/glibc-devel-2.39-33.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/g/glibc-devel-2.39-33.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/g/glibc-devel-2.39-33.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/g/glibc-devel-2.39-33.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "glibc-langpack-en-0__2.39-33.fc40.aarch64",
        sha256 = "7f064ade36f61700a1e7d87f98050df4ba09765b41c4d4579d7aeef3b48c511b",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/g/glibc-langpack-en-2.39-33.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/g/glibc-langpack-en-2.39-33.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/g/glibc-langpack-en-2.39-33.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/g/glibc-langpack-en-2.39-33.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/g/glibc-langpack-en-2.39-33.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "glibc-static-0__2.39-33.fc40.aarch64",
        sha256 = "baca890c7de356277ff28305767a394a5d12c617e8189a93795404c56317b32c",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/g/glibc-static-2.39-33.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/g/glibc-static-2.39-33.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/g/glibc-static-2.39-33.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/g/glibc-static-2.39-33.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/g/glibc-static-2.39-33.fc40.aarch64.rpm",
        ],
    )

    rpm(
        name = "gmp-1__6.2.1-8.fc40.aarch64",
        sha256 = "b9ee08361f4d0936efb771d7391df1838f90703d3d80caf3186c105024391a62",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/g/gmp-6.2.1-8.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/g/gmp-6.2.1-8.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/g/gmp-6.2.1-8.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/g/gmp-6.2.1-8.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/g/gmp-6.2.1-8.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "gnat-srpm-macros-0__6-5.fc40.aarch64",
        sha256 = "35f84a6494aed02d6a2b90f702787232962535c313ab56b3878b264a6c39546c",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/g/gnat-srpm-macros-6-5.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/g/gnat-srpm-macros-6-5.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/g/gnat-srpm-macros-6-5.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/g/gnat-srpm-macros-6-5.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/g/gnat-srpm-macros-6-5.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "gnupg2-0__2.4.4-1.fc40.aarch64",
        sha256 = "7d59174bea245c486b849abe3896f49c4bfe05980a850220337d43f1dfc41065",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/g/gnupg2-2.4.4-1.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/g/gnupg2-2.4.4-1.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/g/gnupg2-2.4.4-1.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/g/gnupg2-2.4.4-1.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/g/gnupg2-2.4.4-1.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "gnutls-0__3.8.6-1.fc40.aarch64",
        sha256 = "428a64ab9491e036c26517229fe39a8c0a5c8d4c347e3754c0f2bea6f0fc0d2d",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/g/gnutls-3.8.6-1.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/g/gnutls-3.8.6-1.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/g/gnutls-3.8.6-1.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/g/gnutls-3.8.6-1.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/g/gnutls-3.8.6-1.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "go-srpm-macros-0__3.5.0-1.fc40.aarch64",
        sha256 = "2968803f0da871cb5b5873efab1360871260c915338f72f11486a1210aafd105",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/g/go-srpm-macros-3.5.0-1.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/g/go-srpm-macros-3.5.0-1.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/g/go-srpm-macros-3.5.0-1.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/g/go-srpm-macros-3.5.0-1.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/g/go-srpm-macros-3.5.0-1.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "grep-0__3.11-7.fc40.aarch64",
        sha256 = "03f4ec56345751bc0017999dc9fde06d7b37f4fea73477f3c2aa4b7bb9136b2e",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/g/grep-3.11-7.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/g/grep-3.11-7.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/g/grep-3.11-7.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/g/grep-3.11-7.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/g/grep-3.11-7.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "groff-base-0__1.23.0-6.fc40.aarch64",
        sha256 = "5473558eebd5530d96a4a1cd69d96854e675076473153ebe7f1fd187c5aa2e65",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/g/groff-base-1.23.0-6.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/g/groff-base-1.23.0-6.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/g/groff-base-1.23.0-6.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/g/groff-base-1.23.0-6.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/g/groff-base-1.23.0-6.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "guile30-0__3.0.7-12.fc40.aarch64",
        sha256 = "5da87e4b41eaa9ebbb8c114777f0395310523809711cbf7eeab4dd07d3b3d6cd",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/g/guile30-3.0.7-12.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/g/guile30-3.0.7-12.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/g/guile30-3.0.7-12.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/g/guile30-3.0.7-12.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/g/guile30-3.0.7-12.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "gzip-0__1.13-1.fc40.aarch64",
        sha256 = "1a688cf61e685dd989f0a2adf8cde29abe5af5da616eac95c0a8d1646e8f0a3a",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/g/gzip-1.13-1.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/g/gzip-1.13-1.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/g/gzip-1.13-1.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/g/gzip-1.13-1.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/g/gzip-1.13-1.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "ipxe-roms-qemu-0__20240119-1.gitde8a0821.fc40.aarch64",
        sha256 = "1d84c5d480e0f23c4dfda72ff6db466d4959941d897fe517a9771112d41203bc",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/i/ipxe-roms-qemu-20240119-1.gitde8a0821.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/i/ipxe-roms-qemu-20240119-1.gitde8a0821.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/i/ipxe-roms-qemu-20240119-1.gitde8a0821.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/i/ipxe-roms-qemu-20240119-1.gitde8a0821.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/i/ipxe-roms-qemu-20240119-1.gitde8a0821.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "jansson-0__2.13.1-9.fc40.aarch64",
        sha256 = "1743dd69769199069f89c47490f2058f739b975641b40476ec10caec59d1895e",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/j/jansson-2.13.1-9.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/j/jansson-2.13.1-9.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/j/jansson-2.13.1-9.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/j/jansson-2.13.1-9.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/j/jansson-2.13.1-9.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "json-c-0__0.17-3.fc40.aarch64",
        sha256 = "45d9a506c790620da6df3fb94965dc1339ee449840af1566a5ea0e081b8f2bfc",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/j/json-c-0.17-3.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/j/json-c-0.17-3.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/j/json-c-0.17-3.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/j/json-c-0.17-3.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/j/json-c-0.17-3.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "kernel-headers-0__6.12.4-100.fc40.aarch64",
        sha256 = "d693e7b57777f2adaf55284fb38f6129cef7b32c46006cbdef4b286667b72d48",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/k/kernel-headers-6.12.4-100.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/k/kernel-headers-6.12.4-100.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/k/kernel-headers-6.12.4-100.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/k/kernel-headers-6.12.4-100.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/k/kernel-headers-6.12.4-100.fc40.aarch64.rpm",
        ],
    )

    rpm(
        name = "kernel-srpm-macros-0__1.0-23.fc40.aarch64",
        sha256 = "95fb5031a23336455d606d05c63855c7f12247ffd4baaac64fb576b420b2a32e",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/k/kernel-srpm-macros-1.0-23.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/k/kernel-srpm-macros-1.0-23.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/k/kernel-srpm-macros-1.0-23.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/k/kernel-srpm-macros-1.0-23.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/k/kernel-srpm-macros-1.0-23.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "keyutils-libs-0__1.6.3-3.fc40.aarch64",
        sha256 = "6d9451953678a00075b99f83d822f9496e3b7b6a0697043a8b8bb9a006ddf9be",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/k/keyutils-libs-1.6.3-3.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/k/keyutils-libs-1.6.3-3.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/k/keyutils-libs-1.6.3-3.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/k/keyutils-libs-1.6.3-3.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/k/keyutils-libs-1.6.3-3.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "kmod-libs-0__31-5.fc40.aarch64",
        sha256 = "2e4b2268fd442687fee1af1478de3c6a6d97db1281500496c47ada4fa31ea6ae",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/k/kmod-libs-31-5.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/k/kmod-libs-31-5.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/k/kmod-libs-31-5.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/k/kmod-libs-31-5.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/k/kmod-libs-31-5.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "krb5-libs-0__1.21.3-2.fc40.aarch64",
        sha256 = "2d302466e8eeaa380583270540642c4e9c06a7b4e57528235cd7a2bf955bd403",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/k/krb5-libs-1.21.3-2.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/k/krb5-libs-1.21.3-2.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/k/krb5-libs-1.21.3-2.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/k/krb5-libs-1.21.3-2.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/k/krb5-libs-1.21.3-2.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "libacl-0__2.3.2-1.fc40.aarch64",
        sha256 = "83b2655e7af5bafa4b0e565b3aefcc61777a49c13d781a2b7f0e40f3583684f1",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libacl-2.3.2-1.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libacl-2.3.2-1.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libacl-2.3.2-1.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libacl-2.3.2-1.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libacl-2.3.2-1.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "libaio-0__0.3.111-19.fc40.aarch64",
        sha256 = "c388d9694d06c590021e1c31eda3e9c89deb2e8d2234206057a9484a2e0fea30",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libaio-0.3.111-19.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libaio-0.3.111-19.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libaio-0.3.111-19.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libaio-0.3.111-19.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libaio-0.3.111-19.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "libarchive-0__3.7.2-7.fc40.aarch64",
        sha256 = "208cdc76a122d8bf094c558d21534be01eaceefdef219c3fdaf67ad972b7ec18",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/l/libarchive-3.7.2-7.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/l/libarchive-3.7.2-7.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/l/libarchive-3.7.2-7.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/l/libarchive-3.7.2-7.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/l/libarchive-3.7.2-7.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "libasan-0__14.2.1-3.fc40.aarch64",
        sha256 = "2dddf4d89a7b0b534987e198821974f05521ca8e323aba77476a6f19459c8704",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/l/libasan-14.2.1-3.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/l/libasan-14.2.1-3.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/l/libasan-14.2.1-3.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/l/libasan-14.2.1-3.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/l/libasan-14.2.1-3.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "libassuan-0__2.5.7-1.fc40.aarch64",
        sha256 = "87797a07fee5a4a7ad92ef6a0e0bd79c688994ead797a651373908d5d8265885",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libassuan-2.5.7-1.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libassuan-2.5.7-1.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libassuan-2.5.7-1.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libassuan-2.5.7-1.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libassuan-2.5.7-1.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "libatomic-0__14.2.1-3.fc40.aarch64",
        sha256 = "f7daf6311766565baeba2c96c8adb528342e9c9ef3cf8a63d0f1bb8dfdcb3d33",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/l/libatomic-14.2.1-3.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/l/libatomic-14.2.1-3.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/l/libatomic-14.2.1-3.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/l/libatomic-14.2.1-3.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/l/libatomic-14.2.1-3.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "libattr-0__2.5.2-3.fc40.aarch64",
        sha256 = "e9a846d5bc41951e1dfd37e09805e7da554b4e171688e5750d6ddf0d4f05855e",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libattr-2.5.2-3.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libattr-2.5.2-3.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libattr-2.5.2-3.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libattr-2.5.2-3.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libattr-2.5.2-3.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "libb2-0__0.98.1-11.fc40.aarch64",
        sha256 = "616bff6f22b8c860442d035cc25f053a61b5a44ce6867022e3c342feedfa3f15",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libb2-0.98.1-11.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libb2-0.98.1-11.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libb2-0.98.1-11.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libb2-0.98.1-11.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libb2-0.98.1-11.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "libblkid-0__2.40.2-1.fc40.aarch64",
        sha256 = "34e8c1aff28bd843759c21ae81a648dae0b1d1326c22f22772e17fe9cf82ad37",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/l/libblkid-2.40.2-1.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/l/libblkid-2.40.2-1.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/l/libblkid-2.40.2-1.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/l/libblkid-2.40.2-1.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/l/libblkid-2.40.2-1.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "libbpf-2__1.2.3-1.fc40.aarch64",
        sha256 = "b8ef77e666082e474766b0aa5b8bdb4e26bf5ce04aca2454041b63fbca08bcc6",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/l/libbpf-1.2.3-1.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/l/libbpf-1.2.3-1.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/l/libbpf-1.2.3-1.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/l/libbpf-1.2.3-1.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/l/libbpf-1.2.3-1.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "libbrotli-0__1.1.0-3.fc40.aarch64",
        sha256 = "618bf1f5510e74828f81d29bf1223f14ad785c6dc1130a4694f54c2ab254e499",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libbrotli-1.1.0-3.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libbrotli-1.1.0-3.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libbrotli-1.1.0-3.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libbrotli-1.1.0-3.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libbrotli-1.1.0-3.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "libcap-0__2.69-8.fc40.aarch64",
        sha256 = "dc57063dd37a71c9b1201b3344ca1e2f3a41522bba80cf74000e9d41c7098204",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/l/libcap-2.69-8.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/l/libcap-2.69-8.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/l/libcap-2.69-8.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/l/libcap-2.69-8.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/l/libcap-2.69-8.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "libcap-ng-0__0.8.4-4.fc40.aarch64",
        sha256 = "c03ba261a23def29f6e0c54a98c76e94c74d29a776418c313c11556199ecc0cc",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libcap-ng-0.8.4-4.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libcap-ng-0.8.4-4.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libcap-ng-0.8.4-4.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libcap-ng-0.8.4-4.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libcap-ng-0.8.4-4.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "libcom_err-0__1.47.0-5.fc40.aarch64",
        sha256 = "42347c3974d2e2f6202a9b8aaa9a9d3bb149bd31ce0a2f7bab054dbfdf377378",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libcom_err-1.47.0-5.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libcom_err-1.47.0-5.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libcom_err-1.47.0-5.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libcom_err-1.47.0-5.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libcom_err-1.47.0-5.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "libcurl-0__8.6.0-10.fc40.aarch64",
        sha256 = "a8cf09f7bed02b82ee1b756c1f4ddac669bf572cacf44f5771510d7c6c684f41",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/l/libcurl-8.6.0-10.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/l/libcurl-8.6.0-10.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/l/libcurl-8.6.0-10.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/l/libcurl-8.6.0-10.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/l/libcurl-8.6.0-10.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "libdb-0__5.3.28-62.fc40.aarch64",
        sha256 = "c1d2eda0e1adfd8294305a6794370530767865802f1735e120774cac2b0a4a05",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/l/libdb-5.3.28-62.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/l/libdb-5.3.28-62.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/l/libdb-5.3.28-62.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/l/libdb-5.3.28-62.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/l/libdb-5.3.28-62.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "libeconf-0__0.6.2-2.fc40.aarch64",
        sha256 = "533262f82b2a732ace8f389a549cfe5bb4d1776be8068d7a13f510a4f9f36c6b",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/l/libeconf-0.6.2-2.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/l/libeconf-0.6.2-2.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/l/libeconf-0.6.2-2.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/l/libeconf-0.6.2-2.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/l/libeconf-0.6.2-2.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "libedit-0__3.1-53.20240808cvs.fc40.aarch64",
        sha256 = "2153535b643c40c7f6e67799db114234d4cc0468fa2298b235fc8f6f9414282b",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/l/libedit-3.1-53.20240808cvs.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/l/libedit-3.1-53.20240808cvs.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/l/libedit-3.1-53.20240808cvs.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/l/libedit-3.1-53.20240808cvs.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/l/libedit-3.1-53.20240808cvs.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "libevent-0__2.1.12-12.fc40.aarch64",
        sha256 = "98e25aa83f69978e75cb2fceaca358daadca1a90f5a8c933167306c5d0d74e4b",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libevent-2.1.12-12.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libevent-2.1.12-12.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libevent-2.1.12-12.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libevent-2.1.12-12.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libevent-2.1.12-12.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "libfdisk-0__2.40.2-1.fc40.aarch64",
        sha256 = "93efe859c2c5fac77ce7e2caf5c6ecb0830cd4bcadbaf226cef13d5cff3636a3",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/l/libfdisk-2.40.2-1.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/l/libfdisk-2.40.2-1.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/l/libfdisk-2.40.2-1.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/l/libfdisk-2.40.2-1.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/l/libfdisk-2.40.2-1.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "libfdt-0__1.7.0-7.fc40.aarch64",
        sha256 = "c1f4759d1372eb8d11986833939f598820f3ed3d273e912dcb5e5934cffe14ec",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libfdt-1.7.0-7.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libfdt-1.7.0-7.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libfdt-1.7.0-7.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libfdt-1.7.0-7.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libfdt-1.7.0-7.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "libffi-0__3.4.4-7.fc40.aarch64",
        sha256 = "133a37ea1f55d8580457209611badd30ea644bdc9b328c77fcdd0c6849a48f0b",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libffi-3.4.4-7.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libffi-3.4.4-7.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libffi-3.4.4-7.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libffi-3.4.4-7.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libffi-3.4.4-7.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "libgcc-0__14.2.1-3.fc40.aarch64",
        sha256 = "6f46237aff8ab4400834a89a259c162eb3ba8431447a36b6be3d50fe20f2f798",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/l/libgcc-14.2.1-3.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/l/libgcc-14.2.1-3.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/l/libgcc-14.2.1-3.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/l/libgcc-14.2.1-3.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/l/libgcc-14.2.1-3.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "libgcrypt-0__1.10.3-3.fc40.aarch64",
        sha256 = "cd19d396e322ac30d99c61d38d076e49a1e16dda65ff378cd31df609d86f5e4d",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libgcrypt-1.10.3-3.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libgcrypt-1.10.3-3.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libgcrypt-1.10.3-3.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libgcrypt-1.10.3-3.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libgcrypt-1.10.3-3.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "libgomp-0__14.2.1-3.fc40.aarch64",
        sha256 = "8507b7f6fc83b7b39c8eb976879864003f711ca3e7bcc52ab96824993fa7d37c",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/l/libgomp-14.2.1-3.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/l/libgomp-14.2.1-3.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/l/libgomp-14.2.1-3.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/l/libgomp-14.2.1-3.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/l/libgomp-14.2.1-3.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "libgpg-error-0__1.49-1.fc40.aarch64",
        sha256 = "c23695f248a678ed9339a064ad7a8a4bdd080c4de048f880fdf940876982f5de",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/l/libgpg-error-1.49-1.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/l/libgpg-error-1.49-1.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/l/libgpg-error-1.49-1.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/l/libgpg-error-1.49-1.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/l/libgpg-error-1.49-1.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "libibverbs-0__48.0-4.fc40.aarch64",
        sha256 = "3d3c18db3f6a74b73d08df2d0e10bca54a7a3c18a89a204cea539eb738d5b638",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libibverbs-48.0-4.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libibverbs-48.0-4.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libibverbs-48.0-4.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libibverbs-48.0-4.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libibverbs-48.0-4.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "libidn2-0__2.3.7-1.fc40.aarch64",
        sha256 = "ce71114618eb8dc7dcac5f3431d11e99e0eb0b377412533a03fc7a5746c3c3f7",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libidn2-2.3.7-1.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libidn2-2.3.7-1.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libidn2-2.3.7-1.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libidn2-2.3.7-1.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libidn2-2.3.7-1.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "libjpeg-turbo-0__3.0.2-1.fc40.aarch64",
        sha256 = "83754a768b06a0a926c794e4ee07c1e09f694990d2c8445679b39c13253150db",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libjpeg-turbo-3.0.2-1.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libjpeg-turbo-3.0.2-1.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libjpeg-turbo-3.0.2-1.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libjpeg-turbo-3.0.2-1.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libjpeg-turbo-3.0.2-1.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "libksba-0__1.6.6-1.fc40.aarch64",
        sha256 = "db0bb13c2624a87b98ff7ae320decbb6c782d0097e9f77f13026b7da2ed66abe",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libksba-1.6.6-1.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libksba-1.6.6-1.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libksba-1.6.6-1.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libksba-1.6.6-1.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libksba-1.6.6-1.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "libmount-0__2.40.2-1.fc40.aarch64",
        sha256 = "3b75009fdca3a0d11bd0d41c7d5e6512f590e6fb2d5dfa5cd5942e89d4086253",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/l/libmount-2.40.2-1.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/l/libmount-2.40.2-1.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/l/libmount-2.40.2-1.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/l/libmount-2.40.2-1.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/l/libmount-2.40.2-1.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "libmpc-0__1.3.1-5.fc40.aarch64",
        sha256 = "0dc2a5f4ee53e67c960e3e579ece8dade1f2329fd32a95a74273321b61c5416e",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libmpc-1.3.1-5.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libmpc-1.3.1-5.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libmpc-1.3.1-5.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libmpc-1.3.1-5.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libmpc-1.3.1-5.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "libnghttp2-0__1.59.0-3.fc40.aarch64",
        sha256 = "f0048920aa021f9244aa77e41057a8319981de91b1c709955ec75972f2d2c07e",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/l/libnghttp2-1.59.0-3.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/l/libnghttp2-1.59.0-3.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/l/libnghttp2-1.59.0-3.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/l/libnghttp2-1.59.0-3.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/l/libnghttp2-1.59.0-3.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "libnl3-0__3.11.0-1.fc40.aarch64",
        sha256 = "61ab82a83b506760cc74ef58d361a7411954634dafe6083fd8650f6abf853c21",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/l/libnl3-3.11.0-1.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/l/libnl3-3.11.0-1.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/l/libnl3-3.11.0-1.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/l/libnl3-3.11.0-1.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/l/libnl3-3.11.0-1.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "libnsl2-0__2.0.1-1.fc40.aarch64",
        sha256 = "c80b5c544889cffbef8c72b9f244f05d8934b124f80f0814109d73c72c8ec362",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libnsl2-2.0.1-1.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libnsl2-2.0.1-1.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libnsl2-2.0.1-1.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libnsl2-2.0.1-1.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libnsl2-2.0.1-1.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "libpkgconf-0__2.1.1-2.fc40.aarch64",
        sha256 = "78b5d7731e0b571e19e3bdbc9234459ff8f5d6c24f979ea296253c91fcdbfbcc",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/l/libpkgconf-2.1.1-2.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/l/libpkgconf-2.1.1-2.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/l/libpkgconf-2.1.1-2.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/l/libpkgconf-2.1.1-2.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/l/libpkgconf-2.1.1-2.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "libpng-2__1.6.40-3.fc40.aarch64",
        sha256 = "1ca5cb0fccb57535567c1dde79360ffd0ec058abc1603d0bf8f138b443225157",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libpng-1.6.40-3.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libpng-1.6.40-3.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libpng-1.6.40-3.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libpng-1.6.40-3.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libpng-1.6.40-3.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "libpsl-0__0.21.5-3.fc40.aarch64",
        sha256 = "d22eef5093cfab189a31e4e96342d6975746e73d5e1a33b3405cb83fdd2fe3f6",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libpsl-0.21.5-3.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libpsl-0.21.5-3.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libpsl-0.21.5-3.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libpsl-0.21.5-3.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libpsl-0.21.5-3.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "libpwquality-0__1.4.5-9.fc40.aarch64",
        sha256 = "5b8ccd23aba57b04dc8fdc8c20bdf4c6719a7eb9cf496a47955f14f6b8ac2fa9",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libpwquality-1.4.5-9.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libpwquality-1.4.5-9.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libpwquality-1.4.5-9.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libpwquality-1.4.5-9.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libpwquality-1.4.5-9.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "librdmacm-0__48.0-4.fc40.aarch64",
        sha256 = "89a9928fb62c82c8be623c472528e14168dbc0a3296a0ed705c68aa249c2f14b",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/librdmacm-48.0-4.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/librdmacm-48.0-4.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/librdmacm-48.0-4.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/librdmacm-48.0-4.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/librdmacm-48.0-4.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "libseccomp-0__2.5.5-1.fc40.aarch64",
        sha256 = "58c77d1209dae9041bd443e07e3e17a1b12d558306d3e51dac67270c19ada065",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/l/libseccomp-2.5.5-1.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/l/libseccomp-2.5.5-1.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/l/libseccomp-2.5.5-1.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/l/libseccomp-2.5.5-1.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/l/libseccomp-2.5.5-1.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "libselinux-0__3.7-5.fc40.aarch64",
        sha256 = "6749cf3d0ef1283e06267793b39f3a6c1d7621f7c801abe77b66042736613016",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/l/libselinux-3.7-5.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/l/libselinux-3.7-5.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/l/libselinux-3.7-5.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/l/libselinux-3.7-5.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/l/libselinux-3.7-5.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "libselinux-utils-0__3.7-5.fc40.aarch64",
        sha256 = "bbb3fedfe27caaa173550d30f3f9c876ccc4852f864603e62b9e3b973bfbcbfb",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/l/libselinux-utils-3.7-5.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/l/libselinux-utils-3.7-5.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/l/libselinux-utils-3.7-5.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/l/libselinux-utils-3.7-5.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/l/libselinux-utils-3.7-5.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "libsemanage-0__3.7-2.fc40.aarch64",
        sha256 = "1390e081daa9ab41d727b0f94711a60a93ba18e517f4cdc259bf302b0d4893a0",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/l/libsemanage-3.7-2.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/l/libsemanage-3.7-2.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/l/libsemanage-3.7-2.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/l/libsemanage-3.7-2.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/l/libsemanage-3.7-2.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "libsepol-0__3.7-2.fc40.aarch64",
        sha256 = "a1d10394b4f6e5a52e7ec3af47fb956e1852cbb98e1c90c73675fbb03d12bbc4",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/l/libsepol-3.7-2.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/l/libsepol-3.7-2.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/l/libsepol-3.7-2.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/l/libsepol-3.7-2.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/l/libsepol-3.7-2.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "libslirp-0__4.7.0-6.fc40.aarch64",
        sha256 = "b639409969fc4ce8a4d85442ed2598e802b0a2b4cca3dcd646f80df88ff5af22",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libslirp-4.7.0-6.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libslirp-4.7.0-6.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libslirp-4.7.0-6.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libslirp-4.7.0-6.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libslirp-4.7.0-6.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "libsmartcols-0__2.40.2-1.fc40.aarch64",
        sha256 = "33ef1577c807aa22ee53154f939d54fd198f83cc176d4a956616b8d07847a4d0",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/l/libsmartcols-2.40.2-1.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/l/libsmartcols-2.40.2-1.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/l/libsmartcols-2.40.2-1.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/l/libsmartcols-2.40.2-1.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/l/libsmartcols-2.40.2-1.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "libssh-0__0.10.6-5.fc40.aarch64",
        sha256 = "05d68c3013e527528ec9979e787b7a27698df191f133eb52a2509d4acedc7acd",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libssh-0.10.6-5.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libssh-0.10.6-5.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libssh-0.10.6-5.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libssh-0.10.6-5.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libssh-0.10.6-5.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "libssh-config-0__0.10.6-5.fc40.aarch64",
        sha256 = "241c73071a373732ec544dad6ba6f4fb054c1f2264d86085c322dd1c1089bbb1",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libssh-config-0.10.6-5.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libssh-config-0.10.6-5.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libssh-config-0.10.6-5.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libssh-config-0.10.6-5.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libssh-config-0.10.6-5.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "libstdc__plus____plus__-0__14.2.1-3.fc40.aarch64",
        sha256 = "ee557bf6eb6bc0e1c6194acb241abf662d5853e8f5f0baadd053a79506aaa002",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/l/libstdc++-14.2.1-3.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/l/libstdc++-14.2.1-3.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/l/libstdc++-14.2.1-3.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/l/libstdc++-14.2.1-3.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/l/libstdc++-14.2.1-3.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "libstdc__plus____plus__-devel-0__14.2.1-3.fc40.aarch64",
        sha256 = "c60a2e95be5f56896d196b5df938e0606c0146c6d523203a49d455db5b378e84",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/l/libstdc++-devel-14.2.1-3.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/l/libstdc++-devel-14.2.1-3.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/l/libstdc++-devel-14.2.1-3.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/l/libstdc++-devel-14.2.1-3.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/l/libstdc++-devel-14.2.1-3.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "libstdc__plus____plus__-static-0__14.2.1-3.fc40.aarch64",
        sha256 = "c5cb34cb0fc0b17c345fd08c4914930361fd88e79f181d8cca05ff061ad96ff5",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/l/libstdc++-static-14.2.1-3.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/l/libstdc++-static-14.2.1-3.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/l/libstdc++-static-14.2.1-3.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/l/libstdc++-static-14.2.1-3.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/l/libstdc++-static-14.2.1-3.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "libtasn1-0__4.19.0-6.fc40.aarch64",
        sha256 = "db3594dc4b74f6e826c25d422e299adb4e01b8002ffe553047dfa9f34136bb0f",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libtasn1-4.19.0-6.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libtasn1-4.19.0-6.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libtasn1-4.19.0-6.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libtasn1-4.19.0-6.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libtasn1-4.19.0-6.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "libtirpc-0__1.3.6-1.fc40.aarch64",
        sha256 = "2479e84de9ab06268c24f20758787cfb5b62473597333fb2b4d578e5ca5bba91",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/l/libtirpc-1.3.6-1.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/l/libtirpc-1.3.6-1.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/l/libtirpc-1.3.6-1.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/l/libtirpc-1.3.6-1.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/l/libtirpc-1.3.6-1.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "libtool-ltdl-0__2.4.7-10.fc40.aarch64",
        sha256 = "f0b063b2abca597c74847ea66791b0ccb4d4adda66866502429bfd64d0cab883",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libtool-ltdl-2.4.7-10.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libtool-ltdl-2.4.7-10.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libtool-ltdl-2.4.7-10.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libtool-ltdl-2.4.7-10.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libtool-ltdl-2.4.7-10.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "libubsan-0__14.2.1-3.fc40.aarch64",
        sha256 = "2eb23572090847661f6fbb4a973e2111b4dde8c3a74253dd059060b194d30270",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/l/libubsan-14.2.1-3.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/l/libubsan-14.2.1-3.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/l/libubsan-14.2.1-3.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/l/libubsan-14.2.1-3.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/l/libubsan-14.2.1-3.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "libunistring-0__1.1-7.fc40.aarch64",
        sha256 = "b3f9b5628b42715eac842859d7f13b66ad36a7b4f21a594788f9b5293614bfa0",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libunistring-1.1-7.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libunistring-1.1-7.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libunistring-1.1-7.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libunistring-1.1-7.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libunistring-1.1-7.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "liburing-0__2.5-3.fc40.aarch64",
        sha256 = "838313b928ac67ead6415d80de7ca4778f5ce12c31d6f125cc805217f19cf1d8",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/liburing-2.5-3.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/liburing-2.5-3.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/liburing-2.5-3.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/liburing-2.5-3.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/liburing-2.5-3.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "libutempter-0__1.2.1-13.fc40.aarch64",
        sha256 = "cc4ad04dc16c67f850a9b5df442bba66835b2b138c299dcf77548c073f4e2f0b",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libutempter-1.2.1-13.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libutempter-1.2.1-13.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libutempter-1.2.1-13.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libutempter-1.2.1-13.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libutempter-1.2.1-13.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "libuuid-0__2.40.2-1.fc40.aarch64",
        sha256 = "5446f61b560a4c9b5bcf6c7c222465660b7aa12d994126b43e02183a615f8298",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/l/libuuid-2.40.2-1.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/l/libuuid-2.40.2-1.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/l/libuuid-2.40.2-1.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/l/libuuid-2.40.2-1.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/l/libuuid-2.40.2-1.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "libuuid-devel-0__2.40.2-1.fc40.aarch64",
        sha256 = "fffe2575e00c39f10c4177ef8999911c1d10f4bf6fe282f49e73483aaabe9835",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/l/libuuid-devel-2.40.2-1.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/l/libuuid-devel-2.40.2-1.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/l/libuuid-devel-2.40.2-1.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/l/libuuid-devel-2.40.2-1.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/l/libuuid-devel-2.40.2-1.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "libverto-0__0.3.2-8.fc40.aarch64",
        sha256 = "d17feb9d8beb1a445c9c45022fb475af435d82d77b5ae9c71b1156146f5e2470",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libverto-0.3.2-8.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libverto-0.3.2-8.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libverto-0.3.2-8.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libverto-0.3.2-8.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libverto-0.3.2-8.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "libxcrypt-0__4.4.37-1.fc40.aarch64",
        sha256 = "021bbdb17a43abb52be41880a378ce090a2c342471cc0b6a78a401365f73640d",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/l/libxcrypt-4.4.37-1.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/l/libxcrypt-4.4.37-1.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/l/libxcrypt-4.4.37-1.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/l/libxcrypt-4.4.37-1.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/l/libxcrypt-4.4.37-1.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "libxcrypt-devel-0__4.4.37-1.fc40.aarch64",
        sha256 = "54f23728e9984565717a4449ec131781733b292681b3f319f6b1583fc88a6fd7",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/l/libxcrypt-devel-4.4.37-1.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/l/libxcrypt-devel-4.4.37-1.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/l/libxcrypt-devel-4.4.37-1.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/l/libxcrypt-devel-4.4.37-1.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/l/libxcrypt-devel-4.4.37-1.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "libxcrypt-static-0__4.4.37-1.fc40.aarch64",
        sha256 = "d9a2f686136a6c7a8cfd2a638ec20ce7cae8fd2304c9d817359e9e964f0e9593",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/l/libxcrypt-static-4.4.37-1.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/l/libxcrypt-static-4.4.37-1.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/l/libxcrypt-static-4.4.37-1.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/l/libxcrypt-static-4.4.37-1.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/l/libxcrypt-static-4.4.37-1.fc40.aarch64.rpm",
        ],
    )

    rpm(
        name = "libxdp-0__1.4.2-1.fc40.aarch64",
        sha256 = "1f29669cd32f0ef075394cdcf28ed71171f8b49520e904750e1c59f0240873e8",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libxdp-1.4.2-1.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libxdp-1.4.2-1.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libxdp-1.4.2-1.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libxdp-1.4.2-1.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/libxdp-1.4.2-1.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "libxml2-0__2.12.9-1.fc40.aarch64",
        sha256 = "1f2d8a4a0a54a083ad6268abaf071d409a2081273fd7f16007bca8c697b08234",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/l/libxml2-2.12.9-1.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/l/libxml2-2.12.9-1.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/l/libxml2-2.12.9-1.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/l/libxml2-2.12.9-1.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/l/libxml2-2.12.9-1.fc40.aarch64.rpm",
        ],
    )

    rpm(
        name = "libzstd-0__1.5.6-1.fc40.aarch64",
        sha256 = "d4f3833bd371dcb3907469bab2feb0405e6153e1ac678b4b445bb31bbfa68163",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/l/libzstd-1.5.6-1.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/l/libzstd-1.5.6-1.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/l/libzstd-1.5.6-1.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/l/libzstd-1.5.6-1.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/l/libzstd-1.5.6-1.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "libzstd-devel-0__1.5.6-1.fc40.aarch64",
        sha256 = "ef7da914b60d1fa6185072f265a2b89b3c46c4f9b54f70336b16d26c0168fe93",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/l/libzstd-devel-1.5.6-1.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/l/libzstd-devel-1.5.6-1.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/l/libzstd-devel-1.5.6-1.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/l/libzstd-devel-1.5.6-1.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/l/libzstd-devel-1.5.6-1.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "lld-0__18.1.8-1.fc40.aarch64",
        sha256 = "82e230168051e4565bc615e798d77e952988eeee16ecda749f6e7d524f737ab5",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/l/lld-18.1.8-1.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/l/lld-18.1.8-1.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/l/lld-18.1.8-1.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/l/lld-18.1.8-1.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/l/lld-18.1.8-1.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "lld-libs-0__18.1.8-1.fc40.aarch64",
        sha256 = "a772b5d579715c3c4cebf914e82e55969cdbd8f5592dff8b0a876b2a31ab77f2",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/l/lld-libs-18.1.8-1.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/l/lld-libs-18.1.8-1.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/l/lld-libs-18.1.8-1.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/l/lld-libs-18.1.8-1.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/l/lld-libs-18.1.8-1.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "llvm-0__18.1.8-2.fc40.aarch64",
        sha256 = "fbfde41912eb947418766e55088dfc3e2a33a712964ef5d2a154386368f4559a",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/l/llvm-18.1.8-2.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/l/llvm-18.1.8-2.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/l/llvm-18.1.8-2.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/l/llvm-18.1.8-2.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/l/llvm-18.1.8-2.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "llvm-libs-0__18.1.8-2.fc40.aarch64",
        sha256 = "faa554653b45d676e570faad024b081dd821e3801ecad5ce784880fa7f380b30",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/l/llvm-libs-18.1.8-2.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/l/llvm-libs-18.1.8-2.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/l/llvm-libs-18.1.8-2.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/l/llvm-libs-18.1.8-2.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/l/llvm-libs-18.1.8-2.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "lua-libs-0__5.4.6-5.fc40.aarch64",
        sha256 = "d26750a2b2a885cfce42f66df28655167854532279fb6cac980e2a494fbb0a67",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/lua-libs-5.4.6-5.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/lua-libs-5.4.6-5.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/lua-libs-5.4.6-5.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/lua-libs-5.4.6-5.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/lua-libs-5.4.6-5.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "lua-srpm-macros-0__1-13.fc40.aarch64",
        sha256 = "959030121201a706bc620d311569f15ab81bafdb9e3de94bf763a72df36d15f0",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/lua-srpm-macros-1-13.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/lua-srpm-macros-1-13.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/lua-srpm-macros-1-13.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/lua-srpm-macros-1-13.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/lua-srpm-macros-1-13.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "lz4-0__1.9.4-6.fc40.aarch64",
        sha256 = "d6dde8a36494226cadbd311ab3138258340cf3d52e81734186ae66665f41f3b5",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/lz4-1.9.4-6.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/lz4-1.9.4-6.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/lz4-1.9.4-6.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/lz4-1.9.4-6.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/lz4-1.9.4-6.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "lz4-libs-0__1.9.4-6.fc40.aarch64",
        sha256 = "4f20d0adcd229dfc54d2681747302fc618895f4445bc8a67b16b4295d0485140",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/lz4-libs-1.9.4-6.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/lz4-libs-1.9.4-6.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/lz4-libs-1.9.4-6.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/lz4-libs-1.9.4-6.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/lz4-libs-1.9.4-6.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "lzo-0__2.10-12.fc40.aarch64",
        sha256 = "3ca4fcfe3eddf8727841ccd725f0e1c5eea38720298a33bd2233f1766a15ebb6",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/lzo-2.10-12.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/lzo-2.10-12.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/lzo-2.10-12.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/lzo-2.10-12.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/l/lzo-2.10-12.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "m4-0__1.4.19-9.fc40.aarch64",
        sha256 = "cbe631d931a719d2c85c627b64cff295f3755ded3670d72533d26d9116086794",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/m/m4-1.4.19-9.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/m/m4-1.4.19-9.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/m/m4-1.4.19-9.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/m/m4-1.4.19-9.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/m/m4-1.4.19-9.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "make-1__4.4.1-6.fc40.aarch64",
        sha256 = "a5bdce585a02c57526695dd7e7166e6562302ce305cc0a2c021530ab5ca7a5cb",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/m/make-4.4.1-6.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/m/make-4.4.1-6.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/m/make-4.4.1-6.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/m/make-4.4.1-6.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/m/make-4.4.1-6.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "mpdecimal-0__2.5.1-9.fc40.aarch64",
        sha256 = "68137e90d691569add6694ae78c3aad1b976a0f03bc9239068bc9f84aaf5c6c6",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/m/mpdecimal-2.5.1-9.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/m/mpdecimal-2.5.1-9.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/m/mpdecimal-2.5.1-9.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/m/mpdecimal-2.5.1-9.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/m/mpdecimal-2.5.1-9.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "mpfr-0__4.2.1-4.fc40.aarch64",
        sha256 = "2fb2998fabde72bdc984ef245867e56ee7e79658ed268871dd30109bcc9943fe",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/m/mpfr-4.2.1-4.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/m/mpfr-4.2.1-4.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/m/mpfr-4.2.1-4.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/m/mpfr-4.2.1-4.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/m/mpfr-4.2.1-4.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "nasm-0__2.16.01-7.fc40.aarch64",
        sha256 = "7387a53632440abf396b83645f0e638ff7a9027e9bbb28ccab20e8b98d6fa457",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/n/nasm-2.16.01-7.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/n/nasm-2.16.01-7.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/n/nasm-2.16.01-7.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/n/nasm-2.16.01-7.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/n/nasm-2.16.01-7.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "ncurses-0__6.4-12.20240127.fc40.aarch64",
        sha256 = "566bb290dd96ed630c6dbb61c473b72f1bdb38f1eaf44fc92109d6d2bdb767b5",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/n/ncurses-6.4-12.20240127.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/n/ncurses-6.4-12.20240127.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/n/ncurses-6.4-12.20240127.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/n/ncurses-6.4-12.20240127.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/n/ncurses-6.4-12.20240127.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "ncurses-base-0__6.4-12.20240127.fc40.aarch64",
        sha256 = "8a93376ce7423bd1a649a13f4b5105f270b4603f5cf3b3e230bdbda7f25dd788",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/n/ncurses-base-6.4-12.20240127.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/n/ncurses-base-6.4-12.20240127.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/n/ncurses-base-6.4-12.20240127.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/n/ncurses-base-6.4-12.20240127.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/n/ncurses-base-6.4-12.20240127.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "ncurses-libs-0__6.4-12.20240127.fc40.aarch64",
        sha256 = "5da1e9b629b555bae7040cd7c722b164d2baa728082813beb5f891b54ba97e93",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/n/ncurses-libs-6.4-12.20240127.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/n/ncurses-libs-6.4-12.20240127.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/n/ncurses-libs-6.4-12.20240127.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/n/ncurses-libs-6.4-12.20240127.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/n/ncurses-libs-6.4-12.20240127.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "nettle-0__3.9.1-6.fc40.aarch64",
        sha256 = "fcf96905af9447ab6933a503ca8f7b0416fcf43cf513126350ba02d18fb9326b",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/n/nettle-3.9.1-6.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/n/nettle-3.9.1-6.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/n/nettle-3.9.1-6.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/n/nettle-3.9.1-6.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/n/nettle-3.9.1-6.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "npth-0__1.7-1.fc40.aarch64",
        sha256 = "06fa439a2ff392d42aded3eae957f07c4e890bd68b8fd0131ee9942e12b121cd",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/n/npth-1.7-1.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/n/npth-1.7-1.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/n/npth-1.7-1.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/n/npth-1.7-1.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/n/npth-1.7-1.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "numactl-libs-0__2.0.16-5.fc40.aarch64",
        sha256 = "727b72bda291e8afc0866e6e945efa2ca0079020a72328218a85bbc01360511a",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/n/numactl-libs-2.0.16-5.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/n/numactl-libs-2.0.16-5.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/n/numactl-libs-2.0.16-5.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/n/numactl-libs-2.0.16-5.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/n/numactl-libs-2.0.16-5.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "ocaml-srpm-macros-0__9-3.fc40.aarch64",
        sha256 = "2d35dbd16fb7c9b306792eddea13d5c863a94ce1b9b9e0c8850cf3c571d56b48",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/o/ocaml-srpm-macros-9-3.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/o/ocaml-srpm-macros-9-3.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/o/ocaml-srpm-macros-9-3.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/o/ocaml-srpm-macros-9-3.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/o/ocaml-srpm-macros-9-3.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "openblas-srpm-macros-0__2-16.fc40.aarch64",
        sha256 = "46ee44ca72fab8e04a7d8c379a550466e7ded1c5a714d14764572fc78b1b5cc5",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/o/openblas-srpm-macros-2-16.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/o/openblas-srpm-macros-2-16.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/o/openblas-srpm-macros-2-16.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/o/openblas-srpm-macros-2-16.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/o/openblas-srpm-macros-2-16.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "openldap-0__2.6.8-1.fc40.aarch64",
        sha256 = "8f4f4bf77f350a991865ba517ed00b11fce46ebbb4944808408f59f3b2ec09fd",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/o/openldap-2.6.8-1.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/o/openldap-2.6.8-1.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/o/openldap-2.6.8-1.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/o/openldap-2.6.8-1.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/o/openldap-2.6.8-1.fc40.aarch64.rpm",
        ],
    )

    rpm(
        name = "openssl-devel-1__3.2.2-3.fc40.aarch64",
        sha256 = "a8c27c4f80d232f1d923bb2f68998ec4b5df5a64d95cbfe33a7acf98b32d4b75",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/o/openssl-devel-3.2.2-3.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/o/openssl-devel-3.2.2-3.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/o/openssl-devel-3.2.2-3.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/o/openssl-devel-3.2.2-3.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/o/openssl-devel-3.2.2-3.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "openssl-libs-1__3.2.2-3.fc40.aarch64",
        sha256 = "a4b77c775c73c1f65f9aa1abe3afd3266f3533f68d425285b557c090160caa03",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/o/openssl-libs-3.2.2-3.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/o/openssl-libs-3.2.2-3.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/o/openssl-libs-3.2.2-3.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/o/openssl-libs-3.2.2-3.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/o/openssl-libs-3.2.2-3.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "p11-kit-0__0.25.5-1.fc40.aarch64",
        sha256 = "605992b2e6fc58f681c090237f20bccb20ee860434a836f14c206c8f2d32cd90",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/p/p11-kit-0.25.5-1.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/p/p11-kit-0.25.5-1.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/p/p11-kit-0.25.5-1.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/p/p11-kit-0.25.5-1.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/p/p11-kit-0.25.5-1.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "p11-kit-trust-0__0.25.5-1.fc40.aarch64",
        sha256 = "ea51ca36df702eb1127593e72bf658777cd11ca07f49b660d3b9b9056ba364df",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/p/p11-kit-trust-0.25.5-1.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/p/p11-kit-trust-0.25.5-1.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/p/p11-kit-trust-0.25.5-1.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/p/p11-kit-trust-0.25.5-1.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/p/p11-kit-trust-0.25.5-1.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "package-notes-srpm-macros-0__0.5-11.fc40.aarch64",
        sha256 = "fb4d7c9f138a9ca7cc6fcb68b0820a99a4d67ee22041b64223430f70cee0240a",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/package-notes-srpm-macros-0.5-11.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/package-notes-srpm-macros-0.5-11.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/package-notes-srpm-macros-0.5-11.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/package-notes-srpm-macros-0.5-11.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/package-notes-srpm-macros-0.5-11.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "pam-0__1.6.1-5.fc40.aarch64",
        sha256 = "d2d472938a5a0cd53de30d2f4dfc9184c7632471d80fe1b79b55e5770451ea8c",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/p/pam-1.6.1-5.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/p/pam-1.6.1-5.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/p/pam-1.6.1-5.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/p/pam-1.6.1-5.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/p/pam-1.6.1-5.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "pam-libs-0__1.6.1-5.fc40.aarch64",
        sha256 = "cdf6fe8c955b8e24c6172ed45473175b16dddf8ba800e93a57e55cdd4f9951dd",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/p/pam-libs-1.6.1-5.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/p/pam-libs-1.6.1-5.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/p/pam-libs-1.6.1-5.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/p/pam-libs-1.6.1-5.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/p/pam-libs-1.6.1-5.fc40.aarch64.rpm",
        ],
    )

    rpm(
        name = "patch-0__2.7.6-24.fc40.aarch64",
        sha256 = "fcc316e1dd91c2729f6b90ba5ec6f1b22d0fd0c549b37411bfd0ea60056b9384",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/patch-2.7.6-24.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/patch-2.7.6-24.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/patch-2.7.6-24.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/patch-2.7.6-24.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/patch-2.7.6-24.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "pcre2-0__10.44-1.fc40.aarch64",
        sha256 = "c8e2d8d4794f0b534e95f91c56b90c701c034b80fdfbf98c0ac213ae83719083",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/p/pcre2-10.44-1.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/p/pcre2-10.44-1.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/p/pcre2-10.44-1.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/p/pcre2-10.44-1.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/p/pcre2-10.44-1.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "pcre2-syntax-0__10.44-1.fc40.aarch64",
        sha256 = "dbec699e88d42fc6fb1df0a8c0b9023941ed1b1b7625694253a612eaf9f2691d",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/p/pcre2-syntax-10.44-1.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/p/pcre2-syntax-10.44-1.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/p/pcre2-syntax-10.44-1.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/p/pcre2-syntax-10.44-1.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/p/pcre2-syntax-10.44-1.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-4__5.38.2-506.fc40.aarch64",
        sha256 = "9c2fae3ee668defd50add1e1c627813927b7a32e0bb0bdda72281d9bc73c64ca",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-5.38.2-506.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-5.38.2-506.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-5.38.2-506.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-5.38.2-506.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-5.38.2-506.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "perl-Algorithm-Diff-0__1.2010-11.fc40.aarch64",
        sha256 = "0c15f155ad3f9ca02482bf70b0d1fd640f2932a5964070106a4a90c62298491e",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Algorithm-Diff-1.2010-11.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Algorithm-Diff-1.2010-11.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Algorithm-Diff-1.2010-11.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Algorithm-Diff-1.2010-11.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Algorithm-Diff-1.2010-11.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Archive-Tar-0__3.02-6.fc40.aarch64",
        sha256 = "b0a57e6b4b9154afea01eb697884b6d30e354258c8ef954ce1a23e6d1603e0a9",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Archive-Tar-3.02-6.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Archive-Tar-3.02-6.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Archive-Tar-3.02-6.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Archive-Tar-3.02-6.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Archive-Tar-3.02-6.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Archive-Zip-0__1.68-14.fc40.aarch64",
        sha256 = "9df5357450fe34cff0c525e54ce7979e990d0da18460a09c65a404d23f3cb89a",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Archive-Zip-1.68-14.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Archive-Zip-1.68-14.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Archive-Zip-1.68-14.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Archive-Zip-1.68-14.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Archive-Zip-1.68-14.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Attribute-Handlers-0__1.03-506.fc40.aarch64",
        sha256 = "c750bbc0d76b38dce225fda305af3728713016af40aa0cc355c01dc984a5df22",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Attribute-Handlers-1.03-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Attribute-Handlers-1.03-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Attribute-Handlers-1.03-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Attribute-Handlers-1.03-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Attribute-Handlers-1.03-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-AutoLoader-0__5.74-506.fc40.aarch64",
        sha256 = "e801f69aa7745987f84f0ad8efa626bd3aea5fb29dc277ed6a5ab157de8878cb",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-AutoLoader-5.74-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-AutoLoader-5.74-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-AutoLoader-5.74-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-AutoLoader-5.74-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-AutoLoader-5.74-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-AutoSplit-0__5.74-506.fc40.aarch64",
        sha256 = "7b4762208435d31674648fddf6556db91ff41fa814f45174b215c0ff2049d1d6",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-AutoSplit-5.74-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-AutoSplit-5.74-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-AutoSplit-5.74-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-AutoSplit-5.74-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-AutoSplit-5.74-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-B-0__1.88-506.fc40.aarch64",
        sha256 = "508e23266262a248605f917237e2febd3f93fbd4f03d0a43bf4f3e0f99e77eec",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-B-1.88-506.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-B-1.88-506.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-B-1.88-506.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-B-1.88-506.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-B-1.88-506.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "perl-Benchmark-0__1.24-506.fc40.aarch64",
        sha256 = "68f04c9f6fcab675933a7f498efe2679a2b214d2f53fea80fb4908981b706329",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Benchmark-1.24-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Benchmark-1.24-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Benchmark-1.24-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Benchmark-1.24-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Benchmark-1.24-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-CPAN-0__2.36-503.fc40.aarch64",
        sha256 = "4b9740e2e7013a95a9962e0c287dd238e8df77336bcb62d32acccf01081aacbc",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-CPAN-2.36-503.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-CPAN-2.36-503.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-CPAN-2.36-503.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-CPAN-2.36-503.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-CPAN-2.36-503.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-CPAN-Meta-0__2.150010-502.fc40.aarch64",
        sha256 = "e8ee0fffaa79bda65bb25b0d51483692c44541982d432a8c25fd650bb8d8ade1",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-CPAN-Meta-2.150010-502.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-CPAN-Meta-2.150010-502.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-CPAN-Meta-2.150010-502.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-CPAN-Meta-2.150010-502.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-CPAN-Meta-2.150010-502.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-CPAN-Meta-Requirements-0__2.143-6.fc40.aarch64",
        sha256 = "5afb26f93a93f7ef39d06344f211688011ece7f15a063a951e7745559452b4ff",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-CPAN-Meta-Requirements-2.143-6.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-CPAN-Meta-Requirements-2.143-6.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-CPAN-Meta-Requirements-2.143-6.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-CPAN-Meta-Requirements-2.143-6.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-CPAN-Meta-Requirements-2.143-6.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-CPAN-Meta-YAML-0__0.018-503.fc40.aarch64",
        sha256 = "8f6613063103ec5d7c588a33a25b956fa340f28df6c5cd5eedd1f67c8f07cd44",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-CPAN-Meta-YAML-0.018-503.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-CPAN-Meta-YAML-0.018-503.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-CPAN-Meta-YAML-0.018-503.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-CPAN-Meta-YAML-0.018-503.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-CPAN-Meta-YAML-0.018-503.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Carp-0__1.54-502.fc40.aarch64",
        sha256 = "a65dd82703e0c5847733f52fcef81d82528381edbc84bf665a7bf53732e7b126",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Carp-1.54-502.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Carp-1.54-502.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Carp-1.54-502.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Carp-1.54-502.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Carp-1.54-502.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Class-Struct-0__0.68-506.fc40.aarch64",
        sha256 = "50c21b40deb69eb7e726f7d9c68e27af906e1eda028559d6e16364ded3625a16",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Class-Struct-0.68-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Class-Struct-0.68-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Class-Struct-0.68-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Class-Struct-0.68-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Class-Struct-0.68-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Compress-Bzip2-0__2.28-17.fc40.aarch64",
        sha256 = "23f000e4ce47129b8be5918c6ac62fad0aea368daf6a676ba60ad5774d6a385b",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Compress-Bzip2-2.28-17.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Compress-Bzip2-2.28-17.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Compress-Bzip2-2.28-17.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Compress-Bzip2-2.28-17.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Compress-Bzip2-2.28-17.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "perl-Compress-Raw-Bzip2-0__2.210-1.fc40.aarch64",
        sha256 = "d274b670cfc6ac01dbe89193add6a3ad7186572089c088fd7e303a748ae89df5",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Compress-Raw-Bzip2-2.210-1.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Compress-Raw-Bzip2-2.210-1.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Compress-Raw-Bzip2-2.210-1.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Compress-Raw-Bzip2-2.210-1.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Compress-Raw-Bzip2-2.210-1.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "perl-Compress-Raw-Lzma-0__2.209-8.fc40.aarch64",
        sha256 = "12ddf2df53640c302327193139cb6b14146b5313c41ece605391c360bbda576c",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Compress-Raw-Lzma-2.209-8.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Compress-Raw-Lzma-2.209-8.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Compress-Raw-Lzma-2.209-8.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Compress-Raw-Lzma-2.209-8.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Compress-Raw-Lzma-2.209-8.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "perl-Compress-Raw-Zlib-0__2.209-1.fc40.aarch64",
        sha256 = "b37c946aa22aff1720dd8c599d9d24eea0a327753fa48573b8d4b93daf794969",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Compress-Raw-Zlib-2.209-1.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Compress-Raw-Zlib-2.209-1.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Compress-Raw-Zlib-2.209-1.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Compress-Raw-Zlib-2.209-1.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Compress-Raw-Zlib-2.209-1.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "perl-Config-Extensions-0__0.03-506.fc40.aarch64",
        sha256 = "acddd0dedec20bc5a7e3e208006e8bd73f887d566d1eec9fdcb6a1061ba0c359",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Config-Extensions-0.03-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Config-Extensions-0.03-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Config-Extensions-0.03-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Config-Extensions-0.03-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Config-Extensions-0.03-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Config-Perl-V-0__0.36-503.fc40.aarch64",
        sha256 = "e619113c6ec1e04dd15968acb7438c6c89f8feb7311e7e5a244f538339dadced",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Config-Perl-V-0.36-503.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Config-Perl-V-0.36-503.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Config-Perl-V-0.36-503.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Config-Perl-V-0.36-503.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Config-Perl-V-0.36-503.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-DBM_Filter-0__0.06-506.fc40.aarch64",
        sha256 = "0a898e82d05169278e37bd626113ea11f5a516b712b0018c4dca82d3f09d3563",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-DBM_Filter-0.06-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-DBM_Filter-0.06-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-DBM_Filter-0.06-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-DBM_Filter-0.06-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-DBM_Filter-0.06-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-DB_File-0__1.859-3.fc40.aarch64",
        sha256 = "5c1440bfa1a466cf9703633d42ea1570c4c63b7845a205d59f7b0e6776c3290d",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-DB_File-1.859-3.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-DB_File-1.859-3.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-DB_File-1.859-3.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-DB_File-1.859-3.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-DB_File-1.859-3.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "perl-Data-Dumper-0__2.188-503.fc40.aarch64",
        sha256 = "cb1e1aa8752567e0cffaf95fa672cd4818bf95d698ba05d0a66c411e5e72e696",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Data-Dumper-2.188-503.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Data-Dumper-2.188-503.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Data-Dumper-2.188-503.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Data-Dumper-2.188-503.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Data-Dumper-2.188-503.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "perl-Data-OptList-0__0.114-4.fc40.aarch64",
        sha256 = "d284d509c99a24e7a4d60d03a9f31dc3be868f2fcc519849defb3351e480a260",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Data-OptList-0.114-4.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Data-OptList-0.114-4.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Data-OptList-0.114-4.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Data-OptList-0.114-4.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Data-OptList-0.114-4.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Data-Section-0__0.200008-5.fc40.aarch64",
        sha256 = "a54ccaa5da958d8988238e6f5c05196dc287a22ae3eba9eba41f72fa11bd46eb",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Data-Section-0.200008-5.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Data-Section-0.200008-5.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Data-Section-0.200008-5.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Data-Section-0.200008-5.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Data-Section-0.200008-5.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Devel-PPPort-0__3.71-503.fc40.aarch64",
        sha256 = "be1f4e6f73aa4bac73cb0c1f13012f96f4722c2ff0627390612108d2ac8b291b",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Devel-PPPort-3.71-503.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Devel-PPPort-3.71-503.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Devel-PPPort-3.71-503.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Devel-PPPort-3.71-503.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Devel-PPPort-3.71-503.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "perl-Devel-Peek-0__1.33-506.fc40.aarch64",
        sha256 = "fc01b4160423641410ec8e53d64cc527bb27e028ba5012c514c420be14079e53",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Devel-Peek-1.33-506.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Devel-Peek-1.33-506.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Devel-Peek-1.33-506.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Devel-Peek-1.33-506.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Devel-Peek-1.33-506.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "perl-Devel-SelfStubber-0__1.06-506.fc40.aarch64",
        sha256 = "085b644fbbf13a82f4848172149d6ffdb027c82ccd1c3b80a5b754bfda6d3c4d",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Devel-SelfStubber-1.06-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Devel-SelfStubber-1.06-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Devel-SelfStubber-1.06-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Devel-SelfStubber-1.06-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Devel-SelfStubber-1.06-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Devel-Size-0__0.84-1.fc40.aarch64",
        sha256 = "b031e243be21bbb9a4538676cfa6e5e1458c111199da00eddf1011886cae1a2a",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/p/perl-Devel-Size-0.84-1.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/p/perl-Devel-Size-0.84-1.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/p/perl-Devel-Size-0.84-1.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/p/perl-Devel-Size-0.84-1.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/p/perl-Devel-Size-0.84-1.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "perl-Digest-0__1.20-502.fc40.aarch64",
        sha256 = "7a3227717f0121273607249d64ea56953f0a1d68eb37e690e8c5f85851e2467a",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Digest-1.20-502.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Digest-1.20-502.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Digest-1.20-502.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Digest-1.20-502.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Digest-1.20-502.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Digest-MD5-0__2.59-3.fc40.aarch64",
        sha256 = "d951c7878435a0974868f1d119858b33ed3ab54798d403b59147e4d75dd75d23",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Digest-MD5-2.59-3.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Digest-MD5-2.59-3.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Digest-MD5-2.59-3.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Digest-MD5-2.59-3.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Digest-MD5-2.59-3.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "perl-Digest-SHA-1__6.04-503.fc40.aarch64",
        sha256 = "f10a051bf34a472195c87979bb0a56ac17cdb2a0efe0cd3dc1eb5ae6e56a16db",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Digest-SHA-6.04-503.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Digest-SHA-6.04-503.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Digest-SHA-6.04-503.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Digest-SHA-6.04-503.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Digest-SHA-6.04-503.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "perl-DirHandle-0__1.05-506.fc40.aarch64",
        sha256 = "6512c08bd7187fef2d8983e27cb832c0a02cf7fd53197c6615df4a46dce2ae45",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-DirHandle-1.05-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-DirHandle-1.05-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-DirHandle-1.05-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-DirHandle-1.05-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-DirHandle-1.05-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Dumpvalue-0__2.27-506.fc40.aarch64",
        sha256 = "bd8f2e9f28453f1723840d323004336fb7ad5e09b9514cfa33493732104b5b4e",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Dumpvalue-2.27-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Dumpvalue-2.27-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Dumpvalue-2.27-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Dumpvalue-2.27-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Dumpvalue-2.27-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-DynaLoader-0__1.54-506.fc40.aarch64",
        sha256 = "c3e45bbc89098327b747a75bd2f5f46fc0f2237b669cdc9bb2708bd1c0545d07",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-DynaLoader-1.54-506.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-DynaLoader-1.54-506.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-DynaLoader-1.54-506.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-DynaLoader-1.54-506.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-DynaLoader-1.54-506.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "perl-Encode-4__3.21-505.fc40.aarch64",
        sha256 = "f8e6dcc53f6a8081913f595514071985bd71ba666a83928cafde64f061487e71",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Encode-3.21-505.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Encode-3.21-505.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Encode-3.21-505.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Encode-3.21-505.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Encode-3.21-505.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "perl-Encode-devel-4__3.21-505.fc40.aarch64",
        sha256 = "46b9c41f4759c221d3adc1e5b15b39ed0bba142e097f573566b2c4674e9d6f76",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Encode-devel-3.21-505.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Encode-devel-3.21-505.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Encode-devel-3.21-505.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Encode-devel-3.21-505.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Encode-devel-3.21-505.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "perl-English-0__1.11-506.fc40.aarch64",
        sha256 = "0a1691e74e99e9d253ce17ba4b608da68be9831b05211d8f8d4ec7f4899642a9",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-English-1.11-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-English-1.11-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-English-1.11-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-English-1.11-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-English-1.11-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Env-0__1.06-502.fc40.aarch64",
        sha256 = "6698622f465a4f06b95351db7f8f2a11b785fad5bd414bb5cc8ccd6c0211c7ad",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Env-1.06-502.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Env-1.06-502.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Env-1.06-502.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Env-1.06-502.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Env-1.06-502.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Errno-0__1.37-506.fc40.aarch64",
        sha256 = "d9f2bcb693980a5203ec28984f8d5b134be50517e0a36722c4e0fc104d564340",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Errno-1.37-506.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Errno-1.37-506.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Errno-1.37-506.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Errno-1.37-506.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Errno-1.37-506.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "perl-Exporter-0__5.78-3.fc40.aarch64",
        sha256 = "8647c554687dbe5dbe010fbc826e897d4cc9a8c691c942e4633195645858ad10",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Exporter-5.78-3.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Exporter-5.78-3.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Exporter-5.78-3.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Exporter-5.78-3.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Exporter-5.78-3.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-ExtUtils-CBuilder-1__0.280238-502.fc40.aarch64",
        sha256 = "e0d83da248ed00c61a8b5097839882cb6bd5279cca04220a02a7a27d2ce93f3a",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-ExtUtils-CBuilder-0.280238-502.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-ExtUtils-CBuilder-0.280238-502.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-ExtUtils-CBuilder-0.280238-502.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-ExtUtils-CBuilder-0.280238-502.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-ExtUtils-CBuilder-0.280238-502.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-ExtUtils-Command-2__7.70-503.fc40.aarch64",
        sha256 = "b76c4c2222b98f38a12630bd6f7d2ea17b6fb39443091f6040baaa1d4d974cd5",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-ExtUtils-Command-7.70-503.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-ExtUtils-Command-7.70-503.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-ExtUtils-Command-7.70-503.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-ExtUtils-Command-7.70-503.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-ExtUtils-Command-7.70-503.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-ExtUtils-Constant-0__0.25-506.fc40.aarch64",
        sha256 = "59a6b0b8c0fd768d6e854adeb6e916162849711d82625d1250c7448dca91a2b1",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-ExtUtils-Constant-0.25-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-ExtUtils-Constant-0.25-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-ExtUtils-Constant-0.25-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-ExtUtils-Constant-0.25-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-ExtUtils-Constant-0.25-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-ExtUtils-Embed-0__1.35-506.fc40.aarch64",
        sha256 = "3b4ea4b7f3d36dfaf042cdf17eee592ffe10888ff654ee0e2d6b064cd5d7fe94",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-ExtUtils-Embed-1.35-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-ExtUtils-Embed-1.35-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-ExtUtils-Embed-1.35-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-ExtUtils-Embed-1.35-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-ExtUtils-Embed-1.35-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-ExtUtils-Install-0__2.22-502.fc40.aarch64",
        sha256 = "43c50ea47a2d6ce1de18c1cad2d753475165a397e8d4546341a8123e365eaceb",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-ExtUtils-Install-2.22-502.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-ExtUtils-Install-2.22-502.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-ExtUtils-Install-2.22-502.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-ExtUtils-Install-2.22-502.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-ExtUtils-Install-2.22-502.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-ExtUtils-MM-Utils-2__7.70-503.fc40.aarch64",
        sha256 = "f56cf2535b3c1a8a79344ac6abc6b2408f25cad47477b00879b14161a0867296",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-ExtUtils-MM-Utils-7.70-503.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-ExtUtils-MM-Utils-7.70-503.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-ExtUtils-MM-Utils-7.70-503.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-ExtUtils-MM-Utils-7.70-503.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-ExtUtils-MM-Utils-7.70-503.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-ExtUtils-MakeMaker-2__7.70-503.fc40.aarch64",
        sha256 = "08d3b88f88f4fb666f96eda69a279543e2fef2e3f3dfc4da4502e9d38dddf16d",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-ExtUtils-MakeMaker-7.70-503.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-ExtUtils-MakeMaker-7.70-503.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-ExtUtils-MakeMaker-7.70-503.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-ExtUtils-MakeMaker-7.70-503.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-ExtUtils-MakeMaker-7.70-503.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-ExtUtils-Manifest-1__1.75-5.fc40.aarch64",
        sha256 = "8c611f2a3d560bbb219a556bb4eb0c9a7fbf45d38f07ff228cbb7f0fe918e2df",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-ExtUtils-Manifest-1.75-5.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-ExtUtils-Manifest-1.75-5.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-ExtUtils-Manifest-1.75-5.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-ExtUtils-Manifest-1.75-5.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-ExtUtils-Manifest-1.75-5.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-ExtUtils-Miniperl-0__1.13-506.fc40.aarch64",
        sha256 = "536f08900d01b6ae7fb6fa1c0ed4243c875aa6475333a5ca4cffc3d24b4dbd03",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-ExtUtils-Miniperl-1.13-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-ExtUtils-Miniperl-1.13-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-ExtUtils-Miniperl-1.13-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-ExtUtils-Miniperl-1.13-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-ExtUtils-Miniperl-1.13-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-ExtUtils-ParseXS-1__3.51-503.fc40.aarch64",
        sha256 = "9bf5620bbd381fe0257b9da589cba7c1c919df199c7ec643af8e52da0bca7bd6",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-ExtUtils-ParseXS-3.51-503.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-ExtUtils-ParseXS-3.51-503.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-ExtUtils-ParseXS-3.51-503.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-ExtUtils-ParseXS-3.51-503.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-ExtUtils-ParseXS-3.51-503.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Fcntl-0__1.15-506.fc40.aarch64",
        sha256 = "4a83b3676b47b2b43dabdd137ee3fda157266e281801ee18872e90eeffeb9b70",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Fcntl-1.15-506.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Fcntl-1.15-506.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Fcntl-1.15-506.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Fcntl-1.15-506.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Fcntl-1.15-506.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "perl-File-Basename-0__2.86-506.fc40.aarch64",
        sha256 = "57164886c006d71b81f930735730ba1bbe56354558596cc582bdebb269d9f2d3",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-File-Basename-2.86-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-File-Basename-2.86-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-File-Basename-2.86-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-File-Basename-2.86-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-File-Basename-2.86-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-File-Compare-0__1.100.700-506.fc40.aarch64",
        sha256 = "17150da64c7c1cf7e03d63b47c4d45b5ed2c5c8a40a52f89c5dbad2d04380c5c",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-File-Compare-1.100.700-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-File-Compare-1.100.700-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-File-Compare-1.100.700-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-File-Compare-1.100.700-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-File-Compare-1.100.700-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-File-Copy-0__2.41-506.fc40.aarch64",
        sha256 = "9f074bce639cbb4e5d7e76466ff2106fb8dc3cd5adbfaadb415051570ab57bbf",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-File-Copy-2.41-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-File-Copy-2.41-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-File-Copy-2.41-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-File-Copy-2.41-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-File-Copy-2.41-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-File-DosGlob-0__1.12-506.fc40.aarch64",
        sha256 = "1ca2ceaa56d275321af4ac37dda20b95359ea6d056f604671c4eb92d7cb7c27c",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-File-DosGlob-1.12-506.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-File-DosGlob-1.12-506.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-File-DosGlob-1.12-506.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-File-DosGlob-1.12-506.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-File-DosGlob-1.12-506.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "perl-File-Fetch-0__1.04-502.fc40.aarch64",
        sha256 = "7a9f7ab914e85b91852bcd77bddf1cfd0532fbe24c17c080c87618d6a1f97691",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-File-Fetch-1.04-502.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-File-Fetch-1.04-502.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-File-Fetch-1.04-502.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-File-Fetch-1.04-502.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-File-Fetch-1.04-502.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-File-Find-0__1.43-506.fc40.aarch64",
        sha256 = "93c539bd75e3fa4a5656c9e341ab82dea1c43b306ef45ff26ae4c599633dbe14",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-File-Find-1.43-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-File-Find-1.43-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-File-Find-1.43-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-File-Find-1.43-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-File-Find-1.43-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-File-HomeDir-0__1.006-12.fc40.aarch64",
        sha256 = "d22d58d6fe4edee5a549c99ceb89c36f9022d6efff6217161ce30fd2eb34f7cf",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-File-HomeDir-1.006-12.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-File-HomeDir-1.006-12.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-File-HomeDir-1.006-12.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-File-HomeDir-1.006-12.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-File-HomeDir-1.006-12.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-File-Path-0__2.18-503.fc40.aarch64",
        sha256 = "8b152cd78a7f56136fd4d2f3b56111a8e5c4ab8192e50069a38df3eb90cdeba8",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-File-Path-2.18-503.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-File-Path-2.18-503.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-File-Path-2.18-503.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-File-Path-2.18-503.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-File-Path-2.18-503.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-File-Temp-1__0.231.100-503.fc40.aarch64",
        sha256 = "a686fe5e5e94ef9876b429099cd2bc85069a83148ea26b17970443a757822fa4",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-File-Temp-0.231.100-503.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-File-Temp-0.231.100-503.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-File-Temp-0.231.100-503.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-File-Temp-0.231.100-503.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-File-Temp-0.231.100-503.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-File-Which-0__1.27-11.fc40.aarch64",
        sha256 = "689e3c08798d1a4385435f2ee0e69c51509be2290de0e552a3f810b3d0790451",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-File-Which-1.27-11.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-File-Which-1.27-11.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-File-Which-1.27-11.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-File-Which-1.27-11.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-File-Which-1.27-11.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-File-stat-0__1.13-506.fc40.aarch64",
        sha256 = "0149f36e83814763c7937eddbecbc53a15959f3a69f2eaed6b380b07366698f5",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-File-stat-1.13-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-File-stat-1.13-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-File-stat-1.13-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-File-stat-1.13-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-File-stat-1.13-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-FileCache-0__1.10-506.fc40.aarch64",
        sha256 = "a9b0d66b720a3867a2c22504d95047785588fb7ee13b728606cc7845f27d47d0",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-FileCache-1.10-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-FileCache-1.10-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-FileCache-1.10-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-FileCache-1.10-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-FileCache-1.10-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-FileHandle-0__2.05-506.fc40.aarch64",
        sha256 = "8579aa3d0f1827c98678007e84a2ade496275f42884be0cc9f999c2de31a533b",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-FileHandle-2.05-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-FileHandle-2.05-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-FileHandle-2.05-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-FileHandle-2.05-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-FileHandle-2.05-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Filter-2__1.64-503.fc40.aarch64",
        sha256 = "39f04dcdf2c23d016075f870dc03ab4d97e39d367ec1648800d6ccee2049809e",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Filter-1.64-503.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Filter-1.64-503.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Filter-1.64-503.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Filter-1.64-503.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Filter-1.64-503.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "perl-Filter-Simple-0__0.96-503.fc40.aarch64",
        sha256 = "08b4bc22ed13283b595ebb153d9bc70e7732d30fb93df561d02a41e3e7136cb1",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Filter-Simple-0.96-503.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Filter-Simple-0.96-503.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Filter-Simple-0.96-503.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Filter-Simple-0.96-503.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Filter-Simple-0.96-503.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-FindBin-0__1.53-506.fc40.aarch64",
        sha256 = "ef6f7bcf631b34bc6092779fd835dbb46532389d08b0b072edb19c0afdd79dfa",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-FindBin-1.53-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-FindBin-1.53-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-FindBin-1.53-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-FindBin-1.53-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-FindBin-1.53-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-GDBM_File-1__1.24-506.fc40.aarch64",
        sha256 = "659bda08344b8fd33d487dd29d5dd687f86b86789593108d567e44a9f94b4a8e",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-GDBM_File-1.24-506.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-GDBM_File-1.24-506.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-GDBM_File-1.24-506.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-GDBM_File-1.24-506.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-GDBM_File-1.24-506.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "perl-Getopt-Long-1__2.57-4.fc40.aarch64",
        sha256 = "c61ce353b34a66009027a2d2e0d819a728b02e888a496c5cf8e63b164b731e6e",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/p/perl-Getopt-Long-2.57-4.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/p/perl-Getopt-Long-2.57-4.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/p/perl-Getopt-Long-2.57-4.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/p/perl-Getopt-Long-2.57-4.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/p/perl-Getopt-Long-2.57-4.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Getopt-Std-0__1.13-506.fc40.aarch64",
        sha256 = "4e46c286d79b208e6111ceae585d2d00b835c913686bd1cfd608f8e225e41a40",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Getopt-Std-1.13-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Getopt-Std-1.13-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Getopt-Std-1.13-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Getopt-Std-1.13-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Getopt-Std-1.13-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-HTTP-Tiny-0__0.088-5.fc40.aarch64",
        sha256 = "d0a5c3099349032e4527c11737c9f54ad7427685b563f10be9e6006b2acee36d",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-HTTP-Tiny-0.088-5.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-HTTP-Tiny-0.088-5.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-HTTP-Tiny-0.088-5.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-HTTP-Tiny-0.088-5.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-HTTP-Tiny-0.088-5.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Hash-Util-0__0.30-506.fc40.aarch64",
        sha256 = "e0e79f7dd42df7f4babc142ef45947b3876e4bc5937e1d54ff6fe71d0d5c8ca9",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Hash-Util-0.30-506.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Hash-Util-0.30-506.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Hash-Util-0.30-506.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Hash-Util-0.30-506.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Hash-Util-0.30-506.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "perl-Hash-Util-FieldHash-0__1.26-506.fc40.aarch64",
        sha256 = "fdd26d81e90ae52caea0f9ac4af9dcf0ea84815b2f98e418a42eeff10e8f1005",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Hash-Util-FieldHash-1.26-506.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Hash-Util-FieldHash-1.26-506.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Hash-Util-FieldHash-1.26-506.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Hash-Util-FieldHash-1.26-506.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Hash-Util-FieldHash-1.26-506.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "perl-I18N-Collate-0__1.02-506.fc40.aarch64",
        sha256 = "d4336be43ce67e258e9743c3184bcb015d985cfc72ee0004ae690451a6f8b6f1",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-I18N-Collate-1.02-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-I18N-Collate-1.02-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-I18N-Collate-1.02-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-I18N-Collate-1.02-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-I18N-Collate-1.02-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-I18N-LangTags-0__0.45-506.fc40.aarch64",
        sha256 = "a5780829fac4152291d60cb5eb06a6a7a1d068b061116b024c4ed73faa4f4e56",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-I18N-LangTags-0.45-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-I18N-LangTags-0.45-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-I18N-LangTags-0.45-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-I18N-LangTags-0.45-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-I18N-LangTags-0.45-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-I18N-Langinfo-0__0.22-506.fc40.aarch64",
        sha256 = "1fa292ecdff658fcc46453381dcb139ce170756bdd2ba73beda629ca0c65b0ae",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-I18N-Langinfo-0.22-506.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-I18N-Langinfo-0.22-506.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-I18N-Langinfo-0.22-506.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-I18N-Langinfo-0.22-506.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-I18N-Langinfo-0.22-506.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "perl-IO-0__1.52-506.fc40.aarch64",
        sha256 = "16baf2e784761040477786545e7633bb2ebdb39fd791b4392032df7de3c45e6a",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-IO-1.52-506.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-IO-1.52-506.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-IO-1.52-506.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-IO-1.52-506.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-IO-1.52-506.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "perl-IO-Compress-0__2.207-1.fc40.aarch64",
        sha256 = "e05036e4a95f2d57814546bad0b2031884fdac6ed88049c03a5e70130626e682",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-IO-Compress-2.207-1.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-IO-Compress-2.207-1.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-IO-Compress-2.207-1.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-IO-Compress-2.207-1.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-IO-Compress-2.207-1.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-IO-Compress-Lzma-0__2.207-1.fc40.aarch64",
        sha256 = "bd4cd6e8050d03ef72f8bc51a653572e0c0ec16b6279b48664d6ad4f729d7608",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-IO-Compress-Lzma-2.207-1.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-IO-Compress-Lzma-2.207-1.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-IO-Compress-Lzma-2.207-1.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-IO-Compress-Lzma-2.207-1.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-IO-Compress-Lzma-2.207-1.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-IO-Socket-IP-0__0.42-2.fc40.aarch64",
        sha256 = "30aa2fa573d6772840ec30431d3e92c78d90442e5349dc9bab14f70816e84ecb",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-IO-Socket-IP-0.42-2.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-IO-Socket-IP-0.42-2.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-IO-Socket-IP-0.42-2.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-IO-Socket-IP-0.42-2.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-IO-Socket-IP-0.42-2.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-IO-Socket-SSL-0__2.085-1.fc40.aarch64",
        sha256 = "660515f32e0985d3ad5d5b58426d77ed07eca255cd30954469dffc6de8516d0e",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-IO-Socket-SSL-2.085-1.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-IO-Socket-SSL-2.085-1.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-IO-Socket-SSL-2.085-1.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-IO-Socket-SSL-2.085-1.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-IO-Socket-SSL-2.085-1.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-IO-Zlib-1__1.15-1.fc40.aarch64",
        sha256 = "282d7aef6b5cad631a03a5a0a28f7302d9ad52a4922b29783ed2d99d5ca0a1fe",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-IO-Zlib-1.15-1.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-IO-Zlib-1.15-1.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-IO-Zlib-1.15-1.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-IO-Zlib-1.15-1.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-IO-Zlib-1.15-1.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-IPC-Cmd-2__1.04-504.fc40.aarch64",
        sha256 = "3a24352aaab55dea0deb0397bc3fd5edce2eee35f34e2e1eaefbd8f026d4d032",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-IPC-Cmd-1.04-504.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-IPC-Cmd-1.04-504.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-IPC-Cmd-1.04-504.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-IPC-Cmd-1.04-504.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-IPC-Cmd-1.04-504.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-IPC-Open3-0__1.22-506.fc40.aarch64",
        sha256 = "c50d0a81b90d11648ea010d72fcb924cd910d9b9c247021c152e1f7ee5ee4e46",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-IPC-Open3-1.22-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-IPC-Open3-1.22-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-IPC-Open3-1.22-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-IPC-Open3-1.22-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-IPC-Open3-1.22-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-IPC-SysV-0__2.09-505.fc40.aarch64",
        sha256 = "6004863277bec95a31fca5066cc5c483f8332984dda35689e6912dbf1600be92",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-IPC-SysV-2.09-505.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-IPC-SysV-2.09-505.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-IPC-SysV-2.09-505.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-IPC-SysV-2.09-505.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-IPC-SysV-2.09-505.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "perl-IPC-System-Simple-0__1.30-13.fc40.aarch64",
        sha256 = "02af6f37e13d21d516a6e152ed6cee163c305975aba24aef4da38d5a1846ecd1",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-IPC-System-Simple-1.30-13.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-IPC-System-Simple-1.30-13.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-IPC-System-Simple-1.30-13.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-IPC-System-Simple-1.30-13.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-IPC-System-Simple-1.30-13.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-JSON-PP-1__4.16-503.fc40.aarch64",
        sha256 = "e3072a5d7b5325c3ded189bb78582231a45ff9dc70f6f27a42ef9a3388dddceb",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-JSON-PP-4.16-503.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-JSON-PP-4.16-503.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-JSON-PP-4.16-503.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-JSON-PP-4.16-503.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-JSON-PP-4.16-503.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Locale-Maketext-0__1.33-503.fc40.aarch64",
        sha256 = "8cc64314cb0124b97c629b0ace1d1f8fc37bf3aa15a654fb1115cf0fb1713386",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Locale-Maketext-1.33-503.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Locale-Maketext-1.33-503.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Locale-Maketext-1.33-503.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Locale-Maketext-1.33-503.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Locale-Maketext-1.33-503.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Locale-Maketext-Simple-1__0.21-506.fc40.aarch64",
        sha256 = "50007a4c207fb30ab09013406d98d2510433786d1deed1bb56bf1abf84e8fbec",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Locale-Maketext-Simple-0.21-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Locale-Maketext-Simple-0.21-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Locale-Maketext-Simple-0.21-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Locale-Maketext-Simple-0.21-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Locale-Maketext-Simple-0.21-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-MIME-Base64-0__3.16-503.fc40.aarch64",
        sha256 = "f304200300976e3be2af9be530ea1bbbc1e330044ca274fe011cc0a964986271",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-MIME-Base64-3.16-503.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-MIME-Base64-3.16-503.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-MIME-Base64-3.16-503.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-MIME-Base64-3.16-503.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-MIME-Base64-3.16-503.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "perl-MRO-Compat-0__0.15-9.fc40.aarch64",
        sha256 = "c714ded9c6fc9a4bc5cee122808d5501c0fe2d443fdfe225705efb2668a61e01",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-MRO-Compat-0.15-9.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-MRO-Compat-0.15-9.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-MRO-Compat-0.15-9.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-MRO-Compat-0.15-9.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-MRO-Compat-0.15-9.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Math-BigInt-1__2.0030.03-1.fc40.aarch64",
        sha256 = "8e31a5a14c24675aef8417af535128ec6096c8f204a1983a773a9e74573aeee3",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/p/perl-Math-BigInt-2.0030.03-1.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/p/perl-Math-BigInt-2.0030.03-1.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/p/perl-Math-BigInt-2.0030.03-1.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/p/perl-Math-BigInt-2.0030.03-1.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/p/perl-Math-BigInt-2.0030.03-1.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Math-BigInt-FastCalc-0__0.501.800-3.fc40.aarch64",
        sha256 = "ba5fcc64d3e6b28179563957c015ac81b4c53e1c065ac686ebe8f39430c70e53",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Math-BigInt-FastCalc-0.501.800-3.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Math-BigInt-FastCalc-0.501.800-3.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Math-BigInt-FastCalc-0.501.800-3.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Math-BigInt-FastCalc-0.501.800-3.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Math-BigInt-FastCalc-0.501.800-3.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "perl-Math-Complex-0__1.62-506.fc40.aarch64",
        sha256 = "76fff240e889f8fceae4e362cfd6c65eb9ffd66a4eb39ffbc4e7923eb8e9feed",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Math-Complex-1.62-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Math-Complex-1.62-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Math-Complex-1.62-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Math-Complex-1.62-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Math-Complex-1.62-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Memoize-0__1.16-506.fc40.aarch64",
        sha256 = "5c0be04f2435f3de7363841a93972de8a550102e812a0a6ea8ea352373d6a8f7",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Memoize-1.16-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Memoize-1.16-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Memoize-1.16-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Memoize-1.16-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Memoize-1.16-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Module-Build-2__0.42.34-5.fc40.aarch64",
        sha256 = "672439cd1b937b4ff1687138d83026d2b491aec1d20fcd0a8ca97dffca005024",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Module-Build-0.42.34-5.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Module-Build-0.42.34-5.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Module-Build-0.42.34-5.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Module-Build-0.42.34-5.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Module-Build-0.42.34-5.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Module-CoreList-1__5.20241120-1.fc40.aarch64",
        sha256 = "a608917869c406c2fdb07d28d926fb09d9dfea83beb0ed052e427cfe4675c95a",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/p/perl-Module-CoreList-5.20241120-1.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/p/perl-Module-CoreList-5.20241120-1.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/p/perl-Module-CoreList-5.20241120-1.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/p/perl-Module-CoreList-5.20241120-1.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/p/perl-Module-CoreList-5.20241120-1.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Module-CoreList-tools-1__5.20241120-1.fc40.aarch64",
        sha256 = "d1b25d44bcc97181ee94b23c12df38a69f433f86892a46eb53dcf174c1b7bef5",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/p/perl-Module-CoreList-tools-5.20241120-1.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/p/perl-Module-CoreList-tools-5.20241120-1.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/p/perl-Module-CoreList-tools-5.20241120-1.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/p/perl-Module-CoreList-tools-5.20241120-1.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/p/perl-Module-CoreList-tools-5.20241120-1.fc40.noarch.rpm",
        ],
    )

    rpm(
        name = "perl-Module-Load-1__0.36-503.fc40.aarch64",
        sha256 = "ffff4d9fa6f9685b36aca24a39f965d4cd94ccb13a4c73e4fec45460733893ef",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Module-Load-0.36-503.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Module-Load-0.36-503.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Module-Load-0.36-503.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Module-Load-0.36-503.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Module-Load-0.36-503.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Module-Load-Conditional-0__0.74-503.fc40.aarch64",
        sha256 = "77bab9d62249a280b9f14e1ae6ed4071dc267234a576cb7fb012533b4ccf116b",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Module-Load-Conditional-0.74-503.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Module-Load-Conditional-0.74-503.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Module-Load-Conditional-0.74-503.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Module-Load-Conditional-0.74-503.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Module-Load-Conditional-0.74-503.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Module-Loaded-1__0.08-506.fc40.aarch64",
        sha256 = "c431dde1f7a5e0af118b300a5e94a965f2376e4fafeda1bac644f225678c6314",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Module-Loaded-0.08-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Module-Loaded-0.08-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Module-Loaded-0.08-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Module-Loaded-0.08-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Module-Loaded-0.08-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Module-Metadata-0__1.000038-5.fc40.aarch64",
        sha256 = "dbefbe5bfd576e1556b257f71deb8dbe83aef33aa639791e1401f74f1de96aca",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Module-Metadata-1.000038-5.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Module-Metadata-1.000038-5.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Module-Metadata-1.000038-5.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Module-Metadata-1.000038-5.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Module-Metadata-1.000038-5.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Module-Signature-0__0.88-9.fc40.aarch64",
        sha256 = "1d89dff0b55c5fdf5aa9abd61552858c7a975a0c51cc5cf25879b12e6fc8f2ef",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Module-Signature-0.88-9.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Module-Signature-0.88-9.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Module-Signature-0.88-9.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Module-Signature-0.88-9.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Module-Signature-0.88-9.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Mozilla-CA-0__20231213-3.fc40.aarch64",
        sha256 = "a53e9503a79437c7585c8f82d2047ad0eb5f53b3a92edfbd218f620d7dd47c98",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Mozilla-CA-20231213-3.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Mozilla-CA-20231213-3.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Mozilla-CA-20231213-3.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Mozilla-CA-20231213-3.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Mozilla-CA-20231213-3.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-NDBM_File-0__1.16-506.fc40.aarch64",
        sha256 = "a8e9d35ab581e218fc4b1f872684e89f20d6c8fad081c4ec85a33343fcfb25a7",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-NDBM_File-1.16-506.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-NDBM_File-1.16-506.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-NDBM_File-1.16-506.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-NDBM_File-1.16-506.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-NDBM_File-1.16-506.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "perl-NEXT-0__0.69-506.fc40.aarch64",
        sha256 = "d860065135bd49e7b3bddb61dd058482c34d1eb618aeceef80a88059195c029d",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-NEXT-0.69-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-NEXT-0.69-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-NEXT-0.69-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-NEXT-0.69-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-NEXT-0.69-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Net-0__1.03-506.fc40.aarch64",
        sha256 = "2095675d99107bc61c726864cd396fb9806c93425afe17017470305c2c227896",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Net-1.03-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Net-1.03-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Net-1.03-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Net-1.03-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Net-1.03-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Net-Ping-0__2.76-502.fc40.aarch64",
        sha256 = "4fcb1cc76c8b4fbe58eb2dc82800fb06b7797582d692c4743cb69f5d0b579421",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Net-Ping-2.76-502.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Net-Ping-2.76-502.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Net-Ping-2.76-502.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Net-Ping-2.76-502.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Net-Ping-2.76-502.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Net-SSLeay-0__1.94-3.fc40.aarch64",
        sha256 = "dc9468d3c16aa57ef6ebd394205cb8009fd911c6bdd9cbb8d523046e06c6d4ff",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Net-SSLeay-1.94-3.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Net-SSLeay-1.94-3.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Net-SSLeay-1.94-3.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Net-SSLeay-1.94-3.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Net-SSLeay-1.94-3.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "perl-ODBM_File-0__1.18-506.fc40.aarch64",
        sha256 = "626cca8971b8a5ba513d039312bc05978c809c87352f3a3499a0d5130b924d04",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-ODBM_File-1.18-506.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-ODBM_File-1.18-506.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-ODBM_File-1.18-506.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-ODBM_File-1.18-506.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-ODBM_File-1.18-506.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "perl-Object-HashBase-0__0.013-1.fc40.aarch64",
        sha256 = "9d44cb7f7a16793be7679c9ec543faae4c893a4887276adf66ef0ad5f0578bbc",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Object-HashBase-0.013-1.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Object-HashBase-0.013-1.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Object-HashBase-0.013-1.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Object-HashBase-0.013-1.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Object-HashBase-0.013-1.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Opcode-0__1.64-506.fc40.aarch64",
        sha256 = "277605909b3df9afcd1d2aa7fb362e679ff493205b768552eb0fa37083aaec60",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Opcode-1.64-506.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Opcode-1.64-506.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Opcode-1.64-506.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Opcode-1.64-506.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Opcode-1.64-506.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "perl-POSIX-0__2.13-506.fc40.aarch64",
        sha256 = "933a79fb8a6bd43194fb8f830e51245cd7f5bc7567180c5cf74f3cf6729afe96",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-POSIX-2.13-506.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-POSIX-2.13-506.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-POSIX-2.13-506.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-POSIX-2.13-506.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-POSIX-2.13-506.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "perl-Package-Generator-0__1.106-31.fc40.aarch64",
        sha256 = "bc41d9ad5b7c28ddfaf66a169d9ad5f4452cc7c360b8f8e8d61f659156b458d4",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Package-Generator-1.106-31.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Package-Generator-1.106-31.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Package-Generator-1.106-31.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Package-Generator-1.106-31.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Package-Generator-1.106-31.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Params-Check-1__0.38-502.fc40.aarch64",
        sha256 = "f1fdf697d7276ce45999bde7ad5a54e52c55b48d633965fe33c66994b8e2ebd3",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Params-Check-0.38-502.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Params-Check-0.38-502.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Params-Check-0.38-502.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Params-Check-0.38-502.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Params-Check-0.38-502.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Params-Util-0__1.102-14.fc40.aarch64",
        sha256 = "60ab3ec57af8cf794aea870f1195e98abf66cd3630ade6e26789fffe86067fc4",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Params-Util-1.102-14.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Params-Util-1.102-14.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Params-Util-1.102-14.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Params-Util-1.102-14.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Params-Util-1.102-14.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "perl-PathTools-0__3.89-502.fc40.aarch64",
        sha256 = "5078544f191517890515bb823127a78c354388adca593803c5fa28c5bb267f7a",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-PathTools-3.89-502.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-PathTools-3.89-502.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-PathTools-3.89-502.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-PathTools-3.89-502.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-PathTools-3.89-502.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "perl-Perl-OSType-0__1.010-503.fc40.aarch64",
        sha256 = "a9b7f734bad66bfe7c82611da67947558bee5d5d400c44f6485211952b74f6fd",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Perl-OSType-1.010-503.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Perl-OSType-1.010-503.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Perl-OSType-1.010-503.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Perl-OSType-1.010-503.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Perl-OSType-1.010-503.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-PerlIO-via-QuotedPrint-0__0.10-502.fc40.aarch64",
        sha256 = "c2eac6a5c9b42fde5593f056fd6c6a952723decae4c24b311964b1272238829a",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-PerlIO-via-QuotedPrint-0.10-502.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-PerlIO-via-QuotedPrint-0.10-502.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-PerlIO-via-QuotedPrint-0.10-502.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-PerlIO-via-QuotedPrint-0.10-502.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-PerlIO-via-QuotedPrint-0.10-502.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Pod-Checker-4__1.77-1.fc40.aarch64",
        sha256 = "01491dd52a63826b44360824a64934c5dfdf03715049c9fd4ff06f60ce00db6f",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Pod-Checker-1.77-1.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Pod-Checker-1.77-1.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Pod-Checker-1.77-1.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Pod-Checker-1.77-1.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Pod-Checker-1.77-1.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Pod-Escapes-1__1.07-503.fc40.aarch64",
        sha256 = "1f96f3aff486ab917ba871eb4875e4f201ae5731b9ec2bd974f11257861fad5c",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Pod-Escapes-1.07-503.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Pod-Escapes-1.07-503.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Pod-Escapes-1.07-503.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Pod-Escapes-1.07-503.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Pod-Escapes-1.07-503.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Pod-Functions-0__1.14-506.fc40.aarch64",
        sha256 = "c9b76089229788021e5a8adda8b5d08a3e463c84436263ae164d64f1f89e0d42",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Pod-Functions-1.14-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Pod-Functions-1.14-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Pod-Functions-1.14-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Pod-Functions-1.14-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Pod-Functions-1.14-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Pod-Html-0__1.34-506.fc40.aarch64",
        sha256 = "f26494f733d23e9658e046bd6025198072041a46f47935c69f09046af9b7ad4b",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Pod-Html-1.34-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Pod-Html-1.34-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Pod-Html-1.34-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Pod-Html-1.34-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Pod-Html-1.34-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Pod-Perldoc-0__3.28.01-503.fc40.aarch64",
        sha256 = "6df25726589106437cf557eae67f3b993e5ff40dad2eeb9150dfc2b23115e6e1",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Pod-Perldoc-3.28.01-503.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Pod-Perldoc-3.28.01-503.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Pod-Perldoc-3.28.01-503.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Pod-Perldoc-3.28.01-503.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Pod-Perldoc-3.28.01-503.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Pod-Simple-1__3.45-6.fc40.aarch64",
        sha256 = "412128c45e763ea21250bae59964120a11f5e29e55e18b3d1f93ab64ab160f6b",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Pod-Simple-3.45-6.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Pod-Simple-3.45-6.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Pod-Simple-3.45-6.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Pod-Simple-3.45-6.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Pod-Simple-3.45-6.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Pod-Usage-4__2.03-504.fc40.aarch64",
        sha256 = "47df4820644fab7febd411bcb3b5dbffdefdcc270951a288500b3fca0b7bb5bf",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/p/perl-Pod-Usage-2.03-504.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/p/perl-Pod-Usage-2.03-504.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/p/perl-Pod-Usage-2.03-504.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/p/perl-Pod-Usage-2.03-504.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/p/perl-Pod-Usage-2.03-504.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Safe-0__2.44-506.fc40.aarch64",
        sha256 = "28836afae302d1b66131bb731187f2a8f2b907e09fa5104b478a9b74911859ce",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Safe-2.44-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Safe-2.44-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Safe-2.44-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Safe-2.44-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Safe-2.44-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Scalar-List-Utils-5__1.63-503.fc40.aarch64",
        sha256 = "d44ef3a9038d836097f7d5eb8c33ee10cff57ce1a5059cd6e282105dac8d170c",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Scalar-List-Utils-1.63-503.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Scalar-List-Utils-1.63-503.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Scalar-List-Utils-1.63-503.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Scalar-List-Utils-1.63-503.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Scalar-List-Utils-1.63-503.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "perl-Search-Dict-0__1.07-506.fc40.aarch64",
        sha256 = "ef996e4c126174c718efce0f3ffc1b043743c6f615627feae7b23100a743720b",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Search-Dict-1.07-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Search-Dict-1.07-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Search-Dict-1.07-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Search-Dict-1.07-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Search-Dict-1.07-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-SelectSaver-0__1.02-506.fc40.aarch64",
        sha256 = "6a45769f7520bc5ef502bdbc2287e1998cc9ca7cd195c845a7c552a2a9b3a650",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-SelectSaver-1.02-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-SelectSaver-1.02-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-SelectSaver-1.02-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-SelectSaver-1.02-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-SelectSaver-1.02-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-SelfLoader-0__1.26-506.fc40.aarch64",
        sha256 = "52e9a5cd271df85c17be959eca0e110ba469a40c3f72b22901bc5345c3c7dc39",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-SelfLoader-1.26-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-SelfLoader-1.26-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-SelfLoader-1.26-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-SelfLoader-1.26-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-SelfLoader-1.26-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Socket-4__2.038-1.fc40.aarch64",
        sha256 = "5c6faeab6f8e924cd2eaef475a938d51be55889bc7d450e4f45e772c8eeeaf3e",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/p/perl-Socket-2.038-1.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/p/perl-Socket-2.038-1.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/p/perl-Socket-2.038-1.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/p/perl-Socket-2.038-1.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/p/perl-Socket-2.038-1.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "perl-Software-License-0__0.104006-1.fc40.aarch64",
        sha256 = "0ce7c97d8327e61877d8639662c03c5223e73209b3fe0ca0ef769085232cd6c8",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Software-License-0.104006-1.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Software-License-0.104006-1.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Software-License-0.104006-1.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Software-License-0.104006-1.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Software-License-0.104006-1.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Storable-1__3.32-502.fc40.aarch64",
        sha256 = "1541f109d37016fc2c12df202cff63e1f8642ace376c7d9abc61f0a01ea22e37",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Storable-3.32-502.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Storable-3.32-502.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Storable-3.32-502.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Storable-3.32-502.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Storable-3.32-502.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "perl-Sub-Exporter-0__0.991-3.fc40.aarch64",
        sha256 = "7f0a77a8c39db8498e070f5094d1e232b5c49700cbbdd2fec8bfaaf9c82d0a2e",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Sub-Exporter-0.991-3.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Sub-Exporter-0.991-3.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Sub-Exporter-0.991-3.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Sub-Exporter-0.991-3.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Sub-Exporter-0.991-3.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Sub-Install-0__0.929-5.fc40.aarch64",
        sha256 = "84626d852eda28dc8d15f99e7057f08fccd4bedb18a11913567f07eb6effd0af",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Sub-Install-0.929-5.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Sub-Install-0.929-5.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Sub-Install-0.929-5.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Sub-Install-0.929-5.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Sub-Install-0.929-5.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Symbol-0__1.09-506.fc40.aarch64",
        sha256 = "31541521adffee5c73c30bdafdda4f380cd9c08421336f30c3036c71160674da",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Symbol-1.09-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Symbol-1.09-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Symbol-1.09-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Symbol-1.09-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Symbol-1.09-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Sys-Hostname-0__1.25-506.fc40.aarch64",
        sha256 = "de63b42b7adf7da47ff4a0a8c61c1951be5320b8ee4cde4133d0190e17eb1294",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Sys-Hostname-1.25-506.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Sys-Hostname-1.25-506.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Sys-Hostname-1.25-506.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Sys-Hostname-1.25-506.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Sys-Hostname-1.25-506.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "perl-Sys-Syslog-0__0.36-504.fc40.aarch64",
        sha256 = "956e44fcecc352a44e2fd8f95038418027051c9e36ff7d1436c4e449e3f5cdb5",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Sys-Syslog-0.36-504.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Sys-Syslog-0.36-504.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Sys-Syslog-0.36-504.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Sys-Syslog-0.36-504.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Sys-Syslog-0.36-504.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "perl-Term-ANSIColor-0__5.01-504.fc40.aarch64",
        sha256 = "79dcb1cb584cdb7bbb7c022b63f4b48e14d59c510825c258caca4b126da2ce53",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Term-ANSIColor-5.01-504.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Term-ANSIColor-5.01-504.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Term-ANSIColor-5.01-504.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Term-ANSIColor-5.01-504.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Term-ANSIColor-5.01-504.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Term-Cap-0__1.18-503.fc40.aarch64",
        sha256 = "b6e98d6c3bc2b72a1b5218d6000cbb5ce9787fbdf852cabc742c3e6e3d1f015b",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Term-Cap-1.18-503.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Term-Cap-1.18-503.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Term-Cap-1.18-503.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Term-Cap-1.18-503.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Term-Cap-1.18-503.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Term-Complete-0__1.403-506.fc40.aarch64",
        sha256 = "72d50b52ce34aa7a28ec6f21dcb5138215b4719840ee7d3bc8dd25f4fb17f4e6",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Term-Complete-1.403-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Term-Complete-1.403-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Term-Complete-1.403-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Term-Complete-1.403-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Term-Complete-1.403-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Term-ReadLine-0__1.17-506.fc40.aarch64",
        sha256 = "8f9fcee1df65ec707a56a5db1c6487644498aca37889a758b664cdd8401d29a3",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Term-ReadLine-1.17-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Term-ReadLine-1.17-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Term-ReadLine-1.17-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Term-ReadLine-1.17-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Term-ReadLine-1.17-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Term-Table-0__0.018-3.fc40.aarch64",
        sha256 = "80ff06ef84ca3cfcc99738c5388726b478de9efbcd2d08fc5bf2c015c3c793a9",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Term-Table-0.018-3.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Term-Table-0.018-3.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Term-Table-0.018-3.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Term-Table-0.018-3.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Term-Table-0.018-3.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Test-0__1.31-506.fc40.aarch64",
        sha256 = "87f8f0b1a7dc09546cc114a793f75ff1e7d338c10b73980bf227852c56a98cf6",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Test-1.31-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Test-1.31-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Test-1.31-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Test-1.31-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Test-1.31-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Test-Harness-1__3.50-1.fc40.aarch64",
        sha256 = "67fbd3f9fc39d8eb7380d8ea64ef53bd1ca0ddea82a0d06df2eb841f076af330",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/p/perl-Test-Harness-3.50-1.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/p/perl-Test-Harness-3.50-1.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/p/perl-Test-Harness-3.50-1.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/p/perl-Test-Harness-3.50-1.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/p/perl-Test-Harness-3.50-1.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Test-Simple-3__1.302198-3.fc40.aarch64",
        sha256 = "50a1496c9d73779b0fe9ac4042f9d0682939245e20da5e1f8443f5dc76381422",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Test-Simple-1.302198-3.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Test-Simple-1.302198-3.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Test-Simple-1.302198-3.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Test-Simple-1.302198-3.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Test-Simple-1.302198-3.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Text-Abbrev-0__1.02-506.fc40.aarch64",
        sha256 = "49be1af475a856ae7d553328d3545f8b5185c94d8105c6388daaef8f78aa507a",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Text-Abbrev-1.02-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Text-Abbrev-1.02-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Text-Abbrev-1.02-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Text-Abbrev-1.02-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Text-Abbrev-1.02-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Text-Balanced-0__2.06-502.fc40.aarch64",
        sha256 = "7bc6da9b4f02264ba87ee83450af7d42c37cda260fb3f7e6be21bb622534bebc",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Text-Balanced-2.06-502.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Text-Balanced-2.06-502.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Text-Balanced-2.06-502.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Text-Balanced-2.06-502.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Text-Balanced-2.06-502.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Text-Diff-0__1.45-21.fc40.aarch64",
        sha256 = "c177c7c5d468a2d656cab043e8d78459fa6aa1a7a7bc64a89b8d31f4724b535e",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Text-Diff-1.45-21.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Text-Diff-1.45-21.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Text-Diff-1.45-21.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Text-Diff-1.45-21.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Text-Diff-1.45-21.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Text-Glob-0__0.11-23.fc40.aarch64",
        sha256 = "131fa277189c06632c3e240fb0a5fef784026310913859db7d25602463fd912c",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Text-Glob-0.11-23.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Text-Glob-0.11-23.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Text-Glob-0.11-23.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Text-Glob-0.11-23.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Text-Glob-0.11-23.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Text-ParseWords-0__3.31-502.fc40.aarch64",
        sha256 = "6911b5d1d519ba25c008d9da7631e7d4e60e7902bbe3bafb35924c440a47080f",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Text-ParseWords-3.31-502.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Text-ParseWords-3.31-502.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Text-ParseWords-3.31-502.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Text-ParseWords-3.31-502.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Text-ParseWords-3.31-502.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Text-Tabs__plus__Wrap-0__2024.001-1.fc40.aarch64",
        sha256 = "9ca255e239f747f40094f3ba0c81079f009a24a244706128f6bfc077f5d3e97d",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Text-Tabs+Wrap-2024.001-1.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Text-Tabs+Wrap-2024.001-1.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Text-Tabs+Wrap-2024.001-1.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Text-Tabs+Wrap-2024.001-1.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Text-Tabs+Wrap-2024.001-1.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Text-Template-0__1.61-5.fc40.aarch64",
        sha256 = "08fc39fb7a3bf50f473e7a4648db696e16d598b313d5df0c1590bc2849e815d3",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Text-Template-1.61-5.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Text-Template-1.61-5.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Text-Template-1.61-5.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Text-Template-1.61-5.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Text-Template-1.61-5.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Thread-0__3.05-506.fc40.aarch64",
        sha256 = "de7779872fcc65fb45930a98adaa732ea0f9ed914c834cd8a0f982a988b8ca78",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Thread-3.05-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Thread-3.05-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Thread-3.05-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Thread-3.05-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Thread-3.05-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Thread-Queue-0__3.14-503.fc40.aarch64",
        sha256 = "174112fe546cb28ca17d8bc2c1a5ff282a634c3eb9141f8c9bd2dc23e79f5ff7",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Thread-Queue-3.14-503.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Thread-Queue-3.14-503.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Thread-Queue-3.14-503.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Thread-Queue-3.14-503.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Thread-Queue-3.14-503.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Thread-Semaphore-0__2.13-506.fc40.aarch64",
        sha256 = "237434e210b6b3bf96b8f0b3cfb86dcb132f9ec15277813af5a3c2862f7473e0",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Thread-Semaphore-2.13-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Thread-Semaphore-2.13-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Thread-Semaphore-2.13-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Thread-Semaphore-2.13-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Thread-Semaphore-2.13-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Tie-0__4.6-506.fc40.aarch64",
        sha256 = "bcf8d4bd280954fc9d8ca113f5a52a62c188f0faa958d14a865ea8b826526fb0",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Tie-4.6-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Tie-4.6-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Tie-4.6-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Tie-4.6-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Tie-4.6-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Tie-File-0__1.07-506.fc40.aarch64",
        sha256 = "d1d415936b536a49f8ae1236f9954fd7798e84953de36f4528327b999d76bbff",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Tie-File-1.07-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Tie-File-1.07-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Tie-File-1.07-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Tie-File-1.07-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Tie-File-1.07-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Tie-Memoize-0__1.1-506.fc40.aarch64",
        sha256 = "2efa634835b838f39693c334a255ed0bdf9502e0afcff6395edea1d32219d0ab",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Tie-Memoize-1.1-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Tie-Memoize-1.1-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Tie-Memoize-1.1-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Tie-Memoize-1.1-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Tie-Memoize-1.1-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Tie-RefHash-0__1.41-1.fc40.aarch64",
        sha256 = "efd2d0b46057f225e059b73270ee8aab84d97481bd580894f698e5fe25546e0d",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/p/perl-Tie-RefHash-1.41-1.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/p/perl-Tie-RefHash-1.41-1.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/p/perl-Tie-RefHash-1.41-1.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/p/perl-Tie-RefHash-1.41-1.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/p/perl-Tie-RefHash-1.41-1.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Time-0__1.03-506.fc40.aarch64",
        sha256 = "7a505448da332cbf40a6f68393d6396686c8fe5dcacc6d09bd5009349b69349a",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Time-1.03-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Time-1.03-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Time-1.03-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Time-1.03-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Time-1.03-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Time-HiRes-4__1.9775-502.fc40.aarch64",
        sha256 = "61ee86e821e56e859e2cc4cb4f16c5145659be7efe634b3c889b475e2e7ab5e2",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Time-HiRes-1.9775-502.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Time-HiRes-1.9775-502.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Time-HiRes-1.9775-502.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Time-HiRes-1.9775-502.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Time-HiRes-1.9775-502.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "perl-Time-Local-2__1.350-5.fc40.aarch64",
        sha256 = "c6cbd7eef0215eceb66866ebc57d8311d220be04857c54ff478dd320f2be146a",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Time-Local-1.350-5.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Time-Local-1.350-5.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Time-Local-1.350-5.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Time-Local-1.350-5.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Time-Local-1.350-5.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Time-Piece-0__1.3401-506.fc40.aarch64",
        sha256 = "8b74550106809e4efba63bf8dd1fbcb1e00fce77716cb2301b3e4a3c19c1d76b",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Time-Piece-1.3401-506.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Time-Piece-1.3401-506.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Time-Piece-1.3401-506.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Time-Piece-1.3401-506.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Time-Piece-1.3401-506.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "perl-URI-0__5.28-1.fc40.aarch64",
        sha256 = "667865fb93f3851228eb29e5403759315e6e34b0e319b43857eb1d46d0e002bb",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/p/perl-URI-5.28-1.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/p/perl-URI-5.28-1.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/p/perl-URI-5.28-1.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/p/perl-URI-5.28-1.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/p/perl-URI-5.28-1.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Unicode-Collate-0__1.31-502.fc40.aarch64",
        sha256 = "3c645b0c28220b98a75f788bdd3d98f9c54abef6a512701066c712d3b562688e",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Unicode-Collate-1.31-502.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Unicode-Collate-1.31-502.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Unicode-Collate-1.31-502.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Unicode-Collate-1.31-502.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Unicode-Collate-1.31-502.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "perl-Unicode-Normalize-0__1.32-502.fc40.aarch64",
        sha256 = "4775c432f319bcf799dfee43a6f06f81337f3fd7841725794262ddc3875e1d9c",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Unicode-Normalize-1.32-502.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Unicode-Normalize-1.32-502.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Unicode-Normalize-1.32-502.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Unicode-Normalize-1.32-502.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Unicode-Normalize-1.32-502.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "perl-Unicode-UCD-0__0.78-506.fc40.aarch64",
        sha256 = "42ff613815f6a56c3b04967260e7fcf1d4a69f6f70a545a09f98157a091ada3d",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Unicode-UCD-0.78-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Unicode-UCD-0.78-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Unicode-UCD-0.78-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Unicode-UCD-0.78-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-Unicode-UCD-0.78-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-User-pwent-0__1.04-506.fc40.aarch64",
        sha256 = "07a55650bfc424e1ee3376bd3503f12629c31da76c24ee0154590107084b2f25",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-User-pwent-1.04-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-User-pwent-1.04-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-User-pwent-1.04-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-User-pwent-1.04-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-User-pwent-1.04-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-autodie-0__2.37-3.fc40.aarch64",
        sha256 = "2ecd5dd679dfa411279b2bd54b0ac25d37b80a2d2275c6d89813c61fa26325aa",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-autodie-2.37-3.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-autodie-2.37-3.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-autodie-2.37-3.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-autodie-2.37-3.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-autodie-2.37-3.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-autouse-0__1.11-506.fc40.aarch64",
        sha256 = "d4385b5a7ef78179e18f6baa3fb323f0494dc8112aa19d16b6c82d7b27adbf11",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-autouse-1.11-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-autouse-1.11-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-autouse-1.11-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-autouse-1.11-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-autouse-1.11-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-base-0__2.27-506.fc40.aarch64",
        sha256 = "a1325c1dc9d7379bb2b786371a7bbdc2c0d217ae73da040960c05279bd46c1ef",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-base-2.27-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-base-2.27-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-base-2.27-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-base-2.27-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-base-2.27-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-bignum-0__0.67-3.fc40.aarch64",
        sha256 = "66f91abbd717c643c9d3fca83356de79c37b5a05e837ecc04e4f642801c48e75",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-bignum-0.67-3.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-bignum-0.67-3.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-bignum-0.67-3.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-bignum-0.67-3.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-bignum-0.67-3.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-blib-0__1.07-506.fc40.aarch64",
        sha256 = "0134ec75e667d9540b06cf36357b0eb770148669e12087cbdfd4eb81c78e72eb",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-blib-1.07-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-blib-1.07-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-blib-1.07-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-blib-1.07-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-blib-1.07-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-constant-0__1.33-503.fc40.aarch64",
        sha256 = "938e497758f54c450e743427cf97e7b4a57399efdf665cbeb35d9a80c7633632",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-constant-1.33-503.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-constant-1.33-503.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-constant-1.33-503.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-constant-1.33-503.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-constant-1.33-503.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-debugger-0__1.60-506.fc40.aarch64",
        sha256 = "ddd5bb33baa41c6cc8787d80ef6a3572ded99721811ebd6236cd1f71e2ac37a6",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-debugger-1.60-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-debugger-1.60-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-debugger-1.60-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-debugger-1.60-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-debugger-1.60-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-deprecate-0__0.04-506.fc40.aarch64",
        sha256 = "11280daf1ceac8e18988ea2d14f094b15f08508c87b7f9124c874f315b81916c",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-deprecate-0.04-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-deprecate-0.04-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-deprecate-0.04-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-deprecate-0.04-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-deprecate-0.04-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-devel-4__5.38.2-506.fc40.aarch64",
        sha256 = "38d3ad230bf12973efbee62b3434df805b46877d46ed27adb3d27f9bb1ed5bc1",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-devel-5.38.2-506.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-devel-5.38.2-506.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-devel-5.38.2-506.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-devel-5.38.2-506.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-devel-5.38.2-506.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "perl-diagnostics-0__1.39-506.fc40.aarch64",
        sha256 = "fab5ac7514c649793d4a618a2a8c53743eafa992ae4926e218b8ca2c8f85a578",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-diagnostics-1.39-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-diagnostics-1.39-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-diagnostics-1.39-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-diagnostics-1.39-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-diagnostics-1.39-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-doc-0__5.38.2-506.fc40.aarch64",
        sha256 = "03aa3c80af55c2e7d4b71fce9e148e41656a62a51974cb7ed4b5f6edf276307a",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-doc-5.38.2-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-doc-5.38.2-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-doc-5.38.2-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-doc-5.38.2-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-doc-5.38.2-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-encoding-4__3.00-505.fc40.aarch64",
        sha256 = "4bb1c49869b7eba8b2d069de21ee1372e731564e3537b9742f17c9a8ae55241a",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-encoding-3.00-505.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-encoding-3.00-505.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-encoding-3.00-505.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-encoding-3.00-505.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-encoding-3.00-505.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "perl-encoding-warnings-0__0.14-506.fc40.aarch64",
        sha256 = "52c022f98b3aaa0397c53473eed4e32efe2ca6a378faf190e62dbde5361eeb62",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-encoding-warnings-0.14-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-encoding-warnings-0.14-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-encoding-warnings-0.14-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-encoding-warnings-0.14-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-encoding-warnings-0.14-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-experimental-0__0.032-1.fc40.aarch64",
        sha256 = "f36fbd8d75427451d0aed0e7e15d59c2a7865e91a2b5137678a9ac37e301042d",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/p/perl-experimental-0.032-1.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/p/perl-experimental-0.032-1.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/p/perl-experimental-0.032-1.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/p/perl-experimental-0.032-1.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/p/perl-experimental-0.032-1.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-fields-0__2.27-506.fc40.aarch64",
        sha256 = "97169a62eaebec86a7f3dcb3bc9b2bae440d091db84fcf6847e59b98243c7efa",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-fields-2.27-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-fields-2.27-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-fields-2.27-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-fields-2.27-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-fields-2.27-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-filetest-0__1.03-506.fc40.aarch64",
        sha256 = "69e7107f05bce5b8574de37dddde94b29674af5fb1df6b5ee9b16f0829733e8b",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-filetest-1.03-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-filetest-1.03-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-filetest-1.03-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-filetest-1.03-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-filetest-1.03-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-if-0__0.61.000-506.fc40.aarch64",
        sha256 = "8861b716151717d5545b97cf3c8f7bb6fab5563c97a939e35b1be98553e193d8",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-if-0.61.000-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-if-0.61.000-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-if-0.61.000-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-if-0.61.000-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-if-0.61.000-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-inc-latest-2__0.500-28.fc40.aarch64",
        sha256 = "6b9b606bc79f133bc1e9efaca9855c8d65b4ec1ce0566fc53dc260bd06111b36",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-inc-latest-0.500-28.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-inc-latest-0.500-28.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-inc-latest-0.500-28.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-inc-latest-0.500-28.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-inc-latest-0.500-28.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-interpreter-4__5.38.2-506.fc40.aarch64",
        sha256 = "2e5853654b9354a2825451a1de182065227776599c183bf59aa027c521ea185b",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-interpreter-5.38.2-506.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-interpreter-5.38.2-506.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-interpreter-5.38.2-506.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-interpreter-5.38.2-506.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-interpreter-5.38.2-506.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "perl-less-0__0.03-506.fc40.aarch64",
        sha256 = "b1802a74fb540e5c58a5682f871b1fe7773ac89cd3149bf3a76cca4d48469dca",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-less-0.03-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-less-0.03-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-less-0.03-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-less-0.03-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-less-0.03-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-lib-0__0.65-506.fc40.aarch64",
        sha256 = "d672e49f9d34b0d36c0de316d73db886c7210251c43dca165e657012a44e6f11",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-lib-0.65-506.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-lib-0.65-506.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-lib-0.65-506.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-lib-0.65-506.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-lib-0.65-506.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "perl-libnet-0__3.15-503.fc40.aarch64",
        sha256 = "a70dcd9f231e55757dce04f454e4cb0109edd3be7a91ab01ac34c43e65398160",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-libnet-3.15-503.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-libnet-3.15-503.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-libnet-3.15-503.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-libnet-3.15-503.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-libnet-3.15-503.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-libnetcfg-4__5.38.2-506.fc40.aarch64",
        sha256 = "89c8e1c30187f7355f61f1b054b4edbf07d99d178ab923741e3d0c0f5aee1a7b",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-libnetcfg-5.38.2-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-libnetcfg-5.38.2-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-libnetcfg-5.38.2-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-libnetcfg-5.38.2-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-libnetcfg-5.38.2-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-libs-4__5.38.2-506.fc40.aarch64",
        sha256 = "a70740ce170598f6446f7283b03496102d8ea38e1f6cbd8f085c82d654517ede",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-libs-5.38.2-506.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-libs-5.38.2-506.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-libs-5.38.2-506.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-libs-5.38.2-506.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-libs-5.38.2-506.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "perl-local-lib-0__2.000029-7.fc40.aarch64",
        sha256 = "9905a6bbac2773979e82474ea1482a40872070cb83b70a47cbe906b06f2afedd",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-local-lib-2.000029-7.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-local-lib-2.000029-7.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-local-lib-2.000029-7.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-local-lib-2.000029-7.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-local-lib-2.000029-7.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-locale-0__1.10-506.fc40.aarch64",
        sha256 = "9113463465c28d02db79ea499f3cfb4a1e508666ca99c776b2068a964c704c49",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-locale-1.10-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-locale-1.10-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-locale-1.10-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-locale-1.10-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-locale-1.10-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-macros-4__5.38.2-506.fc40.aarch64",
        sha256 = "33103a7b62ab22501d690766d4b53540db3f1a36679321062af491111aa0950f",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-macros-5.38.2-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-macros-5.38.2-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-macros-5.38.2-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-macros-5.38.2-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-macros-5.38.2-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-meta-notation-0__5.38.2-506.fc40.aarch64",
        sha256 = "cea0f11c4d3a085d5dbbe35e688a666813baf06706392f3ef6a7dfed5cf617b6",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-meta-notation-5.38.2-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-meta-notation-5.38.2-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-meta-notation-5.38.2-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-meta-notation-5.38.2-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-meta-notation-5.38.2-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-mro-0__1.28-506.fc40.aarch64",
        sha256 = "400d68426452afef477f79cda360aae8dbb8ac836b44145e07dfc5fc2fe8f908",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-mro-1.28-506.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-mro-1.28-506.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-mro-1.28-506.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-mro-1.28-506.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-mro-1.28-506.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "perl-open-0__1.13-506.fc40.aarch64",
        sha256 = "6fc81615c330e36696ffb1ba6ed601d379dc63cdcbdbbfdbc5287dbd248a3383",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-open-1.13-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-open-1.13-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-open-1.13-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-open-1.13-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-open-1.13-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-overload-0__1.37-506.fc40.aarch64",
        sha256 = "32659c3ebd0c02df994c7162e929611cdde63180cc3ab7a8c77d3656389f3157",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-overload-1.37-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-overload-1.37-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-overload-1.37-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-overload-1.37-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-overload-1.37-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-overloading-0__0.02-506.fc40.aarch64",
        sha256 = "97c29dfb29715c6dd3b45d0b5c3d6cd14da1a041b20540db7a4b01f41b6d6ffd",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-overloading-0.02-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-overloading-0.02-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-overloading-0.02-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-overloading-0.02-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-overloading-0.02-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-parent-1__0.241-502.fc40.aarch64",
        sha256 = "43701d6fd82fd42e15823efe40ae1304373e4d17da266ba639b4e9dfb78ba5b4",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-parent-0.241-502.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-parent-0.241-502.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-parent-0.241-502.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-parent-0.241-502.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-parent-0.241-502.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-perlfaq-0__5.20240218-1.fc40.aarch64",
        sha256 = "804bbbf920ad110ac88eeda16dc91b29e2bc07a7a4007803a3848f11be10333a",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-perlfaq-5.20240218-1.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-perlfaq-5.20240218-1.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-perlfaq-5.20240218-1.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-perlfaq-5.20240218-1.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-perlfaq-5.20240218-1.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-ph-0__5.38.2-506.fc40.aarch64",
        sha256 = "ae849b8bf916df203925994f8cd64186fc70b9f85483d8fa22f0db843f2b1cc8",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-ph-5.38.2-506.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-ph-5.38.2-506.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-ph-5.38.2-506.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-ph-5.38.2-506.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-ph-5.38.2-506.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "perl-podlators-1__5.01-502.fc40.aarch64",
        sha256 = "b3d9b83fa34d7b5b448dedaf535027b102c620feaead777af09c46727b0986c3",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-podlators-5.01-502.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-podlators-5.01-502.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-podlators-5.01-502.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-podlators-5.01-502.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-podlators-5.01-502.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-sigtrap-0__1.10-506.fc40.aarch64",
        sha256 = "cd9c4a53a2e865c8e43ce8f5cde4412490fcc55e2ed226239b3c04e6b60c4099",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-sigtrap-1.10-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-sigtrap-1.10-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-sigtrap-1.10-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-sigtrap-1.10-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-sigtrap-1.10-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-sort-0__2.05-506.fc40.aarch64",
        sha256 = "f4109bfab8689144a0c68deff309790069c5a5bfbe985c722491160a4d30ec1d",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-sort-2.05-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-sort-2.05-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-sort-2.05-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-sort-2.05-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-sort-2.05-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-srpm-macros-0__1-53.fc40.aarch64",
        sha256 = "076aab9e67fd58346b9c8ac369aaef8d52b1aeff4d2d21c9550931e03c6fee26",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-srpm-macros-1-53.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-srpm-macros-1-53.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-srpm-macros-1-53.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-srpm-macros-1-53.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-srpm-macros-1-53.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-subs-0__1.04-506.fc40.aarch64",
        sha256 = "a6793e171929f779c3cc3919c8f701912284b8e24d38c8b2a41d19c69658daa7",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-subs-1.04-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-subs-1.04-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-subs-1.04-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-subs-1.04-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-subs-1.04-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-threads-1__2.36-503.fc40.aarch64",
        sha256 = "64ddc7931c79e4eb6f922c60d0e4bc00c2f5b07c1b448fc71b89a7eb3efbd399",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-threads-2.36-503.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-threads-2.36-503.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-threads-2.36-503.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-threads-2.36-503.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-threads-2.36-503.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "perl-threads-shared-0__1.68-502.fc40.aarch64",
        sha256 = "1df25814bf77158bb813d291043cf9f80feb6adb9228e83f581c0e47f8657431",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-threads-shared-1.68-502.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-threads-shared-1.68-502.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-threads-shared-1.68-502.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-threads-shared-1.68-502.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-threads-shared-1.68-502.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "perl-utils-0__5.38.2-506.fc40.aarch64",
        sha256 = "eeff6ef4a3063c29c9080e38a4321a94723f1e0ae4b803edaef6b37fc0d8e7b6",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-utils-5.38.2-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-utils-5.38.2-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-utils-5.38.2-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-utils-5.38.2-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-utils-5.38.2-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-vars-0__1.05-506.fc40.aarch64",
        sha256 = "1c849b40e23b094c8c7d15b26f2e0839e1c1889f0f5892112616bb71a34c9099",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-vars-1.05-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-vars-1.05-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-vars-1.05-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-vars-1.05-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-vars-1.05-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-version-8__0.99.32-1.fc40.aarch64",
        sha256 = "4238fcd8bebdeabd6bc3a8f33197def8deb4b1fa54e42cd032f7957bd42e90e4",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/p/perl-version-0.99.32-1.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/p/perl-version-0.99.32-1.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/p/perl-version-0.99.32-1.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/p/perl-version-0.99.32-1.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/p/perl-version-0.99.32-1.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "perl-vmsish-0__1.04-506.fc40.aarch64",
        sha256 = "c6076c350e5e05ee7192a26d7dfd21b323a51f26d48b78ca7f680998f11a6783",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-vmsish-1.04-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-vmsish-1.04-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-vmsish-1.04-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-vmsish-1.04-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/perl-vmsish-1.04-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "pixman-0__0.43.4-1.fc40.aarch64",
        sha256 = "05bec493ec8d4898cae8e602399cce613b1deccec614fdac6218c505d54d60df",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/p/pixman-0.43.4-1.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/p/pixman-0.43.4-1.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/p/pixman-0.43.4-1.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/p/pixman-0.43.4-1.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/p/pixman-0.43.4-1.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "pkgconf-0__2.1.1-2.fc40.aarch64",
        sha256 = "29ae8965635e9f8fca2049ca494c095a9fb19fac1682949e25a9109c5da9e06b",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/p/pkgconf-2.1.1-2.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/p/pkgconf-2.1.1-2.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/p/pkgconf-2.1.1-2.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/p/pkgconf-2.1.1-2.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/p/pkgconf-2.1.1-2.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "pkgconf-m4-0__2.1.1-2.fc40.aarch64",
        sha256 = "b470bae5560e1d676145e9d53f76136f7c7b02a272d055fe89bd744847b49594",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/p/pkgconf-m4-2.1.1-2.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/p/pkgconf-m4-2.1.1-2.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/p/pkgconf-m4-2.1.1-2.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/p/pkgconf-m4-2.1.1-2.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/p/pkgconf-m4-2.1.1-2.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "pkgconf-pkg-config-0__2.1.1-2.fc40.aarch64",
        sha256 = "80f6afebddf5611f36f75d06f5c6fbd896da8c659d794a3be190fad2444a5019",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/p/pkgconf-pkg-config-2.1.1-2.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/p/pkgconf-pkg-config-2.1.1-2.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/p/pkgconf-pkg-config-2.1.1-2.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/p/pkgconf-pkg-config-2.1.1-2.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/p/pkgconf-pkg-config-2.1.1-2.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "policycoreutils-0__3.7-3.fc40.aarch64",
        sha256 = "527d9648e023b82eeef618da0d8846f2d816c37e3bdc19580ba56ce975531e3e",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/p/policycoreutils-3.7-3.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/p/policycoreutils-3.7-3.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/p/policycoreutils-3.7-3.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/p/policycoreutils-3.7-3.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/p/policycoreutils-3.7-3.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "policycoreutils-python-utils-0__3.7-3.fc40.aarch64",
        sha256 = "7d3aa818a87d3e97fde7fae85e162e07bbe82a3bb5c842aa7e96957b13b110b5",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/p/policycoreutils-python-utils-3.7-3.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/p/policycoreutils-python-utils-3.7-3.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/p/policycoreutils-python-utils-3.7-3.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/p/policycoreutils-python-utils-3.7-3.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/p/policycoreutils-python-utils-3.7-3.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "popt-0__1.19-6.fc40.aarch64",
        sha256 = "81d16164fcb81126bfef02fa20e47d32ed303ad53ab75bfdef7cfdabe4f4e598",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/popt-1.19-6.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/popt-1.19-6.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/popt-1.19-6.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/popt-1.19-6.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/popt-1.19-6.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "publicsuffix-list-dafsa-0__20240107-3.fc40.aarch64",
        sha256 = "cca50802d4f75306bc37126feb92db79fed44dcdabf76c1556853334995b9d3b",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/publicsuffix-list-dafsa-20240107-3.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/publicsuffix-list-dafsa-20240107-3.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/publicsuffix-list-dafsa-20240107-3.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/publicsuffix-list-dafsa-20240107-3.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/publicsuffix-list-dafsa-20240107-3.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "pyproject-srpm-macros-0__1.16.3-1.fc40.aarch64",
        sha256 = "ebf9d071552bc4ee90737be5b38f91433b3ae53075bd325648a934dbce146f8f",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/p/pyproject-srpm-macros-1.16.3-1.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/p/pyproject-srpm-macros-1.16.3-1.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/p/pyproject-srpm-macros-1.16.3-1.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/p/pyproject-srpm-macros-1.16.3-1.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/p/pyproject-srpm-macros-1.16.3-1.fc40.noarch.rpm",
        ],
    )

    rpm(
        name = "python-pip-wheel-0__23.3.2-2.fc40.aarch64",
        sha256 = "7c703b431508f44c5184b5c1df052ed0f49b7439d68aa3597a9a57a5b26bd648",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/p/python-pip-wheel-23.3.2-2.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/p/python-pip-wheel-23.3.2-2.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/p/python-pip-wheel-23.3.2-2.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/p/python-pip-wheel-23.3.2-2.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/p/python-pip-wheel-23.3.2-2.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "python-srpm-macros-0__3.12-8.fc40.aarch64",
        sha256 = "6ea431da8ae16131fcf943610f0bafa6405eea585d96978e4f02854d7a1437cf",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/p/python-srpm-macros-3.12-8.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/p/python-srpm-macros-3.12-8.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/p/python-srpm-macros-3.12-8.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/p/python-srpm-macros-3.12-8.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/p/python-srpm-macros-3.12-8.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "python-unversioned-command-0__3.12.8-2.fc40.aarch64",
        sha256 = "86e17167996c17798e116974f42e63dc2e0ac6bce1c10a47416d421c785a5ea4",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/p/python-unversioned-command-3.12.8-2.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/p/python-unversioned-command-3.12.8-2.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/p/python-unversioned-command-3.12.8-2.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/p/python-unversioned-command-3.12.8-2.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/p/python-unversioned-command-3.12.8-2.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "python3-0__3.12.8-2.fc40.aarch64",
        sha256 = "e5d876bd4d2c6ca974160845534fda6667a36c55261fb4a1c4aace68fe9ee92f",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/p/python3-3.12.8-2.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/p/python3-3.12.8-2.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/p/python3-3.12.8-2.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/p/python3-3.12.8-2.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/p/python3-3.12.8-2.fc40.aarch64.rpm",
        ],
    )

    rpm(
        name = "python3-audit-0__4.0.2-1.fc40.aarch64",
        sha256 = "ac44034f0b3b39d2c28d2fd0fda78b010a829bc6de7e4c600cafd56cfe4a93bb",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/p/python3-audit-4.0.2-1.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/p/python3-audit-4.0.2-1.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/p/python3-audit-4.0.2-1.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/p/python3-audit-4.0.2-1.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/p/python3-audit-4.0.2-1.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "python3-distro-0__1.9.0-3.fc40.aarch64",
        sha256 = "00507cbbee67333b446b0ebce7c8aa6395dffd97e22bf79766ecc7088c6c0d71",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/python3-distro-1.9.0-3.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/python3-distro-1.9.0-3.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/python3-distro-1.9.0-3.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/python3-distro-1.9.0-3.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/python3-distro-1.9.0-3.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "python3-libs-0__3.12.8-2.fc40.aarch64",
        sha256 = "08fd6d820203916b4250e20237fec9dbdb0d67a101246c6e9f3abf5399ba1c08",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/p/python3-libs-3.12.8-2.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/p/python3-libs-3.12.8-2.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/p/python3-libs-3.12.8-2.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/p/python3-libs-3.12.8-2.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/p/python3-libs-3.12.8-2.fc40.aarch64.rpm",
        ],
    )

    rpm(
        name = "python3-libselinux-0__3.7-5.fc40.aarch64",
        sha256 = "ab64e2bfd85875a39ee44378714e2e2405c5fcb95db53a9ab9b797695c3fae2b",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/p/python3-libselinux-3.7-5.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/p/python3-libselinux-3.7-5.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/p/python3-libselinux-3.7-5.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/p/python3-libselinux-3.7-5.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/p/python3-libselinux-3.7-5.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "python3-libsemanage-0__3.7-2.fc40.aarch64",
        sha256 = "02ba539a9eb73441282a7369e473e8e50d09c498add74f7d9aef84d5092925ed",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/p/python3-libsemanage-3.7-2.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/p/python3-libsemanage-3.7-2.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/p/python3-libsemanage-3.7-2.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/p/python3-libsemanage-3.7-2.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/p/python3-libsemanage-3.7-2.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "python3-policycoreutils-0__3.7-3.fc40.aarch64",
        sha256 = "6f225b0c95c58896d646f92289c944d7d79b8603d35b7c4f1a4f9edcc1d01156",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/p/python3-policycoreutils-3.7-3.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/p/python3-policycoreutils-3.7-3.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/p/python3-policycoreutils-3.7-3.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/p/python3-policycoreutils-3.7-3.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/p/python3-policycoreutils-3.7-3.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "python3-pyparsing-0__3.1.2-2.fc40.aarch64",
        sha256 = "dda9238b75b7a6bca8393907089a397f139003434bdeeff7d4d350bee1cc7d39",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/python3-pyparsing-3.1.2-2.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/python3-pyparsing-3.1.2-2.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/python3-pyparsing-3.1.2-2.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/python3-pyparsing-3.1.2-2.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/p/python3-pyparsing-3.1.2-2.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "python3-setools-0__4.5.1-2.fc40.aarch64",
        sha256 = "2db7bb539c6dfa8728aa0eb10a48392c4d58464f08235ce79a4df0582bd92222",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/p/python3-setools-4.5.1-2.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/p/python3-setools-4.5.1-2.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/p/python3-setools-4.5.1-2.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/p/python3-setools-4.5.1-2.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/p/python3-setools-4.5.1-2.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "python3-setuptools-0__69.0.3-4.fc40.aarch64",
        sha256 = "89a75463674f5e878374c7e2bfe094efcbf8bba705d0998f9a68f1cef74f12d5",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/p/python3-setuptools-69.0.3-4.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/p/python3-setuptools-69.0.3-4.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/p/python3-setuptools-69.0.3-4.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/p/python3-setuptools-69.0.3-4.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/p/python3-setuptools-69.0.3-4.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "qemu-common-2__8.2.8-2.fc40.aarch64",
        sha256 = "c34af2a8947708886987397a50b8bdfd47d9ff8fddb5bc21aae5ece125073916",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/q/qemu-common-8.2.8-2.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/q/qemu-common-8.2.8-2.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/q/qemu-common-8.2.8-2.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/q/qemu-common-8.2.8-2.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/q/qemu-common-8.2.8-2.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "qemu-img-2__8.2.8-2.fc40.aarch64",
        sha256 = "ff02a51e6049d3161ac71ec28b14d6df3c2e948597228c401214157ca9ca0078",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/q/qemu-img-8.2.8-2.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/q/qemu-img-8.2.8-2.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/q/qemu-img-8.2.8-2.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/q/qemu-img-8.2.8-2.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/q/qemu-img-8.2.8-2.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "qemu-system-aarch64-core-2__8.2.8-2.fc40.aarch64",
        sha256 = "16a8d8379d6a19f5b1f873fe75731d4c0e19d0f739f6351e7193efc11d411f48",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/q/qemu-system-aarch64-core-8.2.8-2.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/q/qemu-system-aarch64-core-8.2.8-2.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/q/qemu-system-aarch64-core-8.2.8-2.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/q/qemu-system-aarch64-core-8.2.8-2.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/q/qemu-system-aarch64-core-8.2.8-2.fc40.aarch64.rpm",
        ],
    )

    rpm(
        name = "qemu-system-x86-core-2__8.2.8-2.fc40.aarch64",
        sha256 = "201292b31a190c4c1342e6d235f124105a42abea72c0c080c80c5433abcddaed",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/q/qemu-system-x86-core-8.2.8-2.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/q/qemu-system-x86-core-8.2.8-2.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/q/qemu-system-x86-core-8.2.8-2.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/q/qemu-system-x86-core-8.2.8-2.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/q/qemu-system-x86-core-8.2.8-2.fc40.aarch64.rpm",
        ],
    )

    rpm(
        name = "qt5-srpm-macros-0__5.15.15-1.fc40.aarch64",
        sha256 = "3964b93f36be9a4570d882c2886939eba4df0a880132945d7deb47b21b854bd5",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/q/qt5-srpm-macros-5.15.15-1.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/q/qt5-srpm-macros-5.15.15-1.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/q/qt5-srpm-macros-5.15.15-1.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/q/qt5-srpm-macros-5.15.15-1.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/q/qt5-srpm-macros-5.15.15-1.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "qt6-srpm-macros-0__6.7.2-2.fc40.aarch64",
        sha256 = "a7a1e173b543c524249f8a7eef986f942c89030c0ee7b77ab95faa35c0f4372c",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/q/qt6-srpm-macros-6.7.2-2.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/q/qt6-srpm-macros-6.7.2-2.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/q/qt6-srpm-macros-6.7.2-2.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/q/qt6-srpm-macros-6.7.2-2.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/q/qt6-srpm-macros-6.7.2-2.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "readline-0__8.2-8.fc40.aarch64",
        sha256 = "deef4be470d99ff6da3db0af264aed7d09b798b6e1482167f1793ca4564cbe8b",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/r/readline-8.2-8.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/r/readline-8.2-8.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/r/readline-8.2-8.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/r/readline-8.2-8.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/r/readline-8.2-8.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "redhat-rpm-config-0__288-1.fc40.aarch64",
        sha256 = "a71f0902957839e18a7f9e13caf4d37a3d53d1c3f5f51a4a57eec80b3edb948d",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/r/redhat-rpm-config-288-1.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/r/redhat-rpm-config-288-1.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/r/redhat-rpm-config-288-1.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/r/redhat-rpm-config-288-1.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/r/redhat-rpm-config-288-1.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "rpm-0__4.19.1.1-1.fc40.aarch64",
        sha256 = "5e9a1a380024b758fea8f9e8713630440430987abf49c75c10bbbfcdf93f33aa",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/r/rpm-4.19.1.1-1.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/r/rpm-4.19.1.1-1.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/r/rpm-4.19.1.1-1.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/r/rpm-4.19.1.1-1.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/r/rpm-4.19.1.1-1.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "rpm-libs-0__4.19.1.1-1.fc40.aarch64",
        sha256 = "c74d2eeb0e113982f27467de7fc715d41d52920e9596a168f995d7034c787342",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/r/rpm-libs-4.19.1.1-1.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/r/rpm-libs-4.19.1.1-1.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/r/rpm-libs-4.19.1.1-1.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/r/rpm-libs-4.19.1.1-1.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/r/rpm-libs-4.19.1.1-1.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "rpm-plugin-selinux-0__4.19.1.1-1.fc40.aarch64",
        sha256 = "25a309e64d016ece3bf36b5458032d8bf16440d9f07c342823f27cb3ea11d168",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/r/rpm-plugin-selinux-4.19.1.1-1.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/r/rpm-plugin-selinux-4.19.1.1-1.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/r/rpm-plugin-selinux-4.19.1.1-1.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/r/rpm-plugin-selinux-4.19.1.1-1.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/r/rpm-plugin-selinux-4.19.1.1-1.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "rpm-sequoia-0__1.7.0-3.fc40.aarch64",
        sha256 = "d7a5f5265760d8e01a3c7e9b121414b44e951a633eb724f6a42fab7a636524d0",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/r/rpm-sequoia-1.7.0-3.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/r/rpm-sequoia-1.7.0-3.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/r/rpm-sequoia-1.7.0-3.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/r/rpm-sequoia-1.7.0-3.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/r/rpm-sequoia-1.7.0-3.fc40.aarch64.rpm",
        ],
    )

    rpm(
        name = "rsync-0__3.3.0-1.fc40.aarch64",
        sha256 = "7350ee6dfacfa785fbf1c9e6b60120b0f18939d6d47eedcba4e01da17bd8a3bd",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/r/rsync-3.3.0-1.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/r/rsync-3.3.0-1.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/r/rsync-3.3.0-1.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/r/rsync-3.3.0-1.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/r/rsync-3.3.0-1.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "rust-srpm-macros-0__26.3-1.fc40.aarch64",
        sha256 = "5d0470c00b7b6102f383dd8845e7000377040f0bd79e6947034b03f1b84080ef",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/r/rust-srpm-macros-26.3-1.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/r/rust-srpm-macros-26.3-1.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/r/rust-srpm-macros-26.3-1.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/r/rust-srpm-macros-26.3-1.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/r/rust-srpm-macros-26.3-1.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "seabios-bin-0__1.16.3-2.fc40.aarch64",
        sha256 = "cac97b1c51e1ccbf9489c3b67417e018e887287f60a8520dd931578b5e422bf0",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/s/seabios-bin-1.16.3-2.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/s/seabios-bin-1.16.3-2.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/s/seabios-bin-1.16.3-2.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/s/seabios-bin-1.16.3-2.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/s/seabios-bin-1.16.3-2.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "seavgabios-bin-0__1.16.3-2.fc40.aarch64",
        sha256 = "31d20aaa2f430fca6184317a029c076a7405586929632ae6e044308d946e2f30",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/s/seavgabios-bin-1.16.3-2.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/s/seavgabios-bin-1.16.3-2.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/s/seavgabios-bin-1.16.3-2.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/s/seavgabios-bin-1.16.3-2.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/s/seavgabios-bin-1.16.3-2.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "sed-0__4.9-1.fc40.aarch64",
        sha256 = "3407bce63d6700c5b594155262ffb1fca0ea311341956041e7f2c6e22fa106f8",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/s/sed-4.9-1.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/s/sed-4.9-1.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/s/sed-4.9-1.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/s/sed-4.9-1.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/s/sed-4.9-1.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "selinux-policy-0__40.29-2.fc40.aarch64",
        sha256 = "b4e188db51c7ec2d5f0cba79783eb2df7c14a92c2c6e55a9eb490d28d17d123d",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/s/selinux-policy-40.29-2.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/s/selinux-policy-40.29-2.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/s/selinux-policy-40.29-2.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/s/selinux-policy-40.29-2.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/s/selinux-policy-40.29-2.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "selinux-policy-minimum-0__40.29-2.fc40.aarch64",
        sha256 = "89fb3e3e9053f35c9e92d988b1bfe300709fb25be039c7ebb93d99d89de26674",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/s/selinux-policy-minimum-40.29-2.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/s/selinux-policy-minimum-40.29-2.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/s/selinux-policy-minimum-40.29-2.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/s/selinux-policy-minimum-40.29-2.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/s/selinux-policy-minimum-40.29-2.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "setup-0__2.14.5-2.fc40.aarch64",
        sha256 = "89862f646cd64e81497f01a8b69ab30ac8968c47afef92a2c333608fdb90ccc1",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/s/setup-2.14.5-2.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/s/setup-2.14.5-2.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/s/setup-2.14.5-2.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/s/setup-2.14.5-2.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/s/setup-2.14.5-2.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "shadow-utils-2__4.15.1-4.fc40.aarch64",
        sha256 = "8b6f9ca65a58595296c96f46cd4f06a2c01d6150d98635529cd725e66f99e90d",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/s/shadow-utils-4.15.1-4.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/s/shadow-utils-4.15.1-4.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/s/shadow-utils-4.15.1-4.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/s/shadow-utils-4.15.1-4.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/s/shadow-utils-4.15.1-4.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "snappy-0__1.1.10-4.fc40.aarch64",
        sha256 = "382b917404f5f9e3d9219ce4402a859360fadfffea96a91e903b5a4060962f42",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/s/snappy-1.1.10-4.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/s/snappy-1.1.10-4.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/s/snappy-1.1.10-4.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/s/snappy-1.1.10-4.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/s/snappy-1.1.10-4.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "sqlite-libs-0__3.45.1-2.fc40.aarch64",
        sha256 = "ae1426147fb74bbdaa1dc95362bc6899e5e823487387972b8cc1f2746c0bd4b3",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/s/sqlite-libs-3.45.1-2.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/s/sqlite-libs-3.45.1-2.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/s/sqlite-libs-3.45.1-2.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/s/sqlite-libs-3.45.1-2.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/s/sqlite-libs-3.45.1-2.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "systemd-0__255.15-1.fc40.aarch64",
        sha256 = "a374c4d7a3c755f759c3ae89a6739acf94e4a8caf5693a3544a09938a27b612b",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/s/systemd-255.15-1.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/s/systemd-255.15-1.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/s/systemd-255.15-1.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/s/systemd-255.15-1.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/s/systemd-255.15-1.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "systemd-libs-0__255.15-1.fc40.aarch64",
        sha256 = "72e1ca2741c9800a8f9ff14b7f9d2d9c05c96d708a7b812f15b0fcad68dab619",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/s/systemd-libs-255.15-1.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/s/systemd-libs-255.15-1.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/s/systemd-libs-255.15-1.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/s/systemd-libs-255.15-1.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/s/systemd-libs-255.15-1.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "systemd-pam-0__255.15-1.fc40.aarch64",
        sha256 = "d9f5bbd39e05d66cfb4768baea45a1661e99a327b036402986ce667f603f2437",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/s/systemd-pam-255.15-1.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/s/systemd-pam-255.15-1.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/s/systemd-pam-255.15-1.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/s/systemd-pam-255.15-1.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/s/systemd-pam-255.15-1.fc40.aarch64.rpm",
        ],
    )

    rpm(
        name = "systemtap-sdt-devel-0__5.2-1.fc40.aarch64",
        sha256 = "b21b0f61f2e061b08816e0112f85afe50d6df2a74ea7402a48d4f0180bc540f1",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/s/systemtap-sdt-devel-5.2-1.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/s/systemtap-sdt-devel-5.2-1.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/s/systemtap-sdt-devel-5.2-1.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/s/systemtap-sdt-devel-5.2-1.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/s/systemtap-sdt-devel-5.2-1.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "tar-2__1.35-3.fc40.aarch64",
        sha256 = "35bd7bccd3252f7c2ac67c0e6ecff36cf6a03dc9fa782a9bdb9602d1ff0d415c",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/t/tar-1.35-3.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/t/tar-1.35-3.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/t/tar-1.35-3.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/t/tar-1.35-3.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/t/tar-1.35-3.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "tpm2-tss-0__4.1.3-1.fc40.aarch64",
        sha256 = "eb51b5f2e80cb5243bc3bbf85ec09b47d46f367623ba89f57df32f1ab70a909c",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/t/tpm2-tss-4.1.3-1.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/t/tpm2-tss-4.1.3-1.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/t/tpm2-tss-4.1.3-1.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/t/tpm2-tss-4.1.3-1.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/t/tpm2-tss-4.1.3-1.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "tzdata-0__2024a-5.fc40.aarch64",
        sha256 = "0bd358e7dfb2bd730b62c7375c8d8f8d9e2470f085ca8dc4ec626dc0332d5687",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/t/tzdata-2024a-5.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/t/tzdata-2024a-5.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/t/tzdata-2024a-5.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/t/tzdata-2024a-5.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/t/tzdata-2024a-5.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "unzip-0__6.0-63.fc40.aarch64",
        sha256 = "d944bbba138a2ef6cff881309ab2fbac59d942ad6c59b1557581863b84dfa60a",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/u/unzip-6.0-63.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/u/unzip-6.0-63.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/u/unzip-6.0-63.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/u/unzip-6.0-63.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/u/unzip-6.0-63.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "util-linux-0__2.40.2-1.fc40.aarch64",
        sha256 = "977b6f542f73eb324f87d2b7c912acfd9802d913c03f1423306bde9ac2ddba58",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/u/util-linux-2.40.2-1.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/u/util-linux-2.40.2-1.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/u/util-linux-2.40.2-1.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/u/util-linux-2.40.2-1.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/u/util-linux-2.40.2-1.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "util-linux-core-0__2.40.2-1.fc40.aarch64",
        sha256 = "efb493038815c8b228048c97bab1ba2724ad09e0ae45bc06e6e947894278b3b5",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/u/util-linux-core-2.40.2-1.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/u/util-linux-core-2.40.2-1.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/u/util-linux-core-2.40.2-1.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/u/util-linux-core-2.40.2-1.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/u/util-linux-core-2.40.2-1.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "xen-libs-0__4.18.4-1.fc40.aarch64",
        sha256 = "bf634fef64f0db75f7403048e58559a5e9b030f805043d365a52d14123600db3",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/x/xen-libs-4.18.4-1.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/x/xen-libs-4.18.4-1.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/x/xen-libs-4.18.4-1.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/x/xen-libs-4.18.4-1.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/x/xen-libs-4.18.4-1.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "xen-licenses-0__4.18.4-1.fc40.aarch64",
        sha256 = "c3910e797189120cb7a82eda69a7c4c8a5c0cf0035b4726ff50e4a257b93a45b",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/x/xen-licenses-4.18.4-1.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/x/xen-licenses-4.18.4-1.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/x/xen-licenses-4.18.4-1.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/x/xen-licenses-4.18.4-1.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/x/xen-licenses-4.18.4-1.fc40.aarch64.rpm",
        ],
    )

    rpm(
        name = "xxhash-libs-0__0.8.2-4.fc40.aarch64",
        sha256 = "6cd2b7c00e2867b463dc78e46c7b4e2f3befa29c730c5c44096813549d3395fe",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/x/xxhash-libs-0.8.2-4.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/x/xxhash-libs-0.8.2-4.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/x/xxhash-libs-0.8.2-4.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/x/xxhash-libs-0.8.2-4.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/x/xxhash-libs-0.8.2-4.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "xz-1__5.4.6-3.fc40.aarch64",
        sha256 = "44fce3f49bc919462ef1ff0c22c1b6243f6b97d5d11e48285c5f21495a805004",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/x/xz-5.4.6-3.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/x/xz-5.4.6-3.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/x/xz-5.4.6-3.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/x/xz-5.4.6-3.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/x/xz-5.4.6-3.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "xz-libs-1__5.4.6-3.fc40.aarch64",
        sha256 = "7e9c4963d8fa8abaf43ccd34051c428b8dccfec202f42881275db56d909577da",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/x/xz-libs-5.4.6-3.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/x/xz-libs-5.4.6-3.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/x/xz-libs-5.4.6-3.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/x/xz-libs-5.4.6-3.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/x/xz-libs-5.4.6-3.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "yajl-0__2.1.0-23.fc40.aarch64",
        sha256 = "8f09faa2fa8e2b71d41e2200d73cdd5e08e6263780589ee9b55001f98972e0af",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/y/yajl-2.1.0-23.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/y/yajl-2.1.0-23.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/y/yajl-2.1.0-23.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/y/yajl-2.1.0-23.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/y/yajl-2.1.0-23.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "zig-srpm-macros-0__1-2.fc40.aarch64",
        sha256 = "3957667c460ee5ed7c46c401db9e1366bd8a22921ed620ffd9a4d7e79298a8f0",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/z/zig-srpm-macros-1-2.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/z/zig-srpm-macros-1-2.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/z/zig-srpm-macros-1-2.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/z/zig-srpm-macros-1-2.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/z/zig-srpm-macros-1-2.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "zip-0__3.0-40.fc40.aarch64",
        sha256 = "69b043102bf0506bf86dc7f83e2e6d17dfe404915e1f80acbe13f0d3b512e5f3",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/z/zip-3.0-40.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/aarch64/os/Packages/z/zip-3.0-40.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/aarch64/os/Packages/z/zip-3.0-40.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/aarch64/os/Packages/z/zip-3.0-40.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/aarch64/os/Packages/z/zip-3.0-40.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "zlib-ng-compat-0__2.1.7-2.fc40.aarch64",
        sha256 = "45a067e9c06270bb63b6f3e36e0fb87c7f04ba451ff71ae82b42ccc081cde477",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/z/zlib-ng-compat-2.1.7-2.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/z/zlib-ng-compat-2.1.7-2.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/z/zlib-ng-compat-2.1.7-2.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/z/zlib-ng-compat-2.1.7-2.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/z/zlib-ng-compat-2.1.7-2.fc40.aarch64.rpm",
        ],
    )
    rpm(
        name = "zlib-ng-compat-devel-0__2.1.7-2.fc40.aarch64",
        sha256 = "7f8459321677931c886b09d99af036310e846a3abd45beadcccb6f7903c255be",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/aarch64/Packages/z/zlib-ng-compat-devel-2.1.7-2.fc40.aarch64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/aarch64/Packages/z/zlib-ng-compat-devel-2.1.7-2.fc40.aarch64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/aarch64/Packages/z/zlib-ng-compat-devel-2.1.7-2.fc40.aarch64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/aarch64/Packages/z/zlib-ng-compat-devel-2.1.7-2.fc40.aarch64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/aarch64/Packages/z/zlib-ng-compat-devel-2.1.7-2.fc40.aarch64.rpm",
        ],
    )
