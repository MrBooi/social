CREATE EXTENSION  IF NOT EXISTS CITEXT;

CREATE TABLE IF NOT EXISTS users(
    id bigserial Primary Key,
    email CITEXT UNIQUE NOT NULL,
    username varchar(255) UNIQUE NOT NULL,
    password bytea NOT NULL,
    created_at timestamp(0) with time zone Not Null DEFAULT Now()
);