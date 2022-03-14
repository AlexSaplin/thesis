package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
)

type ListFunctionssResponse struct {
	Functions []functionResponse `json:"functions"`
}

func (h *Handlers) ListFunctionsHandler(c echo.Context) error {
	userID := c.Get(MwUserIDKey).(uuid.UUID)

	fns, err := h.ibis.ListFunctions(c.Request().Context(), userID)
	if err != nil {
		return err
	}

	result := make([]functionResponse, 0, len(fns))

	for _, fn := range fns {
		result = append(result, newFunctionResponse(fn))
	}

	return c.JSON(http.StatusOK, ListFunctionssResponse{
		Functions: result,
	})
}
