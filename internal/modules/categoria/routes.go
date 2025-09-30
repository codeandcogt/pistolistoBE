package categoria

import (
	"pistolistoBE/internal/middleware"

	"github.com/gorilla/mux"
)

func SetupCategoriaRoutes(api *mux.Router, handler *CategoriaHandler) {
	categoriaRouter := api.PathPrefix("/categoria").Subrouter()

	categoriaRouter.HandleFunc("", handler.CreateCategoria).Methods("POST")

	// Rutas protegidas -> subrouter con middleware
	protected := categoriaRouter.NewRoute().Subrouter()
	protected.Use(middleware.JWTMiddleware)

	protected.HandleFunc("/all", handler.GetAll).Methods("GET")
	protected.HandleFunc("/{id}", handler.GetCategoriaByID).Methods("GET")
	protected.HandleFunc("/{id}", handler.DeleteCategoria).Methods("DELETE")
}
