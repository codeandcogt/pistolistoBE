package permiso

import (
	"pistolistoBE/internal/middleware"

	"github.com/gorilla/mux"
)

func SetUpPermisoRoutes(api *mux.Router, handler *PermisoHandler) {
	rolRouter := api.PathPrefix("/permiso").Subrouter()

	rolRouter.Use(middleware.AdminJWTMiddleware)

	rolRouter.HandleFunc("", handler.Create).Methods("POST")
	rolRouter.HandleFunc("/all", handler.GetAll).Methods("GET")
	rolRouter.HandleFunc("/{id}", handler.GetByID).Methods("GET")
	rolRouter.HandleFunc("/{id}", handler.Update).Methods("PUT")
	rolRouter.HandleFunc("/{id}", handler.Delete).Methods("DELETE")
}
