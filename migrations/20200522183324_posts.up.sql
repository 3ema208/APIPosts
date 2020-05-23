CREATE TABLE posts (
    id serial PRIMARY KEY,
    title varchar(512) not null,
    link varchar(256) not null,
    author_name VARCHAR(256) not null,
    created_time TIMESTAMP DEFAULT now(),
    amount_upvote INTEGER DEFAULT 0
);