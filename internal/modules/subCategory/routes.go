package subCategory

import (
	"pistolistoBE/internal/middleware"

	"github.com/gorilla/mux"
)

func SetupSubCategoryRoutes(api *mux.Router, handler *SubCategoryHandler) {
	subCategoryRouter := api.PathPrefix("/subcategories").Subrouter()

	subCategoryRouter.HandleFunc("/all", handler.GetAll).Methods("GET")
	subCategoryRouter.HandleFunc("/{id}", handler.GetByID).Methods("GET")

	// Rutas protegidas con JWT
	protected := subCategoryRouter.NewRoute().Subrouter()
	protected.Use(middleware.AdminJWTMiddleware)

	protected.HandleFunc("", handler.Create).Methods("POST")
	protected.HandleFunc("/{id}", handler.Update).Methods("PUT")
	protected.HandleFunc("/{id}", handler.Delete).Methods("DELETE")
}
