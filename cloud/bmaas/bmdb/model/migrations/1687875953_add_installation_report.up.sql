CREATE TYPE machine_os_installation_result AS ENUM (
    'Success',
    'Error'
    );

-- Add column for storing the serialized cloud.bmaas.server.api.OSInstallationReport
-- also add a column to display if the installation was successful or not.
ALTER TABLE machine_os_installation_report
    ADD COLUMN os_installation_result machine_os_installation_result NOT NULL,
    ADD COLUMN os_installation_report_raw BYTEA NOT NULL;