package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

var (
	database map[string][]interface{}
	layout   map[string][]string
	dataFile *os.File
)

func main() {
	fmt.Println("Starting Database")
	database, layout = ReadData("./idea.json")
	fmt.Println("Database is ready to Go")

	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalln(err)
	}
	defer listener.Close()

	for {
		con, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go handleClientRequest(con)
	}
}

func handleClientRequest(con net.Conn) {
	defer con.Close()

	clientReader := bufio.NewReader(con)

	for {
		clientRequest, err := clientReader.ReadBytes('\n')

		switch err {
		case nil:
			clientRequestTrimmed := strings.TrimSpace(string(clientRequest))
			if clientRequestTrimmed == ":QUIT" {
				log.Println("client requested server to close the connection so closing")
				return
			}
		case io.EOF:
			log.Println("client closed the connection by terminating the process")
			return
		default:
			log.Printf("error: %v\n", err)
			return
		}

		HandleQuery(clientRequest, con)
		//con.Write([]byte("done" + "\n"))
	}
}
