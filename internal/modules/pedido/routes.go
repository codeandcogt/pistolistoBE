package pedido

import (
	"pistolistoBE/internal/middleware"

	"github.com/gorilla/mux"
)

func SetupPedidoRoutes(api *mux.Router, handler *PedidoHandler) {
	pedidoRouter := api.PathPrefix("/pedidos").Subrouter()

	protected := pedidoRouter.NewRoute().Subrouter()
	protected.Use(middleware.AdminJWTMiddleware)

	protected.HandleFunc("/checkout", handler.Checkout).Methods("POST")
	protected.HandleFunc("/all", handler.GetAll).Methods("GET")
	protected.HandleFunc("/{id}", handler.GetByID).Methods("GET")
	protected.HandleFunc("/cliente/{idCliente}", handler.GetByCliente).Methods("GET")
	protected.HandleFunc("/{id}/estado", handler.CambiarEstado).Methods("PUT")
	protected.HandleFunc("/{id}/cancelar", handler.CancelarPedido).Methods("PUT")
	protected.HandleFunc("/{id}", handler.CancelarPedido).Methods("DELETE") // Baja lógica
}
