package server

import (
	"pistolistoBE/internal/modules/administrativo"
	"pistolistoBE/internal/modules/articulo"
	"pistolistoBE/internal/modules/auth"
	"pistolistoBE/internal/modules/banco"
	"pistolistoBE/internal/modules/bankAccount"
	"pistolistoBE/internal/modules/carrito"
	"pistolistoBE/internal/modules/categoria"
	"pistolistoBE/internal/modules/cliente"
	"pistolistoBE/internal/modules/cupon"
	"pistolistoBE/internal/modules/departamento"
	"pistolistoBE/internal/modules/descuento"
	"pistolistoBE/internal/modules/direccion"

	"pistolistoBE/internal/modules/factura"
	"pistolistoBE/internal/modules/moneda"
	"pistolistoBE/internal/modules/municipality"
	"pistolistoBE/internal/modules/pago"
	"pistolistoBE/internal/modules/pedido"
	"pistolistoBE/internal/modules/permiso"
	"pistolistoBE/internal/modules/producto"
	"pistolistoBE/internal/modules/resenaEmpresa"
	"pistolistoBE/internal/modules/rol"
	rolpermiso "pistolistoBE/internal/modules/rolPermiso"
	"pistolistoBE/internal/modules/subCategory"
	"pistolistoBE/internal/modules/subsidiary"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Server struct {
	Router *mux.Router
	db     *gorm.DB
}

type Handlers struct {
	Carrito        *carrito.CarritoHandler
	Cliente        *cliente.ClientHandler
	Auth           *auth.AuthHandler
	Rol            *rol.RolHandler
	Subsidiary     *subsidiary.SubsidiaryHandler
	BankAccount    *bankAccount.BankAccountHandler
	Municipality   *municipality.MunicipalityHandler
	Departamento   *departamento.DepartamentoHandler
	Categoria      *categoria.CategoriaHandler
	Descuento      *descuento.DescuentoHandler
	Moneda         *moneda.MonedaHandler
	Banco          *banco.BancoHandler
	Cupon          *cupon.CuponHandler
	Permiso        *permiso.PermisoHandler
	RolPermiso     *rolpermiso.RolPermisoHandler
	SubCategory    *subCategory.SubCategoryHandler
	Direccion      *direccion.DireccionHandler
	Articulo       *articulo.ArticuloHandler
	Administrativo *administrativo.AdministrativoHandler
	ResenaEmpresa  *resenaEmpresa.ResenaEmpresaHandler
	Producto       *producto.ProductoHandler
	Pedido         *pedido.PedidoHandler
	Pago           *pago.PagoHandler
	Factura        *factura.FacturaHandler
}
