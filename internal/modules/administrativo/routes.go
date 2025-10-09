package administrativo

import (
	"pistolistoBE/internal/middleware"

	"github.com/gorilla/mux"
)

func SetUpAdminRoutes(api *mux.Router, handler *AdministrativoHandler) {
	rolRouter := api.PathPrefix("/admin").Subrouter()

	rolRouter.Use(middleware.JWTMiddleware)

	rolRouter.HandleFunc("", handler.Create).Methods("POST")
	rolRouter.HandleFunc("/all", handler.GetAll).Methods("GET")
	rolRouter.HandleFunc("/{id}", handler.GetByID).Methods("GET")
	rolRouter.HandleFunc("/{id}", handler.Update).Methods("PUT")
	rolRouter.HandleFunc("/{id}", handler.Delete).Methods("DELETE")
	rolRouter.HandleFunc("/firstAuth/{id}", handler.FirstAuth).Methods("PUT")
}
