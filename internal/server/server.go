package server

import (
	"pistolistoBE/internal/routes"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func NewServer(db *gorm.DB) *Server {
	s := &Server{
		Router: mux.NewRouter(),
		db:     db,
	}

	handlers := s.initializeHandlers()

	routes.SetupRoutes(s.Router, handlers)

	return s
}
