package server

import (
	"pistolistoBE/internal/modules/administrativo"
	"pistolistoBE/internal/modules/auth"
	"pistolistoBE/internal/modules/banco"
	"pistolistoBE/internal/modules/bankAccount"
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

func (h *Handlers) GetPermisoHandler() *permiso.PermisoHandler {
	return h.Permiso
}

func (h *Handlers) GetRolPermisoHandler() *rolpermiso.RolPermisoHandler {
	return h.RolPermiso
}

func (h *Handlers) GetAdminHandler() *administrativo.AdministrativoHandler {
	return h.Administrativo
}
