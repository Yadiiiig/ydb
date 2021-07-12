package lib

import (
	"context"

	pb "github.com/Yadiiiig/ydb/drivers/go_driver/src/lib/proto"
)

type DeleteQuery struct {
	Details      *TableQuery
	DeleteValues []*pb.DValues
}

func (s TableQuery) Delete(values [][]string) *DeleteQuery {
	var wv []*pb.DValues
	for _, v := range values {
		wv = append(wv, &pb.DValues{Row: v[0], Operator: v[1], Value: v[2]})
	}

	return &DeleteQuery{
		Details:      &s,
		DeleteValues: wv,
	}
}

func (s DeleteQuery) Run() (bool, int32, error) {
	return DeleteQueryF(s.Details.Conn.Services.deleteService, s.Details.Table, s.DeleteValues)
}

func DeleteQueryF(ec pb.DeleteClient, t string, v []*pb.DValues) (bool, int32, error) {
	r, err := ec.DeleteQuery(context.Background(), &pb.DeleteValues{Table: t, Values: v})
	return r.GetResult(), r.GetAmount(), err
}
