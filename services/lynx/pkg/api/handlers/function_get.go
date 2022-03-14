package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
)

func (h *Handlers) GetFunctionHandler(c echo.Context) error {
	userID := c.Get(MwUserIDKey).(uuid.UUID)
	modelName := c.Param("name")

	fn, err := h.ibis.GetFunctionByName(c.Request().Context(), userID, modelName)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, newFunctionResponse(fn))
}
