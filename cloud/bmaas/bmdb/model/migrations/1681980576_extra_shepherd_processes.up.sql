-- Add two more process kinds, ShepherdAgentStart and ShepherdRecovery, for agent
-- start and recovery by the shepherd respectively. These deprecate the previous
-- ShepherdAccess process. The two processes mutually exclude each other.

ALTER TYPE process ADD VALUE IF NOT EXISTS 'ShepherdAgentStart';
ALTER TYPE process ADD VALUE IF NOT EXISTS 'ShepherdRecovery';