package formulario

import (
	"pistolistoBE/internal/middleware"

	"github.com/gorilla/mux"
)

func SetupFormularioRoutes(api *mux.Router, handler *FormularioHandler) {
	formularioRouter := api.PathPrefix("/formularios").Subrouter()

	// Ruta pública (crear formulario + avalúo automático)
	formularioRouter.HandleFunc("", handler.CreateFormulario).Methods("POST")

	// Rutas protegidas para administración
	protected := formularioRouter.NewRoute().Subrouter()
	protected.Use(middleware.JWTMiddleware)

	protected.HandleFunc("", handler.GetAll).Methods("GET")
	protected.HandleFunc("/{id}", handler.GetFormularioByID).Methods("GET")
	protected.HandleFunc("/{id}", handler.UpdateFormulario).Methods("PUT")
	protected.HandleFunc("/{id}", handler.DeleteFormulario).Methods("DELETE")
}
