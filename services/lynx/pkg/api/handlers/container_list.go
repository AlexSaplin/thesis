package handlers

import (
	uuid "github.com/satori/go.uuid"
	"net/http"

	"github.com/labstack/echo/v4"
)

type listContainersResponse struct {
	Containers []containerResponse `json:"containers"`
}

func (h *Handlers) ListContainersHandler(c echo.Context) error {
	result := listContainersResponse{
		Containers: []containerResponse{},
	}

	userID := c.Get(MwUserIDKey).(uuid.UUID)

	containers, err := h.slav.ListContainers(c.Request().Context(), userID)

	if err != nil {
		return err
	}

	for _, container := range containers {
		result.Containers = append(result.Containers, newContainerResponse(container))
	}

	return c.JSON(http.StatusOK, result)
}
