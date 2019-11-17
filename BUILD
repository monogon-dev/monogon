load("@bazel_gazelle//:def.bzl", "gazelle")

# gazelle:prefix git.monogon.dev/source/nexantic.git
# gazelle:exclude core/generated
# gazelle:exclude imports.go
gazelle(name = "gazelle")

# Shortcut for the Go SDK
alias(
    name = "go",
    actual = "@go_sdk//:bin/go",
)
