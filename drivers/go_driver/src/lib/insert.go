package lib

import (
	pb "yadiiig.dev/ydb/go_driver/src/lib/proto"
)

type InsertQuery struct {
	QueryType    string
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
		QueryType:    "insert",
		Details:      &s,
		InsertValues: vals,
	}
}

func (s InsertQuery) Run() (bool, error) {
	return s.Details.Conn.Ctx.InsertQuery(s.Details.Conn.Services.insertService, s.Details.Table, s.InsertValues)
}

func (ctx Ctx) InsertQuery(ec pb.InsertClient, t string, v []*pb.IValues) (bool, error) {
	r, err := ec.InsertQuery(ctx.Context, &pb.InsertValues{Table: t, Values: v})
	return r.GetResult(), err
}