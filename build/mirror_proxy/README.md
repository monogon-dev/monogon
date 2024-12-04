Bazel downloader mirror
===

This is a small tool which acts as a transparent proxy-ish mirror for use in the bazel downloader.
By using a bazel_downloader.cfg we can instruct bazel to rewrite the download URLs and use a custom target instead. We use this to mirror all dependencies to our S3 storage.

Usage
---

This is expected to run with a given bucket name and a hardcoded set of credentials which are used to authenticate requests. When an authenticated request is received, the mirror will download uncached data if it isn't in the cache yet. This is expected to be used by trusted users, e.g. employees.

Users should deploy a .netrc inside their home folder based on the following template to allow bazel to authenticate against the mirror.

`~/.netrc`
```
machine mirror.monogon.dev
login myfancyusername
password mysecretpassword
```