package handlers

import (
	"devtool/internal/storage"
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

func NewHandler(db *repo.Repository) *Handler {
	return &Handler{services: service.NewService(db)}
}

func (h Handler) UpdateMetricByJSONHandler(c *gin.Context) {
	var metrics storage.Metrics
	b, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(b, &metrics)
	if err != nil {
		c.AbortWithStatus(http.StatusNotImplemented)
		return
	}

	count, err := h.services.DB.UpdateMetric(&metrics)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusNotFound, err.Error())
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
		c.AbortWithStatusJSON(http.StatusNotFound, err.Error())
		return
	}

	c.Status(http.StatusOK)
	c.Writer.Write(byteJSON)
}

func (h Handler) UpdateMetricHandler(c *gin.Context) {
	valStr := c.Param("value")
	if valStr == "" {
		// c.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	val, err := strconv.ParseFloat(valStr, 64)
	if err != nil {
		c.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)

		return
	}

	metricType := c.Param("type")
	if metricType == "" || !strings.Contains("gauge,counter", metricType) {
		c.AbortWithStatus(http.StatusNotImplemented)

		return
	}
	name := c.Param("name")

	var mt *storage.Metrics

	switch metricType {
	case "gauge":
		mt = &storage.Metrics{
			ID:    name,
			MType: metricType,
			Value: &val,
		}
	case "counter":
		delta := int64(val)
		mt = &storage.Metrics{
			ID:    name,
			MType: metricType,
			Delta: &delta,
		}
	}

	_, err = h.services.DB.UpdateMetric(mt)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, err.Error())
		return
	}

	c.Status(http.StatusOK)
}

func (h Handler) GetMetricHandler(c *gin.Context) {
	val, err := h.services.DB.GetMetric(c.Param("name"))
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	c.Status(http.StatusOK)
	c.Writer.WriteString(strconv.FormatFloat(val, 'f', -1, 64))
}

func (h Handler) GetMetricByJSONHandler(c *gin.Context) {
	var metric storage.Metrics
	b, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(b, &metric)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	val, err := h.services.DB.GetMetric(metric.ID)
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	if metric.MType == "gauge" {
		metric.Value = &val
	} else if metric.MType == "counter" {
		delta := int64(val)
		metric.Delta = &delta
	} else {
		c.AbortWithStatus(http.StatusNotImplemented)
		return
	}

	outputJSON, err := json.Marshal(metric)
	if err != nil {
		c.AbortWithError(http.StatusNotImplemented, err)
		return
	}

	c.Header("Content-Type", "application/json")
	c.Status(http.StatusOK)
	c.Writer.Write(outputJSON)
}

func (h Handler) GetAllMetricsHandler(c *gin.Context) {
	mt, err := h.services.DB.GetAllMetrics()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.HTML(http.StatusOK, "index.html", gin.H{
		"Title":   "All metrics",
		"Metrics": mt,
	})
}

func (h Handler) CustomNotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{"message": "Page not found"})
}
