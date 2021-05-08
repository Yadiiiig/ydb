package queries

import (
	"fmt"
	"reflect"
	"strings"

	pb "yadiiig.dev/ydb/internals/proto"
	"yadiiig.dev/ydb/internals/reader"
	utils "yadiiig.dev/ydb/internals/utils"
)

func Insert(d *reader.Drivers, in *pb.InsertValues) error {
	//values := in.GetValues()
	//m := map[string]interface{}{}
	tempRowStruct := make([]reflect.StructField, 0, len(d.Layout[in.Table]))

	for _, spec := range d.Layout[in.GetTable()] {
		tempRowStruct = append(tempRowStruct, reflect.StructField{
			Name: strings.Title(spec),
			Type: reflect.TypeOf(""),
			Tag:  reflect.StructTag(fmt.Sprintf(`json:"%s"`, spec)),
		})
	}
	rowStruct := reflect.StructOf(tempRowStruct)
	row := reflect.New(rowStruct).Interface()

	for _, i := range in.Values {
		v := reflect.ValueOf(row).Elem().FieldByName(strings.Title(i.Row))
		if v.IsValid() {
			v.SetString(i.Value)
		}
	}
	d.Database[in.GetTable()] = append(d.Database[in.GetTable()], row)
	utils.WriteRow(in.Table, row, d.Layout[in.Table], d.OpenFile)
	//appendDatabase(query.Table, obj, m)
	// tempString := ""
	// tempString += `"Id":\"50x",\`
	// for _, k := range in.GetValues() {
	// 	tempString += fmt.Sprintf(`"%s":\"%s",\`, strings.Title(k.GetRow()), k.GetValue())
	// }
	// stringValues := fmt.Sprintf("{%s}", strings.TrimSuffix(tempString, `,\`))
	// fmt.Println(stringValues)

	// temp, err := json.Marshal(stringValues)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// obj := reflect.New(rowStruct).Interface()
	// if err := json.Unmarshal(temp, &obj); err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println(obj)

	// for _, k := range values {
	// 	utils.RowExists(k.Row, d.Layout)
	// 	delete(values, k)
	// }

	// for _, v := range query.Data {
	// 	if err := json.Unmarshal([]byte(v), &m); err != nil {
	// 		panic(err)
	// 	}

	// 	for k := range m {
	// 		if !rowExists(k, d.Layout) {
	// 			delete(m, k)
	// 		}
	// 	}

	// temp, err := json.Marshal(m)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// obj := reflect.New(rowStruct).Interface()
	// if err := json.Unmarshal([]byte(temp), &obj); err != nil {
	// 	fmt.Println(err)
	// }

	// 	appendDatabase(query.Table, obj, m)
	// }
	return nil
}
