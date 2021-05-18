Monogon CI
==========

Monogon has a work-in-progress continuous integration / testing pipeline.
Because of historical reasons, some parts of this pipeline are defined in a
separate non-public repository that is managed by Monogon Labs.

In the long term, the entire infrastructure code relating to this will become
public and part of the Monogon repository. In the meantime, this document
should serve as a public reference that explains how that part works and how it
integrates with `//build/ci/...` and the project as a whole.

Builder Image & Container
-------------------------

`//build/ci/Dockerfile` describes a 'builder image'. This image contains a
stable, Fedora-based build environment in which all Monogon components should
be built. It has currently two uses:

1. The build scripts at
   `//scripts/{create_container.sh,destroy_container.sh,/bin/bazel}`. These are
   used by developers to run Bazel against a controlled environment to develop
   Monogon code. The `create_container.sh` script builds the Builder image and
   starts a Builder container. The `bin/bazel` wrapper script launches Bazel in
   it. The `destroy_container.sh` script cleans everything up.

2. The Jenkins based CI uses the Builder image as a base to run Jenkins agents.
   A Monogon Labs developer runs `//build/ci/build_ci_image`, which builds the
   Builder Image and pushes it to a container registry. Then, in another
   repository, that image is used as a base to overlay a Jenkins agent on top,
   and then used to run all Jenkins actions.

As Monogon evolves and gets better build hermeticity using Bazel toolchains,
the need for a Builder image should subdue. Meanwhile, using the same image
ensures that we have the maximum possible reproducibility of builds across
development and CI machines, and gets us a base level of build hermeticity and
reproducibility.

CI usage
--------

When a change on https://review.monogon.dev/ gets opened, it needs to either
be owned by a 'trusted user', or be vouched by one. This is because our current
CI setup is not designed to protect against malicious changes that might
attempt to take over the CI system, or change the CI scripts themselves to skip
tests.

Currently, all Monogon Labs employees (thus, the core Monogon development team)
are marked as 'trusted users'. There is no formal process for community
contributors to become part of this group, but we are more than happy to
formalize such a process when needed, or appoint active community contributors
to this group. Ideally, though, the CI system should be rebuilt to allow any
external contributor to run CI in a secure and sandboxed fashion.

CI implementation
-----------------

The CI system is currently made of a Jenkins instance running on
https://jenkins.monogon.dev/. It runs against open changes that have the
Allow-Run-CI label evaluated to 'ok' Gerrit Prolog rules, and executes the
`//build/ci/jenkins-presubmit.groovy` script on them.

Currently, the Jenkins instance is not publicly available, and thus CI logs are
not publicly available either. This will be fixed very soon.
