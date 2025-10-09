package server

import (
	"pistolistoBE/internal/config"
	"pistolistoBE/internal/modules/auth"
	"pistolistoBE/internal/modules/categoria"
	"pistolistoBE/internal/modules/cliente"
	"pistolistoBE/internal/modules/departamento"
	"pistolistoBE/internal/modules/descuento"
	"pistolistoBE/internal/modules/rol"
	"pistolistoBE/internal/modules/wishlist"
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
	// Departamento module
	departamentoRepo := departamento.NewDepartamento(s.db)
	departamentoService := departamento.NewDepartamentoService(departamentoRepo)
	departamentoHandler := departamento.NewDepartamentoHandler(departamentoService)
	//Categoria module
	categoriaRepo := categoria.NewCategoria(s.db)
	categoriaService := categoria.NewCategoriaService(categoriaRepo)
	categoriaHandler := categoria.NewCategoriaHandler(categoriaService)
	//Descuento module
	descuentoRepo := descuento.NewDescuento(s.db)
	descuentoService := descuento.NewDescuentoService(descuentoRepo)
	descuentoHandler := descuento.NewDescuentoHandler(descuentoService)
	//Wishlist module
	wishlistRepo := wishlist.NewWishlist(s.db)
	WishlistService := wishlist.NewWishlistService(wishlistRepo)
	WishlistHandler := wishlist.NewWishlistHandler(WishlistService)

	return &Handlers{
		Cliente:      clienteHandler,
		Auth:         authHandler,
		Rol:          rolHandler,
		Departamento: departamentoHandler,
		Categoria:    categoriaHandler,
		Descuento:    descuentoHandler,
		Wishlist:     WishlistHandler,
	}
}
