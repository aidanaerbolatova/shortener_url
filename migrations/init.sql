CREATE TABLE IF NOT EXISTS  link_info (
        id SERIAL PRIMARY KEY,
        full_link TEXT NOT NULL,
        shortener_link TEXT NOT NULL,
        visits INT DEFAULT 0,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        CONSTRAINT unique_full_link UNIQUE (full_link),
        CONSTRAINT unique_shortener_link UNIQUE (shortener_link)
);
