package handlers

import (
	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
	"io/ioutil"
	"net/http"

	"lynx/pkg/entities"
)

func (h *Handlers) PostRunFunctionHandler(c echo.Context) error {

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

	reader := c.Request().Body
	defer func() { _ = reader.Close() }()

	if err = verifyFunctionState(function.State); err != nil {
		return err
	}

	inData, err := ioutil.ReadAll(reader)
	if err != nil {
		return err
	}

	result, err := h.rhino.Run(c.Request().Context(), function, inData)
	if err != nil {
		return err
	}

	return c.Blob(http.StatusOK, "application/x-binary", result)
}

func verifyFunctionState(state entities.FunctionState) error {
	switch state {
	case entities.FunctionStateInit:
		return ErrModelStateInit
	case entities.FunctionStateProcessing:
		return ErrModelStateProcessing
	case entities.FunctionStateInvalid:
		return ErrModelStateInvalid
	case entities.FunctionStateDeleted:
		return ErrModelDeleted
	case entities.FunctionStateReady:
		return nil
	default:
		return ErrFunctionStateUnknown
	}
}
