package avaluo

import (
	"pistolistoBE/internal/middleware"

	"github.com/gorilla/mux"
)

func SetupAvaluoRoutes(api *mux.Router, handler *AvaluoHandler) {
	router := api.PathPrefix("/avaluos").Subrouter()
	router.Use(middleware.JWTMiddleware)

	router.HandleFunc("", handler.CrearAvaluoManual).Methods("POST")
	router.HandleFunc("/auto/{formularioId}", handler.CrearAvaluoAutomatico).Methods("POST")
}
