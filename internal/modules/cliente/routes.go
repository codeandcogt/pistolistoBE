package cliente

import "github.com/gorilla/mux"

func SetupClienteRoutes(api *mux.Router, handler *ClientHandler) {
	clienteRouter := api.PathPrefix("/clientes").Subrouter()

	clienteRouter.HandleFunc("", handler.CreateCliente).Methods("POST")
	clienteRouter.HandleFunc("/{id}", handler.GetClienteByID).Methods("GET")

}
