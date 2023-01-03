package routes

import (
	"devtool/internal/handlers"
	"github.com/gin-gonic/gin"
)

func PublicRoutes(r *gin.RouterGroup, h handlers.Handler) {
	r.POST("/update/", h.UpdateMetricByJSONHandler)
	r.POST("/update/:type/:name/:value", h.UpdateMetricHandler)
	r.GET("/value/:type/:name", h.GetMetricHandler)
	r.POST("/value/", h.GetMetricByJSONHandler)
	r.GET("/", h.GetAllMetricsHandler)
	r.POST("/update/:type/", h.CustomNotFound)
}
