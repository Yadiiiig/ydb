package main

type base struct {
	Table  string   `json:"table"`
	Action string   `json:"action"`
	Data   []string `json:"data"`
}
