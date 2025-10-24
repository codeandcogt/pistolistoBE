package inventario

import (
	"encoding/json"
	"net/http"
	"pistolistoBE/internal/common"
	"strconv"

	"github.com/gorilla/mux"
)

type InventarioHandler struct {
	service InventarioService
}

func NewInventarioHandler(service InventarioService) *InventarioHandler {
	return &InventarioHandler{service}
}

func (h *InventarioHandler) CreateInventario(w http.ResponseWriter, r *http.Request) {
	var inv Inventario
	if err := json.NewDecoder(r.Body).Decode(&inv); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.CreateInventario(&inv); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	common.SuccessResponse(w, common.SUCCESS_CREATED, inv, common.HTTP_CREATED)
}

func (h *InventarioHandler) GetInventarioByID(w http.ResponseWriter, r *http.Request) {
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

	inv, err := h.service.GetInventarioByID(uint(id))
	if err != nil {
		common.ErrorResponse(w, http.StatusNotFound, common.HTTP_NOT_FOUND, common.ERR_NOT_FOUND, nil)
		return
	}

	common.SuccessResponse(w, common.SUCCESS_RETRIEVED, inv, common.HTTP_OK)
}

func (h *InventarioHandler) GetAllByAlmacen(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	IdAlmacenStr, exists := vars["id_almacen"]
	if !exists {
		common.ErrorResponse(w, http.StatusBadRequest, common.HTTP_BAD_REQUEST, common.ERR_REQUIRED_FIELD, nil)
		return
	}

	IdAlmacen, err := strconv.ParseUint(IdAlmacenStr, 10, 32)
	if err != nil {
		common.ErrorResponse(w, http.StatusBadRequest, common.HTTP_BAD_REQUEST, common.ERR_VALIDATION, nil)
		return
	}

	inv, err := h.service.GetAllByAlmacen(uint(IdAlmacen))
	if err != nil {
		common.ErrorResponse(w, http.StatusNotFound, common.HTTP_NOT_FOUND, common.ERR_NOT_FOUND, nil)
		return
	}

	common.SuccessResponse(w, common.SUCCESS_RETRIEVED, inv, common.HTTP_OK)
}

func (h *InventarioHandler) GetByTipo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	IdAlmacenStr, exists := vars["id_almacen"]
	if !exists {
		common.ErrorResponse(w, http.StatusBadRequest, common.HTTP_BAD_REQUEST, common.ERR_REQUIRED_FIELD, nil)
		return
	}

	tipoItem, exists := vars["tipo"]
	if !exists {
		common.ErrorResponse(w, http.StatusBadRequest, common.HTTP_BAD_REQUEST, common.ERR_REQUIRED_FIELD, nil)
		return
	}

	IdAlmacen, err := strconv.ParseUint(IdAlmacenStr, 10, 32)
	if err != nil {
		common.ErrorResponse(w, http.StatusBadRequest, common.HTTP_BAD_REQUEST, common.ERR_VALIDATION, nil)
		return
	}

	inv, err := h.service.GetByTipo(uint(IdAlmacen), tipoItem)
	if err != nil {
		common.ErrorResponse(w, http.StatusNotFound, common.HTTP_NOT_FOUND, common.ERR_NOT_FOUND, nil)
		return
	}

	common.SuccessResponse(w, common.SUCCESS_RETRIEVED, inv, common.HTTP_OK)
}

func (h *InventarioHandler) GetByEstado(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	IdAlmacenStr, exists := vars["id_almacen"]
	if !exists {
		common.ErrorResponse(w, http.StatusBadRequest, common.HTTP_BAD_REQUEST, common.ERR_REQUIRED_FIELD, nil)
		return
	}

	estadoInventario, exists := vars["estado"]
	if !exists {
		common.ErrorResponse(w, http.StatusBadRequest, common.HTTP_BAD_REQUEST, common.ERR_REQUIRED_FIELD, nil)
		return
	}

	IdAlmacen, err := strconv.ParseUint(IdAlmacenStr, 10, 32)
	if err != nil {
		common.ErrorResponse(w, http.StatusBadRequest, common.HTTP_BAD_REQUEST, common.ERR_VALIDATION, nil)
		return
	}

	inv, err := h.service.GetByEstado(uint(IdAlmacen), estadoInventario)
	if err != nil {
		common.ErrorResponse(w, http.StatusNotFound, common.HTTP_NOT_FOUND, common.ERR_NOT_FOUND, nil)
		return
	}

	common.SuccessResponse(w, common.SUCCESS_RETRIEVED, inv, common.HTTP_OK)
}

func (h *InventarioHandler) GetProductosDisponibles(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	IdAlmacenStr, exists := vars["id_almacen"]
	if !exists {
		common.ErrorResponse(w, http.StatusBadRequest, common.HTTP_BAD_REQUEST, common.ERR_REQUIRED_FIELD, nil)
		return
	}

	IdAlmacen, err := strconv.ParseUint(IdAlmacenStr, 10, 32)
	if err != nil {
		common.ErrorResponse(w, http.StatusBadRequest, common.HTTP_BAD_REQUEST, common.ERR_VALIDATION, nil)
		return
	}

	inv, err := h.service.GetProductosDisponibles(uint(IdAlmacen))
	if err != nil {
		common.ErrorResponse(w, http.StatusNotFound, common.HTTP_NOT_FOUND, common.ERR_NOT_FOUND, nil)
		return
	}

	common.SuccessResponse(w, common.SUCCESS_RETRIEVED, inv, common.HTTP_OK)
}

func (h *InventarioHandler) GetArticulosEmpenados(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	IdAlmacenStr, exists := vars["id_almacen"]
	if !exists {
		common.ErrorResponse(w, http.StatusBadRequest, common.HTTP_BAD_REQUEST, common.ERR_REQUIRED_FIELD, nil)
		return
	}

	IdAlmacen, err := strconv.ParseUint(IdAlmacenStr, 10, 32)
	if err != nil {
		common.ErrorResponse(w, http.StatusBadRequest, common.HTTP_BAD_REQUEST, common.ERR_VALIDATION, nil)
		return
	}

	inv, err := h.service.GetArticulosEmpenados(uint(IdAlmacen))
	if err != nil {
		common.ErrorResponse(w, http.StatusNotFound, common.HTTP_NOT_FOUND, common.ERR_NOT_FOUND, nil)
		return
	}

	common.SuccessResponse(w, common.SUCCESS_RETRIEVED, inv, common.HTTP_OK)
}

func (h *InventarioHandler) UpdateInventario(w http.ResponseWriter, r *http.Request) {
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

	var updated Inventario
	if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
		common.ErrorResponse(w, http.StatusBadRequest, common.HTTP_BAD_REQUEST, common.ERR_VALIDATION, nil)
		return
	}

	inv, err := h.service.UpdateInventario(uint(id), &updated)
	if err != nil {
		common.ErrorResponse(w, http.StatusInternalServerError, common.HTTP_BAD_REQUEST, common.ERR_DATABASE_ERROR, nil)
		return
	}

	common.SuccessResponse(w, common.SUCCESS_UPDATED, inv, common.HTTP_OK)
}

func (h *InventarioHandler) DeleteInventario(w http.ResponseWriter, r *http.Request) {
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

	text, err := h.service.DeleteInventario(uint(id))
	if err != nil {
		common.ErrorResponse(w, http.StatusNotFound, common.HTTP_NOT_FOUND, common.ERR_NOT_FOUND, nil)
		return
	}

	common.SuccessResponse(w, common.SUCCESS_DELETED, text, common.HTTP_OK)
}
