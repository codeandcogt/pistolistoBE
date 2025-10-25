package routes

import (
	"pistolistoBE/internal/middleware"
	"pistolistoBE/internal/modules/administrativo"
	"pistolistoBE/internal/modules/almacen"
	almacenseccion "pistolistoBE/internal/modules/almacenSeccion"
	"pistolistoBE/internal/modules/auth"
	"pistolistoBE/internal/modules/categoria"
	"pistolistoBE/internal/modules/cliente"
	"pistolistoBE/internal/modules/departamento"
	"pistolistoBE/internal/modules/descuento"
	"pistolistoBE/internal/modules/estadoPedido"
	"pistolistoBE/internal/modules/estadoRuta"
	"pistolistoBE/internal/modules/factura"
	"pistolistoBE/internal/modules/inventario"
	"pistolistoBE/internal/modules/logUbicacion"
	"pistolistoBE/internal/modules/logUbicacionTiempoReal"
	"pistolistoBE/internal/modules/pedido"
	"pistolistoBE/internal/modules/permiso"
	rolpermiso "pistolistoBE/internal/modules/rolPermiso"
	"pistolistoBE/internal/modules/ruta"
	"pistolistoBE/internal/modules/seccion"
	"pistolistoBE/internal/modules/wishlist"
	wishlistitem "pistolistoBE/internal/modules/wishlistItem"

	"pistolistoBE/internal/modules/articulo"
	"pistolistoBE/internal/modules/banco"
	"pistolistoBE/internal/modules/bankAccount"
	"pistolistoBE/internal/modules/carrito"

	"pistolistoBE/internal/modules/cupon"

	"pistolistoBE/internal/modules/moneda"
	"pistolistoBE/internal/modules/municipality"
	"pistolistoBE/internal/modules/pago"
	"pistolistoBE/internal/modules/producto"

	"pistolistoBE/internal/modules/resenaEmpresa"
	"pistolistoBE/internal/modules/rol"

	"pistolistoBE/internal/modules/direccion"

	"pistolistoBE/internal/modules/subCategory"
	"pistolistoBE/internal/modules/subsidiary"

	"pistolistoBE/internal/modules/vehiculo"

	"pistolistoBE/internal/modules/formulario"
	"pistolistoBE/internal/modules/piloto"
	"pistolistoBE/internal/modules/prestamo"

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
	GetWishlistHandler() *wishlist.WishlistHandler
	GetCuponHandler() *cupon.CuponHandler
	GetCarritoHandler() *carrito.CarritoHandler
	GetMunicipalityHandler() *municipality.MunicipalityHandler
	GetBankAccountHandler() *bankAccount.BankAccountHandler
	GetPermisoHandler() *permiso.PermisoHandler
	GetRolPermisoHandler() *rolpermiso.RolPermisoHandler
	GetSubCategoryHandler() *subCategory.SubCategoryHandler
	GetDireccionHandler() *direccion.DireccionHandler
	GetArticuloHandler() *articulo.ArticuloHandler
	GetAdminHandler() *administrativo.AdministrativoHandler
	GetWishListItemHandler() *wishlistitem.WishListItemHandler
	GetAlmacenHandler() *almacen.AlmacenHandler
	GetAlmacenSeccionHandler() *almacenseccion.AlmacenSeccionHandler
	GetSeccionHandler() *seccion.SeccionHandler
	GetInventarioHandler() *inventario.InventarioHandler
	GetResenaEmpresaHandler() *resenaEmpresa.ResenaEmpresaHandler
	GetProductoHandler() *producto.ProductoHandler
	GetVehiculoHandler() *vehiculo.VehiculoHandler
	GetPilotoHandler() *piloto.PilotoHandler
	GetPedidoHandler() *pedido.PedidoHandler
	GetPagoHandler() *pago.PagoHandler
	GetFacturaHandler() *factura.FacturaHandler
	GetEstadoPedidoHandler() *estadoPedido.EstadoPedidoHandler
	GetRutaHandler() *ruta.RutaHandler
	GetEstadoRutaHandler() *estadoRuta.EstadoRutaHandler
	GetLogUbicacionHandler() *logUbicacion.LogUbicacionHandler
	GetFormularioHandler() *formulario.FormularioHandler
	GetPrestamoHandler() *prestamo.PrestamoHandler
}

