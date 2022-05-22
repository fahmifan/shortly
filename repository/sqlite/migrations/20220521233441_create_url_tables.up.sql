CREATE TABLE IF NOT EXISTS urls (
    id TEXT PRIMARY KEY,
    is_public BOOLEAN NOT NULL DEFAULT 0,
    original TEXT NOT NULL,
    shorten TEXT NOT NULL,
    created_by TEXT NOT NULL,
    updated_by TEXT NOT NULL,
    deleted_by TEXT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP NULL
);