package server

import (
	"fmt"
	"net"
	"tonic/internal/handlers"
	// "tonic/internal/server"
)

func HandleRequest(conn net.Conn) {
	buffer := make([]byte, 1024)
	conn.Read(buffer)
	req := string(buffer)
	if req == "" {
		fmt.Println("Empty Request")
		return
	}

	method := handlers.GetRequestMethod(req)
	fmt.Println("Method:", method)

	var response *Response
	switch method {
	case "GET":
		request := handlers.ParseGetRequest(req)
		response = CreateResponse(200, map[string]string{"Content-Type": "text/html; charset=utf-8"}, request.Body)
	case "POST":
		request := handlers.ParsePostRequest(req)
		response = CreateResponse(200, map[string]string{"Content-Type": "application/json"}, request.Body)
	default:
		response = CreateResponse(405, map[string]string{"Content-Type": "text/plain"}, "Method Not Allowed")
	}

	SendResponse(conn, response)
	fmt.Println("Response:", response)
}
