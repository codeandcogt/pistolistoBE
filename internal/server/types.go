package server

import (
	"pistolistoBE/internal/modules/auth"
	"pistolistoBE/internal/modules/categoria"
	"pistolistoBE/internal/modules/cliente"
	"pistolistoBE/internal/modules/departamento"
	"pistolistoBE/internal/modules/descuento"
	"pistolistoBE/internal/modules/rol"
	"pistolistoBE/internal/modules/wishlist"

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
	Departamento *departamento.DepartamentoHandler
	Categoria    *categoria.CategoriaHandler
	Descuento    *descuento.DescuentoHandler
	Wishlist     *wishlist.WishlistHandler
}
