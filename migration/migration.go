package migration

import (
	"fmt"
	"pistolistoBE/db"

	//"pistolistoBE/internal/modules/rol"
	// "pistolistoBE/internal/modules/auth"
	// "pistolistoBE/internal/modules/cliente"
	// "pistolistoBE/internal/modules/cliente"
	"pistolistoBE/internal/modules/subsidiary"
)

func Migration() {
	database := db.Database()
	// err := database.AutoMigrate(&cliente.Cliente{}, &auth.LogLoginCliente{})
	//err := database.AutoMigrate(&rol.Rol{})
	err := database.AutoMigrate(&subsidiary.Subsidiary{})

	// err := database.AutoMigrate()
	// database.Exec("ALTER TABLE log_login_clientes ADD CONSTRAINT fk_log_login_cliente_cliente FOREIGN KEY (id_cliente) REFERENCES clientes(id_cliente)")

	if err != nil {
		panic(err)
	} else {
		fmt.Println("se migro")
	}
}
