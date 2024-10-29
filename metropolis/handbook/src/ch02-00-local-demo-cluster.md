# The demo cluster

This chapter demonstrates the installation process and operation of a Metropolis cluster.

## Prerequisites

### Hardware
An x86_64/amd64 Linux host machine with glibc on which you can run `metroctl`. This will later be expanded to cover many more platforms, but for our first release this is the only supported platform.

And either:
* KVM support on your host machine and a hypervisor capable of running OVMF with TPM2 (like libvirt/virt-manager)
* A physical x86_68/amd64 machine (ideally at least 3 for reboot persistence) with UEFI boot and a TPM2 and a USB thumb drive (>=1G).

### Software

#### metroctl
First, you'll need *metroctl*, the command line utility for working with Metropolis clusters.
You can get it from GitHub Releases (https://github.com/monogon-dev/monogon/releases) with
```shell
curl -L -o metroctl https://github.com/monogon-dev/monogon/releases/download/metropolis-v0.1/metroctl
chmod +x metroctl
```
Optionally you can move the file to a location in PATH, like /usr/local/bin or ~/bin/.

#### The installation bundle

To install Metropolis, you'll need a *bundle*. A *bundle* contains all resources to install or update a Metropolis node.
You can get a prebuilt bundle from GitHub Releases with
```shell
curl -L -o bundle.zip https://github.com/monogon-dev/monogon/releases/download/metropolis-v0.1/bundle.zip
```

## Installation

### The bootstrap node

Let's generate the installer image that we'll use to install the first node of the upcoming cluster. To do that, use the *metroctl* tool in the following way:
```shell
metroctl install genusb bootstrap-node-installer.img --bootstrap --cluster=cluster.internal --bundle=<installation-bundle-path>
```
If you're going to install from a USB stick or other types of removable storage, supply metroctl with a device path:
```shell
metroctl install genusb /dev/sdx --bootstrap --cluster=cluster.internal --bundle=<installation-bundle-path>
```
Since a new GPT will need to be generated for the target device, the image file cannot simply be copied into it.
**Caution:** make sure you'll be using the correct path. *metroctl* will overwrite data on the target device.

Metropolis does not support installation from optical media.

The installer will be paired with your *cluster owner's credentials*, that *metroctl* will save to your XDG config directory. Please note that the resulting installer can be used only to set up the initial node.

If all goes well, this will leave you with the following output.
```
2022/07/07 10:29:44 Generating installer image (this can take a while, see issues/92).
```

Use the installer image to provision the first node. The image will contain an EFI-bootable payload.

**Caution:** the installer will install Metropolis onto the first suitable persistent storage it finds as soon as it boots. The installation process is non-interactive in this version of the OS. If you're going to install on physical hardware, make sure you have backed up all your data from the machine you'll be running it on.

If you'll be using a virtual machine, it is advised to pick smaller storage sizes, eg. 4G. Upon first boot, Metropolis will need to zero its data partition, which can take a long time.

The installer will produce the following output, that will be both sent over the serial interface, and displayed on your screen, if available:
```
Installing to /dev/vdx...
```

Afterwards, it will restart, and the installation media will need to be removed. At this point you should be left with a working bootstrap node.

### Taking ownership of the new cluster

After the first node is set up and running, you can take ownership of the upcoming cluster:
```shell
metroctl takeownership <bootstrap-node-address>
```
This should result in the following output being displayed:
```
2022/07/07 10:42:07 Successfully retrieved owner credentials! You now own this cluster. Setting up kubeconfig now...
2022/07/07 10:42:07 Success! kubeconfig is set up. You can now run kubectl --context=metropolis ... to access the Kubernetes cluster.
```

If this didn't work out the first time you tried, try giving the bootstrap node more time. Depending on available storage size, setting up its data partition might take longer, in which case your connection attempts will be refused.

### Additional nodes

Additional nodes can be provided with the non-bootstrap installer image. It can be generated with *metroctl*. This time, note the lack of the *--bootstrap* flag.
```shell
metroctl --endpoints <bootstrap-node-address> install genusb second-node-installer.img --bundle=<installation-bundle-path>
```

Complete the installation process with one or more nodes.

For the new nodes to join the cluster, you'll need to approve them first. Calling `metroctl approve` with no parameters will list nodes pending approval.
```shell
metroctl --endpoints <bootstrap-node-address> approve
```

You should get a list of node IDs which would look similar to:
```
metropolis-7eec2053798faab726bb9fd9e9444ec9
```

If there are no nodes that have already registered with the cluster, *metroctl* will produce the following output:
```
There are no nodes pending approval at this time.
```

To approve a node, use its node ID as a parameter.
```shell
metroctl --endpoints 192.168.122.238 approve <above-node-id-goes-here>
```

If the node was approved as a result, metroctl will say:
```
Approved node <node-id>
```

## Using the cluster

At this point you can start exploring Metropolis. Try playing with *kubectl*, or take a look at the [Cluster API](https://github.com/monogon-dev/monogon/blob/main/metropolis/handbook/src/ch03-05-cluster-api.md) chapter of this handbook.

The cluster state should be reflected by *kubectl* output:
```shell
kubectl --context=metroctl get nodes

NAME                                          STATUS   ROLES    AGE   VERSION
metropolis-4fb5a2aa4eec34080bea02ac8020028d   Ready    <none>   98m   v1.24.0+mngn
metropolis-7eec2053798faab726bb9fd9e9444ec9   Ready    <none>   86m   v1.24.0+mngn
```

If you need to install kubectl, try [this chapter](https://kubernetes.io/docs/tasks/tools/install-kubectl-linux/) of the official Kubernetes Documentation.

## Caveats

This is a preview version of Metropolis, and there's a couple of things to be aware of.

### Cold start

The cluster recovery flow is still unimplemented. This means that a *cold* cluster, in which all member nodes have been shut down, **will not** start up again. This will be solved in an upcoming release.
