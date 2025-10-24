package server

import (
	"pistolistoBE/internal/config"
	"pistolistoBE/internal/modules/administrativo"
	"pistolistoBE/internal/modules/almacen"
	almacenseccion "pistolistoBE/internal/modules/almacenSeccion"
	"pistolistoBE/internal/modules/articulo"
	"pistolistoBE/internal/modules/auth"
	"pistolistoBE/internal/modules/banco"
	"pistolistoBE/internal/modules/bankAccount"
	"pistolistoBE/internal/modules/carrito"
	"pistolistoBE/internal/modules/categoria"
	"pistolistoBE/internal/modules/cliente"
	"pistolistoBE/internal/modules/cupon"
	"pistolistoBE/internal/modules/departamento"
	"pistolistoBE/internal/modules/descuento"
	"pistolistoBE/internal/modules/direccion"
	"pistolistoBE/internal/modules/factura"
	"pistolistoBE/internal/modules/inventario"
	"pistolistoBE/internal/modules/moneda"
	"pistolistoBE/internal/modules/municipality"
	"pistolistoBE/internal/modules/pago"
	"pistolistoBE/internal/modules/pedido"
	"pistolistoBE/internal/modules/permiso"
	"pistolistoBE/internal/modules/producto"
	"pistolistoBE/internal/modules/resenaEmpresa"
	"pistolistoBE/internal/modules/rol"
	rolpermiso "pistolistoBE/internal/modules/rolPermiso"
	"pistolistoBE/internal/modules/seccion"
	"pistolistoBE/internal/modules/subCategory"
	"pistolistoBE/internal/modules/subsidiary"
	"pistolistoBE/internal/modules/wishlist"
	wishlistitem "pistolistoBE/internal/modules/wishlistItem"
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

	// Rol module
	rolRepo := rol.NewRol(s.db)
	rolService := rol.NewRolService(rolRepo)
	rolHandler := rol.NewRolHandler(rolService)

	// Subsidiary module
	subsidiaryRepo := subsidiary.NewSubsidiaryRepository(s.db)
	subsidiaryService := subsidiary.NewSubsidiaryService(subsidiaryRepo)
	subsidiaryHandler := subsidiary.NewSubsidiaryHandler(subsidiaryService)

	// BankAccount module
	bankAccountRepo := bankAccount.NewBankAccountRepository(s.db)
	bankAccountService := bankAccount.NewBankAccountService(bankAccountRepo)
	bankAccountHandler := bankAccount.NewBankAccountHandler(bankAccountService)

	// Municipality module
	municipalityRepo := municipality.NewMunicipalityRepository(s.db)
	municipalityService := municipality.NewMunicipalityService(municipalityRepo)
	municipalityHandler := municipality.NewMunicipalityHandler(municipalityService)

	// Departamento module
	departamentoRepo := departamento.NewDepartamento(s.db)
	departamentoService := departamento.NewDepartamentoService(departamentoRepo)
	departamentoHandler := departamento.NewDepartamentoHandler(departamentoService)

	// Categoria module
	categoriaRepo := categoria.NewCategoria(s.db)
	categoriaService := categoria.NewCategoriaService(categoriaRepo)
	categoriaHandler := categoria.NewCategoriaHandler(categoriaService)

	// Descuento module
	descuentoRepo := descuento.NewDescuento(s.db)
	descuentoService := descuento.NewDescuentoService(descuentoRepo)
	descuentoHandler := descuento.NewDescuentoHandler(descuentoService)
	//Wishlist module
	wishlistRepo := wishlist.NewWishlist(s.db)
	WishlistService := wishlist.NewWishlistService(wishlistRepo)
	WishlistHandler := wishlist.NewWishlistHandler(WishlistService)

	// Moneda module
	monedaRepo := moneda.NewMonedaRepository(s.db)
	monedaService := moneda.NewMonedaService(monedaRepo)
	monedaHandler := moneda.NewMonedaHandler(monedaService)

	// Banco module
	bancoRepo := banco.NewBancoRepository(s.db)
	bancoService := banco.NewBancoService(bancoRepo)
	bancoHandler := banco.NewBancoHandler(bancoService)

	// Cupon module
	cuponRepo := cupon.NewCuponRepository(s.db)
	cuponService := cupon.NewCuponService(cuponRepo)
	cuponHandler := cupon.NewCuponHandler(cuponService)

	// Carrito module
	carritoRepo := carrito.NewCarritoRepository(s.db)
	carritoService := carrito.NewCarritoService(carritoRepo)
	carritoHandler := carrito.NewCarritoHandler(carritoService)

	// SubCategory module
	subCategoryRepo := subCategory.NewSubCategoryRepository(s.db)
	subCategoryService := subCategory.NewSubCategoryService(subCategoryRepo)
	subCategoryHandler := subCategory.NewSubCategoryHandler(subCategoryService)

	// Direccion module
	direccionRepo := direccion.NewDireccionRepository(s.db)
	direccionService := direccion.NewDireccionService(direccionRepo)
	direccionHandler := direccion.NewDireccionHandler(direccionService)
	//Permiso module
	permisoRepo := permiso.NewPermisoRepository(s.db)
	permisoService := permiso.NewPermisoService(permisoRepo)
	permisoHandler := permiso.NewPermisoHandler(permisoService)

	// Rol Permiso module
	rolPermisoRepo := rolpermiso.NewRolPermiso(s.db)
	rolPermisoService := rolpermiso.NewRolPermosoService(rolPermisoRepo)
	rolPermisoHandler := rolpermiso.NewRolPermisoHandler(rolPermisoService)

	// Articulo module
	articuloRepo := articulo.NewArticuloRepository(s.db)
	articuloService := articulo.NewArticuloService(articuloRepo)
	articuloHandler := articulo.NewArticuloHandler(articuloService)

	// Administrativo module
	adminRepo := administrativo.NewAdminRepository(s.db)
	adminService := administrativo.NewAdministrativoService(adminRepo)
	adminHandler := administrativo.NewAdministrativoHandler(adminService)

	//WishListItem Modulo
	wishListItemRepo := wishlistitem.NewWishListItemRepository(s.db)
	wishListItemService := wishlistitem.NewWishListItemService(wishListItemRepo)
	wishListItemHandler := wishlistitem.NewWishListItemHandler(wishListItemService)

	//Almacen
	almacenRepo := almacen.NewAlmacenRepository(s.db)
	almacenService := almacen.NewAlmacenService(almacenRepo)
	almacenHandler := almacen.NewAlmacenHandler(almacenService)

	//AlmacenSeccion
	almacenSeccionRepo := almacenseccion.NewAlmacenSeccionRepository(s.db)
	almacenSeccionService := almacenseccion.NewAlmacenSeccionService(almacenSeccionRepo)
	almacenSeccionHandler := almacenseccion.NewAlmacenSeccionHandler(almacenSeccionService)

	//Seccion
	seccionRepo := seccion.NewSeccionRepository(s.db)
	seccionService := seccion.NewSeccionService(seccionRepo)
	seccionHandler := seccion.NewSeccionHandler(seccionService)
	// ResenaEmpresa module
	resenaEmpresaRepo := resenaEmpresa.NewResenaEmpresaRepository(s.db)
	resenaEmpresaService := resenaEmpresa.NewResenaEmpresaService(resenaEmpresaRepo)
	resenaEmpresaHandler := resenaEmpresa.NewResenaEmpresaHandler(resenaEmpresaService)

	// Producto module
	productoRepo := producto.NewProductoRepository(s.db)
	productoService := producto.NewProductoService(productoRepo)
	productoHandler := producto.NewProductoHandler(productoService)

	// Pedido module
	pedidoRepo := pedido.NewPedidoRepository(s.db)
	pedidoService := pedido.NewPedidoService(pedidoRepo)
	pedidoHandler := pedido.NewPedidoHandler(pedidoService)

	// Pago module
	pagoRepo := pago.NewPagoRepository(s.db)
	pagoService := pago.NewPagoService(pagoRepo, pedidoRepo)
	pagoHandler := pago.NewPagoHandler(pagoService)

	// Factura module
	facturaRepo := factura.NewFacturaRepository(s.db)
	facturaService := factura.NewFacturaService(facturaRepo)
	facturaHandler := factura.NewFacturaHandler(facturaService)

	//inventario
	inventarioRepo := inventario.NewInventarioRepository(s.db)
	inventarioService := inventario.NewInventarioService(inventarioRepo)
	inventarioHandler := inventario.NewInventarioHandler(inventarioService)

	return &Handlers{
		Cliente:        clienteHandler,
		Auth:           authHandler,
		Rol:            rolHandler,
		Subsidiary:     subsidiaryHandler,
		BankAccount:    bankAccountHandler,
		Municipality:   municipalityHandler,
		Departamento:   departamentoHandler,
		Categoria:      categoriaHandler,
		Descuento:      descuentoHandler,
		Moneda:         monedaHandler,
		Banco:          bancoHandler,
		Cupon:          cuponHandler,
		Carrito:        carritoHandler,
		SubCategory:    subCategoryHandler,
		Direccion:      direccionHandler,
		Permiso:        permisoHandler,
		RolPermiso:     rolPermisoHandler,
		Articulo:       articuloHandler,
		Administrativo: adminHandler,
		Wishlist:       WishlistHandler,
		WishListItem:   wishListItemHandler,
		Almacen:        almacenHandler,
		AlmacenSeccion: almacenSeccionHandler,
		Seccion:        seccionHandler,
		Inventario:     inventarioHandler,
		ResenaEmpresa:  resenaEmpresaHandler,
		Producto:       productoHandler,
		Pedido:         pedidoHandler,
		Pago:           pagoHandler,
		Factura:        facturaHandler,
	}

}
