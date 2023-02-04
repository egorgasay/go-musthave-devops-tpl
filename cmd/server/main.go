package main

import (
	"context"
	"devtool/config"
	"devtool/internal/handlers"
	repo "devtool/internal/repository"
	"devtool/internal/routes"
	"devtool/internal/service"
	store "devtool/internal/storage"
	"devtool/internal/usecase"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	r := gin.Default()

	cfg := config.New()

	srv := &http.Server{
		Addr:    cfg.Host,
		Handler: r,
	}

	storage, err := repo.New(cfg.DBConfig)
	// add context
	go func(storage *repo.Repository, storeInterval time.Duration) {
		if err != nil {
			panic(err)
		}
		for {
			time.Sleep(storeInterval)
			store.StorageRelevance.Mu.RLock()
			if !store.StorageRelevance.Status {
				err = storage.BackupStorage.DoBackup(storage.Store)
				if err != nil {
					log.Println(err)
					return
				}
				store.StorageRelevance.Status = true
			}
			store.StorageRelevance.Mu.RUnlock()
		}
	}(storage, cfg.DBConfig.SaveAfter)

	if err != nil {
		log.Fatalf("Failed to initialize: %s", err.Error())
	}

	Service := service.NewService(storage)
	logic := usecase.New(Service)
	h := handlers.NewHandler(logic, cfg)

	r.Use(gzip.Gzip(gzip.BestSpeed))

	public := r.Group("/")
	routes.PublicRoutes(public, *h)
	r.LoadHTMLGlob("templates/*")

	log.Println(cfg.Host)
	log.Println(cfg.DBConfig)

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	log.Println("Shutdown Server ...")

	shutdown := make(chan struct{}, 1)

	ctx, cancel := context.WithTimeout(context.Background(), 450*time.Millisecond)
	defer cancel()

	go func() {
		store.StorageRelevance.Mu.RLock()
		if !store.StorageRelevance.Status {
			store.StorageRelevance.Mu.RUnlock()

			err := storage.BackupStorage.DoBackup(storage.Store)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			store.StorageRelevance.Mu.RUnlock()
		}

		shutdown <- struct{}{}
	}()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	select {
	case <-ctx.Done():
		log.Println("Timeout exited")
	case <-shutdown:
		log.Println("Finished")
	}

	log.Println("Server exiting")
}
