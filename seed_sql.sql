CREATE TABLE if not exists articles
(
    article_id      INTEGER PRIMARY KEY AUTOINCREMENT,
    title           TEXT        NOT NULL,
    content         TEXT        NOT NULL,
    author_id       INTEGER     NOT NULL,
    published_at    TIMESTAMP            DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMP            DEFAULT CURRENT_TIMESTAMP,
    slug            TEXT UNIQUE NOT NULL,
    cover_image_url TEXT,
    views           INTEGER              DEFAULT 0,
    status          TEXT        NOT NULL DEFAULT 'draft',
    word_count      INTEGER,
    read_time       INTEGER,
    likes_count     INTEGER              DEFAULT 0,
    bookmarks_count INTEGER              DEFAULT 0,
    comments_count  INTEGER              DEFAULT 0
);

-- Create an index on the slug column for faster lookups.
CREATE UNIQUE INDEX if not exists idx_slug ON articles (slug);
CREATE INDEX if not exists idx_author_id ON articles (author_id);

CREATE TABLE if not exists users
(
    id                   SERIAL PRIMARY KEY,
    first_name           VARCHAR(255) NOT NULL,
    last_name            VARCHAR(255) NOT NULL,
    email                VARCHAR(255) NOT NULL UNIQUE,
    password             VARCHAR(255) NOT NULL,
    profile_picture_url  VARCHAR(1000),
    created_at           timestamp DEFAULT CURRENT_TIMESTAMP,
    is_verified          bool      default false,
    verification_code    INT       default 0,
    verification_expired timestamp default '1970-01-01 00:00:00'
);


-- Insert some test users
INSERT INTO users (first_name, last_name, email, password, profile_picture_url, is_verified)
VALUES
    ('John', 'Doe', 'john.doe@example.com', 'password123', 'http://example.com/john.jpg', TRUE),
    ('Jane', 'Doe', 'jane.doe@example.com', 'password456', 'http://example.com/jane.jpg', TRUE);


