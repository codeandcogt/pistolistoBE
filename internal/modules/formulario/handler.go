package formulario

import (
	"encoding/json"
	"net/http"
	"pistolistoBE/internal/common"
	"pistolistoBE/internal/modules/avaluo"
	"strconv"

	"github.com/gorilla/mux"
)

type FormularioHandler struct {
	service   FormularioService
	avaluoSrv avaluo.AvaluoService
}

func NewFormularioHandler(service FormularioService, avaluoSrv avaluo.AvaluoService) *FormularioHandler {
	return &FormularioHandler{service, avaluoSrv}
}

func (h *FormularioHandler) CreateFormulario(w http.ResponseWriter, r *http.Request) {
	var formulario Formulario
	if err := json.NewDecoder(r.Body).Decode(&formulario); err != nil {
		common.ErrorResponse(w, http.StatusBadRequest, common.HTTP_BAD_REQUEST, common.ERR_INVALID_JSON, nil)
		return
	}

	if err := h.service.CreateFormulario(&formulario); err != nil {
		common.ErrorResponse(w, http.StatusInternalServerError, common.HTTP_SERVER_ERROR, "Error al crear formulario", nil)
		return
	}

	avaluoGenerado, err := h.avaluoSrv.CrearAvaluoAutomatico(formulario.IdFormulario)
	if err != nil {
		common.ErrorResponse(w, http.StatusInternalServerError, common.HTTP_SERVER_ERROR, "Error al crear avalúo automático", nil)
		return
	}

	contratoGenerado, err := h.avaluoSrv.GetContratoByAvaluo(avaluoGenerado.IdAvaluo)
	if err != nil {
		common.ErrorResponse(w, http.StatusInternalServerError, common.HTTP_SERVER_ERROR, "Error al obtener contrato", nil)
		return
	}

	articuloData, err := h.avaluoSrv.GetArticuloByFormulario(formulario.IdFormulario)
	if err != nil {
		common.ErrorResponse(w, http.StatusInternalServerError, common.HTTP_SERVER_ERROR, "Error al obtener artículo relacionado", nil)
		return
	}

	response := map[string]interface{}{
		"mensaje":    "Formulario, avalúo, contrato y artículo vinculados correctamente",
		"formulario": formulario,
		"avaluo":     avaluoGenerado,
		"contrato":   contratoGenerado,
		"articulo":   articuloData,
	}

	common.SuccessResponse(w, "Creado correctamente", response, common.HTTP_CREATED)
}

func (h *FormularioHandler) GetFormularioByID(w http.ResponseWriter, r *http.Request) {
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

	formulario, err := h.service.GetFormularioByID(uint(id))
	if err != nil {
		common.ErrorResponse(w, http.StatusNotFound, common.HTTP_NOT_FOUND, common.ERR_NOT_FOUND, nil)
		return
	}

	common.SuccessResponse(w, common.SUCCESS_RETRIEVED, formulario, common.HTTP_OK)
}

func (h *FormularioHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	formularios, err := h.service.GetAll()
	if err != nil {
		common.ErrorResponse(w, http.StatusInternalServerError, common.HTTP_SERVER_ERROR, common.ERR_DATABASE_ERROR, nil)
		return
	}

	common.SuccessResponse(w, common.SUCCESS_RETRIEVED, formularios, common.HTTP_OK)
}

func (h *FormularioHandler) UpdateFormulario(w http.ResponseWriter, r *http.Request) {
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

	var formulario Formulario
	if err := json.NewDecoder(r.Body).Decode(&formulario); err != nil {
		common.ErrorResponse(w, http.StatusBadRequest, common.HTTP_BAD_REQUEST, common.ERR_INVALID_JSON, nil)
		return
	}

	if err := h.service.UpdateFormulario(uint(id), &formulario); err != nil {
		common.ErrorResponse(w, http.StatusInternalServerError, common.HTTP_SERVER_ERROR, common.ERR_INTERNAL_ERROR, nil)
		return
	}

	common.SuccessResponse(w, common.SUCCESS_UPDATED, formulario, common.HTTP_OK)
}

func (h *FormularioHandler) DeleteFormulario(w http.ResponseWriter, r *http.Request) {
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

	if err := h.service.DeleteFormulario(uint(id)); err != nil {
		common.ErrorResponse(w, http.StatusInternalServerError, common.HTTP_SERVER_ERROR, common.ERR_INTERNAL_ERROR, nil)
		return
	}

	common.SuccessResponse(w, common.SUCCESS_DELETED, nil, common.HTTP_OK)
}
