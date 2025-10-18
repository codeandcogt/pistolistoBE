package migration

import (
	"fmt"
	"pistolistoBE/db"
	"pistolistoBE/internal/modules/almacen"
	almacenseccion "pistolistoBE/internal/modules/almacenSeccion"
	"pistolistoBE/internal/modules/seccion"
	"pistolistoBE/internal/modules/wishlist"
	wishlistitem "pistolistoBE/internal/modules/wishlistItem"
	//"pistolistoBE/internal/modules/auth"
	// "pistolistoBE/internal/modules/rol"
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
)

func Migration() {
	database := db.Database()
	// err := database.AutoMigrate(&cliente.Cliente{}, &auth.LogLoginCliente{})
	err := database.AutoMigrate(&wishlist.Wishlist{}, &wishlistitem.WishListItem{}, &almacen.Almacen{}, &almacenseccion.AlmacenSeccion{}, &seccion.Seccion{})

	// database.Exec("ALTER TABLE log_login_clientes ADD CONSTRAINT fk_log_login_cliente_cliente FOREIGN KEY (id_cliente) REFERENCES clientes(id_cliente)")

	// // Foreign key para id_rol que referencia a la tabla rol
	//err := database.Exec("ALTER TABLE log_login_admins ADD CONSTRAINT fk_log_sesion_admin FOREIGN KEY (id_administrativo) REFERENCES administrativos(id_administrativo)")

	// // Foreign key para id_permiso que referencia a la tabla permiso
	// database.Exec("ALTER TABLE rol_permisos ADD CONSTRAINT fk_rol_permiso_permiso FOREIGN KEY (id_permiso) REFERENCES permisos(id_permiso)")
	// // Foreign key para administrativo que referencia a la tabla rol
	// database.Exec("ALTER TABLE administrativos ADD CONSTRAINT fk_rol_admin FOREIGN KEY (id_rol) REFERENCES rols(id_rol)")
	// // Foreign key para administrativo que referencia a la tabla sucursal
	// database.Exec("ALTER TABLE administrativos ADD CONSTRAINT fk_subsidiaries_admin FOREIGN KEY (id_sucursal) REFERENCES subsidiaries(id_sucursal)")
	//database.Exec("ALTER TABLE wishlists ADD CONSTRAINT fk_wishlist_cliente FOREIGN KEY (id_wishlist) REFERENCES clientes(id_cliente)")
	//database.Exec("ALTER TABLE wish_list_items ADD CONSTRAINT fk_wishlistItem_wishlist FOREIGN KEY (id_wish_list_item) REFERENCES wishlists(id_wishlist)")
	//database.Exec("ALTER TABLE almacens ADD CONSTRAINT fk_almacen_sucursal FOREIGN KEY (id_almacen) REFERENCES subsidiaries(id_sucursal)")
	//database.Exec("ALTER TABLE almacen_seccions ADD CONSTRAINT fk_almacenSeccion_almacen FOREIGN KEY (id_almacen_seccion) REFERENCES almacens(id_almacen)")
	//database.Exec("ALTER TABLE almacen_seccions ADD CONSTRAINT fk_almacenSeccion_seccion FOREIGN KEY (id_almacen_seccion) REFERENCES seccions(id_seccion)")

	if err != nil {
		panic(err)
	} else {
		fmt.Println("se migro")
	}
}
