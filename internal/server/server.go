package server

import (
	"pistolistoBE/internal/config"
	"pistolistoBE/internal/modules/auth"
	"pistolistoBE/internal/modules/cliente"
	"pistolistoBE/internal/routes"

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

// Implementa RouteHandlers interface
func (h *Handlers) GetClienteHandler() *cliente.ClientHandler {
	return h.Cliente
}

func (h *Handlers) GetAuthHandler() *auth.AuthHandler {
	return h.Auth
}

func NewServer(db *gorm.DB) *Server {
	s := &Server{
		Router: mux.NewRouter(),
		db:     db,
	}

	// Inicializar handlers
	handlers := s.initializeHandlers()

	// Configurar rutas usando la interface
	routes.SetupRoutes(s.Router, handlers)

	return s
}

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

	return &Handlers{
		Cliente: clienteHandler,
		Auth:    authHandler,
	}
}
