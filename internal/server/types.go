package server

import (
	"pistolistoBE/internal/modules/auth"
	"pistolistoBE/internal/modules/cliente"

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
}
