CREATE TABLE IF NOT EXISTS images (
        id uuid PRIMARY KEY ,
        name varchar(255) UNIQUE NOT NULL
);