package migration

import (
	"fmt"
	"pistolistoBE/db"
	"pistolistoBE/internal/modules/cliente"
)

func Migration() {
	database := db.Database()
	err := database.AutoMigrate(&cliente.Cliente{})
	if err != nil {
		panic(err)
	} else {
		fmt.Println("se migro")
	}
}
