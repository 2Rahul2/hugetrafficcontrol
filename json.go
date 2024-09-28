package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func responsdWithError(w http.ResponseWriter, code int, errMsg string) {
	type errResponse struct {
		Error string `json:"error"`
	}
	fmt.Println(errMsg)
	respondWithJson(w, code, errResponse{
		Error: errMsg,
	})
}
func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(500)
		log.Printf("Failed to Marshal json response :%v", payload)
		return
	}

	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}
