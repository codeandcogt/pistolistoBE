package cuota

import (
	"pistolistoBE/internal/middleware"

	"github.com/gorilla/mux"
)

func SetupCuotaRoutes(api *mux.Router, handler *CuotaHandler) {
	router := api.PathPrefix("/cuotas").Subrouter()
	protected := router.NewRoute().Subrouter()
	protected.Use(middleware.AdminJWTMiddleware)

	protected.HandleFunc("", handler.Create).Methods("POST")
	protected.HandleFunc("", handler.GetAll).Methods("GET")
	protected.HandleFunc("/{id}", handler.GetByID).Methods("GET")
	protected.HandleFunc("/{id}", handler.Update).Methods("PUT")
	protected.HandleFunc("/{id}", handler.Delete).Methods("DELETE")
}
