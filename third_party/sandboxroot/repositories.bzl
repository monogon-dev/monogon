load("@bazeldnf//:deps.bzl", "rpm")

def sandbox_dependencies():
    rpm(
        name = "acpica-tools-0__20220331-4.fc37.x86_64",
        sha256 = "ab044a35844bf56c0a217d298faa7f390915c72be1f5d9fd241877aa76ccb9b7",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/a/acpica-tools-20220331-4.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/a/acpica-tools-20220331-4.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/a/acpica-tools-20220331-4.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/a/acpica-tools-20220331-4.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/a/acpica-tools-20220331-4.fc37.x86_64.rpm",
        ],
    )

    rpm(
        name = "alternatives-0__1.24-1.fc37.x86_64",
        sha256 = "4fd3c9e8aedd1cf2a620c7580999c8510a1b93f5c85be8e24c59ca9b3c27aa8a",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/a/alternatives-1.24-1.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/a/alternatives-1.24-1.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/a/alternatives-1.24-1.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/a/alternatives-1.24-1.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/a/alternatives-1.24-1.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "ansible-srpm-macros-0__1-10.fc37.x86_64",
        sha256 = "4438d542c8ca4253b0da207bbf876615d717d50f764896b54e4a8cd3058acf03",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/a/ansible-srpm-macros-1-10.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/a/ansible-srpm-macros-1-10.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/a/ansible-srpm-macros-1-10.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/a/ansible-srpm-macros-1-10.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/a/ansible-srpm-macros-1-10.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "audit-libs-0__3.1.2-1.fc37.x86_64",
        sha256 = "6193619a12f238f2e94c390d4e584094755f4b23e592ceafd53e65c6e22498ff",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/a/audit-libs-3.1.2-1.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/a/audit-libs-3.1.2-1.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/a/audit-libs-3.1.2-1.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/a/audit-libs-3.1.2-1.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/a/audit-libs-3.1.2-1.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "authselect-0__1.4.2-1.fc37.x86_64",
        sha256 = "c356d05e80f2b57ea2598b45b168fff6da189038e3f3ef0305dd90cfdd2a045f",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/a/authselect-1.4.2-1.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/a/authselect-1.4.2-1.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/a/authselect-1.4.2-1.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/a/authselect-1.4.2-1.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/a/authselect-1.4.2-1.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "authselect-libs-0__1.4.2-1.fc37.x86_64",
        sha256 = "275c282a240a3b7225e98b540a91af3419a9fa527623c5f152c48f8209779146",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/a/authselect-libs-1.4.2-1.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/a/authselect-libs-1.4.2-1.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/a/authselect-libs-1.4.2-1.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/a/authselect-libs-1.4.2-1.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/a/authselect-libs-1.4.2-1.fc37.x86_64.rpm",
        ],
    )

    rpm(
        name = "basesystem-0__11-14.fc37.x86_64",
        sha256 = "38d1877d647bb5f4047d22982a51899c95bdfea1d7b2debbff37c66f0fc0ed44",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/b/basesystem-11-14.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/b/basesystem-11-14.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/b/basesystem-11-14.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/b/basesystem-11-14.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/b/basesystem-11-14.fc37.noarch.rpm",
        ],
    )

    rpm(
        name = "bash-0__5.2.15-1.fc37.x86_64",
        sha256 = "e50ddbdb35ecec1a9bf4e19fd87c6216382be313c3b671704d444053a1cfd183",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/b/bash-5.2.15-1.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/b/bash-5.2.15-1.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/b/bash-5.2.15-1.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/b/bash-5.2.15-1.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/b/bash-5.2.15-1.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "bc-0__1.07.1-16.fc37.x86_64",
        sha256 = "5641d8a1ffc675c13d108fa5218024ae8b35164abff34176ead3289756c79b8c",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/b/bc-1.07.1-16.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/b/bc-1.07.1-16.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/b/bc-1.07.1-16.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/b/bc-1.07.1-16.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/b/bc-1.07.1-16.fc37.x86_64.rpm",
        ],
    )

    rpm(
        name = "binutils-0__2.38-27.fc37.x86_64",
        sha256 = "d9efae81d1c849d7f981089047e9cb22845cc8c41404f8e9e2abdd19eea784b7",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/b/binutils-2.38-27.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/b/binutils-2.38-27.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/b/binutils-2.38-27.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/b/binutils-2.38-27.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/b/binutils-2.38-27.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "binutils-gold-0__2.38-27.fc37.x86_64",
        sha256 = "63f0c543c79929296d415357e88da768ed52e2c092e267671fe8a25cd69b2bc6",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/b/binutils-gold-2.38-27.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/b/binutils-gold-2.38-27.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/b/binutils-gold-2.38-27.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/b/binutils-gold-2.38-27.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/b/binutils-gold-2.38-27.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "bison-0__3.8.2-3.fc37.x86_64",
        sha256 = "2e6094d1f6670f4e99c3334f94b521055cbd780ac6ff0275c95b98a22c34d08a",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/b/bison-3.8.2-3.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/b/bison-3.8.2-3.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/b/bison-3.8.2-3.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/b/bison-3.8.2-3.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/b/bison-3.8.2-3.fc37.x86_64.rpm",
        ],
    )

    rpm(
        name = "bzip2-libs-0__1.0.8-12.fc37.x86_64",
        sha256 = "6e74a8ed5b472cf811f9bf429a999ed3f362e2c88566a461517a12c058abd401",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/b/bzip2-libs-1.0.8-12.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/b/bzip2-libs-1.0.8-12.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/b/bzip2-libs-1.0.8-12.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/b/bzip2-libs-1.0.8-12.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/b/bzip2-libs-1.0.8-12.fc37.x86_64.rpm",
        ],
    )

    rpm(
        name = "ca-certificates-0__2023.2.60-1.0.fc37.x86_64",
        sha256 = "b2dcac3e49cbf75841d41ee1c53f1a91ffa78ba03dab8febb3153dbf76b2c5b2",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/c/ca-certificates-2023.2.60-1.0.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/c/ca-certificates-2023.2.60-1.0.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/c/ca-certificates-2023.2.60-1.0.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/c/ca-certificates-2023.2.60-1.0.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/c/ca-certificates-2023.2.60-1.0.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "capstone-0__4.0.2-11.fc37.x86_64",
        sha256 = "5faaf0c29c0e76456c42b3a8e62b4fdf43150501187653da27974e6384655248",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/c/capstone-4.0.2-11.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/c/capstone-4.0.2-11.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/c/capstone-4.0.2-11.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/c/capstone-4.0.2-11.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/c/capstone-4.0.2-11.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "checkpolicy-0__3.5-1.fc37.x86_64",
        sha256 = "1bd6036081fb219541cdbf1f23ef05381665356f3f45c4af8bb72ff0295f82ce",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/c/checkpolicy-3.5-1.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/c/checkpolicy-3.5-1.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/c/checkpolicy-3.5-1.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/c/checkpolicy-3.5-1.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/c/checkpolicy-3.5-1.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "clang-0__15.0.7-2.fc37.x86_64",
        sha256 = "9f4e9c4f003b285eb2f6988ec9c476c83006041262beaa07702d0b35394c396b",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/c/clang-15.0.7-2.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/c/clang-15.0.7-2.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/c/clang-15.0.7-2.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/c/clang-15.0.7-2.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/c/clang-15.0.7-2.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "clang-libs-0__15.0.7-2.fc37.x86_64",
        sha256 = "9c260e1f9512734eb00a7fdd2cdc156ac9b2ba4c855e661c90d2898853bf9c7f",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/c/clang-libs-15.0.7-2.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/c/clang-libs-15.0.7-2.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/c/clang-libs-15.0.7-2.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/c/clang-libs-15.0.7-2.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/c/clang-libs-15.0.7-2.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "clang-resource-filesystem-0__15.0.7-2.fc37.x86_64",
        sha256 = "505f703e657c33605d3ae395835b0d6d4c132e0cf812a95779aa0b2bb5ce2e77",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/c/clang-resource-filesystem-15.0.7-2.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/c/clang-resource-filesystem-15.0.7-2.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/c/clang-resource-filesystem-15.0.7-2.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/c/clang-resource-filesystem-15.0.7-2.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/c/clang-resource-filesystem-15.0.7-2.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "coreutils-single-0__9.1-8.fc37.x86_64",
        sha256 = "4a0f1b98fd4e00d2b362ebb11ec4ba1702ee363b1ecdcbcd5d53486d580f2edc",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/c/coreutils-single-9.1-8.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/c/coreutils-single-9.1-8.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/c/coreutils-single-9.1-8.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/c/coreutils-single-9.1-8.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/c/coreutils-single-9.1-8.fc37.x86_64.rpm",
        ],
    )

    rpm(
        name = "cpp-0__12.3.1-1.fc37.x86_64",
        sha256 = "845c873ec113f26fd1151083ee83d90a2521805275ec147d684e9f34b8fcc1d9",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/c/cpp-12.3.1-1.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/c/cpp-12.3.1-1.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/c/cpp-12.3.1-1.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/c/cpp-12.3.1-1.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/c/cpp-12.3.1-1.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "cracklib-0__2.9.7-30.fc37.x86_64",
        sha256 = "3847abdc8ff973aeb0fb7e681bdf7c37b19cd49e5df17e8bf6bc35f34615c88f",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/c/cracklib-2.9.7-30.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/c/cracklib-2.9.7-30.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/c/cracklib-2.9.7-30.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/c/cracklib-2.9.7-30.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/c/cracklib-2.9.7-30.fc37.x86_64.rpm",
        ],
    )

    rpm(
        name = "crypto-policies-0__20220815-1.gite4ed860.fc37.x86_64",
        sha256 = "486a11feeaad706c68b05de60a906cc57059454cbce436aeba45f88b84578c0c",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/c/crypto-policies-20220815-1.gite4ed860.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/c/crypto-policies-20220815-1.gite4ed860.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/c/crypto-policies-20220815-1.gite4ed860.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/c/crypto-policies-20220815-1.gite4ed860.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/c/crypto-policies-20220815-1.gite4ed860.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "curl-minimal-0__7.85.0-10.fc37.x86_64",
        sha256 = "ed6b0333085d49fc9b1b652752a926364ee0af3bcab437f01249a17a565bc09f",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/c/curl-minimal-7.85.0-10.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/c/curl-minimal-7.85.0-10.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/c/curl-minimal-7.85.0-10.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/c/curl-minimal-7.85.0-10.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/c/curl-minimal-7.85.0-10.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "cyrus-sasl-lib-0__2.1.28-8.fc37.x86_64",
        sha256 = "4e0e8656faf1f4f5227e4e40cdb4e662a1d78b19e74b90ba2f39f3cdf73e0083",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/c/cyrus-sasl-lib-2.1.28-8.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/c/cyrus-sasl-lib-2.1.28-8.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/c/cyrus-sasl-lib-2.1.28-8.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/c/cyrus-sasl-lib-2.1.28-8.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/c/cyrus-sasl-lib-2.1.28-8.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "daxctl-libs-0__78-1.fc37.x86_64",
        sha256 = "77d7bc0e5edfec318ab6ef62f211f6daa22367ad2ae4b2d5db193123fe6182ac",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/d/daxctl-libs-78-1.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/d/daxctl-libs-78-1.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/d/daxctl-libs-78-1.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/d/daxctl-libs-78-1.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/d/daxctl-libs-78-1.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "dbus-1__1.14.10-1.fc37.x86_64",
        sha256 = "ecf5c78fcd0a060f47ba13f0c719990f801adcf0c8a40bec22b0b5ae285571d9",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/d/dbus-1.14.10-1.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/d/dbus-1.14.10-1.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/d/dbus-1.14.10-1.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/d/dbus-1.14.10-1.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/d/dbus-1.14.10-1.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "dbus-broker-0__33-1.fc37.x86_64",
        sha256 = "069f79144219815854e47cda0bf47c5f5e361d48cbfa652405ac68c0d24d29ee",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/d/dbus-broker-33-1.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/d/dbus-broker-33-1.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/d/dbus-broker-33-1.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/d/dbus-broker-33-1.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/d/dbus-broker-33-1.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "dbus-common-1__1.14.10-1.fc37.x86_64",
        sha256 = "938bd46c3d7bf6a02cdf9507fd6dd47c0fe5d8b52862164cf5ca1cef36951ca5",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/d/dbus-common-1.14.10-1.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/d/dbus-common-1.14.10-1.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/d/dbus-common-1.14.10-1.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/d/dbus-common-1.14.10-1.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/d/dbus-common-1.14.10-1.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "diffutils-0__3.8-3.fc37.x86_64",
        sha256 = "c1374e3372d0d246ecb0e04b36743e23c68ab307c7603c5a267fce654bf05cdd",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/d/diffutils-3.8-3.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/d/diffutils-3.8-3.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/d/diffutils-3.8-3.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/d/diffutils-3.8-3.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/d/diffutils-3.8-3.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "dwz-0__0.14-7.fc37.x86_64",
        sha256 = "82e4d749edaf6e209f8e7eb830454fa12e8b39f88ee190a49fb55496e7eda9af",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/d/dwz-0.14-7.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/d/dwz-0.14-7.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/d/dwz-0.14-7.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/d/dwz-0.14-7.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/d/dwz-0.14-7.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "e2fsprogs-libs-0__1.46.5-3.fc37.x86_64",
        sha256 = "631c5cdd65015cf905cf9c7b9c0213384a524eeb0b1f15da971aae8cc38ed27e",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/e/e2fsprogs-libs-1.46.5-3.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/e/e2fsprogs-libs-1.46.5-3.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/e/e2fsprogs-libs-1.46.5-3.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/e/e2fsprogs-libs-1.46.5-3.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/e/e2fsprogs-libs-1.46.5-3.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "ed-0__1.18-2.fc37.x86_64",
        sha256 = "4a16aed8139002451a96be745b170ac036385afda72b0fac7903979df33e2762",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/e/ed-1.18-2.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/e/ed-1.18-2.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/e/ed-1.18-2.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/e/ed-1.18-2.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/e/ed-1.18-2.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "edk2-ovmf-0__20230524-3.fc37.x86_64",
        sha256 = "be4d38af60d1917c681ea76a5884eba2539aa008090d7c893f17ed8af39efdb0",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/e/edk2-ovmf-20230524-3.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/e/edk2-ovmf-20230524-3.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/e/edk2-ovmf-20230524-3.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/e/edk2-ovmf-20230524-3.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/e/edk2-ovmf-20230524-3.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "efi-srpm-macros-0__5-6.fc37.x86_64",
        sha256 = "d24933b643fabac0d1f63835c39dbb9081c94795427df3b07722dbf0748dd16d",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/e/efi-srpm-macros-5-6.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/e/efi-srpm-macros-5-6.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/e/efi-srpm-macros-5-6.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/e/efi-srpm-macros-5-6.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/e/efi-srpm-macros-5-6.fc37.noarch.rpm",
        ],
    )

    rpm(
        name = "elfutils-debuginfod-client-0__0.189-3.fc37.x86_64",
        sha256 = "e07cb3382bf16c9d452b693015fab9a6da32cdb3581f97882bf91463ec05d706",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/e/elfutils-debuginfod-client-0.189-3.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/e/elfutils-debuginfod-client-0.189-3.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/e/elfutils-debuginfod-client-0.189-3.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/e/elfutils-debuginfod-client-0.189-3.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/e/elfutils-debuginfod-client-0.189-3.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "elfutils-default-yama-scope-0__0.189-3.fc37.x86_64",
        sha256 = "108933bfd4359472c97e6a691053154705e164da743f0f639c04fa61cd52d0a4",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/e/elfutils-default-yama-scope-0.189-3.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/e/elfutils-default-yama-scope-0.189-3.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/e/elfutils-default-yama-scope-0.189-3.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/e/elfutils-default-yama-scope-0.189-3.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/e/elfutils-default-yama-scope-0.189-3.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "elfutils-libelf-0__0.189-3.fc37.x86_64",
        sha256 = "3e77093e7641f4879554287cebf692f1b5ecb44523e108e0a6da3d454b25208e",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/e/elfutils-libelf-0.189-3.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/e/elfutils-libelf-0.189-3.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/e/elfutils-libelf-0.189-3.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/e/elfutils-libelf-0.189-3.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/e/elfutils-libelf-0.189-3.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "elfutils-libelf-devel-0__0.189-3.fc37.x86_64",
        sha256 = "fbd48678fdc48235be3e647e5bff7d9653a389586d7ee1acb580a173f021662e",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/e/elfutils-libelf-devel-0.189-3.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/e/elfutils-libelf-devel-0.189-3.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/e/elfutils-libelf-devel-0.189-3.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/e/elfutils-libelf-devel-0.189-3.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/e/elfutils-libelf-devel-0.189-3.fc37.x86_64.rpm",
        ],
    )

    rpm(
        name = "elfutils-libs-0__0.189-3.fc37.x86_64",
        sha256 = "7e35cc9b4ff561984268e7548e445c1b0e9d69c2e2a37f9acc3c41c966de754e",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/e/elfutils-libs-0.189-3.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/e/elfutils-libs-0.189-3.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/e/elfutils-libs-0.189-3.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/e/elfutils-libs-0.189-3.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/e/elfutils-libs-0.189-3.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "expat-0__2.5.0-1.fc37.x86_64",
        sha256 = "0e49c2393e5507bbaa16ededf0176e731e0196dd3230f6371d67be8b919e3429",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/e/expat-2.5.0-1.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/e/expat-2.5.0-1.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/e/expat-2.5.0-1.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/e/expat-2.5.0-1.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/e/expat-2.5.0-1.fc37.x86_64.rpm",
        ],
    )

    rpm(
        name = "fedora-gpg-keys-0__37-2.x86_64",
        sha256 = "47a0fdf0c8d0aecd3d4b2eee160affec5ba0d12b7ac6647b3f12fdef275e9738",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/f/fedora-gpg-keys-37-2.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/f/fedora-gpg-keys-37-2.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/f/fedora-gpg-keys-37-2.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/f/fedora-gpg-keys-37-2.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/f/fedora-gpg-keys-37-2.noarch.rpm",
        ],
    )

    rpm(
        name = "fedora-release-common-0__37-16.x86_64",
        sha256 = "5887ea74e3b3525a31fc0a685e10b8ef0be80afe223a9d327c53a5a3168e36d7",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/f/fedora-release-common-37-16.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/f/fedora-release-common-37-16.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/f/fedora-release-common-37-16.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/f/fedora-release-common-37-16.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/f/fedora-release-common-37-16.noarch.rpm",
        ],
    )

    rpm(
        name = "fedora-release-container-0__37-16.x86_64",
        sha256 = "2321ec7a64f24b616f6fef130a97f257aff81b7068b1ede4f81938395e8bab56",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/f/fedora-release-container-37-16.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/f/fedora-release-container-37-16.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/f/fedora-release-container-37-16.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/f/fedora-release-container-37-16.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/f/fedora-release-container-37-16.noarch.rpm",
        ],
    )
    rpm(
        name = "fedora-release-identity-container-0__37-16.x86_64",
        sha256 = "c70aef6b122b352afa7e5eff5c80d0a3dc5b019bf878be80862b94ee986d06ec",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/f/fedora-release-identity-container-37-16.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/f/fedora-release-identity-container-37-16.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/f/fedora-release-identity-container-37-16.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/f/fedora-release-identity-container-37-16.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/f/fedora-release-identity-container-37-16.noarch.rpm",
        ],
    )

    rpm(
        name = "fedora-repos-0__37-2.x86_64",
        sha256 = "f43a00322ae512135f695e9378eadcb3f8a8314bd4e290ea40c7c576621297f6",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/f/fedora-repos-37-2.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/f/fedora-repos-37-2.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/f/fedora-repos-37-2.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/f/fedora-repos-37-2.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/f/fedora-repos-37-2.noarch.rpm",
        ],
    )
    rpm(
        name = "file-0__5.42-4.fc37.x86_64",
        sha256 = "c84d46c1df7de7f3574733ca5cacca81619fcc296bf566b218dc116929bbbc4a",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/f/file-5.42-4.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/f/file-5.42-4.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/f/file-5.42-4.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/f/file-5.42-4.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/f/file-5.42-4.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "file-libs-0__5.42-4.fc37.x86_64",
        sha256 = "d5923edd7fd2e5f5cd8aeb08b291b160ea06d9bc8221ffd146111ff6a9982950",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/f/file-libs-5.42-4.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/f/file-libs-5.42-4.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/f/file-libs-5.42-4.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/f/file-libs-5.42-4.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/f/file-libs-5.42-4.fc37.x86_64.rpm",
        ],
    )

    rpm(
        name = "filesystem-0__3.18-2.fc37.x86_64",
        sha256 = "1c28f722e7f3e48dba7ebf4f763ebebc6688b9e0fd58b55ba4fcd884c8180ef4",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/f/filesystem-3.18-2.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/f/filesystem-3.18-2.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/f/filesystem-3.18-2.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/f/filesystem-3.18-2.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/f/filesystem-3.18-2.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "findutils-1__4.9.0-2.fc37.x86_64",
        sha256 = "25cd555f1a70138b3e81ede1cd375cb620e7a3de05680c9ebaa764f1261d0ce3",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/f/findutils-4.9.0-2.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/f/findutils-4.9.0-2.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/f/findutils-4.9.0-2.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/f/findutils-4.9.0-2.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/f/findutils-4.9.0-2.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "flex-0__2.6.4-11.fc37.x86_64",
        sha256 = "7d163eb50d9166fc24c0cc3dacd43ce59b6c524ad58b600eec152b0e604e52c8",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/f/flex-2.6.4-11.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/f/flex-2.6.4-11.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/f/flex-2.6.4-11.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/f/flex-2.6.4-11.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/f/flex-2.6.4-11.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "fonts-srpm-macros-1__2.0.5-9.fc37.x86_64",
        sha256 = "c2eb9a3d0f01a6b5f21fbfd5f84d5185379925bb9e7df2d66220c16288efd83c",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/f/fonts-srpm-macros-2.0.5-9.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/f/fonts-srpm-macros-2.0.5-9.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/f/fonts-srpm-macros-2.0.5-9.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/f/fonts-srpm-macros-2.0.5-9.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/f/fonts-srpm-macros-2.0.5-9.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "fpc-srpm-macros-0__1.3-6.fc37.x86_64",
        sha256 = "712529a76c473868f64cebf808ec69565fb3edabb440f2e96e2cde83fb79821a",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/f/fpc-srpm-macros-1.3-6.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/f/fpc-srpm-macros-1.3-6.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/f/fpc-srpm-macros-1.3-6.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/f/fpc-srpm-macros-1.3-6.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/f/fpc-srpm-macros-1.3-6.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "fuse3-libs-0__3.10.5-5.fc37.x86_64",
        sha256 = "559626f87751e9b10db9237e0bf05589081f69a5436ee88344352cbb5c4ef7cf",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/f/fuse3-libs-3.10.5-5.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/f/fuse3-libs-3.10.5-5.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/f/fuse3-libs-3.10.5-5.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/f/fuse3-libs-3.10.5-5.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/f/fuse3-libs-3.10.5-5.fc37.x86_64.rpm",
        ],
    )

    rpm(
        name = "gawk-0__5.1.1-4.fc37.x86_64",
        sha256 = "6caea2f79e9fadf96e6cd55eac3f8625137b12f6a2ca75fb5e36b453dfe54edd",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/g/gawk-5.1.1-4.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/g/gawk-5.1.1-4.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/g/gawk-5.1.1-4.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/g/gawk-5.1.1-4.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/g/gawk-5.1.1-4.fc37.x86_64.rpm",
        ],
    )

    rpm(
        name = "gc-0__8.0.6-4.fc37.x86_64",
        sha256 = "2dc8d164a0180af6981c5b4f27c97188f4ba9b1ba31920fe17eab0b646273152",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/g/gc-8.0.6-4.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/g/gc-8.0.6-4.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/g/gc-8.0.6-4.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/g/gc-8.0.6-4.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/g/gc-8.0.6-4.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "gcc-0__12.3.1-1.fc37.x86_64",
        sha256 = "c6c5669d75ff4994f810a6750df55936e4a298ae00a5c0272c341297d5887c90",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/g/gcc-12.3.1-1.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/g/gcc-12.3.1-1.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/g/gcc-12.3.1-1.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/g/gcc-12.3.1-1.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/g/gcc-12.3.1-1.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "gcc-c__plus____plus__-0__12.3.1-1.fc37.x86_64",
        sha256 = "426d546e0c346f6803d3f45056c8aa264e2fbb1b62eeb9f95bf1166706a12e96",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/g/gcc-c++-12.3.1-1.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/g/gcc-c++-12.3.1-1.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/g/gcc-c++-12.3.1-1.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/g/gcc-c++-12.3.1-1.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/g/gcc-c++-12.3.1-1.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "gdbm-libs-1__1.23-2.fc37.x86_64",
        sha256 = "32ab362365afcf96144ba3e65c461cf6f8d495651d0c99fb4eeb970fc2b838e5",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/g/gdbm-libs-1.23-2.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/g/gdbm-libs-1.23-2.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/g/gdbm-libs-1.23-2.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/g/gdbm-libs-1.23-2.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/g/gdbm-libs-1.23-2.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "ghc-srpm-macros-0__1.6.1-1.fc37.x86_64",
        sha256 = "de2385311dcb51dd1d1fc9162ee04efbbbe3add24f8736a929f6834b3fc119b6",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/g/ghc-srpm-macros-1.6.1-1.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/g/ghc-srpm-macros-1.6.1-1.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/g/ghc-srpm-macros-1.6.1-1.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/g/ghc-srpm-macros-1.6.1-1.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/g/ghc-srpm-macros-1.6.1-1.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "glib2-0__2.74.7-2.fc37.x86_64",
        sha256 = "5aa73e300b5191bf90a88f6bffcf03df299e37e0dab90b7311d6cae6ffd5395b",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/g/glib2-2.74.7-2.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/g/glib2-2.74.7-2.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/g/glib2-2.74.7-2.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/g/glib2-2.74.7-2.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/g/glib2-2.74.7-2.fc37.x86_64.rpm",
        ],
    )

    rpm(
        name = "glibc-0__2.36-11.fc37.x86_64",
        sha256 = "ee4e34ebf21e27e40f6979f688d2b5978504900f424b9a4051112d5ef5de2031",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/g/glibc-2.36-11.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/g/glibc-2.36-11.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/g/glibc-2.36-11.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/g/glibc-2.36-11.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/g/glibc-2.36-11.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "glibc-common-0__2.36-11.fc37.x86_64",
        sha256 = "dcc9ee24112d6933040e5ca60255e3d692cb439d46da9f6ba2dddf870d7dadcd",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/g/glibc-common-2.36-11.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/g/glibc-common-2.36-11.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/g/glibc-common-2.36-11.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/g/glibc-common-2.36-11.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/g/glibc-common-2.36-11.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "glibc-devel-0__2.36-11.fc37.x86_64",
        sha256 = "f7783f44a0df4ca6b918bff95f5bf56a76483a2461288ac918da55cd63391e27",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/g/glibc-devel-2.36-11.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/g/glibc-devel-2.36-11.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/g/glibc-devel-2.36-11.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/g/glibc-devel-2.36-11.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/g/glibc-devel-2.36-11.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "glibc-headers-x86-0__2.36-11.fc37.x86_64",
        sha256 = "00b28ee6ed5b7850a0527af69650f59d8f92d9abb6dea09c6df6751f8afc4b90",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/g/glibc-headers-x86-2.36-11.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/g/glibc-headers-x86-2.36-11.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/g/glibc-headers-x86-2.36-11.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/g/glibc-headers-x86-2.36-11.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/g/glibc-headers-x86-2.36-11.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "glibc-langpack-en-0__2.36-11.fc37.x86_64",
        sha256 = "75e2d009b8e1572811662a0663f26695e1cb071cf182d0a8070ab46aed7393cb",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/g/glibc-langpack-en-2.36-11.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/g/glibc-langpack-en-2.36-11.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/g/glibc-langpack-en-2.36-11.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/g/glibc-langpack-en-2.36-11.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/g/glibc-langpack-en-2.36-11.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "glibc-static-0__2.36-11.fc37.x86_64",
        sha256 = "307ec5e6764dcca4ab74073969859670408be901e40318f5e28598af12bee456",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/g/glibc-static-2.36-11.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/g/glibc-static-2.36-11.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/g/glibc-static-2.36-11.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/g/glibc-static-2.36-11.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/g/glibc-static-2.36-11.fc37.x86_64.rpm",
        ],
    )

    rpm(
        name = "gmp-1__6.2.1-3.fc37.x86_64",
        sha256 = "42c8a66f1efcdffaf611e70395e16311f6c56ef795ee2a43c2a48c55eef77734",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/g/gmp-6.2.1-3.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/g/gmp-6.2.1-3.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/g/gmp-6.2.1-3.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/g/gmp-6.2.1-3.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/g/gmp-6.2.1-3.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "gnat-srpm-macros-0__5-1.fc37.x86_64",
        sha256 = "7edbc89f82f45b35291f2abf3583b5a62298b53b6caf4f97de9ebd02d3af4729",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/g/gnat-srpm-macros-5-1.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/g/gnat-srpm-macros-5-1.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/g/gnat-srpm-macros-5-1.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/g/gnat-srpm-macros-5-1.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/g/gnat-srpm-macros-5-1.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "gnupg2-0__2.3.8-1.fc37.x86_64",
        sha256 = "87230a42e847ee21330d2bff385ed78031c20aa6c74463350b32e3b28122e331",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/g/gnupg2-2.3.8-1.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/g/gnupg2-2.3.8-1.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/g/gnupg2-2.3.8-1.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/g/gnupg2-2.3.8-1.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/g/gnupg2-2.3.8-1.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "gnutls-0__3.8.1-1.fc37.x86_64",
        sha256 = "43e3ce719dcae9ab7dbbfd1a8299f4b80d27ee86e64573794e39bf4469ca40de",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/g/gnutls-3.8.1-1.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/g/gnutls-3.8.1-1.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/g/gnutls-3.8.1-1.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/g/gnutls-3.8.1-1.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/g/gnutls-3.8.1-1.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "gnutls-dane-0__3.8.1-1.fc37.x86_64",
        sha256 = "617d0468ec726306037a0e8ab405702495adeb4f0c6e38b8c4209db2dfff5004",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/g/gnutls-dane-3.8.1-1.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/g/gnutls-dane-3.8.1-1.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/g/gnutls-dane-3.8.1-1.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/g/gnutls-dane-3.8.1-1.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/g/gnutls-dane-3.8.1-1.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "gnutls-utils-0__3.8.1-1.fc37.x86_64",
        sha256 = "a3ff7ad783c35942d952923eaab0e2c29bfcd903d97d6437590d3472b0d5ba68",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/g/gnutls-utils-3.8.1-1.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/g/gnutls-utils-3.8.1-1.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/g/gnutls-utils-3.8.1-1.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/g/gnutls-utils-3.8.1-1.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/g/gnutls-utils-3.8.1-1.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "go-srpm-macros-0__3.2.0-1.fc37.x86_64",
        sha256 = "aa3c1ee081411f844d22f293cf0fdefbb51e95b54987279aa6c3155aff14207d",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/g/go-srpm-macros-3.2.0-1.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/g/go-srpm-macros-3.2.0-1.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/g/go-srpm-macros-3.2.0-1.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/g/go-srpm-macros-3.2.0-1.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/g/go-srpm-macros-3.2.0-1.fc37.noarch.rpm",
        ],
    )

    rpm(
        name = "grep-0__3.7-4.fc37.x86_64",
        sha256 = "d997786e71f2c7b4a9ed1323b8684ec1802e49a866fb0c1b69101531440cb464",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/g/grep-3.7-4.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/g/grep-3.7-4.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/g/grep-3.7-4.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/g/grep-3.7-4.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/g/grep-3.7-4.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "groff-base-0__1.22.4-10.fc37.x86_64",
        sha256 = "b5b4e759d1c56188fb777926de0d17498c25d3234d2635ce5a8e7b000bfaf7f3",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/g/groff-base-1.22.4-10.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/g/groff-base-1.22.4-10.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/g/groff-base-1.22.4-10.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/g/groff-base-1.22.4-10.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/g/groff-base-1.22.4-10.fc37.x86_64.rpm",
        ],
    )

    rpm(
        name = "guile22-0__2.2.7-6.fc37.x86_64",
        sha256 = "03227ea6ccc2d0dc553d4ed4b66b3fcd0b8f626d0a81cfe965b5ef39c26de059",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/g/guile22-2.2.7-6.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/g/guile22-2.2.7-6.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/g/guile22-2.2.7-6.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/g/guile22-2.2.7-6.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/g/guile22-2.2.7-6.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "gzip-0__1.12-2.fc37.x86_64",
        sha256 = "3ef9e1b938dd19c5268004e370d90f8a8ae0dbc664715457a371ce900ee7736c",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/g/gzip-1.12-2.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/g/gzip-1.12-2.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/g/gzip-1.12-2.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/g/gzip-1.12-2.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/g/gzip-1.12-2.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "ipxe-roms-qemu-0__20220210-2.git64113751.fc37.x86_64",
        sha256 = "541c4cdbb73deecd93b68ef8a5dc50b8bc32ccd0c9f9480c2a35f318dda38f2e",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/i/ipxe-roms-qemu-20220210-2.git64113751.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/i/ipxe-roms-qemu-20220210-2.git64113751.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/i/ipxe-roms-qemu-20220210-2.git64113751.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/i/ipxe-roms-qemu-20220210-2.git64113751.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/i/ipxe-roms-qemu-20220210-2.git64113751.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "json-glib-0__1.6.6-3.fc37.x86_64",
        sha256 = "598ddb966e5ac4f664f2cbe9d0087b62c1a4067a1f1af24f290d7f681956e29c",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/j/json-glib-1.6.6-3.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/j/json-glib-1.6.6-3.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/j/json-glib-1.6.6-3.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/j/json-glib-1.6.6-3.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/j/json-glib-1.6.6-3.fc37.x86_64.rpm",
        ],
    )

    rpm(
        name = "kernel-headers-0__6.4.4-100.fc37.x86_64",
        sha256 = "dbf0a51df7a1daec64506df6c10a72ff7dcc2adc93c0094f8c53547eab1defd8",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/k/kernel-headers-6.4.4-100.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/k/kernel-headers-6.4.4-100.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/k/kernel-headers-6.4.4-100.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/k/kernel-headers-6.4.4-100.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/k/kernel-headers-6.4.4-100.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "kernel-srpm-macros-0__1.0-15.fc37.x86_64",
        sha256 = "54b2d7f1670ee1c3a6786351b5b2b992715ffd298f8b63cc2f93b3fb8ec1f437",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/k/kernel-srpm-macros-1.0-15.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/k/kernel-srpm-macros-1.0-15.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/k/kernel-srpm-macros-1.0-15.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/k/kernel-srpm-macros-1.0-15.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/k/kernel-srpm-macros-1.0-15.fc37.noarch.rpm",
        ],
    )

    rpm(
        name = "keyutils-libs-0__1.6.1-5.fc37.x86_64",
        sha256 = "e3fd19c3020e55d80b8a24edb68506d2adbb07b2db29eecbde91facae1cca59d",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/k/keyutils-libs-1.6.1-5.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/k/keyutils-libs-1.6.1-5.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/k/keyutils-libs-1.6.1-5.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/k/keyutils-libs-1.6.1-5.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/k/keyutils-libs-1.6.1-5.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "kmod-libs-0__30-2.fc37.x86_64",
        sha256 = "73a1a0f041819c1d50501a699945f0121a3b6e1f54df40cd0bf8f94b1b261ef5",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/k/kmod-libs-30-2.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/k/kmod-libs-30-2.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/k/kmod-libs-30-2.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/k/kmod-libs-30-2.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/k/kmod-libs-30-2.fc37.x86_64.rpm",
        ],
    )

    rpm(
        name = "krb5-libs-0__1.19.2-13.fc37.x86_64",
        sha256 = "5f2ffaa4084cb8918d3990ef352dbfdd9ac28d30c2ed2693c1011641199bb369",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/k/krb5-libs-1.19.2-13.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/k/krb5-libs-1.19.2-13.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/k/krb5-libs-1.19.2-13.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/k/krb5-libs-1.19.2-13.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/k/krb5-libs-1.19.2-13.fc37.x86_64.rpm",
        ],
    )

    rpm(
        name = "libacl-0__2.3.1-4.fc37.x86_64",
        sha256 = "15224cb92199b8011fe47dc12e0bbcdbee0c93e0f29553b3b07ae41768b48ce3",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libacl-2.3.1-4.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libacl-2.3.1-4.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libacl-2.3.1-4.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libacl-2.3.1-4.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libacl-2.3.1-4.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "libaio-0__0.3.111-14.fc37.x86_64",
        sha256 = "d4c8bb3a8bb0c529f49ee7fe6c2100674de6b54837aa29bf0a12e08f08575fdd",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libaio-0.3.111-14.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libaio-0.3.111-14.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libaio-0.3.111-14.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libaio-0.3.111-14.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libaio-0.3.111-14.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "libarchive-0__3.6.1-3.fc37.x86_64",
        sha256 = "a21c75bf1af2f299b06879592d9eb89a20168d3c5306365438e6403e1d1064ce",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/l/libarchive-3.6.1-3.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/l/libarchive-3.6.1-3.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/l/libarchive-3.6.1-3.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/l/libarchive-3.6.1-3.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/l/libarchive-3.6.1-3.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "libassuan-0__2.5.5-5.fc37.x86_64",
        sha256 = "337900b23fc2550547c243e11b7a65284c53c844a3882e0b67e2fbb93f8bf1db",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libassuan-2.5.5-5.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libassuan-2.5.5-5.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libassuan-2.5.5-5.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libassuan-2.5.5-5.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libassuan-2.5.5-5.fc37.x86_64.rpm",
        ],
    )

    rpm(
        name = "libattr-0__2.5.1-5.fc37.x86_64",
        sha256 = "3a423be562953538eaa0d1e78ef35890396cdf1ad89561c619aa72d3a59bfb82",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libattr-2.5.1-5.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libattr-2.5.1-5.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libattr-2.5.1-5.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libattr-2.5.1-5.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libattr-2.5.1-5.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "libb2-0__0.98.1-7.fc37.x86_64",
        sha256 = "da6c0a039fb7e2ce0b324c758757c6482c2683f2ff7bd7f9b06cd625d0fae17a",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libb2-0.98.1-7.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libb2-0.98.1-7.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libb2-0.98.1-7.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libb2-0.98.1-7.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libb2-0.98.1-7.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "libblkid-0__2.38.1-1.fc37.x86_64",
        sha256 = "b0388d1a529bf6b54ca648e91529b1e7790e6aaa42e0ac2b7be6640e4f24a21d",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libblkid-2.38.1-1.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libblkid-2.38.1-1.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libblkid-2.38.1-1.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libblkid-2.38.1-1.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libblkid-2.38.1-1.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "libbpf-2__0.8.0-2.fc37.x86_64",
        sha256 = "3722422d69b3fcfc2d1b0e263051aa94c3fed6b89a709b1f1b4ff6627e114c0a",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libbpf-0.8.0-2.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libbpf-0.8.0-2.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libbpf-0.8.0-2.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libbpf-0.8.0-2.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libbpf-0.8.0-2.fc37.x86_64.rpm",
        ],
    )

    rpm(
        name = "libcap-0__2.48-5.fc37.x86_64",
        sha256 = "aa22373907b6ff9fa3d2f7d9e33a9bdefc9ac50486f2dac5251ac4e206a8a61d",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libcap-2.48-5.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libcap-2.48-5.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libcap-2.48-5.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libcap-2.48-5.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libcap-2.48-5.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "libcap-ng-0__0.8.3-3.fc37.x86_64",
        sha256 = "bcca8a17ae16f9f1c8664f9f54e8f2178f028821f6802ebf33cdcd2d4289bf7f",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libcap-ng-0.8.3-3.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libcap-ng-0.8.3-3.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libcap-ng-0.8.3-3.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libcap-ng-0.8.3-3.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libcap-ng-0.8.3-3.fc37.x86_64.rpm",
        ],
    )

    rpm(
        name = "libcom_err-0__1.46.5-3.fc37.x86_64",
        sha256 = "e98643b3299e5a5b9b1e85a0763b567035f1d83164b3b9a4629fd23467667464",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libcom_err-1.46.5-3.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libcom_err-1.46.5-3.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libcom_err-1.46.5-3.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libcom_err-1.46.5-3.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libcom_err-1.46.5-3.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "libcurl-minimal-0__7.85.0-10.fc37.x86_64",
        sha256 = "999d600757237f43daaac3ac2159c1180044a1703ab700db885d1e992f9e1b0b",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/l/libcurl-minimal-7.85.0-10.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/l/libcurl-minimal-7.85.0-10.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/l/libcurl-minimal-7.85.0-10.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/l/libcurl-minimal-7.85.0-10.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/l/libcurl-minimal-7.85.0-10.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "libdb-0__5.3.28-53.fc37.x86_64",
        sha256 = "e89a4a620d5531f30b895694134a982fa37615b3f61c59a21ede6e64a096c5cd",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libdb-5.3.28-53.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libdb-5.3.28-53.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libdb-5.3.28-53.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libdb-5.3.28-53.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libdb-5.3.28-53.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "libeconf-0__0.5.2-1.fc37.x86_64",
        sha256 = "f15e6136be79a5c3c4a76977ac0c92c1317804de78193f107ede045c05735182",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/l/libeconf-0.5.2-1.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/l/libeconf-0.5.2-1.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/l/libeconf-0.5.2-1.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/l/libeconf-0.5.2-1.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/l/libeconf-0.5.2-1.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "libedit-0__3.1-43.20221009cvs.fc37.x86_64",
        sha256 = "7e128e732af0a53585a9cdae8975c423f3079f57ada4374b4d797ef51cae9ce7",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/l/libedit-3.1-43.20221009cvs.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/l/libedit-3.1-43.20221009cvs.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/l/libedit-3.1-43.20221009cvs.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/l/libedit-3.1-43.20221009cvs.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/l/libedit-3.1-43.20221009cvs.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "libevent-0__2.1.12-7.fc37.x86_64",
        sha256 = "eac9405b6177c4778d772b61ef03a5cd571e2ce6ea337929a1e8a10e80422ba7",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libevent-2.1.12-7.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libevent-2.1.12-7.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libevent-2.1.12-7.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libevent-2.1.12-7.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libevent-2.1.12-7.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "libfdisk-0__2.38.1-1.fc37.x86_64",
        sha256 = "7a4bd1f4975a52fc201c9bc978f155dcb97212cb970210525d903b03644a713d",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libfdisk-2.38.1-1.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libfdisk-2.38.1-1.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libfdisk-2.38.1-1.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libfdisk-2.38.1-1.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libfdisk-2.38.1-1.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "libfdt-0__1.6.1-5.fc37.x86_64",
        sha256 = "4a7f47d967d7884439590b5b1a9145dc68e15a201f648b57e3b5f38ada09ea9c",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libfdt-1.6.1-5.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libfdt-1.6.1-5.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libfdt-1.6.1-5.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libfdt-1.6.1-5.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libfdt-1.6.1-5.fc37.x86_64.rpm",
        ],
    )

    rpm(
        name = "libffi-0__3.4.4-1.fc37.x86_64",
        sha256 = "66bae5662d9287e769f5d8b7f723d45eb19f2902d912be40bf9e5dd8d5c68067",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/l/libffi-3.4.4-1.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/l/libffi-3.4.4-1.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/l/libffi-3.4.4-1.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/l/libffi-3.4.4-1.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/l/libffi-3.4.4-1.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "libgcc-0__12.3.1-1.fc37.x86_64",
        sha256 = "527ee4ac72cf16db23bc5510ce37ff744be2ea3946f8a2982dd3481e491d2a1a",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/l/libgcc-12.3.1-1.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/l/libgcc-12.3.1-1.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/l/libgcc-12.3.1-1.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/l/libgcc-12.3.1-1.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/l/libgcc-12.3.1-1.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "libgcrypt-0__1.10.1-4.fc37.x86_64",
        sha256 = "ca802ad5d10b2728ba10bf98bb16796585d69ec775f5452b3a43718e07c4667a",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libgcrypt-1.10.1-4.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libgcrypt-1.10.1-4.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libgcrypt-1.10.1-4.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libgcrypt-1.10.1-4.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libgcrypt-1.10.1-4.fc37.x86_64.rpm",
        ],
    )

    rpm(
        name = "libgomp-0__12.3.1-1.fc37.x86_64",
        sha256 = "10b0f71a8db70c7a9c456c01cd3154a7ab329ebf99e728994eee1646f32f92d2",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/l/libgomp-12.3.1-1.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/l/libgomp-12.3.1-1.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/l/libgomp-12.3.1-1.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/l/libgomp-12.3.1-1.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/l/libgomp-12.3.1-1.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "libgpg-error-0__1.46-1.fc37.x86_64",
        sha256 = "bfa65a9946b2547110994855d168e4434313ad26280cb935c19bb88d2af283d2",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/l/libgpg-error-1.46-1.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/l/libgpg-error-1.46-1.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/l/libgpg-error-1.46-1.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/l/libgpg-error-1.46-1.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/l/libgpg-error-1.46-1.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "libibverbs-0__41.0-1.fc37.x86_64",
        sha256 = "58fc922b01b99cf99809121ca2d3134853a5cc06ec5b8b5f6a0de7eec5c12202",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libibverbs-41.0-1.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libibverbs-41.0-1.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libibverbs-41.0-1.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libibverbs-41.0-1.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libibverbs-41.0-1.fc37.x86_64.rpm",
        ],
    )

    rpm(
        name = "libidn2-0__2.3.4-1.fc37.x86_64",
        sha256 = "e32e2ab71cfb0bedb84611251987db7acdf665917864be335d0786ea6bbd02b4",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/l/libidn2-2.3.4-1.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/l/libidn2-2.3.4-1.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/l/libidn2-2.3.4-1.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/l/libidn2-2.3.4-1.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/l/libidn2-2.3.4-1.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "libjpeg-turbo-0__2.1.3-2.fc37.x86_64",
        sha256 = "a7934e081a697ca28f8ff83b973b33299c7127cdec8d4102128ab9ea69d172f4",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libjpeg-turbo-2.1.3-2.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libjpeg-turbo-2.1.3-2.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libjpeg-turbo-2.1.3-2.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libjpeg-turbo-2.1.3-2.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libjpeg-turbo-2.1.3-2.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "libksba-0__1.6.3-1.fc37.x86_64",
        sha256 = "2301ca3ff0df8a51f1a872b1c68bdb7defedd099271ecb2f00ed1ff2fb347574",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/l/libksba-1.6.3-1.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/l/libksba-1.6.3-1.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/l/libksba-1.6.3-1.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/l/libksba-1.6.3-1.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/l/libksba-1.6.3-1.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "libmount-0__2.38.1-1.fc37.x86_64",
        sha256 = "50c304faa94d7959e5cbc0642b3c77539ad000042e6617ea5da4789c8105496f",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libmount-2.38.1-1.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libmount-2.38.1-1.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libmount-2.38.1-1.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libmount-2.38.1-1.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libmount-2.38.1-1.fc37.x86_64.rpm",
        ],
    )

    rpm(
        name = "libmpc-0__1.2.1-5.fc37.x86_64",
        sha256 = "4f4a872b3d3e322e05ec3cd3e52f4d5cb06604126fd5400757b60aca55913d20",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libmpc-1.2.1-5.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libmpc-1.2.1-5.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libmpc-1.2.1-5.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libmpc-1.2.1-5.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libmpc-1.2.1-5.fc37.x86_64.rpm",
        ],
    )

    rpm(
        name = "libnghttp2-0__1.51.0-1.fc37.x86_64",
        sha256 = "42fbaaacbeb241755d8448dd5672bbbcc48cbe9548c095ce0efef4140bc12520",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/l/libnghttp2-1.51.0-1.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/l/libnghttp2-1.51.0-1.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/l/libnghttp2-1.51.0-1.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/l/libnghttp2-1.51.0-1.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/l/libnghttp2-1.51.0-1.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "libnl3-0__3.7.0-2.fc37.x86_64",
        sha256 = "4543c991e6f536468d9d47527a201b58b9bc049364a6bdfe15a2f910a02e68f6",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libnl3-3.7.0-2.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libnl3-3.7.0-2.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libnl3-3.7.0-2.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libnl3-3.7.0-2.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libnl3-3.7.0-2.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "libnsl2-0__2.0.0-4.fc37.x86_64",
        sha256 = "a1e9428515b0df1c2a423ad3c35bcdf93333172fe346169bb3018a882e27be5f",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libnsl2-2.0.0-4.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libnsl2-2.0.0-4.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libnsl2-2.0.0-4.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libnsl2-2.0.0-4.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libnsl2-2.0.0-4.fc37.x86_64.rpm",
        ],
    )

    rpm(
        name = "libpkgconf-0__1.8.0-3.fc37.x86_64",
        sha256 = "ecd52fd3f3065606ba5164249b29c837cbd172643d13a00a1a72fc657b115af7",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libpkgconf-1.8.0-3.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libpkgconf-1.8.0-3.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libpkgconf-1.8.0-3.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libpkgconf-1.8.0-3.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libpkgconf-1.8.0-3.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "libpmem-0__1.12.0-1.fc37.x86_64",
        sha256 = "af3c9045089110c849ee3191aa506e16f4a69dc625b77abf09583de237df41df",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libpmem-1.12.0-1.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libpmem-1.12.0-1.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libpmem-1.12.0-1.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libpmem-1.12.0-1.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libpmem-1.12.0-1.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "libpng-2__1.6.37-13.fc37.x86_64",
        sha256 = "49a024d34e3c531516562bc51b749dee540db4d34486a95cdd8d85300b7de455",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libpng-1.6.37-13.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libpng-1.6.37-13.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libpng-1.6.37-13.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libpng-1.6.37-13.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libpng-1.6.37-13.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "libpwquality-0__1.4.5-3.fc37.x86_64",
        sha256 = "a9019a471496fdada529757331ec004397db7a0c4347531bd639c127bbaf8300",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/l/libpwquality-1.4.5-3.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/l/libpwquality-1.4.5-3.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/l/libpwquality-1.4.5-3.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/l/libpwquality-1.4.5-3.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/l/libpwquality-1.4.5-3.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "librdmacm-0__41.0-1.fc37.x86_64",
        sha256 = "65edde85818fd605392a74c235e0be8d9bb7032b8bdccd89d3fe34ae2e4a1e7b",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/librdmacm-41.0-1.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/librdmacm-41.0-1.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/librdmacm-41.0-1.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/librdmacm-41.0-1.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/librdmacm-41.0-1.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "libseccomp-0__2.5.3-3.fc37.x86_64",
        sha256 = "017877a97c8222fc7eca7fab77600a3a1fcdec92f9dd39d8df6e64726909fcbe",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libseccomp-2.5.3-3.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libseccomp-2.5.3-3.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libseccomp-2.5.3-3.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libseccomp-2.5.3-3.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libseccomp-2.5.3-3.fc37.x86_64.rpm",
        ],
    )

    rpm(
        name = "libselinux-0__3.5-1.fc37.x86_64",
        sha256 = "43d73a574c3c0838d213c4d5f038766d41e4eb6930c68b09db53d65c30c2de1d",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/l/libselinux-3.5-1.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/l/libselinux-3.5-1.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/l/libselinux-3.5-1.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/l/libselinux-3.5-1.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/l/libselinux-3.5-1.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "libselinux-utils-0__3.5-1.fc37.x86_64",
        sha256 = "723efbdc421150c13f6a2fe47e3d2587f83a26bfae8561e3361985793762b05d",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/l/libselinux-utils-3.5-1.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/l/libselinux-utils-3.5-1.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/l/libselinux-utils-3.5-1.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/l/libselinux-utils-3.5-1.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/l/libselinux-utils-3.5-1.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "libsemanage-0__3.5-2.fc37.x86_64",
        sha256 = "39a7db1792d9b17dd8a4339cef7037e32c100dce4dab242207ca86efffc4b61c",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/l/libsemanage-3.5-2.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/l/libsemanage-3.5-2.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/l/libsemanage-3.5-2.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/l/libsemanage-3.5-2.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/l/libsemanage-3.5-2.fc37.x86_64.rpm",
        ],
    )

    rpm(
        name = "libsepol-0__3.5-1.fc37.x86_64",
        sha256 = "2cdfb41068ac6e211652b3e2ed88c16d606e7374f9f52dfd1248981101501299",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/l/libsepol-3.5-1.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/l/libsepol-3.5-1.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/l/libsepol-3.5-1.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/l/libsepol-3.5-1.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/l/libsepol-3.5-1.fc37.x86_64.rpm",
        ],
    )

    rpm(
        name = "libsigsegv-0__2.14-3.fc37.x86_64",
        sha256 = "0f038b70d155dae3df4824776c5a135f02c423c688b9486d4f84eb6a16a90494",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libsigsegv-2.14-3.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libsigsegv-2.14-3.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libsigsegv-2.14-3.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libsigsegv-2.14-3.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libsigsegv-2.14-3.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "libslirp-0__4.7.0-2.fc37.x86_64",
        sha256 = "775c3da0e9e2961262c552292f98dc80f4b05bde52573539148ff4cb51459a51",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libslirp-4.7.0-2.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libslirp-4.7.0-2.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libslirp-4.7.0-2.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libslirp-4.7.0-2.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libslirp-4.7.0-2.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "libsmartcols-0__2.38.1-1.fc37.x86_64",
        sha256 = "93246c002aefec27bb398aa3397ae555bcc3035b10aebb4937c4bea9268bacf1",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libsmartcols-2.38.1-1.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libsmartcols-2.38.1-1.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libsmartcols-2.38.1-1.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libsmartcols-2.38.1-1.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libsmartcols-2.38.1-1.fc37.x86_64.rpm",
        ],
    )

    rpm(
        name = "libstdc__plus____plus__-0__12.3.1-1.fc37.x86_64",
        sha256 = "cc99d77eec7b347e6867391d306c1641393ef6470ba6dcf4c087c78ce26443ff",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/l/libstdc++-12.3.1-1.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/l/libstdc++-12.3.1-1.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/l/libstdc++-12.3.1-1.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/l/libstdc++-12.3.1-1.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/l/libstdc++-12.3.1-1.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "libstdc__plus____plus__-devel-0__12.3.1-1.fc37.x86_64",
        sha256 = "331c9d83489e05966b1f6f43899afcbd893599bbe7a28e5c3abb1e6760c212a7",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/l/libstdc++-devel-12.3.1-1.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/l/libstdc++-devel-12.3.1-1.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/l/libstdc++-devel-12.3.1-1.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/l/libstdc++-devel-12.3.1-1.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/l/libstdc++-devel-12.3.1-1.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "libstdc__plus____plus__-static-0__12.3.1-1.fc37.x86_64",
        sha256 = "dcc071288b56427543c7a9eb23bb00155564aee085d51510e6b09e46f8cc7fe2",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/l/libstdc++-static-12.3.1-1.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/l/libstdc++-static-12.3.1-1.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/l/libstdc++-static-12.3.1-1.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/l/libstdc++-static-12.3.1-1.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/l/libstdc++-static-12.3.1-1.fc37.x86_64.rpm",
        ],
    )

    rpm(
        name = "libtasn1-0__4.19.0-1.fc37.x86_64",
        sha256 = "35b51a0796af6930b2a8a511df8c51938006cfcfdf74ddfe6482eb9febd87dfa",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/l/libtasn1-4.19.0-1.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/l/libtasn1-4.19.0-1.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/l/libtasn1-4.19.0-1.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/l/libtasn1-4.19.0-1.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/l/libtasn1-4.19.0-1.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "libtirpc-0__1.3.3-1.rc1.fc37.x86_64",
        sha256 = "55a12c31a1f8650f0c517411fff30301a8644a8c3a3505667846e5b59a4de5c8",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/l/libtirpc-1.3.3-1.rc1.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/l/libtirpc-1.3.3-1.rc1.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/l/libtirpc-1.3.3-1.rc1.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/l/libtirpc-1.3.3-1.rc1.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/l/libtirpc-1.3.3-1.rc1.fc37.x86_64.rpm",
        ],
    )

    rpm(
        name = "libtool-ltdl-0__2.4.7-2.fc37.x86_64",
        sha256 = "73b1b4a028077983bb0643a40ddf34a29731c49b0994b29ac77e3bcc4243bafe",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libtool-ltdl-2.4.7-2.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libtool-ltdl-2.4.7-2.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libtool-ltdl-2.4.7-2.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libtool-ltdl-2.4.7-2.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libtool-ltdl-2.4.7-2.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "libtpms-0__0.9.6-1.fc37.x86_64",
        sha256 = "d61fd47b4126e4d89d425af91a224718493b7c7c4eab148508699d577d191dbe",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/l/libtpms-0.9.6-1.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/l/libtpms-0.9.6-1.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/l/libtpms-0.9.6-1.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/l/libtpms-0.9.6-1.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/l/libtpms-0.9.6-1.fc37.x86_64.rpm",
        ],
    )

    rpm(
        name = "libunistring-0__1.0-2.fc37.x86_64",
        sha256 = "acb031577655bba5a41c1fb0ec954bb84e207f9e2d08b2cdb3d4e2b7806b0670",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libunistring-1.0-2.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libunistring-1.0-2.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libunistring-1.0-2.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libunistring-1.0-2.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libunistring-1.0-2.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "liburing-0__2.3-1.fc37.x86_64",
        sha256 = "b7ace7323804d803c8b8c33fc699b9346924fb80841a701471f29e29abeb255b",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/l/liburing-2.3-1.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/l/liburing-2.3-1.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/l/liburing-2.3-1.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/l/liburing-2.3-1.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/l/liburing-2.3-1.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "libutempter-0__1.2.1-7.fc37.x86_64",
        sha256 = "8fc30b0742e939954d6aebd45364dcd1dbb8b9c85e75c799301c3507e22ea56a",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libutempter-1.2.1-7.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libutempter-1.2.1-7.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libutempter-1.2.1-7.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libutempter-1.2.1-7.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libutempter-1.2.1-7.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "libuuid-0__2.38.1-1.fc37.x86_64",
        sha256 = "b054577d98aa9615fe459abec31be46b19ad72e0da620d8d251b4449a6db020d",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libuuid-2.38.1-1.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libuuid-2.38.1-1.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libuuid-2.38.1-1.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libuuid-2.38.1-1.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libuuid-2.38.1-1.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "libuuid-devel-0__2.38.1-1.fc37.x86_64",
        sha256 = "a42450ad26785144969fd5faab10f6a382d13e3db2e7b130a33e8a3b314d5d3f",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libuuid-devel-2.38.1-1.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libuuid-devel-2.38.1-1.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libuuid-devel-2.38.1-1.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libuuid-devel-2.38.1-1.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libuuid-devel-2.38.1-1.fc37.x86_64.rpm",
        ],
    )

    rpm(
        name = "libverto-0__0.3.2-4.fc37.x86_64",
        sha256 = "ca47b52e1ecd8a2ac6eda368d985390816fbb447f43135ec0ba105165997817f",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libverto-0.3.2-4.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libverto-0.3.2-4.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libverto-0.3.2-4.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libverto-0.3.2-4.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libverto-0.3.2-4.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "libxcrypt-0__4.4.36-1.fc37.x86_64",
        sha256 = "eadfd203fc0fdb870af0be9113b077378f79737212b17af222bdb27d4535d14b",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/l/libxcrypt-4.4.36-1.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/l/libxcrypt-4.4.36-1.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/l/libxcrypt-4.4.36-1.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/l/libxcrypt-4.4.36-1.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/l/libxcrypt-4.4.36-1.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "libxcrypt-devel-0__4.4.36-1.fc37.x86_64",
        sha256 = "da63cfcaab4acec68941c7e307db1992d3310319e6e9999cc6f8c97f1509c18c",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/l/libxcrypt-devel-4.4.36-1.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/l/libxcrypt-devel-4.4.36-1.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/l/libxcrypt-devel-4.4.36-1.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/l/libxcrypt-devel-4.4.36-1.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/l/libxcrypt-devel-4.4.36-1.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "libxcrypt-static-0__4.4.36-1.fc37.x86_64",
        sha256 = "ab235b8970405cd9a55af703f1197265dfe6460479e2e458e5cad42136263857",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/l/libxcrypt-static-4.4.36-1.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/l/libxcrypt-static-4.4.36-1.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/l/libxcrypt-static-4.4.36-1.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/l/libxcrypt-static-4.4.36-1.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/l/libxcrypt-static-4.4.36-1.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "libxml2-0__2.10.4-1.fc37.x86_64",
        sha256 = "c2997441f2d1e07318a932b8e01d39c69d891c6f931a6f398d9970ac4de3e6c3",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/l/libxml2-2.10.4-1.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/l/libxml2-2.10.4-1.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/l/libxml2-2.10.4-1.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/l/libxml2-2.10.4-1.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/l/libxml2-2.10.4-1.fc37.x86_64.rpm",
        ],
    )

    rpm(
        name = "libzstd-0__1.5.5-1.fc37.x86_64",
        sha256 = "58af9287e3afe709e2f898e679168485a713382c155f03ed94204094cfa8dd37",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/l/libzstd-1.5.5-1.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/l/libzstd-1.5.5-1.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/l/libzstd-1.5.5-1.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/l/libzstd-1.5.5-1.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/l/libzstd-1.5.5-1.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "libzstd-devel-0__1.5.5-1.fc37.x86_64",
        sha256 = "64c933d39e2b43502ab0faa9b25bf2b5c07493c5160d58be9493208199a1be12",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/l/libzstd-devel-1.5.5-1.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/l/libzstd-devel-1.5.5-1.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/l/libzstd-devel-1.5.5-1.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/l/libzstd-devel-1.5.5-1.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/l/libzstd-devel-1.5.5-1.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "lld-0__15.0.7-1.fc37.x86_64",
        sha256 = "df48da95c6b5c23ba9d32456e11a5c396abd7e6ae0e9b8542feb68cc2661dd0e",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/l/lld-15.0.7-1.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/l/lld-15.0.7-1.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/l/lld-15.0.7-1.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/l/lld-15.0.7-1.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/l/lld-15.0.7-1.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "lld-libs-0__15.0.7-1.fc37.x86_64",
        sha256 = "263ab7c8b283706c8a422f2e43fe20fd5617fded230c53eabe9d47a887e50017",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/l/lld-libs-15.0.7-1.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/l/lld-libs-15.0.7-1.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/l/lld-libs-15.0.7-1.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/l/lld-libs-15.0.7-1.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/l/lld-libs-15.0.7-1.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "llvm-0__15.0.7-2.fc37.x86_64",
        sha256 = "37f989804bc5221653741ec99f2da84c7f84fb7e79f612de16f9f271fa0e5dab",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/l/llvm-15.0.7-2.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/l/llvm-15.0.7-2.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/l/llvm-15.0.7-2.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/l/llvm-15.0.7-2.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/l/llvm-15.0.7-2.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "llvm-libs-0__15.0.7-2.fc37.x86_64",
        sha256 = "5bf61243c4372461c44e52301a93ab8ceefc36a02d0c2f8e4ff3193e58a2d96e",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/l/llvm-libs-15.0.7-2.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/l/llvm-libs-15.0.7-2.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/l/llvm-libs-15.0.7-2.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/l/llvm-libs-15.0.7-2.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/l/llvm-libs-15.0.7-2.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "lua-libs-0__5.4.4-9.fc37.x86_64",
        sha256 = "561ebd5154e2d0d56f6a90283065b27304f81fc57fc881faf485e55d6414fad6",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/l/lua-libs-5.4.4-9.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/l/lua-libs-5.4.4-9.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/l/lua-libs-5.4.4-9.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/l/lua-libs-5.4.4-9.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/l/lua-libs-5.4.4-9.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "lua-srpm-macros-0__1-7.fc37.x86_64",
        sha256 = "550e0e26c2a3fb24381cd14c693592eb565711093d230a9531567897aa857373",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/lua-srpm-macros-1-7.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/lua-srpm-macros-1-7.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/lua-srpm-macros-1-7.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/lua-srpm-macros-1-7.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/lua-srpm-macros-1-7.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "lz4-0__1.9.4-1.fc37.x86_64",
        sha256 = "21ad776601dc88d368a937a1c4be8d3f6f377a9af29c7b0a82bad3aa9ffa12cd",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/l/lz4-1.9.4-1.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/l/lz4-1.9.4-1.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/l/lz4-1.9.4-1.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/l/lz4-1.9.4-1.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/l/lz4-1.9.4-1.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "lz4-libs-0__1.9.4-1.fc37.x86_64",
        sha256 = "f39b8b018fcb2b55477cdbfa4af7c9db9b660c85000a4a42e880b1a951efbe5a",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/l/lz4-libs-1.9.4-1.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/l/lz4-libs-1.9.4-1.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/l/lz4-libs-1.9.4-1.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/l/lz4-libs-1.9.4-1.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/l/lz4-libs-1.9.4-1.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "lzo-0__2.10-7.fc37.x86_64",
        sha256 = "fdde3f48dc7d4f5197d79b765c730aedc86632edc5fcbee3b50e876b2cf39e3e",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/lzo-2.10-7.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/lzo-2.10-7.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/lzo-2.10-7.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/lzo-2.10-7.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/lzo-2.10-7.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "m4-0__1.4.19-4.fc37.x86_64",
        sha256 = "bda2327e5de23e4ab95dd18f164249bd3659f630daa387a650e09516628d5d87",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/m/m4-1.4.19-4.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/m/m4-1.4.19-4.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/m/m4-1.4.19-4.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/m/m4-1.4.19-4.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/m/m4-1.4.19-4.fc37.x86_64.rpm",
        ],
    )

    rpm(
        name = "make-1__4.3-11.fc37.x86_64",
        sha256 = "6bdd9ca1daee43a839f1122fc92b1c42312f824deac45abba2772b45c84cd96e",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/m/make-4.3-11.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/m/make-4.3-11.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/m/make-4.3-11.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/m/make-4.3-11.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/m/make-4.3-11.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "mpdecimal-0__2.5.1-4.fc37.x86_64",
        sha256 = "45764a6773175638883e02215074f084de209d172d1d07be289e89aa5f4131d3",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/m/mpdecimal-2.5.1-4.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/m/mpdecimal-2.5.1-4.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/m/mpdecimal-2.5.1-4.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/m/mpdecimal-2.5.1-4.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/m/mpdecimal-2.5.1-4.fc37.x86_64.rpm",
        ],
    )

    rpm(
        name = "mpfr-0__4.1.0-10.fc37.x86_64",
        sha256 = "3be8cf104424fb5e148846a1df4a9c193527f55ee866bff0963e788450483566",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/m/mpfr-4.1.0-10.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/m/mpfr-4.1.0-10.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/m/mpfr-4.1.0-10.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/m/mpfr-4.1.0-10.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/m/mpfr-4.1.0-10.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "nasm-0__2.15.05-3.fc37.x86_64",
        sha256 = "6df570ecff0bff21b2dd2c8b1809b3e2076ff53d5dca2d51af4ed92381a75c3f",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/n/nasm-2.15.05-3.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/n/nasm-2.15.05-3.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/n/nasm-2.15.05-3.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/n/nasm-2.15.05-3.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/n/nasm-2.15.05-3.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "ncurses-0__6.4-3.20230114.fc37.x86_64",
        sha256 = "54497d6c370fa147c3019b27f977e69de523a6e5bdf47ea0b73671a4b85b0012",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/n/ncurses-6.4-3.20230114.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/n/ncurses-6.4-3.20230114.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/n/ncurses-6.4-3.20230114.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/n/ncurses-6.4-3.20230114.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/n/ncurses-6.4-3.20230114.fc37.x86_64.rpm",
        ],
    )

    rpm(
        name = "ncurses-base-0__6.4-3.20230114.fc37.x86_64",
        sha256 = "bce880d3e30b0a01e1b67457454b217d551e0d4ed8b48176e1a97b761f1fed1c",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/n/ncurses-base-6.4-3.20230114.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/n/ncurses-base-6.4-3.20230114.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/n/ncurses-base-6.4-3.20230114.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/n/ncurses-base-6.4-3.20230114.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/n/ncurses-base-6.4-3.20230114.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "ncurses-libs-0__6.4-3.20230114.fc37.x86_64",
        sha256 = "01815f79546f40b42da5b192a0b781b04a4178b8dac608474463815b232e1373",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/n/ncurses-libs-6.4-3.20230114.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/n/ncurses-libs-6.4-3.20230114.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/n/ncurses-libs-6.4-3.20230114.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/n/ncurses-libs-6.4-3.20230114.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/n/ncurses-libs-6.4-3.20230114.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "ndctl-libs-0__78-1.fc37.x86_64",
        sha256 = "6a4a9cbc70b49b5660a0b69ff02de049f7ace6ba519b88bca25259e30c373abf",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/n/ndctl-libs-78-1.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/n/ndctl-libs-78-1.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/n/ndctl-libs-78-1.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/n/ndctl-libs-78-1.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/n/ndctl-libs-78-1.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "nettle-0__3.8-2.fc37.x86_64",
        sha256 = "8fe2d98578b0c4454536faacbaafd66d1754b8439bb6332d7576a741f4c72208",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/n/nettle-3.8-2.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/n/nettle-3.8-2.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/n/nettle-3.8-2.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/n/nettle-3.8-2.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/n/nettle-3.8-2.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "nim-srpm-macros-0__3-7.fc37.x86_64",
        sha256 = "58f4c0ef46d70a980cc68a451a880202688817972b818295df64c477b3511f1a",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/n/nim-srpm-macros-3-7.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/n/nim-srpm-macros-3-7.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/n/nim-srpm-macros-3-7.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/n/nim-srpm-macros-3-7.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/n/nim-srpm-macros-3-7.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "npth-0__1.6-9.fc37.x86_64",
        sha256 = "6bf0c6a9c2a00e9e06988299543a9bb1aec4445014b11216db6d6c2e2253bda9",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/n/npth-1.6-9.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/n/npth-1.6-9.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/n/npth-1.6-9.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/n/npth-1.6-9.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/n/npth-1.6-9.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "numactl-libs-0__2.0.14-6.fc37.x86_64",
        sha256 = "8f2e423d8f64f3abf33f8660df718d69f785a673a57eb188258a9f79af8f678f",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/n/numactl-libs-2.0.14-6.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/n/numactl-libs-2.0.14-6.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/n/numactl-libs-2.0.14-6.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/n/numactl-libs-2.0.14-6.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/n/numactl-libs-2.0.14-6.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "ocaml-srpm-macros-0__7-2.fc37.x86_64",
        sha256 = "5400fb3ba764f5854715af9e502923ba1020f593e98268e5328d41190622ed2d",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/o/ocaml-srpm-macros-7-2.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/o/ocaml-srpm-macros-7-2.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/o/ocaml-srpm-macros-7-2.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/o/ocaml-srpm-macros-7-2.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/o/ocaml-srpm-macros-7-2.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "openblas-srpm-macros-0__2-12.fc37.x86_64",
        sha256 = "ac34d1143ac424db0f5e16fffb39a255ec76f24e30fa13760941a13edf4ff99c",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/o/openblas-srpm-macros-2-12.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/o/openblas-srpm-macros-2-12.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/o/openblas-srpm-macros-2-12.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/o/openblas-srpm-macros-2-12.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/o/openblas-srpm-macros-2-12.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "openldap-0__2.6.4-1.fc37.x86_64",
        sha256 = "613788ec7bdccd9d14f3ffa97b06c32d43857a5ade51dc54d36d83a57007333c",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/o/openldap-2.6.4-1.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/o/openldap-2.6.4-1.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/o/openldap-2.6.4-1.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/o/openldap-2.6.4-1.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/o/openldap-2.6.4-1.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "openssl-devel-1__3.0.9-1.fc37.x86_64",
        sha256 = "721a14663cba596c6962754d63c1f1135c30fc63ca7eed74474e7061c994abc8",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/o/openssl-devel-3.0.9-1.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/o/openssl-devel-3.0.9-1.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/o/openssl-devel-3.0.9-1.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/o/openssl-devel-3.0.9-1.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/o/openssl-devel-3.0.9-1.fc37.x86_64.rpm",
        ],
    )

    rpm(
        name = "openssl-libs-1__3.0.9-1.fc37.x86_64",
        sha256 = "613667ba22b3705e71e673cedd28a1e4f1e1e975acd5df13b3251c40e5b9d4de",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/o/openssl-libs-3.0.9-1.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/o/openssl-libs-3.0.9-1.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/o/openssl-libs-3.0.9-1.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/o/openssl-libs-3.0.9-1.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/o/openssl-libs-3.0.9-1.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "p11-kit-0__0.25.0-1.fc37.x86_64",
        sha256 = "22cb0f9bd16493d0f15c40851a387d8d6b9896291a98752412a0e99af35f22f4",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/p11-kit-0.25.0-1.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/p11-kit-0.25.0-1.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/p11-kit-0.25.0-1.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/p11-kit-0.25.0-1.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/p11-kit-0.25.0-1.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "p11-kit-trust-0__0.25.0-1.fc37.x86_64",
        sha256 = "a735cc2a9f096362cf44ef6e8c317311c1b7a153cf26b9b8a2d4d8a4de06fb44",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/p11-kit-trust-0.25.0-1.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/p11-kit-trust-0.25.0-1.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/p11-kit-trust-0.25.0-1.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/p11-kit-trust-0.25.0-1.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/p11-kit-trust-0.25.0-1.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "package-notes-srpm-macros-0__0.5-7.fc37.x86_64",
        sha256 = "31ce86bd115ddfc662b3248e80216754b5607599e4ab4db1ab3988ac554e6c44",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/package-notes-srpm-macros-0.5-7.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/package-notes-srpm-macros-0.5-7.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/package-notes-srpm-macros-0.5-7.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/package-notes-srpm-macros-0.5-7.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/package-notes-srpm-macros-0.5-7.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "pam-0__1.5.2-14.fc37.x86_64",
        sha256 = "a66ee1c9f9155c97e77cbd18658ce5129638f7d6e208c01c172c4dd1dfdbbe6d",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/pam-1.5.2-14.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/pam-1.5.2-14.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/pam-1.5.2-14.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/pam-1.5.2-14.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/pam-1.5.2-14.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "pam-libs-0__1.5.2-14.fc37.x86_64",
        sha256 = "ee34422adc6451da744bd16a8cd66c9912a822c4e55227c23ff56960c32980f5",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/pam-libs-1.5.2-14.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/pam-libs-1.5.2-14.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/pam-libs-1.5.2-14.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/pam-libs-1.5.2-14.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/pam-libs-1.5.2-14.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "patch-0__2.7.6-17.fc37.x86_64",
        sha256 = "0ed45c95381e8b979b6c6a02356401de363806a893190ddbc57a5ef1bdadc98e",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/patch-2.7.6-17.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/patch-2.7.6-17.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/patch-2.7.6-17.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/patch-2.7.6-17.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/patch-2.7.6-17.fc37.x86_64.rpm",
        ],
    )

    rpm(
        name = "pcre-0__8.45-1.fc37.2.x86_64",
        sha256 = "86a648e3b88f581b15ca2eda6b441be7c5c3810a9eae25ca940c767029e4e923",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/pcre-8.45-1.fc37.2.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/pcre-8.45-1.fc37.2.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/pcre-8.45-1.fc37.2.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/pcre-8.45-1.fc37.2.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/pcre-8.45-1.fc37.2.x86_64.rpm",
        ],
    )

    rpm(
        name = "pcre2-0__10.40-1.fc37.1.x86_64",
        sha256 = "422de947ec1a7aafcd212a51e64257b64d5b0a02808104a33e7c3cd9ef629148",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/pcre2-10.40-1.fc37.1.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/pcre2-10.40-1.fc37.1.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/pcre2-10.40-1.fc37.1.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/pcre2-10.40-1.fc37.1.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/pcre2-10.40-1.fc37.1.x86_64.rpm",
        ],
    )

    rpm(
        name = "pcre2-syntax-0__10.40-1.fc37.1.x86_64",
        sha256 = "585f339942a0bf4b0eab638ddf825544793485cbcb9f1eaee079b9956d90aafa",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/pcre2-syntax-10.40-1.fc37.1.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/pcre2-syntax-10.40-1.fc37.1.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/pcre2-syntax-10.40-1.fc37.1.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/pcre2-syntax-10.40-1.fc37.1.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/pcre2-syntax-10.40-1.fc37.1.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-4__5.36.1-494.fc37.x86_64",
        sha256 = "1f5cd287839c8e262db394c8bc23a461fd9f49aa4778953a3fdcbb129af86ede",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-5.36.1-494.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-5.36.1-494.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-5.36.1-494.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-5.36.1-494.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-5.36.1-494.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "perl-Algorithm-Diff-0__1.2010-7.fc37.x86_64",
        sha256 = "714ad6947e716a4d6a6f6b24c549fb2fa2f37dd0985bd2db0689bc4224d06b3c",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Algorithm-Diff-1.2010-7.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Algorithm-Diff-1.2010-7.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Algorithm-Diff-1.2010-7.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Algorithm-Diff-1.2010-7.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Algorithm-Diff-1.2010-7.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Archive-Tar-0__2.40-490.fc37.x86_64",
        sha256 = "098fa85dff04e11e8c4feb43e54e20330beeaddc528bf19a0ca42c6168ccc878",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Archive-Tar-2.40-490.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Archive-Tar-2.40-490.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Archive-Tar-2.40-490.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Archive-Tar-2.40-490.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Archive-Tar-2.40-490.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Archive-Zip-0__1.68-9.fc37.x86_64",
        sha256 = "edb3407c4a3440f97726e108e17519f10b0ad6b161a5065a5d46f2d862246965",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Archive-Zip-1.68-9.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Archive-Zip-1.68-9.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Archive-Zip-1.68-9.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Archive-Zip-1.68-9.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Archive-Zip-1.68-9.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Attribute-Handlers-0__1.02-494.fc37.x86_64",
        sha256 = "48b003d2b10dba2d8f6f928b019ad5b6e8c5f81d4ebcfe052d50f6ed40fb0140",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Attribute-Handlers-1.02-494.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Attribute-Handlers-1.02-494.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Attribute-Handlers-1.02-494.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Attribute-Handlers-1.02-494.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Attribute-Handlers-1.02-494.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-AutoLoader-0__5.74-494.fc37.x86_64",
        sha256 = "ee935b51277b2ced900b928307f83d753c9a935744ccb127a045e3fa63b2c89c",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-AutoLoader-5.74-494.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-AutoLoader-5.74-494.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-AutoLoader-5.74-494.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-AutoLoader-5.74-494.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-AutoLoader-5.74-494.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-AutoSplit-0__5.74-494.fc37.x86_64",
        sha256 = "65e5098d529b8f0f942b5affcb04f558208bf564e94f6e23b8355fc95c91b948",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-AutoSplit-5.74-494.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-AutoSplit-5.74-494.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-AutoSplit-5.74-494.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-AutoSplit-5.74-494.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-AutoSplit-5.74-494.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-B-0__1.83-494.fc37.x86_64",
        sha256 = "8ef1be3b5b0db787660b6d8a4d9de717887b73027fa16613390f21f5785cc11f",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-B-1.83-494.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-B-1.83-494.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-B-1.83-494.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-B-1.83-494.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-B-1.83-494.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "perl-Benchmark-0__1.23-494.fc37.x86_64",
        sha256 = "818cda6856e4279ea8c719d694a767dd75249a87be1a029e6dea4cad8816ba49",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Benchmark-1.23-494.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Benchmark-1.23-494.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Benchmark-1.23-494.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Benchmark-1.23-494.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Benchmark-1.23-494.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-CPAN-0__2.36-1.fc37.x86_64",
        sha256 = "4f71ca98763d6b5dc751d264388157fee00fca7054cfb058bd1014b1ef49bdd3",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-CPAN-2.36-1.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-CPAN-2.36-1.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-CPAN-2.36-1.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-CPAN-2.36-1.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-CPAN-2.36-1.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-CPAN-Meta-0__2.150010-489.fc37.x86_64",
        sha256 = "81de5eae891aac095402cf76cc2b4f3f52864db88950f2b1bdd43c2a4d9d7333",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-CPAN-Meta-2.150010-489.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-CPAN-Meta-2.150010-489.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-CPAN-Meta-2.150010-489.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-CPAN-Meta-2.150010-489.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-CPAN-Meta-2.150010-489.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-CPAN-Meta-Requirements-0__2.140-490.fc37.x86_64",
        sha256 = "fcf1d32e91e50232a81e8cf9461f523a76a659733ec008850432454298ebbe12",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-CPAN-Meta-Requirements-2.140-490.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-CPAN-Meta-Requirements-2.140-490.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-CPAN-Meta-Requirements-2.140-490.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-CPAN-Meta-Requirements-2.140-490.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-CPAN-Meta-Requirements-2.140-490.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-CPAN-Meta-YAML-0__0.018-490.fc37.x86_64",
        sha256 = "f5e89acb7ad1be213d7f6d9cfc81fcf4fe5e2e765eb8c372e9d784666b916c6b",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-CPAN-Meta-YAML-0.018-490.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-CPAN-Meta-YAML-0.018-490.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-CPAN-Meta-YAML-0.018-490.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-CPAN-Meta-YAML-0.018-490.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-CPAN-Meta-YAML-0.018-490.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Carp-0__1.52-489.fc37.x86_64",
        sha256 = "c5df34198e7dd39f4f09032beacb9db641c8752d045b8e1f8cacd2637559dd1d",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Carp-1.52-489.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Carp-1.52-489.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Carp-1.52-489.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Carp-1.52-489.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Carp-1.52-489.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Class-Struct-0__0.66-494.fc37.x86_64",
        sha256 = "a250d4c099b3919a35163ac204adbca30a331d4d1108288db176fb1329c7766e",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Class-Struct-0.66-494.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Class-Struct-0.66-494.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Class-Struct-0.66-494.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Class-Struct-0.66-494.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Class-Struct-0.66-494.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Compress-Bzip2-0__2.28-10.fc37.x86_64",
        sha256 = "336fa6c00e0f4a1a078e445a8f70a8cb79a0f783df8479d77a4de600f2abd097",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Compress-Bzip2-2.28-10.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Compress-Bzip2-2.28-10.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Compress-Bzip2-2.28-10.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Compress-Bzip2-2.28-10.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Compress-Bzip2-2.28-10.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "perl-Compress-Raw-Bzip2-0__2.201-2.fc37.x86_64",
        sha256 = "eb4cc36b986a52a9ad30ba5f64a06b3f88bdfc687236532c0d3e442603210034",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Compress-Raw-Bzip2-2.201-2.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Compress-Raw-Bzip2-2.201-2.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Compress-Raw-Bzip2-2.201-2.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Compress-Raw-Bzip2-2.201-2.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Compress-Raw-Bzip2-2.201-2.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "perl-Compress-Raw-Lzma-0__2.201-7.fc37.x86_64",
        sha256 = "3c14f0034f8cc919ddee075d0bb2a651efdbc12860e3edd3ba4090c641e9ff3f",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Compress-Raw-Lzma-2.201-7.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Compress-Raw-Lzma-2.201-7.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Compress-Raw-Lzma-2.201-7.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Compress-Raw-Lzma-2.201-7.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Compress-Raw-Lzma-2.201-7.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "perl-Compress-Raw-Zlib-0__2.202-3.fc37.x86_64",
        sha256 = "101823fac489a2eaba7ba2bd04e21a9daffae7dae98be88913000d5f0a90ad0f",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Compress-Raw-Zlib-2.202-3.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Compress-Raw-Zlib-2.202-3.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Compress-Raw-Zlib-2.202-3.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Compress-Raw-Zlib-2.202-3.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Compress-Raw-Zlib-2.202-3.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "perl-Config-Extensions-0__0.03-494.fc37.x86_64",
        sha256 = "55ad51d0b553a56eb31227a21c01203830711ce0711bab675aa98c1aa7dc7e8f",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Config-Extensions-0.03-494.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Config-Extensions-0.03-494.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Config-Extensions-0.03-494.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Config-Extensions-0.03-494.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Config-Extensions-0.03-494.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Config-Perl-V-0__0.34-1.fc37.x86_64",
        sha256 = "fae66bb345d67beff34771ac5035f42aa1b91272c3ed5f8ee318c343ac9fa79e",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Config-Perl-V-0.34-1.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Config-Perl-V-0.34-1.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Config-Perl-V-0.34-1.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Config-Perl-V-0.34-1.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Config-Perl-V-0.34-1.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-DBM_Filter-0__0.06-494.fc37.x86_64",
        sha256 = "a591414027c5399d77a72d4ce8b1a0d9f1aaa15c7047ffd94537cb1f9dbc20c3",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-DBM_Filter-0.06-494.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-DBM_Filter-0.06-494.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-DBM_Filter-0.06-494.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-DBM_Filter-0.06-494.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-DBM_Filter-0.06-494.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-DB_File-0__1.858-4.fc37.x86_64",
        sha256 = "cd0367bee7e6c35f68fc10fd51a318978572aaec2f049b89acb8e7fd07f19620",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-DB_File-1.858-4.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-DB_File-1.858-4.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-DB_File-1.858-4.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-DB_File-1.858-4.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-DB_File-1.858-4.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "perl-Data-Dumper-0__2.184-490.fc37.x86_64",
        sha256 = "087435f52f5910c5d769ddac9154e9ac54b50c005b2036607d5e2187b3a9b7d9",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Data-Dumper-2.184-490.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Data-Dumper-2.184-490.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Data-Dumper-2.184-490.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Data-Dumper-2.184-490.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Data-Dumper-2.184-490.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "perl-Data-OptList-0__0.112-5.fc37.x86_64",
        sha256 = "ab4aae8175f2ee46cca0d406b7511178b5733f6b67260c3b4194fd07fab7f89a",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Data-OptList-0.112-5.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Data-OptList-0.112-5.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Data-OptList-0.112-5.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Data-OptList-0.112-5.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Data-OptList-0.112-5.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Data-Section-0__0.200007-17.fc37.x86_64",
        sha256 = "97659ba6b8b7e175d0eb235afac957af0f3b7a0fd494924ba800cfa6db8176a9",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Data-Section-0.200007-17.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Data-Section-0.200007-17.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Data-Section-0.200007-17.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Data-Section-0.200007-17.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Data-Section-0.200007-17.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Devel-PPPort-0__3.68-490.fc37.x86_64",
        sha256 = "69811269661d579913e8ac3a97ac8887f382ab561abd621b4d3f71baaf458177",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Devel-PPPort-3.68-490.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Devel-PPPort-3.68-490.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Devel-PPPort-3.68-490.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Devel-PPPort-3.68-490.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Devel-PPPort-3.68-490.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "perl-Devel-Peek-0__1.32-494.fc37.x86_64",
        sha256 = "f302e6eed403107ec2006c2d63692d447a8cad298b2ccc9009255d4c46bb8ef0",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Devel-Peek-1.32-494.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Devel-Peek-1.32-494.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Devel-Peek-1.32-494.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Devel-Peek-1.32-494.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Devel-Peek-1.32-494.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "perl-Devel-SelfStubber-0__1.06-494.fc37.x86_64",
        sha256 = "85bcc5243c80acb0cf19eed20b61a0dcf6bae4e82f19cee9696dbbced650b7cf",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Devel-SelfStubber-1.06-494.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Devel-SelfStubber-1.06-494.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Devel-SelfStubber-1.06-494.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Devel-SelfStubber-1.06-494.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Devel-SelfStubber-1.06-494.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Devel-Size-0__0.83-13.fc37.x86_64",
        sha256 = "1a40670bcdc99252ea6258e647502cbe897fbb34ab8078b32f8b4b1e8b71b15f",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Devel-Size-0.83-13.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Devel-Size-0.83-13.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Devel-Size-0.83-13.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Devel-Size-0.83-13.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Devel-Size-0.83-13.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "perl-Digest-0__1.20-489.fc37.x86_64",
        sha256 = "17acbd974328ca2ac974f45953fb826d09cb071ee381f975928592b72c8b2d9f",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Digest-1.20-489.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Digest-1.20-489.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Digest-1.20-489.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Digest-1.20-489.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Digest-1.20-489.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Digest-MD5-0__2.58-489.fc37.x86_64",
        sha256 = "ab79d34ab61c3f094a7291b2399ddf7de988fa0aa7bbe51cb5d10c8066e83f45",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Digest-MD5-2.58-489.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Digest-MD5-2.58-489.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Digest-MD5-2.58-489.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Digest-MD5-2.58-489.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Digest-MD5-2.58-489.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "perl-Digest-SHA-1__6.03-1.fc37.x86_64",
        sha256 = "a941d164018b903f7f1c0e388a97ef4f591c887dc51a4b1fa240ba794e1fd29b",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Digest-SHA-6.03-1.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Digest-SHA-6.03-1.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Digest-SHA-6.03-1.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Digest-SHA-6.03-1.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Digest-SHA-6.03-1.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "perl-Digest-SHA1-0__2.13-37.fc37.x86_64",
        sha256 = "e037888f738bb4b93e5ff26eda54e3518a3d714f9d0081fc5f245057fddaada3",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Digest-SHA1-2.13-37.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Digest-SHA1-2.13-37.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Digest-SHA1-2.13-37.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Digest-SHA1-2.13-37.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Digest-SHA1-2.13-37.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "perl-DirHandle-0__1.05-494.fc37.x86_64",
        sha256 = "9e16a556068804961a97607007c7aeeffd5d2e5fb5a56d03e348fdbab53b8d4f",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-DirHandle-1.05-494.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-DirHandle-1.05-494.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-DirHandle-1.05-494.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-DirHandle-1.05-494.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-DirHandle-1.05-494.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Dumpvalue-0__2.27-494.fc37.x86_64",
        sha256 = "8b163fcae109702b95fe6425aea1551504c23247641c757b1c20dc8c74d3fcb4",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Dumpvalue-2.27-494.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Dumpvalue-2.27-494.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Dumpvalue-2.27-494.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Dumpvalue-2.27-494.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Dumpvalue-2.27-494.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-DynaLoader-0__1.52-494.fc37.x86_64",
        sha256 = "a11b251a948378e9ea319bfcc63e8daaad6a5406ce567f410478739f44c4214e",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-DynaLoader-1.52-494.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-DynaLoader-1.52-494.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-DynaLoader-1.52-494.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-DynaLoader-1.52-494.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-DynaLoader-1.52-494.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "perl-Encode-4__3.19-492.fc37.x86_64",
        sha256 = "395705071c61bc6faad7cde8eb251d3bff87f543c5096c00139331e9ee2ba856",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Encode-3.19-492.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Encode-3.19-492.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Encode-3.19-492.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Encode-3.19-492.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Encode-3.19-492.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "perl-Encode-devel-4__3.19-492.fc37.x86_64",
        sha256 = "cfad3343efba5b3bb80b56e8cb32d84845c35ba0499e2dc0c0b586dd23d64b49",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Encode-devel-3.19-492.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Encode-devel-3.19-492.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Encode-devel-3.19-492.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Encode-devel-3.19-492.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Encode-devel-3.19-492.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "perl-English-0__1.11-494.fc37.x86_64",
        sha256 = "13f14f71bdc3a7cf3d9388856062582f4f51bcb450f87ad8dfbcf51e5da4e0f7",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-English-1.11-494.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-English-1.11-494.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-English-1.11-494.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-English-1.11-494.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-English-1.11-494.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Env-0__1.05-489.fc37.x86_64",
        sha256 = "03c735e2591c14afeaab7e61da2207794436493bee277163db2b51790aeb4ab1",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Env-1.05-489.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Env-1.05-489.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Env-1.05-489.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Env-1.05-489.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Env-1.05-489.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Errno-0__1.36-494.fc37.x86_64",
        sha256 = "d73a22cd89268fdc6bbacac05980466b96c91a71306866ed4552835b05506a5b",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Errno-1.36-494.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Errno-1.36-494.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Errno-1.36-494.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Errno-1.36-494.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Errno-1.36-494.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "perl-Exporter-0__5.77-489.fc37.x86_64",
        sha256 = "95b26bb47a5b0f52f091cf6a5e6a493b203ed6d1bf8de714ed182c4f78f8b351",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Exporter-5.77-489.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Exporter-5.77-489.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Exporter-5.77-489.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Exporter-5.77-489.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Exporter-5.77-489.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-ExtUtils-CBuilder-1__0.280236-489.fc37.x86_64",
        sha256 = "3f12827f0d6af1fa9312485811bfe94be5a6a49cc60f51c00a7b45418fc8a736",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-ExtUtils-CBuilder-0.280236-489.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-ExtUtils-CBuilder-0.280236-489.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-ExtUtils-CBuilder-0.280236-489.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-ExtUtils-CBuilder-0.280236-489.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-ExtUtils-CBuilder-0.280236-489.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-ExtUtils-Command-2__7.66-1.fc37.x86_64",
        sha256 = "380a1fe3c4546fdcaed2bd10de4eb600c3726914f73b7d4675290c371c2522c0",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-ExtUtils-Command-7.66-1.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-ExtUtils-Command-7.66-1.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-ExtUtils-Command-7.66-1.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-ExtUtils-Command-7.66-1.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-ExtUtils-Command-7.66-1.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-ExtUtils-Constant-0__0.25-494.fc37.x86_64",
        sha256 = "f5fa5939e54b92f89a448c8598112accbde3e33396483c721cc84dfc42556c6c",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-ExtUtils-Constant-0.25-494.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-ExtUtils-Constant-0.25-494.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-ExtUtils-Constant-0.25-494.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-ExtUtils-Constant-0.25-494.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-ExtUtils-Constant-0.25-494.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-ExtUtils-Embed-0__1.35-494.fc37.x86_64",
        sha256 = "9d27313b588a3bb95b931b6fac98fa3d235b88b8f45504f3db379350daa2ff75",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-ExtUtils-Embed-1.35-494.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-ExtUtils-Embed-1.35-494.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-ExtUtils-Embed-1.35-494.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-ExtUtils-Embed-1.35-494.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-ExtUtils-Embed-1.35-494.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-ExtUtils-Install-0__2.20-489.fc37.x86_64",
        sha256 = "fdc9b741640d5449a3736abd433057b35b29bae743b4b34f2896dbe254e13914",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-ExtUtils-Install-2.20-489.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-ExtUtils-Install-2.20-489.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-ExtUtils-Install-2.20-489.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-ExtUtils-Install-2.20-489.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-ExtUtils-Install-2.20-489.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-ExtUtils-MM-Utils-2__7.66-1.fc37.x86_64",
        sha256 = "5c1097fe06bb7d927ba17b379670d1d0b78c57c77a85edac55ef3e4b8ef65624",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-ExtUtils-MM-Utils-7.66-1.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-ExtUtils-MM-Utils-7.66-1.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-ExtUtils-MM-Utils-7.66-1.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-ExtUtils-MM-Utils-7.66-1.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-ExtUtils-MM-Utils-7.66-1.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-ExtUtils-MakeMaker-2__7.66-1.fc37.x86_64",
        sha256 = "d27ad9fb12959eafe80ef6719df1a56730c1ff2123511b401cba1a5dfdcbec5b",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-ExtUtils-MakeMaker-7.66-1.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-ExtUtils-MakeMaker-7.66-1.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-ExtUtils-MakeMaker-7.66-1.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-ExtUtils-MakeMaker-7.66-1.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-ExtUtils-MakeMaker-7.66-1.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-ExtUtils-Manifest-1__1.73-489.fc37.x86_64",
        sha256 = "d23efacaadaabaaa74877985eac9979f5f06ad9fe1b2cacb242070b478a3bb6d",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-ExtUtils-Manifest-1.73-489.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-ExtUtils-Manifest-1.73-489.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-ExtUtils-Manifest-1.73-489.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-ExtUtils-Manifest-1.73-489.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-ExtUtils-Manifest-1.73-489.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-ExtUtils-Miniperl-0__1.11-494.fc37.x86_64",
        sha256 = "81b9ec0d4078df73e5138d748abc3f5038ea649ace8593919eada5bfa850020d",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-ExtUtils-Miniperl-1.11-494.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-ExtUtils-Miniperl-1.11-494.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-ExtUtils-Miniperl-1.11-494.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-ExtUtils-Miniperl-1.11-494.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-ExtUtils-Miniperl-1.11-494.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-ExtUtils-ParseXS-1__3.45-489.fc37.x86_64",
        sha256 = "924465679e012b6896b9372d22650e5b79833d71849d3dec72a64b05546eada7",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-ExtUtils-ParseXS-3.45-489.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-ExtUtils-ParseXS-3.45-489.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-ExtUtils-ParseXS-3.45-489.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-ExtUtils-ParseXS-3.45-489.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-ExtUtils-ParseXS-3.45-489.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Fcntl-0__1.15-494.fc37.x86_64",
        sha256 = "821dc308cdd46f74ae0d3cb75e9dfd8fc4f3ab26d984d6d3bc518c8144d6f73a",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Fcntl-1.15-494.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Fcntl-1.15-494.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Fcntl-1.15-494.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Fcntl-1.15-494.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Fcntl-1.15-494.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "perl-File-Basename-0__2.85-494.fc37.x86_64",
        sha256 = "35a66ded720e73b06f8449f74c53ce5ec5bcdbcc2783808e0b9dbecb91dfab96",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-File-Basename-2.85-494.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-File-Basename-2.85-494.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-File-Basename-2.85-494.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-File-Basename-2.85-494.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-File-Basename-2.85-494.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-File-Compare-0__1.100.700-494.fc37.x86_64",
        sha256 = "e454119581f062a111ee025e70c9ea9fb415f5ebdd85435d6afc8bfa77157550",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-File-Compare-1.100.700-494.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-File-Compare-1.100.700-494.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-File-Compare-1.100.700-494.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-File-Compare-1.100.700-494.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-File-Compare-1.100.700-494.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-File-Copy-0__2.39-494.fc37.x86_64",
        sha256 = "613f5b89c3246b4ebd650038b93df795b71ff57e318de183fbd32b574e264378",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-File-Copy-2.39-494.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-File-Copy-2.39-494.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-File-Copy-2.39-494.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-File-Copy-2.39-494.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-File-Copy-2.39-494.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-File-DosGlob-0__1.12-494.fc37.x86_64",
        sha256 = "43bb5186dbf91a8146f610308976ff73df4847bb20cb8d560f4953f822ca0b2e",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-File-DosGlob-1.12-494.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-File-DosGlob-1.12-494.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-File-DosGlob-1.12-494.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-File-DosGlob-1.12-494.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-File-DosGlob-1.12-494.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "perl-File-Fetch-0__1.04-489.fc37.x86_64",
        sha256 = "f7a51e94040539753ed04fac0cf8ec179dcd3e7dceeb3a38fd889ccbe1b1789c",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-File-Fetch-1.04-489.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-File-Fetch-1.04-489.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-File-Fetch-1.04-489.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-File-Fetch-1.04-489.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-File-Fetch-1.04-489.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-File-Find-0__1.40-494.fc37.x86_64",
        sha256 = "71513593cf19b9d51fd17b5606e5bd3f4f0650bcddf69c2b73830bb2ab5fcabf",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-File-Find-1.40-494.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-File-Find-1.40-494.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-File-Find-1.40-494.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-File-Find-1.40-494.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-File-Find-1.40-494.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-File-HomeDir-0__1.006-7.fc37.x86_64",
        sha256 = "288f5bd2f09143037798cd867fcaf8a6e5c6bb7c5cf47b9de8bcf9ee7606ca12",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-File-HomeDir-1.006-7.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-File-HomeDir-1.006-7.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-File-HomeDir-1.006-7.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-File-HomeDir-1.006-7.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-File-HomeDir-1.006-7.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-File-Path-0__2.18-489.fc37.x86_64",
        sha256 = "d73acee5758b6ac85b46416094fcaf9a3e9f54193996f167d353777e5708745a",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-File-Path-2.18-489.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-File-Path-2.18-489.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-File-Path-2.18-489.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-File-Path-2.18-489.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-File-Path-2.18-489.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-File-Temp-1__0.231.100-489.fc37.x86_64",
        sha256 = "5724aaa686d2cd278da72939aa01678cb6d9ba2b43237425ce2378197fbcc0d0",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-File-Temp-0.231.100-489.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-File-Temp-0.231.100-489.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-File-Temp-0.231.100-489.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-File-Temp-0.231.100-489.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-File-Temp-0.231.100-489.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-File-Which-0__1.27-6.fc37.x86_64",
        sha256 = "f8d6bfb8f2c5ca8539277e80c781d90a1cfa4e72021f1c912067c1dd9836d7be",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-File-Which-1.27-6.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-File-Which-1.27-6.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-File-Which-1.27-6.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-File-Which-1.27-6.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-File-Which-1.27-6.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-File-stat-0__1.12-494.fc37.x86_64",
        sha256 = "09c723c96f4cf3170f8460a5b1290faf55f41c3ae4a54a51ef9d0c3113457085",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-File-stat-1.12-494.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-File-stat-1.12-494.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-File-stat-1.12-494.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-File-stat-1.12-494.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-File-stat-1.12-494.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-FileCache-0__1.10-494.fc37.x86_64",
        sha256 = "33bdd78654deba9f247f45c2945d72d2ea529913a9c7140c1dd9196fc5056376",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-FileCache-1.10-494.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-FileCache-1.10-494.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-FileCache-1.10-494.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-FileCache-1.10-494.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-FileCache-1.10-494.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-FileHandle-0__2.03-494.fc37.x86_64",
        sha256 = "35bc29a12858c7f1ac793fef5e2f3cf4951fcbb0843246a06f445cbd6ef00b5c",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-FileHandle-2.03-494.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-FileHandle-2.03-494.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-FileHandle-2.03-494.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-FileHandle-2.03-494.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-FileHandle-2.03-494.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Filter-2__1.64-1.fc37.x86_64",
        sha256 = "a6400b14b7905ca51354226996a31778c5e720f2ff04c53efb06bf1a72c8cd57",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Filter-1.64-1.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Filter-1.64-1.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Filter-1.64-1.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Filter-1.64-1.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Filter-1.64-1.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "perl-Filter-Simple-0__0.96-489.fc37.x86_64",
        sha256 = "0a5fff5e31a412e05287bb2f1da0a15ba4b3e14e1acb99fda077dda9a965d52e",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Filter-Simple-0.96-489.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Filter-Simple-0.96-489.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Filter-Simple-0.96-489.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Filter-Simple-0.96-489.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Filter-Simple-0.96-489.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-FindBin-0__1.53-494.fc37.x86_64",
        sha256 = "67dbdd42db129bcea067cd151cd54ff0e245c8fa5a6f7cbf7ed202c191efae5e",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-FindBin-1.53-494.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-FindBin-1.53-494.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-FindBin-1.53-494.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-FindBin-1.53-494.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-FindBin-1.53-494.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-GDBM_File-1__1.23-494.fc37.x86_64",
        sha256 = "1a78697276a561680873fd82df2cd5ad12ecc44899c7fe9f5b0ed8cfee34b561",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-GDBM_File-1.23-494.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-GDBM_File-1.23-494.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-GDBM_File-1.23-494.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-GDBM_File-1.23-494.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-GDBM_File-1.23-494.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "perl-Getopt-Long-1__2.54-1.fc37.x86_64",
        sha256 = "49e397448b20d87418f8a7b7b3ae474928b6e33157dbaf300dc8d8ed842804c8",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Getopt-Long-2.54-1.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Getopt-Long-2.54-1.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Getopt-Long-2.54-1.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Getopt-Long-2.54-1.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Getopt-Long-2.54-1.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Getopt-Std-0__1.13-494.fc37.x86_64",
        sha256 = "e7b8e9c6ce42eaf008c044007ee28d6ca1a66e767933034b9c14bd40568c9ce9",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Getopt-Std-1.13-494.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Getopt-Std-1.13-494.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Getopt-Std-1.13-494.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Getopt-Std-1.13-494.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Getopt-Std-1.13-494.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-HTTP-Tiny-0__0.086-1.fc37.x86_64",
        sha256 = "06a15be4faf8f31a6c8fae57c435870d03ac683cdd9ecb5136963b6f21dc397d",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-HTTP-Tiny-0.086-1.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-HTTP-Tiny-0.086-1.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-HTTP-Tiny-0.086-1.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-HTTP-Tiny-0.086-1.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-HTTP-Tiny-0.086-1.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Hash-Util-0__0.28-494.fc37.x86_64",
        sha256 = "7360c97f06ac11656344a6d1eb5184ddf8202496d8bf8963b525291dd9d571b7",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Hash-Util-0.28-494.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Hash-Util-0.28-494.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Hash-Util-0.28-494.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Hash-Util-0.28-494.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Hash-Util-0.28-494.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "perl-Hash-Util-FieldHash-0__1.26-494.fc37.x86_64",
        sha256 = "e1ac0652dc2950ea39fb06be09289eea88487d1d52a0ea0782432bb3cebdda05",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Hash-Util-FieldHash-1.26-494.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Hash-Util-FieldHash-1.26-494.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Hash-Util-FieldHash-1.26-494.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Hash-Util-FieldHash-1.26-494.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Hash-Util-FieldHash-1.26-494.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "perl-I18N-Collate-0__1.02-494.fc37.x86_64",
        sha256 = "4ec3b4c0ca262dabaea7dc5f41e35f1d170a10649cf1e1469cdf4929e56aefac",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-I18N-Collate-1.02-494.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-I18N-Collate-1.02-494.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-I18N-Collate-1.02-494.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-I18N-Collate-1.02-494.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-I18N-Collate-1.02-494.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-I18N-LangTags-0__0.45-494.fc37.x86_64",
        sha256 = "4683c6892310b4af2bc01e30de417ae02ce42626106a6b41391bd9585a4ae896",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-I18N-LangTags-0.45-494.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-I18N-LangTags-0.45-494.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-I18N-LangTags-0.45-494.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-I18N-LangTags-0.45-494.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-I18N-LangTags-0.45-494.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-I18N-Langinfo-0__0.21-494.fc37.x86_64",
        sha256 = "e2dc8c4c49010167517a453d6a38d5200cdee37d402e95ae987c546f28484a28",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-I18N-Langinfo-0.21-494.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-I18N-Langinfo-0.21-494.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-I18N-Langinfo-0.21-494.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-I18N-Langinfo-0.21-494.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-I18N-Langinfo-0.21-494.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "perl-IO-0__1.50-494.fc37.x86_64",
        sha256 = "4ed96a75703de0c51aa26ab3dba42afcdb55cf93c56476eca7c3ccd729f09d4a",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-IO-1.50-494.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-IO-1.50-494.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-IO-1.50-494.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-IO-1.50-494.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-IO-1.50-494.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "perl-IO-Compress-0__2.201-3.fc37.x86_64",
        sha256 = "14defb2171d09871afb150a2b2bd6d86829e294bb3fde518766ccbf48b049c43",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-IO-Compress-2.201-3.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-IO-Compress-2.201-3.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-IO-Compress-2.201-3.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-IO-Compress-2.201-3.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-IO-Compress-2.201-3.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-IO-Compress-Lzma-0__2.201-2.fc37.x86_64",
        sha256 = "7533be1fa034b000f85c463e5d81b80665520cb4d513b86b8ab685bde599bb2b",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-IO-Compress-Lzma-2.201-2.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-IO-Compress-Lzma-2.201-2.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-IO-Compress-Lzma-2.201-2.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-IO-Compress-Lzma-2.201-2.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-IO-Compress-Lzma-2.201-2.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-IO-Socket-IP-0__0.41-490.fc37.x86_64",
        sha256 = "59c3fcf829ce856d6d4f81c4a3d00b632d324af71a3de1801001daf3b241eb1b",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-IO-Socket-IP-0.41-490.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-IO-Socket-IP-0.41-490.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-IO-Socket-IP-0.41-490.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-IO-Socket-IP-0.41-490.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-IO-Socket-IP-0.41-490.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-IO-Zlib-1__1.11-489.fc37.x86_64",
        sha256 = "7fbae28ec6b0925b729c04a2430076371607da2dd38a46cfc3c8d58ac4d19943",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-IO-Zlib-1.11-489.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-IO-Zlib-1.11-489.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-IO-Zlib-1.11-489.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-IO-Zlib-1.11-489.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-IO-Zlib-1.11-489.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-IPC-Cmd-2__1.04-490.fc37.x86_64",
        sha256 = "c4b12c0bc3065a0298dea86e0387a57e291fdd43ebbf5f1881ca4f97d124c5e2",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-IPC-Cmd-1.04-490.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-IPC-Cmd-1.04-490.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-IPC-Cmd-1.04-490.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-IPC-Cmd-1.04-490.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-IPC-Cmd-1.04-490.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-IPC-Open3-0__1.22-494.fc37.x86_64",
        sha256 = "8b67d74e413d20ab80e4b5bdf8a784cb7492734b370099bc105f0faac7bab1ab",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-IPC-Open3-1.22-494.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-IPC-Open3-1.22-494.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-IPC-Open3-1.22-494.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-IPC-Open3-1.22-494.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-IPC-Open3-1.22-494.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-IPC-SysV-0__2.09-490.fc37.x86_64",
        sha256 = "80153b91d4c4c7c0bf5b39fdf42bff9a66472bfa01b32d4c4152b4610af5a83c",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-IPC-SysV-2.09-490.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-IPC-SysV-2.09-490.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-IPC-SysV-2.09-490.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-IPC-SysV-2.09-490.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-IPC-SysV-2.09-490.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "perl-IPC-System-Simple-0__1.30-9.fc37.x86_64",
        sha256 = "9c8f39cc95ff600c5cd60b4e753ff26f7bb2e1664607163063bd4b18da6dea13",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-IPC-System-Simple-1.30-9.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-IPC-System-Simple-1.30-9.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-IPC-System-Simple-1.30-9.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-IPC-System-Simple-1.30-9.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-IPC-System-Simple-1.30-9.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Importer-0__0.026-7.fc37.x86_64",
        sha256 = "6628070ada3140f046e9a03574be0fea3963c5cdd4bda25b630d400c5c14838d",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Importer-0.026-7.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Importer-0.026-7.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Importer-0.026-7.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Importer-0.026-7.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Importer-0.026-7.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-JSON-PP-1__4.11-1.fc37.x86_64",
        sha256 = "1aa8f2ca107295bc2cd5510e8a1c29f98d178d84badb5dfd352c082aede61cdf",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-JSON-PP-4.11-1.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-JSON-PP-4.11-1.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-JSON-PP-4.11-1.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-JSON-PP-4.11-1.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-JSON-PP-4.11-1.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Locale-Maketext-0__1.32-1.fc37.x86_64",
        sha256 = "db61bd3244856beafead63d361b98ecf4cfe6945dd9169ffbf3973f85e419b21",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Locale-Maketext-1.32-1.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Locale-Maketext-1.32-1.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Locale-Maketext-1.32-1.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Locale-Maketext-1.32-1.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Locale-Maketext-1.32-1.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Locale-Maketext-Simple-1__0.21-494.fc37.x86_64",
        sha256 = "13f7264c14d1e1de753e205c40c8bd5a8642e0a0215817540486ec3870fdf125",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Locale-Maketext-Simple-0.21-494.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Locale-Maketext-Simple-0.21-494.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Locale-Maketext-Simple-0.21-494.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Locale-Maketext-Simple-0.21-494.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Locale-Maketext-Simple-0.21-494.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-MIME-Base64-0__3.16-489.fc37.x86_64",
        sha256 = "0823ebce69b5d9df94c9d508669ac612bd1df48e8857cf02103dcc2094df246f",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-MIME-Base64-3.16-489.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-MIME-Base64-3.16-489.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-MIME-Base64-3.16-489.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-MIME-Base64-3.16-489.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-MIME-Base64-3.16-489.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "perl-MRO-Compat-0__0.15-4.fc37.x86_64",
        sha256 = "c92e79de1b50f817638d5b21de3b3426950404280ff32470a1c5ea37f4d7f803",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-MRO-Compat-0.15-4.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-MRO-Compat-0.15-4.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-MRO-Compat-0.15-4.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-MRO-Compat-0.15-4.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-MRO-Compat-0.15-4.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Math-BigInt-1__1.9998.37-2.fc37.x86_64",
        sha256 = "4ae9d5dfd60d270eafe87e0781e836bad36a9b71aa28a5e84d41677600904444",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Math-BigInt-1.9998.37-2.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Math-BigInt-1.9998.37-2.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Math-BigInt-1.9998.37-2.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Math-BigInt-1.9998.37-2.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Math-BigInt-1.9998.37-2.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Math-BigInt-FastCalc-0__0.501.300-3.fc37.x86_64",
        sha256 = "7c2f97cd627fd0a926e7f00daaebc066c8fedfe0e6ebb6c01c42347e0ae6ecec",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Math-BigInt-FastCalc-0.501.300-3.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Math-BigInt-FastCalc-0.501.300-3.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Math-BigInt-FastCalc-0.501.300-3.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Math-BigInt-FastCalc-0.501.300-3.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Math-BigInt-FastCalc-0.501.300-3.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "perl-Math-BigRat-0__0.2624-2.fc37.x86_64",
        sha256 = "837314dcdcae121138b9d14fae2efe2aa68a1313b0ac82d774ebff88eba85197",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Math-BigRat-0.2624-2.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Math-BigRat-0.2624-2.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Math-BigRat-0.2624-2.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Math-BigRat-0.2624-2.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Math-BigRat-0.2624-2.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Math-Complex-0__1.59-494.fc37.x86_64",
        sha256 = "7dc1ba1baa4386f259673974d8dcf60120881fc70c036d526fdfce2451ef8422",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Math-Complex-1.59-494.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Math-Complex-1.59-494.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Math-Complex-1.59-494.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Math-Complex-1.59-494.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Math-Complex-1.59-494.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Memoize-0__1.03-494.fc37.x86_64",
        sha256 = "c831644c3fd5b08a70ee442b7f2e8f66401cddab43551628769626f9ad387207",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Memoize-1.03-494.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Memoize-1.03-494.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Memoize-1.03-494.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Memoize-1.03-494.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Memoize-1.03-494.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Module-Build-2__0.42.31-15.fc37.x86_64",
        sha256 = "e3a45c170655d16f2033b203a8fcac4fc65e6bc920eeeb3eea48b72a0942f783",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Module-Build-0.42.31-15.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Module-Build-0.42.31-15.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Module-Build-0.42.31-15.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Module-Build-0.42.31-15.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Module-Build-0.42.31-15.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Module-CoreList-1__5.20230820-1.fc37.x86_64",
        sha256 = "ba45ee9740e446f6a7f3824e21aa4bdf5d415f806342a65793cdca7385f383e7",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Module-CoreList-5.20230820-1.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Module-CoreList-5.20230820-1.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Module-CoreList-5.20230820-1.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Module-CoreList-5.20230820-1.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Module-CoreList-5.20230820-1.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Module-CoreList-tools-1__5.20230820-1.fc37.x86_64",
        sha256 = "3694bb286ae99d0eee79e678e9dc86081df19c93c231087274def653e26de07e",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Module-CoreList-tools-5.20230820-1.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Module-CoreList-tools-5.20230820-1.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Module-CoreList-tools-5.20230820-1.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Module-CoreList-tools-5.20230820-1.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Module-CoreList-tools-5.20230820-1.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Module-Load-1__0.36-489.fc37.x86_64",
        sha256 = "ee43e582ffdcd8c517a16cde7aea081c92509ef10da7996713b4e03dbc2806bc",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Module-Load-0.36-489.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Module-Load-0.36-489.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Module-Load-0.36-489.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Module-Load-0.36-489.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Module-Load-0.36-489.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Module-Load-Conditional-0__0.74-489.fc37.x86_64",
        sha256 = "b9445196cf8156f908a5ab0e48ad4493d4061e2e6bee63b011da70d12898bcec",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Module-Load-Conditional-0.74-489.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Module-Load-Conditional-0.74-489.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Module-Load-Conditional-0.74-489.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Module-Load-Conditional-0.74-489.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Module-Load-Conditional-0.74-489.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Module-Loaded-1__0.08-494.fc37.x86_64",
        sha256 = "cf4dababe21e9f04cf27b7e7d9d1b7ed3250a9ca255240b4640403d7a7c4c4e0",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Module-Loaded-0.08-494.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Module-Loaded-0.08-494.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Module-Loaded-0.08-494.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Module-Loaded-0.08-494.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Module-Loaded-0.08-494.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Module-Metadata-0__1.000037-489.fc37.x86_64",
        sha256 = "bd8f7944c9d518ad043a31604057d409e6cf735c3046aa3b5e62750baa29c283",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Module-Metadata-1.000037-489.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Module-Metadata-1.000037-489.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Module-Metadata-1.000037-489.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Module-Metadata-1.000037-489.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Module-Metadata-1.000037-489.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Module-Signature-0__0.88-4.fc37.x86_64",
        sha256 = "41ba8e5b325d9dcbb1b2bcf042ea01da282e2654f458d819fbc0df38e57d38bf",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Module-Signature-0.88-4.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Module-Signature-0.88-4.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Module-Signature-0.88-4.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Module-Signature-0.88-4.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Module-Signature-0.88-4.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-NDBM_File-0__1.15-494.fc37.x86_64",
        sha256 = "744f21d6397e3dd9534b3c4655624cd5898560aad623eb0c2b018f670d8a03c9",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-NDBM_File-1.15-494.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-NDBM_File-1.15-494.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-NDBM_File-1.15-494.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-NDBM_File-1.15-494.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-NDBM_File-1.15-494.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "perl-NEXT-0__0.69-494.fc37.x86_64",
        sha256 = "ce83597433fa17af59148e0a5df392926a1e5e13d1c87f9dc056ad15d5dbfc78",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-NEXT-0.69-494.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-NEXT-0.69-494.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-NEXT-0.69-494.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-NEXT-0.69-494.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-NEXT-0.69-494.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Net-0__1.03-494.fc37.x86_64",
        sha256 = "d4106dee5ea10bce794724d5b39d3615cbd34538a0a0d21094f8e734d4aa9fc8",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Net-1.03-494.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Net-1.03-494.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Net-1.03-494.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Net-1.03-494.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Net-1.03-494.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Net-Ping-0__2.75-1.fc37.x86_64",
        sha256 = "e6071fee087a8ec83116efe32cec0a712127211a5edb9c7dfccf4e9a92453d76",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Net-Ping-2.75-1.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Net-Ping-2.75-1.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Net-Ping-2.75-1.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Net-Ping-2.75-1.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Net-Ping-2.75-1.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-ODBM_File-0__1.17-494.fc37.x86_64",
        sha256 = "78184836431e86032f95673b6d4b7c542044240f3487ecb82010efcc83f3e853",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-ODBM_File-1.17-494.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-ODBM_File-1.17-494.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-ODBM_File-1.17-494.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-ODBM_File-1.17-494.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-ODBM_File-1.17-494.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "perl-Object-HashBase-0__0.009-10.fc37.x86_64",
        sha256 = "0bce10ae9482863a391edb2e83dff59df629f323cd284808949c87137f0f22cb",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Object-HashBase-0.009-10.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Object-HashBase-0.009-10.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Object-HashBase-0.009-10.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Object-HashBase-0.009-10.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Object-HashBase-0.009-10.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Opcode-0__1.57-494.fc37.x86_64",
        sha256 = "2a80b4548a5cd4c6d25812efcd5d29bf35c357b686324235d7949f917667498c",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Opcode-1.57-494.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Opcode-1.57-494.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Opcode-1.57-494.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Opcode-1.57-494.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Opcode-1.57-494.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "perl-POSIX-0__2.03-494.fc37.x86_64",
        sha256 = "75f3e8577138a10e0e745dd97e115e4703af3dbc4c6e8a1c99bcbee40637acbf",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-POSIX-2.03-494.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-POSIX-2.03-494.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-POSIX-2.03-494.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-POSIX-2.03-494.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-POSIX-2.03-494.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "perl-Package-Generator-0__1.106-26.fc37.x86_64",
        sha256 = "01533af797e2c6dd68fbc25272fb8e92aa6d53fffbcb460c030971982c0d61cb",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Package-Generator-1.106-26.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Package-Generator-1.106-26.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Package-Generator-1.106-26.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Package-Generator-1.106-26.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Package-Generator-1.106-26.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Params-Check-1__0.38-489.fc37.x86_64",
        sha256 = "f75effeffe02a56097b775842de8a36a29c01906c5983c39333b01b28f87ce61",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Params-Check-0.38-489.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Params-Check-0.38-489.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Params-Check-0.38-489.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Params-Check-0.38-489.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Params-Check-0.38-489.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Params-Util-0__1.102-8.fc37.x86_64",
        sha256 = "d357de4acaca544197cafc8743bc1ce599cfa1f450243541ce2b28298b06e702",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Params-Util-1.102-8.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Params-Util-1.102-8.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Params-Util-1.102-8.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Params-Util-1.102-8.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Params-Util-1.102-8.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "perl-PathTools-0__3.84-489.fc37.x86_64",
        sha256 = "e0a37fde6728e9b662d46a28dcd34399eccda9e5b696858d1447a47a82877de1",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-PathTools-3.84-489.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-PathTools-3.84-489.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-PathTools-3.84-489.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-PathTools-3.84-489.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-PathTools-3.84-489.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "perl-Perl-OSType-0__1.010-490.fc37.x86_64",
        sha256 = "d251ec15e60e1fee18b87e0c5317a7aa96bca90e8464928d7188dd9f2997b4c6",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Perl-OSType-1.010-490.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Perl-OSType-1.010-490.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Perl-OSType-1.010-490.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Perl-OSType-1.010-490.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Perl-OSType-1.010-490.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-PerlIO-via-QuotedPrint-0__0.10-3.fc37.x86_64",
        sha256 = "a5bfec0df83a04a1f0024d815858642c79af0da66b70cf6584e8d9078b37c97d",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-PerlIO-via-QuotedPrint-0.10-3.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-PerlIO-via-QuotedPrint-0.10-3.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-PerlIO-via-QuotedPrint-0.10-3.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-PerlIO-via-QuotedPrint-0.10-3.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-PerlIO-via-QuotedPrint-0.10-3.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Pod-Checker-4__1.75-3.fc37.x86_64",
        sha256 = "cfb3e223566de815650f1ef51bf4819f9a092678184af0e2eb30446be45fe01c",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Pod-Checker-1.75-3.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Pod-Checker-1.75-3.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Pod-Checker-1.75-3.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Pod-Checker-1.75-3.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Pod-Checker-1.75-3.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Pod-Escapes-1__1.07-489.fc37.x86_64",
        sha256 = "08bf74ca4b30a562d26026d0a70b508b7b926a3697d700127accf9d052c30da1",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Pod-Escapes-1.07-489.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Pod-Escapes-1.07-489.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Pod-Escapes-1.07-489.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Pod-Escapes-1.07-489.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Pod-Escapes-1.07-489.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Pod-Functions-0__1.14-494.fc37.x86_64",
        sha256 = "8e5aa9508914951f4ab45c905b0b50df036d03e8f8206ad71745264e8abc9119",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Pod-Functions-1.14-494.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Pod-Functions-1.14-494.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Pod-Functions-1.14-494.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Pod-Functions-1.14-494.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Pod-Functions-1.14-494.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Pod-Html-0__1.33-494.fc37.x86_64",
        sha256 = "014161ace565e0308c4e9d4290ca0fe54aaa384d3b50bfcad87c76d6e5f701aa",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Pod-Html-1.33-494.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Pod-Html-1.33-494.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Pod-Html-1.33-494.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Pod-Html-1.33-494.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Pod-Html-1.33-494.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Pod-Perldoc-0__3.28.01-490.fc37.x86_64",
        sha256 = "0047003615e8d018a4f4a887b414cba325ef294084faab480533f5de78ef58a4",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Pod-Perldoc-3.28.01-490.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Pod-Perldoc-3.28.01-490.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Pod-Perldoc-3.28.01-490.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Pod-Perldoc-3.28.01-490.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Pod-Perldoc-3.28.01-490.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Pod-Simple-1__3.43-490.fc37.x86_64",
        sha256 = "298ba92d8130493373f70271685ba59298377da497a39614d14a6345a44cac8a",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Pod-Simple-3.43-490.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Pod-Simple-3.43-490.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Pod-Simple-3.43-490.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Pod-Simple-3.43-490.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Pod-Simple-3.43-490.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Pod-Usage-4__2.03-3.fc37.x86_64",
        sha256 = "c216ffe76ed543dcb546178dabcceeca4a1054c81beee8de67a741ad39af314e",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Pod-Usage-2.03-3.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Pod-Usage-2.03-3.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Pod-Usage-2.03-3.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Pod-Usage-2.03-3.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Pod-Usage-2.03-3.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Safe-0__2.43-494.fc37.x86_64",
        sha256 = "dc525647470057ffb10fbfdf877dea36e4afcc8265ba7d02cf8575b4ddda4508",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Safe-2.43-494.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Safe-2.43-494.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Safe-2.43-494.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Safe-2.43-494.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Safe-2.43-494.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Scalar-List-Utils-5__1.63-489.fc37.x86_64",
        sha256 = "eed91d8529a1eee7269fd28eccdb636a6fc6dcf30d607bdb7198329f3017a74a",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Scalar-List-Utils-1.63-489.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Scalar-List-Utils-1.63-489.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Scalar-List-Utils-1.63-489.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Scalar-List-Utils-1.63-489.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Scalar-List-Utils-1.63-489.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "perl-Search-Dict-0__1.07-494.fc37.x86_64",
        sha256 = "a3e609d864a9679937fbc5edf16a8a8df8c01b2f415d1259e6ae93d0a222031d",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Search-Dict-1.07-494.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Search-Dict-1.07-494.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Search-Dict-1.07-494.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Search-Dict-1.07-494.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Search-Dict-1.07-494.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-SelectSaver-0__1.02-494.fc37.x86_64",
        sha256 = "6d03a8b30fc34659a68b7552f66133f837c7a1b420ca48e2e3398705a3507081",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-SelectSaver-1.02-494.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-SelectSaver-1.02-494.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-SelectSaver-1.02-494.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-SelectSaver-1.02-494.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-SelectSaver-1.02-494.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-SelfLoader-0__1.26-494.fc37.x86_64",
        sha256 = "8045f8a30e8d7dd9af37df7dd2ea98752cb8f940bd7899977296b603ab6ed05d",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-SelfLoader-1.26-494.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-SelfLoader-1.26-494.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-SelfLoader-1.26-494.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-SelfLoader-1.26-494.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-SelfLoader-1.26-494.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Socket-4__2.036-1.fc37.x86_64",
        sha256 = "5cda2738aaff2c73a850f4359b299028560c2106014c3bb02b340f036d81b564",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Socket-2.036-1.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Socket-2.036-1.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Socket-2.036-1.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Socket-2.036-1.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Socket-2.036-1.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "perl-Software-License-0__0.104002-2.fc37.x86_64",
        sha256 = "31cc5bad7a526176a0ab9e087752b8865c27564acd550fcc74afe2e1b79069b1",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Software-License-0.104002-2.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Software-License-0.104002-2.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Software-License-0.104002-2.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Software-License-0.104002-2.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Software-License-0.104002-2.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Storable-1__3.26-489.fc37.x86_64",
        sha256 = "22e7e312778e31de59a26759b523e4d0916429c17bd23f51a5ddc550d3b33910",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Storable-3.26-489.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Storable-3.26-489.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Storable-3.26-489.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Storable-3.26-489.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Storable-3.26-489.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "perl-Sub-Exporter-0__0.988-5.fc37.x86_64",
        sha256 = "b9cb98e5f30e478be84a1c45256016612202fed34cb08201a5c38dbb8f0c0928",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Sub-Exporter-0.988-5.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Sub-Exporter-0.988-5.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Sub-Exporter-0.988-5.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Sub-Exporter-0.988-5.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Sub-Exporter-0.988-5.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Sub-Install-0__0.928-33.fc37.x86_64",
        sha256 = "c1713bebc234bd8b30ef7b394420febc0e106501f020f901c3f5ea63f4856124",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Sub-Install-0.928-33.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Sub-Install-0.928-33.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Sub-Install-0.928-33.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Sub-Install-0.928-33.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Sub-Install-0.928-33.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Symbol-0__1.09-494.fc37.x86_64",
        sha256 = "3700230721f92da3b5f9ceb6615e9263060eb627be7f8e09f355b5b29feb7125",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Symbol-1.09-494.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Symbol-1.09-494.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Symbol-1.09-494.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Symbol-1.09-494.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Symbol-1.09-494.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Sys-Hostname-0__1.24-494.fc37.x86_64",
        sha256 = "ea57b566e159737d0ccec0b578a51a608950b65837fb9ff5454447b693634d44",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Sys-Hostname-1.24-494.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Sys-Hostname-1.24-494.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Sys-Hostname-1.24-494.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Sys-Hostname-1.24-494.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Sys-Hostname-1.24-494.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "perl-Sys-Syslog-0__0.36-490.fc37.x86_64",
        sha256 = "df45aa2eb6a33191ba5db0d56e2aa637ccb998f625d6f6ad36178f01df8817fe",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Sys-Syslog-0.36-490.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Sys-Syslog-0.36-490.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Sys-Syslog-0.36-490.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Sys-Syslog-0.36-490.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Sys-Syslog-0.36-490.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "perl-Term-ANSIColor-0__5.01-490.fc37.x86_64",
        sha256 = "4810f5377abb3ee2cbd1b2572b1c67421145304f54c491afc770101f90d27138",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Term-ANSIColor-5.01-490.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Term-ANSIColor-5.01-490.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Term-ANSIColor-5.01-490.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Term-ANSIColor-5.01-490.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Term-ANSIColor-5.01-490.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Term-Cap-0__1.17-489.fc37.x86_64",
        sha256 = "65d62ee56cba18a4fcbba844375b6ed49c024863fa6ea948e2ce699e7ab66298",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Term-Cap-1.17-489.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Term-Cap-1.17-489.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Term-Cap-1.17-489.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Term-Cap-1.17-489.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Term-Cap-1.17-489.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Term-Complete-0__1.403-494.fc37.x86_64",
        sha256 = "72aed6dfac21e5fc52079711a44bf87718bfeb701d770d32c2922144fce4b186",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Term-Complete-1.403-494.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Term-Complete-1.403-494.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Term-Complete-1.403-494.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Term-Complete-1.403-494.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Term-Complete-1.403-494.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Term-ReadLine-0__1.17-494.fc37.x86_64",
        sha256 = "c0ac34aa41b589f89df2c70e0c7dd418d71c900a35a9c1a48de1e0ce1aacbb8e",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Term-ReadLine-1.17-494.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Term-ReadLine-1.17-494.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Term-ReadLine-1.17-494.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Term-ReadLine-1.17-494.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Term-ReadLine-1.17-494.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Term-Table-0__0.016-4.fc37.x86_64",
        sha256 = "825835ea8306499ddff57f2295a05269f284b4545113b1a6f5820f1c3c2bfe23",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Term-Table-0.016-4.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Term-Table-0.016-4.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Term-Table-0.016-4.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Term-Table-0.016-4.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Term-Table-0.016-4.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Test-0__1.31-494.fc37.x86_64",
        sha256 = "d523d52a0a3d929715f384d7ca4075dca9cb7b9833f1c81f81aa829353d0ecba",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Test-1.31-494.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Test-1.31-494.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Test-1.31-494.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Test-1.31-494.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Test-1.31-494.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Test-Harness-1__3.44-490.fc37.x86_64",
        sha256 = "09c9f062e7963423ca79183c7e2534c69b74db9890e11a390b33adb0517dbbbe",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Test-Harness-3.44-490.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Test-Harness-3.44-490.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Test-Harness-3.44-490.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Test-Harness-3.44-490.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Test-Harness-3.44-490.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Test-Simple-3__1.302191-2.fc37.x86_64",
        sha256 = "90d514755fe5efbe1f85a5e4d80c8a2baec0032a9b3128a2bfce09b83e7be11e",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Test-Simple-1.302191-2.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Test-Simple-1.302191-2.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Test-Simple-1.302191-2.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Test-Simple-1.302191-2.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Test-Simple-1.302191-2.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Text-Abbrev-0__1.02-494.fc37.x86_64",
        sha256 = "735bd889ef7c76f8e64dd293bd55f94be1407c9b4f0a4533a001cab04c60937a",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Text-Abbrev-1.02-494.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Text-Abbrev-1.02-494.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Text-Abbrev-1.02-494.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Text-Abbrev-1.02-494.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Text-Abbrev-1.02-494.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Text-Balanced-0__2.06-2.fc37.x86_64",
        sha256 = "040ec83e4f9cf9126b301862dc4e963ea2ff77c1b0b9b241c41f137c40180461",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Text-Balanced-2.06-2.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Text-Balanced-2.06-2.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Text-Balanced-2.06-2.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Text-Balanced-2.06-2.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Text-Balanced-2.06-2.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Text-Diff-0__1.45-16.fc37.x86_64",
        sha256 = "76bc4ae7d8b1aaaafb4e943105b38121f22d1603295b0705d33c204d49d3da4e",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Text-Diff-1.45-16.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Text-Diff-1.45-16.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Text-Diff-1.45-16.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Text-Diff-1.45-16.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Text-Diff-1.45-16.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Text-Glob-0__0.11-18.fc37.x86_64",
        sha256 = "16bfd839b78d672bdabf0977be9609f716659bd6d42c903f6178286d18095465",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Text-Glob-0.11-18.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Text-Glob-0.11-18.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Text-Glob-0.11-18.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Text-Glob-0.11-18.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Text-Glob-0.11-18.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Text-ParseWords-0__3.31-489.fc37.x86_64",
        sha256 = "c66cc3f03aed1781132e402db92bc1bbcf27ab08283766428b55c368127fee01",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Text-ParseWords-3.31-489.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Text-ParseWords-3.31-489.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Text-ParseWords-3.31-489.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Text-ParseWords-3.31-489.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Text-ParseWords-3.31-489.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Text-Tabs__plus__Wrap-0__2023.0511-1.fc37.x86_64",
        sha256 = "7cb72df2ae275360fdba457d3c29cc35bc1dd9c086b528cc72fbe550516d8d36",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Text-Tabs+Wrap-2023.0511-1.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Text-Tabs+Wrap-2023.0511-1.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Text-Tabs+Wrap-2023.0511-1.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Text-Tabs+Wrap-2023.0511-1.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Text-Tabs+Wrap-2023.0511-1.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Text-Template-0__1.60-4.fc37.x86_64",
        sha256 = "30d4440da7f8e7a5e8a025261a6beffde1fa2c65e1a8ef89ebcc6a8d9b77c6d5",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Text-Template-1.60-4.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Text-Template-1.60-4.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Text-Template-1.60-4.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Text-Template-1.60-4.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Text-Template-1.60-4.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Thread-0__3.05-494.fc37.x86_64",
        sha256 = "f4fbe36e9da1c9305283ec522e8c0573ceb73a0e55db419c41f1ff9f4260e889",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Thread-3.05-494.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Thread-3.05-494.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Thread-3.05-494.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Thread-3.05-494.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Thread-3.05-494.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Thread-Queue-0__3.14-489.fc37.x86_64",
        sha256 = "c28496a7ccbc411d7edecbe17333a50254e057e7f686dd8be632d4396af95555",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Thread-Queue-3.14-489.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Thread-Queue-3.14-489.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Thread-Queue-3.14-489.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Thread-Queue-3.14-489.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Thread-Queue-3.14-489.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Thread-Semaphore-0__2.13-494.fc37.x86_64",
        sha256 = "0aa049af01d6bde532bbcaf95f43ea81a8a235e07d5ba177a633aa6a064f8cc5",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Thread-Semaphore-2.13-494.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Thread-Semaphore-2.13-494.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Thread-Semaphore-2.13-494.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Thread-Semaphore-2.13-494.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Thread-Semaphore-2.13-494.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Tie-0__4.6-494.fc37.x86_64",
        sha256 = "14115b5811a3be0a9e94b5a522888a6a9e8ec53ed4d7ca77b88f33b9ef911876",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Tie-4.6-494.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Tie-4.6-494.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Tie-4.6-494.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Tie-4.6-494.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Tie-4.6-494.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Tie-File-0__1.06-494.fc37.x86_64",
        sha256 = "a9b7a3dede261a3ac85d44649c5525d66dd0545a14276d47f9cb9b1f3a3715ec",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Tie-File-1.06-494.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Tie-File-1.06-494.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Tie-File-1.06-494.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Tie-File-1.06-494.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Tie-File-1.06-494.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Tie-Memoize-0__1.1-494.fc37.x86_64",
        sha256 = "9997fff923333706be7d381b24959329aef0139b76d7d5704067eb529b221360",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Tie-Memoize-1.1-494.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Tie-Memoize-1.1-494.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Tie-Memoize-1.1-494.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Tie-Memoize-1.1-494.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Tie-Memoize-1.1-494.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Tie-RefHash-0__1.40-489.fc37.x86_64",
        sha256 = "989fa380ed176283736b7c3da4348dcc44affb060d7308d06917961181d820bb",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Tie-RefHash-1.40-489.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Tie-RefHash-1.40-489.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Tie-RefHash-1.40-489.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Tie-RefHash-1.40-489.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Tie-RefHash-1.40-489.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Time-0__1.03-494.fc37.x86_64",
        sha256 = "a152603e478700c97d01244248a0331adbd284a53bc881c245e9e244f51bc06a",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Time-1.03-494.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Time-1.03-494.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Time-1.03-494.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Time-1.03-494.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Time-1.03-494.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Time-HiRes-4__1.9770-489.fc37.x86_64",
        sha256 = "2e7359cf0c3b5105bba1fc6033c91b584fe0db25e378206bb061dd8ecb131431",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Time-HiRes-1.9770-489.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Time-HiRes-1.9770-489.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Time-HiRes-1.9770-489.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Time-HiRes-1.9770-489.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Time-HiRes-1.9770-489.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "perl-Time-Local-2__1.300-489.fc37.x86_64",
        sha256 = "db087f574a6fd314f45daca13a1f4404c76144773fcd6c41df39da7fda862d8d",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Time-Local-1.300-489.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Time-Local-1.300-489.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Time-Local-1.300-489.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Time-Local-1.300-489.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Time-Local-1.300-489.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Time-Piece-0__1.3401-494.fc37.x86_64",
        sha256 = "217464bd7c2fd18f54ebba7702fb2b86c7d9e9449d72e57e5ffa3dea0e063d5f",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Time-Piece-1.3401-494.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Time-Piece-1.3401-494.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Time-Piece-1.3401-494.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Time-Piece-1.3401-494.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Time-Piece-1.3401-494.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "perl-URI-0__5.17-1.fc37.x86_64",
        sha256 = "5984ff3e5bc0ab2df2fd308b85f226e267f760cfd58c36881974b9bca1c192d1",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-URI-5.17-1.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-URI-5.17-1.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-URI-5.17-1.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-URI-5.17-1.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-URI-5.17-1.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Unicode-Collate-0__1.31-489.fc37.x86_64",
        sha256 = "fdb05f59c32f7d68026772e9f849eb52259983ab3b0d2bfeb1a239347456b09d",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Unicode-Collate-1.31-489.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Unicode-Collate-1.31-489.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Unicode-Collate-1.31-489.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Unicode-Collate-1.31-489.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Unicode-Collate-1.31-489.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "perl-Unicode-Normalize-0__1.31-489.fc37.x86_64",
        sha256 = "55da077c84c31fe12266d99462c5cc592156817260eb6234def89f8fca776ad3",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Unicode-Normalize-1.31-489.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Unicode-Normalize-1.31-489.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Unicode-Normalize-1.31-489.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Unicode-Normalize-1.31-489.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Unicode-Normalize-1.31-489.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "perl-Unicode-UCD-0__0.78-494.fc37.x86_64",
        sha256 = "206133a803fdc9ae410fe4bd639a78b17002fa7e08114096499abcbfd55a982e",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Unicode-UCD-0.78-494.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Unicode-UCD-0.78-494.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Unicode-UCD-0.78-494.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Unicode-UCD-0.78-494.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Unicode-UCD-0.78-494.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-User-pwent-0__1.03-494.fc37.x86_64",
        sha256 = "b31cec68fd37ef5dc8ddcf8f452ea9b3b9ee4919fe36854feb7669a19f5ce851",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-User-pwent-1.03-494.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-User-pwent-1.03-494.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-User-pwent-1.03-494.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-User-pwent-1.03-494.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-User-pwent-1.03-494.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-autodie-0__2.34-490.fc37.x86_64",
        sha256 = "1a5c4b040a029e13cf4f2b2ae963f2894e89a5105c6cf36c2d08172e9933cd5f",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-autodie-2.34-490.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-autodie-2.34-490.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-autodie-2.34-490.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-autodie-2.34-490.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-autodie-2.34-490.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-autouse-0__1.11-494.fc37.x86_64",
        sha256 = "9a599666959a6f45056e0e9a3caed42de73959ebcc734987a5126e81ce2c51ce",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-autouse-1.11-494.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-autouse-1.11-494.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-autouse-1.11-494.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-autouse-1.11-494.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-autouse-1.11-494.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-base-0__2.27-494.fc37.x86_64",
        sha256 = "835fdeacb26e0a39c0d7600719072eacc3611f21ed33d2f956c4c09c34e6e181",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-base-2.27-494.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-base-2.27-494.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-base-2.27-494.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-base-2.27-494.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-base-2.27-494.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-bignum-0__0.66-5.fc37.x86_64",
        sha256 = "096e17cde1b8075c993b9815bd93e979483ba8e3ed211688b4b115b38b2f9623",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-bignum-0.66-5.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-bignum-0.66-5.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-bignum-0.66-5.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-bignum-0.66-5.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-bignum-0.66-5.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-blib-0__1.07-494.fc37.x86_64",
        sha256 = "24766c3a21d65490f41938a0f081d87c5d02074c35ed78e69703c8cde6ee3b08",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-blib-1.07-494.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-blib-1.07-494.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-blib-1.07-494.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-blib-1.07-494.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-blib-1.07-494.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-constant-0__1.33-490.fc37.x86_64",
        sha256 = "7898312e5b93625a6fa5f4b60300b06260668dfb12af04602a3354266e8f7850",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-constant-1.33-490.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-constant-1.33-490.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-constant-1.33-490.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-constant-1.33-490.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-constant-1.33-490.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-debugger-0__1.60-494.fc37.x86_64",
        sha256 = "a9e1b639184a98e2e482c510629fef6aee38333ecd396adf97b953ec5744e0c0",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-debugger-1.60-494.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-debugger-1.60-494.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-debugger-1.60-494.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-debugger-1.60-494.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-debugger-1.60-494.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-deprecate-0__0.04-494.fc37.x86_64",
        sha256 = "e928eedcb3afe8ebf89c7fb36bb08211080a3f6f8b2a7217d1e7aa3908768f0d",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-deprecate-0.04-494.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-deprecate-0.04-494.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-deprecate-0.04-494.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-deprecate-0.04-494.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-deprecate-0.04-494.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-devel-4__5.36.1-494.fc37.x86_64",
        sha256 = "70100f544a4d8375ba74fba36b35122c9d40525821a11b3029847db13e3f3a07",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-devel-5.36.1-494.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-devel-5.36.1-494.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-devel-5.36.1-494.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-devel-5.36.1-494.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-devel-5.36.1-494.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "perl-diagnostics-0__1.39-494.fc37.x86_64",
        sha256 = "87ab7e30f8641bbbc5d92861333758f39adf5d3ae31783e147e8b3e76efdc7bf",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-diagnostics-1.39-494.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-diagnostics-1.39-494.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-diagnostics-1.39-494.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-diagnostics-1.39-494.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-diagnostics-1.39-494.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-doc-0__5.36.1-494.fc37.x86_64",
        sha256 = "d737312ea87d1f19b7643ff343a8aaa18999f97e4fd798d362252b2389f6cc3c",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-doc-5.36.1-494.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-doc-5.36.1-494.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-doc-5.36.1-494.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-doc-5.36.1-494.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-doc-5.36.1-494.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-encoding-4__3.00-492.fc37.x86_64",
        sha256 = "8d1716f56b0e97b7ae914f351dec3840512bd0f464e9dbb06356afd8ac7ab582",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-encoding-3.00-492.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-encoding-3.00-492.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-encoding-3.00-492.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-encoding-3.00-492.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-encoding-3.00-492.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "perl-encoding-warnings-0__0.13-494.fc37.x86_64",
        sha256 = "171cc2a67530182711725e78031ac262db85fda1c98cf144a190793797598d84",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-encoding-warnings-0.13-494.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-encoding-warnings-0.13-494.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-encoding-warnings-0.13-494.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-encoding-warnings-0.13-494.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-encoding-warnings-0.13-494.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-experimental-0__0.028-489.fc37.x86_64",
        sha256 = "e636ec425afc7fd8b16a282db226e73b504d6427974223d538c6fae879e24a38",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-experimental-0.028-489.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-experimental-0.028-489.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-experimental-0.028-489.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-experimental-0.028-489.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-experimental-0.028-489.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-fields-0__2.27-494.fc37.x86_64",
        sha256 = "09425523abe82f97125de70bf066b7ecc7ddaa146d3b5b2eaa9a3da62c57287e",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-fields-2.27-494.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-fields-2.27-494.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-fields-2.27-494.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-fields-2.27-494.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-fields-2.27-494.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-filetest-0__1.03-494.fc37.x86_64",
        sha256 = "cb5dd3b71087bbc897e43aa5121755ab3d4822793e2b6b54ae3f27633579ee4f",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-filetest-1.03-494.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-filetest-1.03-494.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-filetest-1.03-494.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-filetest-1.03-494.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-filetest-1.03-494.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-if-0__0.61.000-494.fc37.x86_64",
        sha256 = "29d096c3607e6b70bc3258344d1985ac82ed02a3673f914e2ecd1272637a9845",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-if-0.61.000-494.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-if-0.61.000-494.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-if-0.61.000-494.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-if-0.61.000-494.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-if-0.61.000-494.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-inc-latest-2__0.500-23.fc37.x86_64",
        sha256 = "f75cb2e3c2ae33d3a4efdb1ecbb75afa14104d5d84316b38460a67219ca44be3",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-inc-latest-0.500-23.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-inc-latest-0.500-23.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-inc-latest-0.500-23.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-inc-latest-0.500-23.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-inc-latest-0.500-23.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-interpreter-4__5.36.1-494.fc37.x86_64",
        sha256 = "af3d39a7eaaa963f36e5048ce0e3e89eee4dbf1700ef9a4431aa13e0c29f2042",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-interpreter-5.36.1-494.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-interpreter-5.36.1-494.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-interpreter-5.36.1-494.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-interpreter-5.36.1-494.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-interpreter-5.36.1-494.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "perl-less-0__0.03-494.fc37.x86_64",
        sha256 = "b5913adb7f05223584764b70a7820971a7dd38beedb61eb0254e3348b32ebda8",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-less-0.03-494.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-less-0.03-494.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-less-0.03-494.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-less-0.03-494.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-less-0.03-494.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-lib-0__0.65-494.fc37.x86_64",
        sha256 = "88bf4b3842fc6156d99dfd5c015bd2d018050f75ac31f5523f61e78e526c3a7d",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-lib-0.65-494.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-lib-0.65-494.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-lib-0.65-494.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-lib-0.65-494.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-lib-0.65-494.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "perl-libnet-0__3.14-490.fc37.x86_64",
        sha256 = "301ca6aa190bcc1a6994e0bb9c8c31f1fd75cbef32688aec805f756d46e3d7c9",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-libnet-3.14-490.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-libnet-3.14-490.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-libnet-3.14-490.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-libnet-3.14-490.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-libnet-3.14-490.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-libnetcfg-4__5.36.1-494.fc37.x86_64",
        sha256 = "40546e40ab2b16b0dfec939615b40e84c118b9425d73e8330a299f2644209b53",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-libnetcfg-5.36.1-494.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-libnetcfg-5.36.1-494.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-libnetcfg-5.36.1-494.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-libnetcfg-5.36.1-494.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-libnetcfg-5.36.1-494.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-libs-4__5.36.1-494.fc37.x86_64",
        sha256 = "eb8ba1cc4684b6c8bfe405c5574fe37032f9f8d9f019fc709b06d5fb56573928",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-libs-5.36.1-494.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-libs-5.36.1-494.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-libs-5.36.1-494.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-libs-5.36.1-494.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-libs-5.36.1-494.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "perl-local-lib-0__2.000029-3.fc37.x86_64",
        sha256 = "26865261e7b5c7a3a554fc183b910da91f582e4a4e68f5a8d40bd2a42ceb99df",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-local-lib-2.000029-3.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-local-lib-2.000029-3.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-local-lib-2.000029-3.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-local-lib-2.000029-3.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-local-lib-2.000029-3.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-locale-0__1.10-494.fc37.x86_64",
        sha256 = "fd82ed8982bb6eab6c0baf137755525e6abd2619474388af87cc8ca396b25d3e",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-locale-1.10-494.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-locale-1.10-494.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-locale-1.10-494.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-locale-1.10-494.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-locale-1.10-494.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-macros-4__5.36.1-494.fc37.x86_64",
        sha256 = "a4086025968af1b0572c8487949b09145b1f7005d16b726a5be576202b4d5332",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-macros-5.36.1-494.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-macros-5.36.1-494.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-macros-5.36.1-494.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-macros-5.36.1-494.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-macros-5.36.1-494.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-meta-notation-0__5.36.1-494.fc37.x86_64",
        sha256 = "faddd9fad3c3c8ec891dd83dc9dabf7b74700932a640ea908c174192ea6228e9",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-meta-notation-5.36.1-494.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-meta-notation-5.36.1-494.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-meta-notation-5.36.1-494.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-meta-notation-5.36.1-494.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-meta-notation-5.36.1-494.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-mro-0__1.26-494.fc37.x86_64",
        sha256 = "ecfdc15b4b885245dd3d6243b30a6a1a4c1c6bb787e12021101251085229eaf7",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-mro-1.26-494.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-mro-1.26-494.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-mro-1.26-494.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-mro-1.26-494.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-mro-1.26-494.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "perl-open-0__1.13-494.fc37.x86_64",
        sha256 = "dcc574c9aca8818070654a12d7bcd50c74eb1b11f624b834ee8bbdd4b1eddf9f",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-open-1.13-494.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-open-1.13-494.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-open-1.13-494.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-open-1.13-494.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-open-1.13-494.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-overload-0__1.35-494.fc37.x86_64",
        sha256 = "d59578a532f16931ee1aad20573b1a26eb6b7428a1015ae7e3225e3a2117a472",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-overload-1.35-494.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-overload-1.35-494.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-overload-1.35-494.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-overload-1.35-494.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-overload-1.35-494.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-overloading-0__0.02-494.fc37.x86_64",
        sha256 = "ba31ae4042ed10c2cf28b9c4806546ae5423f4edc5bdc6b4b929a0936be21086",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-overloading-0.02-494.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-overloading-0.02-494.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-overloading-0.02-494.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-overloading-0.02-494.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-overloading-0.02-494.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-parent-1__0.238-489.fc37.x86_64",
        sha256 = "202ebb9e6bf82022838cbe5814a9a38ec60e3d74864264bd48963a48032177df",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-parent-0.238-489.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-parent-0.238-489.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-parent-0.238-489.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-parent-0.238-489.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-parent-0.238-489.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-perlfaq-0__5.20210520-489.fc37.x86_64",
        sha256 = "e32806670a9e739869249d76a0f9c39bffbe10b0bf7e257ffc16912c763cc4b4",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-perlfaq-5.20210520-489.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-perlfaq-5.20210520-489.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-perlfaq-5.20210520-489.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-perlfaq-5.20210520-489.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-perlfaq-5.20210520-489.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-ph-0__5.36.1-494.fc37.x86_64",
        sha256 = "145e6b55abfb36e5efd1e36979fcf980a80092fe0d449695c336aa9a78a01261",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-ph-5.36.1-494.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-ph-5.36.1-494.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-ph-5.36.1-494.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-ph-5.36.1-494.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-ph-5.36.1-494.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "perl-podlators-1__4.14-489.fc37.x86_64",
        sha256 = "1dbfed6c41e81aa507a5a723fbed9407ab7a6a8fef7b054c742c1feadf388cba",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-podlators-4.14-489.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-podlators-4.14-489.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-podlators-4.14-489.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-podlators-4.14-489.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-podlators-4.14-489.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-sigtrap-0__1.10-494.fc37.x86_64",
        sha256 = "0a097f26b643419bc722aae341ae73014b1eab5d819790f1ad94fbdd5d7d6020",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-sigtrap-1.10-494.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-sigtrap-1.10-494.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-sigtrap-1.10-494.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-sigtrap-1.10-494.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-sigtrap-1.10-494.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-sort-0__2.05-494.fc37.x86_64",
        sha256 = "1a270572fa880eacced707342125e1222134aabdb7130219bb52a45ddb1864b1",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-sort-2.05-494.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-sort-2.05-494.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-sort-2.05-494.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-sort-2.05-494.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-sort-2.05-494.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-srpm-macros-0__1-46.fc37.x86_64",
        sha256 = "16a8fe47d15024db2da96e3a380f48be0a7b6f61264b9f81337bb785ad4f8bc4",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-srpm-macros-1-46.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-srpm-macros-1-46.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-srpm-macros-1-46.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-srpm-macros-1-46.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-srpm-macros-1-46.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-subs-0__1.04-494.fc37.x86_64",
        sha256 = "3cf1b173bdb927b84d84f18f9d9932befc41a57169b764c13927429eb463a10a",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-subs-1.04-494.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-subs-1.04-494.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-subs-1.04-494.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-subs-1.04-494.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-subs-1.04-494.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-threads-1__2.27-489.fc37.x86_64",
        sha256 = "ca68be5c4100cc4c5508b2e85e92771ae6cc5a85f877bd9e17490a670f2442fd",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-threads-2.27-489.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-threads-2.27-489.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-threads-2.27-489.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-threads-2.27-489.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-threads-2.27-489.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "perl-threads-shared-0__1.64-489.fc37.x86_64",
        sha256 = "9fc2e275aa70933e5d7e2065258371e430f43c149bb757519e2bef170d1abf83",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-threads-shared-1.64-489.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-threads-shared-1.64-489.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-threads-shared-1.64-489.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-threads-shared-1.64-489.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-threads-shared-1.64-489.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "perl-utils-0__5.36.1-494.fc37.x86_64",
        sha256 = "a94ff7df393cf528f2e16e20cb85c7354834c73421ae9f069532fd138ae31a81",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-utils-5.36.1-494.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-utils-5.36.1-494.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-utils-5.36.1-494.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-utils-5.36.1-494.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-utils-5.36.1-494.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-vars-0__1.05-494.fc37.x86_64",
        sha256 = "c94916a545a2cf5c903577e360051992458b163bfde7a0f8c61fd4e667d0064b",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-vars-1.05-494.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-vars-1.05-494.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-vars-1.05-494.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-vars-1.05-494.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-vars-1.05-494.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-version-8__0.99.29-490.fc37.x86_64",
        sha256 = "e39aee500f3b39ef1374943d2bd91fc6ac89ac10ce4025c580b79bdb90da14f2",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-version-0.99.29-490.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-version-0.99.29-490.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-version-0.99.29-490.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-version-0.99.29-490.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-version-0.99.29-490.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "perl-vmsish-0__1.04-494.fc37.x86_64",
        sha256 = "c79e6196629920c225eaddc25e1215c1fce820eab53afd6f2e7be7157119f758",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-vmsish-1.04-494.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-vmsish-1.04-494.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-vmsish-1.04-494.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-vmsish-1.04-494.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-vmsish-1.04-494.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "pixman-0__0.40.0-6.fc37.x86_64",
        sha256 = "131619876f2f68070ef4e178b5758474f54c577e5a6bf7a88746db54f0d0231f",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/pixman-0.40.0-6.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/pixman-0.40.0-6.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/pixman-0.40.0-6.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/pixman-0.40.0-6.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/pixman-0.40.0-6.fc37.x86_64.rpm",
        ],
    )

    rpm(
        name = "pkgconf-0__1.8.0-3.fc37.x86_64",
        sha256 = "778018594ab5bddc4432e53985b80e6c5a1a1ec1700d38b438848d485f5b357c",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/pkgconf-1.8.0-3.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/pkgconf-1.8.0-3.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/pkgconf-1.8.0-3.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/pkgconf-1.8.0-3.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/pkgconf-1.8.0-3.fc37.x86_64.rpm",
        ],
    )

    rpm(
        name = "pkgconf-m4-0__1.8.0-3.fc37.x86_64",
        sha256 = "dd0356475d0b9106b5a2d577db359aa0290fe6dd9eacea1b6e0cab816ff33566",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/pkgconf-m4-1.8.0-3.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/pkgconf-m4-1.8.0-3.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/pkgconf-m4-1.8.0-3.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/pkgconf-m4-1.8.0-3.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/pkgconf-m4-1.8.0-3.fc37.noarch.rpm",
        ],
    )

    rpm(
        name = "pkgconf-pkg-config-0__1.8.0-3.fc37.x86_64",
        sha256 = "d238b12c750b58ceebc80e25c2074bd929d3f232c1390677f33a94fdadb68f6a",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/pkgconf-pkg-config-1.8.0-3.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/pkgconf-pkg-config-1.8.0-3.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/pkgconf-pkg-config-1.8.0-3.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/pkgconf-pkg-config-1.8.0-3.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/pkgconf-pkg-config-1.8.0-3.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "policycoreutils-0__3.5-1.fc37.x86_64",
        sha256 = "199f28e3ecd24e0650ed7b8bc596f14e7a9cdef79c36989b710da1f557b347d9",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/policycoreutils-3.5-1.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/policycoreutils-3.5-1.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/policycoreutils-3.5-1.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/policycoreutils-3.5-1.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/policycoreutils-3.5-1.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "policycoreutils-python-utils-0__3.5-1.fc37.x86_64",
        sha256 = "6a9db2d2e789177cad01f4cf52bd9ac47562127869c26eaf091134770a7e20a3",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/policycoreutils-python-utils-3.5-1.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/policycoreutils-python-utils-3.5-1.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/policycoreutils-python-utils-3.5-1.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/policycoreutils-python-utils-3.5-1.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/policycoreutils-python-utils-3.5-1.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "popt-0__1.19-1.fc37.x86_64",
        sha256 = "e3c9a6a1611d967fbff4321b5b1ae54377fed22454298859108138c1f64b0c63",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/popt-1.19-1.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/popt-1.19-1.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/popt-1.19-1.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/popt-1.19-1.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/popt-1.19-1.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "protobuf-c-0__1.4.1-2.fc37.x86_64",
        sha256 = "46a9be44b3444815a0197dd85953bf87710d3ea3d8f9fbfff23068ca85885070",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/protobuf-c-1.4.1-2.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/protobuf-c-1.4.1-2.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/protobuf-c-1.4.1-2.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/protobuf-c-1.4.1-2.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/protobuf-c-1.4.1-2.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "pyproject-srpm-macros-0__1.9.0-1.fc37.x86_64",
        sha256 = "f284990a18e5a6f85eda4ab0b089e52f053eea68ce4d1a68db4550b47199b448",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/pyproject-srpm-macros-1.9.0-1.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/pyproject-srpm-macros-1.9.0-1.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/pyproject-srpm-macros-1.9.0-1.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/pyproject-srpm-macros-1.9.0-1.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/pyproject-srpm-macros-1.9.0-1.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "python-pip-wheel-0__22.2.2-3.fc37.x86_64",
        sha256 = "f7800b3f5acca7863bf47981258582728b74861c19b2bef38ae47efe3b042eb4",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/python-pip-wheel-22.2.2-3.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/python-pip-wheel-22.2.2-3.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/python-pip-wheel-22.2.2-3.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/python-pip-wheel-22.2.2-3.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/python-pip-wheel-22.2.2-3.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "python-setuptools-wheel-0__62.6.0-3.fc37.x86_64",
        sha256 = "71eb63ef6b25df748b6d2b66c3fba8c47b4a72dc7e1562fddb1073f5bd8af36a",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/python-setuptools-wheel-62.6.0-3.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/python-setuptools-wheel-62.6.0-3.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/python-setuptools-wheel-62.6.0-3.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/python-setuptools-wheel-62.6.0-3.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/python-setuptools-wheel-62.6.0-3.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "python-srpm-macros-0__3.11-6.fc37.x86_64",
        sha256 = "215f39543052fd3a73ec98750ad692b4fe0d66736da02dc7b2a0c9e16f399188",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/python-srpm-macros-3.11-6.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/python-srpm-macros-3.11-6.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/python-srpm-macros-3.11-6.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/python-srpm-macros-3.11-6.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/python-srpm-macros-3.11-6.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "python-unversioned-command-0__3.11.5-1.fc37.x86_64",
        sha256 = "0320c592389677ade74dbb8c4d3390e77e1a30857fa701c0ba1a344cc685b5ba",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/python-unversioned-command-3.11.5-1.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/python-unversioned-command-3.11.5-1.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/python-unversioned-command-3.11.5-1.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/python-unversioned-command-3.11.5-1.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/python-unversioned-command-3.11.5-1.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "python3-0__3.11.5-1.fc37.x86_64",
        sha256 = "9d67caa894f86665e7172c6c0e71e284a183850fd22f4723bce1d62466c3ac73",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/python3-3.11.5-1.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/python3-3.11.5-1.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/python3-3.11.5-1.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/python3-3.11.5-1.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/python3-3.11.5-1.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "python3-audit-0__3.1.2-1.fc37.x86_64",
        sha256 = "3ac574431a4e03f28dfc2f756393f159dd5e4ff868b24f687b7a34d9ba6ff597",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/python3-audit-3.1.2-1.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/python3-audit-3.1.2-1.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/python3-audit-3.1.2-1.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/python3-audit-3.1.2-1.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/python3-audit-3.1.2-1.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "python3-distro-0__1.7.0-3.fc37.x86_64",
        sha256 = "6427a8b877a51be140e06665b3680a911816c9eb581aa54ecf29c77f51b6e898",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/python3-distro-1.7.0-3.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/python3-distro-1.7.0-3.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/python3-distro-1.7.0-3.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/python3-distro-1.7.0-3.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/python3-distro-1.7.0-3.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "python3-libs-0__3.11.5-1.fc37.x86_64",
        sha256 = "c03cebe52be793cb1c00d3417d906252fcd4eb95ededf36a1e6664268667ff1d",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/python3-libs-3.11.5-1.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/python3-libs-3.11.5-1.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/python3-libs-3.11.5-1.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/python3-libs-3.11.5-1.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/python3-libs-3.11.5-1.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "python3-libselinux-0__3.5-1.fc37.x86_64",
        sha256 = "d6a8fff22472e9629ae8d61c678448215a4f3fb3314b26d0fde0dee70996b3a0",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/python3-libselinux-3.5-1.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/python3-libselinux-3.5-1.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/python3-libselinux-3.5-1.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/python3-libselinux-3.5-1.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/python3-libselinux-3.5-1.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "python3-libsemanage-0__3.5-2.fc37.x86_64",
        sha256 = "5b2a476b4b304c27e9c04c71cbd28b9f4d53d900fc84cff29677f9d0d259ec4b",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/python3-libsemanage-3.5-2.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/python3-libsemanage-3.5-2.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/python3-libsemanage-3.5-2.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/python3-libsemanage-3.5-2.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/python3-libsemanage-3.5-2.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "python3-policycoreutils-0__3.5-1.fc37.x86_64",
        sha256 = "e746e0a43f9a9d57b33964332f595490f73df5df751cc785c66b4b72d4633410",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/python3-policycoreutils-3.5-1.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/python3-policycoreutils-3.5-1.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/python3-policycoreutils-3.5-1.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/python3-policycoreutils-3.5-1.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/python3-policycoreutils-3.5-1.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "python3-pyparsing-0__3.0.9-2.fc37.x86_64",
        sha256 = "2688bcf1fd02d090e63417a9f51e89e6cfc7823a5a64f5dae4576a9bd7744c8a",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/python3-pyparsing-3.0.9-2.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/python3-pyparsing-3.0.9-2.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/python3-pyparsing-3.0.9-2.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/python3-pyparsing-3.0.9-2.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/python3-pyparsing-3.0.9-2.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "python3-setools-0__4.4.0-9.fc37.x86_64",
        sha256 = "b872895c9e0ddbebdf970a42f8c7a25d956fb7bbcdbb18aa4cf5515cff962c62",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/python3-setools-4.4.0-9.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/python3-setools-4.4.0-9.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/python3-setools-4.4.0-9.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/python3-setools-4.4.0-9.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/python3-setools-4.4.0-9.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "python3-setuptools-0__62.6.0-3.fc37.x86_64",
        sha256 = "8096786a1448b24095f14579b715e26a416274002412903da9da181f29945baa",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/python3-setuptools-62.6.0-3.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/python3-setuptools-62.6.0-3.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/python3-setuptools-62.6.0-3.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/python3-setuptools-62.6.0-3.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/p/python3-setuptools-62.6.0-3.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "qemu-common-2__7.0.0-15.fc37.x86_64",
        sha256 = "d9907c540e8cd0c94fe8d541b77f7a4c6173e546b63296e62dd02dfae50f4e1e",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/q/qemu-common-7.0.0-15.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/q/qemu-common-7.0.0-15.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/q/qemu-common-7.0.0-15.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/q/qemu-common-7.0.0-15.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/q/qemu-common-7.0.0-15.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "qemu-system-x86-core-2__7.0.0-15.fc37.x86_64",
        sha256 = "dbed7e4d2d9776063e0b8002a7dc5fe08247f95e5556881c31855d15882fcf15",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/q/qemu-system-x86-core-7.0.0-15.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/q/qemu-system-x86-core-7.0.0-15.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/q/qemu-system-x86-core-7.0.0-15.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/q/qemu-system-x86-core-7.0.0-15.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/q/qemu-system-x86-core-7.0.0-15.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "qt5-srpm-macros-0__5.15.9-1.fc37.x86_64",
        sha256 = "0815803ae72a90fa4ebaad878a75ba1bd80a1dc6691f2ab923373d82b1b430df",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/q/qt5-srpm-macros-5.15.9-1.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/q/qt5-srpm-macros-5.15.9-1.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/q/qt5-srpm-macros-5.15.9-1.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/q/qt5-srpm-macros-5.15.9-1.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/q/qt5-srpm-macros-5.15.9-1.fc37.noarch.rpm",
        ],
    )

    rpm(
        name = "readline-0__8.2-2.fc37.x86_64",
        sha256 = "0663e23dc42a7ce84f60f5f3154ba640460a0e5b7158459abf9d5d0986d69d06",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/r/readline-8.2-2.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/r/readline-8.2-2.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/r/readline-8.2-2.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/r/readline-8.2-2.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/r/readline-8.2-2.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "redhat-rpm-config-0__229-1.fc37.x86_64",
        sha256 = "9678d9f500321a9e69493e47bdb35dd7f9f51737b7fb5c368cdfe82293b6213a",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/r/redhat-rpm-config-229-1.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/r/redhat-rpm-config-229-1.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/r/redhat-rpm-config-229-1.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/r/redhat-rpm-config-229-1.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/r/redhat-rpm-config-229-1.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "rpm-0__4.18.1-2.fc37.x86_64",
        sha256 = "ac3915164eb1e4dd24a94fe7f77f5722c5b7e50ccb3bde29804ffe5c97ae49ae",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/r/rpm-4.18.1-2.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/r/rpm-4.18.1-2.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/r/rpm-4.18.1-2.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/r/rpm-4.18.1-2.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/r/rpm-4.18.1-2.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "rpm-libs-0__4.18.1-2.fc37.x86_64",
        sha256 = "e1a55d458afb2e5c30ceb704ec7424da64a7b7bb872249522590799a5aa786df",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/r/rpm-libs-4.18.1-2.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/r/rpm-libs-4.18.1-2.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/r/rpm-libs-4.18.1-2.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/r/rpm-libs-4.18.1-2.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/r/rpm-libs-4.18.1-2.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "rpm-plugin-selinux-0__4.18.1-2.fc37.x86_64",
        sha256 = "f98bfc30379e94251f50ed2542cf7b3e3c313478f2ab2f6319bbac6a2e945bc3",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/r/rpm-plugin-selinux-4.18.1-2.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/r/rpm-plugin-selinux-4.18.1-2.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/r/rpm-plugin-selinux-4.18.1-2.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/r/rpm-plugin-selinux-4.18.1-2.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/r/rpm-plugin-selinux-4.18.1-2.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "rpmautospec-rpm-macros-0__0.3.5-1.fc37.x86_64",
        sha256 = "75099655c00a959dca4de2b0e7661078991d1d26972563b7ac4588c9ec61e083",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/r/rpmautospec-rpm-macros-0.3.5-1.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/r/rpmautospec-rpm-macros-0.3.5-1.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/r/rpmautospec-rpm-macros-0.3.5-1.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/r/rpmautospec-rpm-macros-0.3.5-1.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/r/rpmautospec-rpm-macros-0.3.5-1.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "rsync-0__3.2.7-1.fc37.x86_64",
        sha256 = "82e02be388b28292136440b3a11e72db8cdc615b2adc8a4073dd3721ae383ef1",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/r/rsync-3.2.7-1.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/r/rsync-3.2.7-1.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/r/rsync-3.2.7-1.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/r/rsync-3.2.7-1.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/r/rsync-3.2.7-1.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "rust-srpm-macros-0__24-4.fc37.x86_64",
        sha256 = "5b2929babe6b0cac40cf83d055480f7b56a99ae96c9874119d9c3bf213063317",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/r/rust-srpm-macros-24-4.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/r/rust-srpm-macros-24-4.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/r/rust-srpm-macros-24-4.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/r/rust-srpm-macros-24-4.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/r/rust-srpm-macros-24-4.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "seabios-bin-0__1.16.2-1.fc37.x86_64",
        sha256 = "095db21fd4f608cd4b195da7ab4c868fa512c2e3ea331adc1ed3bb2f4dee835b",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/s/seabios-bin-1.16.2-1.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/s/seabios-bin-1.16.2-1.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/s/seabios-bin-1.16.2-1.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/s/seabios-bin-1.16.2-1.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/s/seabios-bin-1.16.2-1.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "seavgabios-bin-0__1.16.2-1.fc37.x86_64",
        sha256 = "b666c8034df6c0584f13ed18d7cbe22e01dbbde1e14f81c63fdf8236aa8f029d",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/s/seavgabios-bin-1.16.2-1.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/s/seavgabios-bin-1.16.2-1.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/s/seavgabios-bin-1.16.2-1.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/s/seavgabios-bin-1.16.2-1.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/s/seavgabios-bin-1.16.2-1.fc37.noarch.rpm",
        ],
    )

    rpm(
        name = "sed-0__4.8-11.fc37.x86_64",
        sha256 = "231e782077862f4abecf025aa254a9c391a950490ae856261dcfd229863ac80f",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/s/sed-4.8-11.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/s/sed-4.8-11.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/s/sed-4.8-11.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/s/sed-4.8-11.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/s/sed-4.8-11.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "selinux-policy-0__37.22-1.fc37.x86_64",
        sha256 = "715397068619030c805adafcaee3b524a7807b4b69ed76daf3e07fb3d37841f9",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/s/selinux-policy-37.22-1.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/s/selinux-policy-37.22-1.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/s/selinux-policy-37.22-1.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/s/selinux-policy-37.22-1.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/s/selinux-policy-37.22-1.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "selinux-policy-minimum-0__37.22-1.fc37.x86_64",
        sha256 = "8b6c61cc1f96a21c48482833586aad47b5fdf3eba5c95300a00e25fbfe589051",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/s/selinux-policy-minimum-37.22-1.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/s/selinux-policy-minimum-37.22-1.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/s/selinux-policy-minimum-37.22-1.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/s/selinux-policy-minimum-37.22-1.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/s/selinux-policy-minimum-37.22-1.fc37.noarch.rpm",
        ],
    )

    rpm(
        name = "setup-0__2.14.1-2.fc37.x86_64",
        sha256 = "15d72b2a44f403b3a7ee9138820a8ce7584f954aeafbb43b1251621bca26f785",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/s/setup-2.14.1-2.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/s/setup-2.14.1-2.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/s/setup-2.14.1-2.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/s/setup-2.14.1-2.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/s/setup-2.14.1-2.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "sgabios-bin-1__0.20180715git-9.fc37.x86_64",
        sha256 = "7b18cb14de4338aa5fa937f19a44ac5accceacb18009e5b33afd22755cb7955a",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/s/sgabios-bin-0.20180715git-9.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/s/sgabios-bin-0.20180715git-9.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/s/sgabios-bin-0.20180715git-9.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/s/sgabios-bin-0.20180715git-9.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/s/sgabios-bin-0.20180715git-9.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "shadow-utils-2__4.12.3-6.fc37.x86_64",
        sha256 = "b85714adf81c11f87479305bb324ad15c0cf92ca0f5560cef5c99995f1e9c942",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/s/shadow-utils-4.12.3-6.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/s/shadow-utils-4.12.3-6.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/s/shadow-utils-4.12.3-6.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/s/shadow-utils-4.12.3-6.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/s/shadow-utils-4.12.3-6.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "snappy-0__1.1.9-5.fc37.x86_64",
        sha256 = "46504f3ad77433138805882361af9245a26a74e2b0984f4b35b3509a3b2f91bf",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/s/snappy-1.1.9-5.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/s/snappy-1.1.9-5.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/s/snappy-1.1.9-5.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/s/snappy-1.1.9-5.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/s/snappy-1.1.9-5.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "sqlite-libs-0__3.40.0-1.fc37.x86_64",
        sha256 = "4d1603de146f9bbe90810100df0afa2efe32e13cc86ed42e32528bc50b8f03dd",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/s/sqlite-libs-3.40.0-1.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/s/sqlite-libs-3.40.0-1.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/s/sqlite-libs-3.40.0-1.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/s/sqlite-libs-3.40.0-1.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/s/sqlite-libs-3.40.0-1.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "swtpm-0__0.7.3-2.20220427gitf2268ee.fc37.x86_64",
        sha256 = "4dd0ae80effe40033c02e3d2b9c4f4824c4faa7f58d7e3ba8c946316dc578ba5",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/s/swtpm-0.7.3-2.20220427gitf2268ee.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/s/swtpm-0.7.3-2.20220427gitf2268ee.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/s/swtpm-0.7.3-2.20220427gitf2268ee.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/s/swtpm-0.7.3-2.20220427gitf2268ee.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/s/swtpm-0.7.3-2.20220427gitf2268ee.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "swtpm-libs-0__0.7.3-2.20220427gitf2268ee.fc37.x86_64",
        sha256 = "3b28d0e464f9aefb3c109c56508740c8958a4475235c75ed996f0b80e8caeb0f",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/s/swtpm-libs-0.7.3-2.20220427gitf2268ee.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/s/swtpm-libs-0.7.3-2.20220427gitf2268ee.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/s/swtpm-libs-0.7.3-2.20220427gitf2268ee.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/s/swtpm-libs-0.7.3-2.20220427gitf2268ee.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/s/swtpm-libs-0.7.3-2.20220427gitf2268ee.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "swtpm-tools-0__0.7.3-2.20220427gitf2268ee.fc37.x86_64",
        sha256 = "4e6e001a6c6f8793d4b7abd824396ce7560c9524c636f75346f4721461082d1f",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/s/swtpm-tools-0.7.3-2.20220427gitf2268ee.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/s/swtpm-tools-0.7.3-2.20220427gitf2268ee.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/s/swtpm-tools-0.7.3-2.20220427gitf2268ee.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/s/swtpm-tools-0.7.3-2.20220427gitf2268ee.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/s/swtpm-tools-0.7.3-2.20220427gitf2268ee.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "systemd-0__251.14-2.fc37.x86_64",
        sha256 = "ec407a50153db3d466ed6063bc9274ce73192197e08e22ee88f7c364036ffc66",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/s/systemd-251.14-2.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/s/systemd-251.14-2.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/s/systemd-251.14-2.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/s/systemd-251.14-2.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/s/systemd-251.14-2.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "systemd-libs-0__251.14-2.fc37.x86_64",
        sha256 = "37934c2adce2bf1559a8b74a9e6c62510acb0d0653ab5d80b2676f69388beee5",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/s/systemd-libs-251.14-2.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/s/systemd-libs-251.14-2.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/s/systemd-libs-251.14-2.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/s/systemd-libs-251.14-2.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/s/systemd-libs-251.14-2.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "systemd-pam-0__251.14-2.fc37.x86_64",
        sha256 = "c3f416c1d6cb05491f61e4f9cf2e976b0ccda4cb3eebc2c1b6dfa633e876bbef",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/s/systemd-pam-251.14-2.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/s/systemd-pam-251.14-2.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/s/systemd-pam-251.14-2.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/s/systemd-pam-251.14-2.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/s/systemd-pam-251.14-2.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "systemtap-sdt-devel-0__4.9-2.fc37.x86_64",
        sha256 = "0e4da701910ce174353603a75bf89e0bead813b7ee0aab43679a6a2ebb457074",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/s/systemtap-sdt-devel-4.9-2.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/s/systemtap-sdt-devel-4.9-2.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/s/systemtap-sdt-devel-4.9-2.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/s/systemtap-sdt-devel-4.9-2.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/s/systemtap-sdt-devel-4.9-2.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "tar-2__1.34-6.fc37.x86_64",
        sha256 = "7e7ff9824621df916333c7a85656671d122f761a75a74f6e442d11474455ffed",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/t/tar-1.34-6.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/t/tar-1.34-6.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/t/tar-1.34-6.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/t/tar-1.34-6.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/t/tar-1.34-6.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "trousers-0__0.3.15-7.fc37.x86_64",
        sha256 = "9ec34885483cd25c7ae39b9e5b0af020f6db54123cdc3e38d898badbafb8ca43",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/t/trousers-0.3.15-7.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/t/trousers-0.3.15-7.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/t/trousers-0.3.15-7.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/t/trousers-0.3.15-7.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/t/trousers-0.3.15-7.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "trousers-lib-0__0.3.15-7.fc37.x86_64",
        sha256 = "b33af58d16302786d9b793c4f780aeb3b4d96d944868a998eecdcc37e71cfc50",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/t/trousers-lib-0.3.15-7.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/t/trousers-lib-0.3.15-7.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/t/trousers-lib-0.3.15-7.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/t/trousers-lib-0.3.15-7.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/t/trousers-lib-0.3.15-7.fc37.x86_64.rpm",
        ],
    )

    rpm(
        name = "tzdata-0__2023c-1.fc37.x86_64",
        sha256 = "a7116693b8cc1a857e8a515e6196e2fd3272e56d7c844b65c286d931ad6c3631",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/t/tzdata-2023c-1.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/t/tzdata-2023c-1.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/t/tzdata-2023c-1.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/t/tzdata-2023c-1.fc37.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/t/tzdata-2023c-1.fc37.noarch.rpm",
        ],
    )
    rpm(
        name = "unbound-libs-0__1.17.1-1.fc37.x86_64",
        sha256 = "13f29a4066dde4c0b48de7676275972bd05d8156270e63dadefe9dc2dac82a43",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/u/unbound-libs-1.17.1-1.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/u/unbound-libs-1.17.1-1.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/u/unbound-libs-1.17.1-1.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/u/unbound-libs-1.17.1-1.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/u/unbound-libs-1.17.1-1.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "unzip-0__6.0-58.fc37.x86_64",
        sha256 = "463eb69bf857dc76ce47adde2b02b6f9ce5857c1b897902e09fae9f75b690a4d",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/u/unzip-6.0-58.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/u/unzip-6.0-58.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/u/unzip-6.0-58.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/u/unzip-6.0-58.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/u/unzip-6.0-58.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "util-linux-0__2.38.1-1.fc37.x86_64",
        sha256 = "23f052850cd509743fae6089181a124ee65c2783d6d15f61ffbae1272f5f67ef",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/u/util-linux-2.38.1-1.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/u/util-linux-2.38.1-1.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/u/util-linux-2.38.1-1.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/u/util-linux-2.38.1-1.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/u/util-linux-2.38.1-1.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "util-linux-core-0__2.38.1-1.fc37.x86_64",
        sha256 = "f87ad8fc18f4da254966cc6f99b533dc8125e1ec0eaefd5f89a6b6398cb13a34",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/u/util-linux-core-2.38.1-1.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/u/util-linux-core-2.38.1-1.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/u/util-linux-core-2.38.1-1.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/u/util-linux-core-2.38.1-1.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/u/util-linux-core-2.38.1-1.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "xen-libs-0__4.16.5-1.fc37.x86_64",
        sha256 = "9d8b86f64f8d22a6d94351716f2ff951877401dbc7f3a9a7cbf49a793c708b8e",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/x/xen-libs-4.16.5-1.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/x/xen-libs-4.16.5-1.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/x/xen-libs-4.16.5-1.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/x/xen-libs-4.16.5-1.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/x/xen-libs-4.16.5-1.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "xen-licenses-0__4.16.5-1.fc37.x86_64",
        sha256 = "670c3353dfeaa51ebb75708f541c0fb67d118a30b973f5ee48aede48cc04206a",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/x/xen-licenses-4.16.5-1.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/x/xen-licenses-4.16.5-1.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/x/xen-licenses-4.16.5-1.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/x/xen-licenses-4.16.5-1.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/x/xen-licenses-4.16.5-1.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "xxhash-libs-0__0.8.2-1.fc37.x86_64",
        sha256 = "5ca1d21269d4a568103faef3eb79dd29a3e60d69eeae891355521e1725848d07",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/x/xxhash-libs-0.8.2-1.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/x/xxhash-libs-0.8.2-1.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/x/xxhash-libs-0.8.2-1.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/x/xxhash-libs-0.8.2-1.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/x/xxhash-libs-0.8.2-1.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "xz-0__5.4.1-1.fc37.x86_64",
        sha256 = "7af1096450d0d76dcd5666e31736f18ff44de9908f2e87d89be88592b176c643",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/x/xz-5.4.1-1.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/x/xz-5.4.1-1.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/x/xz-5.4.1-1.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/x/xz-5.4.1-1.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/x/xz-5.4.1-1.fc37.x86_64.rpm",
        ],
    )

    rpm(
        name = "xz-libs-0__5.4.1-1.fc37.x86_64",
        sha256 = "8c06eef8dd28d6dc1406e65e4eb8ee3db359cf6624729be4e426f6b01c4117fd",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/x/xz-libs-5.4.1-1.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/x/xz-libs-5.4.1-1.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/x/xz-libs-5.4.1-1.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/x/xz-libs-5.4.1-1.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/x/xz-libs-5.4.1-1.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "yajl-0__2.1.0-21.fc37.x86_64",
        sha256 = "af829c09f1a4f20e69872f6aad1c0b9dfc8666b2895375358f75149c079cd74f",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/y/yajl-2.1.0-21.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/y/yajl-2.1.0-21.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/y/yajl-2.1.0-21.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/y/yajl-2.1.0-21.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/37/Everything/x86_64/Packages/y/yajl-2.1.0-21.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "zip-0__3.0-33.fc37.x86_64",
        sha256 = "c3be027dae7eeec54aeac1e64d6f969641fa8ff4ef9273b21337112fad7d9974",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/z/zip-3.0-33.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/z/zip-3.0-33.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/z/zip-3.0-33.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/z/zip-3.0-33.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/z/zip-3.0-33.fc37.x86_64.rpm",
        ],
    )

    rpm(
        name = "zlib-0__1.2.12-5.fc37.x86_64",
        sha256 = "7b0eda1ad9e9a06e61d9fe41e5e4e0fbdc8427bc252f06a7d29cd7ba81a71a70",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/z/zlib-1.2.12-5.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/z/zlib-1.2.12-5.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/z/zlib-1.2.12-5.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/z/zlib-1.2.12-5.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/z/zlib-1.2.12-5.fc37.x86_64.rpm",
        ],
    )
    rpm(
        name = "zlib-devel-0__1.2.12-5.fc37.x86_64",
        sha256 = "9eee35ba6c098c1b47d16920e38ce4ed6d4e63775bb3fa74aaef9ddff36c2401",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/z/zlib-devel-1.2.12-5.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/z/zlib-devel-1.2.12-5.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/z/zlib-devel-1.2.12-5.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/z/zlib-devel-1.2.12-5.fc37.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/37/Everything/x86_64/os/Packages/z/zlib-devel-1.2.12-5.fc37.x86_64.rpm",
        ],
    )
