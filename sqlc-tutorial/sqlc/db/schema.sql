CREATE TABLE authors (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    password varchar(64),
    name text NOT NULL,
    age int4range,
    bio text
);