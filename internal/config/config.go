package config

import (
	"os"
	"time"
)

var JwtSecret = []byte(os.Getenv("JWT_SECRET"))
var JwtExpiry, _ = time.ParseDuration(os.Getenv("JWT_EXPIRES_IN"))
