CREATE TABLE machines (
    machine_id UUID NOT NULL DEFAULT gen_random_uuid() PRIMARY KEY,
    machine_created_at TIMESTAMPTZ NOT NULL
);


-- Sessions are maintained by components as they work on the rest of the machine
-- database. Once a session is created, it must be maintained by its owning
-- component by repeatedly 'poking' it, ie. updating the heartbeat_deadline
-- value to be some point in the future.
--
-- TODO: garbage collect old sessions.
CREATE TABLE sessions (
    session_id UUID NOT NULL DEFAULT gen_random_uuid() PRIMARY KEY,
    -- Name of component which created this session. Human-readable.
    session_component_name STRING NOT NULL,
    -- Node name, hostname:port, pod name, whatever. Something to tell where
    -- a particular component is running. Human-readable.
    session_runtime_info STRING NOT NULL,
    -- Time at which this session was created.
    session_created_at TIMESTAMPTZ NOT NULL,
    -- Number of seconds by which session_deadline (counting from now())
    -- is bumped up every time the session is poked.
    session_interval_seconds INT NOT NULL,
    -- Deadline after which this session should not be considered valid anymore.
    session_deadline TIMESTAMPTZ NOT NULL
);

CREATE TYPE process AS ENUM (
    -- Reserved for unit tests.
    'UnitTest1',
    'UnitTest2'
);

-- Work items map a session to work performed on a machine. Multiple work items
-- can exist per session, and thus, a session can back multiple items of work
-- acting on multiple machines. These are optionally created by components to
-- indicate some long-running process being performed on a machine, and will
-- lock out the same process from being run simultaneously, eg. in a
-- concurrently running instance of the same component.
CREATE TABLE work (
    -- Machine that this work is being performed on. Prevent deleting machines
    -- that have active work on them.
    machine_id UUID NOT NULL REFERENCES machines(machine_id) ON DELETE RESTRICT,
    -- Session that this work item is tied to. If the session expires, so does
    -- the work item.
    session_id UUID NOT NULL REFERENCES sessions(session_id) ON DELETE CASCADE,
    -- Human-readable process name.
    process process NOT NULL,
    UNIQUE (machine_id, process),
    CONSTRAINT "primary" PRIMARY KEY (machine_id, session_id, process)
);