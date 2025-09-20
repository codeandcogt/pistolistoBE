package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"pistolistoBE/db"
	"pistolistoBE/internal/config"
	"pistolistoBE/internal/modules/cliente"
	"time"

	// "github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	// migration.Migration()
	db := db.Database()
	// mux := mux.NewRouter()

	clientRepo := cliente.NewClient(db)
	clientService := cliente.NewClientService(clientRepo)
	clientHandler := cliente.NewClientHandler(clientService)

	r := config.SetupRoutes(clientHandler)

	errorVariables := godotenv.Load()
	if errorVariables != nil {
		panic(errorVariables)
	}

	server := &http.Server{
		Addr:         ":" + os.Getenv("PORT"),
		Handler:      r,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}
	fmt.Println("Hello mi ruta es http://192.168.31.57:8080")
	log.Fatal(server.ListenAndServe())
}
