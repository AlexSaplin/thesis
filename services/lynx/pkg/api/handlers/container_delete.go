package handlers

import (
	uuid "github.com/satori/go.uuid"
	"net/http"

	"github.com/labstack/echo/v4"
)

type deleteContainerResponse struct {
}

func (h *Handlers) DeleteContainerHandler(c echo.Context) error {
	result := deleteContainerResponse{}

	userID := c.Get(MwUserIDKey).(uuid.UUID)
	containerName := c.Param("name")

	err := h.slav.DeleteContainer(c.Request().Context(), userID, containerName)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, result)
}
