-- name: MachineAddProvided :exec
INSERT INTO machine_provided (
    machine_id, provider, provider_id
) VALUES (
    $1, $2, $3
);

-- name: MachineSetAgentStarted :exec
INSERT INTO machine_agent_started (
    machine_id, agent_started_at, agent_public_key
) VALUES (
    $1, $2, $3
) ON CONFLICT (machine_id) DO UPDATE SET
    agent_started_at = $2,
    agent_public_key = $3
;

-- name: MachineSetAgentHeartbeat :exec
INSERT INTO machine_agent_heartbeat (
    machine_id, agent_heartbeat_at
) VALUES (
    $1, $2
) ON CONFLICT (machine_id) DO UPDATE SET
    agent_heartbeat_at = $2
;

-- name: MachineSetHardwareReport :exec
INSERT INTO machine_hardware_report (
    machine_id, hardware_report_raw
) VALUES (
    $1, $2
) ON CONFLICT (machine_id) DO UPDATE SET
    hardware_report_raw = $2
;

-- name: MachineSetOSInstallationRequest :exec
INSERT INTO machine_os_installation_request (
    machine_id, generation, os_installation_request_raw
) VALUES (
    $1, $2, $3
) ON CONFLICT (machine_id) DO UPDATE SET
    generation = $2,
    os_installation_request_raw = $3
;

-- name: MachineSetOSInstallationReport :exec
INSERT INTO machine_os_installation_report (
    machine_id, generation
) VALUES (
    $1, $2
) ON CONFLICT (machine_id) DO UPDATE SET
    generation = $2
;


-- name: MachineAgentReset :exec