package estadoPedido

import (
	"pistolistoBE/internal/middleware"

	"github.com/gorilla/mux"
)

func SetupEstadoPedidoRoutes(api *mux.Router, handler *EstadoPedidoHandler) {
	estadoPedidoRouter := api.PathPrefix("/estado-pedido").Subrouter()

	// Rutas protegidas con JWT
	protected := estadoPedidoRouter.NewRoute().Subrouter()
	protected.Use(middleware.AdminJWTMiddleware)

	protected.HandleFunc("", handler.Create).Methods("POST")
	protected.HandleFunc("/all", handler.GetAll).Methods("GET")
	protected.HandleFunc("/{id}", handler.GetByID).Methods("GET")
	protected.HandleFunc("/{id}", handler.Update).Methods("PUT")
	protected.HandleFunc("/{id}", handler.Delete).Methods("DELETE")
}
