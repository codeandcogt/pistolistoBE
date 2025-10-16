package migration

import (
	"fmt"
	"pistolistoBE/db"
	"pistolistoBE/internal/modules/administrativo"
	"pistolistoBE/internal/modules/auth"
	"pistolistoBE/internal/modules/banco"
	"pistolistoBE/internal/modules/carrito"
	"pistolistoBE/internal/modules/cupon"
	"pistolistoBE/internal/modules/moneda"
	"pistolistoBE/internal/modules/permiso"
	"pistolistoBE/internal/modules/resenaEmpresa"
	"pistolistoBE/internal/modules/rol"
	rolpermiso "pistolistoBE/internal/modules/rolPermiso"
	//"pistolistoBE/internal/modules/almacen"
	//almacenseccion "pistolistoBE/internal/modules/almacenSeccion"
	//"pistolistoBE/internal/modules/categoria"
	//"pistolistoBE/internal/modules/descuento"
	//"pistolistoBE/internal/modules/inventario"
	//"pistolistoBE/internal/modules/seccion"
	//"pistolistoBE/internal/modules/wishlist"
	//wishlistitem "pistolistoBE/internal/modules/wishlistItem"
	//"pistolistoBE/internal/modules/auth"
)

func Migration() {
	database := db.Database()
	// err := database.AutoMigrate(&cliente.Cliente{}, &auth.LogLoginCliente{})
	err := database.AutoMigrate(&permiso.Permiso{}, &rolpermiso.RolPermiso{}, &administrativo.Administrativo{}, &rol.Rol{},
		&moneda.Moneda{},
		&banco.Banco{},
		&cupon.Cupon{},
		&carrito.Carrito{},
		&carrito.CarritoItem{},
		&resenaEmpresa.ResenaEmpresa{},
	)
	err = database.AutoMigrate(&auth.LogLoginAdmin{})
	//err := database.AutoMigrate(&rol.Rol{})

	if err != nil {
		panic(err)
	} else {
		fmt.Println("se migro")
	}
}
