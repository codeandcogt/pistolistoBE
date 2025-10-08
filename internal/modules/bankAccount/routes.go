package bankAccount

import (
	"pistolistoBE/internal/middleware"

	"github.com/gorilla/mux"
)

func SetupBankAccountRoutes(api *mux.Router, handler *BankAccountHandler) {
	bankAccountRouter := api.PathPrefix("/bank-accounts").Subrouter()

	// Rutas protegidas con JWT
	protected := bankAccountRouter.NewRoute().Subrouter()
	protected.Use(middleware.JWTMiddleware)

	protected.HandleFunc("", handler.Create).Methods("POST")
	protected.HandleFunc("/all", handler.GetAll).Methods("GET")
	protected.HandleFunc("/{id}", handler.GetByID).Methods("GET")
	protected.HandleFunc("/{id}", handler.Update).Methods("PUT")
	protected.HandleFunc("/{id}", handler.Delete).Methods("DELETE")
}
