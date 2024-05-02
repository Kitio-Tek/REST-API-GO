package main

import (
	"encoding/json"
	"net/http"
)

func WriteJSON(W http.ResponseWriter,status int, v any) {
	W.Header().Set("Content-Type", "application/json")
	W.WriteHeader(status)
	json.NewEncoder(W).Encode(v)
}	