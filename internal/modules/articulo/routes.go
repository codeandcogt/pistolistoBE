package articulo

import (
	"pistolistoBE/internal/middleware"

	"github.com/gorilla/mux"
)

func SetupArticuloRoutes(api *mux.Router, handler *ArticuloHandler) {
	articuloRouter := api.PathPrefix("/articulos").Subrouter()

	// Rutas protegidas con JWT
	protected := articuloRouter.NewRoute().Subrouter()
	protected.Use(middleware.AdminJWTMiddleware)

	protected.HandleFunc("", handler.Create).Methods("POST")
	protected.HandleFunc("/all", handler.GetAll).Methods("GET")
	protected.HandleFunc("/{id}", handler.GetByID).Methods("GET")
	protected.HandleFunc("/{id}", handler.Update).Methods("PUT")
	protected.HandleFunc("/{id}", handler.Delete).Methods("DELETE")
}
