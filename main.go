package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"pistolistoBE/db"
	"pistolistoBE/internal/server"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	// migration.Migration()
	db := db.Database()

	srv := server.NewServer(db)

	errorVariables := godotenv.Load()
	if errorVariables != nil {
		panic(errorVariables)
	}

	server := &http.Server{
		Addr:         ":" + os.Getenv("PORT"),
		Handler:      srv.Router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}
	fmt.Println("Hello mi ruta es http://192.168.31.57:8080")
	log.Fatal(server.ListenAndServe())
}
