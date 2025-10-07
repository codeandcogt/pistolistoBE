package departamento

import (
	"pistolistoBE/internal/middleware"

	"github.com/gorilla/mux"
)

func SetupDepartamentoRoutes(api *mux.Router, handler *DepartamentoHandler) {
	departamentoRouter := api.PathPrefix("/departamento").Subrouter()

	departamentoRouter.HandleFunc("", handler.CreateDepartamento).Methods("POST")

	// Rutas protegidas -> subrouter con middleware
	protected := departamentoRouter.NewRoute().Subrouter()
	protected.Use(middleware.JWTMiddleware)

	protected.HandleFunc("/all", handler.GetAll).Methods("GET")
	protected.HandleFunc("/{id}", handler.GetDepartamentoByID).Methods("GET")
	protected.HandleFunc("/{id}", handler.DeleteDepartamento).Methods("DELETE")
}
