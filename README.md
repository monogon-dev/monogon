# Monogon Monorepo

This is the main repository containing the source code for the [Monogon Project](https://monogon.tech).

*This is pre-release software - feel free to look around, and check back later for our first release!*

## Environment

Our build environment requires a working Podman binary (your distribution should have one).

#### Usage

Spinning up: `scripts/create_container.sh` 

Spinning down: `scripts/destroy_container.sh` 

Running commands: `scripts/run_in_container.sh <...>`

Using bazel using a wrapper script: `scripts/bin/bazel <...>` (add to your local $PATH for convenience)

#### IntelliJ

This repository is compatible with the IntelliJ Bazel plugin, which enables
full autocompletion for external dependencies and generated code. All commands
run inside the container, and necessary paths are mapped into the container.

The following steps are necessary:

- Install Google's [Bazel plugin](https://plugins.jetbrains.com/plugin/8609-bazel) in IntelliJ. On IntelliJ 2020.3 or later,
  you need to install a [beta release](https://github.com/bazelbuild/intellij/issues/2102#issuecomment-801242977) of the plugin.

- Add the absolute path to your `~/.cache/bazel-monogon` folder to your `idea64.vmoptions` (Help → Edit Custom VM Options)
  and restart IntelliJ:

  `-Dbazel.bep.path=/home/leopold/.cache/bazel-monogon`
  
- Set "*Bazel Binary Location*" in Other Settings → Bazel Settings to the absolute path of `scripts/bin/bazel`.
  This is a wrapper that will execute Bazel inside the container.
  
- Use _File → Import Bazel project_... to create a new project from `.bazelproject`.

After running the first sync, everything should now resolve in the IDE, including generated code.

## Metropolis

### Run a single node cluster

Launch the node:

    scripts/bin/bazel run //:launch -c dbg
    
Run a kubectl command:

    scripts/bin/bazel run //metropolis/cli/dbg -c dbg -- kubectl describe
 
Run tests:

    scripts/bin/bazel test -c dbg //...
