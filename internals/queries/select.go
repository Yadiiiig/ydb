package queries

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	pb "yadiiig.dev/ydb/internals/proto"
	"yadiiig.dev/ydb/internals/reader"
)

func Select(d *reader.Drivers, in *pb.SelectValues) string {
	fmt.Println(in)

	result := []interface{}{}
	for _, v := range d.Database[in.GetTable()] {
		tempBool := true
		for _, vq := range in.GetValues() {
			if !operatorQuery(reflect.ValueOf(v).Elem().FieldByName(strings.Title(vq.Row)).String(), vq.Value, vq.Operator) {
				tempBool = false
				break
			}
		}
		if tempBool {
			//fmt.Println(v)
			if len(in.GetFields()) == 1 && in.GetFields()[0] == "*" {
				result = append(result, v)
			} else {
				result = append(result, appendResult(v, in.GetFields()))
			}
		}
	}

	r, _ := json.Marshal(result)

	return string(r)
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

func operatorQuery(rowVal string, val string, conditional string) bool {
	switch conditional {
	case "=":
		return rowVal == val
	case ">":
		return rowVal > val
	case "<":
		return rowVal < val
	case ">=":
		return rowVal >= val
	case "<=":
		return rowVal <= val
	case "!=":
		return rowVal != val
	}
	return false
}
