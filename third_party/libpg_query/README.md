libpg\_query
===

This library provides a C API to parse postgres queries. It consists of some vendored PostgreSQL source code and wrapper header/sources.

Licensing
---

 * libpg\_query itself: BSD 3-clause
 * xxhash: BSD 2-clause
 * protobuf-c: BSD 2-clause (not named, but terms are equal)
 * PostgreSQL: PostgreSQL license (similar to MIT)

Known Issues
---

This library has a very wide include path requirement, that includes its own vendor directories (which contain postgres, xxhash and protobuf-c). These includes are pulled into all dependendents of this library and might break anything that wants eg. both libpg\_query and xxhash. When this happens, we should patch the library to always use absolute includes instead, thereby cleaning up the include directives. We technically have `bazel_cc_fix` for that.

We could also unvendor xxhash, protobuf-c and even postgres. But that might not be worth the effort right now.
