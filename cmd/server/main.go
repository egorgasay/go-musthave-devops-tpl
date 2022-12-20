package main

import (
	"devtool/config"
	"devtool/internal/handlers"
	repo "devtool/internal/repository"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	
	cfg := config.New()
	strg, err := repo.NewMemStorage(cfg.DBConfig)

	if err != nil {
		log.Fatalf("Failed to initialize: %s", err.Error())
	}

	h := handlers.NewHandler(strg.DB)

	r.POST("/update/:type/:name/:value", h.UpdateMetricHandler)
	r.GET("/value/:type/:name", h.GetMetricHandler)
	r.GET("/", h.GetAllMetricsHandler)

	log.Fatal(http.ListenAndServe(":8080", r))
}
