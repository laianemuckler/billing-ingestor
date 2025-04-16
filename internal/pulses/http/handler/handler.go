package handler

import (
	"billing-ingestor/internal/pulses/domain"
	"billing-ingestor/internal/pulses/usecase"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type PulseHandler struct {
	usecase usecase.PulseUsecase
}

func NewPulseHandler(usecase usecase.PulseUsecase) *PulseHandler {
	return &PulseHandler{usecase: usecase}
}

func (h *PulseHandler) ProcessPulse(c echo.Context) error {
	var pulseInput domain.PulseInput
	if err := c.Bind(&pulseInput); err != nil {
		log.Printf("Error binding input: %v", err)
		return c.JSON(http.StatusBadRequest, "Invalid input")
	}

	err := h.usecase.ProcessPulse(c.Request().Context(), pulseInput)
	if err != nil {
		log.Printf("Error processing pulse: %v", err)
		return c.JSON(http.StatusInternalServerError, "Failed to process pulse")
	}

	return c.JSON(http.StatusOK, "Pulse processed successfully")
}

func (h *PulseHandler) CommitAggregatedData(c echo.Context) error {
	err := h.usecase.SendAggregatedData(c.Request().Context())
	if err != nil {
		log.Printf("Error sending aggregated data: %v", err)
		return c.JSON(http.StatusInternalServerError, "Failed to send aggregated data")
	}

	err = h.usecase.ClearAggregatedData(c.Request().Context())
	if err != nil {
		log.Printf("Error clearing aggregated data: %v", err)
		return c.JSON(http.StatusInternalServerError, "Failed to clear aggregated data")
	}

	return c.JSON(http.StatusOK, "Aggregated data committed and cleared successfully")
}

func (h *PulseHandler) GetAggregatedData(c echo.Context) error {
	aggregates, err := h.usecase.GetAggregatedData(c.Request().Context())
	if err != nil {
		log.Printf("Error retrieving aggregated data: %v", err)
		return c.JSON(http.StatusInternalServerError, "Failed to retrieve aggregated data")
	}

	return c.JSON(http.StatusOK, aggregates)
}
