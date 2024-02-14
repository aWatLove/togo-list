CREATE TABLE todo_items (
    id serial PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description VARCHAR(255),
    done BOOLEAN DEFAULT FALSE NOT NULL
);

CREATE TABLE todo_lists (
    id serial PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description VARCHAR(255)
);

CREATE TABLE users (
    id serial PRIMARY KEY,
    NAME VARCHAR(255) NOT NULL,
    username VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL
);

CREATE TABLE users_lists (
    id serial PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users ON DELETE CASCADE,
    list_id INTEGER NOT NULL REFERENCES todo_lists ON DELETE CASCADE
);

CREATE TABLE lists_items (
    id serial PRIMARY KEY,
    item_id INTEGER NOT NULL REFERENCES todo_items ON DELETE CASCADE,
    list_id INTEGER NOT NULL REFERENCES todo_lists ON DELETE CASCADE
);