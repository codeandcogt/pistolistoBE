package descuento

import (
	"pistolistoBE/internal/middleware"

	"github.com/gorilla/mux"
)

func SetupDescuentoRoutes(api *mux.Router, handler *DescuentoHandler) {
	descuentoRouter := api.PathPrefix("/descuento").Subrouter()

	descuentoRouter.HandleFunc("", handler.CreateDescuento).Methods("POST")

	// Rutas protegidas -> subrouter con middleware
	protected := descuentoRouter.NewRoute().Subrouter()
	protected.Use(middleware.JWTMiddleware)

	protected.HandleFunc("/all", handler.GetAll).Methods("GET")
	protected.HandleFunc("/{id}", handler.GetDescuentoByID).Methods("GET")
	protected.HandleFunc("/{id}", handler.DeleteDescuento).Methods("DELETE")
}
