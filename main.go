package main

import (
	"log"
	"net/http"
	"os"
	"pistolistoBE/db"
	"pistolistoBE/internal/config"
	"pistolistoBE/internal/server"

	// "pistolistoBE/migration"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	config.Init()

	db := db.Database()

	// migration.Migration()

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
	log.Fatal(server.ListenAndServe())
}
