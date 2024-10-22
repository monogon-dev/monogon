load("@bazeldnf//:deps.bzl", "rpm")

def sandbox_dependencies():
    rpm(
        name = "acpica-tools-0__20220331-8.fc40.x86_64",
        sha256 = "34bb1ea2cfd28d788de1ada56b1583dce257841bc2d72d74ac13c95d3215ac83",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/a/acpica-tools-20220331-8.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/a/acpica-tools-20220331-8.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/a/acpica-tools-20220331-8.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/a/acpica-tools-20220331-8.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/a/acpica-tools-20220331-8.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "alternatives-0__1.27-1.fc40.x86_64",
        sha256 = "ac860c52abbc65af5835d1bd97400c531a5635d39bc1d68e36a1fe54863385ea",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/a/alternatives-1.27-1.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/a/alternatives-1.27-1.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/a/alternatives-1.27-1.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/a/alternatives-1.27-1.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/a/alternatives-1.27-1.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "ansible-srpm-macros-0__1-16.fc40.x86_64",
        sha256 = "a221968063ee17b8d4ee3e7013d40b2789638a76ce8e94ebd15d694f6b48b4bd",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/a/ansible-srpm-macros-1-16.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/a/ansible-srpm-macros-1-16.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/a/ansible-srpm-macros-1-16.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/a/ansible-srpm-macros-1-16.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/a/ansible-srpm-macros-1-16.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "audit-libs-0__4.0.2-1.fc40.x86_64",
        sha256 = "f4ed40457780c13bebf84c1cf8981550da7e0e728e80250aed179eda8915bc7f",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/a/audit-libs-4.0.2-1.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/a/audit-libs-4.0.2-1.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/a/audit-libs-4.0.2-1.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/a/audit-libs-4.0.2-1.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/a/audit-libs-4.0.2-1.fc40.x86_64.rpm",
        ],
    )

    rpm(
        name = "authselect-0__1.5.0-5.fc40.x86_64",
        sha256 = "0fe4ed8770711ede2fcec43c4545b62461a24f03b3aa836d0e7071f4436e26f1",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/a/authselect-1.5.0-5.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/a/authselect-1.5.0-5.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/a/authselect-1.5.0-5.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/a/authselect-1.5.0-5.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/a/authselect-1.5.0-5.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "authselect-libs-0__1.5.0-5.fc40.x86_64",
        sha256 = "db7d946583f2a91a3301d964a5adc7afb1620e0f72c9a9033ae3a4cfc2f332ad",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/a/authselect-libs-1.5.0-5.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/a/authselect-libs-1.5.0-5.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/a/authselect-libs-1.5.0-5.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/a/authselect-libs-1.5.0-5.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/a/authselect-libs-1.5.0-5.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "basesystem-0__11-20.fc40.x86_64",
        sha256 = "6404b1028262aeaf3e083f08959969abea1301f7f5e8610492cf900b3d13d5db",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/b/basesystem-11-20.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/b/basesystem-11-20.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/b/basesystem-11-20.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/b/basesystem-11-20.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/b/basesystem-11-20.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "bash-0__5.2.26-3.fc40.x86_64",
        sha256 = "156e073308cb28a5a699d6ffafc71cbd28487628fd05471e1978e4b9a5c7a802",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/b/bash-5.2.26-3.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/b/bash-5.2.26-3.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/b/bash-5.2.26-3.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/b/bash-5.2.26-3.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/b/bash-5.2.26-3.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "bc-0__1.07.1-21.fc40.x86_64",
        sha256 = "67bc7eabdc731caf31e57906bc1348754b911313113e69103925c975c9054c4c",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/b/bc-1.07.1-21.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/b/bc-1.07.1-21.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/b/bc-1.07.1-21.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/b/bc-1.07.1-21.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/b/bc-1.07.1-21.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "binutils-0__2.41-37.fc40.x86_64",
        sha256 = "619340f90a77ad1f0f919826d0e2423a10e8b5aea3957b00393c75495997d125",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/b/binutils-2.41-37.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/b/binutils-2.41-37.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/b/binutils-2.41-37.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/b/binutils-2.41-37.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/b/binutils-2.41-37.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "binutils-gold-0__2.41-37.fc40.x86_64",
        sha256 = "f1d034d37740b0180f6b25cca42e6a82d8ad14513bf6158eaff96c8a0c6538e9",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/b/binutils-gold-2.41-37.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/b/binutils-gold-2.41-37.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/b/binutils-gold-2.41-37.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/b/binutils-gold-2.41-37.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/b/binutils-gold-2.41-37.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "bison-0__3.8.2-7.fc40.x86_64",
        sha256 = "347ab37d026999bf8d8bd7bab5e33d4d18dcbb0597c356864bd7d4d191297163",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/b/bison-3.8.2-7.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/b/bison-3.8.2-7.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/b/bison-3.8.2-7.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/b/bison-3.8.2-7.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/b/bison-3.8.2-7.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "bzip2-libs-0__1.0.8-18.fc40.x86_64",
        sha256 = "68a43532d10187888788625d0b6c2224ba95804280eddf2636e5ef700607e7d0",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/b/bzip2-libs-1.0.8-18.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/b/bzip2-libs-1.0.8-18.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/b/bzip2-libs-1.0.8-18.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/b/bzip2-libs-1.0.8-18.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/b/bzip2-libs-1.0.8-18.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "ca-certificates-0__2024.2.69_v8.0.401-1.0.fc40.x86_64",
        sha256 = "1afcf80d5e7b22ee512ec9f24b4f2b148888ef95af3486cf48f2204c3406b12d",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/c/ca-certificates-2024.2.69_v8.0.401-1.0.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/c/ca-certificates-2024.2.69_v8.0.401-1.0.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/c/ca-certificates-2024.2.69_v8.0.401-1.0.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/c/ca-certificates-2024.2.69_v8.0.401-1.0.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/c/ca-certificates-2024.2.69_v8.0.401-1.0.fc40.noarch.rpm",
        ],
    )

    rpm(
        name = "capstone-0__5.0.1-3.fc40.x86_64",
        sha256 = "33e75316755bccb0410019dfe42c0e8f0c5eab10abb328b8160c13343cf04d23",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/c/capstone-5.0.1-3.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/c/capstone-5.0.1-3.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/c/capstone-5.0.1-3.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/c/capstone-5.0.1-3.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/c/capstone-5.0.1-3.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "checkpolicy-0__3.7-2.fc40.x86_64",
        sha256 = "96702c5e5a8a53efea7f0c25b05ed8ff8fd7022280707c2d0c82f1d40edc0064",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/c/checkpolicy-3.7-2.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/c/checkpolicy-3.7-2.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/c/checkpolicy-3.7-2.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/c/checkpolicy-3.7-2.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/c/checkpolicy-3.7-2.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "clang-0__18.1.8-1.fc40.x86_64",
        sha256 = "72eb2e7348f625a2d00232f2f3ed3a12783694156b1d65d47d7fbdabc3ca618a",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/c/clang-18.1.8-1.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/c/clang-18.1.8-1.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/c/clang-18.1.8-1.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/c/clang-18.1.8-1.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/c/clang-18.1.8-1.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "clang-libs-0__18.1.8-1.fc40.x86_64",
        sha256 = "0f69f1b1865b37481a2fb0fa761ec2d2b421d4b7378f5e967e2f466b0b5c3477",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/c/clang-libs-18.1.8-1.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/c/clang-libs-18.1.8-1.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/c/clang-libs-18.1.8-1.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/c/clang-libs-18.1.8-1.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/c/clang-libs-18.1.8-1.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "clang-resource-filesystem-0__18.1.8-1.fc40.x86_64",
        sha256 = "eed9d12d800d7e2d8d5d81fa9f90e11095e21f8f29e1be219a95b50558295db9",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/c/clang-resource-filesystem-18.1.8-1.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/c/clang-resource-filesystem-18.1.8-1.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/c/clang-resource-filesystem-18.1.8-1.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/c/clang-resource-filesystem-18.1.8-1.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/c/clang-resource-filesystem-18.1.8-1.fc40.noarch.rpm",
        ],
    )

    rpm(
        name = "cmake-filesystem-0__3.28.2-1.fc40.x86_64",
        sha256 = "ab937a9fd0b9b27ce34c4fe4779e357706ec0c8fefccc4f899853aa16733f526",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/c/cmake-filesystem-3.28.2-1.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/c/cmake-filesystem-3.28.2-1.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/c/cmake-filesystem-3.28.2-1.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/c/cmake-filesystem-3.28.2-1.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/c/cmake-filesystem-3.28.2-1.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "coreutils-single-0__9.4-8.fc40.x86_64",
        sha256 = "e74a792e74d8467510b859d16927bc951484bee8d3a141795e7dc8cc1b34c183",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/c/coreutils-single-9.4-8.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/c/coreutils-single-9.4-8.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/c/coreutils-single-9.4-8.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/c/coreutils-single-9.4-8.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/c/coreutils-single-9.4-8.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "cpp-0__14.2.1-3.fc40.x86_64",
        sha256 = "7ee02d77a5ef26abd7683e318f9adc68b9b1ff7ba3c5119b3aeb8b1e62aecb2a",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/c/cpp-14.2.1-3.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/c/cpp-14.2.1-3.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/c/cpp-14.2.1-3.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/c/cpp-14.2.1-3.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/c/cpp-14.2.1-3.fc40.x86_64.rpm",
        ],
    )

    rpm(
        name = "cracklib-0__2.9.11-5.fc40.x86_64",
        sha256 = "ea1f43ef9a4b02a9c66726ee386f090145696fb93dff80d593ac82126f8037ec",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/c/cracklib-2.9.11-5.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/c/cracklib-2.9.11-5.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/c/cracklib-2.9.11-5.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/c/cracklib-2.9.11-5.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/c/cracklib-2.9.11-5.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "crypto-policies-0__20240725-1.git28d3e2d.fc40.x86_64",
        sha256 = "2469a287d9fe6ea5f4aa0686fa5f223c14505218743230afbd322fdd90b1d396",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/c/crypto-policies-20240725-1.git28d3e2d.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/c/crypto-policies-20240725-1.git28d3e2d.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/c/crypto-policies-20240725-1.git28d3e2d.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/c/crypto-policies-20240725-1.git28d3e2d.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/c/crypto-policies-20240725-1.git28d3e2d.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "curl-0__8.6.0-10.fc40.x86_64",
        sha256 = "20b0f2923feae4c2f1d339e959d3f03d81f8ca985faa05872377b827d6f30467",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/c/curl-8.6.0-10.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/c/curl-8.6.0-10.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/c/curl-8.6.0-10.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/c/curl-8.6.0-10.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/c/curl-8.6.0-10.fc40.x86_64.rpm",
        ],
    )

    rpm(
        name = "cyrus-sasl-lib-0__2.1.28-19.fc40.x86_64",
        sha256 = "0dff67dfeca59cb68cadafe8d9909b88dfaa2fc0a9a4426352f66a5fe351fbe3",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/c/cyrus-sasl-lib-2.1.28-19.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/c/cyrus-sasl-lib-2.1.28-19.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/c/cyrus-sasl-lib-2.1.28-19.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/c/cyrus-sasl-lib-2.1.28-19.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/c/cyrus-sasl-lib-2.1.28-19.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "daxctl-libs-0__79-1.fc40.x86_64",
        sha256 = "3f7bc7e46ffb64b6724a9d2744003d1f707fc99c43f591b47bdb9c8a64395f94",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/d/daxctl-libs-79-1.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/d/daxctl-libs-79-1.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/d/daxctl-libs-79-1.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/d/daxctl-libs-79-1.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/d/daxctl-libs-79-1.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "dbus-1__1.14.10-3.fc40.x86_64",
        sha256 = "19197df26f76af5e78bd1e3ad2f777bea071eef6dfec1219f6b8ee3c80e10193",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/d/dbus-1.14.10-3.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/d/dbus-1.14.10-3.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/d/dbus-1.14.10-3.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/d/dbus-1.14.10-3.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/d/dbus-1.14.10-3.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "dbus-broker-0__36-2.fc40.x86_64",
        sha256 = "84ca6055aa354df549fdc78d6d9df692ed4d12c14a489a6d2ce844b5f225a502",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/d/dbus-broker-36-2.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/d/dbus-broker-36-2.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/d/dbus-broker-36-2.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/d/dbus-broker-36-2.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/d/dbus-broker-36-2.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "dbus-common-1__1.14.10-3.fc40.x86_64",
        sha256 = "81bade4072aca4f5d22be29a916d9d0cfc9262a6c5d92ddfe750f7b8bf03f7c9",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/d/dbus-common-1.14.10-3.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/d/dbus-common-1.14.10-3.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/d/dbus-common-1.14.10-3.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/d/dbus-common-1.14.10-3.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/d/dbus-common-1.14.10-3.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "diffutils-0__3.10-5.fc40.x86_64",
        sha256 = "6913a547250df04ec388b96b7512977a25ab2fca62ed4345c3a9fc8782ce659f",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/d/diffutils-3.10-5.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/d/diffutils-3.10-5.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/d/diffutils-3.10-5.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/d/diffutils-3.10-5.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/d/diffutils-3.10-5.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "dwz-0__0.15-6.fc40.x86_64",
        sha256 = "c801f65c70a3ae8d5ddbc8d82fe6f78398ea29c7bab40612a4ddddd93c398aeb",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/d/dwz-0.15-6.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/d/dwz-0.15-6.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/d/dwz-0.15-6.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/d/dwz-0.15-6.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/d/dwz-0.15-6.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "e2fsprogs-libs-0__1.47.0-5.fc40.x86_64",
        sha256 = "8476fda117e3cb808129ddc2f975069685a8c7875ee04c3dafa6ceed948a2628",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/e/e2fsprogs-libs-1.47.0-5.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/e/e2fsprogs-libs-1.47.0-5.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/e/e2fsprogs-libs-1.47.0-5.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/e/e2fsprogs-libs-1.47.0-5.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/e/e2fsprogs-libs-1.47.0-5.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "ed-0__1.20.2-1.fc40.x86_64",
        sha256 = "7eb1ab7808024bb294a8ca0730beefec8b9009e1289b844d2880ea5a1e4e0ec0",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/e/ed-1.20.2-1.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/e/ed-1.20.2-1.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/e/ed-1.20.2-1.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/e/ed-1.20.2-1.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/e/ed-1.20.2-1.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "edk2-ovmf-0__20240813-1.fc40.x86_64",
        sha256 = "26f300cc6d58d5b68bd676b7304fc409f8c5e4f7e40d3eb44420163e8bdda61f",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/e/edk2-ovmf-20240813-1.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/e/edk2-ovmf-20240813-1.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/e/edk2-ovmf-20240813-1.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/e/edk2-ovmf-20240813-1.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/e/edk2-ovmf-20240813-1.fc40.noarch.rpm",
        ],
    )

    rpm(
        name = "efi-srpm-macros-0__5-11.fc40.x86_64",
        sha256 = "34ed8bd59f9b299975871ebce1d15208cd66a4383f46a4f0d75e01303bacac2c",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/e/efi-srpm-macros-5-11.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/e/efi-srpm-macros-5-11.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/e/efi-srpm-macros-5-11.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/e/efi-srpm-macros-5-11.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/e/efi-srpm-macros-5-11.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "elfutils-debuginfod-client-0__0.191-4.fc40.x86_64",
        sha256 = "877c66844c68044b2a29b5d7465eb97f429e9f38b56ebaa16d766c0979e93a80",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/e/elfutils-debuginfod-client-0.191-4.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/e/elfutils-debuginfod-client-0.191-4.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/e/elfutils-debuginfod-client-0.191-4.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/e/elfutils-debuginfod-client-0.191-4.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/e/elfutils-debuginfod-client-0.191-4.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "elfutils-default-yama-scope-0__0.191-4.fc40.x86_64",
        sha256 = "3fbe1afd014386a436a25205d6727475a8f1107be734dd92fc40c3d5e0e5971d",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/e/elfutils-default-yama-scope-0.191-4.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/e/elfutils-default-yama-scope-0.191-4.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/e/elfutils-default-yama-scope-0.191-4.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/e/elfutils-default-yama-scope-0.191-4.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/e/elfutils-default-yama-scope-0.191-4.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "elfutils-libelf-0__0.191-4.fc40.x86_64",
        sha256 = "c1eca14924981b987f9b17c01a97511d641f49ac6b2b0f2d8e83563343932302",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/e/elfutils-libelf-0.191-4.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/e/elfutils-libelf-0.191-4.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/e/elfutils-libelf-0.191-4.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/e/elfutils-libelf-0.191-4.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/e/elfutils-libelf-0.191-4.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "elfutils-libelf-devel-0__0.191-4.fc40.x86_64",
        sha256 = "0a322715aefb4c6c8372638218fb8bd71d15de02b8277ce1c12207431bba5769",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/e/elfutils-libelf-devel-0.191-4.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/e/elfutils-libelf-devel-0.191-4.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/e/elfutils-libelf-devel-0.191-4.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/e/elfutils-libelf-devel-0.191-4.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/e/elfutils-libelf-devel-0.191-4.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "elfutils-libs-0__0.191-4.fc40.x86_64",
        sha256 = "d7d1ed3fca0696b8c38effe21bc70c84a94cb66c0b59bb1980c0f455d23b7fec",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/e/elfutils-libs-0.191-4.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/e/elfutils-libs-0.191-4.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/e/elfutils-libs-0.191-4.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/e/elfutils-libs-0.191-4.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/e/elfutils-libs-0.191-4.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "expat-0__2.6.3-1.fc40.x86_64",
        sha256 = "3a5ba168021a01107d6dd4dc7cffe8bb5553c64f236c436979b9fddfdc4cb59d",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/e/expat-2.6.3-1.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/e/expat-2.6.3-1.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/e/expat-2.6.3-1.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/e/expat-2.6.3-1.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/e/expat-2.6.3-1.fc40.x86_64.rpm",
        ],
    )

    rpm(
        name = "fedora-gpg-keys-0__40-2.x86_64",
        sha256 = "849feb04544096f9bbe16bc78c2198708fe658bdafa08575c911e538a7d31c18",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/f/fedora-gpg-keys-40-2.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/f/fedora-gpg-keys-40-2.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/f/fedora-gpg-keys-40-2.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/f/fedora-gpg-keys-40-2.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/f/fedora-gpg-keys-40-2.noarch.rpm",
        ],
    )
    rpm(
        name = "fedora-release-common-0__40-39.x86_64",
        sha256 = "590c9439a81fb9e35a8b4d19dc159ce09b756f8f7f66a6290d8785f424d97003",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/f/fedora-release-common-40-39.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/f/fedora-release-common-40-39.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/f/fedora-release-common-40-39.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/f/fedora-release-common-40-39.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/f/fedora-release-common-40-39.noarch.rpm",
        ],
    )
    rpm(
        name = "fedora-release-container-0__40-39.x86_64",
        sha256 = "28160baf6397f63f6d6e9c430ceab724a9847471721349e1b07f55a46477baa6",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/f/fedora-release-container-40-39.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/f/fedora-release-container-40-39.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/f/fedora-release-container-40-39.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/f/fedora-release-container-40-39.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/f/fedora-release-container-40-39.noarch.rpm",
        ],
    )
    rpm(
        name = "fedora-release-identity-container-0__40-39.x86_64",
        sha256 = "d328cd3f9ad90cd5881af43b1f47f11f8f48e8da888be34becef76c0d4377bfc",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/f/fedora-release-identity-container-40-39.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/f/fedora-release-identity-container-40-39.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/f/fedora-release-identity-container-40-39.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/f/fedora-release-identity-container-40-39.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/f/fedora-release-identity-container-40-39.noarch.rpm",
        ],
    )
    rpm(
        name = "fedora-repos-0__40-2.x86_64",
        sha256 = "e85d69eeea62f4f5a7c6584bc8bae3cb559c1c381838ca89f7d63b28d2368c4b",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/f/fedora-repos-40-2.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/f/fedora-repos-40-2.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/f/fedora-repos-40-2.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/f/fedora-repos-40-2.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/f/fedora-repos-40-2.noarch.rpm",
        ],
    )
    rpm(
        name = "file-0__5.45-4.fc40.x86_64",
        sha256 = "a6f2098fc2ed16df92c9325bd7459cc41479e17306a4f9cddfd5df8a1b80d0f8",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/f/file-5.45-4.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/f/file-5.45-4.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/f/file-5.45-4.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/f/file-5.45-4.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/f/file-5.45-4.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "file-libs-0__5.45-4.fc40.x86_64",
        sha256 = "f76684ee78408660db83ab9932978a1346b280f4210cd744524b00b2e5891fe1",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/f/file-libs-5.45-4.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/f/file-libs-5.45-4.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/f/file-libs-5.45-4.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/f/file-libs-5.45-4.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/f/file-libs-5.45-4.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "filesystem-0__3.18-8.fc40.x86_64",
        sha256 = "063af3db3808bea0d5c07dbb2d8369b275e1d05ad0850c80a8fec0413f47cd64",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/f/filesystem-3.18-8.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/f/filesystem-3.18-8.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/f/filesystem-3.18-8.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/f/filesystem-3.18-8.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/f/filesystem-3.18-8.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "findutils-1__4.9.0-9.fc40.x86_64",
        sha256 = "21725de2a93e1ea19f8d298e32a2428a3a08b9c98f22561cc778a807ed43639f",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/f/findutils-4.9.0-9.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/f/findutils-4.9.0-9.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/f/findutils-4.9.0-9.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/f/findutils-4.9.0-9.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/f/findutils-4.9.0-9.fc40.x86_64.rpm",
        ],
    )

    rpm(
        name = "flex-0__2.6.4-16.fc40.x86_64",
        sha256 = "7bef707ffb9672420bf6179bccecb26c0b86f74fcb0d521cbc3d651dd486ced0",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/f/flex-2.6.4-16.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/f/flex-2.6.4-16.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/f/flex-2.6.4-16.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/f/flex-2.6.4-16.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/f/flex-2.6.4-16.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "fonts-srpm-macros-1__2.0.5-14.fc40.x86_64",
        sha256 = "ebf245973cea76d51b22de0e587fc77b3d6a776fb629c4130971182536afd9d7",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/f/fonts-srpm-macros-2.0.5-14.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/f/fonts-srpm-macros-2.0.5-14.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/f/fonts-srpm-macros-2.0.5-14.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/f/fonts-srpm-macros-2.0.5-14.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/f/fonts-srpm-macros-2.0.5-14.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "forge-srpm-macros-0__0.3.2-1.fc40.x86_64",
        sha256 = "fd7875da8f0a566458ae18420e765f97ee7429c51819c084b5774cbc182e6f83",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/f/forge-srpm-macros-0.3.2-1.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/f/forge-srpm-macros-0.3.2-1.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/f/forge-srpm-macros-0.3.2-1.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/f/forge-srpm-macros-0.3.2-1.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/f/forge-srpm-macros-0.3.2-1.fc40.noarch.rpm",
        ],
    )

    rpm(
        name = "fpc-srpm-macros-0__1.3-12.fc40.x86_64",
        sha256 = "7df65ab4ab462818320c8391aa8b08e63fddba2c60944e40f0b207118effbae5",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/f/fpc-srpm-macros-1.3-12.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/f/fpc-srpm-macros-1.3-12.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/f/fpc-srpm-macros-1.3-12.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/f/fpc-srpm-macros-1.3-12.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/f/fpc-srpm-macros-1.3-12.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "fuse3-libs-0__3.16.2-3.fc40.x86_64",
        sha256 = "a9c6502a5b190aaf169e93afd337c009e0b2e235e31f3da23d29c7d063ad2ff9",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/f/fuse3-libs-3.16.2-3.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/f/fuse3-libs-3.16.2-3.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/f/fuse3-libs-3.16.2-3.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/f/fuse3-libs-3.16.2-3.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/f/fuse3-libs-3.16.2-3.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "gawk-0__5.3.0-3.fc40.x86_64",
        sha256 = "6c80dfdaf7b27ea92c1276856b8b2ae5fde1ae5c391b773805be725515fdc1ac",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/g/gawk-5.3.0-3.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/g/gawk-5.3.0-3.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/g/gawk-5.3.0-3.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/g/gawk-5.3.0-3.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/g/gawk-5.3.0-3.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "gc-0__8.2.2-6.fc40.x86_64",
        sha256 = "641494edfcaad3ca02445dafa37b652efa7c188f40e09d50445a4b5d6d7965df",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/g/gc-8.2.2-6.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/g/gc-8.2.2-6.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/g/gc-8.2.2-6.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/g/gc-8.2.2-6.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/g/gc-8.2.2-6.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "gcc-0__14.2.1-3.fc40.x86_64",
        sha256 = "c828d2ea3d79beb95fab71c59d378f6e9834751ac2189af9c8685e123ee49642",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/g/gcc-14.2.1-3.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/g/gcc-14.2.1-3.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/g/gcc-14.2.1-3.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/g/gcc-14.2.1-3.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/g/gcc-14.2.1-3.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "gcc-c__plus____plus__-0__14.2.1-3.fc40.x86_64",
        sha256 = "34b69d3315d3f588d196cf883b5fc1f612728d9201b2a6e3e88d3794d784229e",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/g/gcc-c++-14.2.1-3.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/g/gcc-c++-14.2.1-3.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/g/gcc-c++-14.2.1-3.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/g/gcc-c++-14.2.1-3.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/g/gcc-c++-14.2.1-3.fc40.x86_64.rpm",
        ],
    )

    rpm(
        name = "gdbm-1__1.23-6.fc40.x86_64",
        sha256 = "21470eb4ec55006c9efeee84c97772462008fceda1ab332e58d2caddfdaa0d1e",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/g/gdbm-1.23-6.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/g/gdbm-1.23-6.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/g/gdbm-1.23-6.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/g/gdbm-1.23-6.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/g/gdbm-1.23-6.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "gdbm-libs-1__1.23-6.fc40.x86_64",
        sha256 = "93450209842a296ea4b295f6d86b69aa52dd8ec45b121ede0d5125aa49bad509",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/g/gdbm-libs-1.23-6.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/g/gdbm-libs-1.23-6.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/g/gdbm-libs-1.23-6.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/g/gdbm-libs-1.23-6.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/g/gdbm-libs-1.23-6.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "ghc-srpm-macros-0__1.9.1-1.fc40.x86_64",
        sha256 = "1509ca46a18243b3f181aac3d77639b805c470816f892fbf62acd0ae96f01f9a",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/g/ghc-srpm-macros-1.9.1-1.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/g/ghc-srpm-macros-1.9.1-1.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/g/ghc-srpm-macros-1.9.1-1.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/g/ghc-srpm-macros-1.9.1-1.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/g/ghc-srpm-macros-1.9.1-1.fc40.noarch.rpm",
        ],
    )

    rpm(
        name = "glib2-0__2.80.3-1.fc40.x86_64",
        sha256 = "0a32c6874ce180375c2c0b1e2f0c8fed38131a598e5c4ba3866cf3aee1f3f5fc",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/g/glib2-2.80.3-1.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/g/glib2-2.80.3-1.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/g/glib2-2.80.3-1.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/g/glib2-2.80.3-1.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/g/glib2-2.80.3-1.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "glibc-0__2.39-22.fc40.x86_64",
        sha256 = "726a1d707dfcf20d1f4c94f76bdba38d166eb574ecc2d83ec438bdc161f3ec27",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/g/glibc-2.39-22.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/g/glibc-2.39-22.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/g/glibc-2.39-22.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/g/glibc-2.39-22.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/g/glibc-2.39-22.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "glibc-common-0__2.39-22.fc40.x86_64",
        sha256 = "0b502c1140a1f6461dbd63d3daefedded6c705769476cec5e69466aab7693ea6",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/g/glibc-common-2.39-22.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/g/glibc-common-2.39-22.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/g/glibc-common-2.39-22.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/g/glibc-common-2.39-22.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/g/glibc-common-2.39-22.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "glibc-devel-0__2.39-22.fc40.x86_64",
        sha256 = "bd2f6ddc6edb65a0c63b952d272f748fea1102630ca43facf25bb0a17b3b3ab9",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/g/glibc-devel-2.39-22.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/g/glibc-devel-2.39-22.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/g/glibc-devel-2.39-22.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/g/glibc-devel-2.39-22.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/g/glibc-devel-2.39-22.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "glibc-headers-x86-0__2.39-22.fc40.x86_64",
        sha256 = "064d8433785d19b73f84bf8e15edb9e8f38458f49604b56f1a892fa171e24849",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/g/glibc-headers-x86-2.39-22.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/g/glibc-headers-x86-2.39-22.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/g/glibc-headers-x86-2.39-22.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/g/glibc-headers-x86-2.39-22.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/g/glibc-headers-x86-2.39-22.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "glibc-langpack-en-0__2.39-22.fc40.x86_64",
        sha256 = "89174efddcfb8c11a9c6eee70998d8d45c6bc72a019a928412d249db165b3935",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/g/glibc-langpack-en-2.39-22.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/g/glibc-langpack-en-2.39-22.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/g/glibc-langpack-en-2.39-22.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/g/glibc-langpack-en-2.39-22.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/g/glibc-langpack-en-2.39-22.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "glibc-static-0__2.39-22.fc40.x86_64",
        sha256 = "ef8cd1b3168deab09f4fe026f25835d9348b6300a7aafe56334919aed46a0873",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/g/glibc-static-2.39-22.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/g/glibc-static-2.39-22.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/g/glibc-static-2.39-22.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/g/glibc-static-2.39-22.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/g/glibc-static-2.39-22.fc40.x86_64.rpm",
        ],
    )

    rpm(
        name = "gmp-1__6.2.1-8.fc40.x86_64",
        sha256 = "b054d6a9ee3477e935686b327aa47379bd1909eac4ce06c4c45dff1a201ecb49",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/g/gmp-6.2.1-8.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/g/gmp-6.2.1-8.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/g/gmp-6.2.1-8.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/g/gmp-6.2.1-8.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/g/gmp-6.2.1-8.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "gnat-srpm-macros-0__6-5.fc40.x86_64",
        sha256 = "35f84a6494aed02d6a2b90f702787232962535c313ab56b3878b264a6c39546c",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/g/gnat-srpm-macros-6-5.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/g/gnat-srpm-macros-6-5.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/g/gnat-srpm-macros-6-5.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/g/gnat-srpm-macros-6-5.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/g/gnat-srpm-macros-6-5.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "gnupg2-0__2.4.4-1.fc40.x86_64",
        sha256 = "0a8b1b3fb625e4d1864ad6726f583e2db5db7f10d9f3564b5916ca7fed1b71cb",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/g/gnupg2-2.4.4-1.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/g/gnupg2-2.4.4-1.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/g/gnupg2-2.4.4-1.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/g/gnupg2-2.4.4-1.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/g/gnupg2-2.4.4-1.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "gnutls-0__3.8.6-1.fc40.x86_64",
        sha256 = "4289ccbb44e4a764ef6f58593a56f2598c6821feebac52be6fa04c771eebf029",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/g/gnutls-3.8.6-1.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/g/gnutls-3.8.6-1.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/g/gnutls-3.8.6-1.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/g/gnutls-3.8.6-1.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/g/gnutls-3.8.6-1.fc40.x86_64.rpm",
        ],
    )

    rpm(
        name = "go-srpm-macros-0__3.5.0-1.fc40.x86_64",
        sha256 = "2968803f0da871cb5b5873efab1360871260c915338f72f11486a1210aafd105",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/g/go-srpm-macros-3.5.0-1.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/g/go-srpm-macros-3.5.0-1.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/g/go-srpm-macros-3.5.0-1.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/g/go-srpm-macros-3.5.0-1.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/g/go-srpm-macros-3.5.0-1.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "grep-0__3.11-7.fc40.x86_64",
        sha256 = "8e2310f6cde324576e537749cf1d4fee8028edfc0c8df3070f147ee162b423ce",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/g/grep-3.11-7.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/g/grep-3.11-7.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/g/grep-3.11-7.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/g/grep-3.11-7.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/g/grep-3.11-7.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "groff-base-0__1.23.0-6.fc40.x86_64",
        sha256 = "30f12d19cb6077b8d0d644e59a94cc0163722e26b8771f9eb14a3edb0e9df25d",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/g/groff-base-1.23.0-6.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/g/groff-base-1.23.0-6.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/g/groff-base-1.23.0-6.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/g/groff-base-1.23.0-6.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/g/groff-base-1.23.0-6.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "guile30-0__3.0.7-12.fc40.x86_64",
        sha256 = "322aa327b35cbd6cd85265bc282505d297e7cc8e7d70dd856705dad805c50af8",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/g/guile30-3.0.7-12.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/g/guile30-3.0.7-12.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/g/guile30-3.0.7-12.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/g/guile30-3.0.7-12.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/g/guile30-3.0.7-12.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "gzip-0__1.13-1.fc40.x86_64",
        sha256 = "6dcc2f8885135fc873c8ab94a6c7df05883060c5b25287956bebb3aa15a84e71",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/g/gzip-1.13-1.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/g/gzip-1.13-1.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/g/gzip-1.13-1.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/g/gzip-1.13-1.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/g/gzip-1.13-1.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "ipxe-roms-qemu-0__20240119-1.gitde8a0821.fc40.x86_64",
        sha256 = "1d84c5d480e0f23c4dfda72ff6db466d4959941d897fe517a9771112d41203bc",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/i/ipxe-roms-qemu-20240119-1.gitde8a0821.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/i/ipxe-roms-qemu-20240119-1.gitde8a0821.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/i/ipxe-roms-qemu-20240119-1.gitde8a0821.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/i/ipxe-roms-qemu-20240119-1.gitde8a0821.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/i/ipxe-roms-qemu-20240119-1.gitde8a0821.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "jansson-0__2.13.1-9.fc40.x86_64",
        sha256 = "9b4f2730a62955650c1e260e1b573f089355faf0155871e2c10381316a3b2e55",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/j/jansson-2.13.1-9.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/j/jansson-2.13.1-9.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/j/jansson-2.13.1-9.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/j/jansson-2.13.1-9.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/j/jansson-2.13.1-9.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "json-c-0__0.17-3.fc40.x86_64",
        sha256 = "77e67991fcd4eea31f5b2844898a7854768548f0ab3abf7beaa91526afbf794b",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/j/json-c-0.17-3.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/j/json-c-0.17-3.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/j/json-c-0.17-3.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/j/json-c-0.17-3.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/j/json-c-0.17-3.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "kernel-headers-0__6.11.3-200.fc40.x86_64",
        sha256 = "e1022a6f80e0968b3289ba0c28acfa9b26bcd7d8911eac51e7fb174690834432",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/k/kernel-headers-6.11.3-200.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/k/kernel-headers-6.11.3-200.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/k/kernel-headers-6.11.3-200.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/k/kernel-headers-6.11.3-200.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/k/kernel-headers-6.11.3-200.fc40.x86_64.rpm",
        ],
    )

    rpm(
        name = "kernel-srpm-macros-0__1.0-23.fc40.x86_64",
        sha256 = "95fb5031a23336455d606d05c63855c7f12247ffd4baaac64fb576b420b2a32e",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/k/kernel-srpm-macros-1.0-23.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/k/kernel-srpm-macros-1.0-23.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/k/kernel-srpm-macros-1.0-23.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/k/kernel-srpm-macros-1.0-23.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/k/kernel-srpm-macros-1.0-23.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "keyutils-libs-0__1.6.3-3.fc40.x86_64",
        sha256 = "387706fa265213dc46e4f818f30333cc93f0c54539cbd2ec4db3bc854077307b",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/k/keyutils-libs-1.6.3-3.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/k/keyutils-libs-1.6.3-3.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/k/keyutils-libs-1.6.3-3.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/k/keyutils-libs-1.6.3-3.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/k/keyutils-libs-1.6.3-3.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "kmod-libs-0__31-5.fc40.x86_64",
        sha256 = "53dd95341767a2ea40b68e4621a231883bd5b69426f0920ce1f1ca94e18765cb",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/k/kmod-libs-31-5.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/k/kmod-libs-31-5.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/k/kmod-libs-31-5.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/k/kmod-libs-31-5.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/k/kmod-libs-31-5.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "krb5-libs-0__1.21.3-1.fc40.x86_64",
        sha256 = "2c32d410b49c6d3d4f66b361169ad76dfd9f75ee01d9866c62b14d1e5dfc5124",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/k/krb5-libs-1.21.3-1.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/k/krb5-libs-1.21.3-1.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/k/krb5-libs-1.21.3-1.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/k/krb5-libs-1.21.3-1.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/k/krb5-libs-1.21.3-1.fc40.x86_64.rpm",
        ],
    )

    rpm(
        name = "libacl-0__2.3.2-1.fc40.x86_64",
        sha256 = "b753174804f57c3c6bae7afeb6145005498f18ae5d1aa0d340f9df5b8d71312f",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libacl-2.3.2-1.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libacl-2.3.2-1.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libacl-2.3.2-1.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libacl-2.3.2-1.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libacl-2.3.2-1.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "libaio-0__0.3.111-19.fc40.x86_64",
        sha256 = "5f5bb334bc8a867320d5f43d27e2b996b76291cd4dbb5470a55ece94028966e1",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libaio-0.3.111-19.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libaio-0.3.111-19.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libaio-0.3.111-19.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libaio-0.3.111-19.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libaio-0.3.111-19.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "libarchive-0__3.7.2-7.fc40.x86_64",
        sha256 = "74d72760c1982830358d676794ee3972ab05550fe7235ae9756a40de8266091f",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/l/libarchive-3.7.2-7.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/l/libarchive-3.7.2-7.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/l/libarchive-3.7.2-7.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/l/libarchive-3.7.2-7.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/l/libarchive-3.7.2-7.fc40.x86_64.rpm",
        ],
    )

    rpm(
        name = "libassuan-0__2.5.7-1.fc40.x86_64",
        sha256 = "e131ab89604dbd4fdc4f80af632099e48bf68bb328dbf0e7dcbef1d1e134dc09",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libassuan-2.5.7-1.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libassuan-2.5.7-1.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libassuan-2.5.7-1.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libassuan-2.5.7-1.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libassuan-2.5.7-1.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "libattr-0__2.5.2-3.fc40.x86_64",
        sha256 = "504cff39c51a04c1d302096899c47dc34ac0eba47524c2fc94c27904149e72cf",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libattr-2.5.2-3.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libattr-2.5.2-3.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libattr-2.5.2-3.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libattr-2.5.2-3.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libattr-2.5.2-3.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "libb2-0__0.98.1-11.fc40.x86_64",
        sha256 = "649cceb60f2e284f8d5dadeec4af8e7035650fe0e5aa75c552354b3fa5708cfe",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libb2-0.98.1-11.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libb2-0.98.1-11.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libb2-0.98.1-11.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libb2-0.98.1-11.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libb2-0.98.1-11.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "libblkid-0__2.40.2-1.fc40.x86_64",
        sha256 = "b506de64d63262d9d957a75fdf2282d82b1e4978cebbdfc191ef93bba37e3b7c",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/l/libblkid-2.40.2-1.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/l/libblkid-2.40.2-1.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/l/libblkid-2.40.2-1.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/l/libblkid-2.40.2-1.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/l/libblkid-2.40.2-1.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "libbpf-2__1.2.3-1.fc40.x86_64",
        sha256 = "fca2d942f6264b630b33991e48dcb605543a4c837371f28f92994bf956677f24",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/l/libbpf-1.2.3-1.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/l/libbpf-1.2.3-1.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/l/libbpf-1.2.3-1.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/l/libbpf-1.2.3-1.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/l/libbpf-1.2.3-1.fc40.x86_64.rpm",
        ],
    )

    rpm(
        name = "libbrotli-0__1.1.0-3.fc40.x86_64",
        sha256 = "97e9e5339bb0ca6ce3d0195c8ebe48384bcfc087ee6bc7a35b1d27d4de23fbfa",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libbrotli-1.1.0-3.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libbrotli-1.1.0-3.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libbrotli-1.1.0-3.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libbrotli-1.1.0-3.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libbrotli-1.1.0-3.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "libcap-0__2.69-8.fc40.x86_64",
        sha256 = "6c92fc0c357964d2b57533a408ec93b7fe5214c1f0b63a6b1c0564b2ba5c481f",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/l/libcap-2.69-8.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/l/libcap-2.69-8.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/l/libcap-2.69-8.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/l/libcap-2.69-8.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/l/libcap-2.69-8.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "libcap-ng-0__0.8.4-4.fc40.x86_64",
        sha256 = "dc22477c3ac762f92ecc322af4f39fee2c5371bedc495ce242f9b94c590c580f",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libcap-ng-0.8.4-4.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libcap-ng-0.8.4-4.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libcap-ng-0.8.4-4.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libcap-ng-0.8.4-4.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libcap-ng-0.8.4-4.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "libcom_err-0__1.47.0-5.fc40.x86_64",
        sha256 = "0d100701976c37fe94e904ed78437db7477ae1dc600ece07bea23fbbd968762c",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libcom_err-1.47.0-5.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libcom_err-1.47.0-5.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libcom_err-1.47.0-5.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libcom_err-1.47.0-5.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libcom_err-1.47.0-5.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "libcurl-0__8.6.0-10.fc40.x86_64",
        sha256 = "b26b1cf74b7b0d6e8f10ee73b0ed5cb33d1a953510814b51a421befb6dd36e2d",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/l/libcurl-8.6.0-10.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/l/libcurl-8.6.0-10.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/l/libcurl-8.6.0-10.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/l/libcurl-8.6.0-10.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/l/libcurl-8.6.0-10.fc40.x86_64.rpm",
        ],
    )

    rpm(
        name = "libdb-0__5.3.28-62.fc40.x86_64",
        sha256 = "03642fc857f0b734ca68dfca6824a09bf7bc8439d2febd1a87f8617ddfba2c1c",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/l/libdb-5.3.28-62.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/l/libdb-5.3.28-62.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/l/libdb-5.3.28-62.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/l/libdb-5.3.28-62.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/l/libdb-5.3.28-62.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "libeconf-0__0.6.2-2.fc40.x86_64",
        sha256 = "2ef764049e121ee2a9fa5d0296e6e2dd0abc7541040b8e49d67960bd9bde74e4",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/l/libeconf-0.6.2-2.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/l/libeconf-0.6.2-2.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/l/libeconf-0.6.2-2.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/l/libeconf-0.6.2-2.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/l/libeconf-0.6.2-2.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "libedit-0__3.1-53.20240808cvs.fc40.x86_64",
        sha256 = "b003de79beac86385d212fce137417439e8ec7cb863115d560e02834c84efd1e",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/l/libedit-3.1-53.20240808cvs.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/l/libedit-3.1-53.20240808cvs.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/l/libedit-3.1-53.20240808cvs.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/l/libedit-3.1-53.20240808cvs.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/l/libedit-3.1-53.20240808cvs.fc40.x86_64.rpm",
        ],
    )

    rpm(
        name = "libevent-0__2.1.12-12.fc40.x86_64",
        sha256 = "c4adcee5dd9e22ea50d6c318ac4936a8df708121741958ce5aa8f038c46c61a9",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libevent-2.1.12-12.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libevent-2.1.12-12.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libevent-2.1.12-12.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libevent-2.1.12-12.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libevent-2.1.12-12.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "libfdisk-0__2.40.2-1.fc40.x86_64",
        sha256 = "aa6a51bbe265bb3d3a50c37557f6513d51298301e4957ce4484e56feb837fa32",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/l/libfdisk-2.40.2-1.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/l/libfdisk-2.40.2-1.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/l/libfdisk-2.40.2-1.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/l/libfdisk-2.40.2-1.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/l/libfdisk-2.40.2-1.fc40.x86_64.rpm",
        ],
    )

    rpm(
        name = "libfdt-0__1.7.0-7.fc40.x86_64",
        sha256 = "38c9fd945b14b1a58c5b7e74e9e6a06f4429cc186dc29b1af2a7e2629a44996f",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libfdt-1.7.0-7.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libfdt-1.7.0-7.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libfdt-1.7.0-7.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libfdt-1.7.0-7.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libfdt-1.7.0-7.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "libffi-0__3.4.4-7.fc40.x86_64",
        sha256 = "25caa7ee56f6013369c2fac26afd3035a7d580af0b919621ba8d495d13a5af86",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libffi-3.4.4-7.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libffi-3.4.4-7.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libffi-3.4.4-7.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libffi-3.4.4-7.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libffi-3.4.4-7.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "libgcc-0__14.2.1-3.fc40.x86_64",
        sha256 = "cd073c42cb4dfcd224e9b4619883f2c7923ab0b083d7c90b01e3052c89f6b814",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/l/libgcc-14.2.1-3.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/l/libgcc-14.2.1-3.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/l/libgcc-14.2.1-3.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/l/libgcc-14.2.1-3.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/l/libgcc-14.2.1-3.fc40.x86_64.rpm",
        ],
    )

    rpm(
        name = "libgcrypt-0__1.10.3-3.fc40.x86_64",
        sha256 = "10c4c12c6539ffea68974cd9b57013d471ac35fe3bef4833c0a22f6b29fbf489",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libgcrypt-1.10.3-3.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libgcrypt-1.10.3-3.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libgcrypt-1.10.3-3.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libgcrypt-1.10.3-3.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libgcrypt-1.10.3-3.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "libgomp-0__14.2.1-3.fc40.x86_64",
        sha256 = "03d5f4d139dec2e7c94714b1b9f59d37236dbda9f09271bdda99c71251f15f0e",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/l/libgomp-14.2.1-3.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/l/libgomp-14.2.1-3.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/l/libgomp-14.2.1-3.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/l/libgomp-14.2.1-3.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/l/libgomp-14.2.1-3.fc40.x86_64.rpm",
        ],
    )

    rpm(
        name = "libgpg-error-0__1.49-1.fc40.x86_64",
        sha256 = "8d0a9840e06e72ccf756fa5a79c49f572dc827b0c75ea5a1f923235150d27ae2",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/l/libgpg-error-1.49-1.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/l/libgpg-error-1.49-1.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/l/libgpg-error-1.49-1.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/l/libgpg-error-1.49-1.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/l/libgpg-error-1.49-1.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "libibverbs-0__48.0-4.fc40.x86_64",
        sha256 = "607dbffbe375f62dc0755457d33ef59538e0f061b28bfa44c25705533ddd7e20",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libibverbs-48.0-4.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libibverbs-48.0-4.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libibverbs-48.0-4.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libibverbs-48.0-4.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libibverbs-48.0-4.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "libidn2-0__2.3.7-1.fc40.x86_64",
        sha256 = "2fd2038b4a94eeede34e46ed0e035e619f77d0e412c70cf4e9bb836957e8f31b",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libidn2-2.3.7-1.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libidn2-2.3.7-1.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libidn2-2.3.7-1.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libidn2-2.3.7-1.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libidn2-2.3.7-1.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "libjpeg-turbo-0__3.0.2-1.fc40.x86_64",
        sha256 = "642019ca5920baf189a25b9305bd715705f4782822e9fac21b5781c51460317d",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libjpeg-turbo-3.0.2-1.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libjpeg-turbo-3.0.2-1.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libjpeg-turbo-3.0.2-1.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libjpeg-turbo-3.0.2-1.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libjpeg-turbo-3.0.2-1.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "libksba-0__1.6.6-1.fc40.x86_64",
        sha256 = "a77eed0fe1b84c11f9175f4642db058753d4eaa1f88e999f01df72e1d10a3826",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libksba-1.6.6-1.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libksba-1.6.6-1.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libksba-1.6.6-1.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libksba-1.6.6-1.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libksba-1.6.6-1.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "libmount-0__2.40.2-1.fc40.x86_64",
        sha256 = "a695daa293bb78b033a2629f5af1284fe212b748227e94efa59a8292eb6b9f40",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/l/libmount-2.40.2-1.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/l/libmount-2.40.2-1.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/l/libmount-2.40.2-1.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/l/libmount-2.40.2-1.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/l/libmount-2.40.2-1.fc40.x86_64.rpm",
        ],
    )

    rpm(
        name = "libmpc-0__1.3.1-5.fc40.x86_64",
        sha256 = "b749c245ecd4d9457a94e2eedbe7196837566bb13b94d0827b45b5135286f6f4",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libmpc-1.3.1-5.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libmpc-1.3.1-5.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libmpc-1.3.1-5.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libmpc-1.3.1-5.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libmpc-1.3.1-5.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "libnghttp2-0__1.59.0-3.fc40.x86_64",
        sha256 = "550160732fc268914a422cfddc3c745febf8da161f8eacbce8649c67117b1476",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/l/libnghttp2-1.59.0-3.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/l/libnghttp2-1.59.0-3.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/l/libnghttp2-1.59.0-3.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/l/libnghttp2-1.59.0-3.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/l/libnghttp2-1.59.0-3.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "libnl3-0__3.10.0-1.fc40.x86_64",
        sha256 = "13c21b26297876c42723a5557d97ecfada4ab2a79a4e4e771a6a3df29ffd5e47",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/l/libnl3-3.10.0-1.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/l/libnl3-3.10.0-1.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/l/libnl3-3.10.0-1.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/l/libnl3-3.10.0-1.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/l/libnl3-3.10.0-1.fc40.x86_64.rpm",
        ],
    )

    rpm(
        name = "libnsl2-0__2.0.1-1.fc40.x86_64",
        sha256 = "fa6dccd7aee4a74a5cfa12c7927c7326485704ebe57c54774b0f157fda639360",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libnsl2-2.0.1-1.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libnsl2-2.0.1-1.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libnsl2-2.0.1-1.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libnsl2-2.0.1-1.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libnsl2-2.0.1-1.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "libpkgconf-0__2.1.1-2.fc40.x86_64",
        sha256 = "3f8bb781f774a99ca828f8c5a56b827cf2ad8578330e8e8ae2ab19905a77718e",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/l/libpkgconf-2.1.1-2.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/l/libpkgconf-2.1.1-2.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/l/libpkgconf-2.1.1-2.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/l/libpkgconf-2.1.1-2.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/l/libpkgconf-2.1.1-2.fc40.x86_64.rpm",
        ],
    )

    rpm(
        name = "libpmem-0__2.0.1-3.fc40.x86_64",
        sha256 = "216f96ee920ff4c1601de1b70d3a24e19e5055909f663571487f764182e98819",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libpmem-2.0.1-3.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libpmem-2.0.1-3.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libpmem-2.0.1-3.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libpmem-2.0.1-3.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libpmem-2.0.1-3.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "libpng-2__1.6.40-3.fc40.x86_64",
        sha256 = "f115b64206304002c658f83c829623aa966e0d99f24de5d60c79a19142803ecb",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libpng-1.6.40-3.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libpng-1.6.40-3.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libpng-1.6.40-3.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libpng-1.6.40-3.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libpng-1.6.40-3.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "libpsl-0__0.21.5-3.fc40.x86_64",
        sha256 = "bb9ceaba0d3283777777524e8c99b8eaa2155e9000d8e3ef5d0ece336f8c1392",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libpsl-0.21.5-3.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libpsl-0.21.5-3.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libpsl-0.21.5-3.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libpsl-0.21.5-3.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libpsl-0.21.5-3.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "libpwquality-0__1.4.5-9.fc40.x86_64",
        sha256 = "210e797a265da7111c1a59eca95f9e301ad05c5c8772aed54af9363e5684950b",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libpwquality-1.4.5-9.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libpwquality-1.4.5-9.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libpwquality-1.4.5-9.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libpwquality-1.4.5-9.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libpwquality-1.4.5-9.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "librdmacm-0__48.0-4.fc40.x86_64",
        sha256 = "413514fb04b2235b0bd6ab7173f0c62e3945e62262a4ec0762a8c3e8173e5ed5",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/librdmacm-48.0-4.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/librdmacm-48.0-4.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/librdmacm-48.0-4.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/librdmacm-48.0-4.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/librdmacm-48.0-4.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "libseccomp-0__2.5.5-1.fc40.x86_64",
        sha256 = "91668f5d08a663948c7d888d7cdef3248285c5d9fbe369ae031d7ca31c6e398c",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/l/libseccomp-2.5.5-1.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/l/libseccomp-2.5.5-1.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/l/libseccomp-2.5.5-1.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/l/libseccomp-2.5.5-1.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/l/libseccomp-2.5.5-1.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "libselinux-0__3.7-5.fc40.x86_64",
        sha256 = "2070bdf786c926400739254f08568ccf564ce613ddacacb36b6a9a499345aa5e",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/l/libselinux-3.7-5.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/l/libselinux-3.7-5.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/l/libselinux-3.7-5.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/l/libselinux-3.7-5.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/l/libselinux-3.7-5.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "libselinux-utils-0__3.7-5.fc40.x86_64",
        sha256 = "aca271d814ee3be14c09963985011c201315a186d3e3b634af8d59cd5eb01208",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/l/libselinux-utils-3.7-5.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/l/libselinux-utils-3.7-5.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/l/libselinux-utils-3.7-5.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/l/libselinux-utils-3.7-5.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/l/libselinux-utils-3.7-5.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "libsemanage-0__3.7-2.fc40.x86_64",
        sha256 = "e200b862d5063f6e85859c5be99c50d5636edae91bd3f603c3a22383b7e2ac88",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/l/libsemanage-3.7-2.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/l/libsemanage-3.7-2.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/l/libsemanage-3.7-2.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/l/libsemanage-3.7-2.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/l/libsemanage-3.7-2.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "libsepol-0__3.7-2.fc40.x86_64",
        sha256 = "85cbaeca877a166cda9637a8ea0d43dd63488fdcc250fe564696cf8beaf8913f",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/l/libsepol-3.7-2.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/l/libsepol-3.7-2.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/l/libsepol-3.7-2.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/l/libsepol-3.7-2.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/l/libsepol-3.7-2.fc40.x86_64.rpm",
        ],
    )

    rpm(
        name = "libslirp-0__4.7.0-6.fc40.x86_64",
        sha256 = "9d552a0d0609305a0a72eaa4470efcda4fb3947b301205fd7d292fb48246e47a",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libslirp-4.7.0-6.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libslirp-4.7.0-6.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libslirp-4.7.0-6.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libslirp-4.7.0-6.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libslirp-4.7.0-6.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "libsmartcols-0__2.40.2-1.fc40.x86_64",
        sha256 = "e9c3e9e3458af7a2f9b5cd6bc45020bb7f2c6cfbd0429b0b1853928bd3e02004",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/l/libsmartcols-2.40.2-1.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/l/libsmartcols-2.40.2-1.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/l/libsmartcols-2.40.2-1.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/l/libsmartcols-2.40.2-1.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/l/libsmartcols-2.40.2-1.fc40.x86_64.rpm",
        ],
    )

    rpm(
        name = "libssh-0__0.10.6-5.fc40.x86_64",
        sha256 = "45695cddc79eafe4c52c44d59d6a8a88850e4bf809fa50d19e56042f1a02f08f",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libssh-0.10.6-5.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libssh-0.10.6-5.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libssh-0.10.6-5.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libssh-0.10.6-5.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libssh-0.10.6-5.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "libssh-config-0__0.10.6-5.fc40.x86_64",
        sha256 = "241c73071a373732ec544dad6ba6f4fb054c1f2264d86085c322dd1c1089bbb1",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libssh-config-0.10.6-5.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libssh-config-0.10.6-5.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libssh-config-0.10.6-5.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libssh-config-0.10.6-5.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libssh-config-0.10.6-5.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "libstdc__plus____plus__-0__14.2.1-3.fc40.x86_64",
        sha256 = "89e7282e0a94d641871dfed423ba2ce6f8b088eaf9aabdea1805708bcafa6a01",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/l/libstdc++-14.2.1-3.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/l/libstdc++-14.2.1-3.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/l/libstdc++-14.2.1-3.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/l/libstdc++-14.2.1-3.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/l/libstdc++-14.2.1-3.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "libstdc__plus____plus__-devel-0__14.2.1-3.fc40.x86_64",
        sha256 = "a8fd68e48b66a6af74c4fb7f5df36be8fb0ab37bbf892be8e215d9c1546f4d3b",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/l/libstdc++-devel-14.2.1-3.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/l/libstdc++-devel-14.2.1-3.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/l/libstdc++-devel-14.2.1-3.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/l/libstdc++-devel-14.2.1-3.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/l/libstdc++-devel-14.2.1-3.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "libstdc__plus____plus__-static-0__14.2.1-3.fc40.x86_64",
        sha256 = "684261627afdba894d150ae35af70b685739cd1af392c6699bb0f199f41c538a",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/l/libstdc++-static-14.2.1-3.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/l/libstdc++-static-14.2.1-3.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/l/libstdc++-static-14.2.1-3.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/l/libstdc++-static-14.2.1-3.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/l/libstdc++-static-14.2.1-3.fc40.x86_64.rpm",
        ],
    )

    rpm(
        name = "libtasn1-0__4.19.0-6.fc40.x86_64",
        sha256 = "d92173d6fbfb7e2af3b35a8554229e247666e15dc5b36cba43b7bbfc4144b781",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libtasn1-4.19.0-6.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libtasn1-4.19.0-6.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libtasn1-4.19.0-6.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libtasn1-4.19.0-6.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libtasn1-4.19.0-6.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "libtirpc-0__1.3.5-0.fc40.x86_64",
        sha256 = "c52aa65956ce5076c7036d486ec29d06832461450c77838c7b9e360c701b6ad2",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/l/libtirpc-1.3.5-0.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/l/libtirpc-1.3.5-0.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/l/libtirpc-1.3.5-0.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/l/libtirpc-1.3.5-0.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/l/libtirpc-1.3.5-0.fc40.x86_64.rpm",
        ],
    )

    rpm(
        name = "libtool-ltdl-0__2.4.7-10.fc40.x86_64",
        sha256 = "e5d150d23f95e4a23288b84145af442607a88bf457c0e04b325b1d1e8e708c2b",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libtool-ltdl-2.4.7-10.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libtool-ltdl-2.4.7-10.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libtool-ltdl-2.4.7-10.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libtool-ltdl-2.4.7-10.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libtool-ltdl-2.4.7-10.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "libunistring-0__1.1-7.fc40.x86_64",
        sha256 = "58719c2f205b23598e31b72144ab55215947ad8fca96af46a641288692c159d2",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libunistring-1.1-7.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libunistring-1.1-7.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libunistring-1.1-7.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libunistring-1.1-7.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libunistring-1.1-7.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "liburing-0__2.5-3.fc40.x86_64",
        sha256 = "7b2df98d3b56482ef87b0751dd7ced32e235e27f4d5083082d283454a7c1e09c",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/liburing-2.5-3.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/liburing-2.5-3.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/liburing-2.5-3.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/liburing-2.5-3.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/liburing-2.5-3.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "libutempter-0__1.2.1-13.fc40.x86_64",
        sha256 = "0093a8d3f490fbbbc71b01e0c8f9b083040dbf7513be31a91a0769d846198c1b",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libutempter-1.2.1-13.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libutempter-1.2.1-13.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libutempter-1.2.1-13.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libutempter-1.2.1-13.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libutempter-1.2.1-13.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "libuuid-0__2.40.2-1.fc40.x86_64",
        sha256 = "b6db3e72ae6575127216145c1f65414ea94acd9db26d08c5081cb5d786101c1f",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/l/libuuid-2.40.2-1.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/l/libuuid-2.40.2-1.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/l/libuuid-2.40.2-1.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/l/libuuid-2.40.2-1.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/l/libuuid-2.40.2-1.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "libuuid-devel-0__2.40.2-1.fc40.x86_64",
        sha256 = "0111f4bc7fc5f2cde04915290dace0b81fd3fdc2bd7a3e5234a871bda1323155",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/l/libuuid-devel-2.40.2-1.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/l/libuuid-devel-2.40.2-1.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/l/libuuid-devel-2.40.2-1.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/l/libuuid-devel-2.40.2-1.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/l/libuuid-devel-2.40.2-1.fc40.x86_64.rpm",
        ],
    )

    rpm(
        name = "libverto-0__0.3.2-8.fc40.x86_64",
        sha256 = "fadf7dd93c5eee57ba78e0628bf041dbd2ea037ace52f0a5cbac55b363234d27",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libverto-0.3.2-8.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libverto-0.3.2-8.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libverto-0.3.2-8.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libverto-0.3.2-8.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libverto-0.3.2-8.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "libxcrypt-0__4.4.36-5.fc40.x86_64",
        sha256 = "26c27a101cf40f84f313d81a28cbca9450e8d901e6fcd315ac6036895a369b92",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libxcrypt-4.4.36-5.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libxcrypt-4.4.36-5.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libxcrypt-4.4.36-5.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libxcrypt-4.4.36-5.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libxcrypt-4.4.36-5.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "libxcrypt-devel-0__4.4.36-5.fc40.x86_64",
        sha256 = "0b384c64ba6bf1c067f6b389181c31088472b9a2bc21ef856953558ae19f8ad5",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libxcrypt-devel-4.4.36-5.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libxcrypt-devel-4.4.36-5.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libxcrypt-devel-4.4.36-5.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libxcrypt-devel-4.4.36-5.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libxcrypt-devel-4.4.36-5.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "libxcrypt-static-0__4.4.36-5.fc40.x86_64",
        sha256 = "e91625081401b39814609635fd87288f747f809da0b61f500527f7e75dc61498",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libxcrypt-static-4.4.36-5.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libxcrypt-static-4.4.36-5.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libxcrypt-static-4.4.36-5.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libxcrypt-static-4.4.36-5.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libxcrypt-static-4.4.36-5.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "libxdp-0__1.4.2-1.fc40.x86_64",
        sha256 = "897efd3e7c74a7ffd7b52aea46a2805b7ec4bb9a0be3e20110f5dd696053df5f",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libxdp-1.4.2-1.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libxdp-1.4.2-1.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libxdp-1.4.2-1.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libxdp-1.4.2-1.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/libxdp-1.4.2-1.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "libxml2-0__2.12.8-1.fc40.x86_64",
        sha256 = "ed8d18570524445954dae5aff6239d9cc987cf8b3313fcd48c42f1b79b8eb247",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/l/libxml2-2.12.8-1.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/l/libxml2-2.12.8-1.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/l/libxml2-2.12.8-1.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/l/libxml2-2.12.8-1.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/l/libxml2-2.12.8-1.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "libzstd-0__1.5.6-1.fc40.x86_64",
        sha256 = "bed3075b9ff919eded25cb45e9e03b8a7c63bcc8e893ec28c999aecaa68c51d3",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/l/libzstd-1.5.6-1.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/l/libzstd-1.5.6-1.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/l/libzstd-1.5.6-1.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/l/libzstd-1.5.6-1.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/l/libzstd-1.5.6-1.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "libzstd-devel-0__1.5.6-1.fc40.x86_64",
        sha256 = "6f019d4b621ac4698ee078c9dba9c2ff98f3031a74b54dda3fdd8c30ef142bc9",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/l/libzstd-devel-1.5.6-1.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/l/libzstd-devel-1.5.6-1.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/l/libzstd-devel-1.5.6-1.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/l/libzstd-devel-1.5.6-1.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/l/libzstd-devel-1.5.6-1.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "lld-0__18.1.8-1.fc40.x86_64",
        sha256 = "bb1237b03528f3855ebc22f3c435ba7552d479485bd8e960fb3da8c952a3de30",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/l/lld-18.1.8-1.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/l/lld-18.1.8-1.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/l/lld-18.1.8-1.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/l/lld-18.1.8-1.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/l/lld-18.1.8-1.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "lld-libs-0__18.1.8-1.fc40.x86_64",
        sha256 = "9895294630322ab3f02d5dfa5f4341d58d5b13eb3dacaecc8dc738ec9b1adabd",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/l/lld-libs-18.1.8-1.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/l/lld-libs-18.1.8-1.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/l/lld-libs-18.1.8-1.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/l/lld-libs-18.1.8-1.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/l/lld-libs-18.1.8-1.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "llvm-0__18.1.8-2.fc40.x86_64",
        sha256 = "e7c0a1c4ed10275f0619a644b698add935a38d5ea9d0792d3739bba1cc4f7e43",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/l/llvm-18.1.8-2.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/l/llvm-18.1.8-2.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/l/llvm-18.1.8-2.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/l/llvm-18.1.8-2.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/l/llvm-18.1.8-2.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "llvm-libs-0__18.1.8-2.fc40.x86_64",
        sha256 = "2bb7b88751f67ceb6e56bf4e31843736e4b156de7025109e57ec1c8943648485",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/l/llvm-libs-18.1.8-2.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/l/llvm-libs-18.1.8-2.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/l/llvm-libs-18.1.8-2.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/l/llvm-libs-18.1.8-2.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/l/llvm-libs-18.1.8-2.fc40.x86_64.rpm",
        ],
    )

    rpm(
        name = "lua-libs-0__5.4.6-5.fc40.x86_64",
        sha256 = "81409455da42a5ffdcf5b8cc711632ce037fec25d5ae00cbfda5010c9db04157",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/lua-libs-5.4.6-5.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/lua-libs-5.4.6-5.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/lua-libs-5.4.6-5.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/lua-libs-5.4.6-5.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/lua-libs-5.4.6-5.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "lua-srpm-macros-0__1-13.fc40.x86_64",
        sha256 = "959030121201a706bc620d311569f15ab81bafdb9e3de94bf763a72df36d15f0",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/lua-srpm-macros-1-13.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/lua-srpm-macros-1-13.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/lua-srpm-macros-1-13.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/lua-srpm-macros-1-13.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/lua-srpm-macros-1-13.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "lz4-0__1.9.4-6.fc40.x86_64",
        sha256 = "65bdae7a87e292a315339ac825e12bc75574f65a1d03709c3944be0adecc0948",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/lz4-1.9.4-6.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/lz4-1.9.4-6.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/lz4-1.9.4-6.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/lz4-1.9.4-6.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/lz4-1.9.4-6.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "lz4-libs-0__1.9.4-6.fc40.x86_64",
        sha256 = "f5f022440c4340b5e7fb1c1dbc382e6b0fd57030b3ff056940f2bb3d254408ec",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/lz4-libs-1.9.4-6.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/lz4-libs-1.9.4-6.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/lz4-libs-1.9.4-6.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/lz4-libs-1.9.4-6.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/lz4-libs-1.9.4-6.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "lzo-0__2.10-12.fc40.x86_64",
        sha256 = "84f01f6d2a134c4a7a6591b68242ed781dbf598e6861ab7acbcf5d77e54dfdac",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/lzo-2.10-12.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/lzo-2.10-12.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/lzo-2.10-12.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/lzo-2.10-12.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/l/lzo-2.10-12.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "m4-0__1.4.19-9.fc40.x86_64",
        sha256 = "cea2880d894f015d80ed2a6dfa9033f3eb154c9f014e7a6de7c24207d462dda7",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/m/m4-1.4.19-9.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/m/m4-1.4.19-9.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/m/m4-1.4.19-9.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/m/m4-1.4.19-9.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/m/m4-1.4.19-9.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "make-1__4.4.1-6.fc40.x86_64",
        sha256 = "a4d2818bc705b4d474552f2461c05740449c9da8e4e9f32c4e4e8eaa6cca2b33",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/m/make-4.4.1-6.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/m/make-4.4.1-6.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/m/make-4.4.1-6.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/m/make-4.4.1-6.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/m/make-4.4.1-6.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "mpdecimal-0__2.5.1-9.fc40.x86_64",
        sha256 = "0a3a3fc2471d2d64cbc85f4b23c93620df6eeee814851a2b69fc5ddf75406b56",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/m/mpdecimal-2.5.1-9.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/m/mpdecimal-2.5.1-9.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/m/mpdecimal-2.5.1-9.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/m/mpdecimal-2.5.1-9.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/m/mpdecimal-2.5.1-9.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "mpfr-0__4.2.1-4.fc40.x86_64",
        sha256 = "bc873693a8b8423d7f82e329abe207c9160a4c746fea9a32ef2a6ae8c912f227",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/m/mpfr-4.2.1-4.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/m/mpfr-4.2.1-4.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/m/mpfr-4.2.1-4.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/m/mpfr-4.2.1-4.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/m/mpfr-4.2.1-4.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "nasm-0__2.16.01-7.fc40.x86_64",
        sha256 = "f0d620322c87e1c08e4ccdced5b3ff5e1c3dd943f184ed09c358cf3855e9096a",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/n/nasm-2.16.01-7.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/n/nasm-2.16.01-7.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/n/nasm-2.16.01-7.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/n/nasm-2.16.01-7.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/n/nasm-2.16.01-7.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "ncurses-0__6.4-12.20240127.fc40.x86_64",
        sha256 = "f1b6c955652d4d16d267edfae6bc875d73efd2591ac9d476c480ee9d1e4ee42c",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/n/ncurses-6.4-12.20240127.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/n/ncurses-6.4-12.20240127.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/n/ncurses-6.4-12.20240127.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/n/ncurses-6.4-12.20240127.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/n/ncurses-6.4-12.20240127.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "ncurses-base-0__6.4-12.20240127.fc40.x86_64",
        sha256 = "8a93376ce7423bd1a649a13f4b5105f270b4603f5cf3b3e230bdbda7f25dd788",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/n/ncurses-base-6.4-12.20240127.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/n/ncurses-base-6.4-12.20240127.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/n/ncurses-base-6.4-12.20240127.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/n/ncurses-base-6.4-12.20240127.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/n/ncurses-base-6.4-12.20240127.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "ncurses-libs-0__6.4-12.20240127.fc40.x86_64",
        sha256 = "a18edf32e89aefd453998d5d0ec3aa1ea193dac43f80b99db195abd7e8cf1a04",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/n/ncurses-libs-6.4-12.20240127.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/n/ncurses-libs-6.4-12.20240127.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/n/ncurses-libs-6.4-12.20240127.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/n/ncurses-libs-6.4-12.20240127.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/n/ncurses-libs-6.4-12.20240127.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "ndctl-libs-0__79-1.fc40.x86_64",
        sha256 = "29ec842748c0abed59f9ca1b3f9cd7a44ee1de013510093c433d95a0f42715ef",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/n/ndctl-libs-79-1.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/n/ndctl-libs-79-1.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/n/ndctl-libs-79-1.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/n/ndctl-libs-79-1.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/n/ndctl-libs-79-1.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "nettle-0__3.9.1-6.fc40.x86_64",
        sha256 = "16172412cfd45453292e18f84fc57e42a3ce92aca72b47ef7e15b44554049cfe",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/n/nettle-3.9.1-6.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/n/nettle-3.9.1-6.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/n/nettle-3.9.1-6.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/n/nettle-3.9.1-6.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/n/nettle-3.9.1-6.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "npth-0__1.7-1.fc40.x86_64",
        sha256 = "784e0fbc9ccb7087c10f4c41edbed13904f94244ff658f308614abe48cdf0d42",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/n/npth-1.7-1.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/n/npth-1.7-1.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/n/npth-1.7-1.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/n/npth-1.7-1.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/n/npth-1.7-1.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "numactl-libs-0__2.0.16-5.fc40.x86_64",
        sha256 = "637179e6df20168b70ceb5f76e4acc56937e5e1808c0d314f59b30eefbd1a30a",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/n/numactl-libs-2.0.16-5.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/n/numactl-libs-2.0.16-5.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/n/numactl-libs-2.0.16-5.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/n/numactl-libs-2.0.16-5.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/n/numactl-libs-2.0.16-5.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "ocaml-srpm-macros-0__9-3.fc40.x86_64",
        sha256 = "2d35dbd16fb7c9b306792eddea13d5c863a94ce1b9b9e0c8850cf3c571d56b48",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/o/ocaml-srpm-macros-9-3.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/o/ocaml-srpm-macros-9-3.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/o/ocaml-srpm-macros-9-3.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/o/ocaml-srpm-macros-9-3.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/o/ocaml-srpm-macros-9-3.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "openblas-srpm-macros-0__2-16.fc40.x86_64",
        sha256 = "46ee44ca72fab8e04a7d8c379a550466e7ded1c5a714d14764572fc78b1b5cc5",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/o/openblas-srpm-macros-2-16.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/o/openblas-srpm-macros-2-16.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/o/openblas-srpm-macros-2-16.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/o/openblas-srpm-macros-2-16.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/o/openblas-srpm-macros-2-16.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "openldap-0__2.6.7-1.fc40.x86_64",
        sha256 = "b09089231ec94ee1b2dc26e34d8d7f19586d411bc40df7d0e495e559ac2d871a",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/o/openldap-2.6.7-1.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/o/openldap-2.6.7-1.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/o/openldap-2.6.7-1.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/o/openldap-2.6.7-1.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/o/openldap-2.6.7-1.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "openssl-devel-1__3.2.2-3.fc40.x86_64",
        sha256 = "f08ecb4ad0b491f0be22d13f156a4bc8dd39d9507b7b7550a91d7311dea49dca",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/o/openssl-devel-3.2.2-3.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/o/openssl-devel-3.2.2-3.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/o/openssl-devel-3.2.2-3.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/o/openssl-devel-3.2.2-3.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/o/openssl-devel-3.2.2-3.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "openssl-libs-1__3.2.2-3.fc40.x86_64",
        sha256 = "e9fca52d76eb6277b9fec3238226faafc0938806318fad1143a527fdd28a16cf",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/o/openssl-libs-3.2.2-3.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/o/openssl-libs-3.2.2-3.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/o/openssl-libs-3.2.2-3.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/o/openssl-libs-3.2.2-3.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/o/openssl-libs-3.2.2-3.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "p11-kit-0__0.25.5-1.fc40.x86_64",
        sha256 = "70fba929aab38a9d69a457cef1b01962161a1df2b78dc5a4e86ff4b994b51079",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/p/p11-kit-0.25.5-1.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/p/p11-kit-0.25.5-1.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/p/p11-kit-0.25.5-1.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/p/p11-kit-0.25.5-1.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/p/p11-kit-0.25.5-1.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "p11-kit-trust-0__0.25.5-1.fc40.x86_64",
        sha256 = "c728dbd90872b7597a8ace70a70555bff576231bb6dbde14b75626d601706af8",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/p/p11-kit-trust-0.25.5-1.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/p/p11-kit-trust-0.25.5-1.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/p/p11-kit-trust-0.25.5-1.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/p/p11-kit-trust-0.25.5-1.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/p/p11-kit-trust-0.25.5-1.fc40.x86_64.rpm",
        ],
    )

    rpm(
        name = "package-notes-srpm-macros-0__0.5-11.fc40.x86_64",
        sha256 = "fb4d7c9f138a9ca7cc6fcb68b0820a99a4d67ee22041b64223430f70cee0240a",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/package-notes-srpm-macros-0.5-11.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/package-notes-srpm-macros-0.5-11.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/package-notes-srpm-macros-0.5-11.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/package-notes-srpm-macros-0.5-11.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/package-notes-srpm-macros-0.5-11.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "pam-0__1.6.1-3.fc40.x86_64",
        sha256 = "33d36e10f465b7b15de75ae1856b403ed37c23f026e3abb80497e496f43718c9",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/p/pam-1.6.1-3.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/p/pam-1.6.1-3.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/p/pam-1.6.1-3.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/p/pam-1.6.1-3.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/p/pam-1.6.1-3.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "pam-libs-0__1.6.1-3.fc40.x86_64",
        sha256 = "fb85b93438336461a0b2b878158e552d30b6fb663404475eb0a050b35fd5d35f",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/p/pam-libs-1.6.1-3.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/p/pam-libs-1.6.1-3.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/p/pam-libs-1.6.1-3.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/p/pam-libs-1.6.1-3.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/p/pam-libs-1.6.1-3.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "patch-0__2.7.6-24.fc40.x86_64",
        sha256 = "6cfc586d1f22841f3c36d5a090f011308414af27498371c9701c556ca929d6ed",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/patch-2.7.6-24.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/patch-2.7.6-24.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/patch-2.7.6-24.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/patch-2.7.6-24.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/patch-2.7.6-24.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "pcre2-0__10.44-1.fc40.x86_64",
        sha256 = "73e50df09266fcffda9c24a3738f579dd365c2c187c294da054ef9915edc3851",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/p/pcre2-10.44-1.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/p/pcre2-10.44-1.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/p/pcre2-10.44-1.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/p/pcre2-10.44-1.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/p/pcre2-10.44-1.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "pcre2-syntax-0__10.44-1.fc40.x86_64",
        sha256 = "dbec699e88d42fc6fb1df0a8c0b9023941ed1b1b7625694253a612eaf9f2691d",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/p/pcre2-syntax-10.44-1.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/p/pcre2-syntax-10.44-1.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/p/pcre2-syntax-10.44-1.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/p/pcre2-syntax-10.44-1.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/p/pcre2-syntax-10.44-1.fc40.noarch.rpm",
        ],
    )

    rpm(
        name = "perl-4__5.38.2-506.fc40.x86_64",
        sha256 = "fcdd8d24e2860db8d909b0bb01a9de66775babdc619aa0fdbd23417879b24695",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-5.38.2-506.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-5.38.2-506.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-5.38.2-506.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-5.38.2-506.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-5.38.2-506.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "perl-Algorithm-Diff-0__1.2010-11.fc40.x86_64",
        sha256 = "0c15f155ad3f9ca02482bf70b0d1fd640f2932a5964070106a4a90c62298491e",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Algorithm-Diff-1.2010-11.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Algorithm-Diff-1.2010-11.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Algorithm-Diff-1.2010-11.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Algorithm-Diff-1.2010-11.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Algorithm-Diff-1.2010-11.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Archive-Tar-0__3.02-6.fc40.x86_64",
        sha256 = "b0a57e6b4b9154afea01eb697884b6d30e354258c8ef954ce1a23e6d1603e0a9",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Archive-Tar-3.02-6.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Archive-Tar-3.02-6.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Archive-Tar-3.02-6.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Archive-Tar-3.02-6.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Archive-Tar-3.02-6.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Archive-Zip-0__1.68-14.fc40.x86_64",
        sha256 = "9df5357450fe34cff0c525e54ce7979e990d0da18460a09c65a404d23f3cb89a",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Archive-Zip-1.68-14.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Archive-Zip-1.68-14.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Archive-Zip-1.68-14.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Archive-Zip-1.68-14.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Archive-Zip-1.68-14.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Attribute-Handlers-0__1.03-506.fc40.x86_64",
        sha256 = "c750bbc0d76b38dce225fda305af3728713016af40aa0cc355c01dc984a5df22",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Attribute-Handlers-1.03-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Attribute-Handlers-1.03-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Attribute-Handlers-1.03-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Attribute-Handlers-1.03-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Attribute-Handlers-1.03-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-AutoLoader-0__5.74-506.fc40.x86_64",
        sha256 = "e801f69aa7745987f84f0ad8efa626bd3aea5fb29dc277ed6a5ab157de8878cb",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-AutoLoader-5.74-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-AutoLoader-5.74-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-AutoLoader-5.74-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-AutoLoader-5.74-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-AutoLoader-5.74-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-AutoSplit-0__5.74-506.fc40.x86_64",
        sha256 = "7b4762208435d31674648fddf6556db91ff41fa814f45174b215c0ff2049d1d6",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-AutoSplit-5.74-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-AutoSplit-5.74-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-AutoSplit-5.74-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-AutoSplit-5.74-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-AutoSplit-5.74-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-B-0__1.88-506.fc40.x86_64",
        sha256 = "38771652f69722bfeb1df019e3204f40f242f93599632564a72286c7a7dedb41",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-B-1.88-506.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-B-1.88-506.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-B-1.88-506.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-B-1.88-506.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-B-1.88-506.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "perl-Benchmark-0__1.24-506.fc40.x86_64",
        sha256 = "68f04c9f6fcab675933a7f498efe2679a2b214d2f53fea80fb4908981b706329",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Benchmark-1.24-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Benchmark-1.24-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Benchmark-1.24-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Benchmark-1.24-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Benchmark-1.24-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-CPAN-0__2.36-503.fc40.x86_64",
        sha256 = "4b9740e2e7013a95a9962e0c287dd238e8df77336bcb62d32acccf01081aacbc",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-CPAN-2.36-503.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-CPAN-2.36-503.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-CPAN-2.36-503.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-CPAN-2.36-503.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-CPAN-2.36-503.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-CPAN-Meta-0__2.150010-502.fc40.x86_64",
        sha256 = "e8ee0fffaa79bda65bb25b0d51483692c44541982d432a8c25fd650bb8d8ade1",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-CPAN-Meta-2.150010-502.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-CPAN-Meta-2.150010-502.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-CPAN-Meta-2.150010-502.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-CPAN-Meta-2.150010-502.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-CPAN-Meta-2.150010-502.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-CPAN-Meta-Requirements-0__2.143-6.fc40.x86_64",
        sha256 = "5afb26f93a93f7ef39d06344f211688011ece7f15a063a951e7745559452b4ff",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-CPAN-Meta-Requirements-2.143-6.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-CPAN-Meta-Requirements-2.143-6.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-CPAN-Meta-Requirements-2.143-6.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-CPAN-Meta-Requirements-2.143-6.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-CPAN-Meta-Requirements-2.143-6.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-CPAN-Meta-YAML-0__0.018-503.fc40.x86_64",
        sha256 = "8f6613063103ec5d7c588a33a25b956fa340f28df6c5cd5eedd1f67c8f07cd44",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-CPAN-Meta-YAML-0.018-503.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-CPAN-Meta-YAML-0.018-503.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-CPAN-Meta-YAML-0.018-503.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-CPAN-Meta-YAML-0.018-503.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-CPAN-Meta-YAML-0.018-503.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Carp-0__1.54-502.fc40.x86_64",
        sha256 = "a65dd82703e0c5847733f52fcef81d82528381edbc84bf665a7bf53732e7b126",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Carp-1.54-502.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Carp-1.54-502.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Carp-1.54-502.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Carp-1.54-502.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Carp-1.54-502.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Class-Struct-0__0.68-506.fc40.x86_64",
        sha256 = "50c21b40deb69eb7e726f7d9c68e27af906e1eda028559d6e16364ded3625a16",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Class-Struct-0.68-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Class-Struct-0.68-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Class-Struct-0.68-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Class-Struct-0.68-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Class-Struct-0.68-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Compress-Bzip2-0__2.28-17.fc40.x86_64",
        sha256 = "53bc59ca25a1166b66298fa0145fc561fb54da2c5b415a9d4ff7dbaed7f990f1",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Compress-Bzip2-2.28-17.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Compress-Bzip2-2.28-17.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Compress-Bzip2-2.28-17.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Compress-Bzip2-2.28-17.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Compress-Bzip2-2.28-17.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "perl-Compress-Raw-Bzip2-0__2.210-1.fc40.x86_64",
        sha256 = "dbafee1e04a29af92e28073aef199577c9b4d3c4a789ecc6be03ad2d88dff53a",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Compress-Raw-Bzip2-2.210-1.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Compress-Raw-Bzip2-2.210-1.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Compress-Raw-Bzip2-2.210-1.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Compress-Raw-Bzip2-2.210-1.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Compress-Raw-Bzip2-2.210-1.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "perl-Compress-Raw-Lzma-0__2.209-8.fc40.x86_64",
        sha256 = "75af0b81aa95649d228645cf061825487cd9fcc03645017ff7b94348ff78bd0f",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Compress-Raw-Lzma-2.209-8.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Compress-Raw-Lzma-2.209-8.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Compress-Raw-Lzma-2.209-8.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Compress-Raw-Lzma-2.209-8.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Compress-Raw-Lzma-2.209-8.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "perl-Compress-Raw-Zlib-0__2.209-1.fc40.x86_64",
        sha256 = "1563b4a49fe678a94bf3237c84b79c5c891c2fcb4b29966660256978cb5face4",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Compress-Raw-Zlib-2.209-1.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Compress-Raw-Zlib-2.209-1.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Compress-Raw-Zlib-2.209-1.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Compress-Raw-Zlib-2.209-1.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Compress-Raw-Zlib-2.209-1.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "perl-Config-Extensions-0__0.03-506.fc40.x86_64",
        sha256 = "acddd0dedec20bc5a7e3e208006e8bd73f887d566d1eec9fdcb6a1061ba0c359",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Config-Extensions-0.03-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Config-Extensions-0.03-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Config-Extensions-0.03-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Config-Extensions-0.03-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Config-Extensions-0.03-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Config-Perl-V-0__0.36-503.fc40.x86_64",
        sha256 = "e619113c6ec1e04dd15968acb7438c6c89f8feb7311e7e5a244f538339dadced",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Config-Perl-V-0.36-503.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Config-Perl-V-0.36-503.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Config-Perl-V-0.36-503.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Config-Perl-V-0.36-503.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Config-Perl-V-0.36-503.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-DBM_Filter-0__0.06-506.fc40.x86_64",
        sha256 = "0a898e82d05169278e37bd626113ea11f5a516b712b0018c4dca82d3f09d3563",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-DBM_Filter-0.06-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-DBM_Filter-0.06-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-DBM_Filter-0.06-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-DBM_Filter-0.06-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-DBM_Filter-0.06-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-DB_File-0__1.859-3.fc40.x86_64",
        sha256 = "034560d8a731a3203a6281d164dd3fad812eb4accaa08fa2cb2f1058ce9b21ee",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-DB_File-1.859-3.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-DB_File-1.859-3.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-DB_File-1.859-3.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-DB_File-1.859-3.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-DB_File-1.859-3.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "perl-Data-Dumper-0__2.188-503.fc40.x86_64",
        sha256 = "7aa596dde17ad70508e7e0e75c154973290a14d1fb56057aa6f580907e837555",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Data-Dumper-2.188-503.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Data-Dumper-2.188-503.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Data-Dumper-2.188-503.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Data-Dumper-2.188-503.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Data-Dumper-2.188-503.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "perl-Data-OptList-0__0.114-4.fc40.x86_64",
        sha256 = "d284d509c99a24e7a4d60d03a9f31dc3be868f2fcc519849defb3351e480a260",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Data-OptList-0.114-4.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Data-OptList-0.114-4.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Data-OptList-0.114-4.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Data-OptList-0.114-4.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Data-OptList-0.114-4.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Data-Section-0__0.200008-5.fc40.x86_64",
        sha256 = "a54ccaa5da958d8988238e6f5c05196dc287a22ae3eba9eba41f72fa11bd46eb",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Data-Section-0.200008-5.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Data-Section-0.200008-5.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Data-Section-0.200008-5.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Data-Section-0.200008-5.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Data-Section-0.200008-5.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Devel-PPPort-0__3.71-503.fc40.x86_64",
        sha256 = "73ba4f2a78ff1d3f5d7b9838d5965b988de41969bdf8b88d9699962cf21017c9",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Devel-PPPort-3.71-503.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Devel-PPPort-3.71-503.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Devel-PPPort-3.71-503.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Devel-PPPort-3.71-503.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Devel-PPPort-3.71-503.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "perl-Devel-Peek-0__1.33-506.fc40.x86_64",
        sha256 = "7e3ae32d01316158602a1a802baee3531116eae6760bc67a3fa84052d6a5a501",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Devel-Peek-1.33-506.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Devel-Peek-1.33-506.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Devel-Peek-1.33-506.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Devel-Peek-1.33-506.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Devel-Peek-1.33-506.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "perl-Devel-SelfStubber-0__1.06-506.fc40.x86_64",
        sha256 = "085b644fbbf13a82f4848172149d6ffdb027c82ccd1c3b80a5b754bfda6d3c4d",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Devel-SelfStubber-1.06-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Devel-SelfStubber-1.06-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Devel-SelfStubber-1.06-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Devel-SelfStubber-1.06-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Devel-SelfStubber-1.06-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Devel-Size-0__0.84-1.fc40.x86_64",
        sha256 = "fa058e61d3844f6b53b160ec54ddbde4e9755d41e2f21caf90688b4f52c2c019",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/p/perl-Devel-Size-0.84-1.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/p/perl-Devel-Size-0.84-1.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/p/perl-Devel-Size-0.84-1.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/p/perl-Devel-Size-0.84-1.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/p/perl-Devel-Size-0.84-1.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "perl-Digest-0__1.20-502.fc40.x86_64",
        sha256 = "7a3227717f0121273607249d64ea56953f0a1d68eb37e690e8c5f85851e2467a",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Digest-1.20-502.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Digest-1.20-502.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Digest-1.20-502.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Digest-1.20-502.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Digest-1.20-502.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Digest-MD5-0__2.59-3.fc40.x86_64",
        sha256 = "af36bb36832421d2092678612d893d27a58f678500ece3a9d268aad43326be59",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Digest-MD5-2.59-3.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Digest-MD5-2.59-3.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Digest-MD5-2.59-3.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Digest-MD5-2.59-3.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Digest-MD5-2.59-3.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "perl-Digest-SHA-1__6.04-503.fc40.x86_64",
        sha256 = "8fde080587a5dde4c175c06632be06d9a0e4731277d15a13c7c91abc8a85fb1a",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Digest-SHA-6.04-503.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Digest-SHA-6.04-503.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Digest-SHA-6.04-503.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Digest-SHA-6.04-503.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Digest-SHA-6.04-503.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "perl-DirHandle-0__1.05-506.fc40.x86_64",
        sha256 = "6512c08bd7187fef2d8983e27cb832c0a02cf7fd53197c6615df4a46dce2ae45",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-DirHandle-1.05-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-DirHandle-1.05-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-DirHandle-1.05-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-DirHandle-1.05-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-DirHandle-1.05-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Dumpvalue-0__2.27-506.fc40.x86_64",
        sha256 = "bd8f2e9f28453f1723840d323004336fb7ad5e09b9514cfa33493732104b5b4e",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Dumpvalue-2.27-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Dumpvalue-2.27-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Dumpvalue-2.27-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Dumpvalue-2.27-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Dumpvalue-2.27-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-DynaLoader-0__1.54-506.fc40.x86_64",
        sha256 = "d38041cf63f98515ade9e0af650b4564dc479f0188686a4abc4a2cd533b2c360",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-DynaLoader-1.54-506.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-DynaLoader-1.54-506.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-DynaLoader-1.54-506.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-DynaLoader-1.54-506.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-DynaLoader-1.54-506.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "perl-Encode-4__3.21-505.fc40.x86_64",
        sha256 = "d0abddb488efa8dc53b1d7c24a9cad49ac9160d6a8180c500884a65ff8cce2a8",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Encode-3.21-505.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Encode-3.21-505.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Encode-3.21-505.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Encode-3.21-505.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Encode-3.21-505.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "perl-Encode-devel-4__3.21-505.fc40.x86_64",
        sha256 = "a0cce51b2b73c24bf18c1e50a25774286bc68cacef6a0a03295026fbf7185fd9",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Encode-devel-3.21-505.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Encode-devel-3.21-505.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Encode-devel-3.21-505.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Encode-devel-3.21-505.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Encode-devel-3.21-505.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "perl-English-0__1.11-506.fc40.x86_64",
        sha256 = "0a1691e74e99e9d253ce17ba4b608da68be9831b05211d8f8d4ec7f4899642a9",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-English-1.11-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-English-1.11-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-English-1.11-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-English-1.11-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-English-1.11-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Env-0__1.06-502.fc40.x86_64",
        sha256 = "6698622f465a4f06b95351db7f8f2a11b785fad5bd414bb5cc8ccd6c0211c7ad",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Env-1.06-502.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Env-1.06-502.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Env-1.06-502.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Env-1.06-502.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Env-1.06-502.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Errno-0__1.37-506.fc40.x86_64",
        sha256 = "37b9ad9c6aef1169ff9cba209df8351408abe01b05ef2262f091dc210351ce17",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Errno-1.37-506.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Errno-1.37-506.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Errno-1.37-506.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Errno-1.37-506.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Errno-1.37-506.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "perl-Exporter-0__5.78-3.fc40.x86_64",
        sha256 = "8647c554687dbe5dbe010fbc826e897d4cc9a8c691c942e4633195645858ad10",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Exporter-5.78-3.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Exporter-5.78-3.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Exporter-5.78-3.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Exporter-5.78-3.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Exporter-5.78-3.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-ExtUtils-CBuilder-1__0.280238-502.fc40.x86_64",
        sha256 = "e0d83da248ed00c61a8b5097839882cb6bd5279cca04220a02a7a27d2ce93f3a",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-ExtUtils-CBuilder-0.280238-502.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-ExtUtils-CBuilder-0.280238-502.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-ExtUtils-CBuilder-0.280238-502.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-ExtUtils-CBuilder-0.280238-502.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-ExtUtils-CBuilder-0.280238-502.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-ExtUtils-Command-2__7.70-503.fc40.x86_64",
        sha256 = "b76c4c2222b98f38a12630bd6f7d2ea17b6fb39443091f6040baaa1d4d974cd5",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-ExtUtils-Command-7.70-503.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-ExtUtils-Command-7.70-503.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-ExtUtils-Command-7.70-503.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-ExtUtils-Command-7.70-503.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-ExtUtils-Command-7.70-503.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-ExtUtils-Constant-0__0.25-506.fc40.x86_64",
        sha256 = "59a6b0b8c0fd768d6e854adeb6e916162849711d82625d1250c7448dca91a2b1",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-ExtUtils-Constant-0.25-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-ExtUtils-Constant-0.25-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-ExtUtils-Constant-0.25-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-ExtUtils-Constant-0.25-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-ExtUtils-Constant-0.25-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-ExtUtils-Embed-0__1.35-506.fc40.x86_64",
        sha256 = "3b4ea4b7f3d36dfaf042cdf17eee592ffe10888ff654ee0e2d6b064cd5d7fe94",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-ExtUtils-Embed-1.35-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-ExtUtils-Embed-1.35-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-ExtUtils-Embed-1.35-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-ExtUtils-Embed-1.35-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-ExtUtils-Embed-1.35-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-ExtUtils-Install-0__2.22-502.fc40.x86_64",
        sha256 = "43c50ea47a2d6ce1de18c1cad2d753475165a397e8d4546341a8123e365eaceb",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-ExtUtils-Install-2.22-502.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-ExtUtils-Install-2.22-502.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-ExtUtils-Install-2.22-502.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-ExtUtils-Install-2.22-502.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-ExtUtils-Install-2.22-502.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-ExtUtils-MM-Utils-2__7.70-503.fc40.x86_64",
        sha256 = "f56cf2535b3c1a8a79344ac6abc6b2408f25cad47477b00879b14161a0867296",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-ExtUtils-MM-Utils-7.70-503.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-ExtUtils-MM-Utils-7.70-503.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-ExtUtils-MM-Utils-7.70-503.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-ExtUtils-MM-Utils-7.70-503.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-ExtUtils-MM-Utils-7.70-503.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-ExtUtils-MakeMaker-2__7.70-503.fc40.x86_64",
        sha256 = "08d3b88f88f4fb666f96eda69a279543e2fef2e3f3dfc4da4502e9d38dddf16d",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-ExtUtils-MakeMaker-7.70-503.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-ExtUtils-MakeMaker-7.70-503.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-ExtUtils-MakeMaker-7.70-503.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-ExtUtils-MakeMaker-7.70-503.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-ExtUtils-MakeMaker-7.70-503.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-ExtUtils-Manifest-1__1.75-5.fc40.x86_64",
        sha256 = "8c611f2a3d560bbb219a556bb4eb0c9a7fbf45d38f07ff228cbb7f0fe918e2df",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-ExtUtils-Manifest-1.75-5.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-ExtUtils-Manifest-1.75-5.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-ExtUtils-Manifest-1.75-5.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-ExtUtils-Manifest-1.75-5.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-ExtUtils-Manifest-1.75-5.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-ExtUtils-Miniperl-0__1.13-506.fc40.x86_64",
        sha256 = "536f08900d01b6ae7fb6fa1c0ed4243c875aa6475333a5ca4cffc3d24b4dbd03",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-ExtUtils-Miniperl-1.13-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-ExtUtils-Miniperl-1.13-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-ExtUtils-Miniperl-1.13-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-ExtUtils-Miniperl-1.13-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-ExtUtils-Miniperl-1.13-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-ExtUtils-ParseXS-1__3.51-503.fc40.x86_64",
        sha256 = "9bf5620bbd381fe0257b9da589cba7c1c919df199c7ec643af8e52da0bca7bd6",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-ExtUtils-ParseXS-3.51-503.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-ExtUtils-ParseXS-3.51-503.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-ExtUtils-ParseXS-3.51-503.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-ExtUtils-ParseXS-3.51-503.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-ExtUtils-ParseXS-3.51-503.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Fcntl-0__1.15-506.fc40.x86_64",
        sha256 = "18f5fdba18cd6e222b85613d2288867e9da41c6bbd3608e7e9f830caba246ea0",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Fcntl-1.15-506.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Fcntl-1.15-506.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Fcntl-1.15-506.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Fcntl-1.15-506.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Fcntl-1.15-506.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "perl-File-Basename-0__2.86-506.fc40.x86_64",
        sha256 = "57164886c006d71b81f930735730ba1bbe56354558596cc582bdebb269d9f2d3",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-File-Basename-2.86-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-File-Basename-2.86-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-File-Basename-2.86-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-File-Basename-2.86-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-File-Basename-2.86-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-File-Compare-0__1.100.700-506.fc40.x86_64",
        sha256 = "17150da64c7c1cf7e03d63b47c4d45b5ed2c5c8a40a52f89c5dbad2d04380c5c",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-File-Compare-1.100.700-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-File-Compare-1.100.700-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-File-Compare-1.100.700-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-File-Compare-1.100.700-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-File-Compare-1.100.700-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-File-Copy-0__2.41-506.fc40.x86_64",
        sha256 = "9f074bce639cbb4e5d7e76466ff2106fb8dc3cd5adbfaadb415051570ab57bbf",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-File-Copy-2.41-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-File-Copy-2.41-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-File-Copy-2.41-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-File-Copy-2.41-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-File-Copy-2.41-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-File-DosGlob-0__1.12-506.fc40.x86_64",
        sha256 = "5d59b93cef21e8789eb8a1c2fbb50345ca56e10e510095aa888b4ee4474f3e3a",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-File-DosGlob-1.12-506.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-File-DosGlob-1.12-506.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-File-DosGlob-1.12-506.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-File-DosGlob-1.12-506.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-File-DosGlob-1.12-506.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "perl-File-Fetch-0__1.04-502.fc40.x86_64",
        sha256 = "7a9f7ab914e85b91852bcd77bddf1cfd0532fbe24c17c080c87618d6a1f97691",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-File-Fetch-1.04-502.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-File-Fetch-1.04-502.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-File-Fetch-1.04-502.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-File-Fetch-1.04-502.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-File-Fetch-1.04-502.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-File-Find-0__1.43-506.fc40.x86_64",
        sha256 = "93c539bd75e3fa4a5656c9e341ab82dea1c43b306ef45ff26ae4c599633dbe14",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-File-Find-1.43-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-File-Find-1.43-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-File-Find-1.43-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-File-Find-1.43-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-File-Find-1.43-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-File-HomeDir-0__1.006-12.fc40.x86_64",
        sha256 = "d22d58d6fe4edee5a549c99ceb89c36f9022d6efff6217161ce30fd2eb34f7cf",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-File-HomeDir-1.006-12.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-File-HomeDir-1.006-12.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-File-HomeDir-1.006-12.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-File-HomeDir-1.006-12.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-File-HomeDir-1.006-12.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-File-Path-0__2.18-503.fc40.x86_64",
        sha256 = "8b152cd78a7f56136fd4d2f3b56111a8e5c4ab8192e50069a38df3eb90cdeba8",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-File-Path-2.18-503.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-File-Path-2.18-503.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-File-Path-2.18-503.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-File-Path-2.18-503.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-File-Path-2.18-503.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-File-Temp-1__0.231.100-503.fc40.x86_64",
        sha256 = "a686fe5e5e94ef9876b429099cd2bc85069a83148ea26b17970443a757822fa4",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-File-Temp-0.231.100-503.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-File-Temp-0.231.100-503.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-File-Temp-0.231.100-503.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-File-Temp-0.231.100-503.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-File-Temp-0.231.100-503.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-File-Which-0__1.27-11.fc40.x86_64",
        sha256 = "689e3c08798d1a4385435f2ee0e69c51509be2290de0e552a3f810b3d0790451",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-File-Which-1.27-11.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-File-Which-1.27-11.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-File-Which-1.27-11.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-File-Which-1.27-11.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-File-Which-1.27-11.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-File-stat-0__1.13-506.fc40.x86_64",
        sha256 = "0149f36e83814763c7937eddbecbc53a15959f3a69f2eaed6b380b07366698f5",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-File-stat-1.13-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-File-stat-1.13-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-File-stat-1.13-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-File-stat-1.13-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-File-stat-1.13-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-FileCache-0__1.10-506.fc40.x86_64",
        sha256 = "a9b0d66b720a3867a2c22504d95047785588fb7ee13b728606cc7845f27d47d0",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-FileCache-1.10-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-FileCache-1.10-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-FileCache-1.10-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-FileCache-1.10-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-FileCache-1.10-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-FileHandle-0__2.05-506.fc40.x86_64",
        sha256 = "8579aa3d0f1827c98678007e84a2ade496275f42884be0cc9f999c2de31a533b",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-FileHandle-2.05-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-FileHandle-2.05-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-FileHandle-2.05-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-FileHandle-2.05-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-FileHandle-2.05-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Filter-2__1.64-503.fc40.x86_64",
        sha256 = "ebe43785e70cedcde209a711e935d1dcca428b1117c668578c21407a79aa5929",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Filter-1.64-503.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Filter-1.64-503.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Filter-1.64-503.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Filter-1.64-503.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Filter-1.64-503.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "perl-Filter-Simple-0__0.96-503.fc40.x86_64",
        sha256 = "08b4bc22ed13283b595ebb153d9bc70e7732d30fb93df561d02a41e3e7136cb1",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Filter-Simple-0.96-503.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Filter-Simple-0.96-503.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Filter-Simple-0.96-503.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Filter-Simple-0.96-503.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Filter-Simple-0.96-503.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-FindBin-0__1.53-506.fc40.x86_64",
        sha256 = "ef6f7bcf631b34bc6092779fd835dbb46532389d08b0b072edb19c0afdd79dfa",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-FindBin-1.53-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-FindBin-1.53-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-FindBin-1.53-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-FindBin-1.53-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-FindBin-1.53-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-GDBM_File-1__1.24-506.fc40.x86_64",
        sha256 = "aeb7aed1b1d0e82b3bd585068ab5595225f9ed136ec26003da6a9a9bb101dddb",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-GDBM_File-1.24-506.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-GDBM_File-1.24-506.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-GDBM_File-1.24-506.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-GDBM_File-1.24-506.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-GDBM_File-1.24-506.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "perl-Getopt-Long-1__2.57-4.fc40.x86_64",
        sha256 = "c61ce353b34a66009027a2d2e0d819a728b02e888a496c5cf8e63b164b731e6e",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/p/perl-Getopt-Long-2.57-4.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/p/perl-Getopt-Long-2.57-4.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/p/perl-Getopt-Long-2.57-4.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/p/perl-Getopt-Long-2.57-4.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/p/perl-Getopt-Long-2.57-4.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Getopt-Std-0__1.13-506.fc40.x86_64",
        sha256 = "4e46c286d79b208e6111ceae585d2d00b835c913686bd1cfd608f8e225e41a40",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Getopt-Std-1.13-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Getopt-Std-1.13-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Getopt-Std-1.13-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Getopt-Std-1.13-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Getopt-Std-1.13-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-HTTP-Tiny-0__0.088-5.fc40.x86_64",
        sha256 = "d0a5c3099349032e4527c11737c9f54ad7427685b563f10be9e6006b2acee36d",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-HTTP-Tiny-0.088-5.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-HTTP-Tiny-0.088-5.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-HTTP-Tiny-0.088-5.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-HTTP-Tiny-0.088-5.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-HTTP-Tiny-0.088-5.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Hash-Util-0__0.30-506.fc40.x86_64",
        sha256 = "f3c7e36f7345cf38e10908b8a88978c6949eb97532d6a958025b51edbccaf363",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Hash-Util-0.30-506.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Hash-Util-0.30-506.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Hash-Util-0.30-506.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Hash-Util-0.30-506.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Hash-Util-0.30-506.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "perl-Hash-Util-FieldHash-0__1.26-506.fc40.x86_64",
        sha256 = "1b3b518b75914bc3376d1eaf56b40dcacf0d1bbb23b7ac4d956e80e668e7673e",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Hash-Util-FieldHash-1.26-506.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Hash-Util-FieldHash-1.26-506.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Hash-Util-FieldHash-1.26-506.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Hash-Util-FieldHash-1.26-506.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Hash-Util-FieldHash-1.26-506.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "perl-I18N-Collate-0__1.02-506.fc40.x86_64",
        sha256 = "d4336be43ce67e258e9743c3184bcb015d985cfc72ee0004ae690451a6f8b6f1",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-I18N-Collate-1.02-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-I18N-Collate-1.02-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-I18N-Collate-1.02-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-I18N-Collate-1.02-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-I18N-Collate-1.02-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-I18N-LangTags-0__0.45-506.fc40.x86_64",
        sha256 = "a5780829fac4152291d60cb5eb06a6a7a1d068b061116b024c4ed73faa4f4e56",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-I18N-LangTags-0.45-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-I18N-LangTags-0.45-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-I18N-LangTags-0.45-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-I18N-LangTags-0.45-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-I18N-LangTags-0.45-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-I18N-Langinfo-0__0.22-506.fc40.x86_64",
        sha256 = "8c066c39c97ebf03a9a7169be3940a37742d345f4a2493d48982f9eef5659706",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-I18N-Langinfo-0.22-506.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-I18N-Langinfo-0.22-506.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-I18N-Langinfo-0.22-506.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-I18N-Langinfo-0.22-506.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-I18N-Langinfo-0.22-506.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "perl-IO-0__1.52-506.fc40.x86_64",
        sha256 = "f154c1c6952d8c2c1340ef1037d6b3da61b1e8468cf9f20d9b5ddebd796c1da6",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-IO-1.52-506.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-IO-1.52-506.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-IO-1.52-506.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-IO-1.52-506.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-IO-1.52-506.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "perl-IO-Compress-0__2.207-1.fc40.x86_64",
        sha256 = "e05036e4a95f2d57814546bad0b2031884fdac6ed88049c03a5e70130626e682",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-IO-Compress-2.207-1.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-IO-Compress-2.207-1.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-IO-Compress-2.207-1.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-IO-Compress-2.207-1.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-IO-Compress-2.207-1.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-IO-Compress-Lzma-0__2.207-1.fc40.x86_64",
        sha256 = "bd4cd6e8050d03ef72f8bc51a653572e0c0ec16b6279b48664d6ad4f729d7608",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-IO-Compress-Lzma-2.207-1.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-IO-Compress-Lzma-2.207-1.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-IO-Compress-Lzma-2.207-1.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-IO-Compress-Lzma-2.207-1.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-IO-Compress-Lzma-2.207-1.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-IO-Socket-IP-0__0.42-2.fc40.x86_64",
        sha256 = "30aa2fa573d6772840ec30431d3e92c78d90442e5349dc9bab14f70816e84ecb",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-IO-Socket-IP-0.42-2.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-IO-Socket-IP-0.42-2.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-IO-Socket-IP-0.42-2.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-IO-Socket-IP-0.42-2.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-IO-Socket-IP-0.42-2.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-IO-Socket-SSL-0__2.085-1.fc40.x86_64",
        sha256 = "660515f32e0985d3ad5d5b58426d77ed07eca255cd30954469dffc6de8516d0e",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-IO-Socket-SSL-2.085-1.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-IO-Socket-SSL-2.085-1.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-IO-Socket-SSL-2.085-1.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-IO-Socket-SSL-2.085-1.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-IO-Socket-SSL-2.085-1.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-IO-Zlib-1__1.15-1.fc40.x86_64",
        sha256 = "282d7aef6b5cad631a03a5a0a28f7302d9ad52a4922b29783ed2d99d5ca0a1fe",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-IO-Zlib-1.15-1.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-IO-Zlib-1.15-1.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-IO-Zlib-1.15-1.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-IO-Zlib-1.15-1.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-IO-Zlib-1.15-1.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-IPC-Cmd-2__1.04-504.fc40.x86_64",
        sha256 = "3a24352aaab55dea0deb0397bc3fd5edce2eee35f34e2e1eaefbd8f026d4d032",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-IPC-Cmd-1.04-504.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-IPC-Cmd-1.04-504.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-IPC-Cmd-1.04-504.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-IPC-Cmd-1.04-504.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-IPC-Cmd-1.04-504.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-IPC-Open3-0__1.22-506.fc40.x86_64",
        sha256 = "c50d0a81b90d11648ea010d72fcb924cd910d9b9c247021c152e1f7ee5ee4e46",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-IPC-Open3-1.22-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-IPC-Open3-1.22-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-IPC-Open3-1.22-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-IPC-Open3-1.22-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-IPC-Open3-1.22-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-IPC-SysV-0__2.09-505.fc40.x86_64",
        sha256 = "01ecb4e9ae94da68f2894b9943f170c3db77d7c04e8f53085ce33ca3c06883f6",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-IPC-SysV-2.09-505.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-IPC-SysV-2.09-505.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-IPC-SysV-2.09-505.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-IPC-SysV-2.09-505.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-IPC-SysV-2.09-505.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "perl-IPC-System-Simple-0__1.30-13.fc40.x86_64",
        sha256 = "02af6f37e13d21d516a6e152ed6cee163c305975aba24aef4da38d5a1846ecd1",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-IPC-System-Simple-1.30-13.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-IPC-System-Simple-1.30-13.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-IPC-System-Simple-1.30-13.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-IPC-System-Simple-1.30-13.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-IPC-System-Simple-1.30-13.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-JSON-PP-1__4.16-503.fc40.x86_64",
        sha256 = "e3072a5d7b5325c3ded189bb78582231a45ff9dc70f6f27a42ef9a3388dddceb",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-JSON-PP-4.16-503.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-JSON-PP-4.16-503.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-JSON-PP-4.16-503.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-JSON-PP-4.16-503.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-JSON-PP-4.16-503.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Locale-Maketext-0__1.33-503.fc40.x86_64",
        sha256 = "8cc64314cb0124b97c629b0ace1d1f8fc37bf3aa15a654fb1115cf0fb1713386",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Locale-Maketext-1.33-503.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Locale-Maketext-1.33-503.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Locale-Maketext-1.33-503.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Locale-Maketext-1.33-503.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Locale-Maketext-1.33-503.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Locale-Maketext-Simple-1__0.21-506.fc40.x86_64",
        sha256 = "50007a4c207fb30ab09013406d98d2510433786d1deed1bb56bf1abf84e8fbec",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Locale-Maketext-Simple-0.21-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Locale-Maketext-Simple-0.21-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Locale-Maketext-Simple-0.21-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Locale-Maketext-Simple-0.21-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Locale-Maketext-Simple-0.21-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-MIME-Base64-0__3.16-503.fc40.x86_64",
        sha256 = "46d39e3f41e3fef98bc85c1bc5237555f21d23453833d7e57341ed3e324a82bf",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-MIME-Base64-3.16-503.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-MIME-Base64-3.16-503.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-MIME-Base64-3.16-503.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-MIME-Base64-3.16-503.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-MIME-Base64-3.16-503.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "perl-MRO-Compat-0__0.15-9.fc40.x86_64",
        sha256 = "c714ded9c6fc9a4bc5cee122808d5501c0fe2d443fdfe225705efb2668a61e01",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-MRO-Compat-0.15-9.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-MRO-Compat-0.15-9.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-MRO-Compat-0.15-9.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-MRO-Compat-0.15-9.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-MRO-Compat-0.15-9.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Math-BigInt-1__2.0030.03-1.fc40.x86_64",
        sha256 = "8e31a5a14c24675aef8417af535128ec6096c8f204a1983a773a9e74573aeee3",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/p/perl-Math-BigInt-2.0030.03-1.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/p/perl-Math-BigInt-2.0030.03-1.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/p/perl-Math-BigInt-2.0030.03-1.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/p/perl-Math-BigInt-2.0030.03-1.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/p/perl-Math-BigInt-2.0030.03-1.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Math-BigInt-FastCalc-0__0.501.800-3.fc40.x86_64",
        sha256 = "5417dcb8e5f84b37e18d4f8eb6b60cec0ed98c596852c6ce561a23c7617c133e",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Math-BigInt-FastCalc-0.501.800-3.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Math-BigInt-FastCalc-0.501.800-3.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Math-BigInt-FastCalc-0.501.800-3.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Math-BigInt-FastCalc-0.501.800-3.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Math-BigInt-FastCalc-0.501.800-3.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "perl-Math-Complex-0__1.62-506.fc40.x86_64",
        sha256 = "76fff240e889f8fceae4e362cfd6c65eb9ffd66a4eb39ffbc4e7923eb8e9feed",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Math-Complex-1.62-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Math-Complex-1.62-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Math-Complex-1.62-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Math-Complex-1.62-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Math-Complex-1.62-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Memoize-0__1.16-506.fc40.x86_64",
        sha256 = "5c0be04f2435f3de7363841a93972de8a550102e812a0a6ea8ea352373d6a8f7",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Memoize-1.16-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Memoize-1.16-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Memoize-1.16-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Memoize-1.16-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Memoize-1.16-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Module-Build-2__0.42.34-5.fc40.x86_64",
        sha256 = "672439cd1b937b4ff1687138d83026d2b491aec1d20fcd0a8ca97dffca005024",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Module-Build-0.42.34-5.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Module-Build-0.42.34-5.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Module-Build-0.42.34-5.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Module-Build-0.42.34-5.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Module-Build-0.42.34-5.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Module-CoreList-1__5.20240920-1.fc40.x86_64",
        sha256 = "9ab2aa8e0cb3afd52243ce91d1f209b7befaae0317baec81f8b898efa2936616",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/p/perl-Module-CoreList-5.20240920-1.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/p/perl-Module-CoreList-5.20240920-1.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/p/perl-Module-CoreList-5.20240920-1.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/p/perl-Module-CoreList-5.20240920-1.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/p/perl-Module-CoreList-5.20240920-1.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Module-CoreList-tools-1__5.20240920-1.fc40.x86_64",
        sha256 = "0bdf5bbc6b69bb68c12cc5b1381a2d14215bf2a5b65233d6ea651bb3de03f9ec",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/p/perl-Module-CoreList-tools-5.20240920-1.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/p/perl-Module-CoreList-tools-5.20240920-1.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/p/perl-Module-CoreList-tools-5.20240920-1.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/p/perl-Module-CoreList-tools-5.20240920-1.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/p/perl-Module-CoreList-tools-5.20240920-1.fc40.noarch.rpm",
        ],
    )

    rpm(
        name = "perl-Module-Load-1__0.36-503.fc40.x86_64",
        sha256 = "ffff4d9fa6f9685b36aca24a39f965d4cd94ccb13a4c73e4fec45460733893ef",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Module-Load-0.36-503.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Module-Load-0.36-503.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Module-Load-0.36-503.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Module-Load-0.36-503.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Module-Load-0.36-503.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Module-Load-Conditional-0__0.74-503.fc40.x86_64",
        sha256 = "77bab9d62249a280b9f14e1ae6ed4071dc267234a576cb7fb012533b4ccf116b",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Module-Load-Conditional-0.74-503.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Module-Load-Conditional-0.74-503.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Module-Load-Conditional-0.74-503.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Module-Load-Conditional-0.74-503.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Module-Load-Conditional-0.74-503.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Module-Loaded-1__0.08-506.fc40.x86_64",
        sha256 = "c431dde1f7a5e0af118b300a5e94a965f2376e4fafeda1bac644f225678c6314",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Module-Loaded-0.08-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Module-Loaded-0.08-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Module-Loaded-0.08-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Module-Loaded-0.08-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Module-Loaded-0.08-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Module-Metadata-0__1.000038-5.fc40.x86_64",
        sha256 = "dbefbe5bfd576e1556b257f71deb8dbe83aef33aa639791e1401f74f1de96aca",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Module-Metadata-1.000038-5.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Module-Metadata-1.000038-5.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Module-Metadata-1.000038-5.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Module-Metadata-1.000038-5.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Module-Metadata-1.000038-5.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Module-Signature-0__0.88-9.fc40.x86_64",
        sha256 = "1d89dff0b55c5fdf5aa9abd61552858c7a975a0c51cc5cf25879b12e6fc8f2ef",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Module-Signature-0.88-9.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Module-Signature-0.88-9.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Module-Signature-0.88-9.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Module-Signature-0.88-9.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Module-Signature-0.88-9.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Mozilla-CA-0__20231213-3.fc40.x86_64",
        sha256 = "a53e9503a79437c7585c8f82d2047ad0eb5f53b3a92edfbd218f620d7dd47c98",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Mozilla-CA-20231213-3.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Mozilla-CA-20231213-3.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Mozilla-CA-20231213-3.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Mozilla-CA-20231213-3.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Mozilla-CA-20231213-3.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-NDBM_File-0__1.16-506.fc40.x86_64",
        sha256 = "49656c4a9bdb4daa20e477465e42a368f10a9a2d625df3cea0163c1971e42454",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-NDBM_File-1.16-506.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-NDBM_File-1.16-506.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-NDBM_File-1.16-506.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-NDBM_File-1.16-506.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-NDBM_File-1.16-506.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "perl-NEXT-0__0.69-506.fc40.x86_64",
        sha256 = "d860065135bd49e7b3bddb61dd058482c34d1eb618aeceef80a88059195c029d",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-NEXT-0.69-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-NEXT-0.69-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-NEXT-0.69-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-NEXT-0.69-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-NEXT-0.69-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Net-0__1.03-506.fc40.x86_64",
        sha256 = "2095675d99107bc61c726864cd396fb9806c93425afe17017470305c2c227896",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Net-1.03-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Net-1.03-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Net-1.03-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Net-1.03-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Net-1.03-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Net-Ping-0__2.76-502.fc40.x86_64",
        sha256 = "4fcb1cc76c8b4fbe58eb2dc82800fb06b7797582d692c4743cb69f5d0b579421",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Net-Ping-2.76-502.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Net-Ping-2.76-502.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Net-Ping-2.76-502.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Net-Ping-2.76-502.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Net-Ping-2.76-502.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Net-SSLeay-0__1.94-3.fc40.x86_64",
        sha256 = "2aba01c37064d99f5c0b7110c418a8f374ffb7c6a6ce142239ed9b7ef429567a",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Net-SSLeay-1.94-3.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Net-SSLeay-1.94-3.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Net-SSLeay-1.94-3.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Net-SSLeay-1.94-3.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Net-SSLeay-1.94-3.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "perl-ODBM_File-0__1.18-506.fc40.x86_64",
        sha256 = "8c48bd8d665645f99a349eadd50b89647439adea37a9a01b95bd3d5449e33289",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-ODBM_File-1.18-506.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-ODBM_File-1.18-506.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-ODBM_File-1.18-506.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-ODBM_File-1.18-506.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-ODBM_File-1.18-506.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "perl-Object-HashBase-0__0.013-1.fc40.x86_64",
        sha256 = "9d44cb7f7a16793be7679c9ec543faae4c893a4887276adf66ef0ad5f0578bbc",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Object-HashBase-0.013-1.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Object-HashBase-0.013-1.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Object-HashBase-0.013-1.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Object-HashBase-0.013-1.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Object-HashBase-0.013-1.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Opcode-0__1.64-506.fc40.x86_64",
        sha256 = "cfc51c4135387790a345b4568bc3d6a8e2e2a488a37a4fd1ce3e34685bade822",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Opcode-1.64-506.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Opcode-1.64-506.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Opcode-1.64-506.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Opcode-1.64-506.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Opcode-1.64-506.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "perl-POSIX-0__2.13-506.fc40.x86_64",
        sha256 = "1225b8dd6ad6ac49a8c749797cfc8d376dc92c895a90a2566124de03e188ed66",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-POSIX-2.13-506.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-POSIX-2.13-506.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-POSIX-2.13-506.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-POSIX-2.13-506.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-POSIX-2.13-506.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "perl-Package-Generator-0__1.106-31.fc40.x86_64",
        sha256 = "bc41d9ad5b7c28ddfaf66a169d9ad5f4452cc7c360b8f8e8d61f659156b458d4",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Package-Generator-1.106-31.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Package-Generator-1.106-31.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Package-Generator-1.106-31.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Package-Generator-1.106-31.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Package-Generator-1.106-31.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Params-Check-1__0.38-502.fc40.x86_64",
        sha256 = "f1fdf697d7276ce45999bde7ad5a54e52c55b48d633965fe33c66994b8e2ebd3",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Params-Check-0.38-502.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Params-Check-0.38-502.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Params-Check-0.38-502.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Params-Check-0.38-502.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Params-Check-0.38-502.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Params-Util-0__1.102-14.fc40.x86_64",
        sha256 = "78dffa230aa3baf893dc36d862ca920ae168c532624fba655ea71e15cb242c7c",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Params-Util-1.102-14.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Params-Util-1.102-14.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Params-Util-1.102-14.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Params-Util-1.102-14.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Params-Util-1.102-14.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "perl-PathTools-0__3.89-502.fc40.x86_64",
        sha256 = "c7790504f73cc2c3227827c1a3a34c018b50dbac0621c7dfcba0c29c64ef15de",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-PathTools-3.89-502.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-PathTools-3.89-502.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-PathTools-3.89-502.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-PathTools-3.89-502.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-PathTools-3.89-502.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "perl-Perl-OSType-0__1.010-503.fc40.x86_64",
        sha256 = "a9b7f734bad66bfe7c82611da67947558bee5d5d400c44f6485211952b74f6fd",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Perl-OSType-1.010-503.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Perl-OSType-1.010-503.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Perl-OSType-1.010-503.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Perl-OSType-1.010-503.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Perl-OSType-1.010-503.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-PerlIO-via-QuotedPrint-0__0.10-502.fc40.x86_64",
        sha256 = "c2eac6a5c9b42fde5593f056fd6c6a952723decae4c24b311964b1272238829a",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-PerlIO-via-QuotedPrint-0.10-502.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-PerlIO-via-QuotedPrint-0.10-502.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-PerlIO-via-QuotedPrint-0.10-502.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-PerlIO-via-QuotedPrint-0.10-502.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-PerlIO-via-QuotedPrint-0.10-502.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Pod-Checker-4__1.77-1.fc40.x86_64",
        sha256 = "01491dd52a63826b44360824a64934c5dfdf03715049c9fd4ff06f60ce00db6f",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Pod-Checker-1.77-1.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Pod-Checker-1.77-1.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Pod-Checker-1.77-1.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Pod-Checker-1.77-1.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Pod-Checker-1.77-1.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Pod-Escapes-1__1.07-503.fc40.x86_64",
        sha256 = "1f96f3aff486ab917ba871eb4875e4f201ae5731b9ec2bd974f11257861fad5c",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Pod-Escapes-1.07-503.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Pod-Escapes-1.07-503.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Pod-Escapes-1.07-503.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Pod-Escapes-1.07-503.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Pod-Escapes-1.07-503.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Pod-Functions-0__1.14-506.fc40.x86_64",
        sha256 = "c9b76089229788021e5a8adda8b5d08a3e463c84436263ae164d64f1f89e0d42",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Pod-Functions-1.14-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Pod-Functions-1.14-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Pod-Functions-1.14-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Pod-Functions-1.14-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Pod-Functions-1.14-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Pod-Html-0__1.34-506.fc40.x86_64",
        sha256 = "f26494f733d23e9658e046bd6025198072041a46f47935c69f09046af9b7ad4b",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Pod-Html-1.34-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Pod-Html-1.34-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Pod-Html-1.34-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Pod-Html-1.34-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Pod-Html-1.34-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Pod-Perldoc-0__3.28.01-503.fc40.x86_64",
        sha256 = "6df25726589106437cf557eae67f3b993e5ff40dad2eeb9150dfc2b23115e6e1",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Pod-Perldoc-3.28.01-503.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Pod-Perldoc-3.28.01-503.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Pod-Perldoc-3.28.01-503.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Pod-Perldoc-3.28.01-503.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Pod-Perldoc-3.28.01-503.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Pod-Simple-1__3.45-6.fc40.x86_64",
        sha256 = "412128c45e763ea21250bae59964120a11f5e29e55e18b3d1f93ab64ab160f6b",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Pod-Simple-3.45-6.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Pod-Simple-3.45-6.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Pod-Simple-3.45-6.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Pod-Simple-3.45-6.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Pod-Simple-3.45-6.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Pod-Usage-4__2.03-504.fc40.x86_64",
        sha256 = "47df4820644fab7febd411bcb3b5dbffdefdcc270951a288500b3fca0b7bb5bf",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/p/perl-Pod-Usage-2.03-504.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/p/perl-Pod-Usage-2.03-504.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/p/perl-Pod-Usage-2.03-504.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/p/perl-Pod-Usage-2.03-504.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/p/perl-Pod-Usage-2.03-504.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Safe-0__2.44-506.fc40.x86_64",
        sha256 = "28836afae302d1b66131bb731187f2a8f2b907e09fa5104b478a9b74911859ce",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Safe-2.44-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Safe-2.44-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Safe-2.44-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Safe-2.44-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Safe-2.44-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Scalar-List-Utils-5__1.63-503.fc40.x86_64",
        sha256 = "9ee3e9fb07acc37ad1aa6ef135ed245622c02118d6c48d5a725f84dd3845e021",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Scalar-List-Utils-1.63-503.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Scalar-List-Utils-1.63-503.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Scalar-List-Utils-1.63-503.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Scalar-List-Utils-1.63-503.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Scalar-List-Utils-1.63-503.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "perl-Search-Dict-0__1.07-506.fc40.x86_64",
        sha256 = "ef996e4c126174c718efce0f3ffc1b043743c6f615627feae7b23100a743720b",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Search-Dict-1.07-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Search-Dict-1.07-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Search-Dict-1.07-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Search-Dict-1.07-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Search-Dict-1.07-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-SelectSaver-0__1.02-506.fc40.x86_64",
        sha256 = "6a45769f7520bc5ef502bdbc2287e1998cc9ca7cd195c845a7c552a2a9b3a650",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-SelectSaver-1.02-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-SelectSaver-1.02-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-SelectSaver-1.02-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-SelectSaver-1.02-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-SelectSaver-1.02-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-SelfLoader-0__1.26-506.fc40.x86_64",
        sha256 = "52e9a5cd271df85c17be959eca0e110ba469a40c3f72b22901bc5345c3c7dc39",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-SelfLoader-1.26-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-SelfLoader-1.26-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-SelfLoader-1.26-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-SelfLoader-1.26-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-SelfLoader-1.26-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Socket-4__2.038-1.fc40.x86_64",
        sha256 = "9ae39e75e08ccae18983d93c9bfbd3c7739975c4fde6265c825d5d257a02757a",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/p/perl-Socket-2.038-1.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/p/perl-Socket-2.038-1.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/p/perl-Socket-2.038-1.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/p/perl-Socket-2.038-1.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/p/perl-Socket-2.038-1.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "perl-Software-License-0__0.104006-1.fc40.x86_64",
        sha256 = "0ce7c97d8327e61877d8639662c03c5223e73209b3fe0ca0ef769085232cd6c8",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Software-License-0.104006-1.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Software-License-0.104006-1.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Software-License-0.104006-1.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Software-License-0.104006-1.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Software-License-0.104006-1.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Storable-1__3.32-502.fc40.x86_64",
        sha256 = "d24094afb476a294b29521ffc39dd73607f17be571157a9abfe6a966fe9dbc9f",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Storable-3.32-502.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Storable-3.32-502.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Storable-3.32-502.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Storable-3.32-502.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Storable-3.32-502.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "perl-Sub-Exporter-0__0.991-3.fc40.x86_64",
        sha256 = "7f0a77a8c39db8498e070f5094d1e232b5c49700cbbdd2fec8bfaaf9c82d0a2e",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Sub-Exporter-0.991-3.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Sub-Exporter-0.991-3.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Sub-Exporter-0.991-3.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Sub-Exporter-0.991-3.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Sub-Exporter-0.991-3.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Sub-Install-0__0.929-5.fc40.x86_64",
        sha256 = "84626d852eda28dc8d15f99e7057f08fccd4bedb18a11913567f07eb6effd0af",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Sub-Install-0.929-5.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Sub-Install-0.929-5.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Sub-Install-0.929-5.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Sub-Install-0.929-5.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Sub-Install-0.929-5.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Symbol-0__1.09-506.fc40.x86_64",
        sha256 = "31541521adffee5c73c30bdafdda4f380cd9c08421336f30c3036c71160674da",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Symbol-1.09-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Symbol-1.09-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Symbol-1.09-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Symbol-1.09-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Symbol-1.09-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Sys-Hostname-0__1.25-506.fc40.x86_64",
        sha256 = "d36294150d8678eb61957fa121b98be306452434df80dd224b29b7f150a17fd7",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Sys-Hostname-1.25-506.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Sys-Hostname-1.25-506.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Sys-Hostname-1.25-506.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Sys-Hostname-1.25-506.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Sys-Hostname-1.25-506.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "perl-Sys-Syslog-0__0.36-504.fc40.x86_64",
        sha256 = "5a04638a95bd69c41f628fe1db39196e4392faefa9ca80fc6066a9735515642f",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Sys-Syslog-0.36-504.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Sys-Syslog-0.36-504.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Sys-Syslog-0.36-504.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Sys-Syslog-0.36-504.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Sys-Syslog-0.36-504.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "perl-Term-ANSIColor-0__5.01-504.fc40.x86_64",
        sha256 = "79dcb1cb584cdb7bbb7c022b63f4b48e14d59c510825c258caca4b126da2ce53",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Term-ANSIColor-5.01-504.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Term-ANSIColor-5.01-504.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Term-ANSIColor-5.01-504.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Term-ANSIColor-5.01-504.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Term-ANSIColor-5.01-504.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Term-Cap-0__1.18-503.fc40.x86_64",
        sha256 = "b6e98d6c3bc2b72a1b5218d6000cbb5ce9787fbdf852cabc742c3e6e3d1f015b",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Term-Cap-1.18-503.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Term-Cap-1.18-503.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Term-Cap-1.18-503.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Term-Cap-1.18-503.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Term-Cap-1.18-503.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Term-Complete-0__1.403-506.fc40.x86_64",
        sha256 = "72d50b52ce34aa7a28ec6f21dcb5138215b4719840ee7d3bc8dd25f4fb17f4e6",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Term-Complete-1.403-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Term-Complete-1.403-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Term-Complete-1.403-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Term-Complete-1.403-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Term-Complete-1.403-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Term-ReadLine-0__1.17-506.fc40.x86_64",
        sha256 = "8f9fcee1df65ec707a56a5db1c6487644498aca37889a758b664cdd8401d29a3",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Term-ReadLine-1.17-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Term-ReadLine-1.17-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Term-ReadLine-1.17-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Term-ReadLine-1.17-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Term-ReadLine-1.17-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Term-Table-0__0.018-3.fc40.x86_64",
        sha256 = "80ff06ef84ca3cfcc99738c5388726b478de9efbcd2d08fc5bf2c015c3c793a9",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Term-Table-0.018-3.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Term-Table-0.018-3.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Term-Table-0.018-3.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Term-Table-0.018-3.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Term-Table-0.018-3.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Test-0__1.31-506.fc40.x86_64",
        sha256 = "87f8f0b1a7dc09546cc114a793f75ff1e7d338c10b73980bf227852c56a98cf6",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Test-1.31-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Test-1.31-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Test-1.31-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Test-1.31-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Test-1.31-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Test-Harness-1__3.50-1.fc40.x86_64",
        sha256 = "67fbd3f9fc39d8eb7380d8ea64ef53bd1ca0ddea82a0d06df2eb841f076af330",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/p/perl-Test-Harness-3.50-1.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/p/perl-Test-Harness-3.50-1.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/p/perl-Test-Harness-3.50-1.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/p/perl-Test-Harness-3.50-1.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/p/perl-Test-Harness-3.50-1.fc40.noarch.rpm",
        ],
    )

    rpm(
        name = "perl-Test-Simple-3__1.302198-3.fc40.x86_64",
        sha256 = "50a1496c9d73779b0fe9ac4042f9d0682939245e20da5e1f8443f5dc76381422",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Test-Simple-1.302198-3.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Test-Simple-1.302198-3.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Test-Simple-1.302198-3.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Test-Simple-1.302198-3.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Test-Simple-1.302198-3.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Text-Abbrev-0__1.02-506.fc40.x86_64",
        sha256 = "49be1af475a856ae7d553328d3545f8b5185c94d8105c6388daaef8f78aa507a",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Text-Abbrev-1.02-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Text-Abbrev-1.02-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Text-Abbrev-1.02-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Text-Abbrev-1.02-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Text-Abbrev-1.02-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Text-Balanced-0__2.06-502.fc40.x86_64",
        sha256 = "7bc6da9b4f02264ba87ee83450af7d42c37cda260fb3f7e6be21bb622534bebc",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Text-Balanced-2.06-502.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Text-Balanced-2.06-502.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Text-Balanced-2.06-502.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Text-Balanced-2.06-502.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Text-Balanced-2.06-502.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Text-Diff-0__1.45-21.fc40.x86_64",
        sha256 = "c177c7c5d468a2d656cab043e8d78459fa6aa1a7a7bc64a89b8d31f4724b535e",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Text-Diff-1.45-21.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Text-Diff-1.45-21.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Text-Diff-1.45-21.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Text-Diff-1.45-21.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Text-Diff-1.45-21.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Text-Glob-0__0.11-23.fc40.x86_64",
        sha256 = "131fa277189c06632c3e240fb0a5fef784026310913859db7d25602463fd912c",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Text-Glob-0.11-23.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Text-Glob-0.11-23.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Text-Glob-0.11-23.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Text-Glob-0.11-23.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Text-Glob-0.11-23.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Text-ParseWords-0__3.31-502.fc40.x86_64",
        sha256 = "6911b5d1d519ba25c008d9da7631e7d4e60e7902bbe3bafb35924c440a47080f",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Text-ParseWords-3.31-502.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Text-ParseWords-3.31-502.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Text-ParseWords-3.31-502.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Text-ParseWords-3.31-502.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Text-ParseWords-3.31-502.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Text-Tabs__plus__Wrap-0__2024.001-1.fc40.x86_64",
        sha256 = "9ca255e239f747f40094f3ba0c81079f009a24a244706128f6bfc077f5d3e97d",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Text-Tabs+Wrap-2024.001-1.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Text-Tabs+Wrap-2024.001-1.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Text-Tabs+Wrap-2024.001-1.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Text-Tabs+Wrap-2024.001-1.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Text-Tabs+Wrap-2024.001-1.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Text-Template-0__1.61-5.fc40.x86_64",
        sha256 = "08fc39fb7a3bf50f473e7a4648db696e16d598b313d5df0c1590bc2849e815d3",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Text-Template-1.61-5.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Text-Template-1.61-5.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Text-Template-1.61-5.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Text-Template-1.61-5.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Text-Template-1.61-5.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Thread-0__3.05-506.fc40.x86_64",
        sha256 = "de7779872fcc65fb45930a98adaa732ea0f9ed914c834cd8a0f982a988b8ca78",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Thread-3.05-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Thread-3.05-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Thread-3.05-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Thread-3.05-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Thread-3.05-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Thread-Queue-0__3.14-503.fc40.x86_64",
        sha256 = "174112fe546cb28ca17d8bc2c1a5ff282a634c3eb9141f8c9bd2dc23e79f5ff7",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Thread-Queue-3.14-503.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Thread-Queue-3.14-503.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Thread-Queue-3.14-503.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Thread-Queue-3.14-503.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Thread-Queue-3.14-503.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Thread-Semaphore-0__2.13-506.fc40.x86_64",
        sha256 = "237434e210b6b3bf96b8f0b3cfb86dcb132f9ec15277813af5a3c2862f7473e0",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Thread-Semaphore-2.13-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Thread-Semaphore-2.13-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Thread-Semaphore-2.13-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Thread-Semaphore-2.13-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Thread-Semaphore-2.13-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Tie-0__4.6-506.fc40.x86_64",
        sha256 = "bcf8d4bd280954fc9d8ca113f5a52a62c188f0faa958d14a865ea8b826526fb0",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Tie-4.6-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Tie-4.6-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Tie-4.6-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Tie-4.6-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Tie-4.6-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Tie-File-0__1.07-506.fc40.x86_64",
        sha256 = "d1d415936b536a49f8ae1236f9954fd7798e84953de36f4528327b999d76bbff",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Tie-File-1.07-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Tie-File-1.07-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Tie-File-1.07-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Tie-File-1.07-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Tie-File-1.07-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Tie-Memoize-0__1.1-506.fc40.x86_64",
        sha256 = "2efa634835b838f39693c334a255ed0bdf9502e0afcff6395edea1d32219d0ab",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Tie-Memoize-1.1-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Tie-Memoize-1.1-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Tie-Memoize-1.1-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Tie-Memoize-1.1-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Tie-Memoize-1.1-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Tie-RefHash-0__1.41-1.fc40.x86_64",
        sha256 = "efd2d0b46057f225e059b73270ee8aab84d97481bd580894f698e5fe25546e0d",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/p/perl-Tie-RefHash-1.41-1.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/p/perl-Tie-RefHash-1.41-1.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/p/perl-Tie-RefHash-1.41-1.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/p/perl-Tie-RefHash-1.41-1.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/p/perl-Tie-RefHash-1.41-1.fc40.noarch.rpm",
        ],
    )

    rpm(
        name = "perl-Time-0__1.03-506.fc40.x86_64",
        sha256 = "7a505448da332cbf40a6f68393d6396686c8fe5dcacc6d09bd5009349b69349a",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Time-1.03-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Time-1.03-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Time-1.03-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Time-1.03-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Time-1.03-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Time-HiRes-4__1.9775-502.fc40.x86_64",
        sha256 = "713a44e1d6379ca6e11849cddf31efb42f6e982a266d3cfc1f1f440f08e3e72f",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Time-HiRes-1.9775-502.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Time-HiRes-1.9775-502.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Time-HiRes-1.9775-502.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Time-HiRes-1.9775-502.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Time-HiRes-1.9775-502.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "perl-Time-Local-2__1.350-5.fc40.x86_64",
        sha256 = "c6cbd7eef0215eceb66866ebc57d8311d220be04857c54ff478dd320f2be146a",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Time-Local-1.350-5.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Time-Local-1.350-5.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Time-Local-1.350-5.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Time-Local-1.350-5.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Time-Local-1.350-5.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Time-Piece-0__1.3401-506.fc40.x86_64",
        sha256 = "11aca663f0fc7d0fd76d83e31d83065ba6c4349fa3195c9d9768c6333010a4bd",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Time-Piece-1.3401-506.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Time-Piece-1.3401-506.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Time-Piece-1.3401-506.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Time-Piece-1.3401-506.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Time-Piece-1.3401-506.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "perl-URI-0__5.28-1.fc40.x86_64",
        sha256 = "667865fb93f3851228eb29e5403759315e6e34b0e319b43857eb1d46d0e002bb",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/p/perl-URI-5.28-1.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/p/perl-URI-5.28-1.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/p/perl-URI-5.28-1.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/p/perl-URI-5.28-1.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/p/perl-URI-5.28-1.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-Unicode-Collate-0__1.31-502.fc40.x86_64",
        sha256 = "72f6e49f256356bbe04e2ce289f1127792b589b851762484fee7cf3a7f19537d",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Unicode-Collate-1.31-502.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Unicode-Collate-1.31-502.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Unicode-Collate-1.31-502.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Unicode-Collate-1.31-502.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Unicode-Collate-1.31-502.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "perl-Unicode-Normalize-0__1.32-502.fc40.x86_64",
        sha256 = "a9e40f3ab2c00e6ec01a1927cd6847ad67bd742998aa46a2aead8dc1f0b24b94",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Unicode-Normalize-1.32-502.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Unicode-Normalize-1.32-502.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Unicode-Normalize-1.32-502.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Unicode-Normalize-1.32-502.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Unicode-Normalize-1.32-502.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "perl-Unicode-UCD-0__0.78-506.fc40.x86_64",
        sha256 = "42ff613815f6a56c3b04967260e7fcf1d4a69f6f70a545a09f98157a091ada3d",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Unicode-UCD-0.78-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Unicode-UCD-0.78-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Unicode-UCD-0.78-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Unicode-UCD-0.78-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-Unicode-UCD-0.78-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-User-pwent-0__1.04-506.fc40.x86_64",
        sha256 = "07a55650bfc424e1ee3376bd3503f12629c31da76c24ee0154590107084b2f25",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-User-pwent-1.04-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-User-pwent-1.04-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-User-pwent-1.04-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-User-pwent-1.04-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-User-pwent-1.04-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-autodie-0__2.37-3.fc40.x86_64",
        sha256 = "2ecd5dd679dfa411279b2bd54b0ac25d37b80a2d2275c6d89813c61fa26325aa",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-autodie-2.37-3.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-autodie-2.37-3.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-autodie-2.37-3.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-autodie-2.37-3.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-autodie-2.37-3.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-autouse-0__1.11-506.fc40.x86_64",
        sha256 = "d4385b5a7ef78179e18f6baa3fb323f0494dc8112aa19d16b6c82d7b27adbf11",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-autouse-1.11-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-autouse-1.11-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-autouse-1.11-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-autouse-1.11-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-autouse-1.11-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-base-0__2.27-506.fc40.x86_64",
        sha256 = "a1325c1dc9d7379bb2b786371a7bbdc2c0d217ae73da040960c05279bd46c1ef",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-base-2.27-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-base-2.27-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-base-2.27-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-base-2.27-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-base-2.27-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-bignum-0__0.67-3.fc40.x86_64",
        sha256 = "66f91abbd717c643c9d3fca83356de79c37b5a05e837ecc04e4f642801c48e75",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-bignum-0.67-3.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-bignum-0.67-3.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-bignum-0.67-3.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-bignum-0.67-3.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-bignum-0.67-3.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-blib-0__1.07-506.fc40.x86_64",
        sha256 = "0134ec75e667d9540b06cf36357b0eb770148669e12087cbdfd4eb81c78e72eb",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-blib-1.07-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-blib-1.07-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-blib-1.07-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-blib-1.07-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-blib-1.07-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-constant-0__1.33-503.fc40.x86_64",
        sha256 = "938e497758f54c450e743427cf97e7b4a57399efdf665cbeb35d9a80c7633632",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-constant-1.33-503.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-constant-1.33-503.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-constant-1.33-503.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-constant-1.33-503.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-constant-1.33-503.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-debugger-0__1.60-506.fc40.x86_64",
        sha256 = "ddd5bb33baa41c6cc8787d80ef6a3572ded99721811ebd6236cd1f71e2ac37a6",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-debugger-1.60-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-debugger-1.60-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-debugger-1.60-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-debugger-1.60-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-debugger-1.60-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-deprecate-0__0.04-506.fc40.x86_64",
        sha256 = "11280daf1ceac8e18988ea2d14f094b15f08508c87b7f9124c874f315b81916c",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-deprecate-0.04-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-deprecate-0.04-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-deprecate-0.04-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-deprecate-0.04-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-deprecate-0.04-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-devel-4__5.38.2-506.fc40.x86_64",
        sha256 = "73f182dfc7095238ab8b6cd38081a35b85cf06f3a9c4787367fca7f93ae6ef23",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-devel-5.38.2-506.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-devel-5.38.2-506.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-devel-5.38.2-506.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-devel-5.38.2-506.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-devel-5.38.2-506.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "perl-diagnostics-0__1.39-506.fc40.x86_64",
        sha256 = "fab5ac7514c649793d4a618a2a8c53743eafa992ae4926e218b8ca2c8f85a578",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-diagnostics-1.39-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-diagnostics-1.39-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-diagnostics-1.39-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-diagnostics-1.39-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-diagnostics-1.39-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-doc-0__5.38.2-506.fc40.x86_64",
        sha256 = "03aa3c80af55c2e7d4b71fce9e148e41656a62a51974cb7ed4b5f6edf276307a",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-doc-5.38.2-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-doc-5.38.2-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-doc-5.38.2-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-doc-5.38.2-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-doc-5.38.2-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-encoding-4__3.00-505.fc40.x86_64",
        sha256 = "96889e6e92bcb5243ed1d4e74612541a491c08e71e9a9acba2a60b904fa6865d",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-encoding-3.00-505.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-encoding-3.00-505.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-encoding-3.00-505.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-encoding-3.00-505.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-encoding-3.00-505.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "perl-encoding-warnings-0__0.14-506.fc40.x86_64",
        sha256 = "52c022f98b3aaa0397c53473eed4e32efe2ca6a378faf190e62dbde5361eeb62",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-encoding-warnings-0.14-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-encoding-warnings-0.14-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-encoding-warnings-0.14-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-encoding-warnings-0.14-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-encoding-warnings-0.14-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-experimental-0__0.032-1.fc40.x86_64",
        sha256 = "f36fbd8d75427451d0aed0e7e15d59c2a7865e91a2b5137678a9ac37e301042d",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/p/perl-experimental-0.032-1.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/p/perl-experimental-0.032-1.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/p/perl-experimental-0.032-1.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/p/perl-experimental-0.032-1.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/p/perl-experimental-0.032-1.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-fields-0__2.27-506.fc40.x86_64",
        sha256 = "97169a62eaebec86a7f3dcb3bc9b2bae440d091db84fcf6847e59b98243c7efa",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-fields-2.27-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-fields-2.27-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-fields-2.27-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-fields-2.27-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-fields-2.27-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-filetest-0__1.03-506.fc40.x86_64",
        sha256 = "69e7107f05bce5b8574de37dddde94b29674af5fb1df6b5ee9b16f0829733e8b",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-filetest-1.03-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-filetest-1.03-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-filetest-1.03-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-filetest-1.03-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-filetest-1.03-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-if-0__0.61.000-506.fc40.x86_64",
        sha256 = "8861b716151717d5545b97cf3c8f7bb6fab5563c97a939e35b1be98553e193d8",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-if-0.61.000-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-if-0.61.000-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-if-0.61.000-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-if-0.61.000-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-if-0.61.000-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-inc-latest-2__0.500-28.fc40.x86_64",
        sha256 = "6b9b606bc79f133bc1e9efaca9855c8d65b4ec1ce0566fc53dc260bd06111b36",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-inc-latest-0.500-28.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-inc-latest-0.500-28.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-inc-latest-0.500-28.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-inc-latest-0.500-28.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-inc-latest-0.500-28.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-interpreter-4__5.38.2-506.fc40.x86_64",
        sha256 = "6a56f5f4f92bff453aca680e3c7791ece6ea5bef689d5db1309f2da157ffc8e8",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-interpreter-5.38.2-506.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-interpreter-5.38.2-506.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-interpreter-5.38.2-506.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-interpreter-5.38.2-506.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-interpreter-5.38.2-506.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "perl-less-0__0.03-506.fc40.x86_64",
        sha256 = "b1802a74fb540e5c58a5682f871b1fe7773ac89cd3149bf3a76cca4d48469dca",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-less-0.03-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-less-0.03-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-less-0.03-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-less-0.03-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-less-0.03-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-lib-0__0.65-506.fc40.x86_64",
        sha256 = "ea3d2f694287cb3d6dfd7c5253524493a26bc1482fd53cc57462c69c72146521",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-lib-0.65-506.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-lib-0.65-506.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-lib-0.65-506.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-lib-0.65-506.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-lib-0.65-506.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "perl-libnet-0__3.15-503.fc40.x86_64",
        sha256 = "a70dcd9f231e55757dce04f454e4cb0109edd3be7a91ab01ac34c43e65398160",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-libnet-3.15-503.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-libnet-3.15-503.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-libnet-3.15-503.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-libnet-3.15-503.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-libnet-3.15-503.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-libnetcfg-4__5.38.2-506.fc40.x86_64",
        sha256 = "89c8e1c30187f7355f61f1b054b4edbf07d99d178ab923741e3d0c0f5aee1a7b",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-libnetcfg-5.38.2-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-libnetcfg-5.38.2-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-libnetcfg-5.38.2-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-libnetcfg-5.38.2-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-libnetcfg-5.38.2-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-libs-4__5.38.2-506.fc40.x86_64",
        sha256 = "5e35bfda17c1a3e8bd1105dc1e77ff207e739b5c9a52990dec63ce420b3e1bda",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-libs-5.38.2-506.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-libs-5.38.2-506.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-libs-5.38.2-506.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-libs-5.38.2-506.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-libs-5.38.2-506.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "perl-local-lib-0__2.000029-7.fc40.x86_64",
        sha256 = "9905a6bbac2773979e82474ea1482a40872070cb83b70a47cbe906b06f2afedd",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-local-lib-2.000029-7.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-local-lib-2.000029-7.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-local-lib-2.000029-7.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-local-lib-2.000029-7.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-local-lib-2.000029-7.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-locale-0__1.10-506.fc40.x86_64",
        sha256 = "9113463465c28d02db79ea499f3cfb4a1e508666ca99c776b2068a964c704c49",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-locale-1.10-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-locale-1.10-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-locale-1.10-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-locale-1.10-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-locale-1.10-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-macros-4__5.38.2-506.fc40.x86_64",
        sha256 = "33103a7b62ab22501d690766d4b53540db3f1a36679321062af491111aa0950f",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-macros-5.38.2-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-macros-5.38.2-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-macros-5.38.2-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-macros-5.38.2-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-macros-5.38.2-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-meta-notation-0__5.38.2-506.fc40.x86_64",
        sha256 = "cea0f11c4d3a085d5dbbe35e688a666813baf06706392f3ef6a7dfed5cf617b6",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-meta-notation-5.38.2-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-meta-notation-5.38.2-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-meta-notation-5.38.2-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-meta-notation-5.38.2-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-meta-notation-5.38.2-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-mro-0__1.28-506.fc40.x86_64",
        sha256 = "08de804c0fe01f433a3c786796fbf84f5ee15a462783271a80b276f384d3d574",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-mro-1.28-506.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-mro-1.28-506.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-mro-1.28-506.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-mro-1.28-506.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-mro-1.28-506.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "perl-open-0__1.13-506.fc40.x86_64",
        sha256 = "6fc81615c330e36696ffb1ba6ed601d379dc63cdcbdbbfdbc5287dbd248a3383",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-open-1.13-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-open-1.13-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-open-1.13-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-open-1.13-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-open-1.13-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-overload-0__1.37-506.fc40.x86_64",
        sha256 = "32659c3ebd0c02df994c7162e929611cdde63180cc3ab7a8c77d3656389f3157",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-overload-1.37-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-overload-1.37-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-overload-1.37-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-overload-1.37-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-overload-1.37-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-overloading-0__0.02-506.fc40.x86_64",
        sha256 = "97c29dfb29715c6dd3b45d0b5c3d6cd14da1a041b20540db7a4b01f41b6d6ffd",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-overloading-0.02-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-overloading-0.02-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-overloading-0.02-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-overloading-0.02-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-overloading-0.02-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-parent-1__0.241-502.fc40.x86_64",
        sha256 = "43701d6fd82fd42e15823efe40ae1304373e4d17da266ba639b4e9dfb78ba5b4",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-parent-0.241-502.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-parent-0.241-502.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-parent-0.241-502.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-parent-0.241-502.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-parent-0.241-502.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-perlfaq-0__5.20240218-1.fc40.x86_64",
        sha256 = "804bbbf920ad110ac88eeda16dc91b29e2bc07a7a4007803a3848f11be10333a",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-perlfaq-5.20240218-1.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-perlfaq-5.20240218-1.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-perlfaq-5.20240218-1.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-perlfaq-5.20240218-1.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-perlfaq-5.20240218-1.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-ph-0__5.38.2-506.fc40.x86_64",
        sha256 = "00f4daf76b60d5f86b9b5621e5d6c973ed4052b1b4cb4cba845e9cf49db17990",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-ph-5.38.2-506.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-ph-5.38.2-506.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-ph-5.38.2-506.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-ph-5.38.2-506.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-ph-5.38.2-506.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "perl-podlators-1__5.01-502.fc40.x86_64",
        sha256 = "b3d9b83fa34d7b5b448dedaf535027b102c620feaead777af09c46727b0986c3",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-podlators-5.01-502.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-podlators-5.01-502.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-podlators-5.01-502.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-podlators-5.01-502.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-podlators-5.01-502.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-sigtrap-0__1.10-506.fc40.x86_64",
        sha256 = "cd9c4a53a2e865c8e43ce8f5cde4412490fcc55e2ed226239b3c04e6b60c4099",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-sigtrap-1.10-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-sigtrap-1.10-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-sigtrap-1.10-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-sigtrap-1.10-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-sigtrap-1.10-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-sort-0__2.05-506.fc40.x86_64",
        sha256 = "f4109bfab8689144a0c68deff309790069c5a5bfbe985c722491160a4d30ec1d",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-sort-2.05-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-sort-2.05-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-sort-2.05-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-sort-2.05-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-sort-2.05-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-srpm-macros-0__1-53.fc40.x86_64",
        sha256 = "076aab9e67fd58346b9c8ac369aaef8d52b1aeff4d2d21c9550931e03c6fee26",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-srpm-macros-1-53.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-srpm-macros-1-53.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-srpm-macros-1-53.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-srpm-macros-1-53.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-srpm-macros-1-53.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-subs-0__1.04-506.fc40.x86_64",
        sha256 = "a6793e171929f779c3cc3919c8f701912284b8e24d38c8b2a41d19c69658daa7",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-subs-1.04-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-subs-1.04-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-subs-1.04-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-subs-1.04-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-subs-1.04-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-threads-1__2.36-503.fc40.x86_64",
        sha256 = "e9c9204219367143d22066988f5d0e9207cbaa58513c1d8eba9d0e57e7e23fe1",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-threads-2.36-503.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-threads-2.36-503.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-threads-2.36-503.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-threads-2.36-503.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-threads-2.36-503.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "perl-threads-shared-0__1.68-502.fc40.x86_64",
        sha256 = "7fd4a14dc11b4c75471aee2f3452692e4109d1715228e1dcebf4e031f43e7039",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-threads-shared-1.68-502.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-threads-shared-1.68-502.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-threads-shared-1.68-502.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-threads-shared-1.68-502.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-threads-shared-1.68-502.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "perl-utils-0__5.38.2-506.fc40.x86_64",
        sha256 = "eeff6ef4a3063c29c9080e38a4321a94723f1e0ae4b803edaef6b37fc0d8e7b6",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-utils-5.38.2-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-utils-5.38.2-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-utils-5.38.2-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-utils-5.38.2-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-utils-5.38.2-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-vars-0__1.05-506.fc40.x86_64",
        sha256 = "1c849b40e23b094c8c7d15b26f2e0839e1c1889f0f5892112616bb71a34c9099",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-vars-1.05-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-vars-1.05-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-vars-1.05-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-vars-1.05-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-vars-1.05-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "perl-version-8__0.99.32-1.fc40.x86_64",
        sha256 = "c4e9f1491315add01b8a0778decee74d84f62e3925f6e1c715d24fcab2bc5e04",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/p/perl-version-0.99.32-1.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/p/perl-version-0.99.32-1.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/p/perl-version-0.99.32-1.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/p/perl-version-0.99.32-1.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/p/perl-version-0.99.32-1.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "perl-vmsish-0__1.04-506.fc40.x86_64",
        sha256 = "c6076c350e5e05ee7192a26d7dfd21b323a51f26d48b78ca7f680998f11a6783",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-vmsish-1.04-506.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-vmsish-1.04-506.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-vmsish-1.04-506.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-vmsish-1.04-506.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/perl-vmsish-1.04-506.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "pixman-0__0.43.4-1.fc40.x86_64",
        sha256 = "ef1ecf553352e1eac512e8a306f4fd8df7572ead32fe307ba3aba679a61c382f",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/p/pixman-0.43.4-1.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/p/pixman-0.43.4-1.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/p/pixman-0.43.4-1.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/p/pixman-0.43.4-1.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/p/pixman-0.43.4-1.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "pkgconf-0__2.1.1-2.fc40.x86_64",
        sha256 = "513cbea187f3e8fab8823da02d133e04de6eea1232e2af5b1b0bf9fcfa27b70f",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/p/pkgconf-2.1.1-2.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/p/pkgconf-2.1.1-2.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/p/pkgconf-2.1.1-2.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/p/pkgconf-2.1.1-2.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/p/pkgconf-2.1.1-2.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "pkgconf-m4-0__2.1.1-2.fc40.x86_64",
        sha256 = "b470bae5560e1d676145e9d53f76136f7c7b02a272d055fe89bd744847b49594",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/p/pkgconf-m4-2.1.1-2.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/p/pkgconf-m4-2.1.1-2.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/p/pkgconf-m4-2.1.1-2.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/p/pkgconf-m4-2.1.1-2.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/p/pkgconf-m4-2.1.1-2.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "pkgconf-pkg-config-0__2.1.1-2.fc40.x86_64",
        sha256 = "6778aa7057eea9243ac150c38924edcfc3c23d85c704fe3be7ca2f6c9eca33c1",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/p/pkgconf-pkg-config-2.1.1-2.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/p/pkgconf-pkg-config-2.1.1-2.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/p/pkgconf-pkg-config-2.1.1-2.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/p/pkgconf-pkg-config-2.1.1-2.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/p/pkgconf-pkg-config-2.1.1-2.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "policycoreutils-0__3.7-3.fc40.x86_64",
        sha256 = "8a0ee0be826338862ecd65d04032b43122cda333ba6bb6891b2ae6aed5208832",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/p/policycoreutils-3.7-3.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/p/policycoreutils-3.7-3.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/p/policycoreutils-3.7-3.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/p/policycoreutils-3.7-3.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/p/policycoreutils-3.7-3.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "policycoreutils-python-utils-0__3.7-3.fc40.x86_64",
        sha256 = "7d3aa818a87d3e97fde7fae85e162e07bbe82a3bb5c842aa7e96957b13b110b5",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/p/policycoreutils-python-utils-3.7-3.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/p/policycoreutils-python-utils-3.7-3.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/p/policycoreutils-python-utils-3.7-3.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/p/policycoreutils-python-utils-3.7-3.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/p/policycoreutils-python-utils-3.7-3.fc40.noarch.rpm",
        ],
    )

    rpm(
        name = "popt-0__1.19-6.fc40.x86_64",
        sha256 = "c03ba1c46e0e2dda36e654941f307aaa0d6574ee5143d6fec6e9af2bdf3252a2",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/popt-1.19-6.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/popt-1.19-6.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/popt-1.19-6.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/popt-1.19-6.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/popt-1.19-6.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "publicsuffix-list-dafsa-0__20240107-3.fc40.x86_64",
        sha256 = "cca50802d4f75306bc37126feb92db79fed44dcdabf76c1556853334995b9d3b",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/publicsuffix-list-dafsa-20240107-3.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/publicsuffix-list-dafsa-20240107-3.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/publicsuffix-list-dafsa-20240107-3.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/publicsuffix-list-dafsa-20240107-3.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/publicsuffix-list-dafsa-20240107-3.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "pyproject-srpm-macros-0__1.15.1-1.fc40.x86_64",
        sha256 = "99ba095342f797d5e75686970f55ecd01cb25f146f4b679e7ddedbd766f017a5",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/p/pyproject-srpm-macros-1.15.1-1.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/p/pyproject-srpm-macros-1.15.1-1.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/p/pyproject-srpm-macros-1.15.1-1.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/p/pyproject-srpm-macros-1.15.1-1.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/p/pyproject-srpm-macros-1.15.1-1.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "python-pip-wheel-0__23.3.2-2.fc40.x86_64",
        sha256 = "7c703b431508f44c5184b5c1df052ed0f49b7439d68aa3597a9a57a5b26bd648",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/p/python-pip-wheel-23.3.2-2.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/p/python-pip-wheel-23.3.2-2.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/p/python-pip-wheel-23.3.2-2.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/p/python-pip-wheel-23.3.2-2.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/p/python-pip-wheel-23.3.2-2.fc40.noarch.rpm",
        ],
    )

    rpm(
        name = "python-srpm-macros-0__3.12-8.fc40.x86_64",
        sha256 = "6ea431da8ae16131fcf943610f0bafa6405eea585d96978e4f02854d7a1437cf",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/p/python-srpm-macros-3.12-8.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/p/python-srpm-macros-3.12-8.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/p/python-srpm-macros-3.12-8.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/p/python-srpm-macros-3.12-8.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/p/python-srpm-macros-3.12-8.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "python-unversioned-command-0__3.12.7-1.fc40.x86_64",
        sha256 = "bcac955e69958e064669ed6e0a394bd9dd2c76e63f558a205ced18a9755012ab",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/p/python-unversioned-command-3.12.7-1.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/p/python-unversioned-command-3.12.7-1.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/p/python-unversioned-command-3.12.7-1.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/p/python-unversioned-command-3.12.7-1.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/p/python-unversioned-command-3.12.7-1.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "python3-0__3.12.7-1.fc40.x86_64",
        sha256 = "6d8342314daafde5c5ec4ec2935e74edb9bea107dc8cd72642e322444f264c7d",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/p/python3-3.12.7-1.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/p/python3-3.12.7-1.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/p/python3-3.12.7-1.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/p/python3-3.12.7-1.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/p/python3-3.12.7-1.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "python3-audit-0__4.0.2-1.fc40.x86_64",
        sha256 = "dd857e78e934557af18d40fd0f78fd1b319e8326171b61bdc2e765347c6c25f0",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/p/python3-audit-4.0.2-1.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/p/python3-audit-4.0.2-1.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/p/python3-audit-4.0.2-1.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/p/python3-audit-4.0.2-1.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/p/python3-audit-4.0.2-1.fc40.x86_64.rpm",
        ],
    )

    rpm(
        name = "python3-distro-0__1.9.0-3.fc40.x86_64",
        sha256 = "00507cbbee67333b446b0ebce7c8aa6395dffd97e22bf79766ecc7088c6c0d71",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/python3-distro-1.9.0-3.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/python3-distro-1.9.0-3.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/python3-distro-1.9.0-3.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/python3-distro-1.9.0-3.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/python3-distro-1.9.0-3.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "python3-libs-0__3.12.7-1.fc40.x86_64",
        sha256 = "839d6dd1d8ac9b55f14b504eca5ac5e66b8330341608f7c9132cb29816116ecb",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/p/python3-libs-3.12.7-1.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/p/python3-libs-3.12.7-1.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/p/python3-libs-3.12.7-1.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/p/python3-libs-3.12.7-1.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/p/python3-libs-3.12.7-1.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "python3-libselinux-0__3.7-5.fc40.x86_64",
        sha256 = "3b2d3a8b1af389a35857c66b55081c5a5cf072671d0de45216794a7cc05d119b",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/p/python3-libselinux-3.7-5.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/p/python3-libselinux-3.7-5.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/p/python3-libselinux-3.7-5.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/p/python3-libselinux-3.7-5.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/p/python3-libselinux-3.7-5.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "python3-libsemanage-0__3.7-2.fc40.x86_64",
        sha256 = "f541054773dcf078f3eca960deeda6b22ab5753e53917fa32d850783f05b8e9a",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/p/python3-libsemanage-3.7-2.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/p/python3-libsemanage-3.7-2.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/p/python3-libsemanage-3.7-2.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/p/python3-libsemanage-3.7-2.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/p/python3-libsemanage-3.7-2.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "python3-policycoreutils-0__3.7-3.fc40.x86_64",
        sha256 = "6f225b0c95c58896d646f92289c944d7d79b8603d35b7c4f1a4f9edcc1d01156",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/p/python3-policycoreutils-3.7-3.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/p/python3-policycoreutils-3.7-3.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/p/python3-policycoreutils-3.7-3.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/p/python3-policycoreutils-3.7-3.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/p/python3-policycoreutils-3.7-3.fc40.noarch.rpm",
        ],
    )

    rpm(
        name = "python3-pyparsing-0__3.1.2-2.fc40.x86_64",
        sha256 = "dda9238b75b7a6bca8393907089a397f139003434bdeeff7d4d350bee1cc7d39",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/python3-pyparsing-3.1.2-2.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/python3-pyparsing-3.1.2-2.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/python3-pyparsing-3.1.2-2.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/python3-pyparsing-3.1.2-2.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/p/python3-pyparsing-3.1.2-2.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "python3-setools-0__4.5.1-2.fc40.x86_64",
        sha256 = "c0f418fe40059909fc20f1c5f9297d1966887efa0facfb5f96435ec4600c737e",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/p/python3-setools-4.5.1-2.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/p/python3-setools-4.5.1-2.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/p/python3-setools-4.5.1-2.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/p/python3-setools-4.5.1-2.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/p/python3-setools-4.5.1-2.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "python3-setuptools-0__69.0.3-4.fc40.x86_64",
        sha256 = "89a75463674f5e878374c7e2bfe094efcbf8bba705d0998f9a68f1cef74f12d5",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/p/python3-setuptools-69.0.3-4.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/p/python3-setuptools-69.0.3-4.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/p/python3-setuptools-69.0.3-4.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/p/python3-setuptools-69.0.3-4.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/p/python3-setuptools-69.0.3-4.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "qemu-common-2__8.2.7-1.fc40.x86_64",
        sha256 = "bc621d7c02a1f55eee16bcbea8a2ab84ee8bbeb5bd146f3f52f38bf503a7ec3d",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/q/qemu-common-8.2.7-1.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/q/qemu-common-8.2.7-1.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/q/qemu-common-8.2.7-1.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/q/qemu-common-8.2.7-1.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/q/qemu-common-8.2.7-1.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "qemu-img-2__8.2.7-1.fc40.x86_64",
        sha256 = "44f58a1d6f13d0fae8aade50c80fb81f86ba8dad92a6c5ef185aa7b6ebe77e1f",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/q/qemu-img-8.2.7-1.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/q/qemu-img-8.2.7-1.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/q/qemu-img-8.2.7-1.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/q/qemu-img-8.2.7-1.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/q/qemu-img-8.2.7-1.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "qemu-system-x86-core-2__8.2.7-1.fc40.x86_64",
        sha256 = "7ec82fc1d15cfaffc69fbb882d22ec655d1eb25188ecc1ed8bad1e99b959e686",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/q/qemu-system-x86-core-8.2.7-1.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/q/qemu-system-x86-core-8.2.7-1.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/q/qemu-system-x86-core-8.2.7-1.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/q/qemu-system-x86-core-8.2.7-1.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/q/qemu-system-x86-core-8.2.7-1.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "qt5-srpm-macros-0__5.15.15-1.fc40.x86_64",
        sha256 = "3964b93f36be9a4570d882c2886939eba4df0a880132945d7deb47b21b854bd5",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/q/qt5-srpm-macros-5.15.15-1.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/q/qt5-srpm-macros-5.15.15-1.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/q/qt5-srpm-macros-5.15.15-1.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/q/qt5-srpm-macros-5.15.15-1.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/q/qt5-srpm-macros-5.15.15-1.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "qt6-srpm-macros-0__6.7.2-2.fc40.x86_64",
        sha256 = "a7a1e173b543c524249f8a7eef986f942c89030c0ee7b77ab95faa35c0f4372c",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/q/qt6-srpm-macros-6.7.2-2.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/q/qt6-srpm-macros-6.7.2-2.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/q/qt6-srpm-macros-6.7.2-2.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/q/qt6-srpm-macros-6.7.2-2.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/q/qt6-srpm-macros-6.7.2-2.fc40.noarch.rpm",
        ],
    )

    rpm(
        name = "readline-0__8.2-8.fc40.x86_64",
        sha256 = "dacd59edbe4744fd9f6823d672e01eff89f871e88537554f16c0a275a17d04e9",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/r/readline-8.2-8.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/r/readline-8.2-8.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/r/readline-8.2-8.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/r/readline-8.2-8.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/r/readline-8.2-8.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "redhat-rpm-config-0__288-1.fc40.x86_64",
        sha256 = "a71f0902957839e18a7f9e13caf4d37a3d53d1c3f5f51a4a57eec80b3edb948d",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/r/redhat-rpm-config-288-1.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/r/redhat-rpm-config-288-1.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/r/redhat-rpm-config-288-1.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/r/redhat-rpm-config-288-1.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/r/redhat-rpm-config-288-1.fc40.noarch.rpm",
        ],
    )

    rpm(
        name = "rpm-0__4.19.1.1-1.fc40.x86_64",
        sha256 = "2fbe0a8f9925ba12b4307fbed8c5c148bab91835f1a3e8797ee08d94d2a0bf83",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/r/rpm-4.19.1.1-1.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/r/rpm-4.19.1.1-1.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/r/rpm-4.19.1.1-1.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/r/rpm-4.19.1.1-1.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/r/rpm-4.19.1.1-1.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "rpm-libs-0__4.19.1.1-1.fc40.x86_64",
        sha256 = "c48c149f4aebfe44d649eea6f7a8eaa229dc8db71ff70b66c7403aa9bd072820",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/r/rpm-libs-4.19.1.1-1.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/r/rpm-libs-4.19.1.1-1.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/r/rpm-libs-4.19.1.1-1.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/r/rpm-libs-4.19.1.1-1.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/r/rpm-libs-4.19.1.1-1.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "rpm-plugin-selinux-0__4.19.1.1-1.fc40.x86_64",
        sha256 = "d400a4e4440bea56566fb1e9582d86d1ac2e07745d37fa6e71f43a8fea05217c",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/r/rpm-plugin-selinux-4.19.1.1-1.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/r/rpm-plugin-selinux-4.19.1.1-1.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/r/rpm-plugin-selinux-4.19.1.1-1.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/r/rpm-plugin-selinux-4.19.1.1-1.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/r/rpm-plugin-selinux-4.19.1.1-1.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "rpm-sequoia-0__1.7.0-1.fc40.x86_64",
        sha256 = "9015e31297a54b708071d347b7877d885a2a97c3b18a89fa31f1481b6406eb06",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/r/rpm-sequoia-1.7.0-1.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/r/rpm-sequoia-1.7.0-1.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/r/rpm-sequoia-1.7.0-1.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/r/rpm-sequoia-1.7.0-1.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/r/rpm-sequoia-1.7.0-1.fc40.x86_64.rpm",
        ],
    )

    rpm(
        name = "rsync-0__3.3.0-1.fc40.x86_64",
        sha256 = "925a9918b5d4157540ab21da866ed992b4e9d3ae4eafa015de1af934c690cb8f",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/r/rsync-3.3.0-1.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/r/rsync-3.3.0-1.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/r/rsync-3.3.0-1.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/r/rsync-3.3.0-1.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/r/rsync-3.3.0-1.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "rust-srpm-macros-0__26.3-1.fc40.x86_64",
        sha256 = "5d0470c00b7b6102f383dd8845e7000377040f0bd79e6947034b03f1b84080ef",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/r/rust-srpm-macros-26.3-1.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/r/rust-srpm-macros-26.3-1.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/r/rust-srpm-macros-26.3-1.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/r/rust-srpm-macros-26.3-1.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/r/rust-srpm-macros-26.3-1.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "seabios-bin-0__1.16.3-2.fc40.x86_64",
        sha256 = "cac97b1c51e1ccbf9489c3b67417e018e887287f60a8520dd931578b5e422bf0",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/s/seabios-bin-1.16.3-2.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/s/seabios-bin-1.16.3-2.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/s/seabios-bin-1.16.3-2.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/s/seabios-bin-1.16.3-2.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/s/seabios-bin-1.16.3-2.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "seavgabios-bin-0__1.16.3-2.fc40.x86_64",
        sha256 = "31d20aaa2f430fca6184317a029c076a7405586929632ae6e044308d946e2f30",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/s/seavgabios-bin-1.16.3-2.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/s/seavgabios-bin-1.16.3-2.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/s/seavgabios-bin-1.16.3-2.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/s/seavgabios-bin-1.16.3-2.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/s/seavgabios-bin-1.16.3-2.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "sed-0__4.9-1.fc40.x86_64",
        sha256 = "6a21b2c132a54fd6d9acb846d0a96289ab739b745cdc4c2b31bdbf6b2434a1a7",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/s/sed-4.9-1.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/s/sed-4.9-1.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/s/sed-4.9-1.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/s/sed-4.9-1.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/s/sed-4.9-1.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "selinux-policy-0__40.28-1.fc40.x86_64",
        sha256 = "696ffbc06ad87337390176888362d7ccbc867d4ecfc6b7040da10bba53d58ba9",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/s/selinux-policy-40.28-1.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/s/selinux-policy-40.28-1.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/s/selinux-policy-40.28-1.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/s/selinux-policy-40.28-1.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/s/selinux-policy-40.28-1.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "selinux-policy-minimum-0__40.28-1.fc40.x86_64",
        sha256 = "635e3f3d4cf06d87af41bec2984cfea7c97859e9840793c5eb18560e00319a25",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/s/selinux-policy-minimum-40.28-1.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/s/selinux-policy-minimum-40.28-1.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/s/selinux-policy-minimum-40.28-1.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/s/selinux-policy-minimum-40.28-1.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/s/selinux-policy-minimum-40.28-1.fc40.noarch.rpm",
        ],
    )

    rpm(
        name = "setup-0__2.14.5-2.fc40.x86_64",
        sha256 = "89862f646cd64e81497f01a8b69ab30ac8968c47afef92a2c333608fdb90ccc1",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/s/setup-2.14.5-2.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/s/setup-2.14.5-2.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/s/setup-2.14.5-2.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/s/setup-2.14.5-2.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/s/setup-2.14.5-2.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "shadow-utils-2__4.15.1-4.fc40.x86_64",
        sha256 = "cfde0d25ecac7e689ee083b330b78df51d346c2b7557c83a189d5df95c4e2c8d",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/s/shadow-utils-4.15.1-4.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/s/shadow-utils-4.15.1-4.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/s/shadow-utils-4.15.1-4.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/s/shadow-utils-4.15.1-4.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/s/shadow-utils-4.15.1-4.fc40.x86_64.rpm",
        ],
    )

    rpm(
        name = "snappy-0__1.1.10-4.fc40.x86_64",
        sha256 = "6cc1d2240e6dcb5e78a7a19418a4b293814139fad0a31a99b83330179a651203",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/s/snappy-1.1.10-4.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/s/snappy-1.1.10-4.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/s/snappy-1.1.10-4.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/s/snappy-1.1.10-4.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/s/snappy-1.1.10-4.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "sqlite-libs-0__3.45.1-2.fc40.x86_64",
        sha256 = "a1e23ae521e93ab19d3df77889a6a418c3432025e4880cfd893e40f7165876a7",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/s/sqlite-libs-3.45.1-2.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/s/sqlite-libs-3.45.1-2.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/s/sqlite-libs-3.45.1-2.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/s/sqlite-libs-3.45.1-2.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/s/sqlite-libs-3.45.1-2.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "systemd-0__255.13-1.fc40.x86_64",
        sha256 = "058089be3d6a5e1ecf0728a82d4c266c0f5bd429625bf1e07dc767fb2b7f5231",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/s/systemd-255.13-1.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/s/systemd-255.13-1.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/s/systemd-255.13-1.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/s/systemd-255.13-1.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/s/systemd-255.13-1.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "systemd-libs-0__255.13-1.fc40.x86_64",
        sha256 = "b23ceced9ce2456a4d7cb10327421f5dfc9c6e18e2046e68f70cb4a7320b4d76",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/s/systemd-libs-255.13-1.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/s/systemd-libs-255.13-1.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/s/systemd-libs-255.13-1.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/s/systemd-libs-255.13-1.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/s/systemd-libs-255.13-1.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "systemd-pam-0__255.13-1.fc40.x86_64",
        sha256 = "680a7a2975b1fbe755774ebf855a7cd777404b81a491c8749d0e8b1d67cfcb73",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/s/systemd-pam-255.13-1.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/s/systemd-pam-255.13-1.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/s/systemd-pam-255.13-1.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/s/systemd-pam-255.13-1.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/s/systemd-pam-255.13-1.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "systemtap-sdt-devel-0__5.2__tilde__pre17250223gd07e4284-1.fc40.x86_64",
        sha256 = "9731680de0bce0c99970a9f3974a005fc6aeaddfa401773ab48051f0a91ca7c9",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/s/systemtap-sdt-devel-5.2~pre17250223gd07e4284-1.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/s/systemtap-sdt-devel-5.2~pre17250223gd07e4284-1.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/s/systemtap-sdt-devel-5.2~pre17250223gd07e4284-1.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/s/systemtap-sdt-devel-5.2~pre17250223gd07e4284-1.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/s/systemtap-sdt-devel-5.2~pre17250223gd07e4284-1.fc40.x86_64.rpm",
        ],
    )

    rpm(
        name = "tar-2__1.35-3.fc40.x86_64",
        sha256 = "65819c502727dc293a71a74b9a5f6b0ba781f12a99c5d5535085f168e5eac56e",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/t/tar-1.35-3.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/t/tar-1.35-3.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/t/tar-1.35-3.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/t/tar-1.35-3.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/t/tar-1.35-3.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "tpm2-tss-0__4.1.3-1.fc40.x86_64",
        sha256 = "c3be8a6d0ea23b1d0bf466b19857b97f7ffde811ad7adec0599161059d84cc74",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/t/tpm2-tss-4.1.3-1.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/t/tpm2-tss-4.1.3-1.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/t/tpm2-tss-4.1.3-1.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/t/tpm2-tss-4.1.3-1.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/t/tpm2-tss-4.1.3-1.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "tzdata-0__2024a-5.fc40.x86_64",
        sha256 = "0bd358e7dfb2bd730b62c7375c8d8f8d9e2470f085ca8dc4ec626dc0332d5687",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/t/tzdata-2024a-5.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/t/tzdata-2024a-5.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/t/tzdata-2024a-5.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/t/tzdata-2024a-5.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/t/tzdata-2024a-5.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "unzip-0__6.0-63.fc40.x86_64",
        sha256 = "8e6642b19621b96ea4811018275f27cf55438c353e50e7c8627e0b30562d5126",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/u/unzip-6.0-63.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/u/unzip-6.0-63.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/u/unzip-6.0-63.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/u/unzip-6.0-63.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/u/unzip-6.0-63.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "util-linux-0__2.40.2-1.fc40.x86_64",
        sha256 = "945aa536bc30050abc1870cef167cb944cf78d6628923476db43201a0054574b",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/u/util-linux-2.40.2-1.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/u/util-linux-2.40.2-1.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/u/util-linux-2.40.2-1.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/u/util-linux-2.40.2-1.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/u/util-linux-2.40.2-1.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "util-linux-core-0__2.40.2-1.fc40.x86_64",
        sha256 = "b1aa4e816c01c08c18924865640f214f717cdfc66837e53a24b8edfb80a86f9d",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/u/util-linux-core-2.40.2-1.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/u/util-linux-core-2.40.2-1.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/u/util-linux-core-2.40.2-1.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/u/util-linux-core-2.40.2-1.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/u/util-linux-core-2.40.2-1.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "xen-libs-0__4.18.3-2.fc40.x86_64",
        sha256 = "7eecf927503f5bd58b6c24e40ce3919efa4d61e198901b9cd73f9c21bef62fee",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/x/xen-libs-4.18.3-2.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/x/xen-libs-4.18.3-2.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/x/xen-libs-4.18.3-2.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/x/xen-libs-4.18.3-2.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/x/xen-libs-4.18.3-2.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "xen-licenses-0__4.18.3-2.fc40.x86_64",
        sha256 = "9efa1793b7139b2991eb48c94694052d523b1a00ca1950a08539d4cb64d1ef75",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/x/xen-licenses-4.18.3-2.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/x/xen-licenses-4.18.3-2.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/x/xen-licenses-4.18.3-2.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/x/xen-licenses-4.18.3-2.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/x/xen-licenses-4.18.3-2.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "xxhash-libs-0__0.8.2-4.fc40.x86_64",
        sha256 = "3e2a33c9d72bf2e4f8a964466ca04f8d36042b82e0e9175a90a510cb18e3df85",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/x/xxhash-libs-0.8.2-4.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/x/xxhash-libs-0.8.2-4.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/x/xxhash-libs-0.8.2-4.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/x/xxhash-libs-0.8.2-4.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/x/xxhash-libs-0.8.2-4.fc40.x86_64.rpm",
        ],
    )

    rpm(
        name = "xz-1__5.4.6-3.fc40.x86_64",
        sha256 = "ee599a1c4d7ee635e54ec137af4dded83f433b9c8a5976f75ecdcd000b5246e3",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/x/xz-5.4.6-3.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/x/xz-5.4.6-3.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/x/xz-5.4.6-3.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/x/xz-5.4.6-3.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/x/xz-5.4.6-3.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "xz-libs-1__5.4.6-3.fc40.x86_64",
        sha256 = "b6ee44b3d7e494b0364f26b7d0b169a8092180af787423cd5e8a47dc0f738a66",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/x/xz-libs-5.4.6-3.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/x/xz-libs-5.4.6-3.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/x/xz-libs-5.4.6-3.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/x/xz-libs-5.4.6-3.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/x/xz-libs-5.4.6-3.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "yajl-0__2.1.0-23.fc40.x86_64",
        sha256 = "9e263e0a9b656178519de20733f3e0950fef494aa056daaa2004b522ba50b952",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/y/yajl-2.1.0-23.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/y/yajl-2.1.0-23.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/y/yajl-2.1.0-23.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/y/yajl-2.1.0-23.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/y/yajl-2.1.0-23.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "zig-srpm-macros-0__1-2.fc40.x86_64",
        sha256 = "3957667c460ee5ed7c46c401db9e1366bd8a22921ed620ffd9a4d7e79298a8f0",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/z/zig-srpm-macros-1-2.fc40.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/z/zig-srpm-macros-1-2.fc40.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/z/zig-srpm-macros-1-2.fc40.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/z/zig-srpm-macros-1-2.fc40.noarch.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/z/zig-srpm-macros-1-2.fc40.noarch.rpm",
        ],
    )
    rpm(
        name = "zip-0__3.0-40.fc40.x86_64",
        sha256 = "feafa5144f815ab92fca16446ec7eea763e116a27e3c5716f7308a314e8138ba",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/z/zip-3.0-40.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/40/Everything/x86_64/os/Packages/z/zip-3.0-40.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/40/Everything/x86_64/os/Packages/z/zip-3.0-40.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/40/Everything/x86_64/os/Packages/z/zip-3.0-40.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/releases/40/Everything/x86_64/os/Packages/z/zip-3.0-40.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "zlib-ng-compat-0__2.1.7-2.fc40.x86_64",
        sha256 = "e50b69054de16d757f5667e3acf2e7439302c91a9c418243467f288dfb79f6ea",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/z/zlib-ng-compat-2.1.7-2.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/z/zlib-ng-compat-2.1.7-2.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/z/zlib-ng-compat-2.1.7-2.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/z/zlib-ng-compat-2.1.7-2.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/z/zlib-ng-compat-2.1.7-2.fc40.x86_64.rpm",
        ],
    )
    rpm(
        name = "zlib-ng-compat-devel-0__2.1.7-2.fc40.x86_64",
        sha256 = "1c959c0ee3c2cd84b8940b48d3cb751fdea7c22ed95fb670bac5e0469dda73ba",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/40/Everything/x86_64/Packages/z/zlib-ng-compat-devel-2.1.7-2.fc40.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/40/Everything/x86_64/Packages/z/zlib-ng-compat-devel-2.1.7-2.fc40.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/40/Everything/x86_64/Packages/z/zlib-ng-compat-devel-2.1.7-2.fc40.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/40/Everything/x86_64/Packages/z/zlib-ng-compat-devel-2.1.7-2.fc40.x86_64.rpm",
            "https://storage.googleapis.com/monogon-infra-public/mirror/fedora/linux/updates/40/Everything/x86_64/Packages/z/zlib-ng-compat-devel-2.1.7-2.fc40.x86_64.rpm",
        ],
    )
