package main

import (
	"log"
	"net/http"

	"github.com/moaabb/broker-service/data"
)

func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	payload := JsonResponse{
		Error:   false,
		Message: "Hit the Broker!",
	}

	err := data.WriteJSON(w, payload, http.StatusAccepted)
	if err != nil {
		log.Println()
		return
	}
}
