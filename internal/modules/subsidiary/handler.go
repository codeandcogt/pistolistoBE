package subsidiary

import (
	"encoding/json"
	"net/http"
	"pistolistoBE/internal/common"
	"strconv"

	"github.com/gorilla/mux"
)

type SubsidiaryHandler struct {
	service SubsidiaryService
}

func NewSubsidiaryHandler(service SubsidiaryService) *SubsidiaryHandler {
	return &SubsidiaryHandler{service}
}

func (h *SubsidiaryHandler) Create(w http.ResponseWriter, r *http.Request) {
	var subsidiary Subsidiary
	if err := json.NewDecoder(r.Body).Decode(&subsidiary); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.Create(&subsidiary); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	common.SuccessResponse(w, common.SUCCESS_CREATED, subsidiary, common.HTTP_CREATED)
}

func (h *SubsidiaryHandler) GetByID(w http.ResponseWriter, r *http.Request) {
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

	subsidiary, err := h.service.GetByID(uint(id))
	if err != nil {
		common.ErrorResponse(w, http.StatusNotFound, common.HTTP_NOT_FOUND, common.ERR_NOT_FOUND, nil)
		return
	}

	common.SuccessResponse(w, common.SUCCESS_RETRIEVED, subsidiary, common.HTTP_OK)
}

func (h *SubsidiaryHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	subsidiaries, err := h.service.GetAll()
	if err != nil {
		common.ErrorResponse(w, http.StatusNotFound, common.HTTP_NOT_FOUND, common.ERR_NOT_FOUND, nil)
		return
	}

	common.SuccessResponse(w, common.SUCCESS_RETRIEVED, subsidiaries, common.HTTP_OK)
}

func (h *SubsidiaryHandler) Update(w http.ResponseWriter, r *http.Request) {
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

	var subsidiary Subsidiary
	if err := json.NewDecoder(r.Body).Decode(&subsidiary); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	subsidiary.IdSucursal = uint(id)

	updatedSubsidiary, err := h.service.Update(&subsidiary)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	common.SuccessResponse(w, common.SUCCESS_UPDATED, updatedSubsidiary, common.HTTP_OK)
}

func (h *SubsidiaryHandler) Delete(w http.ResponseWriter, r *http.Request) {
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

	common.SuccessResponse(w, common.SUCCESS_DELETED, text, common.HTTP_OK)
}
