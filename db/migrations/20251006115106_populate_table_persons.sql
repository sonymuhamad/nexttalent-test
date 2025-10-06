-- migrate:up
INSERT INTO persons (name, country, created_at, updated_at)
VALUES
    ('Adam', 'Kuala Lumpur', NOW(), NOW()),
    ('John', 'Singapore', NOW(), NOW()),
    ('Henry', 'Singapore', NOW(), NOW()),
    ('Dominic', 'Thailand', NOW(), NOW());


-- migrate:down
DELETE FROM persons WHERE name IN ('Adam', 'John', 'Henry', 'Dominic');
