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
CREATE  INDEX if not exists idx_author_id ON articles (author_id);

CREATE TABLE if not exists users
(
    user_id             INTEGER PRIMARY KEY AUTOINCREMENT,
    username            TEXT NOT NULL,
    email               TEXT NOT NULL UNIQUE,
    password_hash       TEXT NOT NULL,
    profile_picture_url TEXT,
    created_at          TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- -- Inserting sample user data
-- INSERT INTO users (user_id, username, email, password_hash, profile_picture_url)
-- VALUES
--     (1, 'user1', 'user1@example.com', 'hashed_password_1', 'https://example.com/user1_profile.jpg'),
--     (2, 'user2', 'user2@example.com', 'hashed_password_2', 'https://example.com/user2_profile.jpg'),
--     (3, 'user3', 'user3@example.com', 'hashed_password_3', 'https://example.com/user3_profile.jpg');
--
--
-- -- Inserting sample article data
-- INSERT INTO articles (title, content, author_id, slug, cover_image_url, status, word_count, read_time)
-- VALUES
--     ('Sample Article 1', 'This is the content of the first article.', 1, 'sample-article-1', 'https://example.com/article1.jpg', 'published', 500, 5),
--     ('Sample Article 2', 'This is the content of the second article.', 2, 'sample-article-2', 'https://example.com/article2.jpg', 'published', 700, 7),
--     ('Sample Article 3', 'This is the content of the third article.', 1, 'sample-article-3', 'https://example.com/article3.jpg', 'draft', 600, 6);
