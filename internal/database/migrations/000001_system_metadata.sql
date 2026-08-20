-- +goose Up
CREATE TABLE system_metadata (
    key TEXT PRIMARY KEY,
    value JSONB NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

INSERT INTO system_metadata (key, value)
VALUES ('schema', '{"phase": 0, "owner": "aetherd"}'::jsonb)
ON CONFLICT (key) DO NOTHING;

-- +goose Down
DROP TABLE system_metadata;

