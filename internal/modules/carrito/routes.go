package carrito

import (
	"pistolistoBE/internal/middleware"

	"github.com/gorilla/mux"
)

func SetupCarritoRoutes(api *mux.Router, handler *CarritoHandler) {
	carritoRouter := api.PathPrefix("/carritos").Subrouter()

	// Rutas protegidas con middleware
	carritoRouter.Use(middleware.JWTMiddleware)

	// Carrito CRUD
	carritoRouter.HandleFunc("", handler.CreateCarrito).Methods("POST")
	carritoRouter.HandleFunc("", handler.GetAll).Methods("GET")
	carritoRouter.HandleFunc("/{id}", handler.GetCarritoByID).Methods("GET")
	carritoRouter.HandleFunc("/cliente/{clienteId}", handler.GetCarritoByClienteID).Methods("GET")
	carritoRouter.HandleFunc("/{id}", handler.UpdateCarrito).Methods("PUT")
	carritoRouter.HandleFunc("/{id}", handler.DeleteCarrito).Methods("DELETE")

	// CarritoItem operations
	carritoRouter.HandleFunc("/{id}/items", handler.GetCarritoItems).Methods("GET")
	carritoRouter.HandleFunc("/items", handler.AddItem).Methods("POST")
	carritoRouter.HandleFunc("/items/{itemId}", handler.UpdateItem).Methods("PUT")
	carritoRouter.HandleFunc("/items/{itemId}", handler.RemoveItem).Methods("DELETE")
}
