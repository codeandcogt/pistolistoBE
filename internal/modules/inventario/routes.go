package inventario

import (
	"pistolistoBE/internal/middleware"

	"github.com/gorilla/mux"
)

func SetupInventarioRoutes(api *mux.Router, handler *InventarioHandler) {
	inventarioRouter := api.PathPrefix("/inventario").Subrouter()

	inventarioRouter.HandleFunc("", handler.CreateInventario).Methods("POST")

	// Rutas protegidas -> subrouter con middleware
	protected := inventarioRouter.NewRoute().Subrouter()
	protected.Use(middleware.JWTMiddleware)

	protected.HandleFunc("/{id}", handler.GetInventarioByID).Methods("GET")
	protected.HandleFunc("/almacen/{id_almacen}", handler.GetAllByAlmacen).Methods("GET")
	protected.HandleFunc("/almacen/{id_almacen}/tipo/{tipo}", handler.GetByTipo).Methods("GET")
	protected.HandleFunc("/almacen/{id_almacen}/estado/{estado}", handler.GetByEstado).Methods("GET")
	protected.HandleFunc("/almacen/{id_almacen}/productos-disponibles", handler.GetProductosDisponibles).Methods("GET")
	protected.HandleFunc("/almacen/{id_almacen}/articulos-empenados", handler.GetArticulosEmpenados).Methods("GET")
	protected.HandleFunc("/{id}", handler.UpdateInventario).Methods("PUT")
	protected.HandleFunc("/{id}", handler.DeleteInventario).Methods("DELETE")
}
