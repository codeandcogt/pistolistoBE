package seccion

import (
	"pistolistoBE/internal/middleware"

	"github.com/gorilla/mux"
)

func SetupSeccionRoutes(api *mux.Router, handler *SeccionHandler) {
	seccionRouter := api.PathPrefix("/seccion").Subrouter()

	// Ruta pública para crear
	seccionRouter.HandleFunc("", handler.CreateSeccion).Methods("POST")

	// Rutas protegidas con middleware JWT
	protected := seccionRouter.NewRoute().Subrouter()
	protected.Use(middleware.JWTMiddleware)

	protected.HandleFunc("", handler.GetAllSecciones).Methods("GET")
	protected.HandleFunc("/{id}", handler.GetSeccionByID).Methods("GET")
	protected.HandleFunc("/{id}", handler.UpdateSeccion).Methods("PUT")
	protected.HandleFunc("/{id}", handler.DeleteSeccion).Methods("DELETE")
}
