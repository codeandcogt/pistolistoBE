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
		common.ErrorResponse(w, http.StatusBadRequest, common.ERR_INVALID_JSON, "Formato JSON inválido", nil)
		return
	}

	if err := h.service.CreateBanco(&banco); err != nil {
		common.ErrorResponse(w, http.StatusInternalServerError, common.ERR_INTERNAL_ERROR, "Error al crear banco", nil)
		return
	}

	common.CreatedResponse(w, "Banco creado exitosamente", banco, common.HTTP_CREATED)
}

func (h *BancoHandler) GetBancoByID(w http.ResponseWriter, r *http.Request) {
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

	banco, err := h.service.GetBancoByID(uint(id))
	if err != nil {
		common.ErrorResponse(w, http.StatusNotFound, common.ERR_NOT_FOUND, "Banco no encontrado", nil)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	common.SuccessResponse(w, "Banco encontrado", banco, common.HTTP_OK)
}

func (h *BancoHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	bancos, err := h.service.GetAll()
	if err != nil {
		common.ErrorResponse(w, http.StatusInternalServerError, common.ERR_DATABASE_ERROR, "Error al obtener bancos", nil)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	common.SuccessResponse(w, "Bancos obtenidos exitosamente", bancos, common.HTTP_OK)
}

func (h *BancoHandler) UpdateBanco(w http.ResponseWriter, r *http.Request) {
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

	var banco Banco
	if err := json.NewDecoder(r.Body).Decode(&banco); err != nil {
		common.ErrorResponse(w, http.StatusBadRequest, common.ERR_INVALID_JSON, "Formato JSON inválido", nil)
		return
	}

	if err := h.service.UpdateBanco(uint(id), &banco); err != nil {
		common.ErrorResponse(w, http.StatusInternalServerError, common.ERR_INTERNAL_ERROR, "Error al actualizar banco", nil)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	common.SuccessResponse(w, "Banco actualizado exitosamente", banco, common.HTTP_OK)
}

func (h *BancoHandler) DeleteBanco(w http.ResponseWriter, r *http.Request) {
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

	if err := h.service.DeleteBanco(uint(id)); err != nil {
		common.ErrorResponse(w, http.StatusInternalServerError, common.ERR_INTERNAL_ERROR, "Error al eliminar banco", nil)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	common.SuccessResponse(w, "Banco eliminado exitosamente", nil, common.HTTP_OK)
}

func (h *BancoHandler) GetBancosByTipo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tipo, exists := vars["tipo"]
	if !exists {
		common.ErrorResponse(w, http.StatusBadRequest, common.ERR_REQUIRED_FIELD, "Tipo requerido", nil)
		return
	}

	bancos, err := h.service.GetBancosByTipo(tipo)
	if err != nil {
		common.ErrorResponse(w, http.StatusInternalServerError, common.ERR_DATABASE_ERROR, "Error al obtener bancos por tipo", nil)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	common.SuccessResponse(w, "Bancos por tipo obtenidos exitosamente", bancos, common.HTTP_OK)
}
