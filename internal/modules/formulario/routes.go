package formulario

import (
	"github.com/gorilla/mux"
)

func SetupFormularioRoutes(api *mux.Router, handler *FormularioHandler) {
	formularioRouter := api.PathPrefix("/formularios").Subrouter()

	// Crear formulario + avalúo automático
	formularioRouter.HandleFunc("", handler.CreateFormulario).Methods("POST")

	// Rutas públicas
	formularioRouter.HandleFunc("", handler.GetAll).Methods("GET")
	formularioRouter.HandleFunc("/{id}", handler.GetFormularioByID).Methods("GET")
	formularioRouter.HandleFunc("/{id}", handler.UpdateFormulario).Methods("PUT")
	formularioRouter.HandleFunc("/{id}", handler.DeleteFormulario).Methods("DELETE")
}
