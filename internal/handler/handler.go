package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"os"
	"strconv"

	"service/internal/service"
)

var Logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))

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
	Email     string  `josn:"email"`
}

func (h *AccountHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req createRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	acc, err := h.accountService.CreateAccount(
		r.Context(),
		req.OwnerName,
		req.Balance,
		req.Email,
		req.Currency)

	if err != nil {
		Logger.Error("create account failed",
			"err", err,
			"path", r.URL.Path,
			"method", r.Method,
		)

		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	Logger.Info("account created",
		"event", "account_created",
		"account_id", acc.ID,
		"email", acc.Email,
	)
	if err := json.NewEncoder(w).Encode(acc); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}

}

func (h *AccountHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	if idStr == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	acc, err := h.accountService.GetByID(r.Context(), id)
	if err != nil {
		Logger.Error("create account failed",
			"err", err,
			"path", r.URL.Path,
			"method", r.Method,
		)

		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	Logger.Info("account created",
		"event", "account_created",
		"account_id", acc.ID,
		"email", acc.Email,
	)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(acc); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}
}
