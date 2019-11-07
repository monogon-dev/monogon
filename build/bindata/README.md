##Bindata

This rule uses [go-bindata](https://github.com/kevinburke/go-bindata) to package arbitrary go files.
Please refer to the documentation there on how to use the packaged data.

Generally this rule is very similar to the `bindata` rule in the default go bazel package.
However this rule also creates an embeddable go library right away.

###How to use

Add the files you want to package to the `srcs` attribute, set the `package` attribute to the 
go package you want the result to be in and embed the rule into a `go_library`.

####Example: Packaging sql migrations

These rules package all `.sql` files into the target and make it accessible at `importpath` in the package `models`. 
```

go_library(
    name = "go_default_library",
    embed = [
        ":migrations_pack",
    ],
    importpath = "git.monogon.dev/source/nexantic.git/golibs/minijob/generated/sql",
    visibility = ["//visibility:public"],
)

bindata(
    name = "migrations_pack",
    package = "models",
    srcs = glob(["*.sql"]),
)

```
