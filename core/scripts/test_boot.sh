#!/usr/bin/expect -f

# Getting the actual path from a sh_test rule is not straight-forward and would involve
# parsing the runfile at $RUNFILES_DIR, so just hardcode it.
#
# We'll want to replace this thing by a proper e2e testing suite sooner than we'll
# have to worry about cross-compilation or varying build environments.
#
# (see https://github.com/bazelbuild/bazel/blob/master/tools/bash/runfiles/runfiles.bash)
set kubectl_path "external/kubernetes/cmd/kubectl/linux_amd64_pure_stripped/kubectl"

set timeout 60

proc print_stderr {msg} {
  send_error "\[TEST\] $msg\n"
}

spawn core/scripts/launch.sh

expect "Network service got IP" {} default {
  print_stderr "Failed while waiting for IP address\n"
  exit 1
}

expect "Initialized encrypted storage" {} default {
  print_stderr "Failed while waiting for encrypted storage\n"
  exit 1
}

# Make an educated guess if the control plane came up
expect -timeout 3 "\n" {
  exp_continue
} timeout {} default {
  print_stderr "Failed while waiting for k8s control plane\n"
  exit 1
}

spawn $kubectl_path cluster-info dump -s https://localhost:6443 --username none --password none --insecure-skip-tls-verify=true

expect "User \"system:anonymous\" cannot list resource \"nodes\" in API group \"\" at the cluster scope" {} default {
  print_stderr "Failed while waiting for encrypted storage\n"
  exit 1
}

print_stderr "Completed successfully"
exit 0
