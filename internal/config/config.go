package config

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
)

var JwtSecret []byte
var JwtExpiry time.Duration

func Init() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("Error cargando el .env: %v\n", err)
		return
	}

	JwtSecret = []byte(os.Getenv("JWT_SECRET"))

	JwtExpiry, err = time.ParseDuration(os.Getenv("JWT_EXPIRES_IN"))
	if err != nil {
		fmt.Printf("Error parseando JWT_EXPIRES_IN: %v\n", err)
		return
	}
}
