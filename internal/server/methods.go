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
	"pistolistoBE/internal/modules/estadoPedido"

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
)

func (h *Handlers) GetClienteHandler() *cliente.ClientHandler {
	return h.Cliente
}

func (h *Handlers) GetAuthHandler() *auth.AuthHandler {
	return h.Auth
}

func (h *Handlers) GetMonedaHandler() *moneda.MonedaHandler {
	return h.Moneda
}

func (h *Handlers) GetBancoHandler() *banco.BancoHandler {
	return h.Banco
}

func (h *Handlers) GetRolHandler() *rol.RolHandler {
	return h.Rol
}

func (h *Handlers) GetSubsidiaryHandler() *subsidiary.SubsidiaryHandler {
	return h.Subsidiary
}

func (h *Handlers) GetBankAccountHandler() *bankAccount.BankAccountHandler {
	return h.BankAccount
}

func (h *Handlers) GetMunicipalityHandler() *municipality.MunicipalityHandler {
	return h.Municipality
}

func (h *Handlers) GetDepartamentoHandler() *departamento.DepartamentoHandler {
	return h.Departamento
}

func (h *Handlers) GetCategoriaHandler() *categoria.CategoriaHandler {
	return h.Categoria
}

func (h *Handlers) GetDescuentoHandler() *descuento.DescuentoHandler {
	return h.Descuento
}

func (h *Handlers) GetCuponHandler() *cupon.CuponHandler {
	return h.Cupon
}

func (h *Handlers) GetCarritoHandler() *carrito.CarritoHandler {
	return h.Carrito
}

func (h *Handlers) GetSubCategoryHandler() *subCategory.SubCategoryHandler {
	return h.SubCategory
}

func (h *Handlers) GetDireccionHandler() *direccion.DireccionHandler {
	return h.Direccion
}

func (h *Handlers) GetPermisoHandler() *permiso.PermisoHandler {
	return h.Permiso
}

func (h *Handlers) GetRolPermisoHandler() *rolpermiso.RolPermisoHandler {
	return h.RolPermiso
}

func (h *Handlers) GetArticuloHandler() *articulo.ArticuloHandler {
	return h.Articulo
}

func (h *Handlers) GetAdminHandler() *administrativo.AdministrativoHandler {
	return h.Administrativo
}

func (h *Handlers) GetResenaEmpresaHandler() *resenaEmpresa.ResenaEmpresaHandler {
	return h.ResenaEmpresa
}

func (h *Handlers) GetProductoHandler() *producto.ProductoHandler {
	return h.Producto
}

func (h *Handlers) GetPedidoHandler() *pedido.PedidoHandler {
	return h.Pedido
}

func (h *Handlers) GetPagoHandler() *pago.PagoHandler {
	return h.Pago
}

func (h *Handlers) GetFacturaHandler() *factura.FacturaHandler {
	return h.Factura
}

func (h *Handlers) GetEstadoPedidoHandler() *estadoPedido.EstadoPedidoHandler {
	return h.EstadoPedido
}
