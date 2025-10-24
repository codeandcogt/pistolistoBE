package producto

import (
	"pistolistoBE/internal/middleware"

	"github.com/gorilla/mux"
)

func SetupProductoRoutes(api *mux.Router, handler *ProductoHandler) {
	productoRouter := api.PathPrefix("/productos").Subrouter()

	productoRouter.HandleFunc("/all", handler.GetAll).Methods("GET")
	productoRouter.HandleFunc("/{id}", handler.GetByID).Methods("GET")

	// Rutas protegidas con JWT
	protected := productoRouter.NewRoute().Subrouter()
	protected.Use(middleware.AdminJWTMiddleware)

	protected.HandleFunc("", handler.Create).Methods("POST")
	protected.HandleFunc("/{id}", handler.Update).Methods("PUT")
	protected.HandleFunc("/{id}", handler.Delete).Methods("DELETE")
}
