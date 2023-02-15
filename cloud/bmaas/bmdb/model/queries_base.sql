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


-- name: WorkHistoryInsert :exec
-- Insert an entry into the work_history audit table.
INSERT INTO work_history (
    machine_id, process, event, timestamp, failed_cause
) VALUES (
    $1, $2, $3, now(), $4
);

-- name: WorkBackoffInsert :exec
-- Upsert a backoff for a given machine/process.
INSERT INTO work_backoff (
    machine_id, process, cause, until
) VALUES (
    $1, $2, $3, now() + (sqlc.arg(seconds)::int * interval '1 second')
) ON CONFLICT (machine_id, process) DO UPDATE SET
    cause = $3,
    until = now() + (sqlc.arg(seconds)::int * interval '1 second')
;

-- name: ListHistoryOf :many
-- Retrieve full audit history of a machine.
SELECT *
FROM work_history
WHERE machine_id = $1
ORDER BY timestamp ASC;

-- name: GetSession :many
-- Retrieve session information by session ID.
SELECT *
FROM sessions
WHERE session_id = $1;