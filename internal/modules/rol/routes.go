package rol

import (
	"pistolistoBE/internal/middleware"

	"github.com/gorilla/mux"
)

func SetUpRolRoutes(api *mux.Router, handler *RolHandler) {
	rolRouter := api.PathPrefix("/rol").Subrouter()

	rolRouter.Use(middleware.JWTMiddleware)

	rolRouter.HandleFunc("", handler.CreateRol).Methods("POST")
	rolRouter.HandleFunc("/all", handler.GetAll).Methods("GET")

}
