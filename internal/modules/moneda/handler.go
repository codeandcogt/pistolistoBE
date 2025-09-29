package moneda

import (
	"encoding/json"
	"net/http"
	"pistolistoBE/internal/common"
	"strconv"

	"github.com/gorilla/mux"
)

type MonedaHandler struct {
	service MonedaService
}

func NewMonedaHandler(service MonedaService) *MonedaHandler {
	return &MonedaHandler{service}
}

func (h *MonedaHandler) CreateMoneda(w http.ResponseWriter, r *http.Request) {
	var moneda Moneda
	if err := json.NewDecoder(r.Body).Decode(&moneda); err != nil {
		common.ErrorResponse(w, http.StatusBadRequest, common.ERR_INVALID_JSON, "Formato JSON inválido", nil)
		return
	}

	if err := h.service.CreateMoneda(&moneda); err != nil {
		common.ErrorResponse(w, http.StatusInternalServerError, common.ERR_INTERNAL_ERROR, "Error al crear moneda", nil)
		return
	}

	common.CreatedResponse(w, "Moneda creada exitosamente", moneda, common.HTTP_CREATED)
}

func (h *MonedaHandler) GetMonedaByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr, exists := vars["id"]
	if !exists {
		common.ErrorResponse(w, http.StatusBadRequest, common.ERR_REQUIRED_FIELD, "ID requerido", nil)
		return
	}

	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		common.ErrorResponse(w, http.StatusBadRequest, common.ERR_VALIDATION, "ID inválido", nil)
		return
	}

	moneda, err := h.service.GetMonedaByID(uint(id))
	if err != nil {
		common.ErrorResponse(w, http.StatusNotFound, common.ERR_NOT_FOUND, "Moneda no encontrada", nil)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	common.SuccessResponse(w, "Moneda encontrada", moneda, common.HTTP_OK)
}

func (h *MonedaHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	monedas, err := h.service.GetAll()
	if err != nil {
		common.ErrorResponse(w, http.StatusInternalServerError, common.ERR_DATABASE_ERROR, "Error al obtener monedas", nil)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	common.SuccessResponse(w, "Monedas obtenidas exitosamente", monedas, common.HTTP_OK)
}

func (h *MonedaHandler) UpdateMoneda(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr, exists := vars["id"]
	if !exists {
		common.ErrorResponse(w, http.StatusBadRequest, common.ERR_REQUIRED_FIELD, "ID requerido", nil)
		return
	}

	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		common.ErrorResponse(w, http.StatusBadRequest, common.ERR_VALIDATION, "ID inválido", nil)
		return
	}

	var moneda Moneda
	if err := json.NewDecoder(r.Body).Decode(&moneda); err != nil {
		common.ErrorResponse(w, http.StatusBadRequest, common.ERR_INVALID_JSON, "Formato JSON inválido", nil)
		return
	}

	if err := h.service.UpdateMoneda(uint(id), &moneda); err != nil {
		common.ErrorResponse(w, http.StatusInternalServerError, common.ERR_INTERNAL_ERROR, "Error al actualizar moneda", nil)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	common.SuccessResponse(w, "Moneda actualizada exitosamente", moneda, common.HTTP_OK)
}

func (h *MonedaHandler) DeleteMoneda(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr, exists := vars["id"]
	if !exists {
		common.ErrorResponse(w, http.StatusBadRequest, common.ERR_REQUIRED_FIELD, "ID requerido", nil)
		return
	}

	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		common.ErrorResponse(w, http.StatusBadRequest, common.ERR_VALIDATION, "ID inválido", nil)
		return
	}

	if err := h.service.DeleteMoneda(uint(id)); err != nil {
		common.ErrorResponse(w, http.StatusInternalServerError, common.ERR_INTERNAL_ERROR, "Error al eliminar moneda", nil)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	common.SuccessResponse(w, "Moneda eliminada exitosamente", nil, common.HTTP_OK)
}
