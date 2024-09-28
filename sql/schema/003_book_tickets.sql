-- +goose Up
create table booked_tickets(
    id UUID primary key,
    concert_id UUID references concert_tickets(id) on delete set null,
    user_id UUID references users(id) on delete cascade
);

-- +goose Down
drop table booked_tickets;