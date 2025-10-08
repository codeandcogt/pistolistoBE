package permiso

import (
	"encoding/json"
	"net/http"
	"pistolistoBE/internal/common"
	"strconv"

	"github.com/gorilla/mux"
)

type PermisoHandler struct {
	service PermisoService
}

func NewPermisoHandler(service PermisoService) *PermisoHandler {
	return &PermisoHandler{service}
}

func (h *PermisoHandler) Create(w http.ResponseWriter, r *http.Request) {
	var permiso Permiso

	if err := json.NewDecoder(r.Body).Decode(&permiso); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.Create(&permiso); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	common.SuccessResponse(w, common.SUCCESS_CREATED, permiso, common.HTTP_CREATED)

}

func (h *PermisoHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	permiso, err := h.service.GetAll()

	if err != nil {
		common.ErrorResponse(w, http.StatusNotFound, common.HTTP_NOT_FOUND, common.ERR_NOT_FOUND, nil)
		return
	}

	w.WriteHeader(http.StatusCreated)
	common.SuccessResponse(w, common.SUCCESS_CREATED, permiso, common.HTTP_CREATED)

}

func (h *PermisoHandler) Update(w http.ResponseWriter, r *http.Request) {
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

	var permiso Permiso
	if err := json.NewDecoder(r.Body).Decode(&permiso); err != nil {
		common.ErrorResponse(w, http.StatusBadRequest, common.HTTP_BAD_REQUEST, common.ERR_VALIDATION, nil)
		return
	}

	permiso.IdPermiso = uint(id)

	updatedText, err := h.service.Update(&permiso)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	common.SuccessResponse(w, common.SUCCESS_UPDATED, updatedText, common.HTTP_OK)

}

func (h *PermisoHandler) GetByID(w http.ResponseWriter, r *http.Request) {
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

	client, err := h.service.GetByID(uint(id))
	if err != nil {
		common.ErrorResponse(w, http.StatusNotFound, common.HTTP_NOT_FOUND, common.ERR_NOT_FOUND, nil)
		return
	}

	common.SuccessResponse(w, common.SUCCESS_RETRIEVED, client, common.HTTP_OK)
}

func (h *PermisoHandler) Delete(w http.ResponseWriter, r *http.Request) {
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

	text, err := h.service.Delete(uint(id))
	if err != nil {
		common.ErrorResponse(w, http.StatusNotFound, common.HTTP_NOT_FOUND, common.ERR_NOT_FOUND, nil)
		return
	}

	common.SuccessResponse(w, common.SUCCESS_RETRIEVED, text, common.HTTP_OK)
}
