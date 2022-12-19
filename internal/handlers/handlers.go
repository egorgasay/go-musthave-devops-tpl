package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"

	repo "devtool/internal/repository"
	"devtool/internal/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(cfg repo.Config) (*Handler, error) {
	srv, err := service.NewService(cfg)
	if err != nil {
		return nil, err
	}

	return &Handler{services: srv}, nil
}

func (h Handler) UpdateMetricHandler(c *gin.Context) {
	valStr := c.Param("value")
	val, err := strconv.ParseFloat(valStr, 64)
	if err != nil {
		c.Error(err)
		c.AbortWithStatus(400)

		return
	}

	mt := &repo.Metrics{
		Type:  c.Param("type"),
		Name:  c.Param("name"),
		Value: val,
	}

	err = h.services.DB.UpdateMetric(mt)
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
	var mt []repo.Metrics

	err := h.services.DB.GetAllMetrics(mt)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	// нужно пройтись по списку и вывести все метрики в html
}
