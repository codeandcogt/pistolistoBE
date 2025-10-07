package server

import (
	"pistolistoBE/internal/modules/auth"
	"pistolistoBE/internal/modules/bankAccount"
	"pistolistoBE/internal/modules/cliente"
	"pistolistoBE/internal/modules/municipality"
	"pistolistoBE/internal/modules/banco"
	"pistolistoBE/internal/modules/categoria"
	"pistolistoBE/internal/modules/cliente"
	"pistolistoBE/internal/modules/cupon"
	"pistolistoBE/internal/modules/departamento"
	"pistolistoBE/internal/modules/descuento"
	"pistolistoBE/internal/modules/moneda"
	"pistolistoBE/internal/modules/rol"
	"pistolistoBE/internal/modules/subsidiary"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Server struct {
	Router *mux.Router
	db     *gorm.DB
}

type Handlers struct {
	Cliente      *cliente.ClientHandler
	Auth         *auth.AuthHandler
	Rol          *rol.RolHandler
	Subsidiary   *subsidiary.SubsidiaryHandler
	BankAccount  *bankAccount.BankAccountHandler
	Municipality *municipality.MunicipalityHandler
	Departamento *departamento.DepartamentoHandler
	Categoria    *categoria.CategoriaHandler
	Descuento    *descuento.DescuentoHandler
	Moneda       *moneda.MonedaHandler
	Banco        *banco.BancoHandler
	Cupon        *cupon.CuponHandler
}
