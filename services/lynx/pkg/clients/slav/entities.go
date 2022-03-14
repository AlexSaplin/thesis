package slav

import (
	slav "lynx/pkg/clients/slav/pb"
	"lynx/pkg/entities"
)

func parseInstanceTypeEntity(in slav.InstanceType) (result entities.ContainerInstanceType) {
	switch in {
	case slav.InstanceType_STARTER:
		return entities.ContainerInstanceTypeStarter
	case slav.InstanceType_INFERENCE:
		return entities.ContainerInstanceTypeInference
	default:
		return entities.ContainerInstanceTypeUnknown
	}
}

func serializeInstanceTypeEntity(in entities.ContainerInstanceType) (result slav.InstanceType) {
	switch in {
	case entities.ContainerInstanceTypeInference:
		return slav.InstanceType_INFERENCE
	default:
		return slav.InstanceType_STARTER
	}
}

func parseStateTypeEntity(in slav.StateType) (result entities.ContainerState) {
	switch in {
	case slav.StateType_RUNNING:
		return entities.ContainerStateRunning
	case slav.StateType_UPDATING:
		return entities.ContainerStateUpdating
	case slav.StateType_ERROR:
		return entities.ContainerStateError
	default:
		return entities.ContainerStateUnknown
	}
}
