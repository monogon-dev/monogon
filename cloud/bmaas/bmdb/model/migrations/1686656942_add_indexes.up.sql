-- Used by the agent gRPC server to retrieve agent information by public key.
CREATE INDEX agent_public_key_idx
ON machine_agent_started (agent_public_key)
INCLUDE (agent_started_at);

-- Used by queries which require a live session.
CREATE INDEX session_id_deadline_idx
ON sessions (session_id, session_deadline)
INCLUDE (session_component_name, session_runtime_info, session_Created_at, session_interval_seconds);

-- Used by work retrieval/scheduling queries to exclude machines that have a given process backed off.
CREATE INDEX process_machine_id_idx
ON work_backoff (process, machine_id)
INCLUDE (until);