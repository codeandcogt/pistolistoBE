package almacenseccion

import (
	"pistolistoBE/internal/middleware"

	"github.com/gorilla/mux"
)

func SetupAlmacenSeccionRoutes(api *mux.Router, handler *AlmacenSeccionHandler) {
	almacenSeccionRouter := api.PathPrefix("/almacenSeccion").Subrouter()

	// Ruta pública para crear
	almacenSeccionRouter.HandleFunc("", handler.CreateAlmacenSeccion).Methods("POST")

	// Rutas protegidas con middleware JWT
	protected := almacenSeccionRouter.NewRoute().Subrouter()
	protected.Use(middleware.JWTMiddleware)

	protected.HandleFunc("/almacen/{id_almacen}", handler.GetAllByAlmacen).Methods("GET")
	protected.HandleFunc("/{id}", handler.GetAlmacenSeccionByID).Methods("GET")
	protected.HandleFunc("/{id}", handler.UpdateAlmacenSeccion).Methods("PUT")
	protected.HandleFunc("/{id}", handler.DeleteAlmacenSeccion).Methods("DELETE")
}
