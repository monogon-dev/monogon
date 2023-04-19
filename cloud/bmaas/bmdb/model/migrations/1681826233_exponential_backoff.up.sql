-- Add interval, in seconds. This is used to calculate subsequent backoff values
-- for exponential backoffs. A future migration should make this field
-- non-nullable.
ALTER TABLE work_backoff
ADD COLUMN last_interval_seconds BIGINT NULL;