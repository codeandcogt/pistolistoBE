package subsidiary

import (
	"pistolistoBE/internal/middleware"

	"github.com/gorilla/mux"
)

func SetupSubsidiaryRoutes(api *mux.Router, handler *SubsidiaryHandler) {
	subsidiaryRouter := api.PathPrefix("/subsidiaries").Subrouter()

	// Rutas protegidas con JWT
	protected := subsidiaryRouter.NewRoute().Subrouter()
	protected.Use(middleware.JWTMiddleware)

	protected.HandleFunc("", handler.Create).Methods("POST")
	protected.HandleFunc("/all", handler.GetAll).Methods("GET")
	protected.HandleFunc("/{id}", handler.GetByID).Methods("GET")
	protected.HandleFunc("/{id}", handler.Update).Methods("PUT")
	protected.HandleFunc("/{id}", handler.Delete).Methods("DELETE")
}
