package pedido

import (
	"encoding/json"
	"net/http"
	"pistolistoBE/internal/common"
	"strconv"

	"github.com/gorilla/mux"
)

type PedidoHandler struct {
	service PedidoService
}

func NewPedidoHandler(service PedidoService) *PedidoHandler {
	return &PedidoHandler{service}
}

// Checkout (crear pedido completo)
func (h *PedidoHandler) Checkout(w http.ResponseWriter, r *http.Request) {
	var pedido Pedido
	if err := json.NewDecoder(r.Body).Decode(&pedido); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newPedido, err := h.service.Checkout(&pedido)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	common.SuccessResponse(w, common.SUCCESS_CREATED, newPedido, common.HTTP_CREATED)
}

// Obtener todos los pedidos
func (h *PedidoHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	pedidos, err := h.service.GetAll()
	if err != nil {
		common.ErrorResponse(w, http.StatusNotFound, common.HTTP_NOT_FOUND, common.ERR_NOT_FOUND, nil)
		return
	}
	common.SuccessResponse(w, common.SUCCESS_RETRIEVED, pedidos, common.HTTP_OK)
}

// Obtener por ID
func (h *PedidoHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		common.ErrorResponse(w, http.StatusBadRequest, common.HTTP_BAD_REQUEST, common.ERR_VALIDATION, nil)
		return
	}

	pedido, err := h.service.GetByID(uint(id))
	if err != nil {
		common.ErrorResponse(w, http.StatusNotFound, common.HTTP_NOT_FOUND, common.ERR_NOT_FOUND, nil)
		return
	}
	common.SuccessResponse(w, common.SUCCESS_RETRIEVED, pedido, common.HTTP_OK)
}

// Listar pedidos por cliente
func (h *PedidoHandler) GetByCliente(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["idCliente"]
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		common.ErrorResponse(w, http.StatusBadRequest, common.HTTP_BAD_REQUEST, common.ERR_VALIDATION, nil)
		return
	}

	pedidos, err := h.service.GetByCliente(uint(id))
	if err != nil {
		common.ErrorResponse(w, http.StatusNotFound, common.HTTP_NOT_FOUND, common.ERR_NOT_FOUND, nil)
		return
	}
	common.SuccessResponse(w, common.SUCCESS_RETRIEVED, pedidos, common.HTTP_OK)
}

// Cambiar estado
func (h *PedidoHandler) CambiarEstado(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		common.ErrorResponse(w, http.StatusBadRequest, common.HTTP_BAD_REQUEST, common.ERR_VALIDATION, nil)
		return
	}

	var body struct {
		IdEstadoPedido int `json:"id_estado_pedido"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.CambiarEstado(uint(id), body.IdEstadoPedido); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	common.SuccessResponse(w, common.SUCCESS_UPDATED, "Estado actualizado", common.HTTP_OK)
}

// Cancelar pedido
func (h *PedidoHandler) CancelarPedido(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		common.ErrorResponse(w, http.StatusBadRequest, common.HTTP_BAD_REQUEST, common.ERR_VALIDATION, nil)
		return
	}

	var body struct {
		Motivo string `json:"motivo"`
	}
	_ = json.NewDecoder(r.Body).Decode(&body)

	if err := h.service.CancelarPedido(uint(id), body.Motivo); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	common.SuccessResponse(w, common.SUCCESS_UPDATED, "Pedido cancelado", common.HTTP_OK)
}
