package config

import (
	"pistolistoBE/internal/modules/cliente"

	"github.com/gorilla/mux"
)

func SetupRoutes(ClientHandler *cliente.ClientHandler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/client", ClientHandler.CreateCliente).Methods("POST")

	return r
}
