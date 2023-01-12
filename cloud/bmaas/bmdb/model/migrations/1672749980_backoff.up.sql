CREATE TYPE work_history_event AS ENUM (
    'Started',
    'Finished',
    'Failed',
    'Canceled'
);

-- Audit trail of work history for a given machine.
CREATE TABLE work_history(
    -- The machine subject to this audit entry. As we want to allow keeping
    -- information about deleted machines, this is not a foreign key.
    machine_id UUID NOT NULL,

    -- TODO(q3k): session history?

    -- Process acting on this machine which caused an audit entry to be created.
    process process NOT NULL,
    -- Process lifecycle event (started, finished, etc) that caused this audit
    -- entry to be created.
    event work_history_event NOT NULL,
    -- Time at which this entry was created.
    timestamp TIMESTAMPTZ NOT NULL,

    -- Failure cause, only set when event == Failed.
    failed_cause STRING
);

CREATE INDEX ON work_history (machine_id);

-- Backoff entries are created by failed work items, and effectively act as
-- a Lockout-tagout entry for a given machine and a given process.
--
-- Currently, there is no way to fully backoff an entire machine, just
-- individual processes from a given machine.
--
-- Backoff entries are only valid as long as 'until' is before now(), after that
-- they are ignored by workflow queries. Insertion queries act as upserts,
-- and thus the backoff entries do not need to be garbage collected, as they do
-- not grow unbounded (maximumum one entry per process/machine).
CREATE TABLE work_backoff(
    -- The machine affected by this backoff.
    machine_id UUID NOT NULL REFERENCES machines(machine_id) ON DELETE CASCADE,
    -- The process that this machine should not be subjected to.
    process process NOT NULL,
    -- Until when the backoff is enforced.
    until TIMESTAMPTZ NOT NULL,

    -- Error reported by process/work when this backoff was inserted.
    -- Human-readable.
    cause STRING NOT NULL,

    UNIQUE(machine_id, process)
);