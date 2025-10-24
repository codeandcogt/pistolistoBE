package almacen

import (
	"pistolistoBE/internal/middleware"

	"github.com/gorilla/mux"
)

func SetupAlmacenRoutes(api *mux.Router, handler *AlmacenHandler) {
	almacenRouter := api.PathPrefix("/almacen").Subrouter()

	// Ruta pública para crear un almacén
	almacenRouter.HandleFunc("", handler.CreateAlmacen).Methods("POST")

	// Subrouter protegido con middleware JWT
	protected := almacenRouter.NewRoute().Subrouter()
	protected.Use(middleware.AdminJWTMiddleware)

	protected.HandleFunc("/sucursal/{id_sucursal}", handler.GetAllBySucursal).Methods("GET")
	protected.HandleFunc("/{id}", handler.GetAlmacenByID).Methods("GET")
	almacenRouter.HandleFunc("", handler.GetAll).Methods("GET")
	protected.HandleFunc("/{id}", handler.UpdateAlmacen).Methods("PUT")
	protected.HandleFunc("/{id}", handler.DeleteAlmacen).Methods("DELETE")
}
