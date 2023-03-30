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
        ],
    )
    
    rpm(
        name = "alternatives-0__1.22-1.fc37.x86_64",
        sha256 = "cf161bb87d597d013444180f4aa26c38e4e85b30f998ad77f2adc25143314055",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/a/alternatives-1.22-1.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/a/alternatives-1.22-1.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/a/alternatives-1.22-1.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/a/alternatives-1.22-1.fc37.x86_64.rpm",
        ],
    )
    
    rpm(
        name = "ansible-srpm-macros-0__1-8.1.fc37.x86_64",
        sha256 = "47eace1b623e365e43f3a4f10bd2c6f7399bc99a8b4376364313a3a8cceb86ae",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/a/ansible-srpm-macros-1-8.1.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/a/ansible-srpm-macros-1-8.1.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/a/ansible-srpm-macros-1-8.1.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/a/ansible-srpm-macros-1-8.1.fc37.noarch.rpm",
        ],
    )
    
    rpm(
        name = "audit-libs-0__3.1-2.fc37.x86_64",
        sha256 = "c58f2c9982f16cc492bba42e7618bbc932a0521b27179e17f6828fcf28c266aa",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/a/audit-libs-3.1-2.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/a/audit-libs-3.1-2.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/a/audit-libs-3.1-2.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/a/audit-libs-3.1-2.fc37.x86_64.rpm",
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
        ],
    )
    
    rpm(
        name = "binutils-0__2.38-25.fc37.x86_64",
        sha256 = "1bd2cad570413a77e4d61198b13d0f3186f45b3e59e51fcb45906adc4485dc94",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/b/binutils-2.38-25.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/b/binutils-2.38-25.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/b/binutils-2.38-25.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/b/binutils-2.38-25.fc37.x86_64.rpm",
        ],
    )
    
    rpm(
        name = "binutils-gold-0__2.38-25.fc37.x86_64",
        sha256 = "879a3745843015f356ebf3147049b78b9350c8f5bd1056933173c625809334bd",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/b/binutils-gold-2.38-25.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/b/binutils-gold-2.38-25.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/b/binutils-gold-2.38-25.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/b/binutils-gold-2.38-25.fc37.x86_64.rpm",
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
        ],
    )
    
    rpm(
        name = "coreutils-single-0__9.1-7.fc37.x86_64",
        sha256 = "414bda840560471cb3d7380923ab00585ee78ca2db4b0d52155e9319a32151bc",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/c/coreutils-single-9.1-7.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/c/coreutils-single-9.1-7.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/c/coreutils-single-9.1-7.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/c/coreutils-single-9.1-7.fc37.x86_64.rpm",
        ],
    )
    
    rpm(
        name = "cpp-0__12.2.1-4.fc37.x86_64",
        sha256 = "ec30f4117248407842024d26b6fa315f6aeef1b58bcf2e5f653f271fcb37c32b",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/c/cpp-12.2.1-4.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/c/cpp-12.2.1-4.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/c/cpp-12.2.1-4.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/c/cpp-12.2.1-4.fc37.x86_64.rpm",
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
        ],
    )
    
    rpm(
        name = "curl-minimal-0__7.85.0-8.fc37.x86_64",
        sha256 = "d331d0c957c9b3c6f4f0bf23959f3654415f55968908864bd70d3cc183821295",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/c/curl-minimal-7.85.0-8.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/c/curl-minimal-7.85.0-8.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/c/curl-minimal-7.85.0-8.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/c/curl-minimal-7.85.0-8.fc37.x86_64.rpm",
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
        ],
    )
    
    rpm(
        name = "daxctl-libs-0__76.1-1.fc37.x86_64",
        sha256 = "d0d4954b1a540de7a3536a5eba0c33fc5dd539c2f2d2dc547b33fa3508a3b7db",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/d/daxctl-libs-76.1-1.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/d/daxctl-libs-76.1-1.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/d/daxctl-libs-76.1-1.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/d/daxctl-libs-76.1-1.fc37.x86_64.rpm",
        ],
    )
    
    rpm(
        name = "dbus-1__1.14.6-1.fc37.x86_64",
        sha256 = "671e7cb382f4cab02530739fd9c57463f6d2649571834e6874c8050abf556e68",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/d/dbus-1.14.6-1.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/d/dbus-1.14.6-1.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/d/dbus-1.14.6-1.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/d/dbus-1.14.6-1.fc37.x86_64.rpm",
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
        ],
    )
    
    rpm(
        name = "dbus-common-1__1.14.6-1.fc37.x86_64",
        sha256 = "2f5d8c77a752f02e4fc98f5ac53ca7ca2811831c8e805907dfce05007be95027",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/d/dbus-common-1.14.6-1.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/d/dbus-common-1.14.6-1.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/d/dbus-common-1.14.6-1.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/d/dbus-common-1.14.6-1.fc37.noarch.rpm",
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
        ],
    )
    
    rpm(
        name = "edk2-ovmf-0__20230301gitf80f052277c8-1.fc37.x86_64",
        sha256 = "49252550260575a19fe48d84981203a3352e89513db43112ac393711a19399ba",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/e/edk2-ovmf-20230301gitf80f052277c8-1.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/e/edk2-ovmf-20230301gitf80f052277c8-1.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/e/edk2-ovmf-20230301gitf80f052277c8-1.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/e/edk2-ovmf-20230301gitf80f052277c8-1.fc37.noarch.rpm",
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
        ],
    )
    
    rpm(
        name = "elfutils-debuginfod-client-0__0.189-1.fc37.x86_64",
        sha256 = "9c9b0cce796ebd712a51290ad4fec8f708e7e8435c52b4318c0cde8be746e53b",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/e/elfutils-debuginfod-client-0.189-1.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/e/elfutils-debuginfod-client-0.189-1.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/e/elfutils-debuginfod-client-0.189-1.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/e/elfutils-debuginfod-client-0.189-1.fc37.x86_64.rpm",
        ],
    )
    
    rpm(
        name = "elfutils-default-yama-scope-0__0.189-1.fc37.x86_64",
        sha256 = "d47043e1562a37dab1674d7fc09b42797cebdbe2cd545008574e85b65a5d011e",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/e/elfutils-default-yama-scope-0.189-1.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/e/elfutils-default-yama-scope-0.189-1.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/e/elfutils-default-yama-scope-0.189-1.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/e/elfutils-default-yama-scope-0.189-1.fc37.noarch.rpm",
        ],
    )
    
    rpm(
        name = "elfutils-libelf-0__0.189-1.fc37.x86_64",
        sha256 = "856a761052deed45559ddd5be420d8f29a6771738c3d3a2eef913a8c5c89c22e",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/e/elfutils-libelf-0.189-1.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/e/elfutils-libelf-0.189-1.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/e/elfutils-libelf-0.189-1.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/e/elfutils-libelf-0.189-1.fc37.x86_64.rpm",
        ],
    )
    
    rpm(
        name = "elfutils-libelf-devel-0__0.189-1.fc37.x86_64",
        sha256 = "8a2ff87057da1630f10afae41fce08dc6b6ce30be59bc025a979a11b73912a74",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/e/elfutils-libelf-devel-0.189-1.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/e/elfutils-libelf-devel-0.189-1.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/e/elfutils-libelf-devel-0.189-1.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/e/elfutils-libelf-devel-0.189-1.fc37.x86_64.rpm",
        ],
    )
    
    rpm(
        name = "elfutils-libs-0__0.189-1.fc37.x86_64",
        sha256 = "cdb6c7b26e4cd92d4c88f8abc76c925d9b509a22062e5c079ddbb8e17453a65a",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/e/elfutils-libs-0.189-1.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/e/elfutils-libs-0.189-1.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/e/elfutils-libs-0.189-1.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/e/elfutils-libs-0.189-1.fc37.x86_64.rpm",
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
        ],
    )
    
    rpm(
        name = "gcc-0__12.2.1-4.fc37.x86_64",
        sha256 = "6fea3e733d6a7b98756a4142382180388afea8f4b2c32c5fea33e53537dee0d7",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/g/gcc-12.2.1-4.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/g/gcc-12.2.1-4.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/g/gcc-12.2.1-4.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/g/gcc-12.2.1-4.fc37.x86_64.rpm",
        ],
    )
    
    rpm(
        name = "gcc-c__plus____plus__-0__12.2.1-4.fc37.x86_64",
        sha256 = "3040cbc9372cbb33eab22915144a8297536b27476fe1ee66a34cbcabfe022d47",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/g/gcc-c++-12.2.1-4.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/g/gcc-c++-12.2.1-4.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/g/gcc-c++-12.2.1-4.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/g/gcc-c++-12.2.1-4.fc37.x86_64.rpm",
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
        ],
    )
    
    rpm(
        name = "glib2-0__2.74.6-1.fc37.x86_64",
        sha256 = "57f55072a259bdfe0fe1bab6f8ae2808bd19858214975885a16727b46213c33f",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/g/glib2-2.74.6-1.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/g/glib2-2.74.6-1.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/g/glib2-2.74.6-1.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/g/glib2-2.74.6-1.fc37.x86_64.rpm",
        ],
    )
    
    rpm(
        name = "glibc-0__2.36-9.fc37.x86_64",
        sha256 = "8c8463cd9f194f03ea1607670399e2fbf068857f566c43dd07d351228c25f187",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/g/glibc-2.36-9.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/g/glibc-2.36-9.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/g/glibc-2.36-9.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/g/glibc-2.36-9.fc37.x86_64.rpm",
        ],
    )
    
    rpm(
        name = "glibc-common-0__2.36-9.fc37.x86_64",
        sha256 = "4237c10e5edacc5d5a9ea88e9fc5fef37249d459b13d4a0715c7836374a8da7a",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/g/glibc-common-2.36-9.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/g/glibc-common-2.36-9.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/g/glibc-common-2.36-9.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/g/glibc-common-2.36-9.fc37.x86_64.rpm",
        ],
    )
    
    rpm(
        name = "glibc-devel-0__2.36-9.fc37.x86_64",
        sha256 = "7c24acc5b5f969a508c7fab06da29ecbbe92a7667d8980ce314a4546f75d2cf8",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/g/glibc-devel-2.36-9.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/g/glibc-devel-2.36-9.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/g/glibc-devel-2.36-9.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/g/glibc-devel-2.36-9.fc37.x86_64.rpm",
        ],
    )
    
    rpm(
        name = "glibc-headers-x86-0__2.36-9.fc37.x86_64",
        sha256 = "6a8fc01c73594de3e70901bb0744f2231d5ad75486514f494143d463701ebcf0",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/g/glibc-headers-x86-2.36-9.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/g/glibc-headers-x86-2.36-9.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/g/glibc-headers-x86-2.36-9.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/g/glibc-headers-x86-2.36-9.fc37.noarch.rpm",
        ],
    )
    
    rpm(
        name = "glibc-langpack-en-0__2.36-9.fc37.x86_64",
        sha256 = "61621e393fc1ccd6b1e34c73de5f6c9f586e493794a53785cbd1531f0486159f",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/g/glibc-langpack-en-2.36-9.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/g/glibc-langpack-en-2.36-9.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/g/glibc-langpack-en-2.36-9.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/g/glibc-langpack-en-2.36-9.fc37.x86_64.rpm",
        ],
    )
    
    rpm(
        name = "glibc-static-0__2.36-9.fc37.x86_64",
        sha256 = "06ba48d3af14ec6a2dd3069b53be6bf91aef1660b500c1e9267d0847eeaf5735",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/g/glibc-static-2.36-9.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/g/glibc-static-2.36-9.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/g/glibc-static-2.36-9.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/g/glibc-static-2.36-9.fc37.x86_64.rpm",
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
        ],
    )
    
    rpm(
        name = "gnutls-0__3.8.0-2.fc37.x86_64",
        sha256 = "91b08de00abe5430f61bc5491ea1f11a23712877da5ab9828865e5c17d4841ee",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/g/gnutls-3.8.0-2.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/g/gnutls-3.8.0-2.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/g/gnutls-3.8.0-2.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/g/gnutls-3.8.0-2.fc37.x86_64.rpm",
        ],
    )
    
    rpm(
        name = "gnutls-dane-0__3.8.0-2.fc37.x86_64",
        sha256 = "1c3e6fa21d05e78d64c3f167a6a20f8a674762cbb5ff10eddb8f934d7058a22e",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/g/gnutls-dane-3.8.0-2.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/g/gnutls-dane-3.8.0-2.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/g/gnutls-dane-3.8.0-2.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/g/gnutls-dane-3.8.0-2.fc37.x86_64.rpm",
        ],
    )
    
    rpm(
        name = "gnutls-utils-0__3.8.0-2.fc37.x86_64",
        sha256 = "3afde268a04428d152408b91541bd61a0e5bb43654df296d198de337029d3b54",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/g/gnutls-utils-3.8.0-2.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/g/gnutls-utils-3.8.0-2.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/g/gnutls-utils-3.8.0-2.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/g/gnutls-utils-3.8.0-2.fc37.x86_64.rpm",
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
        ],
    )
    
    rpm(
        name = "kernel-headers-0__6.2.6-200.fc37.x86_64",
        sha256 = "e5c9a77e4935fd0a4f94374e856ebcad9f1af11dbbccfe0aad6cf13b4369cf98",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/k/kernel-headers-6.2.6-200.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/k/kernel-headers-6.2.6-200.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/k/kernel-headers-6.2.6-200.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/k/kernel-headers-6.2.6-200.fc37.x86_64.rpm",
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
        ],
    )
    
    rpm(
        name = "libcurl-minimal-0__7.85.0-8.fc37.x86_64",
        sha256 = "bb0460b195694c78a58e83ab54268a41cc10f9655ac465d4d0588a5c19a35ab1",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/l/libcurl-minimal-7.85.0-8.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/l/libcurl-minimal-7.85.0-8.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/l/libcurl-minimal-7.85.0-8.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/l/libcurl-minimal-7.85.0-8.fc37.x86_64.rpm",
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
        ],
    )
    
    rpm(
        name = "libeconf-0__0.4.0-4.fc37.x86_64",
        sha256 = "f0cc1addee779f09aade289e3be4e9bd103a274a6bdf11f8331878686f432653",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libeconf-0.4.0-4.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libeconf-0.4.0-4.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libeconf-0.4.0-4.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libeconf-0.4.0-4.fc37.x86_64.rpm",
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
        ],
    )
    
    rpm(
        name = "libgcc-0__12.2.1-4.fc37.x86_64",
        sha256 = "25299b673e7488f538c6d0433ea7fe0ffc8311e41dd7115b5985145e493e4b05",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/l/libgcc-12.2.1-4.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/l/libgcc-12.2.1-4.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/l/libgcc-12.2.1-4.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/l/libgcc-12.2.1-4.fc37.x86_64.rpm",
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
        ],
    )
    
    rpm(
        name = "libgomp-0__12.2.1-4.fc37.x86_64",
        sha256 = "3f2da924fd5168b4f31f56895eb80691778319bf85e408ff02a5ac6714f02f50",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/l/libgomp-12.2.1-4.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/l/libgomp-12.2.1-4.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/l/libgomp-12.2.1-4.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/l/libgomp-12.2.1-4.fc37.x86_64.rpm",
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
        ],
    )
    
    rpm(
        name = "libsemanage-0__3.5-1.fc37.x86_64",
        sha256 = "aeb55e09d224bd6212a8456160cabbcfac61eb2a792a572c01567dba9529c208",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/l/libsemanage-3.5-1.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/l/libsemanage-3.5-1.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/l/libsemanage-3.5-1.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/l/libsemanage-3.5-1.fc37.x86_64.rpm",
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
        ],
    )
    
    rpm(
        name = "libstdc__plus____plus__-0__12.2.1-4.fc37.x86_64",
        sha256 = "ba8009388d86fbb92deff293e04eb57ca9c3b3ba41994932b3e4226533ffb575",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/l/libstdc++-12.2.1-4.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/l/libstdc++-12.2.1-4.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/l/libstdc++-12.2.1-4.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/l/libstdc++-12.2.1-4.fc37.x86_64.rpm",
        ],
    )
    
    rpm(
        name = "libstdc__plus____plus__-devel-0__12.2.1-4.fc37.x86_64",
        sha256 = "0088fcc3a2673acaa17f77ea8ee61ec1523e5a4e19d8bcc33be62d6d35b3f464",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/l/libstdc++-devel-12.2.1-4.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/l/libstdc++-devel-12.2.1-4.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/l/libstdc++-devel-12.2.1-4.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/l/libstdc++-devel-12.2.1-4.fc37.x86_64.rpm",
        ],
    )
    
    rpm(
        name = "libstdc__plus____plus__-static-0__12.2.1-4.fc37.x86_64",
        sha256 = "36eeedc16aaa522f9d5875102a892a8f78c070c876782b73dbdd37d486735b6b",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/l/libstdc++-static-12.2.1-4.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/l/libstdc++-static-12.2.1-4.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/l/libstdc++-static-12.2.1-4.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/l/libstdc++-static-12.2.1-4.fc37.x86_64.rpm",
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
        ],
    )
    
    rpm(
        name = "libtirpc-0__1.3.3-0.fc37.x86_64",
        sha256 = "76dcdfd95452e176f64d6008d114e9415cd8384c5c0d3300fe644c137b6917fa",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libtirpc-1.3.3-0.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libtirpc-1.3.3-0.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libtirpc-1.3.3-0.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/l/libtirpc-1.3.3-0.fc37.x86_64.rpm",
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
        ],
    )
    
    rpm(
        name = "libxcrypt-0__4.4.33-4.fc37.x86_64",
        sha256 = "547b9cffb0211abc4445d159e944f4fb59606b2eddfc14813b8c068859294ba6",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/l/libxcrypt-4.4.33-4.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/l/libxcrypt-4.4.33-4.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/l/libxcrypt-4.4.33-4.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/l/libxcrypt-4.4.33-4.fc37.x86_64.rpm",
        ],
    )
    
    rpm(
        name = "libxcrypt-devel-0__4.4.33-4.fc37.x86_64",
        sha256 = "0dd18ac321ca55e6295c19b0fe0dbd45705673b69e5f2a998620ab08846def23",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/l/libxcrypt-devel-4.4.33-4.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/l/libxcrypt-devel-4.4.33-4.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/l/libxcrypt-devel-4.4.33-4.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/l/libxcrypt-devel-4.4.33-4.fc37.x86_64.rpm",
        ],
    )
    
    rpm(
        name = "libxcrypt-static-0__4.4.33-4.fc37.x86_64",
        sha256 = "58f8943d18d193b727610e17f797e1d4423564b4d2d8e6fa3edf8c19d261b661",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/l/libxcrypt-static-4.4.33-4.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/l/libxcrypt-static-4.4.33-4.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/l/libxcrypt-static-4.4.33-4.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/l/libxcrypt-static-4.4.33-4.fc37.x86_64.rpm",
        ],
    )
    
    rpm(
        name = "libxml2-0__2.10.3-2.fc37.x86_64",
        sha256 = "105e8b221029cc4595682cd837dd80c1124685477efbec280fef2e2bb4974d2d",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/l/libxml2-2.10.3-2.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/l/libxml2-2.10.3-2.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/l/libxml2-2.10.3-2.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/l/libxml2-2.10.3-2.fc37.x86_64.rpm",
        ],
    )
    
    rpm(
        name = "libzstd-0__1.5.4-1.fc37.x86_64",
        sha256 = "d9c9de0b8805782ace29c7fbf5a922dc5d34c3e248f4a13b89a350584045d009",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/l/libzstd-1.5.4-1.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/l/libzstd-1.5.4-1.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/l/libzstd-1.5.4-1.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/l/libzstd-1.5.4-1.fc37.x86_64.rpm",
        ],
    )
    
    rpm(
        name = "libzstd-devel-0__1.5.4-1.fc37.x86_64",
        sha256 = "5df640e0308b8241ccbbf008580a5fb2b19261d8e4b8acb74f9fdd2573250245",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/l/libzstd-devel-1.5.4-1.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/l/libzstd-devel-1.5.4-1.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/l/libzstd-devel-1.5.4-1.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/l/libzstd-devel-1.5.4-1.fc37.x86_64.rpm",
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
        ],
    )
    
    rpm(
        name = "llvm-0__15.0.7-1.fc37.x86_64",
        sha256 = "41ab4712a5bfb19d8c25b65fcf8defa8bcdf6ff3ae84d882a7ed0ca83c6efa1e",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/l/llvm-15.0.7-1.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/l/llvm-15.0.7-1.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/l/llvm-15.0.7-1.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/l/llvm-15.0.7-1.fc37.x86_64.rpm",
        ],
    )
    
    rpm(
        name = "llvm-libs-0__15.0.7-1.fc37.x86_64",
        sha256 = "bb01c1946fccde6933ebe937cef351a8cbbe49921a911f030787396c71d8d77d",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/l/llvm-libs-15.0.7-1.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/l/llvm-libs-15.0.7-1.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/l/llvm-libs-15.0.7-1.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/l/llvm-libs-15.0.7-1.fc37.x86_64.rpm",
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
        ],
    )
    
    rpm(
        name = "ncurses-0__6.3-4.20220501.fc37.x86_64",
        sha256 = "7d90626c613d813fc63a1960985483aabf24ef2ab8b3b8f73cc9d8cac4fa6edd",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/n/ncurses-6.3-4.20220501.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/n/ncurses-6.3-4.20220501.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/n/ncurses-6.3-4.20220501.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/n/ncurses-6.3-4.20220501.fc37.x86_64.rpm",
        ],
    )
    
    rpm(
        name = "ncurses-base-0__6.3-4.20220501.fc37.x86_64",
        sha256 = "000164a9a82458fbb69b3433801dcc0d0e2437e21d7f7d4fd45f63a42a0bc26f",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/n/ncurses-base-6.3-4.20220501.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/n/ncurses-base-6.3-4.20220501.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/n/ncurses-base-6.3-4.20220501.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/n/ncurses-base-6.3-4.20220501.fc37.noarch.rpm",
        ],
    )
    
    rpm(
        name = "ncurses-libs-0__6.3-4.20220501.fc37.x86_64",
        sha256 = "75e51eebcd3fe150b421ec5b1c9a6e918caa5b3c0f243f2b70d445fd434488bb",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/n/ncurses-libs-6.3-4.20220501.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/n/ncurses-libs-6.3-4.20220501.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/n/ncurses-libs-6.3-4.20220501.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/n/ncurses-libs-6.3-4.20220501.fc37.x86_64.rpm",
        ],
    )
    
    rpm(
        name = "ndctl-libs-0__76.1-1.fc37.x86_64",
        sha256 = "8c1c1d79191ba78c7db091a9dc79bc9db1f6bbad8191f8ddbcce015e7183b1e6",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/n/ndctl-libs-76.1-1.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/n/ndctl-libs-76.1-1.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/n/ndctl-libs-76.1-1.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/n/ndctl-libs-76.1-1.fc37.x86_64.rpm",
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
        ],
    )
    
    rpm(
        name = "openssl-devel-1__3.0.8-1.fc37.x86_64",
        sha256 = "28cbab4a2dadfdf33c1510d61f4ef48ef0f33165b22ff9d75233332d7a01df71",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/o/openssl-devel-3.0.8-1.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/o/openssl-devel-3.0.8-1.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/o/openssl-devel-3.0.8-1.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/o/openssl-devel-3.0.8-1.fc37.x86_64.rpm",
        ],
    )
    
    rpm(
        name = "openssl-libs-1__3.0.8-1.fc37.x86_64",
        sha256 = "f250396bc408a880a50a53535e8038d593107594af1d9d348c01aa27a6348dae",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/o/openssl-libs-3.0.8-1.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/o/openssl-libs-3.0.8-1.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/o/openssl-libs-3.0.8-1.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/o/openssl-libs-3.0.8-1.fc37.x86_64.rpm",
        ],
    )
    
    rpm(
        name = "p11-kit-0__0.24.1-3.fc37.x86_64",
        sha256 = "4dad6ac54eb7708cbfc8522d372f2a196cf711e97e279cbddba8cc8b92970dd7",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/p11-kit-0.24.1-3.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/p11-kit-0.24.1-3.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/p11-kit-0.24.1-3.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/p11-kit-0.24.1-3.fc37.x86_64.rpm",
        ],
    )
    
    rpm(
        name = "p11-kit-trust-0__0.24.1-3.fc37.x86_64",
        sha256 = "0fd85eb1ce27615fea745721b18648b4a4585ad4b11a482c1b77fc1785cd5194",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/p11-kit-trust-0.24.1-3.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/p11-kit-trust-0.24.1-3.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/p11-kit-trust-0.24.1-3.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/p11-kit-trust-0.24.1-3.fc37.x86_64.rpm",
        ],
    )
    
    rpm(
        name = "package-notes-srpm-macros-0__0.5-6.fc37.x86_64",
        sha256 = "f565068ef5ce845e1e1e970165a87994409c4748c151ac199fce44f84e19df81",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/package-notes-srpm-macros-0.5-6.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/package-notes-srpm-macros-0.5-6.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/package-notes-srpm-macros-0.5-6.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/package-notes-srpm-macros-0.5-6.fc37.noarch.rpm",
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
        ],
    )
    
    rpm(
        name = "perl-4__5.36.0-492.fc37.x86_64",
        sha256 = "f0aea1adb9b99baaf3c01956c58903603d86f7b63ac8b4a496222de00aa7541d",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-5.36.0-492.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-5.36.0-492.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-5.36.0-492.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-5.36.0-492.fc37.x86_64.rpm",
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
        ],
    )
    
    rpm(
        name = "perl-Attribute-Handlers-0__1.02-492.fc37.x86_64",
        sha256 = "70f26b37bf50d15ca2d702a52db4badabc89ac664c08a0004c30624f1f650c8c",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Attribute-Handlers-1.02-492.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Attribute-Handlers-1.02-492.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Attribute-Handlers-1.02-492.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Attribute-Handlers-1.02-492.fc37.noarch.rpm",
        ],
    )
    
    rpm(
        name = "perl-AutoLoader-0__5.74-492.fc37.x86_64",
        sha256 = "5f7353a0223541047942221e585bd83b23c096a44a4133fe68cb94ac56906d8f",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-AutoLoader-5.74-492.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-AutoLoader-5.74-492.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-AutoLoader-5.74-492.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-AutoLoader-5.74-492.fc37.noarch.rpm",
        ],
    )
    
    rpm(
        name = "perl-AutoSplit-0__5.74-492.fc37.x86_64",
        sha256 = "c4cc3af978917ff71ff2dc186f6e8141d396200449bb38961d594caccd5952b8",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-AutoSplit-5.74-492.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-AutoSplit-5.74-492.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-AutoSplit-5.74-492.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-AutoSplit-5.74-492.fc37.noarch.rpm",
        ],
    )
    
    rpm(
        name = "perl-B-0__1.83-492.fc37.x86_64",
        sha256 = "cbfe6c9df1e503afd59d330b9404232c7b01f3fca5a4a7eb54d5cd1b01ea2768",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-B-1.83-492.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-B-1.83-492.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-B-1.83-492.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-B-1.83-492.fc37.x86_64.rpm",
        ],
    )
    
    rpm(
        name = "perl-Benchmark-0__1.23-492.fc37.x86_64",
        sha256 = "d2525358394d2b28ed483120e8d2a0730857a66d4b7b4e724c4a7757dda09f22",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Benchmark-1.23-492.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Benchmark-1.23-492.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Benchmark-1.23-492.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Benchmark-1.23-492.fc37.noarch.rpm",
        ],
    )
    
    rpm(
        name = "perl-CPAN-0__2.34-4.fc37.x86_64",
        sha256 = "2c460d5a35b0ee8a9ff48dfd1882c120eef8c82a5c24e71987f2fb28ee287a9e",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-CPAN-2.34-4.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-CPAN-2.34-4.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-CPAN-2.34-4.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-CPAN-2.34-4.fc37.noarch.rpm",
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
        ],
    )
    
    rpm(
        name = "perl-Class-Struct-0__0.66-492.fc37.x86_64",
        sha256 = "625c6cc3d5238fd26369e8190ab57e81b14c5c702e45cf191203b042b3e34807",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Class-Struct-0.66-492.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Class-Struct-0.66-492.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Class-Struct-0.66-492.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Class-Struct-0.66-492.fc37.noarch.rpm",
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
        ],
    )
    
    rpm(
        name = "perl-Config-Extensions-0__0.03-492.fc37.x86_64",
        sha256 = "a9febcc42d4d0b8d67504b57c66641ecc2a84c827ab789ef4ac07ee0cfc80721",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Config-Extensions-0.03-492.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Config-Extensions-0.03-492.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Config-Extensions-0.03-492.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Config-Extensions-0.03-492.fc37.noarch.rpm",
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
        ],
    )
    
    rpm(
        name = "perl-DBM_Filter-0__0.06-492.fc37.x86_64",
        sha256 = "af84c245a4ea2eccff2c4fd88e4c414b738a5e371f3864445d77ab43735f9953",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-DBM_Filter-0.06-492.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-DBM_Filter-0.06-492.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-DBM_Filter-0.06-492.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-DBM_Filter-0.06-492.fc37.noarch.rpm",
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
        ],
    )
    
    rpm(
        name = "perl-Devel-Peek-0__1.32-492.fc37.x86_64",
        sha256 = "d6e834937dfb6fba19179e13e816e44ca7d0f34c9b56d3580c358edb00c4629e",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Devel-Peek-1.32-492.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Devel-Peek-1.32-492.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Devel-Peek-1.32-492.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Devel-Peek-1.32-492.fc37.x86_64.rpm",
        ],
    )
    
    rpm(
        name = "perl-Devel-SelfStubber-0__1.06-492.fc37.x86_64",
        sha256 = "900905f342f30341a5f94ae14b465c3483ea7a10888cc7dd0bf3ac7af7368038",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Devel-SelfStubber-1.06-492.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Devel-SelfStubber-1.06-492.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Devel-SelfStubber-1.06-492.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Devel-SelfStubber-1.06-492.fc37.noarch.rpm",
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
        ],
    )
    
    rpm(
        name = "perl-DirHandle-0__1.05-492.fc37.x86_64",
        sha256 = "8c05b3354be6e46fff2cd8fde08195dab10b1fdb0d386f7a9baac115c6e0cf6b",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-DirHandle-1.05-492.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-DirHandle-1.05-492.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-DirHandle-1.05-492.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-DirHandle-1.05-492.fc37.noarch.rpm",
        ],
    )
    
    rpm(
        name = "perl-Dumpvalue-0__2.27-492.fc37.x86_64",
        sha256 = "8e11c03079e7abcea945eab06dce38f0bdfe996a5c8c5cb7c720194358a124cb",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Dumpvalue-2.27-492.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Dumpvalue-2.27-492.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Dumpvalue-2.27-492.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Dumpvalue-2.27-492.fc37.noarch.rpm",
        ],
    )
    
    rpm(
        name = "perl-DynaLoader-0__1.52-492.fc37.x86_64",
        sha256 = "e7373c9bd3e688edc3a68cd34c92999b3d70d8c26ead19bdbd973971d15308a3",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-DynaLoader-1.52-492.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-DynaLoader-1.52-492.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-DynaLoader-1.52-492.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-DynaLoader-1.52-492.fc37.x86_64.rpm",
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
        ],
    )
    
    rpm(
        name = "perl-English-0__1.11-492.fc37.x86_64",
        sha256 = "e2bd78f3eb62cb3704e7d8ddc8457ba509d6938619b8c3eab385a0b93f89dea3",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-English-1.11-492.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-English-1.11-492.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-English-1.11-492.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-English-1.11-492.fc37.noarch.rpm",
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
        ],
    )
    
    rpm(
        name = "perl-Errno-0__1.36-492.fc37.x86_64",
        sha256 = "aaaca92e6353f4cc2d4b81e73efe7b01c3712ed285d6faf940e029a07d853eb4",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Errno-1.36-492.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Errno-1.36-492.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Errno-1.36-492.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Errno-1.36-492.fc37.x86_64.rpm",
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
        ],
    )
    
    rpm(
        name = "perl-ExtUtils-Constant-0__0.25-492.fc37.x86_64",
        sha256 = "87f7630d912268a4c3e2a6d643441087f0b7f05f1bfb8052dddf24a814e86393",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-ExtUtils-Constant-0.25-492.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-ExtUtils-Constant-0.25-492.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-ExtUtils-Constant-0.25-492.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-ExtUtils-Constant-0.25-492.fc37.noarch.rpm",
        ],
    )
    
    rpm(
        name = "perl-ExtUtils-Embed-0__1.35-492.fc37.x86_64",
        sha256 = "6b12c25700d5e71b222f5cbe11649382744852ff94e8f36565e2ddb925beac87",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-ExtUtils-Embed-1.35-492.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-ExtUtils-Embed-1.35-492.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-ExtUtils-Embed-1.35-492.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-ExtUtils-Embed-1.35-492.fc37.noarch.rpm",
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
        ],
    )
    
    rpm(
        name = "perl-ExtUtils-Miniperl-0__1.11-492.fc37.x86_64",
        sha256 = "11de2e98410f4e72b0e6ad8b562bbac3edacb92cff5a32cb0f65aa6d49f32617",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-ExtUtils-Miniperl-1.11-492.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-ExtUtils-Miniperl-1.11-492.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-ExtUtils-Miniperl-1.11-492.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-ExtUtils-Miniperl-1.11-492.fc37.noarch.rpm",
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
        ],
    )
    
    rpm(
        name = "perl-Fcntl-0__1.15-492.fc37.x86_64",
        sha256 = "3e0b22859a921f40ad384d32b212f15ed4dbaecbf41d1077a6a6d1bfc46dc50f",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Fcntl-1.15-492.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Fcntl-1.15-492.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Fcntl-1.15-492.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Fcntl-1.15-492.fc37.x86_64.rpm",
        ],
    )
    
    rpm(
        name = "perl-File-Basename-0__2.85-492.fc37.x86_64",
        sha256 = "aefb2c6be89b24319aeba4ef9e21623f7dbcc478753219ec32f6f07748f6be5c",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-File-Basename-2.85-492.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-File-Basename-2.85-492.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-File-Basename-2.85-492.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-File-Basename-2.85-492.fc37.noarch.rpm",
        ],
    )
    
    rpm(
        name = "perl-File-Compare-0__1.100.700-492.fc37.x86_64",
        sha256 = "11be4ce85ad18e263ac573cf2c7f4e69d525fca1f4127849d1068f93792f38db",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-File-Compare-1.100.700-492.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-File-Compare-1.100.700-492.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-File-Compare-1.100.700-492.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-File-Compare-1.100.700-492.fc37.noarch.rpm",
        ],
    )
    
    rpm(
        name = "perl-File-Copy-0__2.39-492.fc37.x86_64",
        sha256 = "8f15eeb39e9e3fed3aa8cd547e4afdc8db64d2f4c301602dbbe3365da8c47756",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-File-Copy-2.39-492.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-File-Copy-2.39-492.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-File-Copy-2.39-492.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-File-Copy-2.39-492.fc37.noarch.rpm",
        ],
    )
    
    rpm(
        name = "perl-File-DosGlob-0__1.12-492.fc37.x86_64",
        sha256 = "faa65c1601b13d4067774ff600febb2d72bec22eafbcc1d31261c975ed03afcc",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-File-DosGlob-1.12-492.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-File-DosGlob-1.12-492.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-File-DosGlob-1.12-492.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-File-DosGlob-1.12-492.fc37.x86_64.rpm",
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
        ],
    )
    
    rpm(
        name = "perl-File-Find-0__1.40-492.fc37.x86_64",
        sha256 = "8c20c0ca3afea226d1138afdb96a41baa5b47ebe2b818e4ffe6214b13e0ee267",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-File-Find-1.40-492.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-File-Find-1.40-492.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-File-Find-1.40-492.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-File-Find-1.40-492.fc37.noarch.rpm",
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
        ],
    )
    
    rpm(
        name = "perl-File-stat-0__1.12-492.fc37.x86_64",
        sha256 = "dcadef87fb0da0f5dab7f3e051c8423b7c951dfdedd7b18af3be461e05441945",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-File-stat-1.12-492.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-File-stat-1.12-492.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-File-stat-1.12-492.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-File-stat-1.12-492.fc37.noarch.rpm",
        ],
    )
    
    rpm(
        name = "perl-FileCache-0__1.10-492.fc37.x86_64",
        sha256 = "c6dc2770b9e13f1acb082bb62209142ebe9eeb67146ff466ac0d733fe004035d",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-FileCache-1.10-492.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-FileCache-1.10-492.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-FileCache-1.10-492.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-FileCache-1.10-492.fc37.noarch.rpm",
        ],
    )
    
    rpm(
        name = "perl-FileHandle-0__2.03-492.fc37.x86_64",
        sha256 = "7c424594b33c289b2c84b3a38e804629c3f76552874eca63cf6a32cd939c08b9",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-FileHandle-2.03-492.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-FileHandle-2.03-492.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-FileHandle-2.03-492.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-FileHandle-2.03-492.fc37.noarch.rpm",
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
        ],
    )
    
    rpm(
        name = "perl-FindBin-0__1.53-492.fc37.x86_64",
        sha256 = "bb2ac39789b9e6079938b7d0013ee59de21aa10ccb97184a1a72efe5c80855e8",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-FindBin-1.53-492.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-FindBin-1.53-492.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-FindBin-1.53-492.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-FindBin-1.53-492.fc37.noarch.rpm",
        ],
    )
    
    rpm(
        name = "perl-GDBM_File-1__1.23-492.fc37.x86_64",
        sha256 = "f4e3c4c43cf7a4ade2a0cc5717b4622976d5c55bf3df4735fbc03142104eebaa",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-GDBM_File-1.23-492.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-GDBM_File-1.23-492.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-GDBM_File-1.23-492.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-GDBM_File-1.23-492.fc37.x86_64.rpm",
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
        ],
    )
    
    rpm(
        name = "perl-Getopt-Std-0__1.13-492.fc37.x86_64",
        sha256 = "2d8846f2950c6b09ae72e5584291ac7d775faabba897b8db760444ebc402146e",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Getopt-Std-1.13-492.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Getopt-Std-1.13-492.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Getopt-Std-1.13-492.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Getopt-Std-1.13-492.fc37.noarch.rpm",
        ],
    )
    
    rpm(
        name = "perl-HTTP-Tiny-0__0.082-1.fc37.x86_64",
        sha256 = "9e27a04da9f65e27e0eb8bbcad9ec5dfaa986e11731f009700c453db32676455",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-HTTP-Tiny-0.082-1.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-HTTP-Tiny-0.082-1.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-HTTP-Tiny-0.082-1.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-HTTP-Tiny-0.082-1.fc37.noarch.rpm",
        ],
    )
    
    rpm(
        name = "perl-Hash-Util-0__0.28-492.fc37.x86_64",
        sha256 = "d6fe4d6b8179295b6dd28b937aa080ab30e63e7211ddb4e1fdef09b73d71a021",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Hash-Util-0.28-492.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Hash-Util-0.28-492.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Hash-Util-0.28-492.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Hash-Util-0.28-492.fc37.x86_64.rpm",
        ],
    )
    
    rpm(
        name = "perl-Hash-Util-FieldHash-0__1.26-492.fc37.x86_64",
        sha256 = "2c605866cbd2e149850cb8cc2dcbb19a97db145d66ffbc3d16b95357197f8a3d",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Hash-Util-FieldHash-1.26-492.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Hash-Util-FieldHash-1.26-492.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Hash-Util-FieldHash-1.26-492.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Hash-Util-FieldHash-1.26-492.fc37.x86_64.rpm",
        ],
    )
    
    rpm(
        name = "perl-I18N-Collate-0__1.02-492.fc37.x86_64",
        sha256 = "94f63f8cd28d73411c45575f4dc97b55c7198ea1ab2b2072373a28a7edd8d92f",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-I18N-Collate-1.02-492.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-I18N-Collate-1.02-492.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-I18N-Collate-1.02-492.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-I18N-Collate-1.02-492.fc37.noarch.rpm",
        ],
    )
    
    rpm(
        name = "perl-I18N-LangTags-0__0.45-492.fc37.x86_64",
        sha256 = "db2af7e047e9e4d4a0dae1f82206faa61c3661e1a3d6eda574275792415862eb",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-I18N-LangTags-0.45-492.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-I18N-LangTags-0.45-492.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-I18N-LangTags-0.45-492.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-I18N-LangTags-0.45-492.fc37.noarch.rpm",
        ],
    )
    
    rpm(
        name = "perl-I18N-Langinfo-0__0.21-492.fc37.x86_64",
        sha256 = "65ed64aa219dadebb893e2c1f7f555352b0400cdeeb398e3a9cf085c5dace3ca",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-I18N-Langinfo-0.21-492.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-I18N-Langinfo-0.21-492.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-I18N-Langinfo-0.21-492.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-I18N-Langinfo-0.21-492.fc37.x86_64.rpm",
        ],
    )
    
    rpm(
        name = "perl-IO-0__1.50-492.fc37.x86_64",
        sha256 = "86beaf16f888309ca2c55c240a374f1bd3350770973732d90467261fe7d5620e",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-IO-1.50-492.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-IO-1.50-492.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-IO-1.50-492.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-IO-1.50-492.fc37.x86_64.rpm",
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
        ],
    )
    
    rpm(
        name = "perl-IPC-Open3-0__1.22-492.fc37.x86_64",
        sha256 = "7c53e7f6080e0aa53236ee673e2b4b20df87fdf3503620e9eef871c217dcd937",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-IPC-Open3-1.22-492.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-IPC-Open3-1.22-492.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-IPC-Open3-1.22-492.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-IPC-Open3-1.22-492.fc37.noarch.rpm",
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
        ],
    )
    
    rpm(
        name = "perl-Locale-Maketext-Simple-1__0.21-492.fc37.x86_64",
        sha256 = "853378d9d4e6258c48f50b4272a27566ce7cd7dd3438321f79935e4ffc140c54",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Locale-Maketext-Simple-0.21-492.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Locale-Maketext-Simple-0.21-492.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Locale-Maketext-Simple-0.21-492.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Locale-Maketext-Simple-0.21-492.fc37.noarch.rpm",
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
        ],
    )
    
    rpm(
        name = "perl-Math-Complex-0__1.59-492.fc37.x86_64",
        sha256 = "4bca8d6aba65cdf62006c7f239d79bfb40b46b62cc4e17760ea8bc6424e4daeb",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Math-Complex-1.59-492.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Math-Complex-1.59-492.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Math-Complex-1.59-492.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Math-Complex-1.59-492.fc37.noarch.rpm",
        ],
    )
    
    rpm(
        name = "perl-Memoize-0__1.03-492.fc37.x86_64",
        sha256 = "3b6e1f0dded700daf441b85442318f3dbb81b568278904850f7c0a472ba2e47a",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Memoize-1.03-492.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Memoize-1.03-492.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Memoize-1.03-492.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Memoize-1.03-492.fc37.noarch.rpm",
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
        ],
    )
    
    rpm(
        name = "perl-Module-CoreList-1__5.20230220-1.fc37.x86_64",
        sha256 = "95d5258e4c9302250da4bc4c3e8a55717fe83a2f6d566bf4f6dd1fb7c55e2c78",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Module-CoreList-5.20230220-1.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Module-CoreList-5.20230220-1.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Module-CoreList-5.20230220-1.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Module-CoreList-5.20230220-1.fc37.noarch.rpm",
        ],
    )
    
    rpm(
        name = "perl-Module-CoreList-tools-1__5.20230220-1.fc37.x86_64",
        sha256 = "96761f9c6337c2b0f2a569cb6c67c9d1716000ed7a5fcbbe426335a031cf17ae",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Module-CoreList-tools-5.20230220-1.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Module-CoreList-tools-5.20230220-1.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Module-CoreList-tools-5.20230220-1.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/perl-Module-CoreList-tools-5.20230220-1.fc37.noarch.rpm",
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
        ],
    )
    
    rpm(
        name = "perl-Module-Loaded-1__0.08-492.fc37.x86_64",
        sha256 = "8ff08f7fee963eddd622d025df89d41999181636aa7e36b537a27dcfb151fda4",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Module-Loaded-0.08-492.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Module-Loaded-0.08-492.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Module-Loaded-0.08-492.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Module-Loaded-0.08-492.fc37.noarch.rpm",
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
        ],
    )
    
    rpm(
        name = "perl-NDBM_File-0__1.15-492.fc37.x86_64",
        sha256 = "92dacaf3091ed9e0486e2a8eada3ebae25de8aa886875d5962d89900033a7c4b",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-NDBM_File-1.15-492.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-NDBM_File-1.15-492.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-NDBM_File-1.15-492.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-NDBM_File-1.15-492.fc37.x86_64.rpm",
        ],
    )
    
    rpm(
        name = "perl-NEXT-0__0.69-492.fc37.x86_64",
        sha256 = "ca37e599fff4f01c1dfe4f56106a57d00f45b48f1603c92c2af4eb21a424b311",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-NEXT-0.69-492.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-NEXT-0.69-492.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-NEXT-0.69-492.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-NEXT-0.69-492.fc37.noarch.rpm",
        ],
    )
    
    rpm(
        name = "perl-Net-0__1.03-492.fc37.x86_64",
        sha256 = "376e230d0077e245b48fae6461e8e1de5087f2bcd72b5ee7e3a33ac91fa7f828",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Net-1.03-492.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Net-1.03-492.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Net-1.03-492.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Net-1.03-492.fc37.noarch.rpm",
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
        ],
    )
    
    rpm(
        name = "perl-ODBM_File-0__1.17-492.fc37.x86_64",
        sha256 = "31ab3015503a94d0a09d43d7ecd9f0ab82bf961cde1ad652aeabd47e6ed0eb64",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-ODBM_File-1.17-492.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-ODBM_File-1.17-492.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-ODBM_File-1.17-492.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-ODBM_File-1.17-492.fc37.x86_64.rpm",
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
        ],
    )
    
    rpm(
        name = "perl-Opcode-0__1.57-492.fc37.x86_64",
        sha256 = "b49787f95a5bd3e58b7921274e549fef96bd771c4d045b73f6bc4980bef74223",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Opcode-1.57-492.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Opcode-1.57-492.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Opcode-1.57-492.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Opcode-1.57-492.fc37.x86_64.rpm",
        ],
    )
    
    rpm(
        name = "perl-POSIX-0__2.03-492.fc37.x86_64",
        sha256 = "b4b9f1b5e5fd0b533d88d34c28d98115a6a43759ea7be61d6b6d948c47d4e298",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-POSIX-2.03-492.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-POSIX-2.03-492.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-POSIX-2.03-492.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-POSIX-2.03-492.fc37.x86_64.rpm",
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
        ],
    )
    
    rpm(
        name = "perl-Pod-Functions-0__1.14-492.fc37.x86_64",
        sha256 = "21721493960b9ae56cd1f6219a2e5f9e9d890faa7d2b62c7a8d1c2a4d930e816",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Pod-Functions-1.14-492.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Pod-Functions-1.14-492.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Pod-Functions-1.14-492.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Pod-Functions-1.14-492.fc37.noarch.rpm",
        ],
    )
    
    rpm(
        name = "perl-Pod-Html-0__1.33-492.fc37.x86_64",
        sha256 = "ea0ce618eb4ba7c8629329912213413f1f8a3811c7aa770e4501ca81b314251e",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Pod-Html-1.33-492.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Pod-Html-1.33-492.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Pod-Html-1.33-492.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Pod-Html-1.33-492.fc37.noarch.rpm",
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
        ],
    )
    
    rpm(
        name = "perl-Safe-0__2.43-492.fc37.x86_64",
        sha256 = "e65ade165333bcbd9ddec34bd791a756d2b3fecd7c968690b3fe1876381981bf",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Safe-2.43-492.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Safe-2.43-492.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Safe-2.43-492.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Safe-2.43-492.fc37.noarch.rpm",
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
        ],
    )
    
    rpm(
        name = "perl-Search-Dict-0__1.07-492.fc37.x86_64",
        sha256 = "298f5ede6f143d60f5f7dce8e27d6bd09bef70c055b2632a26f18ed8e764fd67",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Search-Dict-1.07-492.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Search-Dict-1.07-492.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Search-Dict-1.07-492.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Search-Dict-1.07-492.fc37.noarch.rpm",
        ],
    )
    
    rpm(
        name = "perl-SelectSaver-0__1.02-492.fc37.x86_64",
        sha256 = "3b54fe377ca9a901ac80d3bcb70518537c383135e10232a6333135f73325a6e2",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-SelectSaver-1.02-492.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-SelectSaver-1.02-492.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-SelectSaver-1.02-492.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-SelectSaver-1.02-492.fc37.noarch.rpm",
        ],
    )
    
    rpm(
        name = "perl-SelfLoader-0__1.26-492.fc37.x86_64",
        sha256 = "87c31681842c06818ebfce3e80c9986c8c0ab230ca9e329c0bc57647fa22b209",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-SelfLoader-1.26-492.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-SelfLoader-1.26-492.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-SelfLoader-1.26-492.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-SelfLoader-1.26-492.fc37.noarch.rpm",
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
        ],
    )
    
    rpm(
        name = "perl-Symbol-0__1.09-492.fc37.x86_64",
        sha256 = "1a52209c5a9bad4c93923470acbeaa8dc4e87a5e03513d01f7bec0147ec9ff7e",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Symbol-1.09-492.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Symbol-1.09-492.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Symbol-1.09-492.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Symbol-1.09-492.fc37.noarch.rpm",
        ],
    )
    
    rpm(
        name = "perl-Sys-Hostname-0__1.24-492.fc37.x86_64",
        sha256 = "7b3ac8ee8319fe22b1ea21200d9e3d6552d4525652d5728fb2007be982900b62",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Sys-Hostname-1.24-492.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Sys-Hostname-1.24-492.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Sys-Hostname-1.24-492.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Sys-Hostname-1.24-492.fc37.x86_64.rpm",
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
        ],
    )
    
    rpm(
        name = "perl-Term-Complete-0__1.403-492.fc37.x86_64",
        sha256 = "6d1819ce387d246cee80dd6cb3248b438585262f8f2bed7da3763a4ec9303afc",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Term-Complete-1.403-492.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Term-Complete-1.403-492.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Term-Complete-1.403-492.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Term-Complete-1.403-492.fc37.noarch.rpm",
        ],
    )
    
    rpm(
        name = "perl-Term-ReadLine-0__1.17-492.fc37.x86_64",
        sha256 = "05153b51d0a9f7f99fa215477fa0d0473f4c76fb06b8ea980d8dc5ac70f6c66f",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Term-ReadLine-1.17-492.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Term-ReadLine-1.17-492.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Term-ReadLine-1.17-492.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Term-ReadLine-1.17-492.fc37.noarch.rpm",
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
        ],
    )
    
    rpm(
        name = "perl-Test-0__1.31-492.fc37.x86_64",
        sha256 = "b7ee6286dd7cad303c1b62c0ba41711a85378fa693d427fe9612686d0b9854d5",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Test-1.31-492.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Test-1.31-492.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Test-1.31-492.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Test-1.31-492.fc37.noarch.rpm",
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
        ],
    )
    
    rpm(
        name = "perl-Text-Abbrev-0__1.02-492.fc37.x86_64",
        sha256 = "1109024f7d75bbac78f4d7933bfcc3876116b7678b9891a6ad649d357a94e88e",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Text-Abbrev-1.02-492.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Text-Abbrev-1.02-492.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Text-Abbrev-1.02-492.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Text-Abbrev-1.02-492.fc37.noarch.rpm",
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
        ],
    )
    
    rpm(
        name = "perl-Text-Tabs__plus__Wrap-0__2021.0814-489.fc37.x86_64",
        sha256 = "2bd5a61328042e939c3360586232a8c5579d4284bfa83d92ec89a578c733c253",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Text-Tabs+Wrap-2021.0814-489.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Text-Tabs+Wrap-2021.0814-489.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Text-Tabs+Wrap-2021.0814-489.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Text-Tabs+Wrap-2021.0814-489.fc37.noarch.rpm",
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
        ],
    )
    
    rpm(
        name = "perl-Thread-0__3.05-492.fc37.x86_64",
        sha256 = "f3cb9d250a3eb509136c268e1ddbc0c4981bfe7c83b59fc2342f1edf5b20787c",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Thread-3.05-492.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Thread-3.05-492.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Thread-3.05-492.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Thread-3.05-492.fc37.noarch.rpm",
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
        ],
    )
    
    rpm(
        name = "perl-Thread-Semaphore-0__2.13-492.fc37.x86_64",
        sha256 = "9ce6ef52ec7509406d610eb5c121cf2013fabab73df0a1ace795607bf5d0866d",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Thread-Semaphore-2.13-492.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Thread-Semaphore-2.13-492.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Thread-Semaphore-2.13-492.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Thread-Semaphore-2.13-492.fc37.noarch.rpm",
        ],
    )
    
    rpm(
        name = "perl-Tie-0__4.6-492.fc37.x86_64",
        sha256 = "68ff45fa18eb140d897557cc00c1224b94aaeb3d1b245ecda5a29bfa1b45fb26",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Tie-4.6-492.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Tie-4.6-492.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Tie-4.6-492.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Tie-4.6-492.fc37.noarch.rpm",
        ],
    )
    
    rpm(
        name = "perl-Tie-File-0__1.06-492.fc37.x86_64",
        sha256 = "407014d52cc8d53cc2c386cf6eefd5b1a031bd3b8045a2157a5e3fa9be999e5d",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Tie-File-1.06-492.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Tie-File-1.06-492.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Tie-File-1.06-492.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Tie-File-1.06-492.fc37.noarch.rpm",
        ],
    )
    
    rpm(
        name = "perl-Tie-Memoize-0__1.1-492.fc37.x86_64",
        sha256 = "8c2ad2228dc85579651b36b9fc7a4f84d637ac00ddfb9cae6814d838f86e4071",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Tie-Memoize-1.1-492.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Tie-Memoize-1.1-492.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Tie-Memoize-1.1-492.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Tie-Memoize-1.1-492.fc37.noarch.rpm",
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
        ],
    )
    
    rpm(
        name = "perl-Time-0__1.03-492.fc37.x86_64",
        sha256 = "829c3c959d332962f59597acc717354e2666240e715f80bb2f0f47b537bcf4a1",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Time-1.03-492.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Time-1.03-492.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Time-1.03-492.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Time-1.03-492.fc37.noarch.rpm",
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
        ],
    )
    
    rpm(
        name = "perl-Time-Piece-0__1.3401-492.fc37.x86_64",
        sha256 = "cc0b5de6e9568ba005f38ccaaddd63370525eb15f402a7a9ee12b61525930bbb",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Time-Piece-1.3401-492.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Time-Piece-1.3401-492.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Time-Piece-1.3401-492.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Time-Piece-1.3401-492.fc37.x86_64.rpm",
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
        ],
    )
    
    rpm(
        name = "perl-Unicode-UCD-0__0.78-492.fc37.x86_64",
        sha256 = "076eaf0aef74477f9e1d021b8725041821cf0a4edc860089650d45ef0dfe56c8",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Unicode-UCD-0.78-492.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Unicode-UCD-0.78-492.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Unicode-UCD-0.78-492.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-Unicode-UCD-0.78-492.fc37.noarch.rpm",
        ],
    )
    
    rpm(
        name = "perl-User-pwent-0__1.03-492.fc37.x86_64",
        sha256 = "e9a1723c8ef3737ded85311d4b0cf82719235b157760e7039efc71572ec2355e",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-User-pwent-1.03-492.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-User-pwent-1.03-492.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-User-pwent-1.03-492.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-User-pwent-1.03-492.fc37.noarch.rpm",
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
        ],
    )
    
    rpm(
        name = "perl-autouse-0__1.11-492.fc37.x86_64",
        sha256 = "57ed31279bf8affcd89439e9c240edb1896a569b0b5525a3dddcf818c2be9d89",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-autouse-1.11-492.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-autouse-1.11-492.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-autouse-1.11-492.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-autouse-1.11-492.fc37.noarch.rpm",
        ],
    )
    
    rpm(
        name = "perl-base-0__2.27-492.fc37.x86_64",
        sha256 = "86ebe8824bb17e449777c3699e5bc0d6a309c8b0a2e3f3a9310ad1719f855be5",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-base-2.27-492.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-base-2.27-492.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-base-2.27-492.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-base-2.27-492.fc37.noarch.rpm",
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
        ],
    )
    
    rpm(
        name = "perl-blib-0__1.07-492.fc37.x86_64",
        sha256 = "7e38d03f3b11771c4c708a121b3a42e9646b140e90299d129ef3bcc3ebba2941",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-blib-1.07-492.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-blib-1.07-492.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-blib-1.07-492.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-blib-1.07-492.fc37.noarch.rpm",
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
        ],
    )
    
    rpm(
        name = "perl-debugger-0__1.60-492.fc37.x86_64",
        sha256 = "848df87be713dc3c94be8da330f3afb740689c45941492fae3b23184f44cce2b",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-debugger-1.60-492.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-debugger-1.60-492.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-debugger-1.60-492.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-debugger-1.60-492.fc37.noarch.rpm",
        ],
    )
    
    rpm(
        name = "perl-deprecate-0__0.04-492.fc37.x86_64",
        sha256 = "6083334765574daf9d243f3fbb5fe9d150080142e979e056e3f5eec1fb77c80a",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-deprecate-0.04-492.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-deprecate-0.04-492.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-deprecate-0.04-492.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-deprecate-0.04-492.fc37.noarch.rpm",
        ],
    )
    
    rpm(
        name = "perl-devel-4__5.36.0-492.fc37.x86_64",
        sha256 = "e438e8e5bb9f82bdc3a55fd100fdf87965ee1000f4f9109efd6a9da4727436e6",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-devel-5.36.0-492.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-devel-5.36.0-492.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-devel-5.36.0-492.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-devel-5.36.0-492.fc37.x86_64.rpm",
        ],
    )
    
    rpm(
        name = "perl-diagnostics-0__1.39-492.fc37.x86_64",
        sha256 = "98828939c0066a5a1f4b8f0408096803da385c79cc436eabe4dc6183a0b93f08",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-diagnostics-1.39-492.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-diagnostics-1.39-492.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-diagnostics-1.39-492.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-diagnostics-1.39-492.fc37.noarch.rpm",
        ],
    )
    
    rpm(
        name = "perl-doc-0__5.36.0-492.fc37.x86_64",
        sha256 = "e4140104bde1caf14794761e8d14eb28bfbf745d5a5b892cba801569e6397dcd",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-doc-5.36.0-492.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-doc-5.36.0-492.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-doc-5.36.0-492.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-doc-5.36.0-492.fc37.noarch.rpm",
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
        ],
    )
    
    rpm(
        name = "perl-encoding-warnings-0__0.13-492.fc37.x86_64",
        sha256 = "d793f4f0a1f4b141b4139d91dece9bdbe07c1d633f473729c849282c2578fd66",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-encoding-warnings-0.13-492.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-encoding-warnings-0.13-492.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-encoding-warnings-0.13-492.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-encoding-warnings-0.13-492.fc37.noarch.rpm",
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
        ],
    )
    
    rpm(
        name = "perl-fields-0__2.27-492.fc37.x86_64",
        sha256 = "aaef9c988a02864aba2a79c3ef0f47528a0078a46813656fde8bba4788500641",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-fields-2.27-492.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-fields-2.27-492.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-fields-2.27-492.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-fields-2.27-492.fc37.noarch.rpm",
        ],
    )
    
    rpm(
        name = "perl-filetest-0__1.03-492.fc37.x86_64",
        sha256 = "f591e958618c514b8c800c3265ff7e4ce29424561e86ac8cea84d093148f6b02",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-filetest-1.03-492.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-filetest-1.03-492.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-filetest-1.03-492.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-filetest-1.03-492.fc37.noarch.rpm",
        ],
    )
    
    rpm(
        name = "perl-if-0__0.61.000-492.fc37.x86_64",
        sha256 = "7d2a698fdc1923e3091359978026b3cbf98f7774c302c8930bf68488ccb83dd2",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-if-0.61.000-492.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-if-0.61.000-492.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-if-0.61.000-492.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-if-0.61.000-492.fc37.noarch.rpm",
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
        ],
    )
    
    rpm(
        name = "perl-interpreter-4__5.36.0-492.fc37.x86_64",
        sha256 = "b6c1c4885bc6ca7b7bfbc95e0ebe7bca41042105fd97a264181534907b6be73e",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-interpreter-5.36.0-492.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-interpreter-5.36.0-492.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-interpreter-5.36.0-492.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-interpreter-5.36.0-492.fc37.x86_64.rpm",
        ],
    )
    
    rpm(
        name = "perl-less-0__0.03-492.fc37.x86_64",
        sha256 = "678367b05cdd9509b2dd27b80eb8256841eb19d1c764407e3c07f415a18a9a9a",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-less-0.03-492.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-less-0.03-492.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-less-0.03-492.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-less-0.03-492.fc37.noarch.rpm",
        ],
    )
    
    rpm(
        name = "perl-lib-0__0.65-492.fc37.x86_64",
        sha256 = "282fc96900724508339cf7af9357848a5885a483db330dc102ea74e5650dfcbd",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-lib-0.65-492.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-lib-0.65-492.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-lib-0.65-492.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-lib-0.65-492.fc37.x86_64.rpm",
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
        ],
    )
    
    rpm(
        name = "perl-libnetcfg-4__5.36.0-492.fc37.x86_64",
        sha256 = "1240d58520aacb5dbeb92b0bd0e69cf4a2aadf9bc61da2041046ca85b8ffb201",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-libnetcfg-5.36.0-492.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-libnetcfg-5.36.0-492.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-libnetcfg-5.36.0-492.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-libnetcfg-5.36.0-492.fc37.noarch.rpm",
        ],
    )
    
    rpm(
        name = "perl-libs-4__5.36.0-492.fc37.x86_64",
        sha256 = "5d41d4194ffc1e78f3816383b83e424984cf81b9c481e76d1453ab37f4298b2a",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-libs-5.36.0-492.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-libs-5.36.0-492.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-libs-5.36.0-492.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-libs-5.36.0-492.fc37.x86_64.rpm",
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
        ],
    )
    
    rpm(
        name = "perl-locale-0__1.10-492.fc37.x86_64",
        sha256 = "cbacb20b0d6288b39742ffabd458c97ed544e9ab3d760b82fde07aa0ffb2a5b0",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-locale-1.10-492.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-locale-1.10-492.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-locale-1.10-492.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-locale-1.10-492.fc37.noarch.rpm",
        ],
    )
    
    rpm(
        name = "perl-macros-4__5.36.0-492.fc37.x86_64",
        sha256 = "dcc1dbc3cc978b9d255d97d3bdfb4f0f72a487547547751f0d21704af9b751a0",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-macros-5.36.0-492.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-macros-5.36.0-492.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-macros-5.36.0-492.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-macros-5.36.0-492.fc37.noarch.rpm",
        ],
    )
    
    rpm(
        name = "perl-meta-notation-0__5.36.0-492.fc37.x86_64",
        sha256 = "15f621ad233470b5eac6d71addb2ffa585993b8a8181d85d1994230bde63baee",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-meta-notation-5.36.0-492.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-meta-notation-5.36.0-492.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-meta-notation-5.36.0-492.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-meta-notation-5.36.0-492.fc37.noarch.rpm",
        ],
    )
    
    rpm(
        name = "perl-mro-0__1.26-492.fc37.x86_64",
        sha256 = "915be0c666707356a59135f0936925aaea35a31df143107188d929163c83051f",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-mro-1.26-492.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-mro-1.26-492.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-mro-1.26-492.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-mro-1.26-492.fc37.x86_64.rpm",
        ],
    )
    
    rpm(
        name = "perl-open-0__1.13-492.fc37.x86_64",
        sha256 = "22ad1d134c02ae5865119f6606007184152e11543ae19ce4ecc5355195a030ff",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-open-1.13-492.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-open-1.13-492.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-open-1.13-492.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-open-1.13-492.fc37.noarch.rpm",
        ],
    )
    
    rpm(
        name = "perl-overload-0__1.35-492.fc37.x86_64",
        sha256 = "1939f3e239071a62d376b00c0a848ddbd500c2c7c1f5321d794d4d7bf52f94fd",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-overload-1.35-492.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-overload-1.35-492.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-overload-1.35-492.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-overload-1.35-492.fc37.noarch.rpm",
        ],
    )
    
    rpm(
        name = "perl-overloading-0__0.02-492.fc37.x86_64",
        sha256 = "089efb5cbe1bc3281e0ae6db4bf5a25e64bfe60952f7598231d567a16092db7d",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-overloading-0.02-492.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-overloading-0.02-492.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-overloading-0.02-492.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-overloading-0.02-492.fc37.noarch.rpm",
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
        ],
    )
    
    rpm(
        name = "perl-ph-0__5.36.0-492.fc37.x86_64",
        sha256 = "57f0a81faf78b0fecc302f03cd439ddaf2aa5d9c84a2c0175419e32d1d5d9723",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-ph-5.36.0-492.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-ph-5.36.0-492.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-ph-5.36.0-492.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-ph-5.36.0-492.fc37.x86_64.rpm",
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
        ],
    )
    
    rpm(
        name = "perl-sigtrap-0__1.10-492.fc37.x86_64",
        sha256 = "2b9f63215ab64e62372608def7c1cd57feab8708b35226dfd1ec613acc25ceaa",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-sigtrap-1.10-492.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-sigtrap-1.10-492.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-sigtrap-1.10-492.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-sigtrap-1.10-492.fc37.noarch.rpm",
        ],
    )
    
    rpm(
        name = "perl-sort-0__2.05-492.fc37.x86_64",
        sha256 = "b607836367fc7835f7d7bc6ffd841ae7a80c50ad29ca3ad2bd40ff6e3751f74a",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-sort-2.05-492.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-sort-2.05-492.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-sort-2.05-492.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-sort-2.05-492.fc37.noarch.rpm",
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
        ],
    )
    
    rpm(
        name = "perl-subs-0__1.04-492.fc37.x86_64",
        sha256 = "c1fc6b645ca97462145c0161e187206c7e216ba3671b141f1b56b4039cf20d20",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-subs-1.04-492.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-subs-1.04-492.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-subs-1.04-492.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-subs-1.04-492.fc37.noarch.rpm",
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
        ],
    )
    
    rpm(
        name = "perl-utils-0__5.36.0-492.fc37.x86_64",
        sha256 = "ac17f83ae81e3f3dede8a5289e98764244f0c81d45711daa4abc546831f2f012",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-utils-5.36.0-492.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-utils-5.36.0-492.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-utils-5.36.0-492.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-utils-5.36.0-492.fc37.noarch.rpm",
        ],
    )
    
    rpm(
        name = "perl-vars-0__1.05-492.fc37.x86_64",
        sha256 = "b32516b678cd965c3e6cf6c6fae8790384ce89bd491f05cb2d5fc02fc225da7f",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-vars-1.05-492.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-vars-1.05-492.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-vars-1.05-492.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-vars-1.05-492.fc37.noarch.rpm",
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
        ],
    )
    
    rpm(
        name = "perl-vmsish-0__1.04-492.fc37.x86_64",
        sha256 = "1087703704aefc77063ef9739b0525e99c90446f74097346780c321d99657be9",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-vmsish-1.04-492.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-vmsish-1.04-492.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-vmsish-1.04-492.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/perl-vmsish-1.04-492.fc37.noarch.rpm",
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
        ],
    )
    
    rpm(
        name = "pyproject-srpm-macros-0__1.6.3-1.fc37.x86_64",
        sha256 = "9587f7185c2ba2ee5186177885fdca7d15bb0d507bf4bf437e33264316df826f",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/pyproject-srpm-macros-1.6.3-1.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/pyproject-srpm-macros-1.6.3-1.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/pyproject-srpm-macros-1.6.3-1.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/pyproject-srpm-macros-1.6.3-1.fc37.noarch.rpm",
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
        ],
    )
    
    rpm(
        name = "python-setuptools-wheel-0__62.6.0-2.fc37.x86_64",
        sha256 = "5a9c2a69949d1bd9293d3fd34719e4d01c8e65d80957d8534ebc23b1deb756c2",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/python-setuptools-wheel-62.6.0-2.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/python-setuptools-wheel-62.6.0-2.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/python-setuptools-wheel-62.6.0-2.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/python-setuptools-wheel-62.6.0-2.fc37.noarch.rpm",
        ],
    )
    
    rpm(
        name = "python-srpm-macros-0__3.11-5.fc37.x86_64",
        sha256 = "176233f538104b82ba5a0afef9feb9cd6f7b8cb1d2f7d6a02fd57342f568cb72",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/python-srpm-macros-3.11-5.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/python-srpm-macros-3.11-5.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/python-srpm-macros-3.11-5.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/python-srpm-macros-3.11-5.fc37.noarch.rpm",
        ],
    )
    
    rpm(
        name = "python-unversioned-command-0__3.11.2-1.fc37.x86_64",
        sha256 = "e99a4700c5a3dc3a671688805ff67ce47f3e503603924c2c0f9c38f482c5cf0e",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/python-unversioned-command-3.11.2-1.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/python-unversioned-command-3.11.2-1.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/python-unversioned-command-3.11.2-1.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/python-unversioned-command-3.11.2-1.fc37.noarch.rpm",
        ],
    )
    
    rpm(
        name = "python3-0__3.11.2-1.fc37.x86_64",
        sha256 = "9eebbf2abbc9791597032fb6136b92e99cadba5fc0b1e673c923b77ff9a5af8c",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/python3-3.11.2-1.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/python3-3.11.2-1.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/python3-3.11.2-1.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/python3-3.11.2-1.fc37.x86_64.rpm",
        ],
    )
    
    rpm(
        name = "python3-audit-0__3.1-2.fc37.x86_64",
        sha256 = "53e518304fe3f9d6b3952b29fe7878961af69178d2e0293dff3e2aacf1abe644",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/python3-audit-3.1-2.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/python3-audit-3.1-2.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/python3-audit-3.1-2.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/python3-audit-3.1-2.fc37.x86_64.rpm",
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
        ],
    )
    
    rpm(
        name = "python3-libs-0__3.11.2-1.fc37.x86_64",
        sha256 = "1681d31085e638e38b2836d0c71bcbf6ddc8b79386824ea403e0886c9fc9f98d",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/python3-libs-3.11.2-1.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/python3-libs-3.11.2-1.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/python3-libs-3.11.2-1.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/python3-libs-3.11.2-1.fc37.x86_64.rpm",
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
        ],
    )
    
    rpm(
        name = "python3-libsemanage-0__3.5-1.fc37.x86_64",
        sha256 = "e7710d9e96771a3becdb771b28d9a2de6bbcb87a3d1a7319e1f4eb9581f242b8",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/python3-libsemanage-3.5-1.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/p/python3-libsemanage-3.5-1.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/p/python3-libsemanage-3.5-1.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/p/python3-libsemanage-3.5-1.fc37.x86_64.rpm",
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
        ],
    )
    
    rpm(
        name = "python3-setuptools-0__62.6.0-2.fc37.x86_64",
        sha256 = "f54b9672f6cdc282263610a619ebef76a69d245bf1588fe767ea136c31c3c93b",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/python3-setuptools-62.6.0-2.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/python3-setuptools-62.6.0-2.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/python3-setuptools-62.6.0-2.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/p/python3-setuptools-62.6.0-2.fc37.noarch.rpm",
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
        ],
    )
    
    rpm(
        name = "qt5-srpm-macros-0__5.15.8-1.fc37.x86_64",
        sha256 = "f090d765f2f1949f5c7b0c02e431aaae6ac1aa652517923e6311373958d3ae35",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/q/qt5-srpm-macros-5.15.8-1.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/q/qt5-srpm-macros-5.15.8-1.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/q/qt5-srpm-macros-5.15.8-1.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/q/qt5-srpm-macros-5.15.8-1.fc37.noarch.rpm",
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
        ],
    )
    
    rpm(
        name = "rpm-0__4.18.0-1.fc37.x86_64",
        sha256 = "7eb9468d77618514bf861da405e2c85b2411efe81577ebc586fd9c25e5ae4194",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/r/rpm-4.18.0-1.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/r/rpm-4.18.0-1.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/r/rpm-4.18.0-1.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/r/rpm-4.18.0-1.fc37.x86_64.rpm",
        ],
    )
    
    rpm(
        name = "rpm-libs-0__4.18.0-1.fc37.x86_64",
        sha256 = "359602208228e24f4d2b4f0ab057ad7ca604ed3f23b0873e7efe395a0c3df25e",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/r/rpm-libs-4.18.0-1.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/r/rpm-libs-4.18.0-1.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/r/rpm-libs-4.18.0-1.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/r/rpm-libs-4.18.0-1.fc37.x86_64.rpm",
        ],
    )
    
    rpm(
        name = "rpm-plugin-selinux-0__4.18.0-1.fc37.x86_64",
        sha256 = "1ab8c75a2f9ee929bbfdb722abc00fb96252e4e76c66f1d292cfe01330f3d56a",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/r/rpm-plugin-selinux-4.18.0-1.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/r/rpm-plugin-selinux-4.18.0-1.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/r/rpm-plugin-selinux-4.18.0-1.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/r/rpm-plugin-selinux-4.18.0-1.fc37.x86_64.rpm",
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
        ],
    )
    
    rpm(
        name = "rust-srpm-macros-0__24-1.fc37.x86_64",
        sha256 = "b902ca5c0f270319854813cc111d5e35eaa2afd69e52ab8a2c852b94be9716d7",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/r/rust-srpm-macros-24-1.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/r/rust-srpm-macros-24-1.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/r/rust-srpm-macros-24-1.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/r/rust-srpm-macros-24-1.fc37.noarch.rpm",
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
        ],
    )
    
    rpm(
        name = "selinux-policy-0__37.19-1.fc37.x86_64",
        sha256 = "8081e5f42dcf1f55cf328ed7b0aca3793ddac5515fc7cacf89f9b826bbda13ba",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/s/selinux-policy-37.19-1.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/s/selinux-policy-37.19-1.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/s/selinux-policy-37.19-1.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/s/selinux-policy-37.19-1.fc37.noarch.rpm",
        ],
    )
    
    rpm(
        name = "selinux-policy-minimum-0__37.19-1.fc37.x86_64",
        sha256 = "4b031956a2e184cf532bd4080e51485afeee3bf475bad904a86df44e6e44e396",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/s/selinux-policy-minimum-37.19-1.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/s/selinux-policy-minimum-37.19-1.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/s/selinux-policy-minimum-37.19-1.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/s/selinux-policy-minimum-37.19-1.fc37.noarch.rpm",
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
        ],
    )
    
    rpm(
        name = "systemd-0__251.13-6.fc37.x86_64",
        sha256 = "d1191eb3a0149638e395591d2004cd6a5d852e5712ab06c4beb7cfd77e2e2488",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/s/systemd-251.13-6.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/s/systemd-251.13-6.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/s/systemd-251.13-6.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/s/systemd-251.13-6.fc37.x86_64.rpm",
        ],
    )
    
    rpm(
        name = "systemd-libs-0__251.13-6.fc37.x86_64",
        sha256 = "20aa751dfa2c65cf5a3cad75867d863ed562694c8d30a317814d3c6e4a0a3e6a",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/s/systemd-libs-251.13-6.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/s/systemd-libs-251.13-6.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/s/systemd-libs-251.13-6.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/s/systemd-libs-251.13-6.fc37.x86_64.rpm",
        ],
    )
    
    rpm(
        name = "systemd-pam-0__251.13-6.fc37.x86_64",
        sha256 = "788dc48f363aaaf40e004475278db7b6408cbe21e20663ca4ed9540b5070ff2a",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/s/systemd-pam-251.13-6.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/s/systemd-pam-251.13-6.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/s/systemd-pam-251.13-6.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/s/systemd-pam-251.13-6.fc37.x86_64.rpm",
        ],
    )
    
    rpm(
        name = "systemtap-sdt-devel-0__4.8__tilde__pre16594741g5bdc37b9-1.fc37.x86_64",
        sha256 = "a7a93ca017fb034c312cfe09b1dabae1e984c9bf20e346799b4a5abc1e1ad260",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/s/systemtap-sdt-devel-4.8~pre16594741g5bdc37b9-1.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/s/systemtap-sdt-devel-4.8~pre16594741g5bdc37b9-1.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/s/systemtap-sdt-devel-4.8~pre16594741g5bdc37b9-1.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/s/systemtap-sdt-devel-4.8~pre16594741g5bdc37b9-1.fc37.x86_64.rpm",
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
        ],
    )
    
    rpm(
        name = "tzdata-0__2022g-1.fc37.x86_64",
        sha256 = "7ff35c66b3478103fbf3941e933e25f60e41f2b0bfd07d43666b40721211c3bb",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/t/tzdata-2022g-1.fc37.noarch.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/t/tzdata-2022g-1.fc37.noarch.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/t/tzdata-2022g-1.fc37.noarch.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/t/tzdata-2022g-1.fc37.noarch.rpm",
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
        ],
    )
    
    rpm(
        name = "xen-libs-0__4.16.3-4.fc37.x86_64",
        sha256 = "990f530caeb686c8cf6340df4ccec898d6b5e524096736f981161dba680701d8",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/x/xen-libs-4.16.3-4.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/x/xen-libs-4.16.3-4.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/x/xen-libs-4.16.3-4.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/x/xen-libs-4.16.3-4.fc37.x86_64.rpm",
        ],
    )
    
    rpm(
        name = "xen-licenses-0__4.16.3-4.fc37.x86_64",
        sha256 = "1813e36375f344d061474204699a9b9d4f2e6632c6880c5df82230d053dfd59d",
        urls = [
            "https://ftp.fau.de/fedora/linux/updates/37/Everything/x86_64/Packages/x/xen-licenses-4.16.3-4.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/updates/37/Everything/x86_64/Packages/x/xen-licenses-4.16.3-4.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/updates/37/Everything/x86_64/Packages/x/xen-licenses-4.16.3-4.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/updates/37/Everything/x86_64/Packages/x/xen-licenses-4.16.3-4.fc37.x86_64.rpm",
        ],
    )
    
    rpm(
        name = "xxhash-libs-0__0.8.1-3.fc37.x86_64",
        sha256 = "f23dc45a9d083793cce0c688700a0ae92731b140736aa4d014fb371c5ebd2d7e",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/x/xxhash-libs-0.8.1-3.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/x/xxhash-libs-0.8.1-3.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/x/xxhash-libs-0.8.1-3.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/x/xxhash-libs-0.8.1-3.fc37.x86_64.rpm",
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
        ],
    )
    
    rpm(
        name = "yajl-0__2.1.0-19.fc37.x86_64",
        sha256 = "b0ca9c6ed5935cde0094694127c13b99a441207eb084f44fb3aa093669c9957c",
        urls = [
            "https://ftp.fau.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/y/yajl-2.1.0-19.fc37.x86_64.rpm",
            "https://ftp.halifax.rwth-aachen.de/fedora/linux/releases/37/Everything/x86_64/os/Packages/y/yajl-2.1.0-19.fc37.x86_64.rpm",
            "https://mirror.23m.com/fedora/linux/releases/37/Everything/x86_64/os/Packages/y/yajl-2.1.0-19.fc37.x86_64.rpm",
            "https://ftp.plusline.net/fedora/linux/releases/37/Everything/x86_64/os/Packages/y/yajl-2.1.0-19.fc37.x86_64.rpm",
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
        ],
    )
