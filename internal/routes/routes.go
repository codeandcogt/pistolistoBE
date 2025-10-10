package routes

import (
	"pistolistoBE/internal/middleware"
	"pistolistoBE/internal/modules/administrativo"
	"pistolistoBE/internal/modules/auth"
	"pistolistoBE/internal/modules/banco"
	"pistolistoBE/internal/modules/bankAccount"
	"pistolistoBE/internal/modules/carrito"
	"pistolistoBE/internal/modules/categoria"
	"pistolistoBE/internal/modules/cliente"
	"pistolistoBE/internal/modules/cupon"
	"pistolistoBE/internal/modules/departamento"
	"pistolistoBE/internal/modules/descuento"
	"pistolistoBE/internal/modules/moneda"
	"pistolistoBE/internal/modules/municipality"
	"pistolistoBE/internal/modules/permiso"
	"pistolistoBE/internal/modules/rol"
	rolpermiso "pistolistoBE/internal/modules/rolPermiso"
	"pistolistoBE/internal/modules/subsidiary"

	"github.com/gorilla/mux"
)

type RouteHandlers interface {
	GetClienteHandler() *cliente.ClientHandler
	GetAuthHandler() *auth.AuthHandler
	GetMonedaHandler() *moneda.MonedaHandler
	GetBancoHandler() *banco.BancoHandler
	GetRolHandler() *rol.RolHandler
	GetSubsidiaryHandler() *subsidiary.SubsidiaryHandler
	GetDepartamentoHandler() *departamento.DepartamentoHandler
	GetCategoriaHandler() *categoria.CategoriaHandler
	GetDescuentoHandler() *descuento.DescuentoHandler
	GetCuponHandler() *cupon.CuponHandler
	GetCarritoHandler() *carrito.CarritoHandler
	GetMunicipalityHandler() *municipality.MunicipalityHandler
	GetBankAccountHandler() *bankAccount.BankAccountHandler
	GetPermisoHandler() *permiso.PermisoHandler
	GetRolPermisoHandler() *rolpermiso.RolPermisoHandler
	GetAdminHandler() *administrativo.AdministrativoHandler
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
	subsidiary.SetupSubsidiaryRoutes(api, handlers.GetSubsidiaryHandler())
	departamento.SetupDepartamentoRoutes(api, handlers.GetDepartamentoHandler())
	categoria.SetupCategoriaRoutes(api, handlers.GetCategoriaHandler())
	descuento.SetupDescuentoRoutes(api, handlers.GetDescuentoHandler())
	cupon.SetupCuponRoutes(api, handlers.GetCuponHandler())
	carrito.SetupCarritoRoutes(api, handlers.GetCarritoHandler())
	municipality.SetupMunicipalityRoutes(api, handlers.GetMunicipalityHandler())
	bankAccount.SetupBankAccountRoutes(api, handlers.GetBankAccountHandler())
	permiso.SetUpPermisoRoutes(api, handlers.GetPermisoHandler())
	rolpermiso.SetUpRolPermisoRoutes(api, handlers.GetRolPermisoHandler())
	administrativo.SetUpAdminRoutes(api, handlers.GetAdminHandler())
}
