CREATE TABLE IF NOT EXISTS Profiles
(
    id SERIAL PRIMARY KEY,
    name VARCHAR(32) NOT NULL,
    user_age INT,
    about TEXT,
    img_path VARCHAR(256)
);