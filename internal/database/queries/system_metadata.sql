-- name: GetSystemMetadata :one
SELECT key, value, updated_at
FROM system_metadata
WHERE key = $1;

-- name: UpsertSystemMetadata :one
INSERT INTO system_metadata (key, value)
VALUES ($1, $2)
ON CONFLICT (key)
DO UPDATE SET
    value = EXCLUDED.value,
    updated_at = NOW()
RETURNING key, value, updated_at;

