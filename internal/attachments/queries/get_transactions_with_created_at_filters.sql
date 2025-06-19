-- Get transactions from 'analytics' table with 'created_at' filters.
SELECT COALESCE(
    json_agg(row_to_json(t)),
    '[]'::json
)
FROM (
    SELECT *
    FROM analytics
    WHERE created_at BETWEEN $1::timestamp AND $2::timestamp
    ORDER BY created_at DESC
) t;