package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/example/stratus-mesh/internal/handlers"
	"github.com/example/stratus-mesh/internal/services"
	"github.com/example/stratus-mesh/internal/store"
	"github.com/example/stratus-mesh/internal/router"
	"github.com/example/stratus-mesh/pkg/config"
)

func main() {
	cfg := config.Load()

	db, err := store.InitDB(cfg.DatabaseDSN)
	if err != nil {
		log.Fatalf("failed to init db: %v", err)
	}
	defer store.CloseDB(db)

	// Initialize services
	discovery := services.NewDiscovery()
	replicator := services.NewReplicator(db, discovery)

	// run replicator background job
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go replicator.Run(ctx)

	h := handlers.NewHandler(db, discovery)
	r := router.NewRouter(h)

	// graceful shutdown
	go func() {
		if err := r.Run(":" + cfg.Port); err != nil {
			log.Fatalf("server error: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("shutting down...")
	ctxShutdown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_ = ctxShutdown
}
