package middleware

import (
	"log"
	"net/http"
	"time"
)

// CORS middleware
func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")

		// Lista de orígenes permitidos (SIN barra final)
		allowedOrigins := map[string]bool{
			"http://localhost:3000":             true,
			"http://localhost:3001":             true, // Por si usas otro puerto
			"https://pistolisto-web.vercel.app": true, // SIN / al final
		}

		// Si el origen está permitido, configura los headers
		if allowedOrigins[origin] {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Access-Control-Allow-Credentials", "true")
		} else {
			// Si no está permitido, NO envías el header Allow-Origin
			// Esto causará el error CORS en el navegador (que es lo que quieres)
			log.Printf("CORS: Origen no permitido: %s", origin)

			// Para debugging, puedes ver qué origen está llegando
			if origin != "" {
				log.Printf("Origen recibido: '%s'", origin)
			}
		}

		// Estos headers siempre se envían
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization, X-CSRF-Token")
		w.Header().Set("Access-Control-Max-Age", "86400")

		// Maneja las solicitudes preflight
		if r.Method == "OPTIONS" {
			if allowedOrigins[origin] {
				w.WriteHeader(http.StatusNoContent)
			} else {
				w.WriteHeader(http.StatusForbidden)
			}
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
