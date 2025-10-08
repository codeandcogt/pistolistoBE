package migration

import (
	"fmt"
	"pistolistoBE/db"

	//"pistolistoBE/internal/modules/rol"
	"pistolistoBE/internal/modules/administrativo"
	"pistolistoBE/internal/modules/permiso"
	"pistolistoBE/internal/modules/rol"
	rolpermiso "pistolistoBE/internal/modules/rolPermiso"
	//"pistolistoBE/internal/modules/administrativo"
	//"pistolistoBE/internal/modules/permiso"
	//"pistolistoBE/internal/modules/rol"
	//"pistolistoBE/internal/modules/rolPermiso"
	// "pistolistoBE/internal/modules/rol"
	//"pistolistoBE/internal/modules/rol"
	//"pistolistoBE/internal/modules/departamento"
	//"pistolistoBE/internal/modules/banco"
	//"pistolistoBE/internal/modules/cupon"
	//"pistolistoBE/internal/modules/moneda"
	// "pistolistoBE/internal/modules/auth"
	// "pistolistoBE/internal/modules/cliente"
	// "pistolistoBE/internal/modules/cliente"
	//"pistolistoBE/internal/modules/subsidiary"
	//"pistolistoBE/internal/modules/bankAccount"
	//"pistolistoBE/internal/modules/municipality"
	//"pistolistoBE/internal/modules/direccion"
	//"pistolistoBE/internal/modules/subCategory"
)

func Migration() {
	database := db.Database()
	// err := database.AutoMigrate(&cliente.Cliente{}, &auth.LogLoginCliente{})
	//err := database.AutoMigrate(&rol.Rol{})

	err := database.AutoMigrate(&permiso.Permiso{}, &rolpermiso.RolPermiso{}, &administrativo.Administrativo{}, &rol.Rol{})
	//err := database.AutoMigrate(&permiso.Permiso{}, &rolpermiso.RolPermiso{}, &administrativo.Administrativo{}, &rol.Rol{})
	//err := database.AutoMigrate(&rol.Rol{})

	//err := database.AutoMigrate(
	//	&subCategory.SubCategory{},
	//	&direccion.Direccion{},
	//)

	// err = database.Exec(`
	// 	ALTER TABLE sub_categories
	// 	ADD CONSTRAINT fk_subcategory_categoria
	// 	FOREIGN KEY (id_categoria)
	// 	REFERENCES categorias(id_categoria)"
	// 	ON UPDATE CASCADE
	// 	ON DELETE RESTRICT
	// `).Error
	// if err != nil {
	// 	fmt.Println("Error se pudo crear el FK")
	// }

	// err = database.Exec(`
	// 	ALTER TABLE direcciones
	// 	ADD CONSTRAINT fk_direccion_municipio
	// 	FOREIGN KEY (id_municipio)
	// 	REFERENCES municipios(id_municipio)
	// 	ON UPDATE CASCADE
	// 	ON DELETE RESTRICT
	// `).Error
	// if err != nil {
	// 	fmt.Println("Error se pudo crear el FK")
	// }

	// err = database.Exec(`
	// 	ALTER TABLE direcciones
	// 	ADD CONSTRAINT fk_direccion_cliente
	// 	FOREIGN KEY (id_cliente)
	// 	REFERENCES clientes(id_cliente)
	// 	ON UPDATE CASCADE
	// 	ON DELETE RESTRICT
	// `).Error
	// if err != nil {
	// 	fmt.Println("Error se pudo crear el FK")
	// }

	// err = database.Exec(`
	// 	ALTER TABLE direcciones
	// 	ADD CONSTRAINT fk_direccion_administrativo
	// 	FOREIGN KEY (id_administrativo)
	// 	REFERENCES administrativos(id_administrativo)
	// 	ON UPDATE CASCADE
	// 	ON DELETE RESTRICT
	// `).Error

	// database.Exec("ALTER TABLE log_login_clientes ADD CONSTRAINT fk_log_login_cliente_cliente FOREIGN KEY (id_cliente) REFERENCES clientes(id_cliente)")
	// // Foreign key para id_rol que referencia a la tabla rol
	// database.Exec("ALTER TABLE rol_permisos ADD CONSTRAINT fk_rol_permiso_rol FOREIGN KEY (id_rol) REFERENCES rols(id_rol)")

	// // Foreign key para id_permiso que referencia a la tabla permiso
	// database.Exec("ALTER TABLE rol_permisos ADD CONSTRAINT fk_rol_permiso_permiso FOREIGN KEY (id_permiso) REFERENCES permisos(id_permiso)")
	// // Foreign key para administrativo que referencia a la tabla rol
	// database.Exec("ALTER TABLE administrativos ADD CONSTRAINT fk_rol_admin FOREIGN KEY (id_rol) REFERENCES rols(id_rol)")
	// // Foreign key para administrativo que referencia a la tabla sucursal
	// database.Exec("ALTER TABLE administrativos ADD CONSTRAINT fk_subsidiaries_admin FOREIGN KEY (id_sucursal) REFERENCES subsidiaries(id_sucursal)")

	if err != nil {
		panic(err)
	} else {
		fmt.Println("se migro")
	}
}
