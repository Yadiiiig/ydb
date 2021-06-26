package handler

import (
	"context"

	pb "github.com/Yadiiiig/ydb/internals/proto"
	q "github.com/Yadiiiig/ydb/internals/queries"
	reader "github.com/Yadiiiig/ydb/internals/reader"
)

type selectService struct {
	pb.UnimplementedSelectServer
	Drivers *reader.Drivers
}

func NewSelectService(d *reader.Drivers) *selectService {
	return &selectService{
		Drivers: d,
	}
}

func (s *selectService) SelectQuery(ctx context.Context, in *pb.SelectValues) (*pb.SelectResponse, error) {
	result := q.Select(s.Drivers, in)
	return &pb.SelectResponse{Result: result}, nil
}
