package main

import (
	"devtool/config"
	"devtool/internal/handlers"
	repo "devtool/internal/repository"
	"devtool/internal/routes"
	store "devtool/internal/storage"
	"flag"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"time"
)

//var host = "localhost:8080"

var (
	host      *string
	path      *string
	saveAfter *string
	restore   *bool
)

func init() {
	host = flag.String("a", "localhost:8080", "-a=host")
	saveAfter = flag.String("i", "5m", "-i=Seconds")
	path = flag.String("f", "/tmp/devops-metrics-db.json", "-f=path")
	restore = flag.Bool("r", true, "-r=restore")
}

func main() {
	flag.Parse()
	r := gin.Default()
	cfg := config.New(*saveAfter, *restore, *path)

	if addr, ok := os.LookupEnv("ADDRESS"); ok {
		host = &addr
	}

	storage, err := repo.New(cfg.DBConfig)
	// add context
	go func(storage *repo.Repository, saveAfter string) {
		storeInterval, err := time.ParseDuration(saveAfter)
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
	}(storage, *saveAfter)

	if err != nil {
		log.Fatalf("Failed to initialize: %s", err.Error())
	}

	h := handlers.NewHandler(storage)

	public := r.Group("/")
	routes.PublicRoutes(public, *h)
	r.LoadHTMLGlob("templates/*")

	r.Run(*host)
}
