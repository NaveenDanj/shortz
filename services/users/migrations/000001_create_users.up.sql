CREATE TABLE users (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    username        TEXT    NOT NULL,
    display_name    TEXT    NOT NULL,
    avatar_key      TEXT    NOT NULL DEFAULT '',
    bio             TEXT    NOT NULL DEFAULT '',
    links           JSONB   NOT NULL DEFAULT '[]',
    is_private      BOOLEAN NOT NULL DEFAULT FALSE,
    follower_count  BIGINT  NOT NULL DEFAULT 0,
    following_count BIGINT  NOT NULL DEFAULT 0,
    video_count     BIGINT  NOT NULL DEFAULT 0,
    created_at      BIGINT  NOT NULL,
    updated_at      BIGINT  NOT NULL,
    deleted_at      BIGINT,
    CONSTRAINT users_username_format CHECK (username ~ '^[A-Za-z0-9_.]{3,30}$')
);

CREATE UNIQUE INDEX users_username_lower_uidx ON users (lower(username)) WHERE deleted_at IS NULL;

CREATE INDEX users_username_prefix_idx ON users (lower(username) text_pattern_ops) WHERE deleted_at IS NULL;
