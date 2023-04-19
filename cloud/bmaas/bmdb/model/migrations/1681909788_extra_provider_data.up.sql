CREATE TYPE provider_status AS ENUM (
    -- The provider has no idea about this machine. This likely means something
    -- went really wrong and should be investigated.
    'Missing',
    -- The provider is attempting to install the initial operating system and
    -- give us control over the machine.
    'Provisioning',
    -- The provider has failed to provide this machine and is not expected to
    -- be able to provision it without deprovisioning it first.
    'ProvisioningFailedPermanent',

    -- The provider sees this machine as running/healthy and ready for use by us.
    'Running',

    -- The provider sees that machine as administratively stopped/down, but not
    -- failed. It can be enabled / turned back on and should become Running.
    'Stopped',

    -- Any other state that we're not able to parse. Likely should be
    -- investigated.
    'Unknown'
);

ALTER TABLE machine_provided
-- Optional hardware reservation ID for this provider. Currently only implemented
-- for Equinix.
ADD COLUMN provider_reservation_id STRING(128) NULL,
-- Optional 'main' IP address as seen by provider. 'Main' is provider specific,
-- but generally should be the IP address that system operators would consider
-- the primary IP address of the machine, generally the one that operators would
-- SSH into. It might be a publicly routable address or might not be. It might
-- be a single IP address or a CIDR. Regardless, it's some human-readable
-- representation of the address, and generally should not be machine-parsed.
--
-- On Equinix, we pick the first IP address marked as 'public'.
ADD COLUMN provider_ip_address STRING(128) NULL,
-- Optional location/region as seen by provider. This is provider-specific: it
-- might be a city name, some internal metro ID, a PoP name, a slug, or even an
-- opaque string.
ADD COLUMN provider_location STRING(128) NULL,
-- Optional status as seen by provider. This is converted from provider-specific
-- data into an internal type. This field is machine-usable and with time should
-- be moved to be non-null.
ADD COLUMN provider_status provider_status NULL;