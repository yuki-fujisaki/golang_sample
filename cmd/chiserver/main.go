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
}
