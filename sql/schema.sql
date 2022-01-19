CREATE DATABASE demo

CREATE TABLE record (
    id              integer PRIMARY KEY NOT NULL,
    title           text NOT NULL,
    content         text NOT NULL,
    create_at       timestamp not null
);
