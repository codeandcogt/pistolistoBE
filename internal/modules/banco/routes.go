package banco

import (
	"pistolistoBE/internal/middleware"

	"github.com/gorilla/mux"
)

func SetupBancoRoutes(api *mux.Router, handler *BancoHandler) {
	bancoRouter := api.PathPrefix("/bancos").Subrouter()

	// Rutas públicas
	bancoRouter.HandleFunc("", handler.GetAll).Methods("GET")
	bancoRouter.HandleFunc("/{id}", handler.GetBancoByID).Methods("GET")
	bancoRouter.HandleFunc("/tipo/{tipo}", handler.GetBancosByTipo).Methods("GET")

	// Rutas protegidas -> subrouter con middleware
	protected := bancoRouter.NewRoute().Subrouter()
	protected.Use(middleware.JWTMiddleware)

	protected.HandleFunc("", handler.CreateBanco).Methods("POST")
	protected.HandleFunc("/{id}", handler.UpdateBanco).Methods("PUT")
	protected.HandleFunc("/{id}", handler.DeleteBanco).Methods("DELETE")
}
