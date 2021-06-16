package handler

import (
	"context"

	pb "yadiiig.dev/ydb/internals/proto"
	q "yadiiig.dev/ydb/internals/queries"
	reader "yadiiig.dev/ydb/internals/reader"
)

type deleteService struct {
	pb.UnimplementedDeleteServer
	Drivers *reader.Drivers
}

func NewDeleteService(d *reader.Drivers) *deleteService {
	return &deleteService{
		Drivers: d,
	}
}

// DeleteQuery will return result (bool) and amount (int).
// Result will only return false if amount equals to zero, so no records were deleted.
// Amount will not be send over if result equals to zero, client has to check for the bool first.
func (s *deleteService) DeleteQuery(ctx context.Context, in *pb.DeleteValues) (*pb.DeleteResponse, error) {
	a := q.Delete(s.Drivers, in)
	if a != 0 {
		return &pb.DeleteResponse{
			Result: true,
			Amount: a,
		}, nil
	} else {
		return &pb.DeleteResponse{
			Result: false,
		}, nil
	}
}
