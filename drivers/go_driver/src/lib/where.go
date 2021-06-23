package lib

import pb "yadiiig.dev/ydb/go_driver/src/lib/proto"

type WhereValues struct {
	Row      string
	Operator string
	Value    string
}

func (s SelectQuery) Where(values [][]string) *SelectQuery {
	var wv []*pb.SValues
	for _, v := range values {
		wv = append(wv, &pb.SValues{Row: v[0], Operator: v[1], Value: v[2]})
	}
	s.WhereValues = wv
	return &s
}
