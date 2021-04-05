package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"
)

var (
	conn *net.TCPConn
)

func main() {
	var err error
	conn, err = connectDatabase("127.0.0.1:8081")
	if err != nil {
		fmt.Println(err)
		log.Fatal("Database could not start.")
	}

	//test()
	selectTest()
}

func selectTest() {
	test := testStruct{}
	test.Table = "users"
	test.Action = "select"

	testSelect := []selectQ{}
	testSelect = append(testSelect, selectQ{Row: "firstname", Value: "James"})
	test.SelectQ = testSelect
	fmt.Println("Selecting user where Firstname == James")
	send(conn, test)
}

// func test() {
// 	data := []exampleData{}

// 	file, _ := ioutil.ReadFile("data.json")
// 	_ = json.Unmarshal([]byte(file), &data)
// 	//fmt.Println(data)
// 	test := testStruct{}
// 	test.Table = "users"
// 	test.Action = "insert"
// 	strs := []interface{}{}
// 	for i := range data {
// 		dataTemp, err := json.Marshal(data[i])
// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		strs = append(strs, string(dataTemp))
// 	}

// 	test.Data = strs
// 	// {"id": "aaaa-1234-bbbb-5678", "firstname": "Bill", "lastname": "Smith", "email": "bill.smith@gmail.com", "password": "supergoodhashedpassword"}
// 	fmt.Println("Sending rn")
// 	send(conn, test)
// }

func connectDatabase(address string) (*net.TCPConn, error) {
	var err error
	TCPAdrr, err := net.ResolveTCPAddr("tcp", address)
	if err != nil {
		return nil, err
	}

	conn, err = net.DialTCP("tcp", nil, TCPAdrr)
	if err != nil {
		println("Dial failed:", err.Error())
		return nil, err
	}

	conn.SetKeepAlive(true)
	return conn, nil
}

func send(conn *net.TCPConn, data testStruct) {
	bytesSend, _ := json.Marshal(data)
	bytesSend = append(bytesSend, '\n')
	conn.Write(bytesSend)
	//fmt.Fprintf(conn, bytesSend+[]byte("\n"))

	//message, _ := bufio.NewReader(conn).ReadString('\n')
	message, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(message)
}

type exampleData struct {
	ID        string `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Company   string `json:"company"`
}

type testStruct struct {
	Table   string        `json:"table"`
	Action  string        `json:"action"`
	Data    []interface{} `json:"data"`
	SelectQ []selectQ     `json:"selectQ"`
}

type selectQ struct {
	Row   string `json:"row"`
	Value string `json:"value"`
}
