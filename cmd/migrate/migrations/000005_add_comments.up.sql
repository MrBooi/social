CREATE TABLE IF Not Exists comments
(
    id      bigserial NOT NULL,
    post_id bigserial NOT NULL,
    user_id bigserial NOT NULL,
    content TEXT      NOT NULL,
    created_at timestamp(0) with time zone Not Null DEFAULT Now(),
    updated_at timestamp(0) with time zone Not Null DEFAULT Now()
);
