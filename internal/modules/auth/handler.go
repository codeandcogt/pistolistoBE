package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pistolistoBE/internal/common"
)

type AuthHandler struct {
	service AuthService
}

func NewAuthHandler(service AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}

type Credentials struct {
	Email      string `json:"email"`
	Contrasena string `json:"contrasena"`
}

func (h *AuthHandler) LoginClient(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		common.ErrorResponse(w, http.StatusBadRequest, common.HTTP_BAD_REQUEST, common.HTTP_BAD_REQUEST, nil)
		return
	}

	data, err := h.service.Login(creds.Email, creds.Contrasena)

	if err != nil {
		common.ErrorResponse(w, http.StatusUnauthorized, common.HTTP_UNAUTHORIZED, common.ERR_UNAUTHORIZED, nil)
		return
	}

	common.SuccessResponse(w, common.SUCCESS_LOGIN, data, common.HTTP_OK)
}

func (h *AuthHandler) LoginAdmin(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		common.ErrorResponse(w, http.StatusBadRequest, common.HTTP_BAD_REQUEST, common.HTTP_BAD_REQUEST, nil)
		return
	}

	data, err := h.service.LoginAdmin(creds.Email, creds.Contrasena)

	if err != nil {
		common.ErrorResponse(w, http.StatusUnauthorized, common.HTTP_UNAUTHORIZED, common.ERR_UNAUTHORIZED, nil)
		return
	}
	fmt.Println(data, "valores de data")
	common.SuccessResponse(w, common.SUCCESS_LOGIN, data, common.HTTP_OK)
}
