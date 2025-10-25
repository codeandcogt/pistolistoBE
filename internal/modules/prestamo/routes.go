package prestamo

import (
	"pistolistoBE/internal/middleware"

	"github.com/gorilla/mux"
)

func SetupPrestamoRoutes(api *mux.Router, handler *PrestamoHandler) {
	prestamoRouter := api.PathPrefix("/prestamos").Subrouter()

	// Rutas protegidas con middleware
	prestamoRouter.Use(middleware.JWTMiddleware)

	prestamoRouter.HandleFunc("", handler.CreatePrestamo).Methods("POST")
	prestamoRouter.HandleFunc("", handler.GetAll).Methods("GET")
	prestamoRouter.HandleFunc("/{id}", handler.GetPrestamoByID).Methods("GET")
	prestamoRouter.HandleFunc("/cliente/{clienteId}", handler.GetPrestamosByClienteID).Methods("GET")
	prestamoRouter.HandleFunc("/{id}", handler.UpdatePrestamo).Methods("PUT")
	prestamoRouter.HandleFunc("/{id}", handler.DeletePrestamo).Methods("DELETE")
}
