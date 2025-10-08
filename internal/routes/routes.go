package routes

import (
	"pistolistoBE/internal/middleware"
	"pistolistoBE/internal/modules/auth"
	"pistolistoBE/internal/modules/banco"
	"pistolistoBE/internal/modules/carrito"
	"pistolistoBE/internal/modules/cliente"
	"pistolistoBE/internal/modules/cupon"
	"pistolistoBE/internal/modules/moneda"
	"pistolistoBE/internal/modules/rol"

	"github.com/gorilla/mux"
)

type RouteHandlers interface {
	GetClienteHandler() *cliente.ClientHandler
	GetAuthHandler() *auth.AuthHandler
	GetMonedaHandler() *moneda.MonedaHandler
	GetBancoHandler() *banco.BancoHandler
	GetRolHandler() *rol.RolHandler
	GetCuponHandler() *cupon.CuponHandler
	GetCarritoHandler() *carrito.CarritoHandler
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
	moneda.SetupMonedaRoutes(api, handlers.GetMonedaHandler())
	banco.SetupBancoRoutes(api, handlers.GetBancoHandler())
	rol.SetUpRolRoutes(api, handlers.GetRolHandler())
	cupon.SetupCuponRoutes(api, handlers.GetCuponHandler())
	carrito.SetupCarritoRoutes(api, handlers.GetCarritoHandler())
}
