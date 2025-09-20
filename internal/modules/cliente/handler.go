package cliente

import (
	"encoding/json"
	"net/http"
)

type ClientHandler struct {
	service ClienteService
}

func NewClientHandler(service ClienteService) *ClientHandler {
	return &ClientHandler{service}
}

func (h *ClientHandler) CreateCliente(w http.ResponseWriter, r *http.Request) {
	var client Cliente
	if err := json.NewDecoder(r.Body).Decode(&client); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.CreateCliente(&client); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(client)
}
