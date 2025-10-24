package estadoRuta

import (
	"encoding/json"
	"net/http"
	"pistolistoBE/internal/common"
	"strconv"

	"github.com/gorilla/mux"
)

type EstadoRutaHandler struct {
	service EstadoRutaService
}

func NewEstadoRutaHandler(service EstadoRutaService) *EstadoRutaHandler {
	return &EstadoRutaHandler{service}
}

func (h *EstadoRutaHandler) Create(w http.ResponseWriter, r *http.Request) {
	var estado EstadoRuta
	if err := json.NewDecoder(r.Body).Decode(&estado); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.Create(&estado); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	common.SuccessResponse(w, common.SUCCESS_CREATED, estado, common.HTTP_CREATED)
}

func (h *EstadoRutaHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		common.ErrorResponse(w, http.StatusBadRequest, common.HTTP_BAD_REQUEST, common.ERR_VALIDATION, nil)
		return
	}

	estado, err := h.service.GetByID(uint(id))
	if err != nil {
		common.ErrorResponse(w, http.StatusNotFound, common.HTTP_NOT_FOUND, common.ERR_NOT_FOUND, nil)
		return
	}

	common.SuccessResponse(w, common.SUCCESS_RETRIEVED, estado, common.HTTP_OK)
}

func (h *EstadoRutaHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	estados, err := h.service.GetAll()
	if err != nil {
		common.ErrorResponse(w, http.StatusInternalServerError, common.HTTP_SERVER_ERROR, common.ERR_DATABASE_ERROR, nil)
		return
	}

	common.SuccessResponse(w, common.SUCCESS_RETRIEVED, estados, common.HTTP_OK)
}

func (h *EstadoRutaHandler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		common.ErrorResponse(w, http.StatusBadRequest, common.HTTP_BAD_REQUEST, common.ERR_VALIDATION, nil)
		return
	}

	var estado EstadoRuta
	if err := json.NewDecoder(r.Body).Decode(&estado); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	estado.IdEstadoRuta = uint(id)

	msg, err := h.service.Update(&estado)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	common.SuccessResponse(w, msg, estado, common.HTTP_OK)
}

func (h *EstadoRutaHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		common.ErrorResponse(w, http.StatusBadRequest, common.HTTP_BAD_REQUEST, common.ERR_VALIDATION, nil)
		return
	}

	msg, err := h.service.Delete(uint(id))
	if err != nil {
		common.ErrorResponse(w, http.StatusNotFound, common.HTTP_NOT_FOUND, common.ERR_NOT_FOUND, nil)
		return
	}

	common.SuccessResponse(w, msg, nil, common.HTTP_OK)
}
