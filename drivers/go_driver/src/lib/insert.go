package lib

import (
	"context"

	pb "github.com/Yadiiiig/ydb/drivers/go_driver/src/lib/proto"
)

type InsertQuery struct {
	Details      *TableQuery
	InsertValues []*pb.IValues
}

func (s TableQuery) Insert(values interface{}) *InsertQuery {
	var vals []*pb.IValues
	tmpMap := StructPrepare(values)
	for k, v := range tmpMap {
		tmp := &pb.IValues{Row: k, Value: v.(string)}
		vals = append(vals, tmp)
	}
	return &InsertQuery{
		Details:      &s,
		InsertValues: vals,
	}
}

func (s InsertQuery) Run() (bool, error) {
	return InsertQueryF(s.Details.Conn.Services.insertService, s.Details.Table, s.InsertValues)
}

func InsertQueryF(ec pb.InsertClient, t string, v []*pb.IValues) (bool, error) {
	r, err := ec.InsertQuery(context.Background(), &pb.InsertValues{Table: t, Values: v})
	return r.GetResult(), err
}
