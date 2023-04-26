-- name: GetProvidedMachines :many
SELECT
    machine_provided.*
FROM machines
INNER JOIN machine_provided ON machines.machine_id = machine_provided.machine_id
WHERE machine_provided.provider = $1;

-- name: GetMachinesForAgentStart :many
-- Get machines that need agent started for the first time. Machine can be
-- assumed to be 'new', with no previous attempts or failures.
-- ONCHANGE(queries_stats.sql): constraints must be kept in sync with StatsMachinesForAgentStart.
SELECT
    machine_provided.*
FROM machines
INNER JOIN machine_provided ON machines.machine_id = machine_provided.machine_id
LEFT JOIN work ON machines.machine_id = work.machine_id AND work.process IN ('ShepherdAccess', 'ShepherdAgentStart', 'ShepherdRecovery')
LEFT JOIN work_backoff ON machines.machine_id = work_backoff.machine_id AND work_backoff.until > now() AND work_backoff.process = 'ShepherdAgentStart'
LEFT JOIN machine_agent_started ON machines.machine_id = machine_agent_started.machine_id
LEFT JOIN machine_os_installation_request ON machines.machine_id = machine_os_installation_request.machine_id
LEFT JOIN machine_os_installation_report ON machines.machine_id = machine_os_installation_report.machine_id
WHERE
  machine_agent_started.machine_id IS NULL
  -- Do not start on machines that have a fulfilled OS installation request.
  AND (
      machine_os_installation_request.machine_id IS NULL
      OR machine_os_installation_request.generation IS DISTINCT FROM machine_os_installation_report.generation
  )
  AND work.machine_id IS NULL
  AND work_backoff.machine_id IS NULL
LIMIT $1;

-- name: GetMachineForAgentRecovery :many
-- Get machines that need agent restarted after something went wrong. Either
-- the agent started but never responded, or the agent stopped responding at
-- some point, or the machine got rebooted or somehow else lost the agent. Assume
-- some work needs to be performed on the shepherd side to diagnose and recover
-- whatever state the machine truly is in.
-- ONCHANGE(queries_stats.sql): constraints must be kept in sync with StatsMachinesForAgentRecovery.
SELECT
    machine_provided.*
FROM machines
INNER JOIN machine_provided ON machines.machine_id = machine_provided.machine_id
LEFT JOIN work ON machines.machine_id = work.machine_id AND work.process IN ('ShepherdAccess', 'ShepherdAgentStart', 'ShepherdRecovery')
LEFT JOIN work_backoff ON machines.machine_id = work_backoff.machine_id AND work_backoff.until > now() AND work_backoff.process = 'ShepherdRecovery'
INNER JOIN machine_agent_started ON machines.machine_id = machine_agent_started.machine_id
LEFT JOIN machine_agent_heartbeat ON machines.machine_id = machine_agent_heartbeat.machine_id
LEFT JOIN machine_os_installation_request ON machines.machine_id = machine_os_installation_request.machine_id
LEFT JOIN machine_os_installation_report ON machines.machine_id = machine_os_installation_report.machine_id
WHERE
  -- Do not recover machines that have a fulfilled OS installation request.
  (
      machine_os_installation_request.machine_id IS NULL
      OR machine_os_installation_request.generation IS DISTINCT FROM machine_os_installation_report.generation
  )
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
  AND work_backoff.machine_id IS NULL
LIMIT $1;

-- name: AuthenticateAgentConnection :many
-- Used by bmdb server to verify incoming connections.
SELECT
    machine_agent_started.*
FROM machines
INNER JOIN machine_agent_started ON machines.machine_id = machine_agent_started.machine_id
WHERE
    machines.machine_id = $1
    AND machine_agent_started.agent_public_key = $2;

-- name: GetExactMachineForOSInstallation :many
-- Get OS installation request for a given machine ID. Used by the bmdb server
-- to tell agent whether there's a pending installation request for the machine
-- it's running on.
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
