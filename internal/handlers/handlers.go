package handlers

import (
	"devtool/config"
	"devtool/internal/usecase"
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	logic usecase.UseCase
	conf  config.Config
}

func NewHandler(logic usecase.UseCase) *Handler {
	return &Handler{logic: logic}
}

func (h Handler) UpdateMetricByJSONHandler(c *gin.Context) {
	cookie, err := getCookies(c)
	if len(h.conf.Key) > 1 {
		if !checkCookies(cookie, h.conf.Key) {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
	}

	if err != nil || !checkCookies(cookie, h.conf.Key) {
		log.Println("Setting new cookies...")
		setCookies(c, h.conf.Key)
	}

	b, err := h.logic.UseGzip(c.Request.Body, c.Request.Header.Get("Content-Type"))
	if err != nil {
		c.Error(err)
		c.AbortWithStatus(http.StatusInternalServerError)

		return
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
	cookie, err := getCookies(c)
	if len(h.conf.Key) > 1 {
		if !checkCookies(cookie, h.conf.Key) {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
	}

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
