package subsidiary

import (
	"pistolistoBE/internal/middleware"

	"github.com/gorilla/mux"
)

func SetupSubsidiaryRoutes(api *mux.Router, handler *SubsidiaryHandler) {
	subsidiaryRouter := api.PathPrefix("/subsidiaries").Subrouter()

	subsidiaryRouter.HandleFunc("/all", handler.GetAll).Methods("GET")
	// Rutas protegidas con JWT
	protected := subsidiaryRouter.NewRoute().Subrouter()
	protected.Use(middleware.AdminJWTMiddleware)

	protected.HandleFunc("", handler.Create).Methods("POST")
	protected.HandleFunc("/{id}", handler.GetByID).Methods("GET")
	protected.HandleFunc("/{id}", handler.Update).Methods("PUT")
	protected.HandleFunc("/{id}", handler.Delete).Methods("DELETE")
}
