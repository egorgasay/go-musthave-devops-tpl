package handlers

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	repo "devtool/internal/repository"
	"devtool/internal/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(db *repo.MemStorage) *Handler {
	return &Handler{services: service.NewService(db)}
}

func (h Handler) UpdateMetricHandler(c *gin.Context) {
	valStr := c.Param("value")
	if valStr == "" {
		// c.Error(err)
		c.AbortWithStatus(400)
		return
	}

	val, err := strconv.ParseFloat(valStr, 64)
	if err != nil {
		c.Error(err)
		c.AbortWithStatus(400)

		return
	}

	metricType := c.Param("type")
	if metricType == "" || !strings.Contains("gauge,counter", metricType) {
		// c.Error(err)
		c.AbortWithStatus(501)

		return
	}

	mt := &repo.Metrics{
		Type:  metricType,
		Name:  c.Param("name"),
		Value: val,
	}

	err = h.services.UpdateMetric.UpdateMetric(mt)
	if err != nil {
		c.AbortWithStatusJSON(404, err.Error())
		return
	}

	c.Status(200)
}

func (h Handler) GetMetricHandler(c *gin.Context) {
	val, err := h.services.GetMetric.GetMetric(c.Param("name"))
	if err != nil {
		c.AbortWithError(404, err)
		return
	}

	c.Status(200)
	c.Writer.WriteString(strconv.FormatFloat(val, 'f', -1, 64))
}

func (h Handler) GetAllMetricsHandler(c *gin.Context) {
	var mt []repo.Metrics

	err := h.services.GetAllMetrics.GetAllMetrics(mt)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	// нужно пройтись по списку и вывести все метрики в html
}

func (h Handler) CustomNotFound(c *gin.Context) {
	c.JSON(404, gin.H{"message": "Page not found"})
	return
}
