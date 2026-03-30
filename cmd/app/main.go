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

	"service/internal/handler"
	"service/internal/router"
	"service/internal/server"
)

func main() {
	// тут создаются config, db, repo, service, handler
	healthHandler := handler.NewHealthHandler()
	r := router.New(router.Dependencies{
		HealthHandler: healthHandler,
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

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("graceful shutdown failed: %v", err)
	}
}
