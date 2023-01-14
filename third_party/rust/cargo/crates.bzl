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
        name = "raze__aho_corasick__0_7_20",
        url = "https://crates.io/api/v1/crates/aho-corasick/0.7.20/download",
        type = "tar.gz",
        sha256 = "cc936419f96fa211c1b9166887b38e5e40b19958e5b895be7c1f93adec7071ac",
        strip_prefix = "aho-corasick-0.7.20",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.aho-corasick-0.7.20.bazel"),
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
        name = "raze__android_system_properties__0_1_5",
        url = "https://crates.io/api/v1/crates/android_system_properties/0.1.5/download",
        type = "tar.gz",
        sha256 = "819e7219dbd41043ac279b19830f2efc897156490d7fd6ea916720117ee66311",
        strip_prefix = "android_system_properties-0.1.5",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.android_system_properties-0.1.5.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__anyhow__1_0_68",
        url = "https://crates.io/api/v1/crates/anyhow/1.0.68/download",
        type = "tar.gz",
        sha256 = "2cb2f989d18dd141ab8ae82f64d1a8cdd37e0840f73a406896cf5e99502fab61",
        strip_prefix = "anyhow-1.0.68",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.anyhow-1.0.68.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__atty__0_2_14",
        url = "https://crates.io/api/v1/crates/atty/0.2.14/download",
        type = "tar.gz",
        sha256 = "d9b39be18770d11421cdb1b9947a45dd3f37e93092cbf377614828a319d5fee8",
        strip_prefix = "atty-0.2.14",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.atty-0.2.14.bazel"),
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
        name = "raze__base64__0_13_1",
        url = "https://crates.io/api/v1/crates/base64/0.13.1/download",
        type = "tar.gz",
        sha256 = "9e1b586273c5702936fe7b7d6896644d8be71e6314cfe09d3167c95f712589e8",
        strip_prefix = "base64-0.13.1",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.base64-0.13.1.bazel"),
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
        name = "raze__block_buffer__0_10_3",
        url = "https://crates.io/api/v1/crates/block-buffer/0.10.3/download",
        type = "tar.gz",
        sha256 = "69cce20737498f97b993470a6e536b8523f0af7892a4f928cceb1ac5e52ebe7e",
        strip_prefix = "block-buffer-0.10.3",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.block-buffer-0.10.3.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__bstr__0_2_17",
        url = "https://crates.io/api/v1/crates/bstr/0.2.17/download",
        type = "tar.gz",
        sha256 = "ba3569f383e8f1598449f1a423e72e99569137b47740b1da11ef19af3d5c3223",
        strip_prefix = "bstr-0.2.17",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.bstr-0.2.17.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__bumpalo__3_11_1",
        url = "https://crates.io/api/v1/crates/bumpalo/3.11.1/download",
        type = "tar.gz",
        sha256 = "572f695136211188308f16ad2ca5c851a712c464060ae6974944458eb83880ba",
        strip_prefix = "bumpalo-3.11.1",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.bumpalo-3.11.1.bazel"),
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
        name = "raze__bytes__1_3_0",
        url = "https://crates.io/api/v1/crates/bytes/1.3.0/download",
        type = "tar.gz",
        sha256 = "dfb24e866b15a1af2a1b663f10c6b6b8f397a84aadb828f12e5b289ec23a3a3c",
        strip_prefix = "bytes-1.3.0",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.bytes-1.3.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__cc__1_0_78",
        url = "https://crates.io/api/v1/crates/cc/1.0.78/download",
        type = "tar.gz",
        sha256 = "a20104e2335ce8a659d6dd92a51a767a0c062599c73b343fd152cb401e828c3d",
        strip_prefix = "cc-1.0.78",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.cc-1.0.78.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__cfg_if__0_1_10",
        url = "https://crates.io/api/v1/crates/cfg-if/0.1.10/download",
        type = "tar.gz",
        sha256 = "4785bdd1c96b2a846b2bd7cc02e86b6b3dbf14e7e53446c4f54c92a361040822",
        strip_prefix = "cfg-if-0.1.10",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.cfg-if-0.1.10.bazel"),
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
        name = "raze__chrono__0_4_23",
        url = "https://crates.io/api/v1/crates/chrono/0.4.23/download",
        type = "tar.gz",
        sha256 = "16b0a3d9ed01224b22057780a37bb8c5dbfe1be8ba48678e7bf57ec4b385411f",
        strip_prefix = "chrono-0.4.23",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.chrono-0.4.23.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__clap__3_2_23",
        url = "https://crates.io/api/v1/crates/clap/3.2.23/download",
        type = "tar.gz",
        sha256 = "71655c45cb9845d3270c9d6df84ebe72b4dad3c2ba3f7023ad47c144e4e473a5",
        strip_prefix = "clap-3.2.23",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.clap-3.2.23.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__clap_complete__3_2_5",
        url = "https://crates.io/api/v1/crates/clap_complete/3.2.5/download",
        type = "tar.gz",
        sha256 = "3f7a2e0a962c45ce25afce14220bc24f9dade0a1787f185cecf96bfba7847cd8",
        strip_prefix = "clap_complete-3.2.5",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.clap_complete-3.2.5.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__clap_lex__0_2_4",
        url = "https://crates.io/api/v1/crates/clap_lex/0.2.4/download",
        type = "tar.gz",
        sha256 = "2850f2f5a82cbf437dd5af4d49848fbdfc27c157c3d010345776f952765261c5",
        strip_prefix = "clap_lex-0.2.4",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.clap_lex-0.2.4.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__codespan_reporting__0_11_1",
        url = "https://crates.io/api/v1/crates/codespan-reporting/0.11.1/download",
        type = "tar.gz",
        sha256 = "3538270d33cc669650c4b093848450d380def10c331d38c768e34cac80576e6e",
        strip_prefix = "codespan-reporting-0.11.1",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.codespan-reporting-0.11.1.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__core_foundation_sys__0_8_3",
        url = "https://crates.io/api/v1/crates/core-foundation-sys/0.8.3/download",
        type = "tar.gz",
        sha256 = "5827cebf4670468b8772dd191856768aedcb1b0278a04f989f7766351917b9dc",
        strip_prefix = "core-foundation-sys-0.8.3",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.core-foundation-sys-0.8.3.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__cpufeatures__0_2_5",
        url = "https://crates.io/api/v1/crates/cpufeatures/0.2.5/download",
        type = "tar.gz",
        sha256 = "28d997bd5e24a5928dd43e46dc529867e207907fe0b239c3477d924f7f2ca320",
        strip_prefix = "cpufeatures-0.2.5",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.cpufeatures-0.2.5.bazel"),
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
        name = "raze__cxx__1_0_86",
        url = "https://crates.io/api/v1/crates/cxx/1.0.86/download",
        type = "tar.gz",
        sha256 = "51d1075c37807dcf850c379432f0df05ba52cc30f279c5cfc43cc221ce7f8579",
        strip_prefix = "cxx-1.0.86",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.cxx-1.0.86.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__cxx_build__1_0_86",
        url = "https://crates.io/api/v1/crates/cxx-build/1.0.86/download",
        type = "tar.gz",
        sha256 = "5044281f61b27bc598f2f6647d480aed48d2bf52d6eb0b627d84c0361b17aa70",
        strip_prefix = "cxx-build-1.0.86",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.cxx-build-1.0.86.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__cxxbridge_flags__1_0_86",
        url = "https://crates.io/api/v1/crates/cxxbridge-flags/1.0.86/download",
        type = "tar.gz",
        sha256 = "61b50bc93ba22c27b0d31128d2d130a0a6b3d267ae27ef7e4fae2167dfe8781c",
        strip_prefix = "cxxbridge-flags-1.0.86",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.cxxbridge-flags-1.0.86.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__cxxbridge_macro__1_0_86",
        url = "https://crates.io/api/v1/crates/cxxbridge-macro/1.0.86/download",
        type = "tar.gz",
        sha256 = "39e61fda7e62115119469c7b3591fd913ecca96fb766cfd3f2e2502ab7bc87a5",
        strip_prefix = "cxxbridge-macro-1.0.86",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.cxxbridge-macro-1.0.86.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__digest__0_10_6",
        url = "https://crates.io/api/v1/crates/digest/0.10.6/download",
        type = "tar.gz",
        sha256 = "8168378f4e5023e7218c89c891c0fd8ecdb5e5e4f18cb78f38cf245dd021e76f",
        strip_prefix = "digest-0.10.6",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.digest-0.10.6.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__elasticlunr_rs__3_0_1",
        url = "https://crates.io/api/v1/crates/elasticlunr-rs/3.0.1/download",
        type = "tar.gz",
        sha256 = "b94d9c8df0fe6879ca12e7633fdfe467c503722cc981fc463703472d2b876448",
        strip_prefix = "elasticlunr-rs-3.0.1",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.elasticlunr-rs-3.0.1.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__env_logger__0_9_3",
        url = "https://crates.io/api/v1/crates/env_logger/0.9.3/download",
        type = "tar.gz",
        sha256 = "a12e6657c4c97ebab115a42dcee77225f7f482cdd841cf7088c657a42e9e00e7",
        strip_prefix = "env_logger-0.9.3",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.env_logger-0.9.3.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__fastrand__1_8_0",
        url = "https://crates.io/api/v1/crates/fastrand/1.8.0/download",
        type = "tar.gz",
        sha256 = "a7a407cfaa3385c4ae6b23e84623d48c2798d06e3e6a1878f7f59f17b3f86499",
        strip_prefix = "fastrand-1.8.0",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.fastrand-1.8.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__filetime__0_2_19",
        url = "https://crates.io/api/v1/crates/filetime/0.2.19/download",
        type = "tar.gz",
        sha256 = "4e884668cd0c7480504233e951174ddc3b382f7c2666e3b7310b5c4e7b0c37f9",
        strip_prefix = "filetime-0.2.19",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.filetime-0.2.19.bazel"),
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
        name = "raze__form_urlencoded__1_1_0",
        url = "https://crates.io/api/v1/crates/form_urlencoded/1.1.0/download",
        type = "tar.gz",
        sha256 = "a9c384f161156f5260c24a097c56119f9be8c798586aecc13afbcbe7b7e26bf8",
        strip_prefix = "form_urlencoded-1.1.0",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.form_urlencoded-1.1.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__fsevent__0_4_0",
        url = "https://crates.io/api/v1/crates/fsevent/0.4.0/download",
        type = "tar.gz",
        sha256 = "5ab7d1bd1bd33cc98b0889831b72da23c0aa4df9cec7e0702f46ecea04b35db6",
        strip_prefix = "fsevent-0.4.0",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.fsevent-0.4.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__fsevent_sys__2_0_1",
        url = "https://crates.io/api/v1/crates/fsevent-sys/2.0.1/download",
        type = "tar.gz",
        sha256 = "f41b048a94555da0f42f1d632e2e19510084fb8e303b0daa2816e733fb3644a0",
        strip_prefix = "fsevent-sys-2.0.1",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.fsevent-sys-2.0.1.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__fuchsia_zircon__0_3_3",
        url = "https://crates.io/api/v1/crates/fuchsia-zircon/0.3.3/download",
        type = "tar.gz",
        sha256 = "2e9763c69ebaae630ba35f74888db465e49e259ba1bc0eda7d06f4a067615d82",
        strip_prefix = "fuchsia-zircon-0.3.3",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.fuchsia-zircon-0.3.3.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__fuchsia_zircon_sys__0_3_3",
        url = "https://crates.io/api/v1/crates/fuchsia-zircon-sys/0.3.3/download",
        type = "tar.gz",
        sha256 = "3dcaa9ae7725d12cdb85b3ad99a434db70b468c09ded17e012d86b5c1010f7a7",
        strip_prefix = "fuchsia-zircon-sys-0.3.3",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.fuchsia-zircon-sys-0.3.3.bazel"),
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
        name = "raze__futures_channel__0_3_25",
        url = "https://crates.io/api/v1/crates/futures-channel/0.3.25/download",
        type = "tar.gz",
        sha256 = "52ba265a92256105f45b719605a571ffe2d1f0fea3807304b522c1d778f79eed",
        strip_prefix = "futures-channel-0.3.25",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.futures-channel-0.3.25.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__futures_core__0_3_25",
        url = "https://crates.io/api/v1/crates/futures-core/0.3.25/download",
        type = "tar.gz",
        sha256 = "04909a7a7e4633ae6c4a9ab280aeb86da1236243a77b694a49eacd659a4bd3ac",
        strip_prefix = "futures-core-0.3.25",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.futures-core-0.3.25.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__futures_macro__0_3_25",
        url = "https://crates.io/api/v1/crates/futures-macro/0.3.25/download",
        type = "tar.gz",
        sha256 = "bdfb8ce053d86b91919aad980c220b1fb8401a9394410e1c289ed7e66b61835d",
        strip_prefix = "futures-macro-0.3.25",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.futures-macro-0.3.25.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__futures_sink__0_3_25",
        url = "https://crates.io/api/v1/crates/futures-sink/0.3.25/download",
        type = "tar.gz",
        sha256 = "39c15cf1a4aa79df40f1bb462fb39676d0ad9e366c2a33b590d7c66f4f81fcf9",
        strip_prefix = "futures-sink-0.3.25",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.futures-sink-0.3.25.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__futures_task__0_3_25",
        url = "https://crates.io/api/v1/crates/futures-task/0.3.25/download",
        type = "tar.gz",
        sha256 = "2ffb393ac5d9a6eaa9d3fdf37ae2776656b706e200c8e16b1bdb227f5198e6ea",
        strip_prefix = "futures-task-0.3.25",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.futures-task-0.3.25.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__futures_util__0_3_25",
        url = "https://crates.io/api/v1/crates/futures-util/0.3.25/download",
        type = "tar.gz",
        sha256 = "197676987abd2f9cadff84926f410af1c183608d36641465df73ae8211dc65d6",
        strip_prefix = "futures-util-0.3.25",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.futures-util-0.3.25.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__generic_array__0_14_6",
        url = "https://crates.io/api/v1/crates/generic-array/0.14.6/download",
        type = "tar.gz",
        sha256 = "bff49e947297f3312447abdca79f45f4738097cc82b06e72054d2223f601f1b9",
        strip_prefix = "generic-array-0.14.6",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.generic-array-0.14.6.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__getrandom__0_2_8",
        url = "https://crates.io/api/v1/crates/getrandom/0.2.8/download",
        type = "tar.gz",
        sha256 = "c05aeb6a22b8f62540c194aac980f2115af067bfe15a0734d7277a768d396b31",
        strip_prefix = "getrandom-0.2.8",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.getrandom-0.2.8.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__gitignore__1_0_7",
        url = "https://crates.io/api/v1/crates/gitignore/1.0.7/download",
        type = "tar.gz",
        sha256 = "78aa90e4620c1498ac434c06ba6e521b525794bbdacf085d490cc794b4a2f9a4",
        strip_prefix = "gitignore-1.0.7",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.gitignore-1.0.7.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__glob__0_3_1",
        url = "https://crates.io/api/v1/crates/glob/0.3.1/download",
        type = "tar.gz",
        sha256 = "d2fabcfbdc87f4758337ca535fb41a6d701b65693ce38287d856d1674551ec9b",
        strip_prefix = "glob-0.3.1",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.glob-0.3.1.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__h2__0_3_15",
        url = "https://crates.io/api/v1/crates/h2/0.3.15/download",
        type = "tar.gz",
        sha256 = "5f9f29bc9dda355256b2916cf526ab02ce0aeaaaf2bad60d65ef3f12f11dd0f4",
        strip_prefix = "h2-0.3.15",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.h2-0.3.15.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__handlebars__4_3_6",
        url = "https://crates.io/api/v1/crates/handlebars/4.3.6/download",
        type = "tar.gz",
        sha256 = "035ef95d03713f2c347a72547b7cd38cbc9af7cd51e6099fb62d586d4a6dee3a",
        strip_prefix = "handlebars-4.3.6",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.handlebars-4.3.6.bazel"),
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
        name = "raze__headers__0_3_8",
        url = "https://crates.io/api/v1/crates/headers/0.3.8/download",
        type = "tar.gz",
        sha256 = "f3e372db8e5c0d213e0cd0b9be18be2aca3d44cf2fe30a9d46a65581cd454584",
        strip_prefix = "headers-0.3.8",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.headers-0.3.8.bazel"),
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
        name = "raze__hermit_abi__0_1_19",
        url = "https://crates.io/api/v1/crates/hermit-abi/0.1.19/download",
        type = "tar.gz",
        sha256 = "62b467343b94ba476dcb2500d242dadbb39557df889310ac77c5d99100aaac33",
        strip_prefix = "hermit-abi-0.1.19",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.hermit-abi-0.1.19.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__hermit_abi__0_2_6",
        url = "https://crates.io/api/v1/crates/hermit-abi/0.2.6/download",
        type = "tar.gz",
        sha256 = "ee512640fe35acbfb4bb779db6f0d80704c2cacfa2e39b601ef3e3f47d1ae4c7",
        strip_prefix = "hermit-abi-0.2.6",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.hermit-abi-0.2.6.bazel"),
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
        name = "raze__http__0_2_8",
        url = "https://crates.io/api/v1/crates/http/0.2.8/download",
        type = "tar.gz",
        sha256 = "75f43d41e26995c17e71ee126451dd3941010b0514a81a9d11f3b341debc2399",
        strip_prefix = "http-0.2.8",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.http-0.2.8.bazel"),
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
        name = "raze__httpdate__1_0_2",
        url = "https://crates.io/api/v1/crates/httpdate/1.0.2/download",
        type = "tar.gz",
        sha256 = "c4a1e36c821dbe04574f602848a19f742f4fb3c98d40449f11bcad18d6b17421",
        strip_prefix = "httpdate-1.0.2",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.httpdate-1.0.2.bazel"),
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
        name = "raze__hyper__0_14_23",
        url = "https://crates.io/api/v1/crates/hyper/0.14.23/download",
        type = "tar.gz",
        sha256 = "034711faac9d2166cb1baf1a2fb0b60b1f277f8492fd72176c17f3515e1abd3c",
        strip_prefix = "hyper-0.14.23",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.hyper-0.14.23.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__iana_time_zone__0_1_53",
        url = "https://crates.io/api/v1/crates/iana-time-zone/0.1.53/download",
        type = "tar.gz",
        sha256 = "64c122667b287044802d6ce17ee2ddf13207ed924c712de9a66a5814d5b64765",
        strip_prefix = "iana-time-zone-0.1.53",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.iana-time-zone-0.1.53.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__iana_time_zone_haiku__0_1_1",
        url = "https://crates.io/api/v1/crates/iana-time-zone-haiku/0.1.1/download",
        type = "tar.gz",
        sha256 = "0703ae284fc167426161c2e3f1da3ea71d94b21bedbcc9494e92b28e334e3dca",
        strip_prefix = "iana-time-zone-haiku-0.1.1",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.iana-time-zone-haiku-0.1.1.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__idna__0_3_0",
        url = "https://crates.io/api/v1/crates/idna/0.3.0/download",
        type = "tar.gz",
        sha256 = "e14ddfc70884202db2244c223200c204c2bda1bc6e0998d11b5e024d657209e6",
        strip_prefix = "idna-0.3.0",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.idna-0.3.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__indexmap__1_9_2",
        url = "https://crates.io/api/v1/crates/indexmap/1.9.2/download",
        type = "tar.gz",
        sha256 = "1885e79c1fc4b10f0e172c475f458b7f7b93061064d98c3293e98c5ba0c8b399",
        strip_prefix = "indexmap-1.9.2",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.indexmap-1.9.2.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__inotify__0_7_1",
        url = "https://crates.io/api/v1/crates/inotify/0.7.1/download",
        type = "tar.gz",
        sha256 = "4816c66d2c8ae673df83366c18341538f234a26d65a9ecea5c348b453ac1d02f",
        strip_prefix = "inotify-0.7.1",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.inotify-0.7.1.bazel"),
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
        name = "raze__instant__0_1_12",
        url = "https://crates.io/api/v1/crates/instant/0.1.12/download",
        type = "tar.gz",
        sha256 = "7a5bbe824c507c5da5956355e86a746d82e0e1464f65d862cc5e71da70e94b2c",
        strip_prefix = "instant-0.1.12",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.instant-0.1.12.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__iovec__0_1_4",
        url = "https://crates.io/api/v1/crates/iovec/0.1.4/download",
        type = "tar.gz",
        sha256 = "b2b3ea6ff95e175473f8ffe6a7eb7c00d054240321b84c57051175fe3c1e075e",
        strip_prefix = "iovec-0.1.4",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.iovec-0.1.4.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__itoa__1_0_5",
        url = "https://crates.io/api/v1/crates/itoa/1.0.5/download",
        type = "tar.gz",
        sha256 = "fad582f4b9e86b6caa621cabeb0963332d92eea04729ab12892c2533951e6440",
        strip_prefix = "itoa-1.0.5",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.itoa-1.0.5.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__js_sys__0_3_60",
        url = "https://crates.io/api/v1/crates/js-sys/0.3.60/download",
        type = "tar.gz",
        sha256 = "49409df3e3bf0856b916e2ceaca09ee28e6871cf7d9ce97a692cacfdb2a25a47",
        strip_prefix = "js-sys-0.3.60",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.js-sys-0.3.60.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__kernel32_sys__0_2_2",
        url = "https://crates.io/api/v1/crates/kernel32-sys/0.2.2/download",
        type = "tar.gz",
        sha256 = "7507624b29483431c0ba2d82aece8ca6cdba9382bff4ddd0f7490560c056098d",
        strip_prefix = "kernel32-sys-0.2.2",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.kernel32-sys-0.2.2.bazel"),
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
        name = "raze__lazycell__1_3_0",
        url = "https://crates.io/api/v1/crates/lazycell/1.3.0/download",
        type = "tar.gz",
        sha256 = "830d08ce1d1d941e6b30645f1a0eb5643013d835ce3779a5fc208261dbe10f55",
        strip_prefix = "lazycell-1.3.0",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.lazycell-1.3.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__libc__0_2_139",
        url = "https://crates.io/api/v1/crates/libc/0.2.139/download",
        type = "tar.gz",
        sha256 = "201de327520df007757c1f0adce6e827fe8562fbc28bfd9c15571c66ca1f5f79",
        strip_prefix = "libc-0.2.139",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.libc-0.2.139.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__link_cplusplus__1_0_8",
        url = "https://crates.io/api/v1/crates/link-cplusplus/1.0.8/download",
        type = "tar.gz",
        sha256 = "ecd207c9c713c34f95a097a5b029ac2ce6010530c7b49d7fea24d977dede04f5",
        strip_prefix = "link-cplusplus-1.0.8",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.link-cplusplus-1.0.8.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__lock_api__0_4_9",
        url = "https://crates.io/api/v1/crates/lock_api/0.4.9/download",
        type = "tar.gz",
        sha256 = "435011366fe56583b16cf956f9df0095b405b82d76425bc8981c0e22e60ec4df",
        strip_prefix = "lock_api-0.4.9",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.lock_api-0.4.9.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__log__0_4_17",
        url = "https://crates.io/api/v1/crates/log/0.4.17/download",
        type = "tar.gz",
        sha256 = "abb12e687cfb44aa40f41fc3978ef76448f9b6038cad6aef4259d3c095a2382e",
        strip_prefix = "log-0.4.17",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.log-0.4.17.bazel"),
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
        name = "raze__mdbook__0_4_22",
        url = "https://crates.io/api/v1/crates/mdbook/0.4.22/download",
        type = "tar.gz",
        sha256 = "6b61566b406cbd75d81c634763d6c90779ca9db80202921c884348d172ada70d",
        strip_prefix = "mdbook-0.4.22",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.mdbook-0.4.22.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__memchr__2_5_0",
        url = "https://crates.io/api/v1/crates/memchr/2.5.0/download",
        type = "tar.gz",
        sha256 = "2dffe52ecf27772e601905b7522cb4ef790d2cc203488bbd0e2fe85fcb74566d",
        strip_prefix = "memchr-2.5.0",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.memchr-2.5.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__mime__0_3_16",
        url = "https://crates.io/api/v1/crates/mime/0.3.16/download",
        type = "tar.gz",
        sha256 = "2a60c7ce501c71e03a9c9c0d35b861413ae925bd979cc7a4e30d060069aaac8d",
        strip_prefix = "mime-0.3.16",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.mime-0.3.16.bazel"),
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
        name = "raze__mio__0_6_23",
        url = "https://crates.io/api/v1/crates/mio/0.6.23/download",
        type = "tar.gz",
        sha256 = "4afd66f5b91bf2a3bc13fad0e21caedac168ca4c707504e75585648ae80e4cc4",
        strip_prefix = "mio-0.6.23",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.mio-0.6.23.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__mio__0_8_5",
        url = "https://crates.io/api/v1/crates/mio/0.8.5/download",
        type = "tar.gz",
        sha256 = "e5d732bc30207a6423068df043e3d02e0735b155ad7ce1a6f76fe2baa5b158de",
        strip_prefix = "mio-0.8.5",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.mio-0.8.5.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__mio_extras__2_0_6",
        url = "https://crates.io/api/v1/crates/mio-extras/2.0.6/download",
        type = "tar.gz",
        sha256 = "52403fe290012ce777c4626790c8951324a2b9e3316b3143779c72b029742f19",
        strip_prefix = "mio-extras-2.0.6",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.mio-extras-2.0.6.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__miow__0_2_2",
        url = "https://crates.io/api/v1/crates/miow/0.2.2/download",
        type = "tar.gz",
        sha256 = "ebd808424166322d4a38da87083bfddd3ac4c131334ed55856112eb06d46944d",
        strip_prefix = "miow-0.2.2",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.miow-0.2.2.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__net2__0_2_38",
        url = "https://crates.io/api/v1/crates/net2/0.2.38/download",
        type = "tar.gz",
        sha256 = "74d0df99cfcd2530b2e694f6e17e7f37b8e26bb23983ac530c0c97408837c631",
        strip_prefix = "net2-0.2.38",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.net2-0.2.38.bazel"),
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
        name = "raze__notify__4_0_17",
        url = "https://crates.io/api/v1/crates/notify/4.0.17/download",
        type = "tar.gz",
        sha256 = "ae03c8c853dba7bfd23e571ff0cff7bc9dceb40a4cd684cd1681824183f45257",
        strip_prefix = "notify-4.0.17",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.notify-4.0.17.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__num_integer__0_1_45",
        url = "https://crates.io/api/v1/crates/num-integer/0.1.45/download",
        type = "tar.gz",
        sha256 = "225d3389fb3509a24c93f5c29eb6bde2586b98d9f016636dff58d7c6f7569cd9",
        strip_prefix = "num-integer-0.1.45",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.num-integer-0.1.45.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__num_traits__0_2_15",
        url = "https://crates.io/api/v1/crates/num-traits/0.2.15/download",
        type = "tar.gz",
        sha256 = "578ede34cf02f8924ab9447f50c28075b4d3e5b269972345e7e0372b38c6cdcd",
        strip_prefix = "num-traits-0.2.15",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.num-traits-0.2.15.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__num_cpus__1_15_0",
        url = "https://crates.io/api/v1/crates/num_cpus/1.15.0/download",
        type = "tar.gz",
        sha256 = "0fac9e2da13b5eb447a6ce3d392f23a29d8694bff781bf03a16cd9ac8697593b",
        strip_prefix = "num_cpus-1.15.0",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.num_cpus-1.15.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__once_cell__1_17_0",
        url = "https://crates.io/api/v1/crates/once_cell/1.17.0/download",
        type = "tar.gz",
        sha256 = "6f61fba1741ea2b3d6a1e3178721804bb716a68a6aeba1149b5d52e3d464ea66",
        strip_prefix = "once_cell-1.17.0",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.once_cell-1.17.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__opener__0_5_0",
        url = "https://crates.io/api/v1/crates/opener/0.5.0/download",
        type = "tar.gz",
        sha256 = "4ea3ebcd72a54701f56345f16785a6d3ac2df7e986d273eb4395c0b01db17952",
        strip_prefix = "opener-0.5.0",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.opener-0.5.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__os_str_bytes__6_4_1",
        url = "https://crates.io/api/v1/crates/os_str_bytes/6.4.1/download",
        type = "tar.gz",
        sha256 = "9b7820b9daea5457c9f21c69448905d723fbd21136ccf521748f23fd49e723ee",
        strip_prefix = "os_str_bytes-6.4.1",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.os_str_bytes-6.4.1.bazel"),
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
        name = "raze__parking_lot_core__0_9_6",
        url = "https://crates.io/api/v1/crates/parking_lot_core/0.9.6/download",
        type = "tar.gz",
        sha256 = "ba1ef8814b5c993410bb3adfad7a5ed269563e4a2f90c41f5d85be7fb47133bf",
        strip_prefix = "parking_lot_core-0.9.6",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.parking_lot_core-0.9.6.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__percent_encoding__2_2_0",
        url = "https://crates.io/api/v1/crates/percent-encoding/2.2.0/download",
        type = "tar.gz",
        sha256 = "478c572c3d73181ff3c2539045f6eb99e5491218eae919370993b890cdbdd98e",
        strip_prefix = "percent-encoding-2.2.0",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.percent-encoding-2.2.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__pest__2_5_3",
        url = "https://crates.io/api/v1/crates/pest/2.5.3/download",
        type = "tar.gz",
        sha256 = "4257b4a04d91f7e9e6290be5d3da4804dd5784fafde3a497d73eb2b4a158c30a",
        strip_prefix = "pest-2.5.3",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.pest-2.5.3.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__pest_derive__2_5_3",
        url = "https://crates.io/api/v1/crates/pest_derive/2.5.3/download",
        type = "tar.gz",
        sha256 = "241cda393b0cdd65e62e07e12454f1f25d57017dcc514b1514cd3c4645e3a0a6",
        strip_prefix = "pest_derive-2.5.3",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.pest_derive-2.5.3.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__pest_generator__2_5_3",
        url = "https://crates.io/api/v1/crates/pest_generator/2.5.3/download",
        type = "tar.gz",
        sha256 = "46b53634d8c8196302953c74d5352f33d0c512a9499bd2ce468fc9f4128fa27c",
        strip_prefix = "pest_generator-2.5.3",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.pest_generator-2.5.3.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__pest_meta__2_5_3",
        url = "https://crates.io/api/v1/crates/pest_meta/2.5.3/download",
        type = "tar.gz",
        sha256 = "0ef4f1332a8d4678b41966bb4cc1d0676880e84183a1ecc3f4b69f03e99c7a51",
        strip_prefix = "pest_meta-2.5.3",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.pest_meta-2.5.3.bazel"),
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
        name = "raze__pin_project__1_0_12",
        url = "https://crates.io/api/v1/crates/pin-project/1.0.12/download",
        type = "tar.gz",
        sha256 = "ad29a609b6bcd67fee905812e544992d216af9d755757c05ed2d0e15a74c6ecc",
        strip_prefix = "pin-project-1.0.12",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.pin-project-1.0.12.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__pin_project_internal__1_0_12",
        url = "https://crates.io/api/v1/crates/pin-project-internal/1.0.12/download",
        type = "tar.gz",
        sha256 = "069bdb1e05adc7a8990dce9cc75370895fbe4e3d58b9b73bf1aee56359344a55",
        strip_prefix = "pin-project-internal-1.0.12",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.pin-project-internal-1.0.12.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__pin_project_lite__0_2_9",
        url = "https://crates.io/api/v1/crates/pin-project-lite/0.2.9/download",
        type = "tar.gz",
        sha256 = "e0a7ae3ac2f1173085d398531c705756c94a4c56843785df85a60c1a0afac116",
        strip_prefix = "pin-project-lite-0.2.9",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.pin-project-lite-0.2.9.bazel"),
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
        name = "raze__proc_macro2__1_0_49",
        url = "https://crates.io/api/v1/crates/proc-macro2/1.0.49/download",
        type = "tar.gz",
        sha256 = "57a8eca9f9c4ffde41714334dee777596264c7825420f521abc92b5b5deb63a5",
        strip_prefix = "proc-macro2-1.0.49",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.proc-macro2-1.0.49.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__pulldown_cmark__0_9_2",
        url = "https://crates.io/api/v1/crates/pulldown-cmark/0.9.2/download",
        type = "tar.gz",
        sha256 = "2d9cc634bc78768157b5cbfe988ffcd1dcba95cd2b2f03a88316c08c6d00ed63",
        strip_prefix = "pulldown-cmark-0.9.2",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.pulldown-cmark-0.9.2.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__quote__1_0_23",
        url = "https://crates.io/api/v1/crates/quote/1.0.23/download",
        type = "tar.gz",
        sha256 = "8856d8364d252a14d474036ea1358d63c9e6965c8e5c1885c18f73d70bff9c7b",
        strip_prefix = "quote-1.0.23",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.quote-1.0.23.bazel"),
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
        name = "raze__redox_syscall__0_2_16",
        url = "https://crates.io/api/v1/crates/redox_syscall/0.2.16/download",
        type = "tar.gz",
        sha256 = "fb5a58c1855b4b6819d59012155603f0b22ad30cad752600aadfcb695265519a",
        strip_prefix = "redox_syscall-0.2.16",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.redox_syscall-0.2.16.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__regex__1_7_1",
        url = "https://crates.io/api/v1/crates/regex/1.7.1/download",
        type = "tar.gz",
        sha256 = "48aaa5748ba571fb95cd2c85c09f629215d3a6ece942baa100950af03a34f733",
        strip_prefix = "regex-1.7.1",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.regex-1.7.1.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__regex_automata__0_1_10",
        url = "https://crates.io/api/v1/crates/regex-automata/0.1.10/download",
        type = "tar.gz",
        sha256 = "6c230d73fb8d8c1b9c0b3135c5142a8acee3a0558fb8db5cf1cb65f8d7862132",
        strip_prefix = "regex-automata-0.1.10",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.regex-automata-0.1.10.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__regex_syntax__0_6_28",
        url = "https://crates.io/api/v1/crates/regex-syntax/0.6.28/download",
        type = "tar.gz",
        sha256 = "456c603be3e8d448b072f410900c09faf164fbce2d480456f50eea6e25f9c848",
        strip_prefix = "regex-syntax-0.6.28",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.regex-syntax-0.6.28.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__remove_dir_all__0_5_3",
        url = "https://crates.io/api/v1/crates/remove_dir_all/0.5.3/download",
        type = "tar.gz",
        sha256 = "3acd125665422973a33ac9d3dd2df85edad0f4ae9b00dafb1a05e43a9f5ef8e7",
        strip_prefix = "remove_dir_all-0.5.3",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.remove_dir_all-0.5.3.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__rustls_pemfile__0_2_1",
        url = "https://crates.io/api/v1/crates/rustls-pemfile/0.2.1/download",
        type = "tar.gz",
        sha256 = "5eebeaeb360c87bfb72e84abdb3447159c0eaececf1bef2aecd65a8be949d1c9",
        strip_prefix = "rustls-pemfile-0.2.1",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.rustls-pemfile-0.2.1.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__ryu__1_0_12",
        url = "https://crates.io/api/v1/crates/ryu/1.0.12/download",
        type = "tar.gz",
        sha256 = "7b4b9743ed687d4b4bcedf9ff5eaa7398495ae14e61cba0a295704edbc7decde",
        strip_prefix = "ryu-1.0.12",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.ryu-1.0.12.bazel"),
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
        name = "raze__scopeguard__1_1_0",
        url = "https://crates.io/api/v1/crates/scopeguard/1.1.0/download",
        type = "tar.gz",
        sha256 = "d29ab0c6d3fc0ee92fe66e2d99f700eab17a8d57d1c1d3b748380fb20baa78cd",
        strip_prefix = "scopeguard-1.1.0",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.scopeguard-1.1.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__scratch__1_0_3",
        url = "https://crates.io/api/v1/crates/scratch/1.0.3/download",
        type = "tar.gz",
        sha256 = "ddccb15bcce173023b3fedd9436f882a0739b8dfb45e4f6b6002bee5929f61b2",
        strip_prefix = "scratch-1.0.3",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.scratch-1.0.3.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__serde__1_0_152",
        url = "https://crates.io/api/v1/crates/serde/1.0.152/download",
        type = "tar.gz",
        sha256 = "bb7d1f0d3021d347a83e556fc4683dea2ea09d87bccdf88ff5c12545d89d5efb",
        strip_prefix = "serde-1.0.152",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.serde-1.0.152.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__serde_derive__1_0_152",
        url = "https://crates.io/api/v1/crates/serde_derive/1.0.152/download",
        type = "tar.gz",
        sha256 = "af487d118eecd09402d70a5d72551860e788df87b464af30e5ea6a38c75c541e",
        strip_prefix = "serde_derive-1.0.152",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.serde_derive-1.0.152.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__serde_json__1_0_91",
        url = "https://crates.io/api/v1/crates/serde_json/1.0.91/download",
        type = "tar.gz",
        sha256 = "877c235533714907a8c2464236f5c4b2a17262ef1bd71f38f35ea592c8da6883",
        strip_prefix = "serde_json-1.0.91",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.serde_json-1.0.91.bazel"),
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
        name = "raze__sha_1__0_10_1",
        url = "https://crates.io/api/v1/crates/sha-1/0.10.1/download",
        type = "tar.gz",
        sha256 = "f5058ada175748e33390e40e872bd0fe59a19f265d0158daa551c5a88a76009c",
        strip_prefix = "sha-1-0.10.1",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.sha-1-0.10.1.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__sha1__0_10_5",
        url = "https://crates.io/api/v1/crates/sha1/0.10.5/download",
        type = "tar.gz",
        sha256 = "f04293dc80c3993519f2d7f6f511707ee7094fe0c6d3406feb330cdb3540eba3",
        strip_prefix = "sha1-0.10.5",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.sha1-0.10.5.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__sha2__0_10_6",
        url = "https://crates.io/api/v1/crates/sha2/0.10.6/download",
        type = "tar.gz",
        sha256 = "82e6b795fe2e3b1e845bafcb27aa35405c4d47cdfc92af5fc8d3002f76cebdc0",
        strip_prefix = "sha2-0.10.6",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.sha2-0.10.6.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__shlex__1_1_0",
        url = "https://crates.io/api/v1/crates/shlex/1.1.0/download",
        type = "tar.gz",
        sha256 = "43b2853a4d09f215c24cc5489c992ce46052d359b5109343cbafbf26bc62f8a3",
        strip_prefix = "shlex-1.1.0",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.shlex-1.1.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__siphasher__0_3_10",
        url = "https://crates.io/api/v1/crates/siphasher/0.3.10/download",
        type = "tar.gz",
        sha256 = "7bd3e3206899af3f8b12af284fafc038cc1dc2b41d1b89dd17297221c5d225de",
        strip_prefix = "siphasher-0.3.10",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.siphasher-0.3.10.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__slab__0_4_7",
        url = "https://crates.io/api/v1/crates/slab/0.4.7/download",
        type = "tar.gz",
        sha256 = "4614a76b2a8be0058caa9dbbaf66d988527d86d003c11a94fbd335d7661edcef",
        strip_prefix = "slab-0.4.7",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.slab-0.4.7.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__smallvec__1_10_0",
        url = "https://crates.io/api/v1/crates/smallvec/1.10.0/download",
        type = "tar.gz",
        sha256 = "a507befe795404456341dfab10cef66ead4c041f62b8b11bbb92bffe5d0953e0",
        strip_prefix = "smallvec-1.10.0",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.smallvec-1.10.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__socket2__0_4_7",
        url = "https://crates.io/api/v1/crates/socket2/0.4.7/download",
        type = "tar.gz",
        sha256 = "02e2d2db9033d13a1567121ddd7a095ee144db4e1ca1b1bda3419bc0da294ebd",
        strip_prefix = "socket2-0.4.7",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.socket2-0.4.7.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__string_cache__0_8_4",
        url = "https://crates.io/api/v1/crates/string_cache/0.8.4/download",
        type = "tar.gz",
        sha256 = "213494b7a2b503146286049378ce02b482200519accc31872ee8be91fa820a08",
        strip_prefix = "string_cache-0.8.4",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.string_cache-0.8.4.bazel"),
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
        name = "raze__syn__1_0_107",
        url = "https://crates.io/api/v1/crates/syn/1.0.107/download",
        type = "tar.gz",
        sha256 = "1f4064b5b16e03ae50984a5a8ed5d4f8803e6bc1fd170a3cda91a1be4b18e3f5",
        strip_prefix = "syn-1.0.107",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.syn-1.0.107.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__tempfile__3_3_0",
        url = "https://crates.io/api/v1/crates/tempfile/3.3.0/download",
        type = "tar.gz",
        sha256 = "5cdb1ef4eaeeaddc8fbd371e5017057064af0911902ef36b39801f67cc6d79e4",
        strip_prefix = "tempfile-3.3.0",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.tempfile-3.3.0.bazel"),
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
        name = "raze__termcolor__1_1_3",
        url = "https://crates.io/api/v1/crates/termcolor/1.1.3/download",
        type = "tar.gz",
        sha256 = "bab24d30b911b2376f3a13cc2cd443142f0c81dda04c118693e35b3835757755",
        strip_prefix = "termcolor-1.1.3",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.termcolor-1.1.3.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__textwrap__0_16_0",
        url = "https://crates.io/api/v1/crates/textwrap/0.16.0/download",
        type = "tar.gz",
        sha256 = "222a222a5bfe1bba4a77b45ec488a741b3cb8872e5e499451fd7d0129c9c7c3d",
        strip_prefix = "textwrap-0.16.0",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.textwrap-0.16.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__thiserror__1_0_38",
        url = "https://crates.io/api/v1/crates/thiserror/1.0.38/download",
        type = "tar.gz",
        sha256 = "6a9cd18aa97d5c45c6603caea1da6628790b37f7a34b6ca89522331c5180fed0",
        strip_prefix = "thiserror-1.0.38",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.thiserror-1.0.38.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__thiserror_impl__1_0_38",
        url = "https://crates.io/api/v1/crates/thiserror-impl/1.0.38/download",
        type = "tar.gz",
        sha256 = "1fb327af4685e4d03fa8cbcf1716380da910eeb2bb8be417e7f9fd3fb164f36f",
        strip_prefix = "thiserror-impl-1.0.38",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.thiserror-impl-1.0.38.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__time__0_1_45",
        url = "https://crates.io/api/v1/crates/time/0.1.45/download",
        type = "tar.gz",
        sha256 = "1b797afad3f312d1c66a56d11d0316f916356d11bd158fbc6ca6389ff6bf805a",
        strip_prefix = "time-0.1.45",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.time-0.1.45.bazel"),
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
        name = "raze__tinyvec_macros__0_1_0",
        url = "https://crates.io/api/v1/crates/tinyvec_macros/0.1.0/download",
        type = "tar.gz",
        sha256 = "cda74da7e1a664f795bb1f8a87ec406fb89a02522cf6e50620d016add6dbbf5c",
        strip_prefix = "tinyvec_macros-0.1.0",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.tinyvec_macros-0.1.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__tokio__1_24_1",
        url = "https://crates.io/api/v1/crates/tokio/1.24.1/download",
        type = "tar.gz",
        sha256 = "1d9f76183f91ecfb55e1d7d5602bd1d979e38a3a522fe900241cf195624d67ae",
        strip_prefix = "tokio-1.24.1",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.tokio-1.24.1.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__tokio_macros__1_8_2",
        url = "https://crates.io/api/v1/crates/tokio-macros/1.8.2/download",
        type = "tar.gz",
        sha256 = "d266c00fde287f55d3f1c3e96c500c362a2b8c695076ec180f27918820bc6df8",
        strip_prefix = "tokio-macros-1.8.2",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.tokio-macros-1.8.2.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__tokio_stream__0_1_11",
        url = "https://crates.io/api/v1/crates/tokio-stream/0.1.11/download",
        type = "tar.gz",
        sha256 = "d660770404473ccd7bc9f8b28494a811bc18542b915c0855c51e8f419d5223ce",
        strip_prefix = "tokio-stream-0.1.11",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.tokio-stream-0.1.11.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__tokio_tungstenite__0_17_2",
        url = "https://crates.io/api/v1/crates/tokio-tungstenite/0.17.2/download",
        type = "tar.gz",
        sha256 = "f714dd15bead90401d77e04243611caec13726c2408afd5b31901dfcdcb3b181",
        strip_prefix = "tokio-tungstenite-0.17.2",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.tokio-tungstenite-0.17.2.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__tokio_util__0_7_4",
        url = "https://crates.io/api/v1/crates/tokio-util/0.7.4/download",
        type = "tar.gz",
        sha256 = "0bb2e075f03b3d66d8d8785356224ba688d2906a371015e225beeb65ca92c740",
        strip_prefix = "tokio-util-0.7.4",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.tokio-util-0.7.4.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__toml__0_5_10",
        url = "https://crates.io/api/v1/crates/toml/0.5.10/download",
        type = "tar.gz",
        sha256 = "1333c76748e868a4d9d1017b5ab53171dfd095f70c712fdb4653a406547f598f",
        strip_prefix = "toml-0.5.10",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.toml-0.5.10.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__topological_sort__0_1_0",
        url = "https://crates.io/api/v1/crates/topological-sort/0.1.0/download",
        type = "tar.gz",
        sha256 = "aa7c7f42dea4b1b99439786f5633aeb9c14c1b53f75e282803c2ec2ad545873c",
        strip_prefix = "topological-sort-0.1.0",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.topological-sort-0.1.0.bazel"),
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
        name = "raze__tracing_core__0_1_30",
        url = "https://crates.io/api/v1/crates/tracing-core/0.1.30/download",
        type = "tar.gz",
        sha256 = "24eb03ba0eab1fd845050058ce5e616558e8f8d8fca633e6b163fe25c797213a",
        strip_prefix = "tracing-core-0.1.30",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.tracing-core-0.1.30.bazel"),
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
        name = "raze__tungstenite__0_17_3",
        url = "https://crates.io/api/v1/crates/tungstenite/0.17.3/download",
        type = "tar.gz",
        sha256 = "e27992fd6a8c29ee7eef28fc78349aa244134e10ad447ce3b9f0ac0ed0fa4ce0",
        strip_prefix = "tungstenite-0.17.3",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.tungstenite-0.17.3.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__typenum__1_16_0",
        url = "https://crates.io/api/v1/crates/typenum/1.16.0/download",
        type = "tar.gz",
        sha256 = "497961ef93d974e23eb6f433eb5fe1b7930b659f06d12dec6fc44a8f554c0bba",
        strip_prefix = "typenum-1.16.0",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.typenum-1.16.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__ucd_trie__0_1_5",
        url = "https://crates.io/api/v1/crates/ucd-trie/0.1.5/download",
        type = "tar.gz",
        sha256 = "9e79c4d996edb816c91e4308506774452e55e95c3c9de07b6729e17e15a5ef81",
        strip_prefix = "ucd-trie-0.1.5",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.ucd-trie-0.1.5.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__unicase__2_6_0",
        url = "https://crates.io/api/v1/crates/unicase/2.6.0/download",
        type = "tar.gz",
        sha256 = "50f37be617794602aabbeee0be4f259dc1778fabe05e2d67ee8f79326d5cb4f6",
        strip_prefix = "unicase-2.6.0",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.unicase-2.6.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__unicode_bidi__0_3_8",
        url = "https://crates.io/api/v1/crates/unicode-bidi/0.3.8/download",
        type = "tar.gz",
        sha256 = "099b7128301d285f79ddd55b9a83d5e6b9e97c92e0ea0daebee7263e932de992",
        strip_prefix = "unicode-bidi-0.3.8",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.unicode-bidi-0.3.8.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__unicode_ident__1_0_6",
        url = "https://crates.io/api/v1/crates/unicode-ident/1.0.6/download",
        type = "tar.gz",
        sha256 = "84a22b9f218b40614adcb3f4ff08b703773ad44fa9423e4e0d346d5db86e4ebc",
        strip_prefix = "unicode-ident-1.0.6",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.unicode-ident-1.0.6.bazel"),
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
        name = "raze__unicode_width__0_1_10",
        url = "https://crates.io/api/v1/crates/unicode-width/0.1.10/download",
        type = "tar.gz",
        sha256 = "c0edd1e5b14653f783770bce4a4dabb4a5108a5370a5f5d8cfe8710c361f6c8b",
        strip_prefix = "unicode-width-0.1.10",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.unicode-width-0.1.10.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__url__2_3_1",
        url = "https://crates.io/api/v1/crates/url/2.3.1/download",
        type = "tar.gz",
        sha256 = "0d68c799ae75762b8c3fe375feb6600ef5602c883c5d21eb51c09f22b83c4643",
        strip_prefix = "url-2.3.1",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.url-2.3.1.bazel"),
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
        name = "raze__version_check__0_9_4",
        url = "https://crates.io/api/v1/crates/version_check/0.9.4/download",
        type = "tar.gz",
        sha256 = "49874b5167b65d7193b8aba1567f5c7d93d001cafc34600cee003eda787e483f",
        strip_prefix = "version_check-0.9.4",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.version_check-0.9.4.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__walkdir__2_3_2",
        url = "https://crates.io/api/v1/crates/walkdir/2.3.2/download",
        type = "tar.gz",
        sha256 = "808cf2735cd4b6866113f648b791c6adc5714537bc222d9347bb203386ffda56",
        strip_prefix = "walkdir-2.3.2",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.walkdir-2.3.2.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__want__0_3_0",
        url = "https://crates.io/api/v1/crates/want/0.3.0/download",
        type = "tar.gz",
        sha256 = "1ce8a968cb1cd110d136ff8b819a556d6fb6d919363c61534f6860c7eb172ba0",
        strip_prefix = "want-0.3.0",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.want-0.3.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__warp__0_3_3",
        url = "https://crates.io/api/v1/crates/warp/0.3.3/download",
        type = "tar.gz",
        sha256 = "ed7b8be92646fc3d18b06147664ebc5f48d222686cb11a8755e561a735aacc6d",
        strip_prefix = "warp-0.3.3",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.warp-0.3.3.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__wasi__0_10_0_wasi_snapshot_preview1",
        url = "https://crates.io/api/v1/crates/wasi/0.10.0+wasi-snapshot-preview1/download",
        type = "tar.gz",
        sha256 = "1a143597ca7c7793eff794def352d41792a93c481eb1042423ff7ff72ba2c31f",
        strip_prefix = "wasi-0.10.0+wasi-snapshot-preview1",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.wasi-0.10.0+wasi-snapshot-preview1.bazel"),
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
        name = "raze__wasm_bindgen__0_2_83",
        url = "https://crates.io/api/v1/crates/wasm-bindgen/0.2.83/download",
        type = "tar.gz",
        sha256 = "eaf9f5aceeec8be17c128b2e93e031fb8a4d469bb9c4ae2d7dc1888b26887268",
        strip_prefix = "wasm-bindgen-0.2.83",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.wasm-bindgen-0.2.83.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__wasm_bindgen_backend__0_2_83",
        url = "https://crates.io/api/v1/crates/wasm-bindgen-backend/0.2.83/download",
        type = "tar.gz",
        sha256 = "4c8ffb332579b0557b52d268b91feab8df3615f265d5270fec2a8c95b17c1142",
        strip_prefix = "wasm-bindgen-backend-0.2.83",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.wasm-bindgen-backend-0.2.83.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__wasm_bindgen_macro__0_2_83",
        url = "https://crates.io/api/v1/crates/wasm-bindgen-macro/0.2.83/download",
        type = "tar.gz",
        sha256 = "052be0f94026e6cbc75cdefc9bae13fd6052cdcaf532fa6c45e7ae33a1e6c810",
        strip_prefix = "wasm-bindgen-macro-0.2.83",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.wasm-bindgen-macro-0.2.83.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__wasm_bindgen_macro_support__0_2_83",
        url = "https://crates.io/api/v1/crates/wasm-bindgen-macro-support/0.2.83/download",
        type = "tar.gz",
        sha256 = "07bc0c051dc5f23e307b13285f9d75df86bfdf816c5721e573dec1f9b8aa193c",
        strip_prefix = "wasm-bindgen-macro-support-0.2.83",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.wasm-bindgen-macro-support-0.2.83.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__wasm_bindgen_shared__0_2_83",
        url = "https://crates.io/api/v1/crates/wasm-bindgen-shared/0.2.83/download",
        type = "tar.gz",
        sha256 = "1c38c045535d93ec4f0b4defec448e4291638ee608530863b1e2ba115d4fff7f",
        strip_prefix = "wasm-bindgen-shared-0.2.83",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.wasm-bindgen-shared-0.2.83.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__winapi__0_2_8",
        url = "https://crates.io/api/v1/crates/winapi/0.2.8/download",
        type = "tar.gz",
        sha256 = "167dc9d6949a9b857f3451275e911c3f44255842c1f7a76f33c55103a909087a",
        strip_prefix = "winapi-0.2.8",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.winapi-0.2.8.bazel"),
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
        name = "raze__winapi_build__0_1_1",
        url = "https://crates.io/api/v1/crates/winapi-build/0.1.1/download",
        type = "tar.gz",
        sha256 = "2d315eee3b34aca4797b2da6b13ed88266e6d612562a0c46390af8299fc699bc",
        strip_prefix = "winapi-build-0.1.1",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.winapi-build-0.1.1.bazel"),
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
        name = "raze__winapi_util__0_1_5",
        url = "https://crates.io/api/v1/crates/winapi-util/0.1.5/download",
        type = "tar.gz",
        sha256 = "70ec6ce85bb158151cae5e5c87f95a8e97d2c0c4b001223f33a334e3ce5de178",
        strip_prefix = "winapi-util-0.1.5",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.winapi-util-0.1.5.bazel"),
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
        name = "raze__windows_sys__0_42_0",
        url = "https://crates.io/api/v1/crates/windows-sys/0.42.0/download",
        type = "tar.gz",
        sha256 = "5a3e1820f08b8513f676f7ab6c1f99ff312fb97b553d30ff4dd86f9f15728aa7",
        strip_prefix = "windows-sys-0.42.0",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.windows-sys-0.42.0.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__windows_aarch64_gnullvm__0_42_1",
        url = "https://crates.io/api/v1/crates/windows_aarch64_gnullvm/0.42.1/download",
        type = "tar.gz",
        sha256 = "8c9864e83243fdec7fc9c5444389dcbbfd258f745e7853198f365e3c4968a608",
        strip_prefix = "windows_aarch64_gnullvm-0.42.1",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.windows_aarch64_gnullvm-0.42.1.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__windows_aarch64_msvc__0_42_1",
        url = "https://crates.io/api/v1/crates/windows_aarch64_msvc/0.42.1/download",
        type = "tar.gz",
        sha256 = "4c8b1b673ffc16c47a9ff48570a9d85e25d265735c503681332589af6253c6c7",
        strip_prefix = "windows_aarch64_msvc-0.42.1",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.windows_aarch64_msvc-0.42.1.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__windows_i686_gnu__0_42_1",
        url = "https://crates.io/api/v1/crates/windows_i686_gnu/0.42.1/download",
        type = "tar.gz",
        sha256 = "de3887528ad530ba7bdbb1faa8275ec7a1155a45ffa57c37993960277145d640",
        strip_prefix = "windows_i686_gnu-0.42.1",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.windows_i686_gnu-0.42.1.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__windows_i686_msvc__0_42_1",
        url = "https://crates.io/api/v1/crates/windows_i686_msvc/0.42.1/download",
        type = "tar.gz",
        sha256 = "bf4d1122317eddd6ff351aa852118a2418ad4214e6613a50e0191f7004372605",
        strip_prefix = "windows_i686_msvc-0.42.1",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.windows_i686_msvc-0.42.1.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__windows_x86_64_gnu__0_42_1",
        url = "https://crates.io/api/v1/crates/windows_x86_64_gnu/0.42.1/download",
        type = "tar.gz",
        sha256 = "c1040f221285e17ebccbc2591ffdc2d44ee1f9186324dd3e84e99ac68d699c45",
        strip_prefix = "windows_x86_64_gnu-0.42.1",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.windows_x86_64_gnu-0.42.1.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__windows_x86_64_gnullvm__0_42_1",
        url = "https://crates.io/api/v1/crates/windows_x86_64_gnullvm/0.42.1/download",
        type = "tar.gz",
        sha256 = "628bfdf232daa22b0d64fdb62b09fcc36bb01f05a3939e20ab73aaf9470d0463",
        strip_prefix = "windows_x86_64_gnullvm-0.42.1",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.windows_x86_64_gnullvm-0.42.1.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__windows_x86_64_msvc__0_42_1",
        url = "https://crates.io/api/v1/crates/windows_x86_64_msvc/0.42.1/download",
        type = "tar.gz",
        sha256 = "447660ad36a13288b1db4d4248e857b510e8c3a225c822ba4fb748c0aafecffd",
        strip_prefix = "windows_x86_64_msvc-0.42.1",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.windows_x86_64_msvc-0.42.1.bazel"),
    )

    maybe(
        http_archive,
        name = "raze__ws2_32_sys__0_2_1",
        url = "https://crates.io/api/v1/crates/ws2_32-sys/0.2.1/download",
        type = "tar.gz",
        sha256 = "d59cefebd0c892fa2dd6de581e937301d8552cb44489cdff035c6187cb63fa5e",
        strip_prefix = "ws2_32-sys-0.2.1",
        build_file = Label("//third_party/rust/cargo/remote:BUILD.ws2_32-sys-0.2.1.bazel"),
    )
