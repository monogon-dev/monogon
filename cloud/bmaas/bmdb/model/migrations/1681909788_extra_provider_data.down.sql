ALTER TABLE machine_provided
DROP COLUMN provider_reservation_id,
DROP COLUMN provider_ip_address,
DROP COLUMN provider_location,
DROP COLUMN provider_status;
DROP type provider_status;
