Fietsje
=======

The little Gazelle that could.

Introduction
------------

Fietsje is a dependency management system for Go dependencies in monogon. It
does not replace either gomods or Gazelle, but instead builds upon both on them
in a way that makes sense for our particular usecase: pulling in a large set of
dependency trees from third\_party projects, and sticking to those as much as
possible.

When run, Fietsje consults rules written themselves in Go (in `deps_.*go`
files), and uses this high-level intent to write a `repositories.bzl` file
that is then consumed by Gazelle. It caches 'locked' versions (ie. Go import
path and version to a particular checksum) in the Shelf, a text proto file
that lives alongside `repositories.bzl`. The Shelf should not be modified
manually.

The effective source of truth used for builds is still the `repositories.bzl`
file in the actual build path. Definitions in Go are in turn the high-level
intent that is used to build `repositories.bzl`.

Running
-------

You should run Fietsje any time you want to update dependencies. The following
should be a no-op if you haven't changed anything in `deps_*.go`:

    scripts/bin/bazel run //:fietsje

Otherwise, if any definition in build/fietsje/deps_*.go has been changed,
third_party/go/repositories.bzl will now reflect that.

Fietsje Definition DSL (collect/use/...)
----------------------------------------

Definitions are kept in pure Go source, with a light DSL focused around a
'planner' builder.

The builder allows for two kinds of actions:
 - import a high level dependency (eg. Kubernetes, google/tpm) at a particular
   version. This is done using the `collect()` call. The dependency will now
   be part of the build, but its transitive dependencies will not. A special
   flavor of collect() is collectOverride(), that explicitely allows for
   overriding a dependency that has already been pulled in by another high
   level dependency.
 - enable a transitive dependency defined by a high-level definition using the `use()`
   call. This can only be done in a `collection` builder context, ie. after a
   `collect()`/`collectOverride()`call.
   
In addition, the builder allows to augment a `collection` context with build flags
(like enabled patches, build tags, etc) that will be applied to the next `.use()`
call only. This is done by calling `.with()`.

In general, `.collect()`/`.collectOverride()` calls should be limited only to
dependencies 'we' (as developers) want. These 'high-level' dependencies are
large projects like Kubernetes, or direct imports from monogon itself. Every
transitive dependency of those should just be enabled by calling `.use()`,
instead of another `.collectOverride()` call that might pin it to a wrong
version.

After updating definitions, run Fietsje as above.

How to: add a new high-level dependency
---------------------------------------

To add a new high-level dependency, first consider making a new `deps_*.go`
file for it. If you're pulling in a separate ecosystem of code (ie. a large
third-party project like kubernetes), it should live in its own file for
clarity. If you're just pulling in a simple dependency (eg. a library low on
transitive dependencies) you can drop it into `main.go`.

The first step is to pick a version of the dependency you want to use. If
possible, pick a tag/release. Otherwise, pick the current master commit hash.
You can find version information by visiting the project's git repository web
viewer, or first cloning the repository locally.

Once you've picked a version, add a line like this:

    p.collect("github.com/example/foo", "1.2.3")

If you now re-run Fietsje and rebuild your code, it should be able to link
against the dependency directly. If this works, you're done. If not, you will
start getting errors about the newly included library trying to link against
missing dependencies (ie. external Bazel workspaces). This means you need to
enable these transitive dependencies for the high-level dependency you've just
included.

If your high-level dependency contains a go.mod/go.sum file, you can call
`.use` on the return of the `collect()` call to enable them. Only enable the
ones that are necessary to build your code. In the future, audit flows might be
implemented to find and eradicate unused transitive dependencies, while enabling
ones that are needed - but for now this has to be done manually - usually by a
cycle of:

 - try to build your code
 - find missing transitive library, enable via .use()
 - repeat until code builds

With our previous example, enabling transitive dependencies would look something
like this:

    p.collect(
        "github.com/example/foo", "1.2.3",
    ).use(
        "github.com/example/libbar",
        "github.com/example/libbaz",
        "github.com/golang/glog",
    )

What this means is that github.com/{example/libbar,example/libbaz,golang/glog}
will now be available to the build at whatever version example/foo defines them
in its go.mod/go.sum.

If your high-level dependency is not go.mod/go.sum compatible, you have
different ways to proceed:

 - if the project uses some alternative resolution/vendoring code, write
   support for it in transitive.go/`getTransitiveDeps`
 - otherwise, if you're not in a rush, try to convince and/or send a PR to
   upstream to enable Go module support
 - if the dependency has little transitive dependencies, use `.inject()` to
   add transitive dependencies manually after your `.collect()` call
 - otherwise, extend fietsje to allow for out-of-tree go.mod/go.sum files kept
   within monogon, or come up with some other solution.

Your new dependency might conflict with existing dependencies, which usually
manifests in build failures due to incompatible types. If this happens, you
will have to start digging to find a way to bring in compatible versions of
the two dependencies that are interacting with eachother. Do also mention any
such constraints in code comments near your `.collect()` call.

How to: update a high-level dependency
--------------------------------------

If you want to update a .collect()/.collectOverride() call, find out the
version you want to bump to and update it in the call. Re-running fietsje
will automatically update all enable transitive dependencies. Build and test
your code. Again, any possible conflicts will have to be resolved manually.

In the future, an audit flow might be provided for checking what the newest
available version of a high-level dependency is, to allow for easier,
semi-automated version bumps.

Version resolution conflicts
----------------------------

Any time a `.collect()`/`.collectOverride()` call is made, Fietsje will note
what transitive dependencies did the specified high-level dependency request.
Then, subsequent `.use()` calls will enable these dependencies in the build. On
subsequent `.collect()`/`.collectOverride()` calls, any transitive dependency
that already has been pulled in will be ignored, and the existing version will
be kept.

This means that Fietsje does not detect or handle version conflicts at a granular
level comparable to gomod. However, it does perform 'well enough', and in general
the Go ecosystem is stable enough that incompatibilites arise rarely - especially as
everything moves forward to versioned to go modules, which allow for multiple
incompatible versions to coexist as fully separate import paths.

It is as such the programmer's job to understand the relationship between imported
high-level dependencies. In the future, helper heuristics can be included that will
help understand and reason about dependency relationships. For now, Fietsje will just
help a user when they call `.use()` on the wrong dependency, ie. when the requested
transitive dependency has not been pulled in by a given high-level dependency.

