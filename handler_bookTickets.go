package main

import (
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"sync"

	"github.com/2Rahul2/trafficControl/internal/database"
	"github.com/google/uuid"
)

//	{
//		"ID": "b10ca959-f6e0-49c1-a7ec-923dc34175a7",
//		"ConcertID": "7cbe24cd-f012-4552-b496-a51af2a8d26d",
//		"UserID": "4135cdbc-397b-47a1-9521-d2856f26273a"
//	  }

func getShard(userId uuid.UUID) int {
	n := new(big.Int)
	// var concertId uuid.UUID = uuid.New()
	n.SetBytes(userId[:])
	shard := n.Mod(n, big.NewInt(2)).Int64() + 1
	return int(shard)
}

// 13.163101195s   1m45.211261984s   2m2.846188592s
func (apiCfg apiConfig) handlerBookTicketsWithShardin(w http.ResponseWriter, r *http.Request, user database.User) {
	type booking_details struct {
		Concert_Id uuid.UUID `json:"concert_id"`
	}
	decoder := json.NewDecoder(r.Body)
	bookingDetails := booking_details{}
	err := decoder.Decode(&bookingDetails)
	if err != nil {
		responsdWithError(w, 400, fmt.Sprintf("Could not decode json data : %v", err))
		return
	}
	// fmt.Println(bookingDetails, user)
	switch shardNum := getShard(user.ID); shardNum {
	case 1:
		bookingTicket1, err := apiCfg.DB.BookConcertTicketsOne(r.Context(), database.BookConcertTicketsOneParams{
			ID:        uuid.New(),
			ConcertID: bookingDetails.Concert_Id,
			UserID:    user.ID,
		})
		if err != nil {
			responsdWithError(w, 400, fmt.Sprintf("Could not book tickets : %v", err))
			return
		}
		respondWithJson(w, 201, bookingTicket1)
	case 2:
		bookingTicket2, err := apiCfg.DB.BookConcertTicketsTwo(r.Context(), database.BookConcertTicketsTwoParams{
			ID:        uuid.New(),
			ConcertID: bookingDetails.Concert_Id,
			UserID:    user.ID,
		})
		if err != nil {
			responsdWithError(w, 400, fmt.Sprintf("Could not book tickets : %v", err))
			return
		}
		respondWithJson(w, 201, bookingTicket2)
	case 3:
		bookingTicket3, err := apiCfg.DB.BookConcertTicketsThree(r.Context(), database.BookConcertTicketsThreeParams{
			ID:        uuid.New(),
			ConcertID: bookingDetails.Concert_Id,
			UserID:    user.ID,
		})
		if err != nil {
			responsdWithError(w, 400, fmt.Sprintf("Could not book tickets : %v", err))
			return
		}
		respondWithJson(w, 201, bookingTicket3)
	case 4:
		bookingTicket4, err := apiCfg.DB.BookConcertTicketsFour(r.Context(), database.BookConcertTicketsFourParams{
			ID:        uuid.New(),
			ConcertID: bookingDetails.Concert_Id,
			UserID:    user.ID,
		})
		if err != nil {
			responsdWithError(w, 400, fmt.Sprintf("Could not book tickets : %v", err))
			return
		}
		respondWithJson(w, 201, bookingTicket4)
	default:
		responsdWithError(w, 400, "Could not get valid shard number")
	}
}

// 3m24.597520958s
func (apiCfg apiConfig) handlerBookTicketsUni(w http.ResponseWriter, r *http.Request, user database.User) {
	type booking_details struct {
		Concert_Id uuid.UUID `json:"concert_id"`
	}

	decoder := json.NewDecoder(r.Body)
	bookingDetails := booking_details{}
	err := decoder.Decode(&bookingDetails)
	if err != nil {
		responsdWithError(w, 400, fmt.Sprintf("Could not decode json data : %v", err))
		return
	}

	bookTicket, err := apiCfg.DB.BookConcertTickets(r.Context(), database.BookConcertTicketsParams{
		ID:        uuid.New(),
		ConcertID: bookingDetails.Concert_Id,
		UserID:    user.ID,
	})
	if err != nil {
		responsdWithError(w, 400, fmt.Sprintf("Could not book tickets : %v", err))
		return
	}
	respondWithJson(w, 201, databaseTicketToTickets(bookTicket))

}

