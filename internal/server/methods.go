package server

import (
	"pistolistoBE/internal/modules/auth"
	"pistolistoBE/internal/modules/categoria"
	"pistolistoBE/internal/modules/cliente"
	"pistolistoBE/internal/modules/departamento"
	"pistolistoBE/internal/modules/descuento"
	"pistolistoBE/internal/modules/rol"
	"pistolistoBE/internal/modules/wishlist"
)

func (h *Handlers) GetClienteHandler() *cliente.ClientHandler {
	return h.Cliente
}

func (h *Handlers) GetAuthHandler() *auth.AuthHandler {
	return h.Auth
}

func (h *Handlers) GetRolHandler() *rol.RolHandler {
	return h.Rol
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

func (h *Handlers) GetWishlistHandler() *wishlist.WishlistHandler {
	return h.Wishlist
}
