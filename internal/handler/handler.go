package handler

import (
	"encoding/json"
	"net/http"

	"service/internal/service"
)

type AccountHandler struct {
	accountService *service.AccountService
}

func NewAccountHandler(s *service.AccountService) *AccountHandler {
	return &AccountHandler{accountService: s}
}

type createRequest struct {
	OwnerName string  `json:"owner_name"`
	Balance   float64 `json:"balance"`
	Currency  string  `json:"currency"`
}

func (h *AccountHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req createRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	acc, err := h.accountService.CreateAccount(r.Context(), req.OwnerName, req.Balance, req.Currency)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(acc)

}
