package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"reflect"
	"strings"
)

func HandleQuery(rec []byte, connection net.Conn) {
	query := base{}
	userB := bytes.NewReader(rec)

	if err := json.NewDecoder(userB).Decode(&query); err != nil {
		panic(err)
	}

	//fmt.Println(len(query.Action))

	switch query.Action {
	case "insert":
		Insert(query, connection)
	case "select":
		Select(query, connection)
	default:
		fmt.Println("Not implemented yet")
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
