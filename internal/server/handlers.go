package server

import (
	"pistolistoBE/internal/config"
	"pistolistoBE/internal/modules/auth"
	"pistolistoBE/internal/modules/banco"
	"pistolistoBE/internal/modules/cliente"
	"pistolistoBE/internal/modules/moneda"
	"pistolistoBE/internal/modules/rol"
)

func (s *Server) initializeHandlers() *Handlers {
	// Auth module
	authRepo := auth.NewAuthRepository(s.db)
	jwtManager := auth.NewJwtManager(config.JwtSecret, config.JwtExpiry)
	authService := auth.NewAuthService(authRepo, jwtManager)
	authHandler := auth.NewAuthHandler(authService)
	// Cliente module
	clienteRepo := cliente.NewClient(s.db)
	clienteService := cliente.NewClientService(clienteRepo)
	clienteHandler := cliente.NewClientHandler(clienteService)
	// Cliente module
	rolRepo := rol.NewRol(s.db)
	rolService := rol.NewRolService(rolRepo)
	rolHandler := rol.NewRolHandler(rolService)

	// Moneda module
	monedaRepo := moneda.NewMonedaRepository(s.db)
	monedaService := moneda.NewMonedaService(monedaRepo)
	monedaHandler := moneda.NewMonedaHandler(monedaService)

	// Banco module
	bancoRepo := banco.NewBancoRepository(s.db)
	bancoService := banco.NewBancoService(bancoRepo)
	bancoHandler := banco.NewBancoHandler(bancoService)

	return &Handlers{
		Cliente: clienteHandler,
		Auth:    authHandler,
		Moneda:  monedaHandler,
		Banco:   bancoHandler,
		Rol:     rolHandler,
	}
}
