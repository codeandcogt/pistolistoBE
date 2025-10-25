package loan

import (
	"encoding/json"
	"net/http"
	"pistolistoBE/internal/common"
	"strconv"

	"github.com/gorilla/mux"
)

type LoanHandler struct {
	service LoanService
}

func NewLoanHandler(service LoanService) *LoanHandler {
	return &LoanHandler{service}
}

// Crear préstamo a partir del contrato
func (h *LoanHandler) CreateLoanFromContrato(w http.ResponseWriter, r *http.Request) {
	var body struct {
		IdContrato uint `json:"idContrato"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		common.ErrorResponse(w, http.StatusBadRequest, common.HTTP_BAD_REQUEST, common.ERR_INVALID_JSON, nil)
		return
	}

	loan, err := h.service.CreateLoanFromContrato(body.IdContrato)
	if err != nil {
		common.ErrorResponse(w, http.StatusInternalServerError, common.HTTP_SERVER_ERROR, err.Error(), nil)
		return
	}

	common.SuccessResponse(w, "Préstamo creado correctamente", loan, common.HTTP_CREATED)
}

// CRUD estándar
func (h *LoanHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, _ := strconv.ParseUint(idStr, 10, 32)
	loan, err := h.service.GetByID(uint(id))
	if err != nil {
		common.ErrorResponse(w, http.StatusNotFound, common.HTTP_NOT_FOUND, "Préstamo no encontrado", nil)
		return
	}
	common.SuccessResponse(w, common.SUCCESS_RETRIEVED, loan, common.HTTP_OK)
}

func (h *LoanHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	loans, err := h.service.GetAll()
	if err != nil {
		common.ErrorResponse(w, http.StatusInternalServerError, common.HTTP_SERVER_ERROR, err.Error(), nil)
		return
	}
	common.SuccessResponse(w, common.SUCCESS_RETRIEVED, loans, common.HTTP_OK)
}

func (h *LoanHandler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, _ := strconv.ParseUint(idStr, 10, 32)

	var loan Loan
	if err := json.NewDecoder(r.Body).Decode(&loan); err != nil {
		common.ErrorResponse(w, http.StatusBadRequest, common.HTTP_BAD_REQUEST, common.ERR_INVALID_JSON, nil)
		return
	}

	if err := h.service.Update(uint(id), &loan); err != nil {
		common.ErrorResponse(w, http.StatusInternalServerError, common.HTTP_SERVER_ERROR, err.Error(), nil)
		return
	}

	common.SuccessResponse(w, common.SUCCESS_UPDATED, loan, common.HTTP_OK)
}

func (h *LoanHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, _ := strconv.ParseUint(idStr, 10, 32)
	if err := h.service.Delete(uint(id)); err != nil {
		common.ErrorResponse(w, http.StatusInternalServerError, common.HTTP_SERVER_ERROR, err.Error(), nil)
		return
	}
	common.SuccessResponse(w, common.SUCCESS_DELETED, nil, common.HTTP_OK)
}
