package estadoRuta

import (
	"pistolistoBE/internal/middleware"

	"github.com/gorilla/mux"
)

func SetupEstadoRutaRoutes(api *mux.Router, handler *EstadoRutaHandler) {
	estadoRutaRouter := api.PathPrefix("/estado-ruta").Subrouter()

	protected := estadoRutaRouter.NewRoute().Subrouter()
	protected.Use(middleware.AdminJWTMiddleware)

	protected.HandleFunc("", handler.Create).Methods("POST")
	protected.HandleFunc("/all", handler.GetAll).Methods("GET")
	protected.HandleFunc("/{id}", handler.GetByID).Methods("GET")
	protected.HandleFunc("/{id}", handler.Update).Methods("PUT")
	protected.HandleFunc("/{id}", handler.Delete).Methods("DELETE")
}
