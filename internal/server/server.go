package server

import (
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
}

// Implementa RouteHandlers interface
func (h *Handlers) GetClienteHandler() *cliente.ClientHandler {
	return h.Cliente
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
	// Cliente module
	clienteRepo := cliente.NewClient(s.db)
	clienteService := cliente.NewClientService(clienteRepo)
	clienteHandler := cliente.NewClientHandler(clienteService)

	return &Handlers{
		Cliente: clienteHandler,
	}
}
