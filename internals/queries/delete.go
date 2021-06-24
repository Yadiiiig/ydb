package queries

import (
	"fmt"
	"reflect"
	"strings"

	pb "yadiiig.dev/ydb/internals/proto"
	"yadiiig.dev/ydb/internals/reader"
	"yadiiig.dev/ydb/internals/utils"
)

func Delete(d *reader.Drivers, in *pb.DeleteValues) int32 {
	var amount int32 = 0
	for i, v := range d.Database[in.GetTable()] {
		tempBool := false
		for _, vq := range in.GetValues() {
			if utils.OperatorQuery(reflect.ValueOf(v).Elem().FieldByName(strings.Title(vq.Row)).String(), vq.Value, vq.Operator) {
				tempBool = true
			} else {
				tempBool = false
				break
			}
		}
		if tempBool {
			fmt.Println(len(d.Database[in.GetTable()]), i)
			if i >= len(d.Database[in.GetTable()]) {
				d.Database[in.GetTable()] = d.Database[in.GetTable()][:len(d.Database[in.GetTable()])-1]
			} else {
				d.Database[in.GetTable()] = append(d.Database[in.GetTable()][:i], d.Database[in.GetTable()][i+1:]...)
			}
			amount += 1
			d.Tracker += 1
		}
	}
	return amount
}

// d.Database[in.GetTable()] = append(d.Database[in.GetTable()][:i], d.Database[in.GetTable()][i+1:]...)
// This code is temp because a few things still have to be figured out, e.g. multiple rows.
// utils.UpdateFile(d.OpenFile, v, d.Layout[in.Table], in.GetTable())
