package handler

import (
	"log"
	"net/http"

	"github.com/vulan1999/go-todo-list/views/pages"
)

func Home(w http.ResponseWriter, r *http.Request) {
	err := pages.Home().Render(r.Context(), w)

	if err != nil {
		log.Printf("Fail to render home page: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
