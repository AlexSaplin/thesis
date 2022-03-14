package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type funcEnvItem struct{
	Name     string `json:"name"`
	Language string `json:"language"`
}

func (h *Handlers) GetFunctionEnvHandler(c echo.Context) error {
	res := struct {
		Envs []funcEnvItem `json:"envs"`
	}{
		Envs: []funcEnvItem{
			{
				Name:     "python3.6",
				Language: "python",
			},
            {
				Name:     "python3.7",
				Language: "python",
			},
			{
				Name:     "python3.6-tensorflow2.1-pytorch-1.6-cuda10.1",
				Language: "python",
			},
			{
				Name:     "python3.7-tensorflow2.1-pytorch-1.6-cuda10.1",
				Language: "python",
			},
			{
			    Name:     "python3.7-tensorflow2.2-pytorch-1.6-cuda10.1",
			    Language: "python",
			},
			{
			    Name:     "python3.7-tensorflow1.13.1-pytorch-1.3-cuda10.0",
			    Language: "python",
			},
			{
			    Name:     "python3.7-mmdetection-pytorch-1.6-cuda10.1",
			    Language: "python",
			},
		},
	}

	return c.JSON(http.StatusOK, res)
}
