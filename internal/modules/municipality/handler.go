package municipality

import (
	"encoding/json"
	"net/http"
	"pistolistoBE/internal/common"
	"strconv"

	"github.com/gorilla/mux"
)

type MunicipalityHandler struct {
	service MunicipalityService
}

func NewMunicipalityHandler(service MunicipalityService) *MunicipalityHandler {
	return &MunicipalityHandler{service}
}

func (h *MunicipalityHandler) Create(w http.ResponseWriter, r *http.Request) {
	var municipality Municipality
	if err := json.NewDecoder(r.Body).Decode(&municipality); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.Create(&municipality); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	common.SuccessResponse(w, common.SUCCESS_CREATED, municipality, common.HTTP_CREATED)
}

func (h *MunicipalityHandler) GetByID(w http.ResponseWriter, r *http.Request) {
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

	municipality, err := h.service.GetByID(uint(id))
	if err != nil {
		common.ErrorResponse(w, http.StatusNotFound, common.HTTP_NOT_FOUND, common.ERR_NOT_FOUND, nil)
		return
	}

	common.SuccessResponse(w, common.SUCCESS_RETRIEVED, municipality, common.HTTP_OK)
}

func (h *MunicipalityHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	municipalities, err := h.service.GetAll()
	if err != nil {
		common.ErrorResponse(w, http.StatusNotFound, common.HTTP_NOT_FOUND, common.ERR_NOT_FOUND, nil)
		return
	}

	common.SuccessResponse(w, common.SUCCESS_RETRIEVED, municipalities, common.HTTP_OK)
}

func (h *MunicipalityHandler) Update(w http.ResponseWriter, r *http.Request) {
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

	var municipality Municipality
	if err := json.NewDecoder(r.Body).Decode(&municipality); err != nil {
		common.ErrorResponse(w, http.StatusBadRequest, common.HTTP_BAD_REQUEST, common.ERR_VALIDATION, nil)
		return
	}

	municipality.IdMunicipio = uint(id)

	updatedText, err := h.service.Update(&municipality)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	common.SuccessResponse(w, common.SUCCESS_UPDATED, updatedText, common.HTTP_OK)
}

func (h *MunicipalityHandler) Delete(w http.ResponseWriter, r *http.Request) {
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
