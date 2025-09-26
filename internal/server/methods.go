package server

import (
	"pistolistoBE/internal/modules/auth"
	"pistolistoBE/internal/modules/cliente"
)

func (h *Handlers) GetClienteHandler() *cliente.ClientHandler {
	return h.Cliente
}

func (h *Handlers) GetAuthHandler() *auth.AuthHandler {
	return h.Auth
}
