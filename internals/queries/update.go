package queries

import (
	"reflect"
	"sort"
	"strings"

	pb "github.com/Yadiiiig/ydb/internals/proto"
	"github.com/Yadiiiig/ydb/internals/reader"
	"github.com/Yadiiiig/ydb/internals/utils"
)

func Update(d *reader.Drivers, in *pb.UpdateValues) int32 {
	var amount int32 = 0
	updator(d, d.Database[in.GetTable()], in, &amount)
	return amount
}

func updator(dTable *reader.Drivers, d []interface{}, in *pb.UpdateValues, amount *int32) {
	i := sort.Search(len(d), func(i int) bool {
		return utils.OperatorQuery(reflect.ValueOf(d[i]).Elem().FieldByName(strings.Title(in.GetMatchers()[0].Row)).String(), in.GetMatchers()[0].Value, in.GetMatchers()[0].Operator)
	})

	if i < len(d) && utils.OperatorQuery(reflect.ValueOf(d[i]).Elem().FieldByName(strings.Title(in.GetMatchers()[0].Row)).String(), in.GetMatchers()[0].Value, in.GetMatchers()[0].Operator) {
		tempBool := true
		for _, v := range in.GetMatchers() {
			if !utils.OperatorQuery(reflect.ValueOf(d[i]).Elem().FieldByName(strings.Title(v.Row)).String(), v.Value, v.Operator) {
				tempBool = false
				break
			}
		}

		if tempBool {
			for _, vz := range in.GetValues() {
				reflect.ValueOf(dTable.Database[in.GetTable()][i]).Elem().FieldByName(strings.Title(vz.Row)).SetString(vz.Value)
			}
			*amount += 1
			dTable.Tracker += 1

			var value interface{} = &d
			sp := value.(*[]interface{})
			*sp = append((*sp)[:i], (*sp)[i+1:]...)
		}

		updator(dTable, d, in, amount)
	}
}
