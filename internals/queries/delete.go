package queries

import (
	"fmt"
	"reflect"
	"sort"
	"strings"

	pb "github.com/Yadiiiig/ydb/internals/proto"
	"github.com/Yadiiiig/ydb/internals/reader"
	"github.com/Yadiiiig/ydb/internals/utils"
)

func Delete(d *reader.Drivers, in *pb.DeleteValues) int32 {
	var amount int32 = 0
	deletor(d, d.Database[in.GetTable()], in, &amount)
	return amount
}

func deletor(dTable *reader.Drivers, d []interface{}, in *pb.DeleteValues, amount *int32) {
	i := sort.Search(len(d), func(i int) bool {
		return utils.OperatorQuery(reflect.ValueOf(d[i]).Elem().FieldByName(strings.Title(in.GetValues()[0].Row)).String(), in.GetValues()[0].Value, in.GetValues()[0].Operator)
	})

	if i < len(d) && utils.OperatorQuery(reflect.ValueOf(d[i]).Elem().FieldByName(strings.Title(in.GetValues()[0].Row)).String(), in.GetValues()[0].Value, in.GetValues()[0].Operator) {
		tempBool := true
		for _, v := range in.GetValues() {
			if !utils.OperatorQuery(reflect.ValueOf(d[i]).Elem().FieldByName(strings.Title(v.Row)).String(), v.Value, v.Operator) {
				fmt.Println(false)
				tempBool = false
				break
			}
		}

		if tempBool {
			if i >= len(dTable.Database[in.GetTable()]) {
				dTable.Database[in.GetTable()] = dTable.Database[in.GetTable()][:len(dTable.Database[in.GetTable()])-1]
			} else {
				dTable.Database[in.GetTable()] = append(dTable.Database[in.GetTable()][:i], dTable.Database[in.GetTable()][i+1:]...)
			}
			*amount += 1
			dTable.Tracker += 1

			var value interface{} = &d
			sp := value.(*[]interface{})
			*sp = append((*sp)[:i], (*sp)[i+1:]...)
		}

		deletor(dTable, d, in, amount)
	}
}
