package migration

import (
	"fmt"
	"pistolistoBE/db"
	"pistolistoBE/internal/modules/cuota"
	//"pistolistoBE/internal/modules/articulo"
	//"pistolistoBE/internal/modules/almacen"
	//almacenseccion "pistolistoBE/internal/modules/almacenSeccion"
	//"pistolistoBE/internal/modules/categoria"
)

func Migration() {
	database := db.Database()
	// err := database.AutoMigrate(&cliente.Cliente{}, &auth.LogLoginCliente{})
	err := database.AutoMigrate(&cuota.Cuota{})
	fmt.Println("Tabla cuotas creada/verificada")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("se migro")
	}
}
