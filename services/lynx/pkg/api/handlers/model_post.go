package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"

	"lynx/pkg/entities"
)

func (h *Handlers) PostModelHandler(c echo.Context) error {
	userID := c.Get(MwUserIDKey).(uuid.UUID)

	err := h.checkPositiveBalance(c, userID)
	if err != nil {
		return err
	}

	modelName := c.Param("name")

	model, err := h.meta.GetModelByName(c.Request().Context(), userID, modelName)
	if err != nil {
		return err
	}

	if model.Path != "" {
		return ErrModelAlreadyUploaded
	}

	reader := c.Request().Body
	defer func() { _ = reader.Close() }()

	path, err := h.data.UploadModelData(model, reader)
	if err != nil {
		return err
	}

	model, err = h.meta.UpdateModelPath(c.Request().Context(), model.ID, path)
	if err != nil {
		return nil
	}

	model, err = h.meta.UpdateModelState(c.Request().Context(), model.ID, entities.ModelStateProcessing)
	if err != nil {
		return err
	}

	if err := h.validator.VerifyModel(model); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, newModelResponse(model))
}
