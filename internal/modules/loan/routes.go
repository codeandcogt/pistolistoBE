package loan

import (
	"pistolistoBE/internal/middleware"

	"github.com/gorilla/mux"
)

func SetupLoanRoutes(api *mux.Router, handler *LoanHandler) {
	router := api.PathPrefix("/prestamos").Subrouter()
	protected := router.NewRoute().Subrouter()
	protected.Use(middleware.AdminJWTMiddleware)

	protected.HandleFunc("", handler.CreateLoanFromContrato).Methods("POST")
	protected.HandleFunc("", handler.GetAll).Methods("GET")
	protected.HandleFunc("/{id}", handler.GetByID).Methods("GET")
	protected.HandleFunc("/{id}", handler.Update).Methods("PUT")
	protected.HandleFunc("/{id}", handler.Delete).Methods("DELETE")
}
