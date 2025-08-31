BEGIN;

-- Rename subscription_link to channel_invite_link for clarity
ALTER TABLE customer RENAME COLUMN subscription_link TO channel_invite_link;

-- Add column to track if user is currently in the channel (for cleanup purposes)
ALTER TABLE customer ADD COLUMN is_in_channel BOOLEAN DEFAULT FALSE;

-- Update purchase table to reflect channel access terminology
ALTER TABLE purchase ADD COLUMN channel_access_granted BOOLEAN DEFAULT FALSE;

COMMIT;