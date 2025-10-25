package contrato

import (
	"pistolistoBE/internal/middleware"

	"github.com/gorilla/mux"
)

func SetupContratoRoutes(api *mux.Router, handler *ContratoHandler) {
	router := api.PathPrefix("/contratos").Subrouter()
	router.Use(middleware.JWTMiddleware)

	router.HandleFunc("/all", handler.GetAll).Methods("GET")
	router.HandleFunc("/avaluo/{avaluoId}", handler.GetByAvaluo).Methods("GET")
	router.HandleFunc("/{id}", handler.GetByID).Methods("GET") // ✅ nuevo
}
