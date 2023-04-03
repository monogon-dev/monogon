sandboxroot mirror
===

Fedora mirrors tend to drop RPMs very quickly. As we don't want to be constantly
chasing every single tiny update, we have decided to set up our own mirror on GCS.

The mirror only contains RPMs that the sandboxroot actually uses, and is managed
by running the `mirror` tool from this directory.

Using the mirror
---

The mirror is enabled by default whenever you use Bazel (see repositories.bzl in this directory).

Updating the mirror
---

Any time you run `third_party/sandboxroot/regenerate.sh`, the last step calls `mirror sync`. If that fails for some reason (eg. you were not logged into GCS), you can run it manually:

```
$ bazel run :mirror sync
```

Checking the mirror
---

If you want to just check whether everything's properly synced, you can run:

```
$ bazel run :mirror check
```

To do a full scan (downloading and checking SHA256 sums) do:

```
$ bazel run :mirror check --deep
```
