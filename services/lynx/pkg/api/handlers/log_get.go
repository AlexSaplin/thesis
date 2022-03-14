package handlers

import (
	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
	"lynx/pkg/entities"
	"net/http"
)

func (h *Handlers) GetLogsHandler(c echo.Context) error {
	userID := c.Get(MwUserIDKey).(uuid.UUID)
	functionName := c.Param("name")

	function, err := h.ibis.GetFunctionByName(c.Request().Context(), userID, functionName)

	if err != nil {
		return err
	}
	logs, err := h.picus.GetFunctionLogs(c.Request().Context(), entities.FunctionQuery{ID: function.ID})
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, logs)
}
