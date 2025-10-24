package departamento

import (
	"pistolistoBE/internal/middleware"

	"github.com/gorilla/mux"
)

func SetupDepartamentoRoutes(api *mux.Router, handler *DepartamentoHandler) {
	departamentoRouter := api.PathPrefix("/departamento").Subrouter()

	// Rutas protegidas -> subrouter con middleware
	protected := departamentoRouter.NewRoute().Subrouter()
	protected.Use(middleware.AdminJWTMiddleware)

	protected.HandleFunc("", handler.CreateDepartamento).Methods("POST")
	protected.HandleFunc("/all", handler.GetAll).Methods("GET")
	protected.HandleFunc("/{id}", handler.GetDepartamentoByID).Methods("GET")
	protected.HandleFunc("/{id}", handler.UpdateDepartamento).Methods("PUT")
	protected.HandleFunc("/{id}", handler.DeleteDepartamento).Methods("DELETE")
}
