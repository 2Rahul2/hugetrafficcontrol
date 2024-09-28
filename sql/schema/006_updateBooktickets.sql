-- +goose Up
alter table booked_tickets
drop constraint if exists booked_tickets_concert_id_fkey;

-- +goose Down
alter table booked_tickets
add constraint booked_tickets_concert_id_fkey
foreign key (concert_id) references concert_tickets(id) on delete set null;