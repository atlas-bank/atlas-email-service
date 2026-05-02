package main

import (
	"encoding/json"
	"net/http"
)

type Payload struct {
	Message string `json:"message"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	var p Payload
	json.NewDecoder(r.Body).Decode(&p)

	w.Write([]byte("ok"))
}
