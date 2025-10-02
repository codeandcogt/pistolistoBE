package banco

import (
	"encoding/json"
	"net/http"
	"pistolistoBE/internal/common"
	"strconv"

	"github.com/gorilla/mux"
)

type BancoHandler struct {
	service BancoService
}

func NewBancoHandler(service BancoService) *BancoHandler {
	return &BancoHandler{service}
}

func (h *BancoHandler) CreateBanco(w http.ResponseWriter, r *http.Request) {
	var banco Banco
	if err := json.NewDecoder(r.Body).Decode(&banco); err != nil {
		common.ErrorResponse(w, http.StatusBadRequest, common.HTTP_BAD_REQUEST, common.ERR_INVALID_JSON, nil)
		return
	}

	if err := h.service.CreateBanco(&banco); err != nil {
		common.ErrorResponse(w, http.StatusInternalServerError, common.HTTP_SERVER_ERROR, common.ERR_INTERNAL_ERROR, nil)
		return
	}

	w.WriteHeader(http.StatusCreated)
	common.SuccessResponse(w, common.SUCCESS_CREATED, banco, common.HTTP_CREATED)
}

func (h *BancoHandler) GetBancoByID(w http.ResponseWriter, r *http.Request) {
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

	banco, err := h.service.GetBancoByID(uint(id))
	if err != nil {
		common.ErrorResponse(w, http.StatusNotFound, common.HTTP_NOT_FOUND, common.ERR_NOT_FOUND, nil)
		return
	}

	common.SuccessResponse(w, common.SUCCESS_RETRIEVED, banco, common.HTTP_OK)
}

func (h *BancoHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	bancos, err := h.service.GetAll()
	if err != nil {
		common.ErrorResponse(w, http.StatusInternalServerError, common.HTTP_SERVER_ERROR, common.ERR_DATABASE_ERROR, nil)
		return
	}

	common.SuccessResponse(w, common.SUCCESS_RETRIEVED, bancos, common.HTTP_OK)
}

func (h *BancoHandler) UpdateBanco(w http.ResponseWriter, r *http.Request) {
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

	var banco Banco
	if err := json.NewDecoder(r.Body).Decode(&banco); err != nil {
		common.ErrorResponse(w, http.StatusBadRequest, common.HTTP_BAD_REQUEST, common.ERR_INVALID_JSON, nil)
		return
	}

	if err := h.service.UpdateBanco(uint(id), &banco); err != nil {
		common.ErrorResponse(w, http.StatusInternalServerError, common.HTTP_SERVER_ERROR, common.ERR_INTERNAL_ERROR, nil)
		return
	}

	common.SuccessResponse(w, common.SUCCESS_UPDATED, banco, common.HTTP_OK)
}

func (h *BancoHandler) DeleteBanco(w http.ResponseWriter, r *http.Request) {
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

	if err := h.service.DeleteBanco(uint(id)); err != nil {
		common.ErrorResponse(w, http.StatusInternalServerError, common.HTTP_SERVER_ERROR, common.ERR_INTERNAL_ERROR, nil)
		return
	}

	common.SuccessResponse(w, common.SUCCESS_DELETED, nil, common.HTTP_OK)
}
