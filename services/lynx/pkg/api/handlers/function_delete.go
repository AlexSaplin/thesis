package handlers

import (
	"lynx/pkg/clients/ibis"
	"net/http"

	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"

	"lynx/pkg/entities"
)

func (h *Handlers) DeleteFunctionHandler(c echo.Context) error {
	userID := c.Get(MwUserIDKey).(uuid.UUID)
	modelName := c.Param("name")

	model, err := h.ibis.GetFunctionByName(c.Request().Context(), userID, modelName)
	if err != nil {
		return err
	}

	st := entities.FunctionStateDeleted
	model, err = h.ibis.UpdateFunction(c.Request().Context(), model.ID, ibis.UpdateFunctionParam{
		State: &st,
	})
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, newFunctionResponse(model))
}
