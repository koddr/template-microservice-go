-- Get all transactions from 'analytics' table.
SELECT COALESCE(
    json_agg(row_to_json(t)),
    '[]'::json
)
FROM (
    SELECT *
    FROM analytics
    ORDER BY created_at DESC
) t;