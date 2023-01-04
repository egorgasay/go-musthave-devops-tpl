package main

import (
	"devtool/config"
	"devtool/internal/handlers"
	repo "devtool/internal/repository"
	"devtool/internal/routes"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

var host = "localhost:8080"

func main() {
	r := gin.Default()

	cfg := config.New()

	if addr, ok := os.LookupEnv("ADDRESS"); ok {
		host = addr
	}

	storage, err := repo.New(cfg.DBConfig)

	if err != nil {
		log.Fatalf("Failed to initialize: %s", err.Error())
	}

	h := handlers.NewHandler(storage)

	public := r.Group("/")
	routes.PublicRoutes(public, *h)
	r.LoadHTMLGlob("templates/*")

	r.Run(host)
}
