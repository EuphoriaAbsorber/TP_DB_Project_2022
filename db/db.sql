CREATE UNLOGGED TABLE users (
	email    TEXT NOT NULL UNIQUE,
    fullname TEXT NOT NULL,
    nickname TEXT NOT NULL UNIQUE PRIMARY KEY, 
    about    TEXT
);

CREATE UNLOGGED TABLE forums (
	title   TEXT NOT NULL,
    user1   TEXT NOT NULL REFERENCES users (nickname) ON DELETE CASCADE,
    slug    TEXT PRIMARY KEY,
    posts   INT DEFAULT 0,
    threads INT DEFAULT 0
);

CREATE UNLOGGED TABLE threads (
    id      SERIAL PRIMARY KEY,
	title   TEXT NOT NULL,
    author  TEXT NOT NULL REFERENCES users (nickname) ON DELETE CASCADE,
    forum   TEXT NOT NULL REFERENCES forums (slug) ON DELETE CASCADE,
    message TEXT NOT NULL,
    votes   INT DEFAULT 0,
    slug    TEXT,
    created TIMESTAMP WITH TIME ZONE
);

CREATE UNLOGGED TABLE posts (
    id  SERIAL PRIMARY KEY,
    parent   INT DEFAULT 0,
    author   TEXT NOT NULL REFERENCES users (nickname) ON DELETE CASCADE,
    message  TEXT NOT NULL,
    forum    TEXT NOT NULL REFERENCES forums (slug) ON DELETE CASCADE,
    isedited BOOLEAN,
    thread   INT REFERENCES threads (id) ON DELETE CASCADE,
    created  TIMESTAMP WITH TIME ZONE,
    path     INT[]  DEFAULT ARRAY []::INTEGER[]
);

CREATE UNLOGGED TABLE votes (
    nickname  TEXT NOT NULL REFERENCES users (nickname) ON DELETE CASCADE,
    thread    INT    NOT NULL REFERENCES threads (id) ON DELETE CASCADE,
    voice     INT    NOT NULL,
    CONSTRAINT vote_key UNIQUE (nickname, thread)
);

--TRIGGERS

CREATE OR REPLACE FUNCTION post_set_path()
    RETURNS TRIGGER AS
$$
DECLARE
parent_post_id posts.id%type := 0;
BEGIN
    NEW.path = (SELECT path FROM posts WHERE id = NEW.parent) || NEW.id;
RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER insert_post
    BEFORE INSERT
    ON posts
    FOR EACH ROW
    EXECUTE PROCEDURE post_set_path();



--INDEXES

CREATE UNIQUE INDEX IF NOT EXISTS votes_key ON votes (thread, nickname);

CREATE INDEX IF NOT EXISTS threads_created ON threads (created);
CREATE INDEX IF NOT EXISTS threads_slug ON threads (lower(slug));

CREATE INDEX IF NOT EXISTS users_nickname ON users (lower(nickname));

CREATE INDEX IF NOT EXISTS posts_thread ON posts (thread);
