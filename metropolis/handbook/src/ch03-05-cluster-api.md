# Cluster API

This chapter goes through all of the services and data types exposed by Cluster API that's used in day-to-day operations. In most cases, it's easier to access this functionality with metroctl tool. However, it still can be used directly whenever custom implementation is needed, eg. when writing automation. The rule of thumb is that metroctl should be used for human interaction, while the direct API access should be used for machine-to-machine communication.

## Services

Most cluster services require authenticated access (in the form of a TLS keypair and certificate), which can be obtained using cluster owner's credentials generated during cluster bootstrapping phase. Escrow service is the single exception to this rule.

### Escrow

Escrow is extensively documented in its [protofile](/metropolis/proto/api/aaa.proto). It's current usecase is exchanging the Initial Owner Credentials (generated before the installation of the first node) into long-term access credentials that are then used to perform further API accesses.

### Management

The Management service is the main cluster manager-facing cluster management API. Management tasks include:
- Addition and removal of cluster nodes
- Querying of cluster and node status
- Node configuration
- Node and cluster bootstrapping

This chapter describes Management service calls together with their expected usage scenarios. For a more technical description, see: [management.proto](/metropolis/proto/api/management.proto)

#### GetRegisterTicket

GetRegisterTicket is used in cluster nodes' Register Flow. An up-to-date register ticket has to be retrieved ahead of time by a cluster manager, and included in Node Parameters supplied to the new node for it to be able to register and join the cluster. The ticket is used to protect the API surface from potential denial of service attacks by limiting the amount of entities that can start the Register Flow against the cluster. It can be regenerated at any time in case it leaks.

#### GetClusterInfo

GetClusterInfo returns summary information about the cluster, currently made up of the Cluster Directory containing node network addresses bundled with node public keys, which can be used to uniquely identify particular nodes.

#### GetNodes

GetNodes retrieves detailed information about cluster member nodes, such as their active roles and health. The call's output can be limited to nodes of interest with a [CEL](opensource.google.com/projects/cel) expression.

The filter expressions operate on [Node protobuf messages](/metropolis/proto/api/management.proto). Here's a couple of examples:
- node.state == NODE_STATE_UP
- has(node.roles.kubernetes_worker)
- node.time_since_heartbeat < duration('6s')

#### ApproveNode

ApproveNode is part of the Register Flow. It's called to admit new nodes into the cluster. For the call to succeed, the target node must have already registered into the cluster using a valid Register Ticket.

#### UpdateNodeRoles

UpdateNodeRoles updates node's roles within the cluster. Currently role separation *is not* implemented. However, in upcoming releases it will become possible to specialize nodes into a cluster consensus member and/or a Kubernetes worker node.

### Node debug

The purpose of this service is to ease the development process of Metropolis. It's disabled in non-debug builds.

See also: [debug.proto](/metropolis/proto/api/debug.proto)

## Authorization

Even though currently the RPC authorization policy is permissive, meaning that cluster managers can access all available calls, this will change in favor of a fine-grained approach. See: [authorization.proto](/metropolis/proto/ext/authorization.proto).

## Configuration datatypes

These datatypes, while not used as part of any remote procedure call, serve as a medium carrying cluster configuration.

See also: [configuration.proto](/metropolis/proto/api/configuration.proto).

### Node Parameters

Node Parameters are prepared for newly installed nodes as part of the installation flow. Depending on their contents, the node will either bootstrap a new Cluster, or join an existing one.
