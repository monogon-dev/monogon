# `//third_party`

Anything related to importing third party code and making it work with our monorepo lives here:

- Vendored code.
- Bazel rules for third party code.
- Patches to thid party code.
- Lock files and similar package manager definitions (whenever possible - things like `go.mod  have to live at the top level, and that's OK too).

First-party code or build rules based *on* a third-party component should not live here.
