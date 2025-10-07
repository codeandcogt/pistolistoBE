package bankAccount

import (
	"encoding/json"
	"net/http"
	"pistolistoBE/internal/common"
	"strconv"

	"github.com/gorilla/mux"
)

type BankAccountHandler struct {
	service BankAccountService
}

func NewBankAccountHandler(service BankAccountService) *BankAccountHandler {
	return &BankAccountHandler{service}
}

func (h *BankAccountHandler) Create(w http.ResponseWriter, r *http.Request) {
	var bankAccount BankAccount
	if err := json.NewDecoder(r.Body).Decode(&bankAccount); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.Create(&bankAccount); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	common.SuccessResponse(w, common.SUCCESS_CREATED, bankAccount, common.HTTP_CREATED)
	return
}

func (h *BankAccountHandler) GetByID(w http.ResponseWriter, r *http.Request) {
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

	bankAccount, err := h.service.GetByID(uint(id))
	if err != nil {
		common.ErrorResponse(w, http.StatusNotFound, common.HTTP_NOT_FOUND, common.ERR_NOT_FOUND, nil)
		return
	}

	common.SuccessResponse(w, common.SUCCESS_RETRIEVED, bankAccount, common.HTTP_OK)
}

func (h *BankAccountHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	bankAccounts, err := h.service.GetAll()
	if err != nil {
		common.ErrorResponse(w, http.StatusNotFound, common.HTTP_NOT_FOUND, common.ERR_NOT_FOUND, nil)
		return
	}

	common.SuccessResponse(w, common.SUCCESS_RETRIEVED, bankAccounts, common.HTTP_OK)
}

func (h *BankAccountHandler) Update(w http.ResponseWriter, r *http.Request) {
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

	var bankAccount BankAccount
	if err := json.NewDecoder(r.Body).Decode(&bankAccount); err != nil {
		common.ErrorResponse(w, http.StatusBadRequest, common.HTTP_BAD_REQUEST, common.ERR_VALIDATION, nil)
		return
	}

	bankAccount.IdCuentaBancaria = uint(id)

	updatedText, err := h.service.Update(&bankAccount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	common.SuccessResponse(w, common.SUCCESS_UPDATED, updatedText, common.HTTP_OK)
}

func (h *BankAccountHandler) Delete(w http.ResponseWriter, r *http.Request) {
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

	text, err := h.service.Delete(uint(id))
	if err != nil {
		common.ErrorResponse(w, http.StatusNotFound, common.HTTP_NOT_FOUND, common.ERR_NOT_FOUND, nil)
		return
	}

	common.SuccessResponse(w, common.SUCCESS_DELETED, text, common.HTTP_OK)
}
