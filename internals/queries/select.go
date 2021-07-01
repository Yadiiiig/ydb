package queries

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"sync"

	pb "github.com/Yadiiiig/ydb/internals/proto"
	"github.com/Yadiiiig/ydb/internals/reader"
	utils "github.com/Yadiiiig/ydb/internals/utils"
)

func Select(d *reader.Drivers, in *pb.SelectValues) string {
	result := []interface{}{}
	amount := len(d.Database[in.GetTable()])
	var wg sync.WaitGroup

	for i := 0; amount > 0; i += 1000 {
		amount -= 1000
		if amount > 0 {
			wg.Add(1)
			go selector(d, in, &result, i, i+1000, &wg)
		} else {
			tmp := 1000 + amount
			wg.Add(1)
			go selector(d, in, &result, i, i+tmp, &wg)
		}
	}
	wg.Wait()
	r, _ := json.Marshal(result)
	return string(r)
}

func selector(d *reader.Drivers, in *pb.SelectValues, result *[]interface{}, x, y int, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, v := range d.Database[in.GetTable()][x:y] {
		tempBool := true
		for _, vq := range in.GetValues() {
			if !utils.OperatorQuery(reflect.ValueOf(v).Elem().FieldByName(strings.Title(vq.Row)).String(), vq.Value, vq.Operator) {
				tempBool = false
				break
			}
		}
		if tempBool {
			if len(in.GetFields()) == 1 && in.GetFields()[0] == "*" {
				*result = append(*result, v)
			} else {
				*result = append(*result, appendResult(v, in.GetFields()))
			}
		}
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
