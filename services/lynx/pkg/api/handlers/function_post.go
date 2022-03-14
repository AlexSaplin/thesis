package handlers

import (
	"github.com/mattn/go-nulltype"
	"lynx/pkg/clients/ibis"
	"net/http"

	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"

	"lynx/pkg/entities"
)

func (h *Handlers) PostFunctionHandler(c echo.Context) error {
	userID := c.Get(MwUserIDKey).(uuid.UUID)

	err := h.checkPositiveBalance(c, userID)
	if err != nil {
		return err
	}

	modelName := c.Param("name")

	function, err := h.ibis.GetFunctionByName(c.Request().Context(), userID, modelName)
	if err != nil {
		return err
	}

	file, err := c.FormFile("repo")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	path, err := h.data.UploadFunctionData(function, src)
	if err != nil {
		return err
	}
	state := entities.FunctionStateProcessing
	updateParam := ibis.UpdateFunctionParam{
		State:    &state,
		CodePath: nulltype.NullStringOf(path),
	}

	function, err = h.ibis.UpdateFunction(c.Request().Context(), function.ID, updateParam)
	if err != nil {
		return err
	}

	functionQuery := entities.FunctionQuery{
		ID: function.ID,
	}

	if err = h.arietes.SubmitFunctionQuery(functionQuery); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, newFunctionResponse(function))
}
