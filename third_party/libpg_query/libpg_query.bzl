load("@rules_cc//cc:defs.bzl", "cc_library")

filegroup(
    name = "all",
    srcs = glob(["**"]),
    visibility = ["//visibility:public"],
)

cc_library(
    name = "libpg_query",
    srcs = glob([
        "src/*.c",
        "src/*.h",
        "src/postgres/include/*.h",
        "src/postgres/include/**/*.h",
    ], [
        "src/pg_query_enum_defs.c",
        "src/pg_query_fingerprint_defs.c",
        "src/pg_query_fingerprint_conds.c",
        "src/pg_query_outfuncs_defs.c",
        "src/pg_query_outfuncs_conds.c",
        "src/pg_query_readfuncs_defs.c",
        "src/pg_query_readfuncs_conds.c",
        "src/pg_query_json_helper.c",
    ]) + [
        "vendor/protobuf-c/protobuf-c.h",
        "vendor/protobuf-c/protobuf-c.c",
        "vendor/xxhash/xxhash.c",
        "protobuf/pg_query.pb-c.c",
        "protobuf/pg_query.pb-c.h",
    ],
    textual_hdrs = [
        "src/pg_query_enum_defs.c",
        "src/pg_query_fingerprint_defs.c",
        "src/pg_query_fingerprint_conds.c",
        "src/pg_query_outfuncs_defs.c",
        "src/pg_query_outfuncs_conds.c",
        "src/pg_query_readfuncs_defs.c",
        "src/pg_query_readfuncs_conds.c",
        "src/pg_query_json_helper.c",
    ],
    hdrs = [
        "pg_query.h",
        "vendor/xxhash/xxhash.h",
    ],
    # Unfortunate. We should patch this library so that this doesn't pollute
    # all dependents.
    includes = [
        "vendor/xxhash",
        "src/postgres/include",
        "vendor",
        "vendor/protobuf-c",
        "src",
    ],
    copts = [
        "-Iexternal/libpg_query/protobuf",
        "-Iexternal/libpg_query/vendor/xxhash",
    ],
    visibility = [
        "@//third_party/libpg_query:__pkg__",
    ],
)
