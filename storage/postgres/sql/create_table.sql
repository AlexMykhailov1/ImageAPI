CREATE TABLE IF NOT EXISTS images (
        id serial PRIMARY KEY ,
        name varchar(255) UNIQUE NOT NULL
);