CREATE TABLE IF NOT EXISTS Users 
(
    id SERIAL PRIMARY KEY,
    login VARCHAR(20) NOT NULL,
    passhash VARCHAR(256) NOT NULL
);