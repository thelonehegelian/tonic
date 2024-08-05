package main

// we use net to make a tcp connection but not http

// 1. handle GET request

import (
	// "bufio"
	"fmt"
	"net"

	"tonic/internal/server"
)

const sampleBody = `<html><body><h1>Hello World</h1></body></html>`
const listeningAddress = "localhost:8080"

/*
request-line = method SP request-URI SP HTTP-version CRLF

headers = header CRLF
body = body

Example:
POST /api/users HTTP/1.1
Host: localhost:8080
User-Agent: curl/7.64.1
Accept:
Content-Type: application/json
Content-Length: 54

{"name": "John Doe", "email": "john.doe@example.com"}
*/

func main() {
	// create a listener
	listener, err := net.Listen("tcp", listeningAddress)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
	}

	fmt.Println("Listening on", listeningAddress)
	defer listener.Close()

	// keep the connection open
	for {

		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting:", err.Error())
		}
		fmt.Println("Accepted connection")

		go server.HandleRequest(conn)
	}
}
