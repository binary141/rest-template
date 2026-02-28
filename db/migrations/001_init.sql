CREATE TABLE IF NOT EXISTS users (
    id           SERIAL PRIMARY KEY,
    email        VARCHAR(255) NOT NULL UNIQUE,
    display_name VARCHAR(255),
    password     TEXT NOT NULL,
    can_login    BOOLEAN NOT NULL DEFAULT true,
    is_admin     BOOLEAN NOT NULL DEFAULT false,
    created_at   TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at   TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at   TIMESTAMPTZ
);
