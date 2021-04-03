package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

func HandleQuery(rec []byte) {
	query := base{}
	userB := bytes.NewReader(rec)
	m := map[string]interface{}{}

	if err := json.NewDecoder(userB).Decode(&query); err != nil {
		panic(err)
	}

	fmt.Println(len(query.Data))

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

func appendDatabase(tableName string, row interface{}, rawRow map[string]interface{}) {
	var tempString string

	for i := range layout[tableName] {
		tempElement := reflect.ValueOf(row).Elem().FieldByName(strings.Title(layout[tableName][i])).String()
		tempString += fmt.Sprintf(`"%s": "%s", `, layout[tableName][i], tempElement)
	}

	fullString := fmt.Sprintf("%s: {%s}\n", tableName, strings.TrimSuffix(tempString, ", "))
	database[tableName] = append(database[tableName], row)

	if _, err := dataFile.WriteString(fullString); err != nil {
		fmt.Println(err)
	}
}

func rowExists(name string, table string) bool {
	for _, v := range layout[table] {
		if v == name {
			return true
		}
	}
	return false
}
