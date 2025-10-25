package loan

import (
	"github.com/gorilla/mux"
)

func SetupLoanRoutes(api *mux.Router, handler *LoanHandler) {
	router := api.PathPrefix("/prestamos").Subrouter()

	// Todas las rutas públicas
	router.HandleFunc("", handler.CreateLoanFromContrato).Methods("POST")
	router.HandleFunc("", handler.GetAll).Methods("GET")
	router.HandleFunc("/{id}", handler.GetByID).Methods("GET")
	router.HandleFunc("/{id}", handler.Update).Methods("PUT")
	router.HandleFunc("/{id}", handler.Delete).Methods("DELETE")
}
