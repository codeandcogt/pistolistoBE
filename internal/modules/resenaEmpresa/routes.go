package resenaEmpresa

import (
	"pistolistoBE/internal/middleware"

	"github.com/gorilla/mux"
)

func SetupResenaEmpresaRoutes(api *mux.Router, handler *ResenaEmpresaHandler) {
	resenaRouter := api.PathPrefix("/resenas-empresa").Subrouter()

	// Rutas públicas
	resenaRouter.HandleFunc("", handler.GetAll).Methods("GET")
	resenaRouter.HandleFunc("/{id}", handler.GetResenaByID).Methods("GET")
	resenaRouter.HandleFunc("/promedio", handler.GetPromedioCalificacion).Methods("GET")
	resenaRouter.HandleFunc("/cliente/{clienteId}", handler.GetResenasByClienteID).Methods("GET")
	resenaRouter.HandleFunc("/factura/{facturaId}", handler.GetResenaByFacturaID).Methods("GET")

	// Rutas protegidas
	protected := resenaRouter.NewRoute().Subrouter()
	protected.Use(middleware.AdminJWTMiddleware)

	protected.HandleFunc("", handler.CreateResena).Methods("POST")
	protected.HandleFunc("/{id}", handler.UpdateResena).Methods("PUT")
	protected.HandleFunc("/{id}", handler.DeleteResena).Methods("DELETE")
}
