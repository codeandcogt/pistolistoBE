package migration

import (
	"fmt"
	"pistolistoBE/db"

	//"pistolistoBE/internal/modules/rol"
	// "pistolistoBE/internal/modules/auth"
	// "pistolistoBE/internal/modules/cliente"
	// "pistolistoBE/internal/modules/cliente"
	//"pistolistoBE/internal/modules/subsidiary"
	"pistolistoBE/internal/modules/bankAccount"
	"pistolistoBE/internal/modules/municipality"
)

func Migration() {
	database := db.Database()
	// err := database.AutoMigrate(&cliente.Cliente{}, &auth.LogLoginCliente{})
	//err := database.AutoMigrate(&rol.Rol{})
	err := database.AutoMigrate(
		&municipality.Municipality{},
		&bankAccount.BankAccount{},
	)
	// err := database.AutoMigrate()

	// --- Municipality → Departamento
	// err = database.Exec(`
	// 	ALTER TABLE municipalities
	// 	ADD CONSTRAINT fk_municipality_departamento
	// 	FOREIGN KEY (id_departamento)
	// 	REFERENCES departamentos(id_departamento)
	// 	ON UPDATE CASCADE
	// 	ON DELETE RESTRICT;
	// `).Error
	// if err != nil {
	// 	fmt.Println("⚠️ Advertencia: No se pudo crear FK municipality → departamento:", err)
	// }

	// // --- BankAccount → Cliente
	// err = database.Exec(`
	// 	ALTER TABLE bank_accounts
	// 	ADD CONSTRAINT fk_bank_account_cliente
	// 	FOREIGN KEY (id_cliente)
	// 	REFERENCES clientes(id_cliente)
	// 	ON UPDATE CASCADE
	// 	ON DELETE RESTRICT;
	// `).Error
	// if err != nil {
	// 	fmt.Println("⚠️ Advertencia: No se pudo crear FK bank_account → cliente:", err)
	// }

	// // --- BankAccount → Moneda
	// err = database.Exec(`
	// 	ALTER TABLE bank_accounts
	// 	ADD CONSTRAINT fk_bank_account_moneda
	// 	FOREIGN KEY (id_moneda)
	// 	REFERENCES monedas(id_moneda)
	// 	ON UPDATE CASCADE
	// 	ON DELETE RESTRICT;
	// `).Error
	// if err != nil {
	// 	fmt.Println("⚠️ Advertencia: No se pudo crear FK bank_account → moneda:", err)
	// }

	// // --- BankAccount → Banco
	// err = database.Exec(`
	// 	ALTER TABLE bank_accounts
	// 	ADD CONSTRAINT fk_bank_account_banco
	// 	FOREIGN KEY (id_banco)
	// 	REFERENCES bancos(id_banco)
	// 	ON UPDATE CASCADE
	// 	ON DELETE RESTRICT;
	// `).Error
	// if err != nil {
	// 	fmt.Println("⚠️ Advertencia: No se pudo crear FK bank_account → banco:", err)
	// }

	// database.Exec("ALTER TABLE log_login_clientes ADD CONSTRAINT fk_log_login_cliente_cliente FOREIGN KEY (id_cliente) REFERENCES clientes(id_cliente)")

	if err != nil {
		panic(err)
	} else {
		fmt.Println("se migro")
	}
}
