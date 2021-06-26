package queries

import (
	"fmt"
	"reflect"
	"strings"

	pb "github.com/Yadiiiig/ydb/internals/proto"
	"github.com/Yadiiiig/ydb/internals/reader"
	utils "github.com/Yadiiiig/ydb/internals/utils"
	"github.com/google/uuid"
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

	v := reflect.ValueOf(row).Elem().FieldByName("Id")
	if v.IsValid() {
		v.SetString(uuid.New().String())
	}

	d.Database[in.GetTable()] = append(d.Database[in.GetTable()], row)
	if err := utils.WriteRow(in.Table, row, d.Layout[in.Table], d.OpenFile); err != nil {
		return false, err
	}
	return true, nil
}
