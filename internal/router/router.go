package router

import (
	"net/http"

	"service/internal/handler"
)

type Dependencies struct {
	HealthHandler  *handler.HealthHandler
	AccountHandler *handler.AccountHandler
	//ProductHandler *handler.ProductHandler
	//UserHandler    *handler.UserHandler
}

func New(deps Dependencies) http.Handler {
	mux := http.NewServeMux()

	// --- Health ---
	mux.HandleFunc("GET /health", deps.HealthHandler.Check)
	mux.HandleFunc("POST /accounts", deps.AccountHandler.Create)
	mux.HandleFunc("GET /accounts/{id}", deps.AccountHandler.GetByID)

	//// --- Users ---
	//mux.HandleFunc("GET /users", deps.UserHandler.List)
	//mux.HandleFunc("GET /users/{id}", deps.UserHandler.Get)
	//mux.HandleFunc("POST /users", deps.UserHandler.Create)
	//
	//// --- Products ---
	//mux.HandleFunc("GET /products", deps.ProductHandler.List)
	//mux.HandleFunc("GET /products/{id}", deps.ProductHandler.Get)
	//mux.HandleFunc("POST /products", deps.ProductHandler.Create)

	return mux
}
