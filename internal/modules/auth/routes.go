package auth

import (
	"github.com/gorilla/mux"
)

func SetUpAuthRoutes(api *mux.Router, handler *AuthHandler) {
	auth := api.PathPrefix("/auth").Subrouter()

	auth.HandleFunc("/client", handler.LoginClient).Methods("POST")
}
