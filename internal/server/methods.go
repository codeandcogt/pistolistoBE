package server

import (
	"pistolistoBE/internal/modules/auth"
	"pistolistoBE/internal/modules/banco"
	"pistolistoBE/internal/modules/carrito"
	"pistolistoBE/internal/modules/cliente"
	"pistolistoBE/internal/modules/cupon"
	"pistolistoBE/internal/modules/moneda"
	"pistolistoBE/internal/modules/rol"
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

func (h *Handlers) GetCuponHandler() *cupon.CuponHandler {
	return h.Cupon
}

func (h *Handlers) GetCarritoHandler() *carrito.CarritoHandler {
	return h.Carrito
}
