package vehiculo

import (
	"pistolistoBE/internal/middleware"

	"github.com/gorilla/mux"
)

func SetupVehiculoRoutes(api *mux.Router, handler *VehiculoHandler) {
	vehiculoRouter := api.PathPrefix("/vehiculos").Subrouter()

	// Rutas protegidas con middleware
	vehiculoRouter.Use(middleware.AdminJWTMiddleware)

	vehiculoRouter.HandleFunc("", handler.CreateVehiculo).Methods("POST")
	vehiculoRouter.HandleFunc("", handler.GetAll).Methods("GET")
	vehiculoRouter.HandleFunc("/{id}", handler.GetVehiculoByID).Methods("GET")
	vehiculoRouter.HandleFunc("/piloto/{pilotoId}", handler.GetVehiculosByPilotoID).Methods("GET")
	vehiculoRouter.HandleFunc("/{id}", handler.UpdateVehiculo).Methods("PUT")
	vehiculoRouter.HandleFunc("/{id}", handler.DeleteVehiculo).Methods("DELETE")
}
