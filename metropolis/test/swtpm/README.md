swtpm enhancements
==================

Metropolis uses [swtpm](https://github.com/stefanberger/swtpm) for emulating a
TPM device when running tests in qemu, eg. end-to-end-tests.

swtpm consists of a runtime emulator (`swtpm`) which runs against a state
directory and exposes TPM functionality over the socket; and of tooling
designed to create said state directory (`swtpm_setup`, `swtpm_localca`, etc).

Getting the former to be built with Bazel is generally trivial, as it mostly
depends on libraries we are already building (glib, openssl/boringssll, etc).
However, the tooling is another story: it depends heavily on GnuTLS, both as a
library to link against and as a runtime tool (`certtool`). We already have one
C implementation of cryptographic primitives in `//third_party` (boringssl),
dragging another one in would be shameful.

The tooling is also not a single C binary, but a handful of different ones that
call eachother based on the requested functionality (presumably as a way to
implement modularity to allow creating swtpm secrets using a HSM, etc).

This subdirectory contains bits and pieces that allow us to use the
aforementioned tooling without depending on GnuTLS. This is done by patching
some tools to rip out GnuTLS support, and by replacing other with native Go
reimplementations.

swtpm_cert
----------

This is a reimplementation of swtpm_cert in Go. The upstream swtpm_cert is implemented in C and has a hard dependency on
GnuTLS and libtasn1. Rewriting it in Go and using plain stdlib functions seems like the correct solution here (the
alternative being either bringing in GnuTLS/libtasn1 into `third_party`, or rewriting swtpm_cert to use
OpenSSL/BoringSSL).

certtool
--------

This is a minimal GnuTLS certtool reimplementation in Go. It's used by `swtpm_localca` to generate TLS certificates. An
alternative to this would be to rewrite `swtpm_localca` entirely to Go, but that seems like a bit too much effort for
now.