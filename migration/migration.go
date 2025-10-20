package migration

import (
	"fmt"
	"pistolistoBE/db"
	"pistolistoBE/internal/modules/estadoPedido"
)

func Migration() {
	database := db.Database()
	// err := database.AutoMigrate(&cliente.Cliente{}, &auth.LogLoginCliente{})
	// err := database.AutoMigrate(&permiso.Permiso{}, &rolpermiso.RolPermiso{}, &administrativo.Administrativo{}, &rol.Rol{},
	// 	&moneda.Moneda{},
	// 	&banco.Banco{},
	// 	&cupon.Cupon{},
	// 	&carrito.Carrito{},
	// 	&carrito.CarritoItem{},
	// 	&resenaEmpresa.ResenaEmpresa{},
	// )
	err := database.AutoMigrate(&estadoPedido.EstadoPedido{})

	//err := database.AutoMigrate(&rol.Rol{})

	if err != nil {
		panic(err)
	} else {
		fmt.Println("se migro")
	}
}
