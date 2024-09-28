package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/2Rahul2/trafficControl/internal/database"
	"github.com/google/uuid"
)

func (apiCfg apiConfig) handlerCreateShardsConcert(w http.ResponseWriter, r *http.Request) {
	type concertDatabase struct {
		id                     uuid.UUID
		total_tickets          int32
		concert_name           string
		ticket_available_count int32
	}
	var concertDatabaseList []concertDatabase
	type getConcertDetails struct {
		TotalTickets int    `json:"total_tickets"`
		ConcertName  string `json:"concert_name"`
	}

	decoder := json.NewDecoder(r.Body)
	get_concert_details := getConcertDetails{}
	err := decoder.Decode(&get_concert_details)
	if err != nil {
		responsdWithError(w, 400, fmt.Sprintf("ERR: Could not decode the json data : %v", err))
		return
	}
	var concertID uuid.UUID = uuid.New()
	var shardTotalTickets int = get_concert_details.TotalTickets / 4

	// Concert One
	concertOne, err := apiCfg.DB.CreateconcertsOne(r.Context(), database.CreateconcertsOneParams{
		ID: concertID,
		TotalTickets: sql.NullInt32{
			Int32: int32(shardTotalTickets),
			Valid: true,
		},
		ConcertName: sql.NullString{
			String: get_concert_details.ConcertName,
			Valid:  true,
		},
		TicketAvailableCount: sql.NullInt32{
			Int32: int32(shardTotalTickets),
			Valid: true,
		},
	})
	if err != nil {
		responsdWithError(w, 400, fmt.Sprintf("ERR: Could not add concert details : %v", err))
		return
	}
	concertDatabaseList = append(concertDatabaseList, concertDatabase{
		id:                     concertOne.ID,
		total_tickets:          concertOne.TotalTickets.Int32,
		concert_name:           concertOne.ConcertName.String,
		ticket_available_count: concertOne.TicketAvailableCount.Int32,
	})

	// Concert Two

	concertTwo, err := apiCfg.DB.CreateconcertsTwo(r.Context(), database.CreateconcertsTwoParams{
		ID: concertID,
		TotalTickets: sql.NullInt32{
			Int32: int32(shardTotalTickets),
			Valid: true,
		},
		ConcertName: sql.NullString{
			String: get_concert_details.ConcertName,
			Valid:  true,
		},
		TicketAvailableCount: sql.NullInt32{
			Int32: int32(shardTotalTickets),
			Valid: true,
		},
	})
	if err != nil {
		responsdWithError(w, 400, fmt.Sprintf("ERR: Could not add concert details : %v", err))
		return
	}
	concertDatabaseList = append(concertDatabaseList, concertDatabase{
		id:                     concertTwo.ID,
		total_tickets:          concertTwo.TotalTickets.Int32,
		concert_name:           concertTwo.ConcertName.String,
		ticket_available_count: concertTwo.TicketAvailableCount.Int32,
	})

	// Concert Three
	concertThree, err := apiCfg.DB.CreateconcertsThree(r.Context(), database.CreateconcertsThreeParams{
		ID: concertID,
		TotalTickets: sql.NullInt32{
			Int32: int32(shardTotalTickets),
			Valid: true,
		},
		ConcertName: sql.NullString{
			String: get_concert_details.ConcertName,
			Valid:  true,
		},
		TicketAvailableCount: sql.NullInt32{
			Int32: int32(shardTotalTickets),
			Valid: true,
		},
	})
	if err != nil {
		responsdWithError(w, 400, fmt.Sprintf("ERR: Could not add concert details : %v", err))
		return
	}
	concertDatabaseList = append(concertDatabaseList, concertDatabase{
		id:                     concertThree.ID,
		total_tickets:          concertThree.TotalTickets.Int32,
		concert_name:           concertThree.ConcertName.String,
		ticket_available_count: concertThree.TicketAvailableCount.Int32,
	})

	// Concert Four

	concertFour, err := apiCfg.DB.CreateconcertsFour(r.Context(), database.CreateconcertsFourParams{
		ID: concertID,
		TotalTickets: sql.NullInt32{
			Int32: int32(shardTotalTickets),
			Valid: true,
		},
		ConcertName: sql.NullString{
			String: get_concert_details.ConcertName,
			Valid:  true,
		},
		TicketAvailableCount: sql.NullInt32{
			Int32: int32(shardTotalTickets),
			Valid: true,
		},
	})
	if err != nil {
		responsdWithError(w, 400, fmt.Sprintf("ERR: Could not add concert details : %v", err))
		return
	}
	concertDatabaseList = append(concertDatabaseList, concertDatabase{
		id:                     concertFour.ID,
		total_tickets:          concertFour.TotalTickets.Int32,
		concert_name:           concertFour.ConcertName.String,
		ticket_available_count: concertFour.TicketAvailableCount.Int32,
	})

	respondWithJson(w, 201, concertDatabaseList)
}

// "ID": "7cbe24cd-f012-4552-b496-a51af2a8d26d",

func (apiCfg apiConfig) handlerCreateConcert(w http.ResponseWriter, r *http.Request) {
	type getConcertDetails struct {
		TotalTickets int    `json:"total_tickets"`
		ConcertName  string `json:"concert_name"`
	}

	decoder := json.NewDecoder(r.Body)
	get_concert_details := getConcertDetails{}
	err := decoder.Decode(&get_concert_details)
	if err != nil {
		responsdWithError(w, 400, fmt.Sprintf("ERR: Could not decode the json data : %v", err))
		return
	}

	concert, err := apiCfg.DB.Createconcerts(r.Context(), database.CreateconcertsParams{
		ID: uuid.New(),
		ConcertName: sql.NullString{
			String: get_concert_details.ConcertName,
			Valid:  get_concert_details.ConcertName != "",
		},
		TotalTickets: sql.NullInt32{
			Int32: int32(get_concert_details.TotalTickets),
			Valid: true,
		},
		TicketAvailableCount: sql.NullInt32{
			Int32: int32(get_concert_details.TotalTickets),
			Valid: true,
		},
	})
	if err != nil {
		responsdWithError(w, 400, fmt.Sprintf("ERR: Could not add concert details : %v", err))
		return
	}
	respondWithJson(w, 201, concert)
}
