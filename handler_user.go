package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/2Rahul2/trafficControl/internal/database"
	"github.com/google/uuid"
)

// rohit ID : 4135cdbc-397b-47a1-9521-d2856f26273a

func (apiCfg apiConfig) handlerUser(w http.ResponseWriter, r *http.Request) {
	type userDetails struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)
	userdetails := userDetails{}
	err := decoder.Decode(&userdetails)
	if err != nil {
		log.Println("Unable to decode json body")
		responsdWithError(w, 400, "ERR: Error parsing json :")
		return
	}

	user, err := apiCfg.DB.Createuser(r.Context(), database.CreateuserParams{
		ID:        uuid.New(),
		Name:      userdetails.Name,
		CreatedAt: time.Now().UTC(),
	})

	if err != nil {
		log.Printf("Error creating user: %v", err)
		responsdWithError(w, 400, fmt.Sprintf("Error creating user: %v", err))
		return
	}
	respondWithJson(w, 201, user)
}
