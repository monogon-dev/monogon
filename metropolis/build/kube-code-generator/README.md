kube-code-generator
===================

A small Bazel rule library for dealing with k8s.io/code-generators.

See defs.bzl for documentation, and `//metropolis/vm/kube/apis` for an example of usage.

Current Limitations
-------------------

 - Clientset-gen's `versioned/fake` is not generated.
 - Only the following generators are ran: deepcopy, clientset, informer, lister.
 - Bazel BUILDfiles for the generated structure must be crafted manually.
 - Go packages must follow upstream format (group/version). This influences
   Bazel target structure, which can then look somewhat awkward in a
   project-oriented monorepo (eg. //foo/bar/widget/kube/apis/widget/v1 has a
   'widget' stutter.
