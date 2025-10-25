package avaluo

import (
	"encoding/json"
	"net/http"
	"pistolistoBE/internal/common"
	"strconv"

	"github.com/gorilla/mux"
)

type AvaluoHandler struct {
	service AvaluoService
}

func NewAvaluoHandler(service AvaluoService) *AvaluoHandler {
	return &AvaluoHandler{service}
}

func (h *AvaluoHandler) CrearAvaluoManual(w http.ResponseWriter, r *http.Request) {
	var avaluo Avaluo
	if err := json.NewDecoder(r.Body).Decode(&avaluo); err != nil {
		common.ErrorResponse(w, http.StatusBadRequest, common.HTTP_BAD_REQUEST, common.ERR_INVALID_JSON, nil)
		return
	}
	if err := h.service.CrearAvaluoManual(&avaluo); err != nil {
		common.ErrorResponse(w, http.StatusInternalServerError, common.HTTP_SERVER_ERROR, common.ERR_DATABASE_ERROR, nil)
		return
	}
	common.SuccessResponse(w, common.SUCCESS_CREATED, avaluo, common.HTTP_CREATED)
}

func (h *AvaluoHandler) CrearAvaluoAutomatico(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	formIDStr := vars["formularioId"]
	formID, err := strconv.ParseUint(formIDStr, 10, 32)
	if err != nil {
		common.ErrorResponse(w, http.StatusBadRequest, common.HTTP_BAD_REQUEST, common.ERR_VALIDATION, nil)
		return
	}

	avaluo, err := h.service.CrearAvaluoAutomatico(uint(formID))
	if err != nil {
		common.ErrorResponse(w, http.StatusInternalServerError, common.HTTP_SERVER_ERROR, err.Error(), nil)
		return
	}
	common.SuccessResponse(w, "Avalúo generado automáticamente", avaluo, common.HTTP_CREATED)
}
