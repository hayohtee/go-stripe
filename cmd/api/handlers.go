package main

import (
	"encoding/json"
	"net/http"
)

type stripePayload struct {
	Currency string `json:"currency"`
	Amount   string `json:"amount"`
}

type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
	Content string `json:"content"`
	ID      int    `json:"id"`	
}

func (app *application) getPaymentIntent(w http.ResponseWriter, r *http.Request) {
	resp := jsonResponse{
		OK: true,

	}

	out, err := json.MarshalIndent(resp, "", "\t")
	if err != nil {
		app.errorLog.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}