package server

import (
	"pistolistoBE/internal/modules/auth"
	"pistolistoBE/internal/modules/bankAccount"
	"pistolistoBE/internal/modules/cliente"
	"pistolistoBE/internal/modules/municipality"
	"pistolistoBE/internal/modules/rol"
	"pistolistoBE/internal/modules/subsidiary"
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

func (h *Handlers) GetSubsidiaryHandler() *subsidiary.SubsidiaryHandler {
	return h.Subsidiary
}

func (h *Handlers) GetBankAccountHandler() *bankAccount.BankAccountHandler {
	return h.BankAccount
}

func (h *Handlers) GetMunicipalityHandler() *municipality.MunicipalityHandler {
	return h.Municipality
}
