package wishlistitem

import (
	"pistolistoBE/internal/middleware"

	"github.com/gorilla/mux"
)

func SetupWishListItemRoutes(api *mux.Router, handler *WishListItemHandler) {
	// Subrouter para /wishListItem
	wishListItemRouter := api.PathPrefix("/wishListItem").Subrouter()

	// Ruta pública para crear un ítem
	wishListItemRouter.HandleFunc("", handler.AddWishListItem).Methods("POST")

	// Subrouter protegido con middleware JWT
	protected := wishListItemRouter.NewRoute().Subrouter()
	protected.Use(middleware.JWTMiddleware)

	// Obtener todos los ítems de una wishlist por cliente
	protected.HandleFunc("/wishlist/{wishlist_id}", handler.GetWishListItemByWishlist).Methods("GET")

	// Obtener un ítem por ID
	protected.HandleFunc("/{id}", handler.GetWishListItemByID).Methods("GET")

	// Actualizar un ítem
	protected.HandleFunc("/{id}", handler.UpdateWishListItem).Methods("PUT")

	// Eliminar (desactivar) un ítem
	protected.HandleFunc("/{id}", handler.DeleteWishListItem).Methods("DELETE")
}
