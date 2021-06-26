package handler

import (
	"context"

	pb "github.com/Yadiiiig/ydb/internals/proto"
	q "github.com/Yadiiiig/ydb/internals/queries"
	reader "github.com/Yadiiiig/ydb/internals/reader"
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
