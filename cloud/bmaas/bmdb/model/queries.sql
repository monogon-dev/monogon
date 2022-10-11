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

-- Example tag processing queries follow.

-- name: MachineAddProvided :exec
INSERT INTO machine_provided (
    machine_id, provider, provider_id
) VALUES (
    $1, $2, $3
);

-- name: GetMachinesNeedingInstall :many
SELECT
    machine_provided.*
FROM machines
INNER JOIN machine_provided ON machines.machine_id = machine_provided.machine_id
LEFT JOIN work ON machines.machine_id = work.machine_id AND work.process = 'NecromancerInstall'
LEFT JOIN machine_agent_installed ON machines.machine_id = machine_agent_installed.machine_id
WHERE machine_agent_installed.machine_id IS NULL
  AND work.machine_id IS NULL
LIMIT $1;

