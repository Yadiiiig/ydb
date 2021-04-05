package main

import (
	"encoding/json"
	"fmt"
	"net"
	"reflect"
	"strings"
)

var (
	bNewline = []byte("\n")
)

func Select(query base, connection net.Conn) {
	testingV := []interface{}{}
	for _, v := range database[query.Table] {
		for _, vq := range query.Select {
			if strings.EqualFold(reflect.ValueOf(v).Elem().FieldByName(strings.Title(vq.Row)).String(), vq.Value) {
				testingV = append(testingV, v)
			}
		}
	}
	send(testingV, connection)
}

func send(results interface{}, connection net.Conn) {
	tempRes, err := json.Marshal(results)
	if err != nil {
		fmt.Println(err)
	}

	finalR := append(tempRes, bNewline...)
	connection.Write(finalR)
}
