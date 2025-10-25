package migration

import (
	"fmt"
	"pistolistoBE/db"
	"pistolistoBE/internal/modules/administrativo"
	"pistolistoBE/internal/modules/articulo"

	//"pistolistoBE/internal/modules/articulo"
	"pistolistoBE/internal/modules/banco"
	"pistolistoBE/internal/modules/carrito"
	"pistolistoBE/internal/modules/cupon"
	"pistolistoBE/internal/modules/formulario"
	"pistolistoBE/internal/modules/moneda"
	"pistolistoBE/internal/modules/permiso"
	"pistolistoBE/internal/modules/piloto"
	"pistolistoBE/internal/modules/producto"
	"pistolistoBE/internal/modules/resenaEmpresa"
	"pistolistoBE/internal/modules/rol"
	rolpermiso "pistolistoBE/internal/modules/rolPermiso"
	"pistolistoBE/internal/modules/vehiculo"
	//"pistolistoBE/internal/modules/almacen"
	//almacenseccion "pistolistoBE/internal/modules/almacenSeccion"
	//"pistolistoBE/internal/modules/categoria"
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
		&producto.Producto{},
		&articulo.Articulo{},
		&vehiculo.Vehiculo{},
		&piloto.Piloto{},
		&formulario.Formulario{},
	)

	// AVALUO → USUARIO (ADMINISTRATIVO)
	err = database.Exec(`
		ALTER TABLE avaluos
		ADD CONSTRAINT fk_avaluo_usuario
		FOREIGN KEY (id_usuario)
		REFERENCES administrativos(id_administrativo)
		ON UPDATE CASCADE
		ON DELETE RESTRICT;
	`).Error
	if err != nil {
		fmt.Println("⚠️ No se pudo crear FK fk_avaluo_usuario:", err)
	} else {
		fmt.Println("✅ FK fk_avaluo_usuario creada correctamente")
	}

	if err != nil {
		panic(err)
	} else {
		fmt.Println("se migro")
	}
}
