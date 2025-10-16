package server

import (
	"net/http"
	"pistolistoBE/internal/middleware"
	"pistolistoBE/internal/routes"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func NewServer(db *gorm.DB) *Server {
	s := &Server{
		Router: mux.NewRouter(),
		db:     db,
	}

	s.Router.Use(middleware.CORS)

	handlers := s.initializeHandlers()

	routes.SetupRoutes(s.Router, handlers)

	// Manejar preflight globalmente
	s.Router.Methods(http.MethodOptions).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})

	return s
}
