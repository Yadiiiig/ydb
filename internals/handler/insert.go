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
	r, err := q.Insert(i.Drivers, in)
	return &pb.InsertResponse{
		Result: r,
	}, err
}
