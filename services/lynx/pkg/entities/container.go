package entities

import (
	"github.com/mattn/go-nulltype"
)

type ContainerInstanceType string

const (
	ContainerInstanceTypeStarter      ContainerInstanceType = "starter"
	ContainerInstanceTypeInference                          = "inference"
	ContainerInstanceTypeInferencePRO                       = "inference_pro"
	ContainerInstanceTypeUnknown                            = "unknown"
)

func (t ContainerInstanceType) String() string {
	return string(t)
}

type ContainerState uint8

const (
	ContainerStateRunning ContainerState = iota
	ContainerStateUpdating
	ContainerStateError
	ContainerStateUnknown
)

func (s ContainerState) String() string {
	switch s {
	case ContainerStateRunning:
		return "RUNNING"
	case ContainerStateUpdating:
		return "UPDATING"
	case ContainerStateError:
		return "ERROR"
	default:
		return "UNKNOWN"
	}
}

type Container struct {
	Name         string
	Scale        uint32
	InstanceType ContainerInstanceType
	Image        string
	Port         uint32
	URL          string
	Env          map[string]string
	Auth         []string
}

type ContainerFull struct {
	Name         string
	Scale        uint32
	InstanceType ContainerInstanceType
	Image        string
	Port         uint32
	URL          string
	Env          map[string]string
	Auth         []string

	State ContainerState
	Error nulltype.NullString
}

type ContainerAuth struct {
	Username string
	Password string
	Registry string
}

type NullString struct {
	Value   string
	IsValid bool
}

type NullUInt32 struct {
	Value   uint32
	IsValid bool
}

type NullInstance struct {
	Value   ContainerInstanceType
	IsValid bool
}
