package rol

import (
	"encoding/json"
	"net/http"
	"pistolistoBE/internal/common"
)

type RolHandler struct {
	service RolService
}

func NewRolHandler(service RolService) *RolHandler {
	return &RolHandler{service}
}

func (h *RolHandler) CreateRol(w http.ResponseWriter, r *http.Request) {
	var rol Rol

	if err := json.NewDecoder(r.Body).Decode(&rol); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.CreateRol(&rol); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	common.SuccessResponse(w, common.SUCCESS_CREATED, rol, common.HTTP_CREATED)

}

func (h *RolHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	rol, err := h.service.GetAll()

	if err != nil {
		common.ErrorResponse(w, http.StatusNotFound, common.HTTP_NOT_FOUND, common.ERR_NOT_FOUND, nil)
		return
	}

	w.WriteHeader(http.StatusCreated)
	common.SuccessResponse(w, common.SUCCESS_CREATED, rol, common.HTTP_CREATED)

}
