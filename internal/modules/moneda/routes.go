package moneda

import (
	"pistolistoBE/internal/middleware"

	"github.com/gorilla/mux"
)

func SetupMonedaRoutes(api *mux.Router, handler *MonedaHandler) {
	monedaRouter := api.PathPrefix("/monedas").Subrouter()

	// Rutas públicas
	monedaRouter.HandleFunc("", handler.GetAll).Methods("GET")
	monedaRouter.HandleFunc("/{id}", handler.GetMonedaByID).Methods("GET")

	// Rutas protegidas -> subrouter con middleware
	protected := monedaRouter.NewRoute().Subrouter()
	protected.Use(middleware.JWTMiddleware)

	protected.HandleFunc("", handler.CreateMoneda).Methods("POST")
	protected.HandleFunc("/{id}", handler.UpdateMoneda).Methods("PUT")
	protected.HandleFunc("/{id}", handler.DeleteMoneda).Methods("DELETE")
}
