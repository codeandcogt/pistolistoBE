package factura

import (
	"pistolistoBE/internal/middleware"

	"github.com/gorilla/mux"
)

func SetupFacturaRoutes(api *mux.Router, handler *FacturaHandler) {
	facturaRouter := api.PathPrefix("/facturas").Subrouter()

	protected := facturaRouter.NewRoute().Subrouter()
	protected.Use(middleware.JWTMiddleware)

	protected.HandleFunc("/emitir", handler.EmitirFactura).Methods("POST")
	protected.HandleFunc("/all", handler.GetAll).Methods("GET")
	protected.HandleFunc("/{id}", handler.GetByID).Methods("GET")
	protected.HandleFunc("/pedido/{pedidoId}", handler.GetByPedido).Methods("GET")
	protected.HandleFunc("/{id}", handler.Update).Methods("PUT")
	protected.HandleFunc("/{id}", handler.Delete).Methods("DELETE")
}
