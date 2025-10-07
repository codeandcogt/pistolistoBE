package server

import (
	"pistolistoBE/internal/modules/auth"
	"pistolistoBE/internal/modules/bankAccount"
	"pistolistoBE/internal/modules/cliente"
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
	Cliente     *cliente.ClientHandler
	Auth        *auth.AuthHandler
	Rol         *rol.RolHandler
	Subsidiary  *subsidiary.SubsidiaryHandler
	BankAccount *bankAccount.BankAccountHandler
}
