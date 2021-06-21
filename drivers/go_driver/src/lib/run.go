package lib

import (
	"encoding/json"

	pb "yadiiig.dev/ydb/go_driver/src/lib/proto"
)

// Will change
func (s Query) Run() {
	var vals []*pb.IValues
	tmpMap := structPrepare(s.InsertValues)
	for k, v := range tmpMap {
		tmp := &pb.IValues{Row: k, Value: v.(string)}
		vals = append(vals, tmp)
	}
	s.Details.Conn.Ctx.InsertQuery(s.Details.Conn.Services.insertService, s.Details.Table, vals)
}

func structPrepare(s interface{}) map[string]interface{} {
	var sm map[string]interface{}
	inrec, _ := json.Marshal(s)
	json.Unmarshal(inrec, &sm)
	return sm
}
