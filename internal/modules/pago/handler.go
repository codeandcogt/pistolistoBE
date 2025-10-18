package pago

import (
	"encoding/json"
	"net/http"
	"pistolistoBE/internal/common"
	"strconv"

	"github.com/gorilla/mux"
)

type PagoHandler struct {
	service PagoService
}

func NewPagoHandler(service PagoService) *PagoHandler {
	return &PagoHandler{service}
}

// Crear y procesar pago
func (h *PagoHandler) ProcesarPago(w http.ResponseWriter, r *http.Request) {
	var pago Pago
	if err := json.NewDecoder(r.Body).Decode(&pago); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newPago, err := h.service.ProcesarPago(&pago)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	common.SuccessResponse(w, "Pago procesado exitosamente", newPago, common.HTTP_CREATED)
}

// Obtener todos los pagos
func (h *PagoHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	pagos, err := h.service.GetAll()
	if err != nil {
		common.ErrorResponse(w, http.StatusNotFound, common.HTTP_NOT_FOUND, common.ERR_NOT_FOUND, nil)
		return
	}
	common.SuccessResponse(w, common.SUCCESS_RETRIEVED, pagos, common.HTTP_OK)
}

// Obtener pago por ID
func (h *PagoHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		common.ErrorResponse(w, http.StatusBadRequest, common.HTTP_BAD_REQUEST, common.ERR_VALIDATION, nil)
		return
	}

	pago, err := h.service.GetByID(uint(id))
	if err != nil {
		common.ErrorResponse(w, http.StatusNotFound, common.HTTP_NOT_FOUND, common.ERR_NOT_FOUND, nil)
		return
	}

	common.SuccessResponse(w, common.SUCCESS_RETRIEVED, pago, common.HTTP_OK)
}

// Obtener pagos por pedido
func (h *PagoHandler) GetByPedido(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["pedidoId"]
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		common.ErrorResponse(w, http.StatusBadRequest, common.HTTP_BAD_REQUEST, common.ERR_VALIDATION, nil)
		return
	}

	pagos, err := h.service.GetByPedido(uint(id))
	if err != nil {
		common.ErrorResponse(w, http.StatusNotFound, common.HTTP_NOT_FOUND, common.ERR_NOT_FOUND, nil)
		return
	}

	common.SuccessResponse(w, common.SUCCESS_RETRIEVED, pagos, common.HTTP_OK)
}

// Actualizar pago
func (h *PagoHandler) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		common.ErrorResponse(w, http.StatusBadRequest, common.HTTP_BAD_REQUEST, common.ERR_VALIDATION, nil)
		return
	}

	var pago Pago
	if err := json.NewDecoder(r.Body).Decode(&pago); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	pago.IdPago = uint(id)

	msg, err := h.service.Update(&pago)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	common.SuccessResponse(w, msg, pago, common.HTTP_OK)
}

// Eliminar (baja lógica)
func (h *PagoHandler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
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
