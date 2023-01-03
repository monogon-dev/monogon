-- name: NewMachine :one
INSERT INTO machines (
    machine_created_at
) VALUES (
    now()
)
RETURNING *;

-- name: NewSession :one
INSERT INTO sessions (
    session_component_name, session_runtime_info, session_created_at, session_interval_seconds, session_deadline
) VALUES (
    $1, $2, now(), $3, (now() + $3 * interval '1 second')
)
RETURNING *;

-- name: SessionPoke :exec
-- Update a given session with a new deadline. Must be called in the same
-- transaction as SessionCheck to ensure the session is still alive.
UPDATE sessions
SET session_deadline = now() + session_interval_seconds * interval '1 second'
WHERE session_id = $1;

-- name: SessionCheck :many
-- SessionCheck returns a session by ID if that session is still valid (ie. its
-- deadline hasn't expired).
SELECT *
FROM sessions
WHERE session_id = $1
AND session_deadline > now();

-- name: StartWork :exec
INSERT INTO work (
    machine_id, session_id, process
) VALUES (
    $1, $2, $3
);

-- name: FinishWork :exec
DELETE FROM work
WHERE machine_id = $1
  AND session_id = $2
  AND process = $3;

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

-- name: GetMachinesForAgentStart :many
-- Get machines that need agent installed for the first time. Machine can be
-- assumed to be 'new', with no previous attempts or failures.
SELECT
    machine_provided.*
FROM machines
INNER JOIN machine_provided ON machines.machine_id = machine_provided.machine_id
LEFT JOIN work ON machines.machine_id = work.machine_id AND work.process = 'ShepherdInstall'
LEFT JOIN machine_agent_started ON machines.machine_id = machine_agent_started.machine_id
WHERE
  machine_agent_started.machine_id IS NULL
  -- TODO(q3k): exclude machines which are not expected to run the agent (eg.
  -- are already exposed to a user).
  AND work.machine_id IS NULL
LIMIT $1;

-- name: GetMachineForAgentRecovery :many
-- Get machines that need agent installed after something went wrong. Either
-- the agent started but never responded, or the agent stopped responding at
-- some point, or the machine is being reinstalled after failure. Assume some
-- work needs to be performed on the shepherd side to diagnose and recover
-- whatever state the machine truly is in.
SELECT
    machine_provided.*
FROM machines
INNER JOIN machine_provided ON machines.machine_id = machine_provided.machine_id
LEFT JOIN work ON machines.machine_id = work.machine_id AND work.process = 'ShepherdInstall'
LEFT JOIN machine_agent_started ON machines.machine_id = machine_agent_started.machine_id
LEFT JOIN machine_agent_heartbeat ON machines.machine_id = machine_agent_heartbeat.machine_id
WHERE
  -- Only act on machines where the agent is expected to have been started.
  machine_agent_started.machine_id IS NOT NULL
  AND (
    -- No heartbeat 30 minutes after starting the agent.
    (
        machine_agent_heartbeat.machine_id IS NULL
        AND now() > (machine_agent_started.agent_started_at + interval '30 minutes')
    )
    -- Heartbeats ceased for 10 minutes.
    OR (
        machine_agent_heartbeat.machine_id IS NOT NULL
        AND now() > (machine_agent_heartbeat.agent_heartbeat_at + interval '10 minutes')
    )
  )
  AND work.machine_id IS NULL
LIMIT $1;

-- name: GetExactMachineForOSInstallation :many
SELECT
    machine_os_installation_request.*
FROM machines
LEFT JOIN machine_os_installation_request ON machines.machine_id = machine_os_installation_request.machine_id
LEFT JOIN machine_os_installation_report ON machines.machine_id = machine_os_installation_report.machine_id
WHERE
    -- We are only interested in one concrete machine.
    machines.machine_id = $1
    -- We must have an installation request.
    AND machine_os_installation_request.machine_id IS NOT NULL
    -- And we either must have no installation report, or the installation
    -- report's generation must not match the installation request's generation.
    AND (
        machine_os_installation_report.machine_id IS NULL
        OR machine_os_installation_report.generation != machine_os_installation_request.generation
    )
LIMIT $2;

-- name: AuthenticateAgentConnection :many
SELECT
    machine_agent_started.*
FROM machines
INNER JOIN machine_agent_started ON machines.machine_id = machine_agent_started.machine_id
WHERE
    machines.machine_id = $1
    AND machine_agent_started.agent_public_key = $2;
