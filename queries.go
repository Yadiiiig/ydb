package main

type base struct {
	Table         string        `json:"table"`
	Action        string        `json:"action"`
	Data          []string      `json:"data"` // Needs to be replaced
	SelectDetails selectDetails `json:"selectDetails"`
}

// type insertQ struct {

// }

// Everything for select queries
type selectDetails struct {
	Action string    `json:"action"` // Need to implement
	Fields sFields   `json:"fields"`
	Values []sValues `json:"values"`
	//Select []selectV `json:"selectV"`
}

type sFields struct {
	Everything string   `json:"everything"`
	Amount     int      `json:"amount"`
	Fields     []string `json:"fields"`
}

type sValues struct {
	Operator string `json:"operator"`
	Row      string `json:"row"`
	Value    string `json:"value"`
}
