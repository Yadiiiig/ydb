package queries

import (
	"fmt"
	"reflect"
	"strings"

	pb "yadiiig.dev/ydb/internals/proto"
	"yadiiig.dev/ydb/internals/reader"
	utils "yadiiig.dev/ydb/internals/utils"
)

func Insert(d *reader.Drivers, in *pb.InsertValues) (bool, error) {
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
	if err := utils.WriteRow(in.Table, row, d.Layout[in.Table], d.OpenFile); err != nil {
		return false, err
	}
	return true, nil
}
