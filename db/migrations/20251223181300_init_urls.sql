-- +goose Up
CREATE TABLE IF NOT EXISTS urls (
    id BIGSERIAL PRIMARY KEY,
    code TEXT NOT NULL UNIQUE,
    original_url TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    clicks BIGINT NOT NULL DEFAULT 0
    );

CREATE INDEX IF NOT EXISTS idx_urls_code ON urls(code);

-- +goose Down
DROP TABLE IF EXISTS urls;
