package logUbicacion

import (
	"encoding/json"
	"net/http"
	"pistolistoBE/internal/common"
	"strconv"

	"github.com/gorilla/mux"
)

type LogUbicacionHandler struct {
	service LogUbicacionService
}

func NewLogUbicacionHandler(service LogUbicacionService) *LogUbicacionHandler {
	return &LogUbicacionHandler{service}
}

// POST /api/log-ubicaciones
func (h *LogUbicacionHandler) RegistrarUbicacion(w http.ResponseWriter, r *http.Request) {
	var log LogUbicacion
	if err := json.NewDecoder(r.Body).Decode(&log); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.RegistrarUbicacion(&log); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	common.SuccessResponse(w, "Ubicación registrada y emitida correctamente", log, common.HTTP_CREATED)
}

// GET /api/log-ubicaciones/piloto/{id}
func (h *LogUbicacionHandler) GetByPiloto(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["pilotoId"]
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		common.ErrorResponse(w, http.StatusBadRequest, common.HTTP_BAD_REQUEST, common.ERR_VALIDATION, nil)
		return
	}

	logs, err := h.service.GetByPiloto(uint(id))
	if err != nil {
		common.ErrorResponse(w, http.StatusInternalServerError, common.HTTP_SERVER_ERROR, common.ERR_DATABASE_ERROR, nil)
		return
	}

	common.SuccessResponse(w, common.SUCCESS_RETRIEVED, logs, common.HTTP_OK)
}
