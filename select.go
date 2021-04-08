package main

import (
	"encoding/json"
	"fmt"
	"net"
	"reflect"
	"strings"
)

var (
	bNewline = []byte("\n")
)

func Select(query base, connection net.Conn) {
	result := []interface{}{}
	for _, v := range database[query.Table] {
		tempBool := true
		for _, vq := range query.SelectDetails.Values {
			if !operatorQuery(reflect.ValueOf(v).Elem().FieldByName(strings.Title(vq.Row)).String(), vq.Value, vq.Operator) {
				tempBool = false
				break
			}
		}
		if tempBool {
			fmt.Println(v)
			if query.SelectDetails.Fields.Everything == "*" {
				result = append(result, v)
			} else {
				result = append(result, appendResult(v, query.SelectDetails.Fields))
			}
		}
	}
	fmt.Println(result)
	send(result, connection)
}

func send(results interface{}, connection net.Conn) {
	tempRes, err := json.Marshal(results)
	if err != nil {
		fmt.Println(err)
	}

	finalR := append(tempRes, bNewline...)
	connection.Write(finalR)
}

func appendResult(value interface{}, fields sFields) interface{} {
	res := []interface{}{}
	temp := ""
	//var testing interface{}
	// tempRes := `"%s": "%s",`
	for _, v := range fields.Fields {
		//tempRow := fmt.Sprintf(`"%s": "%s", `, v, reflect.ValueOf(value).Elem().FieldByName(strings.Title(v)).String())
		//res = append(res, strings.TrimSuffix(tempRow, ", "))
		temp += fmt.Sprintf(`"%s": "%s", `, v, reflect.ValueOf(value).Elem().FieldByName(strings.Title(v)).String())
		// Need to change this to include error checking, user can add row that doesn;t exist
	}
	res = append(res, fmt.Sprintf("{%s}", strings.TrimSuffix(temp, ", ")))
	fmt.Println(res)
	//res = append(res, strings.TrimSuffix(temp, ", "))
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

// for vs :=  vq. range {
// 	if operator(reflect.ValueOf(v).Elem().FieldByName(strings.Title(vq.Row),
// 			query.Select.sValues.Value, query.Select.sValues.Operator) {
// 		testingV = append(testingV, v)
// 	}
// }
// if strings.EqualFold(reflect.ValueOf(v).Elem().FieldByName(strings.Title(vq.Row)).String(), vq.Value) {
// 	testingV = append(testingV, v)
// }
