package handler

import (
	"context"

	pb "yadiiig.dev/ydb/internals/proto"
	q "yadiiig.dev/ydb/internals/queries"
	"yadiiig.dev/ydb/internals/reader"
)

type insertService struct {
	pb.UnimplementedInsertServer
	Drivers *reader.Drivers
}

func NewInsertService(d *reader.Drivers) *insertService {
	return &insertService{
		Drivers: d,
	}
}

func (i *insertService) InsertQuery(ctx context.Context, in *pb.InsertValues) (*pb.InsertResponse, error) {
	_ = q.Insert(i.Drivers, in)
	return &pb.InsertResponse{
		Result: true,
	}, nil
}

// type insertService struct {
// 	pb.UnimplementedSelectServer
// 	Drivers reader.Drivers
// }

// func NewInsertService(d reader.Drivers) *selectService {
// 	return &selectService{
// 		Drivers: d,
// 	}
// }

// func (s *selectService) InsertQuery(ctx context.Context, in *pb.SelectValues) (*pb.SelectResponse, error) {
// 	result := q.Insert(s.Drivers, in)
// 	return &pb.SelectResponse{Result: result}, nil
// }
