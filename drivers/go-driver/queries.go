package main

import (
	pb "yadiiig.dev/ydb/go-driver/proto"
)

type Query struct {
	QueryType    string
	Details      TableQuery
	InsertValues interface{}
}

func (s TableQuery) insert(values interface{}) *Query {
	return &Query{
		QueryType:    "insert",
		Details:      s,
		InsertValues: values,
	}
}

func (ctx Ctx) Insert(ec pb.InsertClient, t string, v []*pb.IValues) (bool, error) {
	r, err := ec.InsertQuery(ctx.Context, &pb.InsertValues{Table: t, Values: v})
	return r.GetResult(), err
}
