# BIOS Bootcode

This package contains legacy bootcode which is displayed to non-UEFI users.
It's sole purpose is to explain users their wrongdoing and tell them to use UEFI.
It also shows a cute ascii-art logo.

## Build
 
Bazel generates the logo content with `genlogo`.
It takes a black/white png-file and converts it to RLE encoded data,
which is rendered as ascii-art at runtime.
