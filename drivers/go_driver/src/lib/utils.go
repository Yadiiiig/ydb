package lib

import "encoding/json"

func StructPrepare(s interface{}) map[string]interface{} {
	var sm map[string]interface{}
	inrec, _ := json.Marshal(s)
	json.Unmarshal(inrec, &sm)
	return sm
}
