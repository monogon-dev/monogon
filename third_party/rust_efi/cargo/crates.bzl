"""
@generated
cargo-raze generated Bazel file.

DO NOT EDIT! Replaced on runs of cargo-raze
"""

load("@bazel_tools//tools/build_defs/repo:git.bzl", "new_git_repository")  # buildifier: disable=load
load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")  # buildifier: disable=load
load("@bazel_tools//tools/build_defs/repo:utils.bzl", "maybe")  # buildifier: disable=load

def rsefi_fetch_remote_crates():
    """This function defines a collection of repos and should be called in a WORKSPACE file"""
    maybe(
        http_archive,
        name = "rsefi__anyhow__1_0_75",
        url = "https://crates.io/api/v1/crates/anyhow/1.0.75/download",
        type = "tar.gz",
        sha256 = "a4668cab20f66d8d020e1fbc0ebe47217433c1b6c8f2040faf858554e394ace6",
        strip_prefix = "anyhow-1.0.75",
        build_file = Label("//third_party/rust_efi/cargo/remote:BUILD.anyhow-1.0.75.bazel"),
    )

    maybe(
        http_archive,
        name = "rsefi__bit_field__0_10_2",
        url = "https://crates.io/api/v1/crates/bit_field/0.10.2/download",
        type = "tar.gz",
        sha256 = "dc827186963e592360843fb5ba4b973e145841266c1357f7180c43526f2e5b61",
        strip_prefix = "bit_field-0.10.2",
        build_file = Label("//third_party/rust_efi/cargo/remote:BUILD.bit_field-0.10.2.bazel"),
    )

    maybe(
        http_archive,
        name = "rsefi__bitflags__2_4_0",
        url = "https://crates.io/api/v1/crates/bitflags/2.4.0/download",
        type = "tar.gz",
        sha256 = "b4682ae6287fcf752ecaabbfcc7b6f9b72aa33933dc23a554d853aea8eea8635",
        strip_prefix = "bitflags-2.4.0",
        build_file = Label("//third_party/rust_efi/cargo/remote:BUILD.bitflags-2.4.0.bazel"),
    )

    maybe(
        http_archive,
        name = "rsefi__bytes__1_5_0",
        url = "https://crates.io/api/v1/crates/bytes/1.5.0/download",
        type = "tar.gz",
        sha256 = "a2bd12c1caf447e69cd4528f47f94d203fd2582878ecb9e9465484c4148a8223",
        strip_prefix = "bytes-1.5.0",
        build_file = Label("//third_party/rust_efi/cargo/remote:BUILD.bytes-1.5.0.bazel"),
    )

    maybe(
        http_archive,
        name = "rsefi__cfg_if__1_0_0",
        url = "https://crates.io/api/v1/crates/cfg-if/1.0.0/download",
        type = "tar.gz",
        sha256 = "baf1de4339761588bc0619e3cbc0120ee582ebb74b53b4efbf79117bd2da40fd",
        strip_prefix = "cfg-if-1.0.0",
        build_file = Label("//third_party/rust_efi/cargo/remote:BUILD.cfg-if-1.0.0.bazel"),
    )

    maybe(
        http_archive,
        name = "rsefi__either__1_9_0",
        url = "https://crates.io/api/v1/crates/either/1.9.0/download",
        type = "tar.gz",
        sha256 = "a26ae43d7bcc3b814de94796a5e736d4029efb0ee900c12e2d54c993ad1a1e07",
        strip_prefix = "either-1.9.0",
        build_file = Label("//third_party/rust_efi/cargo/remote:BUILD.either-1.9.0.bazel"),
    )

    maybe(
        http_archive,
        name = "rsefi__itertools__0_11_0",
        url = "https://crates.io/api/v1/crates/itertools/0.11.0/download",
        type = "tar.gz",
        sha256 = "b1c173a5686ce8bfa551b3563d0c2170bf24ca44da99c7ca4bfdab5418c3fe57",
        strip_prefix = "itertools-0.11.0",
        build_file = Label("//third_party/rust_efi/cargo/remote:BUILD.itertools-0.11.0.bazel"),
    )

    maybe(
        http_archive,
        name = "rsefi__log__0_4_20",
        url = "https://crates.io/api/v1/crates/log/0.4.20/download",
        type = "tar.gz",
        sha256 = "b5e6163cb8c49088c2c36f57875e58ccd8c87c7427f7fbd50ea6710b2f3f2e8f",
        strip_prefix = "log-0.4.20",
        build_file = Label("//third_party/rust_efi/cargo/remote:BUILD.log-0.4.20.bazel"),
    )

    maybe(
        http_archive,
        name = "rsefi__proc_macro2__1_0_67",
        url = "https://crates.io/api/v1/crates/proc-macro2/1.0.67/download",
        type = "tar.gz",
        sha256 = "3d433d9f1a3e8c1263d9456598b16fec66f4acc9a74dacffd35c7bb09b3a1328",
        strip_prefix = "proc-macro2-1.0.67",
        build_file = Label("//third_party/rust_efi/cargo/remote:BUILD.proc-macro2-1.0.67.bazel"),
    )

    maybe(
        http_archive,
        name = "rsefi__prost__0_12_1",
        url = "https://crates.io/api/v1/crates/prost/0.12.1/download",
        type = "tar.gz",
        sha256 = "f4fdd22f3b9c31b53c060df4a0613a1c7f062d4115a2b984dd15b1858f7e340d",
        strip_prefix = "prost-0.12.1",
        build_file = Label("//third_party/rust_efi/cargo/remote:BUILD.prost-0.12.1.bazel"),
    )

    maybe(
        http_archive,
        name = "rsefi__prost_derive__0_12_1",
        url = "https://crates.io/api/v1/crates/prost-derive/0.12.1/download",
        type = "tar.gz",
        sha256 = "265baba7fabd416cf5078179f7d2cbeca4ce7a9041111900675ea7c4cb8a4c32",
        strip_prefix = "prost-derive-0.12.1",
        build_file = Label("//third_party/rust_efi/cargo/remote:BUILD.prost-derive-0.12.1.bazel"),
    )

    maybe(
        http_archive,
        name = "rsefi__prost_types__0_12_1",
        url = "https://crates.io/api/v1/crates/prost-types/0.12.1/download",
        type = "tar.gz",
        sha256 = "e081b29f63d83a4bc75cfc9f3fe424f9156cf92d8a4f0c9407cce9a1b67327cf",
        strip_prefix = "prost-types-0.12.1",
        build_file = Label("//third_party/rust_efi/cargo/remote:BUILD.prost-types-0.12.1.bazel"),
    )

    maybe(
        http_archive,
        name = "rsefi__ptr_meta__0_2_0",
        url = "https://crates.io/api/v1/crates/ptr_meta/0.2.0/download",
        type = "tar.gz",
        sha256 = "bcada80daa06c42ed5f48c9a043865edea5dc44cbf9ac009fda3b89526e28607",
        strip_prefix = "ptr_meta-0.2.0",
        build_file = Label("//third_party/rust_efi/cargo/remote:BUILD.ptr_meta-0.2.0.bazel"),
    )

    maybe(
        http_archive,
        name = "rsefi__ptr_meta_derive__0_2_0",
        url = "https://crates.io/api/v1/crates/ptr_meta_derive/0.2.0/download",
        type = "tar.gz",
        sha256 = "bca9224df2e20e7c5548aeb5f110a0f3b77ef05f8585139b7148b59056168ed2",
        strip_prefix = "ptr_meta_derive-0.2.0",
        build_file = Label("//third_party/rust_efi/cargo/remote:BUILD.ptr_meta_derive-0.2.0.bazel"),
    )

    maybe(
        http_archive,
        name = "rsefi__quote__1_0_33",
        url = "https://crates.io/api/v1/crates/quote/1.0.33/download",
        type = "tar.gz",
        sha256 = "5267fca4496028628a95160fc423a33e8b2e6af8a5302579e322e4b520293cae",
        strip_prefix = "quote-1.0.33",
        build_file = Label("//third_party/rust_efi/cargo/remote:BUILD.quote-1.0.33.bazel"),
    )

    maybe(
        http_archive,
        name = "rsefi__syn__1_0_109",
        url = "https://crates.io/api/v1/crates/syn/1.0.109/download",
        type = "tar.gz",
        sha256 = "72b64191b275b66ffe2469e8af2c1cfe3bafa67b529ead792a6d0160888b4237",
        strip_prefix = "syn-1.0.109",
        build_file = Label("//third_party/rust_efi/cargo/remote:BUILD.syn-1.0.109.bazel"),
    )

    maybe(
        http_archive,
        name = "rsefi__syn__2_0_37",
        url = "https://crates.io/api/v1/crates/syn/2.0.37/download",
        type = "tar.gz",
        sha256 = "7303ef2c05cd654186cb250d29049a24840ca25d2747c25c0381c8d9e2f582e8",
        strip_prefix = "syn-2.0.37",
        build_file = Label("//third_party/rust_efi/cargo/remote:BUILD.syn-2.0.37.bazel"),
    )

    maybe(
        http_archive,
        name = "rsefi__ucs2__0_3_2",
        url = "https://crates.io/api/v1/crates/ucs2/0.3.2/download",
        type = "tar.gz",
        sha256 = "bad643914094137d475641b6bab89462505316ec2ce70907ad20102d28a79ab8",
        strip_prefix = "ucs2-0.3.2",
        build_file = Label("//third_party/rust_efi/cargo/remote:BUILD.ucs2-0.3.2.bazel"),
    )

    maybe(
        http_archive,
        name = "rsefi__uefi__0_24_0",
        url = "https://crates.io/api/v1/crates/uefi/0.24.0/download",
        type = "tar.gz",
        sha256 = "3b63e82686b4bdb0db74f18b2abbd60a0470354fb640aa69e115598d714d0a10",
        strip_prefix = "uefi-0.24.0",
        build_file = Label("//third_party/rust_efi/cargo/remote:BUILD.uefi-0.24.0.bazel"),
    )

    maybe(
        http_archive,
        name = "rsefi__uefi_macros__0_12_0",
        url = "https://crates.io/api/v1/crates/uefi-macros/0.12.0/download",
        type = "tar.gz",
        sha256 = "023d94ef8e135d068b9a3bd94614ef2610b2b0419ade0a9d8f3501fa9cd08e95",
        strip_prefix = "uefi-macros-0.12.0",
        build_file = Label("//third_party/rust_efi/cargo/remote:BUILD.uefi-macros-0.12.0.bazel"),
    )

    maybe(
        http_archive,
        name = "rsefi__uefi_raw__0_3_0",
        url = "https://crates.io/api/v1/crates/uefi-raw/0.3.0/download",
        type = "tar.gz",
        sha256 = "62642516099c6441a5f41b0da8486d5fc3515a0603b0fdaea67b31600e22082e",
        strip_prefix = "uefi-raw-0.3.0",
        build_file = Label("//third_party/rust_efi/cargo/remote:BUILD.uefi-raw-0.3.0.bazel"),
    )

    maybe(
        http_archive,
        name = "rsefi__uefi_services__0_21_0",
        url = "https://crates.io/api/v1/crates/uefi-services/0.21.0/download",
        type = "tar.gz",
        sha256 = "44b32954ebbb4be5ebfde0df6699c2091f04e9f9c3762c65f3435dfb1a90a668",
        strip_prefix = "uefi-services-0.21.0",
        build_file = Label("//third_party/rust_efi/cargo/remote:BUILD.uefi-services-0.21.0.bazel"),
    )

    maybe(
        http_archive,
        name = "rsefi__uguid__2_0_1",
        url = "https://crates.io/api/v1/crates/uguid/2.0.1/download",
        type = "tar.gz",
        sha256 = "16dfbd255defbd727b3a30e8950695d2e6d045841ee250ff0f1f7ced17917f8d",
        strip_prefix = "uguid-2.0.1",
        build_file = Label("//third_party/rust_efi/cargo/remote:BUILD.uguid-2.0.1.bazel"),
    )

    maybe(
        http_archive,
        name = "rsefi__unicode_ident__1_0_12",
        url = "https://crates.io/api/v1/crates/unicode-ident/1.0.12/download",
        type = "tar.gz",
        sha256 = "3354b9ac3fae1ff6755cb6db53683adb661634f67557942dea4facebec0fee4b",
        strip_prefix = "unicode-ident-1.0.12",
        build_file = Label("//third_party/rust_efi/cargo/remote:BUILD.unicode-ident-1.0.12.bazel"),
    )
