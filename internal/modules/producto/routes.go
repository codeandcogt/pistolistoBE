package producto

import (
	"pistolistoBE/internal/middleware"

	"github.com/gorilla/mux"
)

func SetupProductoRoutes(api *mux.Router, handler *ProductoHandler) {
	productoRouter := api.PathPrefix("/productos").Subrouter()

	// Rutas protegidas con JWT
	protected := productoRouter.NewRoute().Subrouter()
	protected.Use(middleware.AdminJWTMiddleware)

	protected.HandleFunc("", handler.Create).Methods("POST")
	protected.HandleFunc("/all", handler.GetAll).Methods("GET")
	protected.HandleFunc("/{id}", handler.GetByID).Methods("GET")
	protected.HandleFunc("/{id}", handler.Update).Methods("PUT")
	protected.HandleFunc("/{id}", handler.Delete).Methods("DELETE")
}
