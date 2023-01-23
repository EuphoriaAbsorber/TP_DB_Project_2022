CREATE UNLOGGED TABLE users (
    id  SERIAL PRIMARY KEY,
	email VARCHAR (40) NOT NULL,
    fullname VARCHAR (100) NOT NULL,
    nickname VARCHAR (40), 
    about VARCHAR (600)
);

CREATE UNLOGGED TABLE forums (
    id  SERIAL PRIMARY KEY,
	title VARCHAR (80) NOT NULL,
    user VARCHAR (40) NOT NULL,
    slug VARCHAR (100) NOT NULL,
    posts INT
    threads INT
);

CREATE UNLOGGED TABLE threads (
    id  SERIAL PRIMARY KEY,
	title VARCHAR (80) NOT NULL,
    author VARCHAR (40) NOT NULL,
    forum VARCHAR (80),
    message VARCHAR (400) NOT NULL,
    votes INT,
    slug VARCHAR (100),
    created TIMESTAMP
);

CREATE UNLOGGED TABLE posts (
    id  SERIAL PRIMARY KEY,
    parent  INT,
    author VARCHAR (40) NOT NULL,
    message VARCHAR (400) NOT NULL,
    forum VARCHAR (80),
    isedited BOOLEAN,
    thread INT REFERENCES threads (id) ON DELETE CASCADE,
    created TIMESTAMP
);



