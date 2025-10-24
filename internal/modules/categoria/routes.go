package categoria

import (
	"pistolistoBE/internal/middleware"

	"github.com/gorilla/mux"
)

func SetupCategoriaRoutes(api *mux.Router, handler *CategoriaHandler) {
	categoriaRouter := api.PathPrefix("/categoria").Subrouter()

	categoriaRouter.HandleFunc("", handler.CreateCategoria).Methods("POST")

	categoriaRouter.HandleFunc("/all", handler.GetAll).Methods("GET")
	// Rutas protegidas -> subrouter con middleware
	protected := categoriaRouter.NewRoute().Subrouter()
	protected.Use(middleware.AdminJWTMiddleware)

	protected.HandleFunc("/{id}", handler.GetCategoriaByID).Methods("GET")
	protected.HandleFunc("/{id}", handler.DeleteCategoria).Methods("DELETE")
}
