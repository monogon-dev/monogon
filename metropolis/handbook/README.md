Metropolis Handbook
===

This directory contains the sources of the Metropolis Handbook end-user documentation.

Layout
---

Everything within `//monogon/handbook/src` will be used to generate documentation with [mdbook](https://rust-lang.github.io/mdBook/format/index.html).

Compared to upstream mdbook we do not have a static `book.toml` file, one is instead generated as part of the build process. See the definition of the `//metropolis/handbook` target to change some of the options.

Building
---

    bazel build //metropolis/handbook

Then, you can visit the following file in your browser:

    bazel-bin/metropolis/handbook/handbook/index.html

To view the built documentation.

Interactive editing
---

For faster edit/check loops of the handbook, you can use `ibazel`:

    ibazel build //metropolis/handbook

This will automatically rebuild the handbook any time some source changes.

You will still need to manually refresh your browser to see any changes. This could be made better, if needed, by injecting some [ibazel-compatible live reload javascript](https://github.com/bazelbuild/bazel-watcher/blob/84cab6f15f64850fb972ea88701e634c8b611301/example_client/example_client.go#L24) to automatically reload the page on changes, or by adding a target which launches `mdbook serve`.

Publishing
---

We currently do not build the handbook automatically and/or publish it anywhere.
