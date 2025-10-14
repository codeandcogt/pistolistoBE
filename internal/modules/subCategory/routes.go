package subCategory

import (
	"pistolistoBE/internal/middleware"

	"github.com/gorilla/mux"
)

func SetupSubCategoryRoutes(api *mux.Router, handler *SubCategoryHandler) {
	subCategoryRouter := api.PathPrefix("/subcategories").Subrouter()

	// Rutas protegidas con JWT
	protected := subCategoryRouter.NewRoute().Subrouter()
	protected.Use(middleware.JWTMiddleware)

	protected.HandleFunc("", handler.Create).Methods("POST")
	protected.HandleFunc("/all", handler.GetAll).Methods("GET")
	protected.HandleFunc("/{id}", handler.GetByID).Methods("GET")
	protected.HandleFunc("/{id}", handler.Update).Methods("PUT")
	protected.HandleFunc("/{id}", handler.Delete).Methods("DELETE")
}
