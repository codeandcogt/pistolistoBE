package direccion

import (
	"pistolistoBE/internal/middleware"

	"github.com/gorilla/mux"
)

func SetupDireccionRoutes(api *mux.Router, handler *DireccionHandler) {
	direccionRouter := api.PathPrefix("/direcciones").Subrouter()

	direccionRouter.HandleFunc("", handler.Create).Methods("POST")

	// Rutas protegidas con JWT
	protected := direccionRouter.NewRoute().Subrouter()
	protected.Use(middleware.JWTMiddleware)

	protected.HandleFunc("/all", handler.GetAll).Methods("GET")
	protected.HandleFunc("/{id}", handler.GetByID).Methods("GET")
	protected.HandleFunc("/{id}", handler.Update).Methods("PUT")
	protected.HandleFunc("/{id}", handler.Delete).Methods("DELETE")
}
