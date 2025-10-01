package server

import (
	"pistolistoBE/internal/modules/auth"
	"pistolistoBE/internal/modules/cliente"
	"pistolistoBE/internal/modules/departamento"
	"pistolistoBE/internal/modules/rol"
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
