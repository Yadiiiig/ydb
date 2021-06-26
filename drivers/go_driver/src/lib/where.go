package lib

import pb "github.com/Yadiiiig/ydb/go_driver/src/lib/proto"

type WhereValues struct {
	Row      string
	Operator string
	Value    string
}

// Could probably make this a generic function, although will the where function
// ever be used in another query?
func (s SelectQuery) Where(values [][]string) *SelectQuery {
	var wv []*pb.SValues
	for _, v := range values {
		wv = append(wv, &pb.SValues{Row: v[0], Operator: v[1], Value: v[2]})
	}
	s.WhereValues = wv
	return &s
}
