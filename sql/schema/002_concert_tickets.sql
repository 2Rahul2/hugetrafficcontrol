-- +goose Up
create table concert_tickets(
    id UUID primary key,
    total_tickets int,
    concert_name text,
    ticket_available_count int
);

-- +goose Down
drop table concert_tickets;