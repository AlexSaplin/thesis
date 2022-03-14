package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"

	"lynx/pkg/entities"
)

func (h *Handlers) DeleteModelHandler(c echo.Context) error {
	userID := c.Get(MwUserIDKey).(uuid.UUID)
	modelName := c.Param("name")

	model, err := h.meta.GetModelByName(c.Request().Context(), userID, modelName)
	if err != nil {
		return err
	}

	model, err = h.meta.UpdateModelState(c.Request().Context(), model.ID, entities.ModelStateDeleted)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, newModelResponse(model))
}
