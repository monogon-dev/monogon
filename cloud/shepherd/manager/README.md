Equinix Shepherd
===

Manages Equinix machines in sync with BMDB contents. Made up of two components:

Provisioner
---

Brings up machines from hardware reservations and populates BMDB with new Provided machines.

Initializer
---

Starts the Agent over SSH (wherever necessary per the BMDB) and reports success into the BMDB.


Running
===

Unit Tests
---

The Shepherd has some basic smoke tests which run against a Fakequinix.

Manual Testing
---

If you have Equinix credentials, you can run:

```
$ bazel build //cloud/shepherd/provider/equinix
$ bazel build //cloud/shepherd/manager/test_agent
$ bazel-bin/cloud/shepherd/provider/equinix/equinix_/equinix \
    -bmdb_eat_my_data \
    -equinix_project_id FIXME \
    -equinix_api_username FIXME \
    -equinix_api_key FIXME \
    -agent_executable_path bazel-bin/cloud/shepherd/manager/test_agent/test_agent_/test_agent \
    -agent_endpoint example.com \
    -equinix_ssh_key_label $USER-FIXME \
    -equinix_device_prefix $USER-FIXME- \
    -provisioner_assimilate -provisioner_max_machines 10
```

Replace $USER-FIXME with `<your username>-test` or some other unique name/prefix.

This will start a single instance of the provisioner accompanied by a single instance of the initializer.

A persistent SSH key will be created in your current working directory.

Prod Deployment
---

TODO(q3k): split server binary into separate provisioner/initializer for initializer scalability, as that's the main bottleneck.