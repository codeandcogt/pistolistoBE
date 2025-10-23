package piloto

import (
	"pistolistoBE/internal/middleware"

	"github.com/gorilla/mux"
)

func SetupPilotoRoutes(api *mux.Router, handler *PilotoHandler) {
	pilotoRouter := api.PathPrefix("/pilotos").Subrouter()

	// Rutas protegidas con middleware
	pilotoRouter.Use(middleware.JWTMiddleware)

	pilotoRouter.HandleFunc("", handler.CreatePiloto).Methods("POST")
	pilotoRouter.HandleFunc("", handler.GetAll).Methods("GET")
	pilotoRouter.HandleFunc("/{id}", handler.GetPilotoByID).Methods("GET")
	pilotoRouter.HandleFunc("/administrativo/{administrativoId}", handler.GetPilotosByAdministrativoID).Methods("GET")
	pilotoRouter.HandleFunc("/{id}", handler.UpdatePiloto).Methods("PUT")
	pilotoRouter.HandleFunc("/{id}", handler.DeletePiloto).Methods("DELETE")
}