func SetupRoutes(router *mux.Router, handlers RouteHandlers) {
	router.Use(middleware.Recovery)
	router.Use(middleware.Logger)
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
	wishlist.SetupWishlistRoutes(api, handlers.GetWishlistHandler())
	cupon.SetupCuponRoutes(api, handlers.GetCuponHandler())
	carrito.SetupCarritoRoutes(api, handlers.GetCarritoHandler())
	municipality.SetupMunicipalityRoutes(api, handlers.GetMunicipalityHandler())
	bankAccount.SetupBankAccountRoutes(api, handlers.GetBankAccountHandler())
	subCategory.SetupSubCategoryRoutes(api, handlers.GetSubCategoryHandler())
	direccion.SetupDireccionRoutes(api, handlers.GetDireccionHandler())
	permiso.SetUpPermisoRoutes(api, handlers.GetPermisoHandler())
	rolpermiso.SetUpRolPermisoRoutes(api, handlers.GetRolPermisoHandler())
	articulo.SetupArticuloRoutes(api, handlers.GetArticuloHandler())
	administrativo.SetUpAdminRoutes(api, handlers.GetAdminHandler())
	wishlistitem.SetupWishListItemRoutes(api, handlers.GetWishListItemHandler())
	almacen.SetupAlmacenRoutes(api, handlers.GetAlmacenHandler())
	almacenseccion.SetupAlmacenSeccionRoutes(api, handlers.GetAlmacenSeccionHandler())
	seccion.SetupSeccionRoutes(api, handlers.GetSeccionHandler())
	inventario.SetupInventarioRoutes(api, handlers.GetInventarioHandler())
	resenaEmpresa.SetupResenaEmpresaRoutes(api, handlers.GetResenaEmpresaHandler())
	producto.SetupProductoRoutes(api, handlers.GetProductoHandler())
	vehiculo.SetupVehiculoRoutes(api, handlers.GetVehiculoHandler())
	piloto.SetupPilotoRoutes(api, handlers.GetPilotoHandler())
	pedido.SetupPedidoRoutes(api, handlers.GetPedidoHandler())
	pago.SetupPagoRoutes(api, handlers.GetPagoHandler())
	factura.SetupFacturaRoutes(api, handlers.GetFacturaHandler())
	resenaEmpresa.SetupResenaEmpresaRoutes(api, handlers.GetResenaEmpresaHandler())
	producto.SetupProductoRoutes(api, handlers.GetProductoHandler())
	vehiculo.SetupVehiculoRoutes(api, handlers.GetVehiculoHandler())
	piloto.SetupPilotoRoutes(api, handlers.GetPilotoHandler())
	pedido.SetupPedidoRoutes(api, handlers.GetPedidoHandler())
	pago.SetupPagoRoutes(api, handlers.GetPagoHandler())
	factura.SetupFacturaRoutes(api, handlers.GetFacturaHandler())
	estadoPedido.SetupEstadoPedidoRoutes(api, handlers.GetEstadoPedidoHandler())
	ruta.SetupRutaRoutes(api, handlers.GetRutaHandler())
	estadoRuta.SetupEstadoRutaRoutes(api, handlers.GetEstadoRutaHandler())
	logUbicacion.SetupLogUbicacionRoutes(api, handlers.GetLogUbicacionHandler())
	logUbicacionTiempoReal.SetupLogUbicacionTiempoRealRoutes(api)
	formulario.SetupFormularioRoutes(api, handlers.GetFormularioHandler())
	prestamo.SetupPrestamoRoutes(api, handlers.GetPrestamoHandler())

}
