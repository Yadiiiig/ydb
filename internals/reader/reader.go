package reader

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"

	utils "github.com/Yadiiiig/ydb/internals/utils"
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
	Path     string
	Tracker  int
}

func ReadData(path string) (*Drivers, error) {
	if path[len(path)-1:] != "/" {
		path = path + "/"
	}

	var temp_data []TempData
	file, err := os.OpenFile(path+"/structure/layout.json", os.O_RDONLY, 0)
	if err != nil {
		return nil, err
	}

	data, err := os.Open(path + "data.ydb")
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
		return nil, err
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
					return nil, err
				}

				res[table] = append(res[table], obj)
			}
		}
	}

	data.Close()
	openFile := utils.OpenData(path + "data.ydb")

	return &Drivers{
		Database: res,
		Layout:   layout,
		OpenFile: openFile,
		Path:     path,
		Tracker:  0,
	}, nil
}

// Code for an upcoming extra release
// {"name": "ydb_offset", "type": "string"},
//             {"name": "ydb_location", "type": "string"},
// func ReadData(path string) *Drivers {
// 	var temp_data []TempData

// 	files, err := ioutil.ReadDir(path)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	for _, f := range files {
// 		data, err := os.Open(path + f.Name())
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		defer data.Close()

// 		scanner := bufio.NewScanner(data)
// 		offset := 0
// 		for scanner.Scan() {
// 			temp_row := strings.SplitN(scanner.Text(), ":", 2)
// 			offset += len(scanner.Bytes())
// 			temp_row[1] = strings.TrimRight(temp_row[1], "}") + fmt.Sprintf(`, "ydb_offset": "%s", "ydb_location": "%s"}`, strconv.Itoa(offset), f.Name())
// 			fmt.Println(temp_row[1])
// 			fmt.Println(offset)
// 			temp_data = append(temp_data, TempData{Name: temp_row[0], Row: temp_row[1]})
// 		}
// 	}

// 	layoutFile, err := os.OpenFile(path+"/structure/layout.json", os.O_RDONLY, 0)
// 	if err != nil {
// 		return nil, err
// 	}

// 	container := DataContainer{}
// 	if err := json.NewDecoder(layoutFile).Decode(&container); err != nil {
// 		return nil, err
// 	}
// 	fmt.Println(container)
// 	res := map[string][]interface{}{}
// 	layout := map[string][]string{}

// 	for table := range container.Tables {
// 		res[table] = []interface{}{}
// 		layout[table] = make([]string, 0)
// 		rowFields := make([]reflect.StructField, 0, len(container.Tables[table]))
// 		fmt.Println(table)
// 		for _, spec := range container.Tables[table] {

// 			rowFields = append(rowFields, reflect.StructField{
// 				Name: strings.Title(spec.Name),
// 				Type: GetType(spec.Type),
// 				Tag:  reflect.StructTag(fmt.Sprintf(`json:"%s"`, spec.Name)),
// 			})
// 			fmt.Println("hey 3")
// 			layout[table] = append(layout[table], spec.Name)
// 			fmt.Println(layout[table], spec.Name)
// 		}

// 		rowStruct := reflect.StructOf(rowFields)

// 		for i := range temp_data {
// 			if temp_data[i].Name == table {
// 				obj := reflect.New(rowStruct).Interface()
// 				if err := json.Unmarshal([]byte(temp_data[i].Row), &obj); err != nil {
// 					return nil, err
// 				}

// 				res[table] = append(res[table], obj)
// 			}
// 		}
// 	}

// 	//data.Close()
// 	//openFile := OpenData(ydb)

// 	return &Drivers{
// 		Database: res,
// 		Layout:   layout,
// 		// OpenFile: openFile,
// 	}
// }
