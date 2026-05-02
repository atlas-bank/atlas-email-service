package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type EmailPayload struct {
	To      string `json:"to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

func emailHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		fmt.Printf("Método HTTP inválido! Esperado: POST, recebido: %s", r.Method)
		return
	}

	var payload EmailPayload

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, "Erro ao realizar parse do JSON", http.StatusBadRequest)
		return
	}
	fmt.Printf("Lendo da fila: Enviando e-mail para %s...\n", payload.To)
	err = sendEmail(payload)

	if err != nil {
		http.Error(w, "Falha ao enviar e-mail", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Email enviado com sucesso!"))
}

func sendEmail(p EmailPayload) error {
	fmt.Printf("Assunto: %s | Conteúdo: %s\n", p.Subject, p.Body)
	return nil
}

func main() {
	http.HandleFunc("/atlas-web-hook/email-service", emailHandler)

	fmt.Println("Servidor de e-mail rodando na porta 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
