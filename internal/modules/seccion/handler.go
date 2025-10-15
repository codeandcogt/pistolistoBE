package seccion

import (
	"encoding/json"
	"net/http"
	"pistolistoBE/internal/common"
	"strconv"

	"github.com/gorilla/mux"
)

type SeccionHandler struct {
	service SeccionService
}

func NewSeccionHandler(service SeccionService) *SeccionHandler {
	return &SeccionHandler{service}
}

func (h *SeccionHandler) CreateSeccion(w http.ResponseWriter, r *http.Request) {
	var seccion Seccion
	if err := json.NewDecoder(r.Body).Decode(&seccion); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.Create(&seccion); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	common.SuccessResponse(w, common.SUCCESS_CREATED, seccion, common.HTTP_CREATED)
}

func (h *SeccionHandler) GetSeccionByID(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		common.ErrorResponse(w, http.StatusBadRequest, common.HTTP_BAD_REQUEST, common.ERR_VALIDATION, nil)
		return
	}

	seccion, err := h.service.GetByID(uint(id))
	if err != nil {
		common.ErrorResponse(w, http.StatusNotFound, common.HTTP_NOT_FOUND, common.ERR_NOT_FOUND, nil)
		return
	}

	common.SuccessResponse(w, common.SUCCESS_RETRIEVED, seccion, common.HTTP_OK)
}

func (h *SeccionHandler) GetAllSecciones(w http.ResponseWriter, r *http.Request) {
	secciones, err := h.service.GetAll()
	if err != nil {
		common.ErrorResponse(w, http.StatusInternalServerError, common.HTTP_BAD_REQUEST, common.ERR_DATABASE_ERROR, nil)
		return
	}

	common.SuccessResponse(w, common.SUCCESS_RETRIEVED, secciones, common.HTTP_OK)
}

func (h *SeccionHandler) UpdateSeccion(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		common.ErrorResponse(w, http.StatusBadRequest, common.HTTP_BAD_REQUEST, common.ERR_VALIDATION, nil)
		return
	}

	var updated Seccion
	if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	seccion, err := h.service.UpdateSeccion(uint(id), &updated)
	if err != nil {
		common.ErrorResponse(w, http.StatusInternalServerError, common.HTTP_BAD_REQUEST, common.ERR_DATABASE_ERROR, nil)
		return
	}

	common.SuccessResponse(w, common.SUCCESS_UPDATED, seccion, common.HTTP_OK)
}

func (h *SeccionHandler) DeleteSeccion(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		common.ErrorResponse(w, http.StatusBadRequest, common.HTTP_BAD_REQUEST, common.ERR_VALIDATION, nil)
		return
	}

	msg, err := h.service.DeleteSeccion(uint(id))
	if err != nil {
		common.ErrorResponse(w, http.StatusNotFound, common.HTTP_NOT_FOUND, common.ERR_NOT_FOUND, nil)
		return
	}

	common.SuccessResponse(w, common.SUCCESS_DELETED, msg, common.HTTP_OK)
}
