-- name: CountActiveWork :many
-- Return number of active work, grouped by process.
SELECT COUNT(*), work.process
FROM work
GROUP BY (work.process);

-- name: CountActiveBackoffs :many
-- Return number of active backoffs, grouped by process.
SELECT COUNT(*), work_backoff.process
FROM work_backoff
GROUP BY (work_backoff.process);

-- name: CountMachines :one
SELECT COUNT(*)
FROM machines;

-- name: CountMachinesProvided :one
SELECT COUNT(*)
FROM machine_provided;

-- name: CountMachinesAgentHeartbeating :one
SELECT COUNT(*)
FROM machines
    INNER JOIN machine_provided ON machines.machine_id = machine_provided.machine_id
    INNER JOIN machine_agent_heartbeat ON machines.machine_id = machine_agent_heartbeat.machine_id
WHERE
    now() < machine_agent_heartbeat.agent_heartbeat_at + interval '10 minute';

-- name: CountMachinesInstallationPending :one
SELECT COUNT(*)
FROM machines
    INNER JOIN machine_provided ON machines.machine_id = machine_provided.machine_id
    INNER JOIN machine_os_installation_request ON machines.machine_id = machine_os_installation_request.machine_id
    LEFT JOIN machine_os_installation_report ON machines.machine_id = machine_os_installation_report.machine_id
WHERE
        machine_os_installation_request.generation IS DISTINCT FROM machine_os_installation_report.generation;

-- name: CountMachinesInstallationComplete :one
SELECT COUNT(*)
FROM machines
         INNER JOIN machine_provided ON machines.machine_id = machine_provided.machine_id
         INNER JOIN machine_os_installation_request ON machines.machine_id = machine_os_installation_request.machine_id
         INNER JOIN machine_os_installation_report ON machines.machine_id = machine_os_installation_report.machine_id
WHERE
    machine_os_installation_request.generation IS NOT DISTINCT FROM machine_os_installation_report.generation;

-- name: CountMachinesForAgentStart :one
-- Return number of machines eligible for agent start.
-- ONCHANGE(queries_workflows.sql): constraints must be kept in sync with GetMachinesForAgentStart.
SELECT COUNT(machine_provided)
FROM machines
         INNER JOIN machine_provided ON machines.machine_id = machine_provided.machine_id
         LEFT JOIN work ON machines.machine_id = work.machine_id AND work.process IN ('ShepherdAccess', 'ShepherdAgentStart', 'ShepherdRecovery')
         LEFT JOIN work_backoff ON machines.machine_id = work_backoff.machine_id AND work_backoff.until > now() AND work_backoff.process = 'ShepherdAgentStart'
         LEFT JOIN machine_agent_started ON machines.machine_id = machine_agent_started.machine_id
         LEFT JOIN machine_os_installation_request ON machines.machine_id = machine_os_installation_request.machine_id
         LEFT JOIN machine_os_installation_report ON machines.machine_id = machine_os_installation_report.machine_id
WHERE
    work.machine_id IS NULL
    AND work_backoff.machine_id IS NULL
    AND machine_agent_started.machine_id IS NULL
    -- If there is a pending installation request, it must not have been fulfilled already.
    AND (
        machine_os_installation_request.machine_id IS NULL
        OR machine_os_installation_report.generation IS DISTINCT FROM machine_os_installation_request.generation
    );

-- name: CountMachinesForAgentRecovery :one
-- Return number of machines eligible for agent recovery.
-- ONCHANGE(queries_workflows.sql): constraints must be kept in sync with GetMachineForAgentRecovery.
SELECT COUNT(machine_provided)
FROM machines
         INNER JOIN machine_provided ON machines.machine_id = machine_provided.machine_id
         LEFT JOIN work ON machines.machine_id = work.machine_id AND work.process IN ('ShepherdAccess', 'ShepherdAgentStart', 'ShepherdRecovery')
         LEFT JOIN work_backoff ON machines.machine_id = work_backoff.machine_id AND work_backoff.until > now() AND work_backoff.process = 'ShepherdRecovery'
         INNER JOIN machine_agent_started ON machines.machine_id = machine_agent_started.machine_id
         LEFT JOIN machine_agent_heartbeat ON machines.machine_id = machine_agent_heartbeat.machine_id
         LEFT JOIN machine_os_installation_request ON machines.machine_id = machine_os_installation_request.machine_id
         LEFT JOIN machine_os_installation_report ON machines.machine_id = machine_os_installation_report.machine_id
WHERE
    work.machine_id IS NULL
    AND work_backoff.machine_id IS NULL
    -- Only act on machines where the agent is expected to have been started:
    -- 1. If there is a pending installation request, it must not have been fulfilled already.
    AND (
        machine_os_installation_request.machine_id IS NULL
        OR machine_os_installation_report.generation IS DISTINCT FROM machine_os_installation_request.generation
    )
    -- 2. The agent must have never heartbeat or must have stopped heartbeating.
    AND (
        -- No heartbeat 30 minutes after starting the agent.
        ( machine_agent_heartbeat.machine_id IS NULL
          AND now() > (machine_agent_started.agent_started_at + interval '30 minutes')
        )
        -- Heartbeats ceased for 10 minutes.
        OR ( machine_agent_heartbeat.machine_id IS NOT NULL
          AND now() > (machine_agent_heartbeat.agent_heartbeat_at + interval '10 minutes')
        )
    );

-- name: ListMachineHardware :many
SELECT
    machine_provided.*,
    machine_hardware_report.*
FROM machines
         INNER JOIN machine_provided ON machines.machine_id = machine_provided.machine_id
         INNER JOIN machine_hardware_report ON machines.machine_id = machine_hardware_report.machine_id
WHERE machines.machine_id > $1
ORDER BY machines.machine_id ASC
LIMIT $2;
