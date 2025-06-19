-- Create 'analytics' table.
CREATE TABLE IF NOT EXISTS analytics (
   id BIGSERIAL PRIMARY KEY,
   profile_id VARCHAR(64) NOT NULL,
   messenger_id VARCHAR(64) NOT NULL,
   messenger_name VARCHAR(16) NOT NULL,
   event_id SMALLINT NOT NULL,
   event_type VARCHAR(16) NOT NULL,
   utm_source VARCHAR(32) NULL,
   utm_medium VARCHAR(32) NULL,
   utm_campaign TEXT NULL,
   utm_content VARCHAR(64) NULL,
   utm_term VARCHAR(32) NULL,
   created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Create indexes.
CREATE INDEX IF NOT EXISTS idx_analytics_profile_id ON analytics (profile_id);
CREATE INDEX IF NOT EXISTS idx_analytics_messenger_id ON analytics (messenger_id);
CREATE INDEX IF NOT EXISTS idx_analytics_created_at ON analytics (created_at DESC);