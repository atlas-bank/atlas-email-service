package handler // Pode ser package handler ou main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type EmailPayload struct {
	To      string `json:"to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var payload EmailPayload
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, "Erro ao realizar parse do JSON", http.StatusBadRequest)
		return
	}

	fmt.Printf("Recebido webhook para: %s\n", payload.To)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Webhook processado pela Vercel!"))
}
