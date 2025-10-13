package resenaEmpresa

import (
	"encoding/json"
	"net/http"
	"pistolistoBE/internal/common"
	"strconv"

	"github.com/gorilla/mux"
)

type ResenaEmpresaHandler struct {
	service ResenaEmpresaService
}

func NewResenaEmpresaHandler(service ResenaEmpresaService) *ResenaEmpresaHandler {
	return &ResenaEmpresaHandler{service}
}

func (h *ResenaEmpresaHandler) CreateResena(w http.ResponseWriter, r *http.Request) {
	var resena ResenaEmpresa
	if err := json.NewDecoder(r.Body).Decode(&resena); err != nil {
		common.ErrorResponse(w, http.StatusBadRequest, common.HTTP_BAD_REQUEST, common.ERR_INVALID_JSON, nil)
		return
	}

	if err := h.service.CreateResena(&resena); err != nil {
		common.ErrorResponse(w, http.StatusInternalServerError, common.HTTP_SERVER_ERROR, common.ERR_INTERNAL_ERROR, nil)
		return
	}

	w.WriteHeader(http.StatusCreated)
	common.SuccessResponse(w, common.SUCCESS_CREATED, resena, common.HTTP_CREATED)
}

func (h *ResenaEmpresaHandler) GetResenaByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr, exists := vars["id"]
	if !exists {
		common.ErrorResponse(w, http.StatusBadRequest, common.HTTP_BAD_REQUEST, common.ERR_REQUIRED_FIELD, nil)
		return
	}

	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		common.ErrorResponse(w, http.StatusBadRequest, common.HTTP_BAD_REQUEST, common.ERR_VALIDATION, nil)
		return
	}

	resena, err := h.service.GetResenaByID(uint(id))
	if err != nil {
		common.ErrorResponse(w, http.StatusNotFound, common.HTTP_NOT_FOUND, common.ERR_NOT_FOUND, nil)
		return
	}

	common.SuccessResponse(w, common.SUCCESS_RETRIEVED, resena, common.HTTP_OK)
}

func (h *ResenaEmpresaHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	resenas, err := h.service.GetAll()
	if err != nil {
		common.ErrorResponse(w, http.StatusInternalServerError, common.HTTP_SERVER_ERROR, common.ERR_DATABASE_ERROR, nil)
		return
	}

	common.SuccessResponse(w, common.SUCCESS_RETRIEVED, resenas, common.HTTP_OK)
}

func (h *ResenaEmpresaHandler) GetResenasByClienteID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	clienteIdStr, exists := vars["clienteId"]
	if !exists {
		common.ErrorResponse(w, http.StatusBadRequest, common.HTTP_BAD_REQUEST, common.ERR_REQUIRED_FIELD, nil)
		return
	}

	clienteId, err := strconv.ParseUint(clienteIdStr, 10, 32)
	if err != nil {
		common.ErrorResponse(w, http.StatusBadRequest, common.HTTP_BAD_REQUEST, common.ERR_VALIDATION, nil)
		return
	}

	resenas, err := h.service.GetResenasByClienteID(uint(clienteId))
	if err != nil {
		common.ErrorResponse(w, http.StatusInternalServerError, common.HTTP_SERVER_ERROR, common.ERR_DATABASE_ERROR, nil)
		return
	}

	common.SuccessResponse(w, common.SUCCESS_RETRIEVED, resenas, common.HTTP_OK)
}

func (h *ResenaEmpresaHandler) GetResenaByFacturaID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	facturaIdStr, exists := vars["facturaId"]
	if !exists {
		common.ErrorResponse(w, http.StatusBadRequest, common.HTTP_BAD_REQUEST, common.ERR_REQUIRED_FIELD, nil)
		return
	}

	facturaId, err := strconv.ParseUint(facturaIdStr, 10, 32)
	if err != nil {
		common.ErrorResponse(w, http.StatusBadRequest, common.HTTP_BAD_REQUEST, common.ERR_VALIDATION, nil)
		return
	}

	resena, err := h.service.GetResenaByFacturaID(uint(facturaId))
	if err != nil {
		common.ErrorResponse(w, http.StatusNotFound, common.HTTP_NOT_FOUND, common.ERR_NOT_FOUND, nil)
		return
	}

	common.SuccessResponse(w, common.SUCCESS_RETRIEVED, resena, common.HTTP_OK)
}

func (h *ResenaEmpresaHandler) UpdateResena(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr, exists := vars["id"]
	if !exists {
		common.ErrorResponse(w, http.StatusBadRequest, common.HTTP_BAD_REQUEST, common.ERR_REQUIRED_FIELD, nil)
		return
	}

	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		common.ErrorResponse(w, http.StatusBadRequest, common.HTTP_BAD_REQUEST, common.ERR_VALIDATION, nil)
		return
	}

	var resena ResenaEmpresa
	if err := json.NewDecoder(r.Body).Decode(&resena); err != nil {
		common.ErrorResponse(w, http.StatusBadRequest, common.HTTP_BAD_REQUEST, common.ERR_INVALID_JSON, nil)
		return
	}

	if err := h.service.UpdateResena(uint(id), &resena); err != nil {
		common.ErrorResponse(w, http.StatusInternalServerError, common.HTTP_SERVER_ERROR, common.ERR_INTERNAL_ERROR, nil)
		return
	}

	common.SuccessResponse(w, common.SUCCESS_UPDATED, resena, common.HTTP_OK)
}

func (h *ResenaEmpresaHandler) DeleteResena(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr, exists := vars["id"]
	if !exists {
		common.ErrorResponse(w, http.StatusBadRequest, common.HTTP_BAD_REQUEST, common.ERR_REQUIRED_FIELD, nil)
		return
	}

	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		common.ErrorResponse(w, http.StatusBadRequest, common.HTTP_BAD_REQUEST, common.ERR_VALIDATION, nil)
		return
	}

	if err := h.service.DeleteResena(uint(id)); err != nil {
		common.ErrorResponse(w, http.StatusInternalServerError, common.HTTP_SERVER_ERROR, common.ERR_INTERNAL_ERROR, nil)
		return
	}

	common.SuccessResponse(w, common.SUCCESS_DELETED, nil, common.HTTP_OK)
}

func (h *ResenaEmpresaHandler) GetPromedioCalificacion(w http.ResponseWriter, r *http.Request) {
	promedio, err := h.service.GetPromedioCalificacion()
	if err != nil {
		common.ErrorResponse(w, http.StatusInternalServerError, common.HTTP_SERVER_ERROR, common.ERR_DATABASE_ERROR, nil)
		return
	}

	response := map[string]float64{"promedioCalificacion": promedio}
	common.SuccessResponse(w, common.SUCCESS_RETRIEVED, response, common.HTTP_OK)
}
