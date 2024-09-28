package main

import (
	"github.com/2Rahul2/trafficControl/internal/database"
	"github.com/google/uuid"
)

type bookTicket struct {
	ID        uuid.UUID
	ConcertID uuid.UUID
	UserID    uuid.UUID
}

func databaseTicketToTickets(dbTicket database.BookedTicket) bookTicket {
	return bookTicket{
		ID:        dbTicket.ID,
		ConcertID: dbTicket.ConcertID,
		UserID:    dbTicket.UserID,
	}
}
