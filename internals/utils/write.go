package utils

import (
	"fmt"
	"os"
	"reflect"
	"strings"
)

func WriteRow(tableName, row interface{}, layout []string, file *os.File) {
	var temp string

	for i := range layout {
		temp += fmt.Sprintf(`"%s": "%s", `, layout[i], reflect.ValueOf(row).Elem().FieldByName(strings.Title(layout[i])))
	}

	res := fmt.Sprintf("%s: {%s}\n", tableName, strings.TrimSuffix(temp, ", "))
	if _, err := file.WriteString(res); err != nil {
		fmt.Println(err)
	}
}

// func WriteRow(tableName string, row interface{}, rawRow map[string]interface{}) {
// 	var tempString string

// 	for i := range layout[tableName] {
// 		tempElement := reflect.ValueOf(row).Elem().FieldByName(strings.Title(layout[tableName][i])).String()
// 		tempString += fmt.Sprintf(`"%s": "%s", `, layout[tableName][i], tempElement)
// 	}

// 	fullString := fmt.Sprintf("%s: {%s}\n", tableName, strings.TrimSuffix(tempString, ", "))

// 	if _, err := dataFile.WriteString(fullString); err != nil {
// 		fmt.Println(err)
// 	}
// }
