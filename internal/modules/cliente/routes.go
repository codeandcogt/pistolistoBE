package cliente

import (
	"pistolistoBE/internal/middleware"

	"github.com/gorilla/mux"
)

func SetupClienteRoutes(api *mux.Router, handler *ClientHandler) {
	clienteRouter := api.PathPrefix("/clientes").Subrouter()

	clienteRouter.HandleFunc("", handler.CreateCliente).Methods("POST")

	// Rutas protegidas -> subrouter con middleware
	protected := clienteRouter.NewRoute().Subrouter()
	protected.Use(middleware.JWTMiddleware)

	protected.HandleFunc("/all", handler.GetAll).Methods("GET")
	protected.HandleFunc("/{id}", handler.GetClienteByID).Methods("GET")
}
