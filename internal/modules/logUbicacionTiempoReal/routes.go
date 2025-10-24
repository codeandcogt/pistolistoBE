package logUbicacionTiempoReal

import "github.com/gorilla/mux"

func SetupLogUbicacionTiempoRealRoutes(api *mux.Router) {
	api.HandleFunc("/ws/ubicaciones", HandleConnections)
	go HandleMessages()
}
