##SQLboiler

This rule allows to generate sqlboiler models and ORM code for golang from a stack of SQL migrations.

It uses `sql-migrate`, `sqlboiler` and `sqlboiker-crdb`.

###How to use

Create a package and create a `0_initial.sql` file with the following template:

```
-- +migrate Up

Your initial SQL goes here

-- +migrate Down

```

Then create a `BUILD` file with the following rules:

```
load("//build/sqlboiler:sqlboiler.bzl", "go_sqlboiler_library", "sqlboiler")
load("@io_bazel_rules_go//go:def.bzl", "go_library")

# gazelle:ignore .

sqlboiler(
    name = "sqlboiler",
    srcs = glob(["*.sql"]),
    tables = [
        <your table names>
    ],
)

go_sqlboiler_library(
    name = "sqlboiler_lib",
    importpath = "git.monogon.dev/source/nexantic.git/<target>",
    sqlboiler = ":sqlboiler",
)

go_library(
    name = "go_default_library",
    embed = [":sqlboiler_lib"],
    importpath = "git.monogon.dev/source/nexantic.git/<target>",
    visibility = ["//visibility:public"],
)

```

Replace `target` with your intended importpath and add all created tables to the `tables` argument of the sqlboiler rule.

Running the tasks will apply the migrations to a temporary database using `sql-migrate` and generate sqlboiler code from it.
The code will be importable from the specified `importpath`.

When making changes to the schema please generate new migrations using the `<n>_<description>.sql` pattern for file names.
The migrations will automatically be applied and the models updated.

Make sure to also update the `tables` argument when creating new tables in migrations.
