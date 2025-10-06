-- migrate:up
CREATE TABLE IF NOT EXISTS persons (
    id              BIGSERIAL PRIMARY KEY,
    name            TEXT NOT NULL,
    country         TEXT NOT NULL,
    created_at      TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at      TIMESTAMP WITH TIME ZONE NOT NULL
);

-- migrate:down
DROP TABLE IF EXISTS persons;
