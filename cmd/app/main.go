package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"service/db"
	"service/internal/config"
	"service/internal/handler"
	"service/internal/repository"
	"service/internal/router"
	"service/internal/server"
	"service/internal/service"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}
	pool, err := db.NewPool(ctx, cfg.ConnString())
	if err != nil {
		log.Fatal(err)
	}
	acc := repository.NewAccountPostgres(pool)
	service := service.NewAccountService(acc)
	handlerAccount := handler.NewAccountHandler(service)
	healthHandler := handler.NewHealthHandler()

	r := router.New(router.Dependencies{
		HealthHandler:  healthHandler,
		AccountHandler: handlerAccount,
	})

	srv := server.New(":8080", r)

	go func() {
		log.Println("server started on :8080")
		if err := srv.Run(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("server failed: %v", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	log.Println("shutting down server...")

	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("graceful shutdown failed: %v", err)
	}
}
