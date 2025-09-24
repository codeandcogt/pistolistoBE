package cliente

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pistolistoBE/internal/common"
	"strconv"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
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
	fmt.Println("cliente", client)
	hash, _ := bcrypt.GenerateFromPassword([]byte(client.Contrasena), bcrypt.DefaultCost)
	fmt.Println("cliente hash", hash)
	client.Contrasena = string(hash)

	if err := h.service.CreateCliente(&client); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	common.SuccessResponse(w, common.SUCCESS_CREATED, client, common.HTTP_CREATED)

}

func (h *ClientHandler) GetClienteByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr, exists := vars["id"]
	if !exists {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	client, err := h.service.GetClienteByID(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	common.SuccessResponse(w, common.SUCCESS_RETRIEVED, client, common.HTTP_OK)
}
