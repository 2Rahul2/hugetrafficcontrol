-- +goose Up
create table concert_tickets_1(
    id UUID primary key,
    total_tickets int,
    concert_name text,
    ticket_available_count int
);
create table concert_tickets_2(
    id UUID primary key,
    total_tickets int,
    concert_name text,
    ticket_available_count int
);
create table concert_tickets_3(
    id UUID primary key,
    total_tickets int,
    concert_name text,
    ticket_available_count int
);
create table concert_tickets_4(
    id UUID primary key,
    total_tickets int,
    concert_name text,
    ticket_available_count int
);
-- +goose Down
drop table concert_tickets_1;
drop table concert_tickets_2;
drop table concert_tickets_3;
drop table concert_tickets_4;