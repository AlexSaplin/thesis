package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"

	"lynx/pkg/entities"
)


type postContainerAuth struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Registry string `json:"registry"`
}

type postContainerRequest struct {
	Scale        uint32              `json:"scale"`
	InstanceType string              `json:"instance_type"`
	Image        string              `json:"image"`
	Port         uint32              `json:"port"`
	Env          map[string]string   `json:"env"`
	Auth         []postContainerAuth `json:"auth"`
}

type postContainerResponse struct {
	Container containerResponse `json:"container"`
}

func (h *Handlers) PostContainerHandler(c echo.Context) error {
	result := postContainerResponse{}

	userID := c.Get(MwUserIDKey).(uuid.UUID)
	containerName := c.Param("name")
	if err := h.checkContainerName(containerName); err != nil {
		return err
	}

	req := postContainerRequest{}
	if err := c.Bind(&req); err != nil {
		return err
	}

	if req.Scale < 1 || req.Scale > 10 {
		return ErrInvalidContainerScale
	}

	instanceType := entities.ContainerInstanceType(req.InstanceType)
	err := h.checkInstanceType(instanceType)
	if err != nil {
		return err
	}

	err = h.checkImageName(req.Image)
	if err != nil {
		return err
	}

	container, err := h.slav.CreateContainer(c.Request().Context(), userID, containerName, req.Scale,
		instanceType, req.Image, req.Port, req.Env, h.parseAuth(req.Auth))

	if err != nil {
		return err
	}

	result.Container = newContainerResponse(container)

	return c.JSON(http.StatusOK, result)
}

func (h *Handlers) parseAuth(in []postContainerAuth) (out []entities.ContainerAuth) {
	for _, item := range in {
		out = append(out, entities.ContainerAuth{
			Username: item.Username,
			Password: item.Password,
			Registry: item.Registry,
		})
	}
	return
}
