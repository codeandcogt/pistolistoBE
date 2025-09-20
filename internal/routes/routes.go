package routes

import (
	"pistolistoBE/internal/middleware"
	"pistolistoBE/internal/modules/cliente"

	"github.com/gorilla/mux"
)

type RouteHandlers interface {
	GetClienteHandler() *cliente.ClientHandler
}

func SetupRoutes(router *mux.Router, handlers RouteHandlers) {
	router.Use(middleware.Recovery)
	router.Use(middleware.Logger)
	router.Use(middleware.CORS)
	router.Use(middleware.ContentTypeJSON)
	// API versioning
	api := router.PathPrefix("/api").Subrouter()

	// Configurar rutas por módulos
	cliente.SetupClienteRoutes(api, handlers.GetClienteHandler())
}
