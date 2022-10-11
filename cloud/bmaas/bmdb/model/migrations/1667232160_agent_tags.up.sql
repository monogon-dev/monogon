CREATE TYPE provider AS ENUM (
    'Equinix'
    -- More providers will follow in subsequent migrations.
);

-- tag MachineProvided {
--     Provider Provider
--     ProviderID String
-- }
-- Represents the fact that a machine is backed by a machine from a given
-- provider identified there with a given provider id.
CREATE TABLE machine_provided (
    machine_id UUID NOT NULL REFERENCES machines(machine_id) ON DELETE RESTRICT,
    provider provider NOT NULL,
    provider_id STRING(128) NOT NULL,
    CONSTRAINT "primary" PRIMARY KEY (machine_id),
    UNIQUE (provider, provider_id)
);

-- tag AgentStarted {
--     StartedAt time.Time
--     PublicKey []byte
-- }
-- Represents the fact that a machine has had the Agent started on it at some
-- given time, and that the agent returned a given public key which it will use
-- to authenticate itself to the bmdb API server.
CREATE TABLE machine_agent_started (
    machine_id UUID NOT NULL REFERENCES machines(machine_id) ON DELETE RESTRICT,
    agent_started_at TIMESTAMPTZ NOT NULL,
    agent_public_key BYTES NOT NULL,
    CONSTRAINT "primary" PRIMARY KEY(machine_id)
);

-- tag AgentHeartbeat {
--     At time.Time
-- }
-- Represents a successful heartbeat send by the Agent running on a machine at
-- some given time.
CREATE TABLE machine_agent_heartbeat (
    machine_id UUID NOT NULL REFERENCES machines(machine_id) ON DELETE RESTRICT,
    agent_heartbeat_at TIMESTAMPTZ NOT NULL,
    CONSTRAINT "primary" PRIMARY KEY(machine_id)
);

-- tag HardwareReport {
--     Raw []byte
-- }
-- Represents a hardware report submitted by an Agent running on a machine.
-- Usually a report is submitted only once after an agent has been started.
CREATE TABLE machine_hardware_report (
    machine_id UUID NOT NULL REFERENCES machines(machine_id) ON DELETE RESTRICT,
    hardware_report_raw BYTES NOT NULL,
    CONSTRAINT "primary" PRIMARY KEY(machine_id)
);

-- Used by the Shepherd when performing direct actions against a machine.
ALTER TYPE process ADD VALUE IF NOT EXISTS 'ShepherdInstall';