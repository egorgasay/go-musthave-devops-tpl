package main

import (
	"devtool/config"
	"devtool/internal/handlers"
	repo "devtool/internal/repository"
	"devtool/internal/routes"
	"flag"
	"github.com/gin-gonic/gin"
	"log"
	"os"
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
	saveAfter = flag.String("i", "300", "-i=URL")
	path = flag.String("f", "/tmp/devops-metrics-db.json", "-f=path")
	restore = flag.Bool("r", true, "-r=restore")
}

func main() {
	r := gin.Default()

	cfg := config.New(*saveAfter, *restore, *path)

	if addr, ok := os.LookupEnv("ADDRESS"); ok {
		host = &addr
	}

	storage, err := repo.New(cfg.DBConfig)

	if err != nil {
		log.Fatalf("Failed to initialize: %s", err.Error())
	}

	h := handlers.NewHandler(storage)

	public := r.Group("/")
	routes.PublicRoutes(public, *h)
	r.LoadHTMLGlob("templates/*")

	r.Run(*host)
}
