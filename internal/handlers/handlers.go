package handlers

import (
	"devtool/config"
	"devtool/internal/storage"
	"devtool/internal/usecase"
	"errors"
	"github.com/goccy/go-json"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	logic usecase.UseCase
	conf  *config.Config
}

func NewHandler(logic usecase.UseCase, conf *config.Config) *Handler {
	return &Handler{logic: logic, conf: conf}
}

func (h Handler) UpdateMetricByJSONHandler(c *gin.Context) {
	b, err := h.logic.UseGzip(c.Request.Body, c.Request.Header.Get("Content-Type"))
	if err != nil {
		c.Error(err)
		c.AbortWithStatus(http.StatusInternalServerError)

		return
	}

	var metric storage.Metrics
	err = json.Unmarshal(b, &metric)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	//log.Println(string(b))

	isValid := checkCookies(metric.Hash, metric, []byte(h.conf.Key))
	if h.conf.Key != "" {
		if !isValid {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
	}

	byteJSON, err := h.logic.UpdateMetricByJSON(b)
	if err != nil {
		if errors.Is(err, usecase.ErrNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, err.Error())
			return
		}

		c.AbortWithStatus(http.StatusNotImplemented)
		return
	}

	c.Status(http.StatusOK)
	c.Writer.Write(byteJSON)
}

func (h Handler) UpdateMetricHandler(c *gin.Context) {
	valStr := c.Param("value")
	if valStr == "" {
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

	//var metric = storage.Metrics{
	//	ID:    name,
	//	MType: metricType,
	//	Delta: int64(val),
	//	Value: val,
	//}

	err = h.logic.UpdateMetric(val, metricType, name)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, err.Error())
		return
	}

	c.Status(http.StatusOK)
}

func (h Handler) GetMetricHandler(c *gin.Context) {
	val, err := h.logic.GetMetric(c.Param("name"))
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	c.Status(http.StatusOK)
	c.Writer.WriteString(val)
}

func (h Handler) GetMetricByJSONHandler(c *gin.Context) {
	b, err := h.logic.UseGzip(c.Request.Body, c.Request.Header.Get("Content-Type"))
	if err != nil {
		c.Error(err)
		c.AbortWithStatus(http.StatusInternalServerError)

		return
	}

	outputJSON, err := h.logic.GetMetricByJSON(b)
	if err != nil {
		if errors.Is(err, usecase.ErrNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, err.Error())
			return
		}

		c.AbortWithStatus(http.StatusNotImplemented)
		return
	}

	c.Header("Content-Type", "application/json")
	c.Status(http.StatusOK)
	c.Writer.Write(outputJSON)
}

func (h Handler) GetAllMetricsHandler(c *gin.Context) {
	mt, err := h.logic.GetAllMetrics()
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
