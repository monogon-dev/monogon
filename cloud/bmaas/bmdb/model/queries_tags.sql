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
    machine_id, generation, os_installation_result, os_installation_report_raw
) VALUES (
    $1, $2, $3, $4
) ON CONFLICT (machine_id) DO UPDATE SET
    generation = $2,
    os_installation_result = $3,
    os_installation_report_raw = $4
;


-- name: MachineDeleteAgentStarted :exec
DELETE FROM machine_agent_started
WHERE machine_id = $1;

-- name: MachineDeleteAgentHeartbeat :exec
DELETE FROM machine_agent_heartbeat
WHERE machine_id = $1;

-- name: MachineUpdateProviderStatus :exec
UPDATE machine_provided
SET
    provider_reservation_id = COALESCE($3, provider_reservation_id),
    provider_ip_address = COALESCE($4, provider_ip_address),
    provider_location = COALESCE($5, provider_location),
    provider_status = COALESCE($6, provider_status)
WHERE provider_id = $1
AND   provider = $2;