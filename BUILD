load("@io_bazel_rules_go//go:def.bzl", "go_library")
load("@bazel_gazelle//:def.bzl", "gazelle")

# gazelle:prefix git.monogon.dev/source/nexantic.git
# gazelle:exclude core/generated
gazelle(name = "gazelle")

go_library(
    name = "go_default_library",
    srcs = ["imports.go"],
    importpath = "git.monogon.dev/source/nexantic.git",
    visibility = ["//visibility:public"],
    deps = [
        "@com_github_lopezator_sqlboiler_crdb//:go_default_library",
        "@com_github_rubenv_sql_migrate//sql-migrate:go_default_library",
        "@com_github_volatiletech_sqlboiler//:go_default_library",
        "@com_github_volatiletech_sqlboiler//queries/qmhelper:go_default_library",
        "@com_github_volatiletech_sqlboiler//randomize:go_default_library",
        "@com_github_volatiletech_sqlboiler//types:go_default_library",
    ],
)
