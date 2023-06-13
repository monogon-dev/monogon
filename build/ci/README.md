Monogon CI
==========

Monogon has a work-in-progress continuous integration / testing pipeline.
Because of historical reasons, some parts of this pipeline are defined in a
separate non-public repository that is managed by Monogon SE.

In the long term, the entire infrastructure code relating to this will become
public and part of the Monogon repository. In the meantime, this document
should serve as a public reference that explains how that part works and how it
integrates with `//build/ci/...` and the project as a whole.

CI usage
--------

When a change on https://review.monogon.dev/ gets opened, it needs to either
be owned by a 'trusted user', or be vouched by one. This is because our current
CI setup is not designed to protect against malicious changes that might
attempt to take over the CI system, or change the CI scripts themselves to skip
tests.

Currently, all Monogon SE employees (thus, the core Monogon development team)
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
not publicly available either. This will be fixed soon.
