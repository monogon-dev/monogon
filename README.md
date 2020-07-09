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

We check the entire .ijwb project directory into the repository, which requires everyone to use the latest
version of both IntelliJ and the Bazel plugin, but eliminates manual setup steps.

The following steps are necessary:

- Install Google's official Bazel plugin in IntelliJ.

- Add the absolute path to your ~/.cache/bazel-nxt folder to your idea64.vmoptions (Help → Edit Custom VM Options)
  and restart IntelliJ:

  `-Dbazel.bep.path=/home/leopold/.cache/bazel-nxt`
  
- Set "*Bazel Binary Location*" in Other Settings → Bazel Settings to the absolute path of scripts/bin/bazel.
  This is a wrapper that will execute Bazel inside the container.

- Open the `.ijwb` folder as IntelliJ project.
  
- Disable Vgo support for the project.

- Run a non-incremental sync in IntelliJ 

The plugin will automatically resolve paths for generated files.

If you do not use IntelliJ, you need to use the scripts/bazel_copy_generated_for_ide.sh script to copy files locally.
