package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
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

func (h Handler) UpdateMetricByJSONHandler(c *gin.Context) {
	var metrics repo.Metrics
	b, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.AbortWithStatus(400)
		return
	}

	err = json.Unmarshal(b, &metrics)
	if err != nil {
		c.AbortWithStatus(501)
		return
	}

	count, err := h.services.DB.UpdateMetric(&metrics)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(404, err.Error())
		return
	}

	if metrics.MType == "gauge" {
		metrics.Value = &count
	} else {
		delta := int64(count)
		metrics.Delta = &delta
	}

	byteJSON, err := json.MarshalIndent(metrics, "", "    ")
	if err != nil {
		c.AbortWithStatusJSON(404, err.Error())
		return
	}

	c.Status(200)
	c.Writer.Write(byteJSON)
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
		c.AbortWithStatus(501)

		return
	}
	name := c.Param("name")

	var mt *repo.Metrics

	switch metricType {
	case "gauge":
		mt = &repo.Metrics{
			ID:    name,
			MType: metricType,
			Value: &val,
		}
	case "counter":
		delta := int64(val)
		mt = &repo.Metrics{
			ID:    name,
			MType: metricType,
			Delta: &delta,
		}
	}

	_, err = h.services.DB.UpdateMetric(mt)
	if err != nil {
		c.AbortWithStatusJSON(404, err.Error())
		return
	}

	c.Status(200)
}

func (h Handler) GetMetricHandler(c *gin.Context) {
	val, err := h.services.DB.GetMetric(c.Param("name"))
	if err != nil {
		c.AbortWithError(404, err)
		return
	}

	c.Status(200)
	c.Writer.WriteString(strconv.FormatFloat(val, 'f', -1, 64))
}

func (h Handler) GetAllMetricsHandler(c *gin.Context) {
	mt, err := h.services.DB.GetAllMetrics()
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	c.HTML(http.StatusOK, "index.html", gin.H{
		"Title":   "All metrics",
		"Metrics": mt,
	})
}

func (h Handler) CustomNotFound(c *gin.Context) {
	c.JSON(404, gin.H{"message": "Page not found"})
}
