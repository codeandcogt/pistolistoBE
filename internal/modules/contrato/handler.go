package contrato

import (
	"net/http"
	"pistolistoBE/internal/common"
	"strconv"

	"github.com/gorilla/mux"
)

type ContratoHandler struct {
	service ContratoService
}

func NewContratoHandler(service ContratoService) *ContratoHandler {
	return &ContratoHandler{service}
}

func (h *ContratoHandler) GetByAvaluo(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["avaluoId"]
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		common.ErrorResponse(w, http.StatusBadRequest, common.HTTP_BAD_REQUEST, common.ERR_VALIDATION, nil)
		return
	}

	contrato, err := h.service.GetByAvaluo(uint(id))
	if err != nil {
		common.ErrorResponse(w, http.StatusNotFound, common.HTTP_NOT_FOUND, common.ERR_NOT_FOUND, nil)
		return
	}

	common.SuccessResponse(w, common.SUCCESS_RETRIEVED, contrato, common.HTTP_OK)
}

func (h *ContratoHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	contratos, err := h.service.GetAll()
	if err != nil {
		common.ErrorResponse(w, http.StatusInternalServerError, common.HTTP_SERVER_ERROR, common.ERR_DATABASE_ERROR, nil)
		return
	}
	common.SuccessResponse(w, common.SUCCESS_RETRIEVED, contratos, common.HTTP_OK)
}
