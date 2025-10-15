package almacenseccion

import (
	"encoding/json"
	"net/http"
	"pistolistoBE/internal/common"
	"strconv"

	"github.com/gorilla/mux"
)

type AlmacenSeccionHandler struct {
	service AlmacenSeccionService
}

func NewAlmacenSeccionHandler(service AlmacenSeccionService) *AlmacenSeccionHandler {
	return &AlmacenSeccionHandler{service}
}

func (h *AlmacenSeccionHandler) CreateAlmacenSeccion(w http.ResponseWriter, r *http.Request) {
	var almseccion AlmacenSeccion
	if err := json.NewDecoder(r.Body).Decode(&almseccion); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.Create(&almseccion); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	common.SuccessResponse(w, common.SUCCESS_CREATED, almseccion, common.HTTP_CREATED)
}

func (h *AlmacenSeccionHandler) GetAlmacenSeccionByID(w http.ResponseWriter, r *http.Request) {
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

func (h *AlmacenSeccionHandler) GetAllByAlmacen(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id_almacen"]
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		common.ErrorResponse(w, http.StatusBadRequest, common.HTTP_BAD_REQUEST, common.ERR_VALIDATION, nil)
		return
	}

	secciones, err := h.service.GetByAlmacen(uint(id))
	if err != nil {
		common.ErrorResponse(w, http.StatusNotFound, common.HTTP_NOT_FOUND, common.ERR_NOT_FOUND, nil)
		return
	}

	common.SuccessResponse(w, common.SUCCESS_RETRIEVED, secciones, common.HTTP_OK)
}

func (h *AlmacenSeccionHandler) UpdateAlmacenSeccion(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		common.ErrorResponse(w, http.StatusBadRequest, common.HTTP_BAD_REQUEST, common.ERR_VALIDATION, nil)
		return
	}

	var updated AlmacenSeccion
	if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	seccion, err := h.service.UpdateAlmacenSeccion(uint(id), &updated)
	if err != nil {
		common.ErrorResponse(w, http.StatusInternalServerError, common.HTTP_BAD_REQUEST, common.ERR_DATABASE_ERROR, nil)
		return
	}

	common.SuccessResponse(w, common.SUCCESS_UPDATED, seccion, common.HTTP_OK)
}
func (h *AlmacenSeccionHandler) DeleteAlmacenSeccion(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		common.ErrorResponse(w, http.StatusBadRequest, common.HTTP_BAD_REQUEST, common.ERR_VALIDATION, nil)
		return
	}

	msg, err := h.service.DeleteAlmacenSeccion(uint(id))
	if err != nil {
		common.ErrorResponse(w, http.StatusNotFound, common.HTTP_NOT_FOUND, common.ERR_NOT_FOUND, nil)
		return
	}

	common.SuccessResponse(w, common.SUCCESS_DELETED, msg, common.HTTP_OK)
}
