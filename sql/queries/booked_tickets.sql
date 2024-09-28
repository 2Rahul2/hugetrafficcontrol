-- name: CreateBookedTickets :one
insert into booked_tickets(id ,concert_id ,user_id)
values($1 ,$2 ,$3)
RETURNING *;

-- name: BookConcertTickets :one
with concert_tickets as(
    update concert_tickets 
    set ticket_available_count = ticket_available_count - 1
    where ticket_available_count > 0 and id = $2
    RETURNING *
)
insert into booked_tickets(id, concert_id ,user_id)
values($1 , $2 ,$3)
RETURNING *;

-- name: BookConcertTicketsOne :one
with concert_tickets_1 as(
    update concert_tickets_1 
    set ticket_available_count = ticket_available_count - 1
    where ticket_available_count > 0 and id = $2
    RETURNING *
)
insert into booked_tickets(id, concert_id ,user_id)
values($1 , $2 ,$3)
RETURNING *;

-- name: BookConcertTicketsTwo :one
with concert_tickets_2 as(
    update concert_tickets_2
    set ticket_available_count = ticket_available_count - 1
    where ticket_available_count > 0 and id = $2
    RETURNING *
)
insert into booked_tickets(id, concert_id ,user_id)
values($1 , $2 ,$3)
RETURNING *;

-- name: BookConcertTicketsThree :one
with concert_tickets_3 as(
    update concert_tickets_3
    set ticket_available_count = ticket_available_count - 1
    where ticket_available_count > 0 and id = $2
    RETURNING *
)
insert into booked_tickets(id, concert_id ,user_id)
values($1 , $2 ,$3)
RETURNING *;

-- name: BookConcertTicketsFour :one
with concert_tickets_4 as(
    update concert_tickets_4
    set ticket_available_count = ticket_available_count - 1
    where ticket_available_count > 0 and id = $2
    RETURNING *
)
insert into booked_tickets(id, concert_id ,user_id)
values($1 , $2 ,$3)
RETURNING *;
