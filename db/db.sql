CREATE UNLOGGED TABLE users (
	email VARCHAR (40) NOT NULL UNIQUE,
    fullname VARCHAR (100) NOT NULL,
    nickname VARCHAR (40) NOT NULL UNIQUE PRIMARY KEY, 
    about VARCHAR (600)
);

CREATE UNLOGGED TABLE forums (
	title VARCHAR (80) NOT NULL,
    user1 VARCHAR (40) NOT NULL REFERENCES users (nickname) ON DELETE CASCADE,
    slug VARCHAR (100) PRIMARY KEY,
    posts INT DEFAULT 0,
    threads INT DEFAULT 0
);

CREATE UNLOGGED TABLE threads (
    id  SERIAL PRIMARY KEY,
	title VARCHAR (80) NOT NULL,
    author VARCHAR (40) NOT NULL REFERENCES users (nickname) ON DELETE CASCADE,
    forum VARCHAR (80) NOT NULL REFERENCES forums (slug) ON DELETE CASCADE,
    message VARCHAR (400) NOT NULL,
    votes INT DEFAULT 0,
    slug VARCHAR (100) UNIQUE,
    created TIMESTAMP
);

CREATE UNLOGGED TABLE posts (
    id  SERIAL PRIMARY KEY,
    parent  INT,
    author VARCHAR (40) NOT NULL REFERENCES users (nickname) ON DELETE CASCADE,
    message VARCHAR (400) NOT NULL,
    forum VARCHAR (100) NOT NULL REFERENCES forums (slug) ON DELETE CASCADE,
    isedited BOOLEAN,
    thread INT REFERENCES threads (id) ON DELETE CASCADE,
    created TIMESTAMP
);

CREATE UNLOGGED TABLE votes (
    nickname  VARCHAR (40) NOT NULL REFERENCES users (nickname) ON DELETE CASCADE,
    thread    INT    NOT NULL REFERENCES threads (id) ON DELETE CASCADE,
    voice     INT    NOT NULL,
    CONSTRAINT vote_key UNIQUE (nickname, thread)
);



