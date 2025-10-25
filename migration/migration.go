package migration

import (
	"fmt"
	"pistolistoBE/db"
	"pistolistoBE/internal/modules/loan"
	//"pistolistoBE/internal/modules/articulo"
	//"pistolistoBE/internal/modules/almacen"
	//almacenseccion "pistolistoBE/internal/modules/almacenSeccion"
	//"pistolistoBE/internal/modules/categoria"
)

func Migration() {
	database := db.Database()
	// err := database.AutoMigrate(&cliente.Cliente{}, &auth.LogLoginCliente{})
	err := database.AutoMigrate(&loan.Loan{})
	fmt.Println("Tabla préstamos creada/verificada")

	err = database.Exec(`
	ALTER TABLE loans
	ADD CONSTRAINT fk_prestamo_contrato
	FOREIGN KEY (id_contrato)
	REFERENCES contratos(id_contrato)
	ON UPDATE CASCADE
	ON DELETE RESTRICT;
`).Error
	if err != nil {
		fmt.Println("⚠️ No se pudo crear FK fk_prestamo_contrato:", err)
	}
	if err != nil {
		panic(err)
	} else {
		fmt.Println("se migro")
	}
}
