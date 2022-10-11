BMaaS Server
===

Background
---

This server provides an interface to the BMaaS database/state over a gRPC API. Most components of the BMaaS system talk to the database directly whenever possible. Everything else communicates through this server.

Currently this is:

1. Agents running on machines, as they should only be allowed to access/update information about the machine they're running on, and they're generally considered untrusted since they run on end-user available machines.

In the future this server will likely also take care of:

1. A debug web API for developers/administrators to inspect database/BMDB state.
2. Periodic batch jobs across the entire BMDB, like consistency checks.
3. Exporting BMDB state into monitoring systems.
4. Coordinating access to the BMDB systems if the current direct-access-to-database architecture stops scaling.

Running
---

    bazel run //cloud/bmaas/server/cmd -- -srv_dev_certs -bmdb_eat_my_data

Although that's not very useful in itself currently. Instead, most functionality is currently exercised through automated tests.

TODO(q3k): document complete BMaaS dev deployment (multi-component, single BMDB).