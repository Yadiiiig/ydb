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

func UpdateFile(file *os.File, d map[string][]interface{}, l map[string][]string, path string) bool {
	if err := CopyFile(path); err != nil {
		return false
	}

	err := file.Truncate(0)
	if err != nil {
		fmt.Println(err)
	}

	for k, v := range d {
		for _, x := range v {
			WriteRow(k, x, l[k], file)
		}
	}
	return true
}

func UpdatePeriodically(file *os.File, d map[string][]interface{}, l map[string][]string, path string) {

}

// func UpdateFile(file *os.File, row interface{}, layout []string, tableName string) bool {
// 	var temp string
// 	// Need to replace this into seperate function, this needs to happen before this function is called
// 	for i := range layout {
// 		temp += fmt.Sprintf(`"%s": "%s", `, layout[i], reflect.ValueOf(row).Elem().FieldByName(strings.Title(layout[i])))
// 	}

// 	res := fmt.Sprintf("%s: {%s}\n", tableName, strings.TrimSuffix(temp, ", "))

// 	f, _ := os.Open("data.ydb")
// 	defer f.Close()

// 	pattern := filepath.Base("data.ydb") + "-temp-*"
// 	tmp, err := ioutil.TempFile("./", pattern)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	defer os.Remove(tmp.Name())
// 	defer tmp.Close()

// 	if _, err := io.Copy(tmp, transform.NewReader(f, replace.String(res, ""))); err != nil {
// 		fmt.Println(err)
// 	}

// 	if err := tmp.Close(); err != nil {
// 		fmt.Println(err)
// 	}

// 	if err := f.Close(); err != nil {
// 		fmt.Println(err)
// 	}

// 	os.Rename(tmp.Name(), "data.ydb")

// 	return true
// }
