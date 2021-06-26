package lib

import (
	pb "yadiiig.dev/ydb/go_driver/src/lib/proto"
)

type UpdateQuery struct {
	Details      *TableQuery
	MatchValues  []*pb.MatchValues
	UpdateValues []*pb.UValues
}

func (s TableQuery) Update(match [][]string, update [][]string) *UpdateQuery {
	var mv []*pb.MatchValues
	var uv []*pb.UValues

	for _, v := range match {
		mv = append(mv, &pb.MatchValues{Row: v[0], Operator: v[1], Value: v[2]})
	}

	for _, v := range update {
		uv = append(uv, &pb.UValues{Row: v[0], Value: v[1]})
	}

	return &UpdateQuery{
		Details:      &s,
		MatchValues:  mv,
		UpdateValues: uv,
	}
}

func (s UpdateQuery) Run() (bool, int32, error) {
	return s.Details.Conn.Ctx.UpdateQuery(s.Details.Conn.Services.updateService, s.Details.Table, s.MatchValues, s.UpdateValues)
}

func (ctx Ctx) UpdateQuery(ec pb.UpdateClient, t string, v []*pb.MatchValues, d []*pb.UValues) (bool, int32, error) {
	r, err := ec.UpdateQuery(ctx.Context, &pb.UpdateValues{Table: t, Matchers: v, Values: d})
	return r.GetResult(), r.GetAmount(), err
}
