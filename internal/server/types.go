package server

import (
	"pistolistoBE/internal/modules/administrativo"
	"pistolistoBE/internal/modules/almacen"
	almacenseccion "pistolistoBE/internal/modules/almacenSeccion"
	"pistolistoBE/internal/modules/articulo"
	"pistolistoBE/internal/modules/auth"
	"pistolistoBE/internal/modules/banco"
	"pistolistoBE/internal/modules/bankAccount"
	"pistolistoBE/internal/modules/categoria"
	"pistolistoBE/internal/modules/cliente"
	"pistolistoBE/internal/modules/cupon"
	"pistolistoBE/internal/modules/departamento"
	"pistolistoBE/internal/modules/descuento"
	"pistolistoBE/internal/modules/direccion"
	"pistolistoBE/internal/modules/moneda"
	"pistolistoBE/internal/modules/municipality"
	"pistolistoBE/internal/modules/permiso"
	"pistolistoBE/internal/modules/producto"
	"pistolistoBE/internal/modules/rol"
	rolpermiso "pistolistoBE/internal/modules/rolPermiso"
	"pistolistoBE/internal/modules/seccion"
	"pistolistoBE/internal/modules/subCategory"
	"pistolistoBE/internal/modules/subsidiary"
	"pistolistoBE/internal/modules/wishlist"
	wishlistitem "pistolistoBE/internal/modules/wishlistItem"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Server struct {
	Router *mux.Router
	db     *gorm.DB
}

type Handlers struct {
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
	Wishlist       *wishlist.WishlistHandler
	WishListItem   *wishlistitem.WishListItemHandler
	Almacen        *almacen.AlmacenHandler
	AlmacenSeccion *almacenseccion.AlmacenSeccionHandler
	Seccion        *seccion.SeccionHandler
	Producto       *producto.ProductoHandler
}
