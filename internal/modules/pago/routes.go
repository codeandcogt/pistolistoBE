package pago

import (
	"pistolistoBE/internal/middleware"

	"github.com/gorilla/mux"
)

func SetupPagoRoutes(api *mux.Router, handler *PagoHandler) {
	pagoRouter := api.PathPrefix("/pagos").Subrouter()

	protected := pagoRouter.NewRoute().Subrouter()
	protected.Use(middleware.JWTMiddleware)

	protected.HandleFunc("/procesar", handler.ProcesarPago).Methods("POST")
	protected.HandleFunc("/all", handler.GetAll).Methods("GET")
	protected.HandleFunc("/{id}", handler.GetByID).Methods("GET")
	protected.HandleFunc("/pedido/{pedidoId}", handler.GetByPedido).Methods("GET")
	protected.HandleFunc("/{id}", handler.Update).Methods("PUT")
	protected.HandleFunc("/{id}", handler.Delete).Methods("DELETE")
}
