package main

import (
	"log"
	"net/http"
	"time"

	"github.com/vulan1999/go-todo-list/internal/handler"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", handler.Healthcheck)
	mux.HandleFunc("GET /", handler.Home)
	mux.Handle(
		"GET /static/",
		http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))),
	)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Println("Server listening on port 8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Fail to start server: %v", err)
	}
}
