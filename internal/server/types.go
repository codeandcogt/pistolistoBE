package server

import (
	"pistolistoBE/internal/modules/auth"
	"pistolistoBE/internal/modules/banco"
	"pistolistoBE/internal/modules/cliente"
	"pistolistoBE/internal/modules/moneda"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Server struct {
	Router *mux.Router
	db     *gorm.DB
}

type Handlers struct {
	Cliente *cliente.ClientHandler
	Auth    *auth.AuthHandler
	Moneda  *moneda.MonedaHandler
	Banco   *banco.BancoHandler
}
