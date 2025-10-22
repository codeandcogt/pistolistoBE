package factura

import (
	"encoding/json"
	"net/http"
	"pistolistoBE/internal/common"
	"strconv"

	"github.com/gorilla/mux"
)

type FacturaHandler struct {
	service FacturaService
}

func NewFacturaHandler(service FacturaService) *FacturaHandler {
	return &FacturaHandler{service}
}

// Crear factura (emitir)
func (h *FacturaHandler) EmitirFactura(w http.ResponseWriter, r *http.Request) {
	var factura Factura
	if err := json.NewDecoder(r.Body).Decode(&factura); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newFactura, err := h.service.EmitirFactura(&factura)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	common.SuccessResponse(w, "Factura emitida correctamente", newFactura, common.HTTP_CREATED)
}

// Obtener todas
func (h *FacturaHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	facturas, err := h.service.GetAll()
	if err != nil {
		common.ErrorResponse(w, http.StatusNotFound, common.HTTP_NOT_FOUND, common.ERR_NOT_FOUND, nil)
		return
	}
	common.SuccessResponse(w, common.SUCCESS_RETRIEVED, facturas, common.HTTP_OK)
}

// Obtener por ID
func (h *FacturaHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		common.ErrorResponse(w, http.StatusBadRequest, common.HTTP_BAD_REQUEST, common.ERR_VALIDATION, nil)
		return
	}

	factura, err := h.service.GetByID(uint(id))
	if err != nil {
		common.ErrorResponse(w, http.StatusNotFound, common.HTTP_NOT_FOUND, common.ERR_NOT_FOUND, nil)
		return
	}
	common.SuccessResponse(w, common.SUCCESS_RETRIEVED, factura, common.HTTP_OK)
}

// Obtener por pedido
func (h *FacturaHandler) GetByPedido(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["pedidoId"]
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		common.ErrorResponse(w, http.StatusBadRequest, common.HTTP_BAD_REQUEST, common.ERR_VALIDATION, nil)
		return
	}

	factura, err := h.service.GetByPedido(uint(id))
	if err != nil {
		common.ErrorResponse(w, http.StatusNotFound, common.HTTP_NOT_FOUND, common.ERR_NOT_FOUND, nil)
		return
	}
	common.SuccessResponse(w, common.SUCCESS_RETRIEVED, factura, common.HTTP_OK)
}

// Update
func (h *FacturaHandler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		common.ErrorResponse(w, http.StatusBadRequest, common.HTTP_BAD_REQUEST, common.ERR_VALIDATION, nil)
		return
	}

	var factura Factura
	if err := json.NewDecoder(r.Body).Decode(&factura); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	factura.IdFactura = uint(id)

	msg, err := h.service.Update(&factura)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	common.SuccessResponse(w, msg, factura, common.HTTP_OK)
}

func (h *FacturaHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		common.ErrorResponse(w, http.StatusBadRequest, common.HTTP_BAD_REQUEST, common.ERR_VALIDATION, nil)
		return
	}

	msg, err := h.service.Delete(uint(id))
	if err != nil {
		common.ErrorResponse(w, http.StatusNotFound, common.HTTP_NOT_FOUND, common.ERR_NOT_FOUND, nil)
		return
	}

	common.SuccessResponse(w, msg, nil, common.HTTP_OK)
}
