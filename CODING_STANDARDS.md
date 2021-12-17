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

    "dev.source.monogon/some/internal/library"
    apkg "dev.source.monogon/some/other/pkg"
)
```

Ie., all stdlib imports must come before all non-Monogon ('global') imports, which must come before all Monogon ('local') imports. Within each group, imports must be sorted. 

There can be multiple groups of a given class, but they must be in the right order:

```go
import (
    "errors"
    "net/http"

    "dev.source.monogon/some/internal/library"
    "dev.source.monogon/some/internal/pkg"

    "dev.source.monogon/other/subtree/a"
    "dev.source.monogon/other/subtree/b"
    foo "dev.source.monogon/other/subtree/c"
)
```
