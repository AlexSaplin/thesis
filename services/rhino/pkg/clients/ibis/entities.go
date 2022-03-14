package ibis

import (
	ibis "rhino/pkg/clients/ibis/pb"
	"rhino/pkg/entities"
)

func parseFunctionStateEntity(in ibis.FunctionState) (result entities.FunctionState, err error) {
	switch in {
	case ibis.FunctionState_STATE_INIT:
		result = entities.FunctionStateInit
	case ibis.FunctionState_STATE_PROCESSING:
		result = entities.FunctionStateProcessing
	case ibis.FunctionState_STATE_READY:
		result = entities.FunctionStateReady
	case ibis.FunctionState_STATE_INVALID:
		result = entities.FunctionStateInvalid
	case ibis.FunctionState_STATE_DELETED:
		result = entities.FunctionStateDeleted
	default:
		result = entities.FunctionStateUnknown
	}
	return
}

