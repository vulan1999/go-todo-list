package handler

import (
	"log"
	"net/http"

	"github.com/vulan1999/go-todo-list/internal/model"
	"github.com/vulan1999/go-todo-list/views/pages"
)

func Home(w http.ResponseWriter, r *http.Request) {
	todos := []model.Todo{
		{ID: 1, Title: "Todo 1", Done: false},
		{ID: 2, Title: "Todo 2", Done: true},
		{ID: 3, Title: "Todo 3", Done: false},
	}

	err := pages.Home(todos).Render(r.Context(), w)

	if err != nil {
		log.Printf("Fail to render home page: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
