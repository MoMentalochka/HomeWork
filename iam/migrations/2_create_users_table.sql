-- +goose Up
create table users (
    id serial primary key,
    login varchar not null,
    password varchar not null,
    email varchar,
    notification_methods text
);

-- +goose Down
drop table users;