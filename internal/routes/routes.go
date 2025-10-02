package routes

import (
	"pistolistoBE/internal/middleware"
	"pistolistoBE/internal/modules/auth"
	"pistolistoBE/internal/modules/categoria"
	"pistolistoBE/internal/modules/cliente"
	"pistolistoBE/internal/modules/departamento"
	"pistolistoBE/internal/modules/descuento"

	"pistolistoBE/internal/modules/rol"

	"github.com/gorilla/mux"
)

type RouteHandlers interface {
	GetClienteHandler() *cliente.ClientHandler
	GetAuthHandler() *auth.AuthHandler
	GetRolHandler() *rol.RolHandler
	GetDepartamentoHandler() *departamento.DepartamentoHandler
	GetCategoriaHandler() *categoria.CategoriaHandler
	GetDescuentoHandler() *descuento.DescuentoHandler
}

func SetupRoutes(router *mux.Router, handlers RouteHandlers) {
	router.Use(middleware.Recovery)
	router.Use(middleware.Logger)
	router.Use(middleware.CORS)
	router.Use(middleware.ContentTypeJSON)
	// API versioning
	api := router.PathPrefix("/api").Subrouter()

	// Configurar rutas por módulos
	cliente.SetupClienteRoutes(api, handlers.GetClienteHandler())
	auth.SetUpAuthRoutes(api, handlers.GetAuthHandler())
	rol.SetUpRolRoutes(api, handlers.GetRolHandler())
	departamento.SetupDepartamentoRoutes(api, handlers.GetDepartamentoHandler())
	categoria.SetupCategoriaRoutes(api, handlers.GetCategoriaHandler())
	descuento.SetupDescuentoRoutes(api, handlers.GetDescuentoHandler())
}
