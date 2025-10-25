package cuota

import (
	"encoding/json"
	"net/http"
	"pistolistoBE/internal/common"
	"strconv"

	"github.com/gorilla/mux"
)

type CuotaHandler struct {
	service CuotaService
}

func NewCuotaHandler(service CuotaService) *CuotaHandler {
	return &CuotaHandler{service}
}

// Crear nueva cuota
func (h *CuotaHandler) Create(w http.ResponseWriter, r *http.Request) {
	var cuota Cuota
	if err := json.NewDecoder(r.Body).Decode(&cuota); err != nil {
		common.ErrorResponse(w, http.StatusBadRequest, common.HTTP_BAD_REQUEST, common.ERR_INVALID_JSON, nil)
		return
	}

	if err := h.service.Create(&cuota); err != nil {
		common.ErrorResponse(w, http.StatusInternalServerError, common.HTTP_SERVER_ERROR, err.Error(), nil)
		return
	}

	common.SuccessResponse(w, common.SUCCESS_CREATED, cuota, common.HTTP_CREATED)
}

// Obtener cuota por ID
func (h *CuotaHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, _ := strconv.ParseUint(idStr, 10, 32)

	cuota, err := h.service.GetByID(uint(id))
	if err != nil {
		common.ErrorResponse(w, http.StatusNotFound, common.HTTP_NOT_FOUND, "Cuota no encontrada", nil)
		return
	}

	common.SuccessResponse(w, common.SUCCESS_RETRIEVED, cuota, common.HTTP_OK)
}

// Obtener todas las cuotas
func (h *CuotaHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	cuotas, err := h.service.GetAll()
	if err != nil {
		common.ErrorResponse(w, http.StatusInternalServerError, common.HTTP_SERVER_ERROR, err.Error(), nil)
		return
	}

	common.SuccessResponse(w, common.SUCCESS_RETRIEVED, cuotas, common.HTTP_OK)
}

// Actualizar cuota
func (h *CuotaHandler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, _ := strconv.ParseUint(idStr, 10, 32)

	var cuota Cuota
	if err := json.NewDecoder(r.Body).Decode(&cuota); err != nil {
		common.ErrorResponse(w, http.StatusBadRequest, common.HTTP_BAD_REQUEST, common.ERR_INVALID_JSON, nil)
		return
	}

	if err := h.service.Update(uint(id), &cuota); err != nil {
		common.ErrorResponse(w, http.StatusInternalServerError, common.HTTP_SERVER_ERROR, err.Error(), nil)
		return
	}

	common.SuccessResponse(w, common.SUCCESS_UPDATED, cuota, common.HTTP_OK)
}

// Eliminar cuota (baja lógica)
func (h *CuotaHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, _ := strconv.ParseUint(idStr, 10, 32)

	if err := h.service.Delete(uint(id)); err != nil {
		common.ErrorResponse(w, http.StatusInternalServerError, common.HTTP_SERVER_ERROR, err.Error(), nil)
		return
	}

	common.SuccessResponse(w, common.SUCCESS_DELETED, nil, common.HTTP_OK)
}
