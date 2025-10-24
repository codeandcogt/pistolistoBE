package municipality

import (
	"pistolistoBE/internal/middleware"

	"github.com/gorilla/mux"
)

func SetupMunicipalityRoutes(api *mux.Router, handler *MunicipalityHandler) {
	municipalityRouter := api.PathPrefix("/municipalities").Subrouter()

	municipalityRouter.HandleFunc("/all", handler.GetAll).Methods("GET")
	municipalityRouter.HandleFunc("/{id}", handler.GetByID).Methods("GET")

	// Rutas protegidas con JWT
	protected := municipalityRouter.NewRoute().Subrouter()
	protected.Use(middleware.AdminJWTMiddleware)

	protected.HandleFunc("", handler.Create).Methods("POST")
	protected.HandleFunc("/{id}", handler.Update).Methods("PUT")
	protected.HandleFunc("/{id}", handler.Delete).Methods("DELETE")
}
