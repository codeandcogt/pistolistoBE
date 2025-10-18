package wishlist

import (
	"pistolistoBE/internal/middleware"

	"github.com/gorilla/mux"
)

func SetupWishlistRoutes(api *mux.Router, handler *WishlistHandler) {
	wishlistRouter := api.PathPrefix("/wishlist").Subrouter()

	// Ruta pública para crear
	wishlistRouter.HandleFunc("", handler.CreateWishlist).Methods("POST")

	// Rutas protegidas con middleware JWT
	protected := wishlistRouter.NewRoute().Subrouter()
	protected.Use(middleware.JWTMiddleware)

	protected.HandleFunc("/cliente/{id_cliente}", handler.GetAllByCliente).Methods("GET")
	protected.HandleFunc("/{id}", handler.GetWishlistByID).Methods("GET")
	protected.HandleFunc("/{id}", handler.UpdateWishlist).Methods("PUT")
	protected.HandleFunc("/{id}", handler.DeleteWishlist).Methods("DELETE")
}
