-- name: Createconcerts :one
insert into concert_tickets(id , total_tickets ,concert_name ,ticket_available_count)
values($1 ,$2 ,$3 ,$4)
RETURNING *;

-- name: CreateconcertsOne :one
insert into concert_tickets_1(id , total_tickets ,concert_name ,ticket_available_count)
values($1 ,$2 ,$3 ,$4)
RETURNING *;

-- name: CreateconcertsTwo :one
insert into concert_tickets_2(id , total_tickets ,concert_name ,ticket_available_count)
values($1 ,$2 ,$3 ,$4)
RETURNING *;

-- name: CreateconcertsThree :one
insert into concert_tickets_3(id , total_tickets ,concert_name ,ticket_available_count)
values($1 ,$2 ,$3 ,$4)
RETURNING *;

-- name: CreateconcertsFour :one
insert into concert_tickets_4(id , total_tickets ,concert_name ,ticket_available_count)
values($1 ,$2 ,$3 ,$4)
RETURNING *;

-- name: IncrementConcertCount :exec
update concert_tickets
set ticket_available_count = ticket_available_count - 1
where id=$1;
