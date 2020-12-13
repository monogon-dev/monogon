# Nexantic monorepo

This is the monorepo storing all of nexantic's internal projects and libraries.

## Environment

We assume a Fedora host system provisioned using rW, and IntelliJ as the IDE.

For better reproducibility, all builds are executed in containers.

#### Usage

Spinning up: `scripts/create_container.sh` 

Spinning down: `scripts/destroy_container.sh` 

Running commands: `scripts/run_in_container.sh <...>`

Using bazel using a wrapper script: `scripts/bin/bazel <...>` (add to your local $PATH for convenience)

#### Run a single node cluster

Launch the node:

    bazel run //:launch
    
Run a kubectl command:

    bazel run //core/cmd/dbg -- kubectl describe
 
#### IntelliJ

This repository is compatible with the IntelliJ Bazel plugin. All commands run inside the container, and
necessary paths are mapped into the container.

The following steps are necessary:

- Install Google's [Bazel plugin](https://plugins.jetbrains.com/plugin/8609-bazel) in IntelliJ.

- Add the absolute path to your `~/.cache/bazel-nxt` folder to your `idea64.vmoptions` (Help → Edit Custom VM Options)
  and restart IntelliJ:

  `-Dbazel.bep.path=/home/leopold/.cache/bazel-nxt`
  
- Set "*Bazel Binary Location*" in Other Settings → Bazel Settings to the absolute path of `scripts/bin/bazel`.
  This is a wrapper that will execute Bazel inside the container.
  
- Use _File → Import Bazel project_... to create a new project from `.bazelproject`.

After running the first sync, everything should now resolve in the IDE, including generated code.
