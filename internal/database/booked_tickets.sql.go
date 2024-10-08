// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: booked_tickets.sql

package database

import (
	"context"

	"github.com/google/uuid"
)

const bookConcertTickets = `-- name: BookConcertTickets :one
with concert_tickets as(
    update concert_tickets 
    set ticket_available_count = ticket_available_count - 1
    where ticket_available_count > 0 and id = $2
    RETURNING id, total_tickets, concert_name, ticket_available_count
)
insert into booked_tickets(id, concert_id ,user_id)
values($1 , $2 ,$3)
RETURNING id, concert_id, user_id
`

type BookConcertTicketsParams struct {
	ID        uuid.UUID
	ConcertID uuid.UUID
	UserID    uuid.UUID
}

func (q *Queries) BookConcertTickets(ctx context.Context, arg BookConcertTicketsParams) (BookedTicket, error) {
	row := q.db.QueryRowContext(ctx, bookConcertTickets, arg.ID, arg.ConcertID, arg.UserID)
	var i BookedTicket
	err := row.Scan(&i.ID, &i.ConcertID, &i.UserID)
	return i, err
}

const bookConcertTicketsFour = `-- name: BookConcertTicketsFour :one
with concert_tickets_4 as(
    update concert_tickets_4
    set ticket_available_count = ticket_available_count - 1
    where ticket_available_count > 0 and id = $2
    RETURNING id, total_tickets, concert_name, ticket_available_count
)
insert into booked_tickets(id, concert_id ,user_id)
values($1 , $2 ,$3)
RETURNING id, concert_id, user_id
`

type BookConcertTicketsFourParams struct {
	ID        uuid.UUID
	ConcertID uuid.UUID
	UserID    uuid.UUID
}

func (q *Queries) BookConcertTicketsFour(ctx context.Context, arg BookConcertTicketsFourParams) (BookedTicket, error) {
	row := q.db.QueryRowContext(ctx, bookConcertTicketsFour, arg.ID, arg.ConcertID, arg.UserID)
	var i BookedTicket
	err := row.Scan(&i.ID, &i.ConcertID, &i.UserID)
	return i, err
}

const bookConcertTicketsOne = `-- name: BookConcertTicketsOne :one
with concert_tickets_1 as(
    update concert_tickets_1 
    set ticket_available_count = ticket_available_count - 1
    where ticket_available_count > 0 and id = $2
    RETURNING id, total_tickets, concert_name, ticket_available_count
)
insert into booked_tickets(id, concert_id ,user_id)
values($1 , $2 ,$3)
RETURNING id, concert_id, user_id
`

type BookConcertTicketsOneParams struct {
	ID        uuid.UUID
	ConcertID uuid.UUID
	UserID    uuid.UUID
}

func (q *Queries) BookConcertTicketsOne(ctx context.Context, arg BookConcertTicketsOneParams) (BookedTicket, error) {
	row := q.db.QueryRowContext(ctx, bookConcertTicketsOne, arg.ID, arg.ConcertID, arg.UserID)
	var i BookedTicket
	err := row.Scan(&i.ID, &i.ConcertID, &i.UserID)
	return i, err
}

const bookConcertTicketsThree = `-- name: BookConcertTicketsThree :one
with concert_tickets_3 as(
    update concert_tickets_3
    set ticket_available_count = ticket_available_count - 1
    where ticket_available_count > 0 and id = $2
    RETURNING id, total_tickets, concert_name, ticket_available_count
)
insert into booked_tickets(id, concert_id ,user_id)
values($1 , $2 ,$3)
RETURNING id, concert_id, user_id
`

type BookConcertTicketsThreeParams struct {
	ID        uuid.UUID
	ConcertID uuid.UUID
	UserID    uuid.UUID
}

func (q *Queries) BookConcertTicketsThree(ctx context.Context, arg BookConcertTicketsThreeParams) (BookedTicket, error) {
	row := q.db.QueryRowContext(ctx, bookConcertTicketsThree, arg.ID, arg.ConcertID, arg.UserID)
	var i BookedTicket
	err := row.Scan(&i.ID, &i.ConcertID, &i.UserID)
	return i, err
}

const bookConcertTicketsTwo = `-- name: BookConcertTicketsTwo :one
with concert_tickets_2 as(
    update concert_tickets_2
    set ticket_available_count = ticket_available_count - 1
    where ticket_available_count > 0 and id = $2
    RETURNING id, total_tickets, concert_name, ticket_available_count
)
insert into booked_tickets(id, concert_id ,user_id)
values($1 , $2 ,$3)
RETURNING id, concert_id, user_id
`

type BookConcertTicketsTwoParams struct {
	ID        uuid.UUID
	ConcertID uuid.UUID
	UserID    uuid.UUID
}

func (q *Queries) BookConcertTicketsTwo(ctx context.Context, arg BookConcertTicketsTwoParams) (BookedTicket, error) {
	row := q.db.QueryRowContext(ctx, bookConcertTicketsTwo, arg.ID, arg.ConcertID, arg.UserID)
	var i BookedTicket
	err := row.Scan(&i.ID, &i.ConcertID, &i.UserID)
	return i, err
}

const createBookedTickets = `-- name: CreateBookedTickets :one
insert into booked_tickets(id ,concert_id ,user_id)
values($1 ,$2 ,$3)
RETURNING id, concert_id, user_id
`

type CreateBookedTicketsParams struct {
	ID        uuid.UUID
	ConcertID uuid.UUID
	UserID    uuid.UUID
}

func (q *Queries) CreateBookedTickets(ctx context.Context, arg CreateBookedTicketsParams) (BookedTicket, error) {
	row := q.db.QueryRowContext(ctx, createBookedTickets, arg.ID, arg.ConcertID, arg.UserID)
	var i BookedTicket
	err := row.Scan(&i.ID, &i.ConcertID, &i.UserID)
	return i, err
}
