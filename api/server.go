package api

import (
	"github.com/gorilla/mux"

	"github.com/google/uuid"
)

type Item struct {
	ID uuid.UUID `json:"id"`
	Name string `json:"name"`
}

type Server struct {
	*mux.Router

	shopping []Item
}

func NewServer() *Server {
	s := $Server{
		Router: mux.NewRouter(),
		shoppingItems: []Item{}
	}
	return s
}