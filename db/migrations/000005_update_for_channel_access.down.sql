BEGIN;

-- Revert the changes made in the up migration
ALTER TABLE customer DROP COLUMN IF EXISTS channel_access_granted;
ALTER TABLE customer DROP COLUMN IF EXISTS is_in_channel;
ALTER TABLE customer RENAME COLUMN channel_invite_link TO subscription_link;

-- Revert purchase table changes
ALTER TABLE purchase DROP COLUMN IF EXISTS channel_access_granted;

COMMIT;