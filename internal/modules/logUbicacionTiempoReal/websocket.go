package logUbicacionTiempoReal

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

var (
	clients   = make(map[*websocket.Conn]bool)
	broadcast = make(chan *LogUbicacionTiempoReal)
	upgrader  = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}
	mutex sync.Mutex
)

// --- Manejo de conexiones WebSocket ---
func HandleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "No se pudo establecer WebSocket", http.StatusInternalServerError)
		return
	}
	defer ws.Close()

	mutex.Lock()
	clients[ws] = true
	mutex.Unlock()

	fmt.Println("📡 Cliente conectado al WebSocket")

	for {
		var msg LogUbicacionTiempoReal
		err := ws.ReadJSON(&msg)
		if err != nil {
			mutex.Lock()
			delete(clients, ws)
			mutex.Unlock()
			fmt.Println("❌ Cliente desconectado del WebSocket")
			break
		}
		// Si quisieras permitir enviar ubicaciones desde WebSocket directamente:
		broadcast <- &msg
	}
}

// --- Envío de mensajes a todos los clientes conectados ---
func HandleMessages() {
	for {
		msg := <-broadcast
		mutex.Lock()
		for client := range clients {
			client.WriteJSON(msg)
		}
		mutex.Unlock()
	}
}

// --- Función para emitir ubicación desde cualquier módulo (como LogUbicacionService) ---
func BroadcastUbicacion(pilotoId int, lat, lon float64) {
	msg := &LogUbicacionTiempoReal{
		PilotoId: pilotoId,
		Latitud:  lat,
		Longitud: lon,
	}
	data, _ := json.Marshal(msg)

	mutex.Lock()
	for client := range clients {
		client.WriteMessage(websocket.TextMessage, data)
	}
	mutex.Unlock()

	fmt.Printf("📤 Broadcast enviado → piloto:%d lat:%.6f lon:%.6f @%s\n", pilotoId, lat, lon, time.Now().Format(time.RFC3339))
}
