package queries

import (
	"reflect"
	"strings"

	pb "github.com/Yadiiiig/ydb/internals/proto"
	"github.com/Yadiiiig/ydb/internals/reader"
	"github.com/Yadiiiig/ydb/internals/utils"
)

func Update(d *reader.Drivers, in *pb.UpdateValues) int32 {
	var amount int32 = 0
	for _, v := range d.Database[in.GetTable()] {
		tempBool := false
		for _, vq := range in.GetMatchers() {
			if utils.OperatorQuery(reflect.ValueOf(v).Elem().FieldByName(strings.Title(vq.Row)).String(), vq.Value, vq.Operator) {
				tempBool = true
			} else {
				tempBool = false
				break
			}
		}
		if tempBool {
			// Should a field be updated even tho it has the exact same value?
			for _, vz := range in.GetValues() {
				reflect.ValueOf(v).Elem().FieldByName(strings.Title(vz.Row)).SetString(vz.Value)
			}
			amount += 1
			d.Tracker += 1
		}
	}
	return amount
}
