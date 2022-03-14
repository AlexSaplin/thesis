package handlers

import (
	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
	"net/http"

	"lynx/pkg/entities"
)

type patchContainerRequest struct {
	Scale        *uint32 `json:"scale,omitempty"`
	InstanceType *string `json:"instance_type,omitempty"`
	Image        *string `json:"image,omitempty"`
}

type patchContainerResponse struct {
	Container containerResponse `json:"container"`
}

func (h *Handlers) PatchContainerRequest(c echo.Context) error {
	result := patchContainerResponse{}

	userID := c.Get(MwUserIDKey).(uuid.UUID)
	containerName := c.Param("name")

	req := patchContainerRequest{}
	if err := c.Bind(&req); err != nil {
		return err
	}

	var scale entities.NullUInt32
	if req.Scale == nil {
		scale = entities.NullUInt32{
			Value:   0,
			IsValid: false,
		}
	} else {
		scale = entities.NullUInt32{
			Value:   *req.Scale,
			IsValid: true,
		}
	}

	var instanceType entities.NullInstance
	if req.InstanceType == nil {
		instanceType = entities.NullInstance{
			Value:   entities.ContainerInstanceTypeStarter,
			IsValid: true,
		}
	} else {
		instanceType = entities.NullInstance{
			Value:   entities.ContainerInstanceType(*req.InstanceType),
			IsValid: true,
		}
	}

	var image entities.NullString
	if req.Image == nil {
		image = entities.NullString{
			Value:   "",
			IsValid: false,
		}
	} else {
		image = entities.NullString{
			Value:   *req.Image,
			IsValid: true,
		}
	}

	if scale.IsValid && scale.Value < 1 || scale.Value > 10 {
		return ErrInvalidContainerScale
	}

	if instanceType.IsValid {
		instanceType := instanceType.Value
		err := h.checkInstanceType(instanceType)
		if err != nil {
			return err
		}
	}

	if image.IsValid {
		err := h.checkImageName(image.Value)
		if err != nil {
			return err
		}
	}

	container, err := h.slav.UpdateContainer(c.Request().Context(), userID, containerName, scale, instanceType, image)
	if err != nil {
		return err
	}

	result.Container = newContainerResponse(container)

	return c.JSON(http.StatusOK, result)
}
