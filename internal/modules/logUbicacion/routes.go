package logUbicacion

import (
	"pistolistoBE/internal/middleware"

	"github.com/gorilla/mux"
)

func SetupLogUbicacionRoutes(api *mux.Router, handler *LogUbicacionHandler) {
	router := api.PathPrefix("/log-ubicaciones").Subrouter()
	protected := router.NewRoute().Subrouter()
	protected.Use(middleware.AdminJWTMiddleware)

	protected.HandleFunc("", handler.RegistrarUbicacion).Methods("POST")
	protected.HandleFunc("/piloto/{pilotoId}", handler.GetByPiloto).Methods("GET")
}
