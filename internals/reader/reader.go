package reader

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"

	utils "yadiiig.dev/ydb/internals/utils"
)

type DataContainer struct {
	Tables map[string][]FieldSpec `json:"tables"`
}

type FieldSpec struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type TempData struct {
	Name string
	Row  string
}

type Drivers struct {
	Database map[string][]interface{}
	Layout   map[string][]string
	OpenFile *os.File
}

func ReadData(structure string, ydb string) *Drivers {
	var temp_data []TempData
	file, err := os.OpenFile(structure, os.O_RDONLY, 0)
	if err != nil {
		panic(err)
	}

	data, err := os.Open(ydb)
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()

	scanner := bufio.NewScanner(data)

	for scanner.Scan() {
		temp_row := strings.SplitN(scanner.Text(), ":", 2)
		temp_data = append(temp_data, TempData{Name: temp_row[0], Row: temp_row[1]})
	}

	container := DataContainer{}
	if err := json.NewDecoder(file).Decode(&container); err != nil {
		panic(err)
	}

	res := map[string][]interface{}{}
	layout := map[string][]string{}

	for table := range container.Tables {
		res[table] = []interface{}{}
		layout[table] = make([]string, 0)

		rowFields := make([]reflect.StructField, 0, len(container.Tables[table]))
		for _, spec := range container.Tables[table] {
			rowFields = append(rowFields, reflect.StructField{
				Name: strings.Title(spec.Name),
				Type: utils.GetType(spec.Type),
				Tag:  reflect.StructTag(fmt.Sprintf(`json:"%s"`, spec.Name)),
			})
			layout[table] = append(layout[table], spec.Name)
		}

		rowStruct := reflect.StructOf(rowFields)

		for i := range temp_data {
			if temp_data[i].Name == table {
				obj := reflect.New(rowStruct).Interface()
				if err := json.Unmarshal([]byte(temp_data[i].Row), &obj); err != nil {
					panic(err)
				}

				res[table] = append(res[table], obj)
			}
		}
	}

	data.Close()
	openFile := utils.OpenData(ydb)

	return &Drivers{
		Database: res,
		Layout:   layout,
		OpenFile: openFile,
	}
}
