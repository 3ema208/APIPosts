CREATE TABLE comment (
    id serial PRIMARY KEY,
    author varchar(256) not null,
    content text not null,
    created_time TIMESTAMP DEFAULT now(),
    post_id INTEGER,
    FOREIGN KEY (id) REFERENCES posts (id)
);