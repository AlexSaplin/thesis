package ibis

import (
	"github.com/mattn/go-nulltype"
	uuid "github.com/satori/go.uuid"
	ibis "lynx/pkg/clients/ibis/pb"
	"lynx/pkg/entities"
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

func makeFunctionStatePb(in entities.FunctionState) (result ibis.FunctionState, err error) {
	switch in {
	case entities.FunctionStateInit:
		result = ibis.FunctionState_STATE_INIT
	case entities.FunctionStateProcessing:
		result = ibis.FunctionState_STATE_PROCESSING
	case entities.FunctionStateReady:
		result = ibis.FunctionState_STATE_READY
	case entities.FunctionStateInvalid:
		result = ibis.FunctionState_STATE_INVALID
	case entities.FunctionStateDeleted:
		result = ibis.FunctionState_STATE_DELETED
	default:
		result = ibis.FunctionState_STATE_UNKNOWN
	}
	return
}

func makeFunctionEntity(in *ibis.Function) (result entities.Function, err error) {
	modelID, err := uuid.FromString(in.ID)
	if err != nil {
		return
	}

	ownerID, err := uuid.FromString(in.OwnerID)
	if err != nil {
		return
	}

	state, err := parseFunctionStateEntity(in.State)
	if err != nil {
		return
	}
	var errStr nulltype.NullString
	if in.ErrStr != "" {
		errStr = nulltype.NullStringOf(in.ErrStr)
	}

	result = entities.Function{
		ID:       modelID,
		OwnerID:  ownerID,
		Name:     in.Name,
		State:    state,
		ImageURL: nulltype.NullString{},
		CodePath: nulltype.NullString{},
		Metadata: nulltype.NullString{},
		Error:    errStr,
	}
	return
}