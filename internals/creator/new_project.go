package creator

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func Create(path string) error {
	template := `{
		"tables": {
			"example": [
				{"name": "id", "type": "string"},
				{"name": "name", "type": "string"}
			]
		}
	}`

	if path[len(path)-1:] != "/" {
		path = path + "/"
	}

	err := os.Mkdir(path+"database/", 0755)
	if err != nil {
		return err
	}

	err = os.Mkdir(path+"database/structure/", 0755)
	if err != nil {
		return err
	}

	f, err := os.OpenFile(path+"database/data.ydb", os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	f.Close()

	var obj map[string]interface{}
	json.Unmarshal([]byte(template), &obj)

	file, _ := json.MarshalIndent(obj, "", "    ")
	err = ioutil.WriteFile(path+"database/structure/layout.json", file, 0644)
	if err != nil {
		return err
	}
	return nil
}
