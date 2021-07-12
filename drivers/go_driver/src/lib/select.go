package lib

import (
	"context"
	"encoding/json"

	pb "github.com/Yadiiiig/ydb/drivers/go_driver/src/lib/proto"
)

type SelectQuery struct {
	Details     *TableQuery
	Rows        []string
	Limit       int
	WhereValues []*pb.SValues
	CastDest    interface{}
}

func (s TableQuery) Select(dest interface{}, tables ...string) *SelectQuery {
	return &SelectQuery{
		Details:  &s,
		Rows:     tables,
		CastDest: dest,
	}
}

func (s SelectQuery) Run() error {
	// Still need to implement a .limit
	if s.Rows == nil {
		r, err := SelectQueryF(s.Details.Conn.Services.selectService, s.Details.Table, []string{"*"}, s.WhereValues)
		if err != nil {
			return err
		}

		err = json.Unmarshal([]byte(r), &s.CastDest)
		return err

	} else {
		r, err := SelectQueryF(s.Details.Conn.Services.selectService, s.Details.Table, s.Rows, s.WhereValues)
		if err != nil {
			return err
		}

		err = json.Unmarshal([]byte(r), &s.CastDest)
		return err
	}
}

func SelectQueryF(ec pb.SelectClient, t string, f []string, v []*pb.SValues) (string, error) {
	r, err := ec.SelectQuery(context.Background(), &pb.SelectValues{Table: t, Fields: f, Values: v})
	return r.GetResult(), err
}
