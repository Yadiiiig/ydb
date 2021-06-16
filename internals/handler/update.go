package handler

import (
	"context"

	pb "yadiiig.dev/ydb/internals/proto"
	q "yadiiig.dev/ydb/internals/queries"
	reader "yadiiig.dev/ydb/internals/reader"
)

type updateService struct {
	pb.UnimplementedUpdateServer
	Drivers *reader.Drivers
}

func NewUpdateService(d *reader.Drivers) *updateService {
	return &updateService{
		Drivers: d,
	}
}

func (s *updateService) UpdateQuery(ctx context.Context, in *pb.UpdateValues) (*pb.UpdateResponse, error) {
	a := q.Update(s.Drivers, in)
	if a != 0 {
		return &pb.UpdateResponse{
			Result: true,
			Amount: a,
		}, nil
	} else {
		return &pb.UpdateResponse{
			Result: false,
		}, nil
	}
}
