package articulo

import (
	"github.com/gorilla/mux"
)

func SetupArticuloRoutes(api *mux.Router, handler *ArticuloHandler) {
	articuloRouter := api.PathPrefix("/articulos").Subrouter()

	// Todas las rutas públicas
	articuloRouter.HandleFunc("/all", handler.GetAll).Methods("GET")
	articuloRouter.HandleFunc("", handler.Create).Methods("POST")
	articuloRouter.HandleFunc("/{id}", handler.GetByID).Methods("GET")
	articuloRouter.HandleFunc("/{id}", handler.Update).Methods("PUT")
	articuloRouter.HandleFunc("/{id}", handler.Delete).Methods("DELETE")
}
