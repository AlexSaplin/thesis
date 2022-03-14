package handlers

import (
	"bufio"
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"

	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"

	"lynx/pkg/entities"
)

type ShapePart struct {
	Shape [][]int64 `json:"shape"`
}

func (h *Handlers) PostRunModelHandler(c echo.Context) error {

	userID := c.Get(MwUserIDKey).(uuid.UUID)

	err := h.checkPositiveBalance(c, userID)
	if err != nil {
		return err
	}

	modelName := c.Param("name")

	input, err := h.postRunModelParams(c)
	if err != nil {
		return err
	}
	if len(input) == 0 {
		return ErrEmptyInput
	}

	model, err := h.meta.GetModelByName(c.Request().Context(), userID, modelName)
	if err != nil {
		return err
	}

	if err = verifyModelState(model.State); err != nil {
		return err
	}

	if !input.ConformsToShape(model.InputShape) {
		return ErrShapeMismatch
	}

	if input[0].Type != model.ValueType {
		return ErrTypeMismatch
	}

	reader := c.Request().Body
	defer func() { _ = reader.Close() }()

	result, shape, err := h.runner.Run(c.Request().Context(), model, input)
	if err != nil {
		return err
	}

	var buffer bytes.Buffer
	mw := multipart.NewWriter(&buffer)

	var contentType string
	contentType = mw.FormDataContentType()

	resultField, err := mw.CreateFormField("result")
	if err != nil {
		return ErrMultipartBuilding
	}
	if _, err := resultField.Write(result); err != nil {
		return ErrMultipartBuilding
	}

	shapeField, err := mw.CreateFormField("shape")
	if err != nil {
		return ErrMultipartBuilding
	}

	shapePart := ShapePart{Shape: shape}

	shapeBytes, err := json.Marshal(shapePart)

	if err != nil {
		return ErrInvalidShapeFormat
	}

	if _, err := shapeField.Write(shapeBytes); err != nil {
		return ErrMultipartBuilding
	}

	if err := mw.Close(); err != nil {
		return ErrMultipartBuilding
	}

	return c.Blob(http.StatusOK, contentType, buffer.Bytes())
}

func (h *Handlers) postRunModelParams(c echo.Context) (tensors entities.TensorList, err error) {
	reqShape := c.FormValue("shape")
	var shape entities.IOShape

	if err = json.Unmarshal([]byte(reqShape), &shape); err != nil {
		var shapeOld entities.Shape
		if err = json.Unmarshal([]byte(reqShape), &shapeOld); err != nil {
			err = ErrInvalidShapeFormat
			return
		}
		shape = entities.IOShape{shapeOld}
	}
	if err = shape.Valid(); err != nil {
		return
	}

	reqType := c.FormValue("data_type")
	parsedType, err := h.parseValueType(reqType)
	if err != nil {
		return
	}

	reqTensor, err := c.FormFile("tensor")
	if err != nil {
		return
	}
	tensorReader, err := reqTensor.Open()
	if err != nil {
		return
	}
	defer tensorReader.Close()
	bufReader := bufio.NewReaderSize(tensorReader, 1024)

	tensors = make(entities.TensorList, 0, len(shape))

	for _, shapeItem := range shape {
		var expectedTensorSize int64
		expectedTensorSize, err = parsedType.Size()
		if err != nil {
			return
		}

		var currentShape entities.TensorShape
		for _, elem := range shapeItem {
			currentShape = append(currentShape, elem.Int64Value())
		}
		expectedTensorSize *= currentShape.Product()

		buf := make([]byte, expectedTensorSize)
		_, err = io.ReadFull(bufReader, buf)
		if err != nil {
			return
		}
		tmpResult := entities.Tensor{
			Type:  parsedType,
			Shape: currentShape,
			Data:  buf,
		}
		if err = tmpResult.Valid(); err != nil {
			return
		}
		tensors = append(tensors, tmpResult)
	}

	return
}

func verifyModelState(state entities.ModelState) error {
	switch state {
	case entities.ModelStateInit:
		return ErrModelStateInit
	case entities.ModelStateProcessing:
		return ErrModelStateProcessing
	case entities.ModelStateInvalid:
		return ErrModelStateInvalid
	case entities.ModelStateDeleted:
		return ErrModelDeleted
	case entities.ModelStateReady:
		return nil
	default:
		return ErrModelStateUnknown
	}
}
