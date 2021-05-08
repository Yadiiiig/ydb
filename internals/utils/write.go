package utils

import (
	"fmt"
	"os"
	"reflect"
	"strings"
)

func WriteRow(tableName, row interface{}, layout []string, file *os.File) error {
	var temp string

	for i := range layout {
		temp += fmt.Sprintf(`"%s": "%s", `, layout[i], reflect.ValueOf(row).Elem().FieldByName(strings.Title(layout[i])))
	}

	res := fmt.Sprintf("%s: {%s}\n", tableName, strings.TrimSuffix(temp, ", "))
	if _, err := file.WriteString(res); err != nil {
		return err
	}
	return nil
}
