CREATE TABLE IF NOT EXISTS posts(
    id bigserial Primary Key,
    title text NOT NULL,
    user_id bigint NOT NULL,
    content text NOT NULL,
    created_at timestamp(0) with time zone Not Null DEFAULT Now()
    );