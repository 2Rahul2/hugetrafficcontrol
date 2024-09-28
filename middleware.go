package main

import (
	"fmt"
	"net/http"

	"github.com/2Rahul2/trafficControl/internal/database"
	"github.com/google/uuid"
)

type userHandler func(w http.ResponseWriter, r *http.Request, user database.User)

func (apiCfg *apiConfig) middlewareHandler(handler userHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var id string = r.Header.Get("user_id")
		uid, err := uuid.Parse(id)
		if err != nil {
			responsdWithError(w, 400, fmt.Sprintf("ERR : could not parse ID : %v", err))
			return
		}
		user, err := apiCfg.DB.GetUser(r.Context(), uid)
		if err != nil {
			responsdWithError(w, 400, fmt.Sprintf("ERR : could not get the user : %v", err))
			return
		}
		handler(w, r, user)

	}
}
