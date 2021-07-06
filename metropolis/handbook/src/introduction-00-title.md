# Metropolis, a cluster operating system

> Note: Metropolis is currently in **heavy development**. This documentation is written to *reflect our goals*, not necessarily the current state of the product. You are welcome to give Metropolis a try, but we cannot recommend running it anywhere near production workloads.

Welcome to the *Metropolis Handbook*, the primary documentation resource for Metropolis. Metropolis is a cluster operating system, meaning its goal is to run on a fleet of machines (be it physical or virtual) and pool their resources together into a unified API for operations and developer teams.

Metropolis stands on the shoulders of giants, and takes the best of battle-tested software like the Linux kernel and Kubernetes to build a cohesive, stable, reliable and secure platform.

## What makes Metropolis unique

 1. **A self-contained operating system**: Metropolis is a full software stack, including the Linux kernel, userspace code, Kubernetes distribution and cluster management system. In contrast to traditional cluster administration, there are no puzzles to put together from a dozen vendors. The entire stack is tested as a single deployable unit.
 1. **Eliminates state**: Metropolis nodes don't have a traditional read-write filesystem, all of their state is contained on a separate partition with clear per-component ownership of data. All node configuration is managed declaratively on a per-node basis, and all cluster operations are all done by gRPC API.
 1. **No shell, no one-off hacks, no configuration drift**: Metropolis nodes do not run SSH nor depend on low-level system administration tools for day-to-day operations, even debugging.
 1. **Opinionated on production readiness**: Metropolis does not attempt to support every possible software configuration, instead focusing on scenarios that make for a high quality production experience .
 1. **Robust**: Metropolis builds upon proven technology and does not take risks. Cluster consensus is maintained using the Raft protocol, user and node communication use well-defined gRPC services, while system services are limited in complexity and purpose-built for Metropolis.
 1. **Secure at rest**: Metropolis nodes by default encrypt their data partitions and check the integrity of running code, providing tamper resistance and preventing data exfiltration even if an attacker can access a node's disk drives.
 1. **Self-locking**: Metropolis can be configured to use TPM hardware attestation, in which cluster membership is limited to nodes that are running trusted versions of the software on trusted hardware.
 1. **Not magic**: Metropolis clusters are complex, distributed systems. Managing any distributed system like Metropolis requires some knowledge of core concepts and components involved, and the Metropolis does not attempt to hide that complexity away. Limited internal abstractions and well documented source code lets anyone easily troubleshoot any deeper issues.

## Kubernetes on Metropolis

While we aim to make Metropolis run various kinds of workloads in the future, Metropolis strongly focuses on using Kubernetes as an application platform for users. Workloads can be scheduled on Metropolis using any Kubernetes tools like kubecfg, Tanka or even Helm.

In comparison to other Kubernetes distributions, *Metropolis does not attempt to simplify Kubernetes* by providing extra wrappers or shortcuts for users. Instead, we believe that users should understand the Kubernetes production model and aim to be proficient in its API, as any high-level wrappers only paradoxically introduce complexity.