// 3m37.444000299s
func (apiCfg *apiConfig) handlerBookTicketsNoGo(w http.ResponseWriter, r *http.Request, user database.User) {
	type booking_details struct {
		Concert_Id uuid.UUID `json:"concert_id"`
	}

	decoder := json.NewDecoder(r.Body)
	fmt.Println(r.Body)
	bookingDetails := booking_details{}
	err := decoder.Decode(&bookingDetails)
	if err != nil {
		responsdWithError(w, 400, fmt.Sprintf("Could not decode json data : %v", err))
		return
	}

	tx, err := apiCfg.DBConn.BeginTx(r.Context(), nil)
	if err != nil {
		responsdWithError(w, 400, fmt.Sprintf("Could not issue a transaction : %v", err))
		return
	}

	defer tx.Rollback()
	q := apiCfg.DB.WithTx(tx)
	bookTicket, err := q.CreateBookedTickets(r.Context(), database.CreateBookedTicketsParams{
		ID:        uuid.New(),
		ConcertID: bookingDetails.Concert_Id,
		UserID:    user.ID,
	})
	if err != nil {
		tx.Rollback()
		responsdWithError(w, 400, err.Error())
		return
	}
	err = q.IncrementConcertCount(r.Context(), bookingDetails.Concert_Id)
	if err != nil {
		tx.Rollback()
		responsdWithError(w, 400, fmt.Sprintf("Could not find the concert : %v", err))
		return
	}

	if err = tx.Commit(); err != nil {
		responsdWithError(w, 400, fmt.Sprintf("Error during commiting transaction : %v", err))
		return
	}
	respondWithJson(w, 201, databaseTicketToTickets(bookTicket))

}

// 3m38.980306375s 3m49.22022261s
func (apiCfg *apiConfig) handlerBookTickets(w http.ResponseWriter, r *http.Request, user database.User) {
	type booking_details struct {
		Concert_Id uuid.UUID `json:"concert_id"`
	}
	decoder := json.NewDecoder(r.Body)
	bookingDetails := booking_details{}
	err := decoder.Decode(&bookingDetails)
	if err != nil {
		responsdWithError(w, 400, fmt.Sprintf("Could not decode json data : %v", err))
		return
	}

	tx, err := apiCfg.DBConn.BeginTx(r.Context(), nil)
	if err != nil {
		responsdWithError(w, 400, fmt.Sprintf("Could not issue a transaction : %v", err))
		return
	}

	q := apiCfg.DB.WithTx(tx)
	var wg sync.WaitGroup
	var bookingError error
	var bookTicket database.BookedTicket

	wg.Add(1)
	go func() {
		defer wg.Done()
		bookTicket, bookingError = q.CreateBookedTickets(r.Context(), database.CreateBookedTicketsParams{
			ID:        uuid.New(),
			ConcertID: bookingDetails.Concert_Id,
			UserID:    user.ID,
		})
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		err := q.IncrementConcertCount(r.Context(), bookingDetails.Concert_Id)
		if err != nil {
			responsdWithError(w, 400, fmt.Sprintf("Could not find the concert : %v", err))
			return
		}

	}()

	wg.Wait()

	if bookingError != nil {
		tx.Rollback()
		responsdWithError(w, 400, bookingError.Error())
		return
	}

	if err := tx.Commit(); err != nil {
		responsdWithError(w, 400, fmt.Sprintf("Error during commiting transaction : %v", err))
		return
	}

	respondWithJson(w, 201, databaseTicketToTickets(bookTicket))
}
