package migration

import (
	"fmt"
	"pistolistoBE/db"
	//"pistolistoBE/internal/modules/almacen"
	//almacenseccion "pistolistoBE/internal/modules/almacenSeccion"
	//"pistolistoBE/internal/modules/categoria"
	//"pistolistoBE/internal/modules/descuento"
	//"pistolistoBE/internal/modules/inventario"
	//"pistolistoBE/internal/modules/seccion"
	//"pistolistoBE/internal/modules/wishlist"
	//wishlistitem "pistolistoBE/internal/modules/wishlistItem"
	//"pistolistoBE/internal/modules/auth"
	//"pistolistoBE/internal/modules/rol"
	//"pistolistoBE/internal/modules/administrativo"
	//"pistolistoBE/internal/modules/permiso"
	//"pistolistoBE/internal/modules/rol"
	//rolpermiso "pistolistoBE/internal/modules/rolPermiso"
	//"pistolistoBE/internal/modules/administrativo"
	//"pistolistoBE/internal/modules/permiso"
	//"pistolistoBE/internal/modules/rolPermiso"
	// "pistolistoBE/internal/modules/rol"
	//"pistolistoBE/internal/modules/rol"
	//"pistolistoBE/internal/modules/auth"
	// "pistolistoBE/internal/modules/rol"
	//"pistolistoBE/internal/modules/rol"
	//"pistolistoBE/internal/modules/administrativo"
	//"pistolistoBE/internal/modules/permiso"
	//"pistolistoBE/internal/modules/rol"
	//rolpermiso "pistolistoBE/internal/modules/rolPermiso"
	//"pistolistoBE/internal/modules/administrativo"
	//"pistolistoBE/internal/modules/permiso"
	//"pistolistoBE/internal/modules/rol"
	//rolpermiso "pistolistoBE/internal/modules/rolPermiso"
	//"pistolistoBE/internal/modules/administrativo"
	//"pistolistoBE/internal/modules/permiso"
	//"pistolistoBE/internal/modules/rolPermiso"
	// "pistolistoBE/internal/modules/rol"
	//"pistolistoBE/internal/modules/rol"
	//"pistolistoBE/internal/modules/auth"
	// "pistolistoBE/internal/modules/rol"
	//"pistolistoBE/internal/modules/rol"
	//"pistolistoBE/internal/modules/administrativo"
	//"pistolistoBE/internal/modules/permiso"
	//"pistolistoBE/internal/modules/rol"
	//rolpermiso "pistolistoBE/internal/modules/rolPermiso"
	//"pistolistoBE/internal/modules/administrativo"
	//"pistolistoBE/internal/modules/permiso"
	//"pistolistoBE/internal/modules/rol"
	//rolpermiso "pistolistoBE/internal/modules/rolPermiso"
	//"pistolistoBE/internal/modules/administrativo"
	//"pistolistoBE/internal/modules/permiso"
	//"pistolistoBE/internal/modules/rol"
	//"pistolistoBE/internal/modules/rolPermiso"
	// "pistolistoBE/internal/modules/rol"
	//"pistolistoBE/internal/modules/rol"
	//"pistolistoBE/internal/modules/departamento"
	//"pistolistoBE/internal/modules/banco"
	//"pistolistoBE/internal/modules/moneda"
	// "pistolistoBE/internal/modules/auth"
	// "pistolistoBE/internal/modules/cliente"
	// "pistolistoBE/internal/modules/cliente"
	//"pistolistoBE/internal/modules/subsidiary"
	//"pistolistoBE/internal/modules/bankAccount"
	//"pistolistoBE/internal/modules/municipality"
	//"pistolistoBE/internal/modules/direccion"
	//"pistolistoBE/internal/modules/subCategory"
	//"pistolistoBE/internal/modules/articulo"
	//"pistolistoBE/internal/modules/producto"
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

	//err := database.AutoMigrate(&rol.Rol{})

	//err = database.AutoMigrate(&piloto.Piloto{})
	//fmt.Println("Tabla pilotos creada/verificada")

	//err := database.AutoMigrate(&producto.Producto{})
	//err := database.AutoMigrate(
	//	&subCategory.SubCategory{},
	//	&direccion.Direccion{},
	//)

	// err = database.Exec(`
	// 	ALTER TABLE productos
	// 	ADD CONSTRAINT fk_producto_articulo
	// 	FOREIGN KEY (id_articulo)
	// 	REFERENCES articulos(id_articulo)
	// 	ON UPDATE CASCADE
	// 	ON DELETE RESTRICT
	// `).Error
	// if err != nil {
	// 	fmt.Println("No se pudo crear FK fk_producto_articulo:", err)
	// }

	// err = database.Exec(`
	// 	ALTER TABLE productos
	// 	ADD CONSTRAINT fk_producto_descuento
	// 	FOREIGN KEY (id_descuento)
	// 	REFERENCES descuentos(id_descuento)
	// 	ON UPDATE CASCADE
	// 	ON DELETE SET NULL
	// `).Error
	// if err != nil {
	// 	fmt.Println("No se pudo crear FK fk_producto_descuento:", err)
	// }

	// err = database.Exec(`
	// 	ALTER TABLE productos
	// 	ADD CONSTRAINT fk_producto_articulo
	// 	FOREIGN KEY (id_articulo)
	// 	REFERENCES articulos(id_articulo)
	// 	ON UPDATE CASCADE
	// 	ON DELETE RESTRICT
	// `).Error
	// if err != nil {
	// 	fmt.Println("No se pudo crear FK fk_producto_articulo:", err)
	// }

	// err = database.Exec(`
	// 	ALTER TABLE productos
	// 	ADD CONSTRAINT fk_producto_descuento
	// 	FOREIGN KEY (id_descuento)
	// 	REFERENCES descuentos(id_descuento)
	// 	ON UPDATE CASCADE
	// 	ON DELETE SET NULL
	// `).Error
	// if err != nil {
	// 	fmt.Println("No se pudo crear FK fk_producto_descuento:", err)
	// }

	// err = database.Exec(`
	// 	ALTER TABLE productos
	// 	ADD CONSTRAINT fk_producto_articulo
	// 	FOREIGN KEY (id_articulo)
	// 	REFERENCES articulos(id_articulo)
	// 	ON UPDATE CASCADE
	// 	ON DELETE RESTRICT
	// `).Error
	// if err != nil {
	// 	fmt.Println("No se pudo crear FK fk_producto_articulo:", err)
	// }

	// err = database.Exec(`
	// 	ALTER TABLE productos
	// 	ADD CONSTRAINT fk_producto_descuento
	// 	FOREIGN KEY (id_descuento)
	// 	REFERENCES descuentos(id_descuento)
	// 	ON UPDATE CASCADE
	// 	ON DELETE SET NULL
	// `).Error
	// if err != nil {
	// 	fmt.Println("No se pudo crear FK fk_producto_descuento:", err)
	// }

	// err = database.Exec(`
	// 	ALTER TABLE productos
	// 	ADD CONSTRAINT fk_producto_articulo
	// 	FOREIGN KEY (id_articulo)
	// 	REFERENCES articulos(id_articulo)
	// 	ON UPDATE CASCADE
	// 	ON DELETE RESTRICT
	// `).Error
	// if err != nil {
	// 	fmt.Println("No se pudo crear FK fk_producto_articulo:", err)
	// }

	// err = database.Exec(`
	// 	ALTER TABLE productos
	// 	ADD CONSTRAINT fk_producto_descuento
	// 	FOREIGN KEY (id_descuento)
	// 	REFERENCES descuentos(id_descuento)
	// 	ON UPDATE CASCADE
	// 	ON DELETE SET NULL
	// `).Error
	// if err != nil {
	// 	fmt.Println("No se pudo crear FK fk_producto_descuento:", err)
	// }

	// err = database.Exec(`
	// 	ALTER TABLE productos
	// 	ADD CONSTRAINT fk_producto_articulo
	// 	FOREIGN KEY (id_articulo)
	// 	REFERENCES articulos(id_articulo)
	// 	ON UPDATE CASCADE
	// 	ON DELETE RESTRICT
	// `).Error
	// if err != nil {
	// 	fmt.Println("No se pudo crear FK fk_producto_articulo:", err)
	// }

	// err = database.Exec(`
	// 	ALTER TABLE productos
	// 	ADD CONSTRAINT fk_producto_descuento
	// 	FOREIGN KEY (id_descuento)
	// 	REFERENCES descuentos(id_descuento)
	// 	ON UPDATE CASCADE
	// 	ON DELETE SET NULL
	// `).Error
	// if err != nil {
	// 	fmt.Println("No se pudo crear FK fk_producto_descuento:", err)
	// }

	// err = database.Exec(`
	// 	ALTER TABLE productos
	// 	ADD CONSTRAINT fk_producto_articulo
	// 	FOREIGN KEY (id_articulo)
	// 	REFERENCES articulos(id_articulo)
	// 	ON UPDATE CASCADE
	// 	ON DELETE RESTRICT
	// `).Error
	// if err != nil {
	// 	fmt.Println("No se pudo crear FK fk_producto_articulo:", err)
	// }

	// err = database.Exec(`
	// 	ALTER TABLE productos
	// 	ADD CONSTRAINT fk_producto_descuento
	// 	FOREIGN KEY (id_descuento)
	// 	REFERENCES descuentos(id_descuento)
	// 	ON UPDATE CASCADE
	// 	ON DELETE SET NULL
	// `).Error
	// if err != nil {
	// 	fmt.Println("No se pudo crear FK fk_producto_descuento:", err)
	// }

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
	//err := database.AutoMigrate(&permiso.Permiso{}, &rolpermiso.RolPermiso{}, &administrativo.Administrativo{}, &rol.Rol{})
	//err := database.AutoMigrate(&auth.LogLoginAdmin{})
	//err := database.AutoMigrate(&auth.LogLoginAdmin{})
	//err := database.AutoMigrate(&rol.Rol{})
	//err := database.AutoMigrate(&wishlist.Wishlist{}, &wishlistitem.WishListItem{}, &almacen.Almacen{}, &almacenseccion.AlmacenSeccion{}, &seccion.Seccion{})

	// database.Exec("ALTER TABLE log_login_clientes ADD CONSTRAINT fk_log_login_cliente_cliente FOREIGN KEY (id_cliente) REFERENCES clientes(id_cliente)")
	// // Foreign key para id_rol que referencia a la tabla rol
	// database.Exec("ALTER TABLE rol_permisos ADD CONSTRAINT fk_rol_permiso_rol FOREIGN KEY (id_rol) REFERENCES rols(id_rol)")

	// // Foreign key para id_permiso que referencia a la tabla permiso
	// database.Exec("ALTER TABLE rol_permisos ADD CONSTRAINT fk_rol_permiso_permiso FOREIGN KEY (id_permiso) REFERENCES permisos(id_permiso)")
	// // Foreign key para administrativo que referencia a la tabla rol
	// database.Exec("ALTER TABLE administrativos ADD CONSTRAINT fk_rol_admin FOREIGN KEY (id_rol) REFERENCES rols(id_rol)")
	// // Foreign key para administrativo que referencia a la tabla sucursal
	// database.Exec("ALTER TABLE administrativos ADD CONSTRAINT fk_subsidiaries_admin FOREIGN KEY (id_sucursal) REFERENCES subsidiaries(id_sucursal)")

	// // Foreign key para id_rol que referencia a la tabla rol
	//err := database.Exec("ALTER TABLE log_login_admins ADD CONSTRAINT fk_log_sesion_admin FOREIGN KEY (id_administrativo) REFERENCES administrativos(id_administrativo)")
	//err := database.AutoMigrate(&permiso.Permiso{}, &rolpermiso.RolPermiso{}, &administrativo.Administrativo{}, &rol.Rol{})
	//err := database.AutoMigrate(&permiso.Permiso{}, &rolpermiso.RolPermiso{}, &administrativo.Administrativo{}, &rol.Rol{})

	//err := database.AutoMigrate(&articulo.Articulo{})

	//err := database.AutoMigrate(&permiso.Permiso{}, &rolpermiso.RolPermiso{}, &administrativo.Administrativo{}, &rol.Rol{})
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
	//err := database.AutoMigrate(&permiso.Permiso{}, &rolpermiso.RolPermiso{}, &administrativo.Administrativo{}, &rol.Rol{})
	//err := database.AutoMigrate(&auth.LogLoginAdmin{})
	//err := database.AutoMigrate(&auth.LogLoginAdmin{})
	//err := database.AutoMigrate(&rol.Rol{})
	//err := database.AutoMigrate(&auth.LogLoginAdmin{})

	// database.Exec("ALTER TABLE log_login_clientes ADD CONSTRAINT fk_log_login_cliente_cliente FOREIGN KEY (id_cliente) REFERENCES clientes(id_cliente)")
	// // Foreign key para id_rol que referencia a la tabla rol
	// database.Exec("ALTER TABLE rol_permisos ADD CONSTRAINT fk_rol_permiso_rol FOREIGN KEY (id_rol) REFERENCES rols(id_rol)")

	// // Foreign key para id_permiso que referencia a la tabla permiso
	// database.Exec("ALTER TABLE rol_permisos ADD CONSTRAINT fk_rol_permiso_permiso FOREIGN KEY (id_permiso) REFERENCES permisos(id_permiso)")
	// // Foreign key para administrativo que referencia a la tabla rol
	// database.Exec("ALTER TABLE administrativos ADD CONSTRAINT fk_rol_admin FOREIGN KEY (id_rol) REFERENCES rols(id_rol)")
	// // Foreign key para administrativo que referencia a la tabla sucursal
	// database.Exec("ALTER TABLE administrativos ADD CONSTRAINT fk_subsidiaries_admin FOREIGN KEY (id_sucursal) REFERENCES subsidiaries(id_sucursal)")

	// // Foreign key para id_rol que referencia a la tabla rol
	//err := database.Exec("ALTER TABLE log_login_admins ADD CONSTRAINT fk_log_sesion_admin FOREIGN KEY (id_administrativo) REFERENCES administrativos(id_administrativo)")
	//err := database.AutoMigrate(&permiso.Permiso{}, &rolpermiso.RolPermiso{}, &administrativo.Administrativo{}, &rol.Rol{})
	// database.Exec("ALTER TABLE log_login_clientes ADD CONSTRAINT fk_log_login_cliente_cliente FOREIGN KEY (id_cliente) REFERENCES clientes(id_cliente)")
	// // Foreign key para id_rol que referencia a la tabla rol
	// database.Exec("ALTER TABLE rol_permisos ADD CONSTRAINT fk_rol_permiso_rol FOREIGN KEY (id_rol) REFERENCES rols(id_rol)")

	// // Foreign key para id_permiso que referencia a la tabla permiso
	// database.Exec("ALTER TABLE rol_permisos ADD CONSTRAINT fk_rol_permiso_permiso FOREIGN KEY (id_permiso) REFERENCES permisos(id_permiso)")
	// // Foreign key para administrativo que referencia a la tabla rol
	// database.Exec("ALTER TABLE administrativos ADD CONSTRAINT fk_rol_admin FOREIGN KEY (id_rol) REFERENCES rols(id_rol)")
	// // Foreign key para administrativo que referencia a la tabla sucursal
	// database.Exec("ALTER TABLE administrativos ADD CONSTRAINT fk_subsidiaries_admin FOREIGN KEY (id_sucursal) REFERENCES subsidiaries(id_sucursal)")
	//database.Exec("ALTER TABLE wishlists ADD CONSTRAINT fk_wishlist_cliente FOREIGN KEY (id_wishlist) REFERENCES clientes(id_cliente)")
	//database.Exec("ALTER TABLE wish_list_items ADD CONSTRAINT fk_wishlistItem_wishlist FOREIGN KEY (id_wish_list_item) REFERENCES wishlists(id_wishlist)")
	//database.Exec("ALTER TABLE almacens ADD CONSTRAINT fk_almacen_sucursal FOREIGN KEY (id_almacen) REFERENCES subsidiaries(id_sucursal)")
	err := database.Exec(`
    ALTER TABLE almacens
    ADD CONSTRAINT fk_almacen_sucursal
    FOREIGN KEY (id_sucursal)
    REFERENCES subsidiaries(id_sucursal)
    ON UPDATE CASCADE
    ON DELETE RESTRICT;
`).Error

	if err != nil {
		fmt.Println("❌ No se pudo crear FK fk_almacen_sucursal:", err)
	} else {
		fmt.Println("✅ FK fk_almacen_sucursal creada correctamente")
	}
	//database.Exec("ALTER TABLE almacen_seccions ADD CONSTRAINT fk_almacenSeccion_almacen FOREIGN KEY (id_almacen_seccion) REFERENCES almacens(id_almacen)")
	//database.Exec("ALTER TABLE almacen_seccions ADD CONSTRAINT fk_almacenSeccion_seccion FOREIGN KEY (id_almacen_seccion) REFERENCES seccions(id_seccion)")

	// // Foreign key para id_rol que referencia a la tabla rol
	//err := database.Exec("ALTER TABLE log_login_admins ADD CONSTRAINT fk_log_sesion_admin FOREIGN KEY (id_administrativo) REFERENCES administrativos(id_administrativo)")
	//err := database.AutoMigrate(&permiso.Permiso{}, &rolpermiso.RolPermiso{}, &administrativo.Administrativo{}, &rol.Rol{})
	//err := database.AutoMigrate(&permiso.Permiso{}, &rolpermiso.RolPermiso{}, &administrativo.Administrativo{}, &rol.Rol{})

	//err := database.AutoMigrate(&articulo.Articulo{})

	//err := database.AutoMigrate(&permiso.Permiso{}, &rolpermiso.RolPermiso{}, &administrativo.Administrativo{}, &rol.Rol{})
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
	//err := database.AutoMigrate(&permiso.Permiso{}, &rolpermiso.RolPermiso{}, &administrativo.Administrativo{}, &rol.Rol{})
	//err := database.AutoMigrate(&auth.LogLoginAdmin{})
	//err := database.AutoMigrate(&auth.LogLoginAdmin{})
	//err := database.AutoMigrate(&rol.Rol{})

	// database.Exec("ALTER TABLE log_login_clientes ADD CONSTRAINT fk_log_login_cliente_cliente FOREIGN KEY (id_cliente) REFERENCES clientes(id_cliente)")
	// // Foreign key para id_rol que referencia a la tabla rol
	// database.Exec("ALTER TABLE rol_permisos ADD CONSTRAINT fk_rol_permiso_rol FOREIGN KEY (id_rol) REFERENCES rols(id_rol)")

	// // Foreign key para id_permiso que referencia a la tabla permiso
	// database.Exec("ALTER TABLE rol_permisos ADD CONSTRAINT fk_rol_permiso_permiso FOREIGN KEY (id_permiso) REFERENCES permisos(id_permiso)")
	// // Foreign key para administrativo que referencia a la tabla rol
	// database.Exec("ALTER TABLE administrativos ADD CONSTRAINT fk_rol_admin FOREIGN KEY (id_rol) REFERENCES rols(id_rol)")
	// // Foreign key para administrativo que referencia a la tabla sucursal
	// database.Exec("ALTER TABLE administrativos ADD CONSTRAINT fk_subsidiaries_admin FOREIGN KEY (id_sucursal) REFERENCES subsidiaries(id_sucursal)")

	// // Foreign key para id_rol que referencia a la tabla rol
	//err := database.Exec("ALTER TABLE log_login_admins ADD CONSTRAINT fk_log_sesion_admin FOREIGN KEY (id_administrativo) REFERENCES administrativos(id_administrativo)")
	//err := database.AutoMigrate(&permiso.Permiso{}, &rolpermiso.RolPermiso{}, &administrativo.Administrativo{}, &rol.Rol{})
	//err := database.AutoMigrate(&permiso.Permiso{}, &rolpermiso.RolPermiso{}, &administrativo.Administrativo{}, &rol.Rol{})

	//err := database.AutoMigrate(&articulo.Articulo{})

	//err := database.AutoMigrate(&permiso.Permiso{}, &rolpermiso.RolPermiso{}, &administrativo.Administrativo{}, &rol.Rol{})
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
	//err := database.AutoMigrate(&permiso.Permiso{}, &rolpermiso.RolPermiso{}, &administrativo.Administrativo{}, &rol.Rol{})
	//err := database.AutoMigrate(&auth.LogLoginAdmin{})

	// database.Exec("ALTER TABLE log_login_clientes ADD CONSTRAINT fk_log_login_cliente_cliente FOREIGN KEY (id_cliente) REFERENCES clientes(id_cliente)")
	// // Foreign key para id_rol que referencia a la tabla rol
	// database.Exec("ALTER TABLE rol_permisos ADD CONSTRAINT fk_rol_permiso_rol FOREIGN KEY (id_rol) REFERENCES rols(id_rol)")

	// // Foreign key para id_permiso que referencia a la tabla permiso
	// database.Exec("ALTER TABLE rol_permisos ADD CONSTRAINT fk_rol_permiso_permiso FOREIGN KEY (id_permiso) REFERENCES permisos(id_permiso)")
	// // Foreign key para administrativo que referencia a la tabla rol
	// database.Exec("ALTER TABLE administrativos ADD CONSTRAINT fk_rol_admin FOREIGN KEY (id_rol) REFERENCES rols(id_rol)")
	// // Foreign key para administrativo que referencia a la tabla sucursal
	// database.Exec("ALTER TABLE administrativos ADD CONSTRAINT fk_subsidiaries_admin FOREIGN KEY (id_sucursal) REFERENCES subsidiaries(id_sucursal)")
	//database.Exec("ALTER TABLE wishlists ADD CONSTRAINT fk_wishlist_cliente FOREIGN KEY (id_wishlist) REFERENCES clientes(id_cliente)")
	//database.Exec("ALTER TABLE wish_list_items ADD CONSTRAINT fk_wishlistItem_wishlist FOREIGN KEY (id_wish_list_item) REFERENCES wishlists(id_wishlist)")
	//database.Exec("ALTER TABLE almacens ADD CONSTRAINT fk_almac en_sucursal FOREIGN KEY (id_almacen) REFERENCES subsidiaries(id_sucursal)")
	//database.Exec("ALTER TABLE almacen_seccions ADD CONSTRAINT fk_almacenSeccion_almacen FOREIGN KEY (id_almacen_seccion) REFERENCES almacens(id_almacen)")
	//database.Exec("ALTER TABLE almacen_seccions ADD CONSTRAINT fk_almacenSeccion_seccion FOREIGN KEY (id_almacen_seccion) REFERENCES seccions(id_seccion)")
	//database.Exec("ALTER TABLE invetarios ADD CONSTRAINT fk_invetario FOREIGN KEY (id_invetario) REFERENCES almacens(id_almacen)")
	//database.Exec("ALTER TABLE invetarios ADD CONSTRAINT fk_invetario FOREIGN KEY (id_invetario) REFERENCES productos(id_producto)")
	//database.Exec("ALTER TABLE invetarios ADD CONSTRAINT fk_invetario FOREIGN KEY (id_invetario) REFERENCES articulos(id_articulo)")

	if err != nil {
		panic(err)
	} else {
		fmt.Println("se migro")
	}
}
