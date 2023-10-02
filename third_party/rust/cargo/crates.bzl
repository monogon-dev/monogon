"""
@generated
cargo-raze generated Bazel file.

DO NOT EDIT! Replaced on runs of cargo-raze
"""

load("@bazel_tools//tools/build_defs/repo:git.bzl", "new_git_repository")  # buildifier: disable=load
load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")  # buildifier: disable=load
load("@bazel_tools//tools/build_defs/repo:utils.bzl", "maybe")  # buildifier: disable=load

def raze_fetch_remote_crates():
    """This function defines a collection of repos and should be called in a WORKSPACE file"""
    maybe(
        http_archive,
        name = "raze__addr2line__0_21_0",
        url = "https://crates.io/api/v1/crates/addr2line/0.21.0/download",
        type = "tar.gz",
        sha256 = "8a30b2e23b9e17a9f90641c7ab1549cd9b44f296d3ccbf309d2863cfe398a0cb",
        strip_prefix = "addr2line-0.21.0",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.addr2line-0.21.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__adler__1_0_2",
        url = "https://crates.io/api/v1/crates/adler/1.0.2/download",
        type = "tar.gz",
        sha256 = "f26201604c87b1e01bd3d98f8d5d9a8fcbb815e8cedb41ffccbeb4bf593a35fe",
        strip_prefix = "adler-1.0.2",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.adler-1.0.2.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__aho_corasick__1_1_1",
        url = "https://crates.io/api/v1/crates/aho-corasick/1.1.1/download",
        type = "tar.gz",
        sha256 = "ea5d730647d4fadd988536d06fecce94b7b4f2a7efdae548f1cf4b63205518ab",
        strip_prefix = "aho-corasick-1.1.1",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.aho-corasick-1.1.1.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__ammonia__3_3_0",
        url = "https://crates.io/api/v1/crates/ammonia/3.3.0/download",
        type = "tar.gz",
        sha256 = "64e6d1c7838db705c9b756557ee27c384ce695a1c51a6fe528784cb1c6840170",
        strip_prefix = "ammonia-3.3.0",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.ammonia-3.3.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__android_tzdata__0_1_1",
        url = "https://crates.io/api/v1/crates/android-tzdata/0.1.1/download",
        type = "tar.gz",
        sha256 = "e999941b234f3131b00bc13c22d06e8c5ff726d1b6318ac7eb276997bbb4fef0",
        strip_prefix = "android-tzdata-0.1.1",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.android-tzdata-0.1.1.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__android_system_properties__0_1_5",
        url = "https://crates.io/api/v1/crates/android_system_properties/0.1.5/download",
        type = "tar.gz",
        sha256 = "819e7219dbd41043ac279b19830f2efc897156490d7fd6ea916720117ee66311",
        strip_prefix = "android_system_properties-0.1.5",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.android_system_properties-0.1.5.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__anstream__0_6_4",
        url = "https://crates.io/api/v1/crates/anstream/0.6.4/download",
        type = "tar.gz",
        sha256 = "2ab91ebe16eb252986481c5b62f6098f3b698a45e34b5b98200cf20dd2484a44",
        strip_prefix = "anstream-0.6.4",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.anstream-0.6.4.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__anstyle__1_0_4",
        url = "https://crates.io/api/v1/crates/anstyle/1.0.4/download",
        type = "tar.gz",
        sha256 = "7079075b41f533b8c61d2a4d073c4676e1f8b249ff94a393b0595db304e0dd87",
        strip_prefix = "anstyle-1.0.4",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.anstyle-1.0.4.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__anstyle_parse__0_2_2",
        url = "https://crates.io/api/v1/crates/anstyle-parse/0.2.2/download",
        type = "tar.gz",
        sha256 = "317b9a89c1868f5ea6ff1d9539a69f45dffc21ce321ac1fd1160dfa48c8e2140",
        strip_prefix = "anstyle-parse-0.2.2",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.anstyle-parse-0.2.2.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__anstyle_query__1_0_0",
        url = "https://crates.io/api/v1/crates/anstyle-query/1.0.0/download",
        type = "tar.gz",
        sha256 = "5ca11d4be1bab0c8bc8734a9aa7bf4ee8316d462a08c6ac5052f888fef5b494b",
        strip_prefix = "anstyle-query-1.0.0",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.anstyle-query-1.0.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__anstyle_wincon__3_0_1",
        url = "https://crates.io/api/v1/crates/anstyle-wincon/3.0.1/download",
        type = "tar.gz",
        sha256 = "f0699d10d2f4d628a98ee7b57b289abbc98ff3bad977cb3152709d4bf2330628",
        strip_prefix = "anstyle-wincon-3.0.1",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.anstyle-wincon-3.0.1.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__anyhow__1_0_75",
        url = "https://crates.io/api/v1/crates/anyhow/1.0.75/download",
        type = "tar.gz",
        sha256 = "a4668cab20f66d8d020e1fbc0ebe47217433c1b6c8f2040faf858554e394ace6",
        strip_prefix = "anyhow-1.0.75",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.anyhow-1.0.75.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__autocfg__1_1_0",
        url = "https://crates.io/api/v1/crates/autocfg/1.1.0/download",
        type = "tar.gz",
        sha256 = "d468802bab17cbc0cc575e9b053f41e72aa36bfa6b7f55e3529ffa43161b97fa",
        strip_prefix = "autocfg-1.1.0",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.autocfg-1.1.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__backtrace__0_3_69",
        url = "https://crates.io/api/v1/crates/backtrace/0.3.69/download",
        type = "tar.gz",
        sha256 = "2089b7e3f35b9dd2d0ed921ead4f6d318c27680d4a5bd167b3ee120edb105837",
        strip_prefix = "backtrace-0.3.69",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.backtrace-0.3.69.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__base64__0_21_4",
        url = "https://crates.io/api/v1/crates/base64/0.21.4/download",
        type = "tar.gz",
        sha256 = "9ba43ea6f343b788c8764558649e08df62f86c6ef251fdaeb1ffd010a9ae50a2",
        strip_prefix = "base64-0.21.4",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.base64-0.21.4.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__bitflags__1_3_2",
        url = "https://crates.io/api/v1/crates/bitflags/1.3.2/download",
        type = "tar.gz",
        sha256 = "bef38d45163c2f1dde094a7dfd33ccf595c92905c8f8f4fdc18d06fb1037718a",
        strip_prefix = "bitflags-1.3.2",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.bitflags-1.3.2.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__bitflags__2_4_0",
        url = "https://crates.io/api/v1/crates/bitflags/2.4.0/download",
        type = "tar.gz",
        sha256 = "b4682ae6287fcf752ecaabbfcc7b6f9b72aa33933dc23a554d853aea8eea8635",
        strip_prefix = "bitflags-2.4.0",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.bitflags-2.4.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__block_buffer__0_10_4",
        url = "https://crates.io/api/v1/crates/block-buffer/0.10.4/download",
        type = "tar.gz",
        sha256 = "3078c7629b62d3f0439517fa394996acacc5cbc91c5a20d8c658e77abd503a71",
        strip_prefix = "block-buffer-0.10.4",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.block-buffer-0.10.4.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__bstr__1_6_2",
        url = "https://crates.io/api/v1/crates/bstr/1.6.2/download",
        type = "tar.gz",
        sha256 = "4c2f7349907b712260e64b0afe2f84692af14a454be26187d9df565c7f69266a",
        strip_prefix = "bstr-1.6.2",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.bstr-1.6.2.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__bumpalo__3_14_0",
        url = "https://crates.io/api/v1/crates/bumpalo/3.14.0/download",
        type = "tar.gz",
        sha256 = "7f30e7476521f6f8af1a1c4c0b8cc94f0bee37d91763d0ca2665f299b6cd8aec",
        strip_prefix = "bumpalo-3.14.0",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.bumpalo-3.14.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__byteorder__1_4_3",
        url = "https://crates.io/api/v1/crates/byteorder/1.4.3/download",
        type = "tar.gz",
        sha256 = "14c189c53d098945499cdfa7ecc63567cf3886b3332b312a5b4585d8d3a6a610",
        strip_prefix = "byteorder-1.4.3",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.byteorder-1.4.3.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__bytes__1_5_0",
        url = "https://crates.io/api/v1/crates/bytes/1.5.0/download",
        type = "tar.gz",
        sha256 = "a2bd12c1caf447e69cd4528f47f94d203fd2582878ecb9e9465484c4148a8223",
        strip_prefix = "bytes-1.5.0",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.bytes-1.5.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__cc__1_0_83",
        url = "https://crates.io/api/v1/crates/cc/1.0.83/download",
        type = "tar.gz",
        sha256 = "f1174fb0b6ec23863f8b971027804a42614e347eafb0a95bf0b12cdae21fc4d0",
        strip_prefix = "cc-1.0.83",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.cc-1.0.83.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__cfg_if__1_0_0",
        url = "https://crates.io/api/v1/crates/cfg-if/1.0.0/download",
        type = "tar.gz",
        sha256 = "baf1de4339761588bc0619e3cbc0120ee582ebb74b53b4efbf79117bd2da40fd",
        strip_prefix = "cfg-if-1.0.0",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.cfg-if-1.0.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__chrono__0_4_31",
        url = "https://crates.io/api/v1/crates/chrono/0.4.31/download",
        type = "tar.gz",
        sha256 = "7f2c685bad3eb3d45a01354cedb7d5faa66194d1d58ba6e267a8de788f79db38",
        strip_prefix = "chrono-0.4.31",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.chrono-0.4.31.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__clap__4_4_6",
        url = "https://crates.io/api/v1/crates/clap/4.4.6/download",
        type = "tar.gz",
        sha256 = "d04704f56c2cde07f43e8e2c154b43f216dc5c92fc98ada720177362f953b956",
        strip_prefix = "clap-4.4.6",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.clap-4.4.6.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__clap_builder__4_4_6",
        url = "https://crates.io/api/v1/crates/clap_builder/4.4.6/download",
        type = "tar.gz",
        sha256 = "0e231faeaca65ebd1ea3c737966bf858971cd38c3849107aa3ea7de90a804e45",
        strip_prefix = "clap_builder-4.4.6",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.clap_builder-4.4.6.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__clap_complete__4_4_3",
        url = "https://crates.io/api/v1/crates/clap_complete/4.4.3/download",
        type = "tar.gz",
        sha256 = "e3ae8ba90b9d8b007efe66e55e48fb936272f5ca00349b5b0e89877520d35ea7",
        strip_prefix = "clap_complete-4.4.3",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.clap_complete-4.4.3.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__clap_lex__0_5_1",
        url = "https://crates.io/api/v1/crates/clap_lex/0.5.1/download",
        type = "tar.gz",
        sha256 = "cd7cc57abe963c6d3b9d8be5b06ba7c8957a930305ca90304f24ef040aa6f961",
        strip_prefix = "clap_lex-0.5.1",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.clap_lex-0.5.1.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__colorchoice__1_0_0",
        url = "https://crates.io/api/v1/crates/colorchoice/1.0.0/download",
        type = "tar.gz",
        sha256 = "acbf1af155f9b9ef647e42cdc158db4b64a1b61f743629225fde6f3e0be2a7c7",
        strip_prefix = "colorchoice-1.0.0",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.colorchoice-1.0.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__core_foundation_sys__0_8_4",
        url = "https://crates.io/api/v1/crates/core-foundation-sys/0.8.4/download",
        type = "tar.gz",
        sha256 = "e496a50fda8aacccc86d7529e2c1e0892dbd0f898a6b5645b5561b89c3210efa",
        strip_prefix = "core-foundation-sys-0.8.4",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.core-foundation-sys-0.8.4.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__cpufeatures__0_2_9",
        url = "https://crates.io/api/v1/crates/cpufeatures/0.2.9/download",
        type = "tar.gz",
        sha256 = "a17b76ff3a4162b0b27f354a0c87015ddad39d35f9c0c36607a3bdd175dde1f1",
        strip_prefix = "cpufeatures-0.2.9",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.cpufeatures-0.2.9.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__crossbeam_channel__0_5_8",
        url = "https://crates.io/api/v1/crates/crossbeam-channel/0.5.8/download",
        type = "tar.gz",
        sha256 = "a33c2bf77f2df06183c3aa30d1e96c0695a313d4f9c453cc3762a6db39f99200",
        strip_prefix = "crossbeam-channel-0.5.8",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.crossbeam-channel-0.5.8.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__crossbeam_utils__0_8_16",
        url = "https://crates.io/api/v1/crates/crossbeam-utils/0.8.16/download",
        type = "tar.gz",
        sha256 = "5a22b2d63d4d1dc0b7f1b6b2747dd0088008a9be28b6ddf0b1e7d335e3037294",
        strip_prefix = "crossbeam-utils-0.8.16",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.crossbeam-utils-0.8.16.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__crypto_common__0_1_6",
        url = "https://crates.io/api/v1/crates/crypto-common/0.1.6/download",
        type = "tar.gz",
        sha256 = "1bfb12502f3fc46cca1bb51ac28df9d618d813cdc3d2f25b9fe775a34af26bb3",
        strip_prefix = "crypto-common-0.1.6",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.crypto-common-0.1.6.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__data_encoding__2_4_0",
        url = "https://crates.io/api/v1/crates/data-encoding/2.4.0/download",
        type = "tar.gz",
        sha256 = "c2e66c9d817f1720209181c316d28635c050fa304f9c79e47a520882661b7308",
        strip_prefix = "data-encoding-2.4.0",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.data-encoding-2.4.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__digest__0_10_7",
        url = "https://crates.io/api/v1/crates/digest/0.10.7/download",
        type = "tar.gz",
        sha256 = "9ed9a281f7bc9b7576e61468ba615a66a5c8cfdff42420a70aa82701a3b1e292",
        strip_prefix = "digest-0.10.7",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.digest-0.10.7.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__either__1_9_0",
        url = "https://crates.io/api/v1/crates/either/1.9.0/download",
        type = "tar.gz",
        sha256 = "a26ae43d7bcc3b814de94796a5e736d4029efb0ee900c12e2d54c993ad1a1e07",
        strip_prefix = "either-1.9.0",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.either-1.9.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__elasticlunr_rs__3_0_2",
        url = "https://crates.io/api/v1/crates/elasticlunr-rs/3.0.2/download",
        type = "tar.gz",
        sha256 = "41e83863a500656dfa214fee6682de9c5b9f03de6860fec531235ed2ae9f6571",
        strip_prefix = "elasticlunr-rs-3.0.2",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.elasticlunr-rs-3.0.2.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__env_logger__0_10_0",
        url = "https://crates.io/api/v1/crates/env_logger/0.10.0/download",
        type = "tar.gz",
        sha256 = "85cdab6a89accf66733ad5a1693a4dcced6aeff64602b634530dd73c1f3ee9f0",
        strip_prefix = "env_logger-0.10.0",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.env_logger-0.10.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__equivalent__1_0_1",
        url = "https://crates.io/api/v1/crates/equivalent/1.0.1/download",
        type = "tar.gz",
        sha256 = "5443807d6dff69373d433ab9ef5378ad8df50ca6298caf15de6e52e24aaf54d5",
        strip_prefix = "equivalent-1.0.1",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.equivalent-1.0.1.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__errno__0_3_4",
        url = "https://crates.io/api/v1/crates/errno/0.3.4/download",
        type = "tar.gz",
        sha256 = "add4f07d43996f76ef320709726a556a9d4f965d9410d8d0271132d2f8293480",
        strip_prefix = "errno-0.3.4",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.errno-0.3.4.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__errno_dragonfly__0_1_2",
        url = "https://crates.io/api/v1/crates/errno-dragonfly/0.1.2/download",
        type = "tar.gz",
        sha256 = "aa68f1b12764fab894d2755d2518754e71b4fd80ecfb822714a1206c2aab39bf",
        strip_prefix = "errno-dragonfly-0.1.2",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.errno-dragonfly-0.1.2.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__fastrand__2_0_1",
        url = "https://crates.io/api/v1/crates/fastrand/2.0.1/download",
        type = "tar.gz",
        sha256 = "25cbce373ec4653f1a01a31e8a5e5ec0c622dc27ff9c4e6606eefef5cbbed4a5",
        strip_prefix = "fastrand-2.0.1",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.fastrand-2.0.1.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__filetime__0_2_22",
        url = "https://crates.io/api/v1/crates/filetime/0.2.22/download",
        type = "tar.gz",
        sha256 = "d4029edd3e734da6fe05b6cd7bd2960760a616bd2ddd0d59a0124746d6272af0",
        strip_prefix = "filetime-0.2.22",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.filetime-0.2.22.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__fixedbitset__0_4_2",
        url = "https://crates.io/api/v1/crates/fixedbitset/0.4.2/download",
        type = "tar.gz",
        sha256 = "0ce7134b9999ecaf8bcd65542e436736ef32ddca1b3e06094cb6ec5755203b80",
        strip_prefix = "fixedbitset-0.4.2",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.fixedbitset-0.4.2.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__fnv__1_0_7",
        url = "https://crates.io/api/v1/crates/fnv/1.0.7/download",
        type = "tar.gz",
        sha256 = "3f9eec918d3f24069decb9af1554cad7c880e2da24a9afd88aca000531ab82c1",
        strip_prefix = "fnv-1.0.7",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.fnv-1.0.7.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__form_urlencoded__1_2_0",
        url = "https://crates.io/api/v1/crates/form_urlencoded/1.2.0/download",
        type = "tar.gz",
        sha256 = "a62bc1cf6f830c2ec14a513a9fb124d0a213a629668a4186f329db21fe045652",
        strip_prefix = "form_urlencoded-1.2.0",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.form_urlencoded-1.2.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__fsevent_sys__4_1_0",
        url = "https://crates.io/api/v1/crates/fsevent-sys/4.1.0/download",
        type = "tar.gz",
        sha256 = "76ee7a02da4d231650c7cea31349b889be2f45ddb3ef3032d2ec8185f6313fd2",
        strip_prefix = "fsevent-sys-4.1.0",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.fsevent-sys-4.1.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__futf__0_1_5",
        url = "https://crates.io/api/v1/crates/futf/0.1.5/download",
        type = "tar.gz",
        sha256 = "df420e2e84819663797d1ec6544b13c5be84629e7bb00dc960d6917db2987843",
        strip_prefix = "futf-0.1.5",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.futf-0.1.5.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__futures_channel__0_3_28",
        url = "https://crates.io/api/v1/crates/futures-channel/0.3.28/download",
        type = "tar.gz",
        sha256 = "955518d47e09b25bbebc7a18df10b81f0c766eaf4c4f1cccef2fca5f2a4fb5f2",
        strip_prefix = "futures-channel-0.3.28",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.futures-channel-0.3.28.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__futures_core__0_3_28",
        url = "https://crates.io/api/v1/crates/futures-core/0.3.28/download",
        type = "tar.gz",
        sha256 = "4bca583b7e26f571124fe5b7561d49cb2868d79116cfa0eefce955557c6fee8c",
        strip_prefix = "futures-core-0.3.28",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.futures-core-0.3.28.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__futures_macro__0_3_28",
        url = "https://crates.io/api/v1/crates/futures-macro/0.3.28/download",
        type = "tar.gz",
        sha256 = "89ca545a94061b6365f2c7355b4b32bd20df3ff95f02da9329b34ccc3bd6ee72",
        strip_prefix = "futures-macro-0.3.28",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.futures-macro-0.3.28.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__futures_sink__0_3_28",
        url = "https://crates.io/api/v1/crates/futures-sink/0.3.28/download",
        type = "tar.gz",
        sha256 = "f43be4fe21a13b9781a69afa4985b0f6ee0e1afab2c6f454a8cf30e2b2237b6e",
        strip_prefix = "futures-sink-0.3.28",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.futures-sink-0.3.28.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__futures_task__0_3_28",
        url = "https://crates.io/api/v1/crates/futures-task/0.3.28/download",
        type = "tar.gz",
        sha256 = "76d3d132be6c0e6aa1534069c705a74a5997a356c0dc2f86a47765e5617c5b65",
        strip_prefix = "futures-task-0.3.28",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.futures-task-0.3.28.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__futures_util__0_3_28",
        url = "https://crates.io/api/v1/crates/futures-util/0.3.28/download",
        type = "tar.gz",
        sha256 = "26b01e40b772d54cf6c6d721c1d1abd0647a0106a12ecaa1c186273392a69533",
        strip_prefix = "futures-util-0.3.28",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.futures-util-0.3.28.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__generic_array__0_14_7",
        url = "https://crates.io/api/v1/crates/generic-array/0.14.7/download",
        type = "tar.gz",
        sha256 = "85649ca51fd72272d7821adaf274ad91c288277713d9c18820d8499a7ff69e9a",
        strip_prefix = "generic-array-0.14.7",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.generic-array-0.14.7.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__getrandom__0_2_10",
        url = "https://crates.io/api/v1/crates/getrandom/0.2.10/download",
        type = "tar.gz",
        sha256 = "be4136b2a15dd319360be1c07d9933517ccf0be8f16bf62a3bee4f0d618df427",
        strip_prefix = "getrandom-0.2.10",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.getrandom-0.2.10.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__gimli__0_28_0",
        url = "https://crates.io/api/v1/crates/gimli/0.28.0/download",
        type = "tar.gz",
        sha256 = "6fb8d784f27acf97159b40fc4db5ecd8aa23b9ad5ef69cdd136d3bc80665f0c0",
        strip_prefix = "gimli-0.28.0",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.gimli-0.28.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__globset__0_4_13",
        url = "https://crates.io/api/v1/crates/globset/0.4.13/download",
        type = "tar.gz",
        sha256 = "759c97c1e17c55525b57192c06a267cda0ac5210b222d6b82189a2338fa1c13d",
        strip_prefix = "globset-0.4.13",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.globset-0.4.13.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__h2__0_3_21",
        url = "https://crates.io/api/v1/crates/h2/0.3.21/download",
        type = "tar.gz",
        sha256 = "91fc23aa11be92976ef4729127f1a74adf36d8436f7816b185d18df956790833",
        strip_prefix = "h2-0.3.21",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.h2-0.3.21.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__handlebars__4_4_0",
        url = "https://crates.io/api/v1/crates/handlebars/4.4.0/download",
        type = "tar.gz",
        sha256 = "c39b3bc2a8f715298032cf5087e58573809374b08160aa7d750582bdb82d2683",
        strip_prefix = "handlebars-4.4.0",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.handlebars-4.4.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__hashbrown__0_12_3",
        url = "https://crates.io/api/v1/crates/hashbrown/0.12.3/download",
        type = "tar.gz",
        sha256 = "8a9ee70c43aaf417c914396645a0fa852624801b24ebb7ae78fe8272889ac888",
        strip_prefix = "hashbrown-0.12.3",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.hashbrown-0.12.3.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__hashbrown__0_14_1",
        url = "https://crates.io/api/v1/crates/hashbrown/0.14.1/download",
        type = "tar.gz",
        sha256 = "7dfda62a12f55daeae5015f81b0baea145391cb4520f86c248fc615d72640d12",
        strip_prefix = "hashbrown-0.14.1",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.hashbrown-0.14.1.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__headers__0_3_9",
        url = "https://crates.io/api/v1/crates/headers/0.3.9/download",
        type = "tar.gz",
        sha256 = "06683b93020a07e3dbcf5f8c0f6d40080d725bea7936fc01ad345c01b97dc270",
        strip_prefix = "headers-0.3.9",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.headers-0.3.9.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__headers_core__0_2_0",
        url = "https://crates.io/api/v1/crates/headers-core/0.2.0/download",
        type = "tar.gz",
        sha256 = "e7f66481bfee273957b1f20485a4ff3362987f85b2c236580d81b4eb7a326429",
        strip_prefix = "headers-core-0.2.0",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.headers-core-0.2.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__heck__0_4_1",
        url = "https://crates.io/api/v1/crates/heck/0.4.1/download",
        type = "tar.gz",
        sha256 = "95505c38b4572b2d910cecb0281560f54b440a19336cbbcb27bf6ce6adc6f5a8",
        strip_prefix = "heck-0.4.1",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.heck-0.4.1.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__hermit_abi__0_3_3",
        url = "https://crates.io/api/v1/crates/hermit-abi/0.3.3/download",
        type = "tar.gz",
        sha256 = "d77f7ec81a6d05a3abb01ab6eb7590f6083d08449fe5a1c8b1e620283546ccb7",
        strip_prefix = "hermit-abi-0.3.3",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.hermit-abi-0.3.3.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__home__0_5_5",
        url = "https://crates.io/api/v1/crates/home/0.5.5/download",
        type = "tar.gz",
        sha256 = "5444c27eef6923071f7ebcc33e3444508466a76f7a2b93da00ed6e19f30c1ddb",
        strip_prefix = "home-0.5.5",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.home-0.5.5.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__html5ever__0_26_0",
        url = "https://crates.io/api/v1/crates/html5ever/0.26.0/download",
        type = "tar.gz",
        sha256 = "bea68cab48b8459f17cf1c944c67ddc572d272d9f2b274140f223ecb1da4a3b7",
        strip_prefix = "html5ever-0.26.0",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.html5ever-0.26.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__http__0_2_9",
        url = "https://crates.io/api/v1/crates/http/0.2.9/download",
        type = "tar.gz",
        sha256 = "bd6effc99afb63425aff9b05836f029929e345a6148a14b7ecd5ab67af944482",
        strip_prefix = "http-0.2.9",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.http-0.2.9.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__http_body__0_4_5",
        url = "https://crates.io/api/v1/crates/http-body/0.4.5/download",
        type = "tar.gz",
        sha256 = "d5f38f16d184e36f2408a55281cd658ecbd3ca05cce6d6510a176eca393e26d1",
        strip_prefix = "http-body-0.4.5",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.http-body-0.4.5.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__httparse__1_8_0",
        url = "https://crates.io/api/v1/crates/httparse/1.8.0/download",
        type = "tar.gz",
        sha256 = "d897f394bad6a705d5f4104762e116a75639e470d80901eed05a860a95cb1904",
        strip_prefix = "httparse-1.8.0",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.httparse-1.8.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__httpdate__1_0_3",
        url = "https://crates.io/api/v1/crates/httpdate/1.0.3/download",
        type = "tar.gz",
        sha256 = "df3b46402a9d5adb4c86a0cf463f42e19994e3ee891101b1841f30a545cb49a9",
        strip_prefix = "httpdate-1.0.3",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.httpdate-1.0.3.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__humantime__2_1_0",
        url = "https://crates.io/api/v1/crates/humantime/2.1.0/download",
        type = "tar.gz",
        sha256 = "9a3a5bfb195931eeb336b2a7b4d761daec841b97f947d34394601737a7bba5e4",
        strip_prefix = "humantime-2.1.0",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.humantime-2.1.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__hyper__0_14_27",
        url = "https://crates.io/api/v1/crates/hyper/0.14.27/download",
        type = "tar.gz",
        sha256 = "ffb1cfd654a8219eaef89881fdb3bb3b1cdc5fa75ded05d6933b2b382e395468",
        strip_prefix = "hyper-0.14.27",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.hyper-0.14.27.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__iana_time_zone__0_1_57",
        url = "https://crates.io/api/v1/crates/iana-time-zone/0.1.57/download",
        type = "tar.gz",
        sha256 = "2fad5b825842d2b38bd206f3e81d6957625fd7f0a361e345c30e01a0ae2dd613",
        strip_prefix = "iana-time-zone-0.1.57",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.iana-time-zone-0.1.57.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__iana_time_zone_haiku__0_1_2",
        url = "https://crates.io/api/v1/crates/iana-time-zone-haiku/0.1.2/download",
        type = "tar.gz",
        sha256 = "f31827a206f56af32e590ba56d5d2d085f558508192593743f16b2306495269f",
        strip_prefix = "iana-time-zone-haiku-0.1.2",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.iana-time-zone-haiku-0.1.2.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__idna__0_4_0",
        url = "https://crates.io/api/v1/crates/idna/0.4.0/download",
        type = "tar.gz",
        sha256 = "7d20d6b07bfbc108882d88ed8e37d39636dcc260e15e30c45e6ba089610b917c",
        strip_prefix = "idna-0.4.0",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.idna-0.4.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__ignore__0_4_20",
        url = "https://crates.io/api/v1/crates/ignore/0.4.20/download",
        type = "tar.gz",
        sha256 = "dbe7873dab538a9a44ad79ede1faf5f30d49f9a5c883ddbab48bce81b64b7492",
        strip_prefix = "ignore-0.4.20",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.ignore-0.4.20.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__indexmap__1_9_3",
        url = "https://crates.io/api/v1/crates/indexmap/1.9.3/download",
        type = "tar.gz",
        sha256 = "bd070e393353796e801d209ad339e89596eb4c8d430d18ede6a1cced8fafbd99",
        strip_prefix = "indexmap-1.9.3",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.indexmap-1.9.3.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__indexmap__2_0_2",
        url = "https://crates.io/api/v1/crates/indexmap/2.0.2/download",
        type = "tar.gz",
        sha256 = "8adf3ddd720272c6ea8bf59463c04e0f93d0bbf7c5439b691bca2987e0270897",
        strip_prefix = "indexmap-2.0.2",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.indexmap-2.0.2.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__inotify__0_9_6",
        url = "https://crates.io/api/v1/crates/inotify/0.9.6/download",
        type = "tar.gz",
        sha256 = "f8069d3ec154eb856955c1c0fbffefbf5f3c40a104ec912d4797314c1801abff",
        strip_prefix = "inotify-0.9.6",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.inotify-0.9.6.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__inotify_sys__0_1_5",
        url = "https://crates.io/api/v1/crates/inotify-sys/0.1.5/download",
        type = "tar.gz",
        sha256 = "e05c02b5e89bff3b946cedeca278abc628fe811e604f027c45a8aa3cf793d0eb",
        strip_prefix = "inotify-sys-0.1.5",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.inotify-sys-0.1.5.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__is_terminal__0_4_9",
        url = "https://crates.io/api/v1/crates/is-terminal/0.4.9/download",
        type = "tar.gz",
        sha256 = "cb0889898416213fab133e1d33a0e5858a48177452750691bde3666d0fdbaf8b",
        strip_prefix = "is-terminal-0.4.9",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.is-terminal-0.4.9.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__itertools__0_10_5",
        url = "https://crates.io/api/v1/crates/itertools/0.10.5/download",
        type = "tar.gz",
        sha256 = "b0fd2260e829bddf4cb6ea802289de2f86d6a7a690192fbe91b3f46e0f2c8473",
        strip_prefix = "itertools-0.10.5",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.itertools-0.10.5.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__itoa__1_0_9",
        url = "https://crates.io/api/v1/crates/itoa/1.0.9/download",
        type = "tar.gz",
        sha256 = "af150ab688ff2122fcef229be89cb50dd66af9e01a4ff320cc137eecc9bacc38",
        strip_prefix = "itoa-1.0.9",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.itoa-1.0.9.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__js_sys__0_3_64",
        url = "https://crates.io/api/v1/crates/js-sys/0.3.64/download",
        type = "tar.gz",
        sha256 = "c5f195fe497f702db0f318b07fdd68edb16955aed830df8363d837542f8f935a",
        strip_prefix = "js-sys-0.3.64",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.js-sys-0.3.64.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__kqueue__1_0_8",
        url = "https://crates.io/api/v1/crates/kqueue/1.0.8/download",
        type = "tar.gz",
        sha256 = "7447f1ca1b7b563588a205fe93dea8df60fd981423a768bc1c0ded35ed147d0c",
        strip_prefix = "kqueue-1.0.8",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.kqueue-1.0.8.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__kqueue_sys__1_0_4",
        url = "https://crates.io/api/v1/crates/kqueue-sys/1.0.4/download",
        type = "tar.gz",
        sha256 = "ed9625ffda8729b85e45cf04090035ac368927b8cebc34898e7c120f52e4838b",
        strip_prefix = "kqueue-sys-1.0.4",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.kqueue-sys-1.0.4.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__lazy_static__1_4_0",
        url = "https://crates.io/api/v1/crates/lazy_static/1.4.0/download",
        type = "tar.gz",
        sha256 = "e2abad23fbc42b3700f2f279844dc832adb2b2eb069b2df918f455c4e18cc646",
        strip_prefix = "lazy_static-1.4.0",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.lazy_static-1.4.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__libc__0_2_148",
        url = "https://crates.io/api/v1/crates/libc/0.2.148/download",
        type = "tar.gz",
        sha256 = "9cdc71e17332e86d2e1d38c1f99edcb6288ee11b815fb1a4b049eaa2114d369b",
        strip_prefix = "libc-0.2.148",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.libc-0.2.148.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__linux_raw_sys__0_4_8",
        url = "https://crates.io/api/v1/crates/linux-raw-sys/0.4.8/download",
        type = "tar.gz",
        sha256 = "3852614a3bd9ca9804678ba6be5e3b8ce76dfc902cae004e3e0c44051b6e88db",
        strip_prefix = "linux-raw-sys-0.4.8",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.linux-raw-sys-0.4.8.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__lock_api__0_4_10",
        url = "https://crates.io/api/v1/crates/lock_api/0.4.10/download",
        type = "tar.gz",
        sha256 = "c1cc9717a20b1bb222f333e6a92fd32f7d8a18ddc5a3191a11af45dcbf4dcd16",
        strip_prefix = "lock_api-0.4.10",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.lock_api-0.4.10.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__log__0_4_20",
        url = "https://crates.io/api/v1/crates/log/0.4.20/download",
        type = "tar.gz",
        sha256 = "b5e6163cb8c49088c2c36f57875e58ccd8c87c7427f7fbd50ea6710b2f3f2e8f",
        strip_prefix = "log-0.4.20",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.log-0.4.20.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__mac__0_1_1",
        url = "https://crates.io/api/v1/crates/mac/0.1.1/download",
        type = "tar.gz",
        sha256 = "c41e0c4fef86961ac6d6f8a82609f55f31b05e4fce149ac5710e439df7619ba4",
        strip_prefix = "mac-0.1.1",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.mac-0.1.1.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__maplit__1_0_2",
        url = "https://crates.io/api/v1/crates/maplit/1.0.2/download",
        type = "tar.gz",
        sha256 = "3e2e65a1a2e43cfcb47a895c4c8b10d1f4a61097f9f254f183aee60cad9c651d",
        strip_prefix = "maplit-1.0.2",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.maplit-1.0.2.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__markup5ever__0_11_0",
        url = "https://crates.io/api/v1/crates/markup5ever/0.11.0/download",
        type = "tar.gz",
        sha256 = "7a2629bb1404f3d34c2e921f21fd34ba00b206124c81f65c50b43b6aaefeb016",
        strip_prefix = "markup5ever-0.11.0",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.markup5ever-0.11.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__mdbook__0_4_35",
        url = "https://crates.io/api/v1/crates/mdbook/0.4.35/download",
        type = "tar.gz",
        sha256 = "1c3f88addd34930bc5f01b9dc19f780447e51c92bf2536e3ded058018271775d",
        strip_prefix = "mdbook-0.4.35",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.mdbook-0.4.35.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__memchr__2_6_4",
        url = "https://crates.io/api/v1/crates/memchr/2.6.4/download",
        type = "tar.gz",
        sha256 = "f665ee40bc4a3c5590afb1e9677db74a508659dfd71e126420da8274909a0167",
        strip_prefix = "memchr-2.6.4",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.memchr-2.6.4.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__mime__0_3_17",
        url = "https://crates.io/api/v1/crates/mime/0.3.17/download",
        type = "tar.gz",
        sha256 = "6877bb514081ee2a7ff5ef9de3281f14a4dd4bceac4c09388074a6b5df8a139a",
        strip_prefix = "mime-0.3.17",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.mime-0.3.17.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__mime_guess__2_0_4",
        url = "https://crates.io/api/v1/crates/mime_guess/2.0.4/download",
        type = "tar.gz",
        sha256 = "4192263c238a5f0d0c6bfd21f336a313a4ce1c450542449ca191bb657b4642ef",
        strip_prefix = "mime_guess-2.0.4",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.mime_guess-2.0.4.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__miniz_oxide__0_7_1",
        url = "https://crates.io/api/v1/crates/miniz_oxide/0.7.1/download",
        type = "tar.gz",
        sha256 = "e7810e0be55b428ada41041c41f32c9f1a42817901b4ccf45fa3d4b6561e74c7",
        strip_prefix = "miniz_oxide-0.7.1",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.miniz_oxide-0.7.1.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__mio__0_8_8",
        url = "https://crates.io/api/v1/crates/mio/0.8.8/download",
        type = "tar.gz",
        sha256 = "927a765cd3fc26206e66b296465fa9d3e5ab003e651c1b3c060e7956d96b19d2",
        strip_prefix = "mio-0.8.8",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.mio-0.8.8.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__multimap__0_8_3",
        url = "https://crates.io/api/v1/crates/multimap/0.8.3/download",
        type = "tar.gz",
        sha256 = "e5ce46fe64a9d73be07dcbe690a38ce1b293be448fd8ce1e6c1b8062c9f72c6a",
        strip_prefix = "multimap-0.8.3",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.multimap-0.8.3.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__new_debug_unreachable__1_0_4",
        url = "https://crates.io/api/v1/crates/new_debug_unreachable/1.0.4/download",
        type = "tar.gz",
        sha256 = "e4a24736216ec316047a1fc4252e27dabb04218aa4a3f37c6e7ddbf1f9782b54",
        strip_prefix = "new_debug_unreachable-1.0.4",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.new_debug_unreachable-1.0.4.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__normpath__1_1_1",
        url = "https://crates.io/api/v1/crates/normpath/1.1.1/download",
        type = "tar.gz",
        sha256 = "ec60c60a693226186f5d6edf073232bfb6464ed97eb22cf3b01c1e8198fd97f5",
        strip_prefix = "normpath-1.1.1",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.normpath-1.1.1.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__notify__6_1_1",
        url = "https://crates.io/api/v1/crates/notify/6.1.1/download",
        type = "tar.gz",
        sha256 = "6205bd8bb1e454ad2e27422015fb5e4f2bcc7e08fa8f27058670d208324a4d2d",
        strip_prefix = "notify-6.1.1",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.notify-6.1.1.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__notify_debouncer_mini__0_3_0",
        url = "https://crates.io/api/v1/crates/notify-debouncer-mini/0.3.0/download",
        type = "tar.gz",
        sha256 = "e55ee272914f4563a2f8b8553eb6811f3c0caea81c756346bad15b7e3ef969f0",
        strip_prefix = "notify-debouncer-mini-0.3.0",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.notify-debouncer-mini-0.3.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__num_traits__0_2_16",
        url = "https://crates.io/api/v1/crates/num-traits/0.2.16/download",
        type = "tar.gz",
        sha256 = "f30b0abd723be7e2ffca1272140fac1a2f084c77ec3e123c192b66af1ee9e6c2",
        strip_prefix = "num-traits-0.2.16",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.num-traits-0.2.16.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__num_cpus__1_16_0",
        url = "https://crates.io/api/v1/crates/num_cpus/1.16.0/download",
        type = "tar.gz",
        sha256 = "4161fcb6d602d4d2081af7c3a45852d875a03dd337a6bfdd6e06407b61342a43",
        strip_prefix = "num_cpus-1.16.0",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.num_cpus-1.16.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__object__0_32_1",
        url = "https://crates.io/api/v1/crates/object/0.32.1/download",
        type = "tar.gz",
        sha256 = "9cf5f9dd3933bd50a9e1f149ec995f39ae2c496d31fd772c1fd45ebc27e902b0",
        strip_prefix = "object-0.32.1",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.object-0.32.1.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__once_cell__1_18_0",
        url = "https://crates.io/api/v1/crates/once_cell/1.18.0/download",
        type = "tar.gz",
        sha256 = "dd8b5dd2ae5ed71462c540258bedcb51965123ad7e7ccf4b9a8cafaa4a63576d",
        strip_prefix = "once_cell-1.18.0",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.once_cell-1.18.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__opener__0_6_1",
        url = "https://crates.io/api/v1/crates/opener/0.6.1/download",
        type = "tar.gz",
        sha256 = "6c62dcb6174f9cb326eac248f07e955d5d559c272730b6c03e396b443b562788",
        strip_prefix = "opener-0.6.1",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.opener-0.6.1.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__parking_lot__0_12_1",
        url = "https://crates.io/api/v1/crates/parking_lot/0.12.1/download",
        type = "tar.gz",
        sha256 = "3742b2c103b9f06bc9fff0a37ff4912935851bee6d36f3c02bcc755bcfec228f",
        strip_prefix = "parking_lot-0.12.1",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.parking_lot-0.12.1.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__parking_lot_core__0_9_8",
        url = "https://crates.io/api/v1/crates/parking_lot_core/0.9.8/download",
        type = "tar.gz",
        sha256 = "93f00c865fe7cabf650081affecd3871070f26767e7b2070a3ffae14c654b447",
        strip_prefix = "parking_lot_core-0.9.8",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.parking_lot_core-0.9.8.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__percent_encoding__2_3_0",
        url = "https://crates.io/api/v1/crates/percent-encoding/2.3.0/download",
        type = "tar.gz",
        sha256 = "9b2a4787296e9989611394c33f193f676704af1686e70b8f8033ab5ba9a35a94",
        strip_prefix = "percent-encoding-2.3.0",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.percent-encoding-2.3.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__pest__2_7_4",
        url = "https://crates.io/api/v1/crates/pest/2.7.4/download",
        type = "tar.gz",
        sha256 = "c022f1e7b65d6a24c0dbbd5fb344c66881bc01f3e5ae74a1c8100f2f985d98a4",
        strip_prefix = "pest-2.7.4",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.pest-2.7.4.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__pest_derive__2_7_4",
        url = "https://crates.io/api/v1/crates/pest_derive/2.7.4/download",
        type = "tar.gz",
        sha256 = "35513f630d46400a977c4cb58f78e1bfbe01434316e60c37d27b9ad6139c66d8",
        strip_prefix = "pest_derive-2.7.4",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.pest_derive-2.7.4.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__pest_generator__2_7_4",
        url = "https://crates.io/api/v1/crates/pest_generator/2.7.4/download",
        type = "tar.gz",
        sha256 = "bc9fc1b9e7057baba189b5c626e2d6f40681ae5b6eb064dc7c7834101ec8123a",
        strip_prefix = "pest_generator-2.7.4",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.pest_generator-2.7.4.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__pest_meta__2_7_4",
        url = "https://crates.io/api/v1/crates/pest_meta/2.7.4/download",
        type = "tar.gz",
        sha256 = "1df74e9e7ec4053ceb980e7c0c8bd3594e977fde1af91daba9c928e8e8c6708d",
        strip_prefix = "pest_meta-2.7.4",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.pest_meta-2.7.4.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__petgraph__0_6_4",
        url = "https://crates.io/api/v1/crates/petgraph/0.6.4/download",
        type = "tar.gz",
        sha256 = "e1d3afd2628e69da2be385eb6f2fd57c8ac7977ceeff6dc166ff1657b0e386a9",
        strip_prefix = "petgraph-0.6.4",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.petgraph-0.6.4.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__phf__0_10_1",
        url = "https://crates.io/api/v1/crates/phf/0.10.1/download",
        type = "tar.gz",
        sha256 = "fabbf1ead8a5bcbc20f5f8b939ee3f5b0f6f281b6ad3468b84656b658b455259",
        strip_prefix = "phf-0.10.1",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.phf-0.10.1.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__phf_codegen__0_10_0",
        url = "https://crates.io/api/v1/crates/phf_codegen/0.10.0/download",
        type = "tar.gz",
        sha256 = "4fb1c3a8bc4dd4e5cfce29b44ffc14bedd2ee294559a294e2a4d4c9e9a6a13cd",
        strip_prefix = "phf_codegen-0.10.0",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.phf_codegen-0.10.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__phf_generator__0_10_0",
        url = "https://crates.io/api/v1/crates/phf_generator/0.10.0/download",
        type = "tar.gz",
        sha256 = "5d5285893bb5eb82e6aaf5d59ee909a06a16737a8970984dd7746ba9283498d6",
        strip_prefix = "phf_generator-0.10.0",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.phf_generator-0.10.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__phf_shared__0_10_0",
        url = "https://crates.io/api/v1/crates/phf_shared/0.10.0/download",
        type = "tar.gz",
        sha256 = "b6796ad771acdc0123d2a88dc428b5e38ef24456743ddb1744ed628f9815c096",
        strip_prefix = "phf_shared-0.10.0",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.phf_shared-0.10.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__pin_project__1_1_3",
        url = "https://crates.io/api/v1/crates/pin-project/1.1.3/download",
        type = "tar.gz",
        sha256 = "fda4ed1c6c173e3fc7a83629421152e01d7b1f9b7f65fb301e490e8cfc656422",
        strip_prefix = "pin-project-1.1.3",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.pin-project-1.1.3.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__pin_project_internal__1_1_3",
        url = "https://crates.io/api/v1/crates/pin-project-internal/1.1.3/download",
        type = "tar.gz",
        sha256 = "4359fd9c9171ec6e8c62926d6faaf553a8dc3f64e1507e76da7911b4f6a04405",
        strip_prefix = "pin-project-internal-1.1.3",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.pin-project-internal-1.1.3.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__pin_project_lite__0_2_13",
        url = "https://crates.io/api/v1/crates/pin-project-lite/0.2.13/download",
        type = "tar.gz",
        sha256 = "8afb450f006bf6385ca15ef45d71d2288452bc3683ce2e2cacc0d18e4be60b58",
        strip_prefix = "pin-project-lite-0.2.13",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.pin-project-lite-0.2.13.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__pin_utils__0_1_0",
        url = "https://crates.io/api/v1/crates/pin-utils/0.1.0/download",
        type = "tar.gz",
        sha256 = "8b870d8c151b6f2fb93e84a13146138f05d02ed11c7e7c54f8826aaaf7c9f184",
        strip_prefix = "pin-utils-0.1.0",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.pin-utils-0.1.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__ppv_lite86__0_2_17",
        url = "https://crates.io/api/v1/crates/ppv-lite86/0.2.17/download",
        type = "tar.gz",
        sha256 = "5b40af805b3121feab8a3c29f04d8ad262fa8e0561883e7653e024ae4479e6de",
        strip_prefix = "ppv-lite86-0.2.17",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.ppv-lite86-0.2.17.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__precomputed_hash__0_1_1",
        url = "https://crates.io/api/v1/crates/precomputed-hash/0.1.1/download",
        type = "tar.gz",
        sha256 = "925383efa346730478fb4838dbe9137d2a47675ad789c546d150a6e1dd4ab31c",
        strip_prefix = "precomputed-hash-0.1.1",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.precomputed-hash-0.1.1.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__proc_macro2__1_0_67",
        url = "https://crates.io/api/v1/crates/proc-macro2/1.0.67/download",
        type = "tar.gz",
        sha256 = "3d433d9f1a3e8c1263d9456598b16fec66f4acc9a74dacffd35c7bb09b3a1328",
        strip_prefix = "proc-macro2-1.0.67",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.proc-macro2-1.0.67.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__prost__0_11_9",
        url = "https://crates.io/api/v1/crates/prost/0.11.9/download",
        type = "tar.gz",
        sha256 = "0b82eaa1d779e9a4bc1c3217db8ffbeabaae1dca241bf70183242128d48681cd",
        strip_prefix = "prost-0.11.9",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.prost-0.11.9.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__prost_build__0_11_9",
        url = "https://crates.io/api/v1/crates/prost-build/0.11.9/download",
        type = "tar.gz",
        sha256 = "119533552c9a7ffacc21e099c24a0ac8bb19c2a2a3f363de84cd9b844feab270",
        strip_prefix = "prost-build-0.11.9",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.prost-build-0.11.9.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__prost_derive__0_11_9",
        url = "https://crates.io/api/v1/crates/prost-derive/0.11.9/download",
        type = "tar.gz",
        sha256 = "e5d2d8d10f3c6ded6da8b05b5fb3b8a5082514344d56c9f871412d29b4e075b4",
        strip_prefix = "prost-derive-0.11.9",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.prost-derive-0.11.9.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__prost_types__0_11_9",
        url = "https://crates.io/api/v1/crates/prost-types/0.11.9/download",
        type = "tar.gz",
        sha256 = "213622a1460818959ac1181aaeb2dc9c7f63df720db7d788b3e24eacd1983e13",
        strip_prefix = "prost-types-0.11.9",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.prost-types-0.11.9.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__protoc_gen_prost__0_2_3",
        url = "https://crates.io/api/v1/crates/protoc-gen-prost/0.2.3/download",
        type = "tar.gz",
        sha256 = "10dfa031ad41fdcfb180de73ece3ed076250f1132a13ad6bba218699f612fb95",
        strip_prefix = "protoc-gen-prost-0.2.3",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.protoc-gen-prost-0.2.3.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__pulldown_cmark__0_9_3",
        url = "https://crates.io/api/v1/crates/pulldown-cmark/0.9.3/download",
        type = "tar.gz",
        sha256 = "77a1a2f1f0a7ecff9c31abbe177637be0e97a0aef46cf8738ece09327985d998",
        strip_prefix = "pulldown-cmark-0.9.3",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.pulldown-cmark-0.9.3.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__quote__1_0_33",
        url = "https://crates.io/api/v1/crates/quote/1.0.33/download",
        type = "tar.gz",
        sha256 = "5267fca4496028628a95160fc423a33e8b2e6af8a5302579e322e4b520293cae",
        strip_prefix = "quote-1.0.33",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.quote-1.0.33.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__rand__0_8_5",
        url = "https://crates.io/api/v1/crates/rand/0.8.5/download",
        type = "tar.gz",
        sha256 = "34af8d1a0e25924bc5b7c43c079c942339d8f0a8b57c39049bef581b46327404",
        strip_prefix = "rand-0.8.5",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.rand-0.8.5.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__rand_chacha__0_3_1",
        url = "https://crates.io/api/v1/crates/rand_chacha/0.3.1/download",
        type = "tar.gz",
        sha256 = "e6c10a63a0fa32252be49d21e7709d4d4baf8d231c2dbce1eaa8141b9b127d88",
        strip_prefix = "rand_chacha-0.3.1",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.rand_chacha-0.3.1.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__rand_core__0_6_4",
        url = "https://crates.io/api/v1/crates/rand_core/0.6.4/download",
        type = "tar.gz",
        sha256 = "ec0be4795e2f6a28069bec0b5ff3e2ac9bafc99e6a9a7dc3547996c5c816922c",
        strip_prefix = "rand_core-0.6.4",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.rand_core-0.6.4.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__redox_syscall__0_3_5",
        url = "https://crates.io/api/v1/crates/redox_syscall/0.3.5/download",
        type = "tar.gz",
        sha256 = "567664f262709473930a4bf9e51bf2ebf3348f2e748ccc50dea20646858f8f29",
        strip_prefix = "redox_syscall-0.3.5",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.redox_syscall-0.3.5.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__regex__1_9_6",
        url = "https://crates.io/api/v1/crates/regex/1.9.6/download",
        type = "tar.gz",
        sha256 = "ebee201405406dbf528b8b672104ae6d6d63e6d118cb10e4d51abbc7b58044ff",
        strip_prefix = "regex-1.9.6",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.regex-1.9.6.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__regex_automata__0_3_9",
        url = "https://crates.io/api/v1/crates/regex-automata/0.3.9/download",
        type = "tar.gz",
        sha256 = "59b23e92ee4318893fa3fe3e6fb365258efbfe6ac6ab30f090cdcbb7aa37efa9",
        strip_prefix = "regex-automata-0.3.9",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.regex-automata-0.3.9.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__regex_syntax__0_7_5",
        url = "https://crates.io/api/v1/crates/regex-syntax/0.7.5/download",
        type = "tar.gz",
        sha256 = "dbb5fb1acd8a1a18b3dd5be62d25485eb770e05afb408a9627d14d451bae12da",
        strip_prefix = "regex-syntax-0.7.5",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.regex-syntax-0.7.5.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__rustc_demangle__0_1_23",
        url = "https://crates.io/api/v1/crates/rustc-demangle/0.1.23/download",
        type = "tar.gz",
        sha256 = "d626bb9dae77e28219937af045c257c28bfd3f69333c512553507f5f9798cb76",
        strip_prefix = "rustc-demangle-0.1.23",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.rustc-demangle-0.1.23.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__rustix__0_38_15",
        url = "https://crates.io/api/v1/crates/rustix/0.38.15/download",
        type = "tar.gz",
        sha256 = "d2f9da0cbd88f9f09e7814e388301c8414c51c62aa6ce1e4b5c551d49d96e531",
        strip_prefix = "rustix-0.38.15",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.rustix-0.38.15.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__rustls_pemfile__1_0_3",
        url = "https://crates.io/api/v1/crates/rustls-pemfile/1.0.3/download",
        type = "tar.gz",
        sha256 = "2d3987094b1d07b653b7dfdc3f70ce9a1da9c51ac18c1b06b662e4f9a0e9f4b2",
        strip_prefix = "rustls-pemfile-1.0.3",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.rustls-pemfile-1.0.3.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__ryu__1_0_15",
        url = "https://crates.io/api/v1/crates/ryu/1.0.15/download",
        type = "tar.gz",
        sha256 = "1ad4cc8da4ef723ed60bced201181d83791ad433213d8c24efffda1eec85d741",
        strip_prefix = "ryu-1.0.15",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.ryu-1.0.15.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__same_file__1_0_6",
        url = "https://crates.io/api/v1/crates/same-file/1.0.6/download",
        type = "tar.gz",
        sha256 = "93fc1dc3aaa9bfed95e02e6eadabb4baf7e3078b0bd1b4d7b6b0b68378900502",
        strip_prefix = "same-file-1.0.6",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.same-file-1.0.6.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__scoped_tls__1_0_1",
        url = "https://crates.io/api/v1/crates/scoped-tls/1.0.1/download",
        type = "tar.gz",
        sha256 = "e1cf6437eb19a8f4a6cc0f7dca544973b0b78843adbfeb3683d1a94a0024a294",
        strip_prefix = "scoped-tls-1.0.1",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.scoped-tls-1.0.1.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__scopeguard__1_2_0",
        url = "https://crates.io/api/v1/crates/scopeguard/1.2.0/download",
        type = "tar.gz",
        sha256 = "94143f37725109f92c262ed2cf5e59bce7498c01bcc1502d7b9afe439a4e9f49",
        strip_prefix = "scopeguard-1.2.0",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.scopeguard-1.2.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__serde__1_0_188",
        url = "https://crates.io/api/v1/crates/serde/1.0.188/download",
        type = "tar.gz",
        sha256 = "cf9e0fcba69a370eed61bcf2b728575f726b50b55cba78064753d708ddc7549e",
        strip_prefix = "serde-1.0.188",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.serde-1.0.188.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__serde_derive__1_0_188",
        url = "https://crates.io/api/v1/crates/serde_derive/1.0.188/download",
        type = "tar.gz",
        sha256 = "4eca7ac642d82aa35b60049a6eccb4be6be75e599bd2e9adb5f875a737654af2",
        strip_prefix = "serde_derive-1.0.188",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.serde_derive-1.0.188.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__serde_json__1_0_107",
        url = "https://crates.io/api/v1/crates/serde_json/1.0.107/download",
        type = "tar.gz",
        sha256 = "6b420ce6e3d8bd882e9b243c6eed35dbc9a6110c9769e74b584e0d68d1f20c65",
        strip_prefix = "serde_json-1.0.107",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.serde_json-1.0.107.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__serde_urlencoded__0_7_1",
        url = "https://crates.io/api/v1/crates/serde_urlencoded/0.7.1/download",
        type = "tar.gz",
        sha256 = "d3491c14715ca2294c4d6a88f15e84739788c1d030eed8c110436aafdaa2f3fd",
        strip_prefix = "serde_urlencoded-0.7.1",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.serde_urlencoded-0.7.1.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__sha1__0_10_6",
        url = "https://crates.io/api/v1/crates/sha1/0.10.6/download",
        type = "tar.gz",
        sha256 = "e3bf829a2d51ab4a5ddf1352d8470c140cadc8301b2ae1789db023f01cedd6ba",
        strip_prefix = "sha1-0.10.6",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.sha1-0.10.6.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__sha2__0_10_8",
        url = "https://crates.io/api/v1/crates/sha2/0.10.8/download",
        type = "tar.gz",
        sha256 = "793db75ad2bcafc3ffa7c68b215fee268f537982cd901d132f89c6343f3a3dc8",
        strip_prefix = "sha2-0.10.8",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.sha2-0.10.8.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__shlex__1_2_0",
        url = "https://crates.io/api/v1/crates/shlex/1.2.0/download",
        type = "tar.gz",
        sha256 = "a7cee0529a6d40f580e7a5e6c495c8fbfe21b7b52795ed4bb5e62cdf92bc6380",
        strip_prefix = "shlex-1.2.0",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.shlex-1.2.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__siphasher__0_3_11",
        url = "https://crates.io/api/v1/crates/siphasher/0.3.11/download",
        type = "tar.gz",
        sha256 = "38b58827f4464d87d377d175e90bf58eb00fd8716ff0a62f80356b5e61555d0d",
        strip_prefix = "siphasher-0.3.11",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.siphasher-0.3.11.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__slab__0_4_9",
        url = "https://crates.io/api/v1/crates/slab/0.4.9/download",
        type = "tar.gz",
        sha256 = "8f92a496fb766b417c996b9c5e57daf2f7ad3b0bebe1ccfca4856390e3d3bb67",
        strip_prefix = "slab-0.4.9",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.slab-0.4.9.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__smallvec__1_11_1",
        url = "https://crates.io/api/v1/crates/smallvec/1.11.1/download",
        type = "tar.gz",
        sha256 = "942b4a808e05215192e39f4ab80813e599068285906cc91aa64f923db842bd5a",
        strip_prefix = "smallvec-1.11.1",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.smallvec-1.11.1.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__socket2__0_4_9",
        url = "https://crates.io/api/v1/crates/socket2/0.4.9/download",
        type = "tar.gz",
        sha256 = "64a4a911eed85daf18834cfaa86a79b7d266ff93ff5ba14005426219480ed662",
        strip_prefix = "socket2-0.4.9",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.socket2-0.4.9.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__socket2__0_5_4",
        url = "https://crates.io/api/v1/crates/socket2/0.5.4/download",
        type = "tar.gz",
        sha256 = "4031e820eb552adee9295814c0ced9e5cf38ddf1e8b7d566d6de8e2538ea989e",
        strip_prefix = "socket2-0.5.4",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.socket2-0.5.4.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__string_cache__0_8_7",
        url = "https://crates.io/api/v1/crates/string_cache/0.8.7/download",
        type = "tar.gz",
        sha256 = "f91138e76242f575eb1d3b38b4f1362f10d3a43f47d182a5b359af488a02293b",
        strip_prefix = "string_cache-0.8.7",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.string_cache-0.8.7.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__string_cache_codegen__0_5_2",
        url = "https://crates.io/api/v1/crates/string_cache_codegen/0.5.2/download",
        type = "tar.gz",
        sha256 = "6bb30289b722be4ff74a408c3cc27edeaad656e06cb1fe8fa9231fa59c728988",
        strip_prefix = "string_cache_codegen-0.5.2",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.string_cache_codegen-0.5.2.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__strsim__0_10_0",
        url = "https://crates.io/api/v1/crates/strsim/0.10.0/download",
        type = "tar.gz",
        sha256 = "73473c0e59e6d5812c5dfe2a064a6444949f089e20eec9a2e5506596494e4623",
        strip_prefix = "strsim-0.10.0",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.strsim-0.10.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__syn__1_0_109",
        url = "https://crates.io/api/v1/crates/syn/1.0.109/download",
        type = "tar.gz",
        sha256 = "72b64191b275b66ffe2469e8af2c1cfe3bafa67b529ead792a6d0160888b4237",
        strip_prefix = "syn-1.0.109",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.syn-1.0.109.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__syn__2_0_37",
        url = "https://crates.io/api/v1/crates/syn/2.0.37/download",
        type = "tar.gz",
        sha256 = "7303ef2c05cd654186cb250d29049a24840ca25d2747c25c0381c8d9e2f582e8",
        strip_prefix = "syn-2.0.37",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.syn-2.0.37.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__tempfile__3_8_0",
        url = "https://crates.io/api/v1/crates/tempfile/3.8.0/download",
        type = "tar.gz",
        sha256 = "cb94d2f3cc536af71caac6b6fcebf65860b347e7ce0cc9ebe8f70d3e521054ef",
        strip_prefix = "tempfile-3.8.0",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.tempfile-3.8.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__tendril__0_4_3",
        url = "https://crates.io/api/v1/crates/tendril/0.4.3/download",
        type = "tar.gz",
        sha256 = "d24a120c5fc464a3458240ee02c299ebcb9d67b5249c8848b09d639dca8d7bb0",
        strip_prefix = "tendril-0.4.3",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.tendril-0.4.3.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__termcolor__1_3_0",
        url = "https://crates.io/api/v1/crates/termcolor/1.3.0/download",
        type = "tar.gz",
        sha256 = "6093bad37da69aab9d123a8091e4be0aa4a03e4d601ec641c327398315f62b64",
        strip_prefix = "termcolor-1.3.0",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.termcolor-1.3.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__terminal_size__0_3_0",
        url = "https://crates.io/api/v1/crates/terminal_size/0.3.0/download",
        type = "tar.gz",
        sha256 = "21bebf2b7c9e0a515f6e0f8c51dc0f8e4696391e6f1ff30379559f8365fb0df7",
        strip_prefix = "terminal_size-0.3.0",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.terminal_size-0.3.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__thiserror__1_0_49",
        url = "https://crates.io/api/v1/crates/thiserror/1.0.49/download",
        type = "tar.gz",
        sha256 = "1177e8c6d7ede7afde3585fd2513e611227efd6481bd78d2e82ba1ce16557ed4",
        strip_prefix = "thiserror-1.0.49",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.thiserror-1.0.49.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__thiserror_impl__1_0_49",
        url = "https://crates.io/api/v1/crates/thiserror-impl/1.0.49/download",
        type = "tar.gz",
        sha256 = "10712f02019e9288794769fba95cd6847df9874d49d871d062172f9dd41bc4cc",
        strip_prefix = "thiserror-impl-1.0.49",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.thiserror-impl-1.0.49.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__thread_local__1_1_7",
        url = "https://crates.io/api/v1/crates/thread_local/1.1.7/download",
        type = "tar.gz",
        sha256 = "3fdd6f064ccff2d6567adcb3873ca630700f00b5ad3f060c25b5dcfd9a4ce152",
        strip_prefix = "thread_local-1.1.7",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.thread_local-1.1.7.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__tinyvec__1_6_0",
        url = "https://crates.io/api/v1/crates/tinyvec/1.6.0/download",
        type = "tar.gz",
        sha256 = "87cc5ceb3875bb20c2890005a4e226a4651264a5c75edb2421b52861a0a0cb50",
        strip_prefix = "tinyvec-1.6.0",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.tinyvec-1.6.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__tinyvec_macros__0_1_1",
        url = "https://crates.io/api/v1/crates/tinyvec_macros/0.1.1/download",
        type = "tar.gz",
        sha256 = "1f3ccbac311fea05f86f61904b462b55fb3df8837a366dfc601a0161d0532f20",
        strip_prefix = "tinyvec_macros-0.1.1",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.tinyvec_macros-0.1.1.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__tokio__1_32_0",
        url = "https://crates.io/api/v1/crates/tokio/1.32.0/download",
        type = "tar.gz",
        sha256 = "17ed6077ed6cd6c74735e21f37eb16dc3935f96878b1fe961074089cc80893f9",
        strip_prefix = "tokio-1.32.0",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.tokio-1.32.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__tokio_macros__2_1_0",
        url = "https://crates.io/api/v1/crates/tokio-macros/2.1.0/download",
        type = "tar.gz",
        sha256 = "630bdcf245f78637c13ec01ffae6187cca34625e8c63150d424b59e55af2675e",
        strip_prefix = "tokio-macros-2.1.0",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.tokio-macros-2.1.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__tokio_stream__0_1_14",
        url = "https://crates.io/api/v1/crates/tokio-stream/0.1.14/download",
        type = "tar.gz",
        sha256 = "397c988d37662c7dda6d2208364a706264bf3d6138b11d436cbac0ad38832842",
        strip_prefix = "tokio-stream-0.1.14",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.tokio-stream-0.1.14.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__tokio_tungstenite__0_20_1",
        url = "https://crates.io/api/v1/crates/tokio-tungstenite/0.20.1/download",
        type = "tar.gz",
        sha256 = "212d5dcb2a1ce06d81107c3d0ffa3121fe974b73f068c8282cb1c32328113b6c",
        strip_prefix = "tokio-tungstenite-0.20.1",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.tokio-tungstenite-0.20.1.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__tokio_util__0_7_9",
        url = "https://crates.io/api/v1/crates/tokio-util/0.7.9/download",
        type = "tar.gz",
        sha256 = "1d68074620f57a0b21594d9735eb2e98ab38b17f80d3fcb189fca266771ca60d",
        strip_prefix = "tokio-util-0.7.9",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.tokio-util-0.7.9.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__toml__0_5_11",
        url = "https://crates.io/api/v1/crates/toml/0.5.11/download",
        type = "tar.gz",
        sha256 = "f4f7f0dd8d50a853a531c426359045b1998f04219d88799810762cd4ad314234",
        strip_prefix = "toml-0.5.11",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.toml-0.5.11.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__topological_sort__0_2_2",
        url = "https://crates.io/api/v1/crates/topological-sort/0.2.2/download",
        type = "tar.gz",
        sha256 = "ea68304e134ecd095ac6c3574494fc62b909f416c4fca77e440530221e549d3d",
        strip_prefix = "topological-sort-0.2.2",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.topological-sort-0.2.2.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__tower_service__0_3_2",
        url = "https://crates.io/api/v1/crates/tower-service/0.3.2/download",
        type = "tar.gz",
        sha256 = "b6bc1c9ce2b5135ac7f93c72918fc37feb872bdc6a5533a8b85eb4b86bfdae52",
        strip_prefix = "tower-service-0.3.2",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.tower-service-0.3.2.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__tracing__0_1_37",
        url = "https://crates.io/api/v1/crates/tracing/0.1.37/download",
        type = "tar.gz",
        sha256 = "8ce8c33a8d48bd45d624a6e523445fd21ec13d3653cd51f681abf67418f54eb8",
        strip_prefix = "tracing-0.1.37",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.tracing-0.1.37.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__tracing_core__0_1_31",
        url = "https://crates.io/api/v1/crates/tracing-core/0.1.31/download",
        type = "tar.gz",
        sha256 = "0955b8137a1df6f1a2e9a37d8a6656291ff0297c1a97c24e0d8425fe2312f79a",
        strip_prefix = "tracing-core-0.1.31",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.tracing-core-0.1.31.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__try_lock__0_2_4",
        url = "https://crates.io/api/v1/crates/try-lock/0.2.4/download",
        type = "tar.gz",
        sha256 = "3528ecfd12c466c6f163363caf2d02a71161dd5e1cc6ae7b34207ea2d42d81ed",
        strip_prefix = "try-lock-0.2.4",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.try-lock-0.2.4.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__tungstenite__0_20_1",
        url = "https://crates.io/api/v1/crates/tungstenite/0.20.1/download",
        type = "tar.gz",
        sha256 = "9e3dac10fd62eaf6617d3a904ae222845979aec67c615d1c842b4002c7666fb9",
        strip_prefix = "tungstenite-0.20.1",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.tungstenite-0.20.1.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__typenum__1_17_0",
        url = "https://crates.io/api/v1/crates/typenum/1.17.0/download",
        type = "tar.gz",
        sha256 = "42ff0bf0c66b8238c6f3b578df37d0b7848e55df8577b3f74f92a69acceeb825",
        strip_prefix = "typenum-1.17.0",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.typenum-1.17.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__ucd_trie__0_1_6",
        url = "https://crates.io/api/v1/crates/ucd-trie/0.1.6/download",
        type = "tar.gz",
        sha256 = "ed646292ffc8188ef8ea4d1e0e0150fb15a5c2e12ad9b8fc191ae7a8a7f3c4b9",
        strip_prefix = "ucd-trie-0.1.6",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.ucd-trie-0.1.6.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__unicase__2_7_0",
        url = "https://crates.io/api/v1/crates/unicase/2.7.0/download",
        type = "tar.gz",
        sha256 = "f7d2d4dafb69621809a81864c9c1b864479e1235c0dd4e199924b9742439ed89",
        strip_prefix = "unicase-2.7.0",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.unicase-2.7.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__unicode_bidi__0_3_13",
        url = "https://crates.io/api/v1/crates/unicode-bidi/0.3.13/download",
        type = "tar.gz",
        sha256 = "92888ba5573ff080736b3648696b70cafad7d250551175acbaa4e0385b3e1460",
        strip_prefix = "unicode-bidi-0.3.13",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.unicode-bidi-0.3.13.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__unicode_ident__1_0_12",
        url = "https://crates.io/api/v1/crates/unicode-ident/1.0.12/download",
        type = "tar.gz",
        sha256 = "3354b9ac3fae1ff6755cb6db53683adb661634f67557942dea4facebec0fee4b",
        strip_prefix = "unicode-ident-1.0.12",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.unicode-ident-1.0.12.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__unicode_normalization__0_1_22",
        url = "https://crates.io/api/v1/crates/unicode-normalization/0.1.22/download",
        type = "tar.gz",
        sha256 = "5c5713f0fc4b5db668a2ac63cdb7bb4469d8c9fed047b1d0292cc7b0ce2ba921",
        strip_prefix = "unicode-normalization-0.1.22",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.unicode-normalization-0.1.22.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__url__2_4_1",
        url = "https://crates.io/api/v1/crates/url/2.4.1/download",
        type = "tar.gz",
        sha256 = "143b538f18257fac9cad154828a57c6bf5157e1aa604d4816b5995bf6de87ae5",
        strip_prefix = "url-2.4.1",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.url-2.4.1.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__utf_8__0_7_6",
        url = "https://crates.io/api/v1/crates/utf-8/0.7.6/download",
        type = "tar.gz",
        sha256 = "09cc8ee72d2a9becf2f2febe0205bbed8fc6615b7cb429ad062dc7b7ddd036a9",
        strip_prefix = "utf-8-0.7.6",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.utf-8-0.7.6.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__utf8parse__0_2_1",
        url = "https://crates.io/api/v1/crates/utf8parse/0.2.1/download",
        type = "tar.gz",
        sha256 = "711b9620af191e0cdc7468a8d14e709c3dcdb115b36f838e601583af800a370a",
        strip_prefix = "utf8parse-0.2.1",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.utf8parse-0.2.1.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__version_check__0_9_4",
        url = "https://crates.io/api/v1/crates/version_check/0.9.4/download",
        type = "tar.gz",
        sha256 = "49874b5167b65d7193b8aba1567f5c7d93d001cafc34600cee003eda787e483f",
        strip_prefix = "version_check-0.9.4",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.version_check-0.9.4.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__walkdir__2_4_0",
        url = "https://crates.io/api/v1/crates/walkdir/2.4.0/download",
        type = "tar.gz",
        sha256 = "d71d857dc86794ca4c280d616f7da00d2dbfd8cd788846559a6813e6aa4b54ee",
        strip_prefix = "walkdir-2.4.0",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.walkdir-2.4.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__want__0_3_1",
        url = "https://crates.io/api/v1/crates/want/0.3.1/download",
        type = "tar.gz",
        sha256 = "bfa7760aed19e106de2c7c0b581b509f2f25d3dacaf737cb82ac61bc6d760b0e",
        strip_prefix = "want-0.3.1",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.want-0.3.1.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__warp__0_3_6",
        url = "https://crates.io/api/v1/crates/warp/0.3.6/download",
        type = "tar.gz",
        sha256 = "c1e92e22e03ff1230c03a1a8ee37d2f89cd489e2e541b7550d6afad96faed169",
        strip_prefix = "warp-0.3.6",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.warp-0.3.6.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__wasi__0_11_0_wasi_snapshot_preview1",
        url = "https://crates.io/api/v1/crates/wasi/0.11.0+wasi-snapshot-preview1/download",
        type = "tar.gz",
        sha256 = "9c8d87e72b64a3b4db28d11ce29237c246188f4f51057d65a7eab63b7987e423",
        strip_prefix = "wasi-0.11.0+wasi-snapshot-preview1",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.wasi-0.11.0+wasi-snapshot-preview1.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__wasm_bindgen__0_2_87",
        url = "https://crates.io/api/v1/crates/wasm-bindgen/0.2.87/download",
        type = "tar.gz",
        sha256 = "7706a72ab36d8cb1f80ffbf0e071533974a60d0a308d01a5d0375bf60499a342",
        strip_prefix = "wasm-bindgen-0.2.87",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.wasm-bindgen-0.2.87.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__wasm_bindgen_backend__0_2_87",
        url = "https://crates.io/api/v1/crates/wasm-bindgen-backend/0.2.87/download",
        type = "tar.gz",
        sha256 = "5ef2b6d3c510e9625e5fe6f509ab07d66a760f0885d858736483c32ed7809abd",
        strip_prefix = "wasm-bindgen-backend-0.2.87",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.wasm-bindgen-backend-0.2.87.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__wasm_bindgen_macro__0_2_87",
        url = "https://crates.io/api/v1/crates/wasm-bindgen-macro/0.2.87/download",
        type = "tar.gz",
        sha256 = "dee495e55982a3bd48105a7b947fd2a9b4a8ae3010041b9e0faab3f9cd028f1d",
        strip_prefix = "wasm-bindgen-macro-0.2.87",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.wasm-bindgen-macro-0.2.87.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__wasm_bindgen_macro_support__0_2_87",
        url = "https://crates.io/api/v1/crates/wasm-bindgen-macro-support/0.2.87/download",
        type = "tar.gz",
        sha256 = "54681b18a46765f095758388f2d0cf16eb8d4169b639ab575a8f5693af210c7b",
        strip_prefix = "wasm-bindgen-macro-support-0.2.87",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.wasm-bindgen-macro-support-0.2.87.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__wasm_bindgen_shared__0_2_87",
        url = "https://crates.io/api/v1/crates/wasm-bindgen-shared/0.2.87/download",
        type = "tar.gz",
        sha256 = "ca6ad05a4870b2bf5fe995117d3728437bd27d7cd5f06f13c17443ef369775a1",
        strip_prefix = "wasm-bindgen-shared-0.2.87",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.wasm-bindgen-shared-0.2.87.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__which__4_4_2",
        url = "https://crates.io/api/v1/crates/which/4.4.2/download",
        type = "tar.gz",
        sha256 = "87ba24419a2078cd2b0f2ede2691b6c66d8e47836da3b6db8265ebad47afbfc7",
        strip_prefix = "which-4.4.2",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.which-4.4.2.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__winapi__0_3_9",
        url = "https://crates.io/api/v1/crates/winapi/0.3.9/download",
        type = "tar.gz",
        sha256 = "5c839a674fcd7a98952e593242ea400abe93992746761e38641405d28b00f419",
        strip_prefix = "winapi-0.3.9",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.winapi-0.3.9.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__winapi_i686_pc_windows_gnu__0_4_0",
        url = "https://crates.io/api/v1/crates/winapi-i686-pc-windows-gnu/0.4.0/download",
        type = "tar.gz",
        sha256 = "ac3b87c63620426dd9b991e5ce0329eff545bccbbb34f3be09ff6fb6ab51b7b6",
        strip_prefix = "winapi-i686-pc-windows-gnu-0.4.0",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.winapi-i686-pc-windows-gnu-0.4.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__winapi_util__0_1_6",
        url = "https://crates.io/api/v1/crates/winapi-util/0.1.6/download",
        type = "tar.gz",
        sha256 = "f29e6f9198ba0d26b4c9f07dbe6f9ed633e1f3d5b8b414090084349e46a52596",
        strip_prefix = "winapi-util-0.1.6",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.winapi-util-0.1.6.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__winapi_x86_64_pc_windows_gnu__0_4_0",
        url = "https://crates.io/api/v1/crates/winapi-x86_64-pc-windows-gnu/0.4.0/download",
        type = "tar.gz",
        sha256 = "712e227841d057c1ee1cd2fb22fa7e5a5461ae8e48fa2ca79ec42cfc1931183f",
        strip_prefix = "winapi-x86_64-pc-windows-gnu-0.4.0",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.winapi-x86_64-pc-windows-gnu-0.4.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__windows__0_48_0",
        url = "https://crates.io/api/v1/crates/windows/0.48.0/download",
        type = "tar.gz",
        sha256 = "e686886bc078bc1b0b600cac0147aadb815089b6e4da64016cbd754b6342700f",
        strip_prefix = "windows-0.48.0",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.windows-0.48.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__windows_sys__0_48_0",
        url = "https://crates.io/api/v1/crates/windows-sys/0.48.0/download",
        type = "tar.gz",
        sha256 = "677d2418bec65e3338edb076e806bc1ec15693c5d0104683f2efe857f61056a9",
        strip_prefix = "windows-sys-0.48.0",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.windows-sys-0.48.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__windows_targets__0_48_5",
        url = "https://crates.io/api/v1/crates/windows-targets/0.48.5/download",
        type = "tar.gz",
        sha256 = "9a2fa6e2155d7247be68c096456083145c183cbbbc2764150dda45a87197940c",
        strip_prefix = "windows-targets-0.48.5",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.windows-targets-0.48.5.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__windows_aarch64_gnullvm__0_48_5",
        url = "https://crates.io/api/v1/crates/windows_aarch64_gnullvm/0.48.5/download",
        type = "tar.gz",
        sha256 = "2b38e32f0abccf9987a4e3079dfb67dcd799fb61361e53e2882c3cbaf0d905d8",
        strip_prefix = "windows_aarch64_gnullvm-0.48.5",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.windows_aarch64_gnullvm-0.48.5.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__windows_aarch64_msvc__0_48_5",
        url = "https://crates.io/api/v1/crates/windows_aarch64_msvc/0.48.5/download",
        type = "tar.gz",
        sha256 = "dc35310971f3b2dbbf3f0690a219f40e2d9afcf64f9ab7cc1be722937c26b4bc",
        strip_prefix = "windows_aarch64_msvc-0.48.5",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.windows_aarch64_msvc-0.48.5.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__windows_i686_gnu__0_48_5",
        url = "https://crates.io/api/v1/crates/windows_i686_gnu/0.48.5/download",
        type = "tar.gz",
        sha256 = "a75915e7def60c94dcef72200b9a8e58e5091744960da64ec734a6c6e9b3743e",
        strip_prefix = "windows_i686_gnu-0.48.5",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.windows_i686_gnu-0.48.5.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__windows_i686_msvc__0_48_5",
        url = "https://crates.io/api/v1/crates/windows_i686_msvc/0.48.5/download",
        type = "tar.gz",
        sha256 = "8f55c233f70c4b27f66c523580f78f1004e8b5a8b659e05a4eb49d4166cca406",
        strip_prefix = "windows_i686_msvc-0.48.5",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.windows_i686_msvc-0.48.5.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__windows_x86_64_gnu__0_48_5",
        url = "https://crates.io/api/v1/crates/windows_x86_64_gnu/0.48.5/download",
        type = "tar.gz",
        sha256 = "53d40abd2583d23e4718fddf1ebec84dbff8381c07cae67ff7768bbf19c6718e",
        strip_prefix = "windows_x86_64_gnu-0.48.5",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.windows_x86_64_gnu-0.48.5.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__windows_x86_64_gnullvm__0_48_5",
        url = "https://crates.io/api/v1/crates/windows_x86_64_gnullvm/0.48.5/download",
        type = "tar.gz",
        sha256 = "0b7b52767868a23d5bab768e390dc5f5c55825b6d30b86c844ff2dc7414044cc",
        strip_prefix = "windows_x86_64_gnullvm-0.48.5",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.windows_x86_64_gnullvm-0.48.5.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__windows_x86_64_msvc__0_48_5",
        url = "https://crates.io/api/v1/crates/windows_x86_64_msvc/0.48.5/download",
        type = "tar.gz",
        sha256 = "ed94fce61571a4006852b7389a063ab983c02eb1bb37b47f8272ce92d06d9538",
        strip_prefix = "windows_x86_64_msvc-0.48.5",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.windows_x86_64_msvc-0.48.5.bazel"),
    )
