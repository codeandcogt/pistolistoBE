package server

import (
	"pistolistoBE/internal/config"
	"pistolistoBE/internal/modules/auth"
	"pistolistoBE/internal/modules/banco"
	"pistolistoBE/internal/modules/bankAccount"
	"pistolistoBE/internal/modules/categoria"
	"pistolistoBE/internal/modules/cliente"
	"pistolistoBE/internal/modules/cupon"
	"pistolistoBE/internal/modules/departamento"
	"pistolistoBE/internal/modules/descuento"
	"pistolistoBE/internal/modules/direccion"
	"pistolistoBE/internal/modules/moneda"
	"pistolistoBE/internal/modules/municipality"
	"pistolistoBE/internal/modules/permiso"
	"pistolistoBE/internal/modules/rol"
	rolpermiso "pistolistoBE/internal/modules/rolPermiso"
	"pistolistoBE/internal/modules/subCategory"
	"pistolistoBE/internal/modules/subsidiary"
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
	//Categoria module
	categoriaRepo := categoria.NewCategoria(s.db)
	categoriaService := categoria.NewCategoriaService(categoriaRepo)
	categoriaHandler := categoria.NewCategoriaHandler(categoriaService)
	//Descuento module
	descuentoRepo := descuento.NewDescuento(s.db)
	descuentoService := descuento.NewDescuentoService(descuentoRepo)
	descuentoHandler := descuento.NewDescuentoHandler(descuentoService)

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

	//Rol Permiso modulo
	rolPermisoRepo := rolpermiso.NewRolPermiso(s.db)
	rolPermisoService := rolpermiso.NewRolPermosoService(rolPermisoRepo)
	rolPermisoHandler := rolpermiso.NewRolPermisoHandler(rolPermisoService)

	// SubCategory module
	subCategoryRepo := subCategory.NewSubCategoryRepository(s.db)
	subCategoryService := subCategory.NewSubCategoryService(subCategoryRepo)
	subCategoryHandler := subCategory.NewSubCategoryHandler(subCategoryService)

	// Direccion module
	direccionRepo := direccion.NewDireccionRepository(s.db)
	direccionService := direccion.NewDireccionService(direccionRepo)
	direccionHandler := direccion.NewDireccionHandler(direccionService)

	return &Handlers{
		Cliente:      clienteHandler,
		Auth:         authHandler,
		Rol:          rolHandler,
		Subsidiary:   subsidiaryHandler,
		BankAccount:  bankAccountHandler,
		Municipality: municipalityHandler,
		Departamento: departamentoHandler,
		Categoria:    categoriaHandler,
		Descuento:    descuentoHandler,
		Moneda:       monedaHandler,
		Banco:        bancoHandler,
		Cupon:        cuponHandler,
		SubCategory:  subCategoryHandler,
		Direccion:    direccionHandler,
		Permiso:      permisoHandler,
		RolPermiso:   rolPermisoHandler,
		SubCategory:  subCategoryHandler,
		Direccion:    direccionHandler,
	}
}
