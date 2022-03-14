package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/satori/go.uuid"
	"net/http"
)

func (h *Handlers) PutFunctionHandler(c echo.Context) error {
	userID := c.Get(MwUserIDKey).(uuid.UUID)

	err := h.checkPositiveBalance(c, userID)
	if err != nil {
		return err
	}

	modelName := c.Param("name")

	function, err := h.ibis.CreateFunction(
		c.Request().Context(), userID, modelName,
	)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, newFunctionResponse(function))
}
