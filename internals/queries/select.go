package queries

import (
	"encoding/json"
	"fmt"
	"reflect"
	"sort"
	"strings"

	pb "github.com/Yadiiiig/ydb/internals/proto"
	"github.com/Yadiiiig/ydb/internals/reader"
	"github.com/Yadiiiig/ydb/internals/utils"
)

func Select(d *reader.Drivers, in *pb.SelectValues) string {
	tl := d.Database[in.GetTable()]
	result := []interface{}{}

	selector(tl, in, &result)

	r, _ := json.Marshal(result)
	return string(r)
}

func selector(d []interface{}, in *pb.SelectValues, result *[]interface{}) {
	i := sort.Search(len(d), func(i int) bool {
		return utils.OperatorQuery(reflect.ValueOf(d[i]).Elem().FieldByName(strings.Title(in.GetValues()[0].Row)).String(), in.GetValues()[0].Value, in.GetValues()[0].Operator)
	})

	if i < len(d) && utils.OperatorQuery(reflect.ValueOf(d[i]).Elem().FieldByName(strings.Title(in.GetValues()[0].Row)).String(), in.GetValues()[0].Value, in.GetValues()[0].Operator) {
		tempBool := true
		for _, v := range in.GetValues() {
			if !utils.OperatorQuery(reflect.ValueOf(d[i]).Elem().FieldByName(strings.Title(v.Row)).String(), v.Value, v.Operator) {
				tempBool = false
				break
			}
		}
		if tempBool {
			if len(in.GetFields()) == 1 && in.GetFields()[0] == "*" {
				*result = append(*result, d[i])
			} else {
				*result = append(*result, appendResult(d[i], in.GetFields()))
			}

			var value interface{} = &d
			sp := value.(*[]interface{})
			*sp = append((*sp)[:i], (*sp)[i+1:]...)
		}
		selector(d, in, result)
	}
}

func appendResult(value interface{}, fields []string) interface{} {
	res := []interface{}{}
	temp := ""

	for _, v := range fields {
		temp += fmt.Sprintf(`"%s": "%s", `, v, reflect.ValueOf(value).Elem().FieldByName(strings.Title(v)).String())
	}
	res = append(res, fmt.Sprintf("{%s}", strings.TrimSuffix(temp, ", ")))
	return res[0]

}
