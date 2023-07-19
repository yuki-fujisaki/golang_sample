package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

type Message struct {
	Message string `json:"message"`
}

func main() {
	r := chi.NewRouter()

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		msg := Message{Message: "hello"}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(msg)
	})

	fmt.Println("server is running...")
	http.ListenAndServe(":3000", r)
}
