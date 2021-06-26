package lib

import (
	pb "github.com/Yadiiiig/ydb/go_driver/src/lib/proto"
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
	return s.Details.Conn.Ctx.DeleteQuery(s.Details.Conn.Services.deleteService, s.Details.Table, s.DeleteValues)
}

func (ctx Ctx) DeleteQuery(ec pb.DeleteClient, t string, v []*pb.DValues) (bool, int32, error) {
	r, err := ec.DeleteQuery(ctx.Context, &pb.DeleteValues{Table: t, Values: v})
	return r.GetResult(), r.GetAmount(), err
}
