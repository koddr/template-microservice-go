-- Insert new transaction to 'analytics' table.
INSERT INTO analytics (
    profile_id,
    messenger_id,
    messenger_name,
    event_id,
    event_type,
    utm_source,
    utm_medium,
    utm_campaign,
    utm_content,
    utm_term
) VALUES (
    $1::varchar,
    $2::varchar,
    $3::varchar,
    $4::smallint,
    $5::varchar,
    $6::varchar,
    $7::varchar,
    $8::varchar,
    $9::varchar,
    $10::varchar
);