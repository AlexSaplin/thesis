package handlers

import (
	uuid "github.com/satori/go.uuid"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handlers) GetContainerHandler(c echo.Context) error {
	userID := c.Get(MwUserIDKey).(uuid.UUID)
	containerName := c.Param("name")

	container, err := h.slav.GetContainer(c.Request().Context(), userID, containerName)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, newContainerFullResponse(container))
}
