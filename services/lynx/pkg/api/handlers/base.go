package handlers

import (
	"lynx/pkg/clients/arietes"
	"lynx/pkg/clients/ibis"
	"lynx/pkg/clients/picus"
	"lynx/pkg/clients/prom"
	"lynx/pkg/clients/rhino"
	"regexp"

	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"

	"lynx/pkg/clients/ardea"
	"lynx/pkg/clients/auth"
	"lynx/pkg/clients/gorilla"
	"lynx/pkg/clients/hippo"
	"lynx/pkg/clients/ovis"
	"lynx/pkg/clients/s3"
	"lynx/pkg/clients/slav"
	"lynx/pkg/config"
	"lynx/pkg/entities"
)

type Handlers struct {
	users []config.UserToken

	meta      ardea.Client
	balance   gorilla.Client
	data      s3.S3Client
	runner    hippo.Client
	validator ovis.Client
	auth      auth.Client
	slav      slav.Client
	ibis      ibis.Client
	rhino     rhino.Client
	arietes   arietes.Client
	picus     picus.Client
	prom      prom.Client
}

func NewHandlers(
	users []config.UserToken, meta ardea.Client, balance gorilla.Client,
	data s3.S3Client, runner hippo.Client, validator ovis.Client,
	auth auth.Client, slav slav.Client, ibis ibis.Client, rhino rhino.Client, arietes arietes.Client, picus picus.Client, prom prom.Client,
) *Handlers {
	return &Handlers{
		users:     users,
		meta:      meta,
		balance:   balance,
		data:      data,
		runner:    runner,
		validator: validator,
		auth:      auth,
		slav:      slav,
		ibis:      ibis,
		rhino:     rhino,
		arietes:   arietes,
		picus:     picus,
		prom:      prom,
	}
}

func (h *Handlers) parseValueType(s string) (result entities.ValueType, err error) {
	switch s {
	case "UINT8":
		result = entities.ValueTypeUInt8
	case "INT8":
		result = entities.ValueTypeInt8
	case "FLOAT16":
		result = entities.ValueTypeFloat16
	case "UINT16":
		result = entities.ValueTypeUInt16
	case "INT16":
		result = entities.ValueTypeInt16
	case "FLOAT32":
		result = entities.ValueTypeFloat32
	case "UINT32":
		result = entities.ValueTypeUInt32
	case "INT32":
		result = entities.ValueTypeInt32
	case "FLOAT64":
		result = entities.ValueTypeFloat64
	case "UINT64":
		result = entities.ValueTypeUint64
	case "INT64":
		result = entities.ValueTypeInt64
	case "COMPLEX64":
		result = entities.ValueTypeComplex64
	case "COMPLEX128":
		result = entities.ValueTypeComplex128
	default:
		err = ErrInvalidTypeName
	}
	return
}

func (h *Handlers) checkImplementedValueType(valueType entities.ValueType) error {
	switch valueType {
	case entities.ValueTypeUInt8, entities.ValueTypeInt8,
		entities.ValueTypeUInt16, entities.ValueTypeInt16,
		entities.ValueTypeUInt32, entities.ValueTypeInt32, entities.ValueTypeFloat32,
		entities.ValueTypeUint64, entities.ValueTypeInt64, entities.ValueTypeFloat64:
		return nil
	default:
		return ErrUnsupportedValueType
	}
}

func (h *Handlers) checkContainerName(containerName string) error {
	re := regexp.MustCompile("[a-z][a-z0-9\\-]{2,25}")
	if re.MatchString(containerName) {
		return nil
	}
	return ErrInvalidContainerName
}

func (h *Handlers) checkImageName(imageName string) error {
	// https://github.com/docker/distribution/blob/2800ab02245e2eafc10e338939511dd1aeb5e135/reference/regexp.go#L37
	re := regexp.MustCompile("[\\w][\\w.-]{0,127}")
	if re.MatchString(imageName) {
		return nil
	}
	return ErrInvalidImageName
}

func (h *Handlers) checkInstanceType(instance entities.ContainerInstanceType) error {
	switch instance {
	case entities.ContainerInstanceTypeStarter, entities.ContainerInstanceTypeInference,
		entities.ContainerInstanceTypeInferencePRO:
		return nil
	default:
		return ErrUnsupportedInstanceType
	}
}

func (h *Handlers) checkPositiveBalance(c echo.Context, ownerID uuid.UUID) error {
	if c.Get(MWConstTokenKey).(bool) {
		return nil
	}
	balance, err := h.balance.GetBalance(c.Request().Context(), ownerID)
	if err != nil {
		return err
	}
	if balance <= 0 {
		return ErrInsufficientBalance
	}
	return nil
}

type modelResponse struct {
	Name     string `json:"name"`
	DataType string `json:"data_type"`
	State    string `json:"state"`

	Input  entities.IOShape `json:"input"`
	Output entities.IOShape `json:"output"`

	// Legacy, for compatibility with python SDK <= 0.3.1
	InputShape  entities.Shape `json:"input_shape"`
	OutputShape entities.Shape `json:"output_shape"`

	Error string `json:"error,omitempty"`
}

func newModelResponse(model entities.Model) (result modelResponse) {
	result = modelResponse{
		Name:     model.Name,
		DataType: model.ValueType.String(),
		State:    model.State.String(),
		Input:    model.InputShape,
		Output:   model.OutputShape,
		Error:    model.Error.String(),
	}
	if len(model.InputShape) > 0 {
		result.InputShape = model.InputShape[0]
	}
	if len(model.OutputShape) > 0 {
		result.OutputShape = model.OutputShape[0]
	}
	return
}

type containerFullResponse struct {
	Name         string            `json:"name"`
	Scale        uint32            `json:"scale"`
	InstanceType string            `json:"instance_type"`
	Image        string            `json:"image"`
	Port         uint32            `json:"port"`
	URL          string            `json:"url"`
	Env          map[string]string `json:"env"`
	Auth         []string          `json:"auth"`

	State string `json:"state"`
	Error string `json:"error,omitempty"`
}

func newContainerFullResponse(container entities.ContainerFull) (result containerFullResponse) {
	result = containerFullResponse{
		Name:         container.Name,
		State:        container.State.String(),
		Scale:        container.Scale,
		InstanceType: container.InstanceType.String(),
		Image:        container.Image,
		Port:         container.Port,
		URL:          container.URL,
		Error:        container.Error.String(),
		Env:          container.Env,
		Auth:         container.Auth,
	}
	return
}

type containerResponse struct {
	Name         string            `json:"name"`
	Scale        uint32            `json:"scale"`
	InstanceType string            `json:"instance_type"`
	Image        string            `json:"image"`
	Port         uint32            `json:"port"`
	URL          string            `json:"url"`
	Env          map[string]string `json:"env"`
	Auth         []string          `json:"auth"`
}

func newContainerResponse(container entities.Container) (result containerResponse) {
	result = containerResponse{
		Name:         container.Name,
		Scale:        container.Scale,
		InstanceType: container.InstanceType.String(),
		Image:        container.Image,
		Port:         container.Port,
		URL:          container.URL,
		Env:          container.Env,
		Auth:         container.Auth,
	}
	return
}

type functionResponse struct {
	Id    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	State string    `json:"state"`

	Error string    `json:"error,omitempty"`
}

func newFunctionResponse(fn entities.Function) (result functionResponse) {
	result = functionResponse{
		Id:    fn.ID,
		Name:  fn.Name,
		State: fn.State.String(),
		Error: fn.Error.String(),
	}
	return
}
