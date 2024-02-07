# Monogon Operating System

A cluster operating system. Codename: metropolis. Linux kernel, stateless userland, API-driven management, high
integrity. Designed to run Kubernetes and other workload scheduling systems.

## Documentation

The canonical documentation for Monogon OS is the [Monogon OS Handbook](https://docs.monogon.dev/metropolis-v0.1/handbook/index.html).


## Developer Quick Start

Follow the setup instructions in the top-level README.md. We recommend using `nix-shell` either via Nix installed on an existing distribution or on NixOS.

Start a test cluster by running:

```
$ bazel run //metropolis:launch-cluster
```

This will build all the required code and run a fully userspace test cluster consisting of four qemu VMs (three for
Monogon OS nodes, one for a swich/router emulator). No root access on the host required.

```
.--------. .--------. .--------.
| node 0 | | node 1 | | node 2 |
'--------' '--------' '--------'
     ^          ^          ^
     |          |          | Virtual Ethernet (10.1.0.0/24)
     V          V          V
   .-------------------------.
   | nanoswitch              |
   |-------------------------|    .-------------------.
   | Router, switch,         |--->| Internet via host |
   | SOCKS proxy             |    '-------------------'
   '-------------------------'
     ^
     | gRPC over SOCKS to nodes
   .----------.
   | metroctl |
   '----------'
```

The launch tool will output information on how to connect to the cluster:

```
Launch: Cluster running!
  To access cluster use: metroctl --config /tmp/metroctl-3733429479 --proxy 127.0.0.1:42385 --endpoints 10.1.0.2 --endpoints 10.1.0.3 --endpoints 10.1.0.4
  Or use this handy wrapper: /tmp/metroctl-3733429479/metroctl.sh
  To access Kubernetes, use kubectl --context=launch-cluster
```

You can use the metroctl wrapper to then look at the node list per the Monogon OS cluster control plane:

```
$ alias metroctl=/tmp/metroctl-3733429479/metroctl.sh
$ metroctl node describe
NODE ID                                       STATE   ADDRESS    HEALTH    ROLES                                  TPM   VERSION                         HEARTBEAT   
metropolis-067651202d00b79fffe92df0001aabff   UP      10.1.0.4   HEALTHY                                          yes   v0.1.0-dev494.g0d8a8a4f.dirty   1s          
metropolis-7ccd2437c50696ea9a9b6543dc163f84   UP      10.1.0.3   HEALTHY                                          yes   v0.1.0-dev494.g0d8a8a4f.dirty   3s          
metropolis-ec101152c48c5f761534c1910cf66200   UP      10.1.0.2   HEALTHY   ConsensusMember,KubernetesController   yes   v0.1.0-dev494.g0d8a8a4f.dirty   3s       
```

We have a node running the Monogon OS control plane (ConsensusMember role) and Kubernetes control plane (
KubernetesController role), but no Kubernetes worker nodes. But changing that is a simple API call (or metroctl
invocation) away:

```
$ metroctl node add role KubernetesWorker metropolis-067651202d00b79fffe92df0001aabff
2024/02/12 17:42:33 Updated node metropolis-067651202d00b79fffe92df0001aabff.
$ metroctl node add role KubernetesWorker metropolis-7ccd2437c50696ea9a9b6543dc163f84
2024/02/12 17:42:36 Updated node metropolis-7ccd2437c50696ea9a9b6543dc163f84.
$ metroctl node describe
NODE ID                                       STATE   ADDRESS    HEALTH    ROLES                                  TPM   VERSION                         HEARTBEAT   
metropolis-067651202d00b79fffe92df0001aabff   UP      10.1.0.4   HEALTHY   KubernetesWorker                       yes   v0.1.0-dev494.g0d8a8a4f.dirty   0s          
metropolis-7ccd2437c50696ea9a9b6543dc163f84   UP      10.1.0.3   HEALTHY   KubernetesWorker                       yes   v0.1.0-dev494.g0d8a8a4f.dirty   3s          
metropolis-ec101152c48c5f761534c1910cf66200   UP      10.1.0.2   HEALTHY   ConsensusMember,KubernetesController   yes   v0.1.0-dev494.g0d8a8a4f.dirty   3s     
```

And just like that, we can now see these nodes in Kubernetes, too:

```
$ kubectl --context=launch-cluster get nodes
NAME                                          STATUS   ROLES    AGE   VERSION
metropolis-067651202d00b79fffe92df0001aabff   Ready    <none>   15s   v1.24.2+mngn
metropolis-7ccd2437c50696ea9a9b6543dc163f84   Ready    <none>   13s   v1.24.2+mngn
$ kubectl --context=launch-cluster run -it --image=ubuntu:22.04 test -- bash
If you don't see a command prompt, try pressing enter.
root@test:/# uname -a
Linux test 6.1.69-metropolis #1 SMP PREEMPT_DYNAMIC Tue Jan 30 14:43:23 UTC 2024 x86_64 x86_64 x86_64 GNU/Linux
root@test:/# 

```

With the test launch tooling, you can now start iterating on the codebase. Regardless of whether you're changing the
Linux kernel config or implementing a new RPC, testing your changes interactively is a single `bazel` command away.

## End-to-end tests

We have an end-to-end test suite. It's run automatically on CI. Any new logic should be exercised there.

```
$ bazel run //metropolis/test/e2e:e2e_test
```

These tests operate on a fully virtualized cluster just like the launch tooling, so be patient.
