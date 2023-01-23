CREATE TABLE users (
    id  SERIAL PRIMARY KEY,
	email VARCHAR (40) NOT NULL,
    fullname VARCHAR (100) NOT NULL,
    nickname VARCHAR (40), 
    about VARCHAR (600)
);