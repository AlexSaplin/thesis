package grpc

import (
	"github.com/mattn/go-nulltype"
	"ibis/pkg/entities"
	pb "ibis/pkg/grpc/pb"
)

func parseFunctionStateEntity(in pb.FunctionState) (result entities.FunctionState, err error) {
	switch in {
	case pb.FunctionState_STATE_INIT:
		result = entities.FunctionStateInit
	case pb.FunctionState_STATE_PROCESSING:
		result = entities.FunctionStateProcessing
	case pb.FunctionState_STATE_READY:
		result = entities.FunctionStateReady
	case pb.FunctionState_STATE_INVALID:
		result = entities.FunctionStateInvalid
	case pb.FunctionState_STATE_DELETED:
		result = entities.FunctionStateDeleted
	default:
		result = entities.FunctionStateUnknown
	}
	return
}

func serializeFunctionStateEntity(in entities.FunctionState) (result pb.FunctionState, err error) {
	switch in {
	case entities.FunctionStateInit:
		result = pb.FunctionState_STATE_INIT
	case entities.FunctionStateProcessing:
		result = pb.FunctionState_STATE_PROCESSING
	case entities.FunctionStateReady:
		result = pb.FunctionState_STATE_READY
	case entities.FunctionStateInvalid:
		result = pb.FunctionState_STATE_INVALID
	case entities.FunctionStateDeleted:
		result = pb.FunctionState_STATE_DELETED
	default:
		result = pb.FunctionState_STATE_UNKNOWN
	}
	return
}

func serializeFunctionEntity(in entities.Function) (result pb.Function, err error) {
	var state pb.FunctionState
	state, err = serializeFunctionStateEntity(in.State)
	if err != nil {
		return
	}
	return pb.Function{
		ID:          in.ID.String(),
		OwnerID:     in.OwnerID.String(),
		State:       state,
		Name:        in.Name,
		ErrStr:      in.Error.String(),
		ImageURL:    in.ImageURL.String(),
		CodePath:    in.CodePath.String(),
		Metadata:    in.Metadata.String(),
	}, nil
}

func parseUpdateParams(in []*pb.UpdateFunctionParam) (result entities.UpdateFunctionParam, err error) {
	for _, v := range in {
		if v == nil {
			continue
		}
		switch vt := v.Param.(type) {
		case *pb.UpdateFunctionParam_CodePath:
			result.CodePath = nulltype.NullStringOf(vt.CodePath)
		case *pb.UpdateFunctionParam_ErrStr:
			result.ErrStr = nulltype.NullStringOf(vt.ErrStr)
		case *pb.UpdateFunctionParam_ImageURL:
			result.ImageURL = nulltype.NullStringOf(vt.ImageURL)
		case *pb.UpdateFunctionParam_State:
			var state entities.FunctionState
			state, err = parseFunctionStateEntity(vt.State)
			result.State = new(entities.FunctionState)
			*result.State = state
		case *pb.UpdateFunctionParam_Metadata:
			result.Metadata = nulltype.NullStringOf(vt.Metadata)
		}
	}
	return
}