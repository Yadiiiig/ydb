package handler

import (
	"context"

	pb "yadiiig.dev/ydb/internals/proto"
	q "yadiiig.dev/ydb/internals/queries"
	reader "yadiiig.dev/ydb/internals/reader"
)

type selectService struct {
	pb.UnimplementedSelectServer
	Drivers reader.Drivers
}

func NewSelectService(d reader.Drivers) *selectService {
	return &selectService{
		Drivers: d,
	}
}

func (s *selectService) SelectQuery(ctx context.Context, in *pb.SelectValues) (*pb.SelectResponse, error) {
	result := q.Select(s.Drivers, in)
	return &pb.SelectResponse{Result: result}, nil
}
