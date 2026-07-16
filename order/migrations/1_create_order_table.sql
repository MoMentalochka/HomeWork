-- +goose Up
create table orders (
    id serial primary key,
    order_uuid varchar not null,
    user_uuid varchar,
    part_uuids text,
    total_price double precision,
    transaction_uuid varchar,
    payment_method varchar,
    status varchar
);

-- +goose Down
drop table orders;