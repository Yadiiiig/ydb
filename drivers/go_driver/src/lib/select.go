package lib

import (
	pb "yadiiig.dev/ydb/go_driver/src/lib/proto"
)

type SelectQuery struct {
	QueryType   string
	Details     *TableQuery
	Rows        []string
	Limit       int
	WhereValues []*pb.SValues
}

func (s TableQuery) Select(tables ...string) *SelectQuery {
	return &SelectQuery{
		QueryType: "select",
		Details:   &s,
		Rows:      tables,
	}
}

func (s SelectQuery) Run() (string, error) {
	// Still need to implement a .limit
	// Being able to cast into a struct
	if s.Rows == nil {
		return s.Details.Conn.Ctx.SelectQuery(s.Details.Conn.Services.selectService, s.Details.Table, []string{"*"}, s.WhereValues)
	} else {
		return s.Details.Conn.Ctx.SelectQuery(s.Details.Conn.Services.selectService, s.Details.Table, s.Rows, s.WhereValues)
	}
}

func (ctx Ctx) SelectQuery(ec pb.SelectClient, t string, f []string, v []*pb.SValues) (string, error) {
	r, err := ec.SelectQuery(ctx.Context, &pb.SelectValues{Table: t, Fields: f, Values: v})
	return r.GetResult(), err
}
