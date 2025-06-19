-- Drop indexes.
DROP INDEX IF EXISTS idx_analytics_profile_id;
DROP INDEX IF EXISTS idx_analytics_messenger_id;
DROP INDEX IF EXISTS idx_analytics_created_at;

-- Drop 'analytics' table.
DROP TABLE IF EXISTS analytics;