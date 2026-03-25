package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"rest-notes-api/internal/config"
	"rest-notes-api/internal/httpapi"
	"rest-notes-api/internal/repo"
	"rest-notes-api/internal/service"
	"syscall"
	"time"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	pqRepo, err := repo.NewPostgresRepo(cfg.DbDSN)
	if err != nil {
		log.Fatal(err)
	}
	defer pqRepo.Close()

	if err := pqRepo.Init(context.Background()); err != nil {
		log.Fatal(err)
	}

	svc := service.NewNoteService(pqRepo)

	handler := httpapi.NewHandler(svc)

	srv := &http.Server{
		Addr:              cfg.ServerPort,
		Handler:           handler,
		ReadHeaderTimeout: 5 * time.Second,
	}

	go func() {
		log.Printf("server started on %s", cfg.ServerPort)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quik := make(chan os.Signal, 1)
	signal.Notify(quik, syscall.SIGINT, syscall.SIGTERM)

	<-quik
	log.Println("shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("server forced to shutdown:", err)
	}

	log.Println("server exited properly")
}
