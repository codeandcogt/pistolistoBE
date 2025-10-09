package cupon

import (
	"pistolistoBE/internal/middleware"

	"github.com/gorilla/mux"
)

func SetupCuponRoutes(api *mux.Router, handler *CuponHandler) {
	cuponRouter := api.PathPrefix("/cupones").Subrouter()

	// Rutas públicas
	cuponRouter.HandleFunc("", handler.GetAll).Methods("GET")
	cuponRouter.HandleFunc("/{id}", handler.GetCuponByID).Methods("GET")
	cuponRouter.HandleFunc("/codigo/{codigo}", handler.GetCuponByCodigo).Methods("GET")

	// Rutas protegidas
	protected := cuponRouter.NewRoute().Subrouter()
	protected.Use(middleware.JWTMiddleware)

	protected.HandleFunc("", handler.CreateCupon).Methods("POST")
	protected.HandleFunc("/{id}", handler.UpdateCupon).Methods("PUT")
	protected.HandleFunc("/{id}", handler.DeleteCupon).Methods("DELETE")
}
