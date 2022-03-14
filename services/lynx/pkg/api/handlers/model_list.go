package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
)

type ListModelsResponse struct {
	Models []modelResponse `json:"models"`
}

func (h *Handlers) ListModelsHandler(c echo.Context) error {
	userID := c.Get(MwUserIDKey).(uuid.UUID)

	models, err := h.meta.ListModels(c.Request().Context(), userID)
	if err != nil {
		return err
	}

	result := make([]modelResponse, 0, len(models))

	for _, model := range models {
		result = append(result, newModelResponse(model))
	}

	return c.JSON(http.StatusOK, ListModelsResponse{
		Models: result,
	})
}
