CREATE TABLE machine_os_installation_request (
    machine_id UUID NOT NULL REFERENCES machines(machine_id) ON DELETE RESTRICT,
    -- Version of this request, for example monotonic epoch counter. Used to
    -- match successful installation (represented by OS installation report) to
    -- pending request, making sure that we don't perform spurious re-installs.
    generation INT NOT NULL,
    -- Serialized cloud.bmaas.server.api.OSInstallationRequest.
    os_installation_request_raw BYTEA NOT NULL,
    CONSTRAINT "primary" PRIMARY KEY (machine_id)
);

CREATE TABLE machine_os_installation_report (
    machine_id UUID NOT NULL REFERENCES machines(machine_id) ON DELETE RESTRICT,
    -- Matches generation in machine_os_installation_request. Not constrained on
    -- purpose, as a mismatch between generations implies an actionable
    -- installation request and is a valid state of the system.
    generation INT NOT NULL,
    CONSTRAINT "primary" PRIMARY KEY (machine_id)
);
