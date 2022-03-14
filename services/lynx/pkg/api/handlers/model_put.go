package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/satori/go.uuid"

	"lynx/pkg/entities"
)

type putModelRequest struct {
	Input  entities.IOShape `json:"input"`
	Output entities.IOShape `json:"output"`

	// Legacy, for compatibility with python SDK <= 0.3.1
	InputShape  entities.Shape `json:"input_shape"`
	OutputShape entities.Shape `json:"output_shape"`

	DataType string `json:"data_type"`
}

func (h *Handlers) PutModelHandler(c echo.Context) error {
	userID := c.Get(MwUserIDKey).(uuid.UUID)

	err := h.checkPositiveBalance(c, userID)
	if err != nil {
		return err
	}

	req := putModelRequest{}
	if err := c.Bind(&req); err != nil {
		return err
	}

	modelName := c.Param("name")

	valueType, err := h.parseValueType(req.DataType)
	if err != nil {
		return err
	}
	if err = h.checkImplementedValueType(valueType); err != nil {
		return err
	}

	var (
		inputShape  entities.IOShape
		outputShape entities.IOShape
	)

	if len(req.InputShape) > 0 { // Legacy check
		inputShape = entities.IOShape{req.InputShape}
		outputShape = entities.IOShape{req.OutputShape}
	} else {
		inputShape = req.Input
		outputShape = req.Output
	}

	if err = inputShape.Valid(); err != nil {
		return err
	}

	if err = outputShape.Valid(); err != nil {
		return err
	}

	// create model
	model, err := h.meta.CreateModel(
		c.Request().Context(), userID, modelName, inputShape, outputShape, valueType,
	)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, newModelResponse(model))
}
