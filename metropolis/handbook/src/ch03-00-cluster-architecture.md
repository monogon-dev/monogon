Introduction
---

Each Metropolis deployment (Cluster) is fully self-contained and independent from other Clusters. 

A Cluster is made up of Nodes. Nodes are machines (be it physical or virtual) running an instance of Metropolis. A Node can be part of only one Cluster.


```
             Cluster
 .-----------._.'._.------------.
 '                              '
 .--------. .--------. .--------.
 | Node A | | Node B | | Node C |
 '--------' '--------' '--------'
```

Nodes
---

Each Node runs the Linux kernel and Metropolis userspace. The userspace is comprised of Metropolis code on a signed read-only partition, and of persistent user data on an encrypted read-write partition. The signed read-only filesystem (the System filesystem) is verified by the Linux kernel, which in turn is signed and verified by a Node's firmware (EFI) via Secure Boot.

```
          
.--------------------.         .--------------------.
| Platform Firmware  |-------->| Secure Boot Image  |
|       (EFI)        | Checks  |--------------------|        
|--------------------|         |  .--------------.  |        .-------------------.
|      PK/KEK        |         |  | Linux Kernel |---------->| System FS (erofs) |
| Signature Database |         |  |--------------|  | Checks |-------------------|
'--------------------'         |  |  System FS   |  |        |    Node Code      |
                               |  |  Signature   |  |        '-------------------'
                               |  |  (dm-verity) |  |        
                               |  '--------------'  |        
                               '--------------------'

```

When booting, a Node needs to become part of a cluster (by either Bootstrapping a new one, Registering into an existing one for the first time, or Joining after reboot) to gather all the key material needed to mount the encrypted data partition. One part of the key is stored on the EFI System Partition encrypted by the TPM (sealed), and will only decrypt correctly if the Node's Secure Boot settings have not been tampered with. The other part of the key is stored by the Cluster, enforcing active communication (and possibly hardware attestation) with the Cluster before a Node can boot.

```
.-------------------.  Measures Secure Boot settings
| Platform Firmware |<----------.
'-------------------'           |
         | Checks               |
         v                      |
.-------------------.           |
| Secure Boot Image |           |
'-------------------'           |
         | Checks           .-------.
         v                  |  TPM  |
.-------------------.       '-------'
|     System FS     |           |
'-------------------'           | Seals/Unseals
         | Mounts               v
         |           .---------------------.        .------------------------.
         | .---------| Node Encryption Key |        |    Running Cluster     |
         |/          '---------------------'        |------------------------|
         | .----------------------------------------| Cluster Encryption Key |
         |/                                         |       (per node)       |
         |                                          '------------------------'
         v
.---------------------------.
| Data Partition (xfs/LUKS) |
'---------------------------'
 
```

The Node boot, disk setup and security model are described in mode detail in the [Node](ch03-01-node.md) chapter.

Each Node has the same minimal userland implemented in Go. However, this userland is unlike an usual GNU/Linux distribution, or most Linux-based operating systems for that matter. Metropolis does not have an LSB-compliant filesystem root (no /bin, /etc...) and does not run a standard init system / syslog. Instead, all process management is performed within a supervision tree (where supervised processes are called Runnables), and logging is performed within that supervision tree.

The supervision tree and log tree have some strict properties that are unlike a traditional Unix-like init system. Most importantly, any time a runnable restarts due to some unhandled error (or when it explicitly exits), all subordinate runnables will also be restarted.

In a more practical example, when working with Metropolis, you will see log messages like the following:

```
root.enrolment                   I0228 13:30:45.996960 cluster_bootstrap.go:48] Bootstrapping: mounting new storage...
root.network.interfaces          I0228 13:30:45.997359 main.go:252] Starting network interface management
root.time                        R 2022-02-28T13:30:45Z chronyd version 4.1-monogon starting (NTP RTC SCFILTER ASYNCDNS)
root.network.interfaces.dhcp     I0228 13:30:46.006082 dhcpc.go:632] DISCOVERING => REQUESTING
root.network.interfaces.dhcp     I0228 13:30:46.008871 dhcpc.go:632] REQUESTING => BOUND
```

The first column represents a runnable's Distinguished Name. It shows, for example, that the `DISCOVERING => REQUESTING` log line was emitted by a supervision tree runnable named `dhcp`, which was spawned by another runnable named `interfaces`, which in turn was spawned by a runnable named `network`, which in turn was started by the root of the Metropolis Node code.

The Node runnables axioms, supervision tree and log tree are described in more detail in the [Node Runnables and Logging](ch03-02-node-runnables.md) chapter.

Node roles and control plane
---

Each Node has a set of Roles assigned to it. These roles include, for example running the cluster control plane, running Kubernetes workloads, etc. At runtime, Nodes continuously retrieve the set of roles assigned to them by the cluster and maintain services which are required to fulfill the roles assigned to them. For example, if a node has a 'kubernetes worker' role, it will attempt to run the Kubelet service, amongst others.

```

   .-----------------------.
   | Cluster Control Plane |
   |-----------------------|
   |  Node Configuration   |
   |    & Node Status      |
   '-----------------------'
 Assigned   |      ^ Status
    roles   v      | updates
         .------------.
         |   Node A   |
         |------------|
         |            |
         |  Kubelet   |
         |            |
         '------------'
     
```

Nodes which have the 'control plane' role run core cluster services which other nodes depend on. These services make up a multi-node consensus which manages cluster configuration and management state. This effectively makes the cluster self-managed and self-contained. That is, the control plane block pictured above is in fact running on nodes in the same way as the Kubelet.

```

.---------------. .---------------. .---------------.
|    Node A     | |     Node B    | |    Node C     |
|---------------| |---------------| |---------------|
| Control Plane | | Control Plane | |       Kubelet |
| ^             | | ^     Kubelet | |               |
'-|-------------' '-|-------------' '---------------'
  |       |         |       |                  |
  '-------+---------+-------+------------------'
           Assigned roles & Status updates
```

The control plane services are described in more detail in the [Cluster](ch03-04-cluster.md) chapter.

The Control Plane services serve requests from Nodes (like the aforementioned retrieval of roles) and Users/Operators (like management requests) over gRPC, via an API named [Cluster API](ch03-05-cluster-api.md).

Identity & Authentication
---

When Nodes or Users/Operators contact the Cluster API, they need to prove their identity to the Node handling the request. In addition, nodes handling these requests need to prove their identity to the client. This is performed by a providing both sides of the connection with TLS certificates, and with some early communication (when certificates are not yet available) being performed over self-signed certificates to prove ownership of a key.

The TLS Public Key Infrastructure (CA and certificates) is fully self-managed by the Cluster Control Plane, and Users or Operators never have access to the underlying private keys of nodes or the CA. These keys are also stored encrypted within the Node's data partition, so are only available to nodes that have successfully become part of the Cluster. This model is explained and documented further in the [Identity and Authentication](ch-03-06-identity-and-authentication.md) chapter.

In the future, we plan to implement TPM-based Hardware Attestation as part of the early connections of a Node to a Cluster, allowing full cross-node verification, and optionally connections from a User/Manager to a Cluster.

