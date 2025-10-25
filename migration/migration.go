package migration

import (
	"fmt"
	"pistolistoBE/db"
	"pistolistoBE/internal/modules/contrato"
	//"pistolistoBE/internal/modules/articulo"
	//"pistolistoBE/internal/modules/almacen"
	//almacenseccion "pistolistoBE/internal/modules/almacenSeccion"
	//"pistolistoBE/internal/modules/categoria"
)

func Migration() {
	database := db.Database()
	// err := database.AutoMigrate(&cliente.Cliente{}, &auth.LogLoginCliente{})
	err := database.AutoMigrate(&contrato.Contrato{})
	if err != nil {
		panic(err)
	}
	fmt.Println("✅ Tabla contratos creada/verificada")

	// FK: contrato → avaluo
	err = database.Exec(`
	ALTER TABLE contratos
	ADD CONSTRAINT fk_contrato_avaluo
	FOREIGN KEY (id_avaluo)
	REFERENCES avaluos(id_avaluo)
	ON UPDATE CASCADE
	ON DELETE RESTRICT;
`).Error
	if err != nil {
		fmt.Println("⚠️ No se pudo crear FK fk_contrato_avaluo:", err)
	} else {
		fmt.Println("✅ FK fk_contrato_avaluo creada correctamente")
	}

	if err != nil {
		panic(err)
	} else {
		fmt.Println("se migro")
	}
}
