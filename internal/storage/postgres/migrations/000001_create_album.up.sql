CREATE TABLE album
(
    id     SERIAL PRIMARY KEY,
    title  VARCHAR(128)  NOT NULL,
    artist VARCHAR(255)  NOT NULL,
    price  NUMERIC(5, 2) NOT NULL
);