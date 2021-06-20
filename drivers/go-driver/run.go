package main

import (
	"encoding/json"

	pb "yadiiig.dev/ydb/go-driver/proto"
)

// Will change
func (s Query) run() {
	var vals []*pb.IValues
	tmpMap := structPrepare(s.InsertValues)
	for k, v := range tmpMap {
		tmp := &pb.IValues{Row: k, Value: v.(string)}
		vals = append(vals, tmp)
	}
	s.Details.Conn.Ctx.Insert(s.Details.Conn.Services.insertService, s.Details.Table, vals)
}

func structPrepare(s interface{}) map[string]interface{} {
	var sm map[string]interface{}
	inrec, _ := json.Marshal(s)
	json.Unmarshal(inrec, &sm)
	return sm
}
