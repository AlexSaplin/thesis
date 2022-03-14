package grpc

import (
	"gorilla/pkg/entities"
	pb "gorilla/pkg/grpc/pb"
)

func makeDeltaPb(in entities.Delta) (result pb.Delta) {
	return pb.Delta{
		Date:       in.Date.Unix(),
		Category:   in.Category,
		Balance:    in.Balance,
		OwnerID:    in.OwnerID.String(),
		ObjectID:   in.ObjectID,
		ObjectType: in.ObjectType,
	}
}
