CREATE UNLOGGED TABLE users (
	email TEXT NOT NULL UNIQUE,
    fullname TEXT NOT NULL,
    nickname TEXT NOT NULL UNIQUE PRIMARY KEY, 
    about TEXT
);

CREATE UNLOGGED TABLE forums (
	title TEXT NOT NULL,
    user1 TEXT NOT NULL REFERENCES users (nickname) ON DELETE CASCADE,
    slug TEXT PRIMARY KEY,
    posts INT DEFAULT 0,
    threads INT DEFAULT 0
);

CREATE UNLOGGED TABLE threads (
    id  SERIAL PRIMARY KEY,
	title TEXT NOT NULL,
    author TEXT NOT NULL REFERENCES users (nickname) ON DELETE CASCADE,
    forum TEXT NOT NULL REFERENCES forums (slug) ON DELETE CASCADE,
    message TEXT NOT NULL,
    votes INT DEFAULT 0,
    slug TEXT,
    created TIMESTAMP
);

CREATE UNLOGGED TABLE posts (
    id  SERIAL PRIMARY KEY,
    parent  INT,
    author TEXT NOT NULL REFERENCES users (nickname) ON DELETE CASCADE,
    message TEXT NOT NULL,
    forum TEXT NOT NULL REFERENCES forums (slug) ON DELETE CASCADE,
    isedited BOOLEAN,
    thread INT REFERENCES threads (id) ON DELETE CASCADE,
    created TIMESTAMP
);

CREATE UNLOGGED TABLE votes (
    nickname  TEXT NOT NULL REFERENCES users (nickname) ON DELETE CASCADE,
    thread    INT    NOT NULL REFERENCES threads (id) ON DELETE CASCADE,
    voice     INT    NOT NULL,
    CONSTRAINT vote_key UNIQUE (nickname, thread)
);



