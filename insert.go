package main

import (
	"encoding/json"
	"fmt"
	"net"
	"reflect"
	"strings"
)

func Insert(query base, connection net.Conn) {
	m := map[string]interface{}{}
	tempRowStruct := make([]reflect.StructField, 0, len(layout[query.Table]))

	for _, spec := range layout[query.Table] {
		tempRowStruct = append(tempRowStruct, reflect.StructField{
			Name: strings.Title(spec),
			Type: reflect.TypeOf(""),
			Tag:  reflect.StructTag(fmt.Sprintf(`json:"%s"`, spec)),
		})
	}
	rowStruct := reflect.StructOf(tempRowStruct)

	for _, v := range query.Data {
		if err := json.Unmarshal([]byte(v), &m); err != nil {
			panic(err)
		}

		for k := range m {
			if !rowExists(k, query.Table) {
				delete(m, k)
			}
		}

		temp, err := json.Marshal(m)
		if err != nil {
			fmt.Println(err)
		}

		obj := reflect.New(rowStruct).Interface()
		if err := json.Unmarshal([]byte(temp), &obj); err != nil {
			fmt.Println(err)
		}

		appendDatabase(query.Table, obj, m)
	}
}
