package routes

import (
	"billing-ingestor/internal/pulses/http/handler"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo, h *handler.PulseHandler) {
	e.POST("/pulses", h.ProcessPulse)
	e.GET("/aggregates", h.GetAggregatedData)
	e.POST("/aggregates/commit", h.CommitAggregatedData)
}
