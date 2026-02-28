CREATE TABLE IF NOT EXISTS sessions (
    id         SERIAL PRIMARY KEY,
    session_id TEXT NOT NULL UNIQUE,
    expires_at BIGINT NOT NULL,
    is_valid   BOOLEAN NOT NULL DEFAULT true,
    user_id    INTEGER NOT NULL REFERENCES users(id)
);
