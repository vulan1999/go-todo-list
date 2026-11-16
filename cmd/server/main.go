package main

import (
	"fmt"
	"net/http"

	"github.com/vulan1999/go-todo-list/internal/handler"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", handler.Healthcheck)
	fmt.Println("Listening on port 3000")
	http.ListenAndServe(":3000", mux)
}
