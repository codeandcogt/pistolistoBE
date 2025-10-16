package middleware

import (
	"log"
	"net/http"
	"strings"
	"time"
)

// CORS middleware compatible con Next.js y redes locales (IP dinámica o localhost)
func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")

		// Lista de orígenes permitidos
		allowedOrigins := map[string]bool{
			"http://localhost:3000":              true, // para desarrollo local
			"http://192.168.31.57:3000":          true,
			"https://pistolisto-web.vercel.app/": true, // cambia por el dominio real de tu frontend
		}

		// Si el origen está permitido, se agrega el header
		if allowedOrigins[origin] {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Vary", "Origin") // Evita problemas de cache
		}

		if origin != "" &&
			(strings.HasPrefix(origin, "http://localhost:") ||
				strings.HasPrefix(origin, "http://169.254.") ||
				strings.HasPrefix(origin, "https://pistolisto-web.vercel.app") ||
				strings.HasPrefix(origin, "https://")) {

			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Vary", "Origin")
		} else {
			log.Println("Origin no permitido:", origin)
		}

		// Cabeceras permitidas
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With, Accept, Origin")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		// Responder al preflight (OPTIONS)
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// Logging middleware
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)

		log.Printf(
			"%s %s %s %v",
			r.RemoteAddr,
			r.Method,
			r.URL.Path,
			time.Since(start),
		)
	})
}

// Content-Type middleware
func ContentTypeJSON(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

// Recovery middleware para panic handling
func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic recovered: %v", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
