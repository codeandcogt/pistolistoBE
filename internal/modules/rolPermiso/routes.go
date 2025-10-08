package rolpermiso

import (
	"pistolistoBE/internal/middleware"

	"github.com/gorilla/mux"
)

func SetUpRolPermisoRoutes(api *mux.Router, handler *RolPermisoHandler) {
	rolPermisoRouter := api.PathPrefix("/rolPermiso").Subrouter()

	rolPermisoRouter.Use(middleware.JWTMiddleware)

	rolPermisoRouter.HandleFunc("", handler.Create).Methods("POST")
	rolPermisoRouter.HandleFunc("/all", handler.GetAll).Methods("GET")
	rolPermisoRouter.HandleFunc("/{id}", handler.GetByID).Methods("GET")
	rolPermisoRouter.HandleFunc("/{id}", handler.Update).Methods("PUT")
	rolPermisoRouter.HandleFunc("/{id}", handler.Delete).Methods("PUT")

}
