# Monogon coding standards

## Programming languages

### Go
While working on a change, please adhere to practices documented in the chapter about [frequently occurring code review comments](https://github.com/golang/go/wiki/CodeReviewComments) of [the Go wiki](https://github.com/golang/go/wiki/). In addition, some Monogon-specific style is enforced at build time as follows:

#### Line Length

Comment lines must not be longer than 80 characters. Non-comment lines may be any (reasonable) length.

#### Imports

Imports must be grouped and sorted in the following way:

```go
import (
    "errors"
    "net/http"

    "example.com/some/external/thing"
    "golang.org/x/crypto"

    "source.monogon.dev/some/internal/library"
    apkg "source.monogon.dev/some/other/pkg"
)
```

Ie., all stdlib imports must come before all non-Monogon ('global') imports, which must come before all Monogon ('local') imports. Within each group, imports must be sorted. 

There can be multiple groups of a given class, but they must be in the right order:

```go
import (
    "errors"
    "net/http"

    "source.monogon.dev/some/internal/library"
    "source.monogon.dev/some/internal/pkg"

    "source.monogon.dev/other/subtree/a"
    "source.monogon.dev/other/subtree/b"
    foo "source.monogon.dev/other/subtree/c"
)
```

A styleguide compliant fork of `goimports` (itself a superset of `gofmt`)  can be built by running:

    $ bazel build //:goimports

The resulting binary can then be copied to anywhere in the filesystem (eg. $HOME/bin/goimports-monogon) and any editor which supports gofmt/goimports integration can be pointed at this tool to automatically reformat files to the required format.

When setting up integration with a text editor (or calling the binary manually), you must make it so that goimports get called `-local source.monogon.dev/`. Otherwise goimports will not correctly split away local/Monogon (`source.monogon.dev`) imports from other (global) imports.
