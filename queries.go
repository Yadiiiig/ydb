package main

type base struct {
	Table         string        `json:"table"`
	Action        string        `json:"action"`
	Data          []string      `json:"data"` // Needs to be replaced
	selectDetails selectDetails `json:"selectDetails"`
	Select        []sValues     `json:"selectQ"`
}

// type insertQ struct {

// }

// Everything for select queries
type selectDetails struct {
	Action string  `json:"action"`
	Fields sFields `json:"fields"`
	//Select []selectV `json:"selectV"`
}

type sFields struct {
	Amount int      `json:"amount"`
	Fields []string `json:"fields"`
}

type sValues struct {
	Operator string `json:"operator"`
	Row      string `json:"row"`
	Value    string `json:"value"`
}
