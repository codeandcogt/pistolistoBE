package ruta

import (
	"pistolistoBE/internal/middleware"

	"github.com/gorilla/mux"
)

func SetupRutaRoutes(api *mux.Router, handler *RutaHandler) {
	rutaRouter := api.PathPrefix("/rutas").Subrouter()

	protected := rutaRouter.NewRoute().Subrouter()
	protected.Use(middleware.AdminJWTMiddleware)

	protected.HandleFunc("", handler.Create).Methods("POST")
	protected.HandleFunc("/all", handler.GetAll).Methods("GET")
	protected.HandleFunc("/{id}", handler.GetByID).Methods("GET")
	protected.HandleFunc("/{id}", handler.Update).Methods("PUT")
	protected.HandleFunc("/{id}/estado", handler.CambiarEstado).Methods("PUT")
	protected.HandleFunc("/{id}", handler.Delete).Methods("DELETE")
}
