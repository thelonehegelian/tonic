package main

// we use net to make a tcp connection but not http

import (
	// "bufio"
	"fmt"
	"net"
	"strings"
)

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
	fmt.Println(listener)

	defer listener.Close()

	// keep the connection open
	for {

		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting:", err.Error())
		}
		go handleRequest(conn)

	}
}

func handleRequest(conn net.Conn) {
	buffer := make([]byte, 1024)
	conn.Read(buffer)
	request := string(buffer)
	if request == "" {
		fmt.Println("Empty Request")
	}
	parseHtmlRequest(request)
	sendResponse(conn)
	// if r.Method == "POST" {
	// }

}

func parseJsonFromString(s string) {

}
func handlePostRequest(r *IncomingRequest) {
	if r.Headers["Content-Type"] == "application/json" {
		parseJsonFromString(r.Body)
	}
}

type IncomingRequest struct {
	Method  string
	URI     string
	Version string

	Headers map[string]string
	Body    string
}

type Response struct {
	StatusCode int
	Headers    map[string]string
	Body       string
}

func sendResponse(conn net.Conn) {
	body := "<html><body><h1>Hello World</h1></body></html>"
	statusline := "HTTP/1.1 200 OK"
	headers := make(map[string]string)
	headers["Content-Type"] = "text/html; charset=utf-8"
	headers["Content-Length"] = fmt.Sprint(len(body))
	/*
		Example:
		HTTP/1.1 200 OK
		Content-Type: text/html; charset=utf-8
		Content-Length: 13

		<html><body><h1>Hello World</h1></body></html>
	*/

	conn.Write([]byte(statusline + "\r\n"))
	for k, v := range headers {
		conn.Write([]byte(k + ": " + v + "\r\n"))
	}
	conn.Write([]byte("\r\n"))
	conn.Write([]byte(body))
}

func parseHtmlRequest(req string) *IncomingRequest {
	lines := strings.Split(req, "\r\n")
	requestLine := lines[0]
	parts := strings.Fields(requestLine)

	headers := make(map[string]string)
	for _, line := range lines[1:] {
		// if we hit an empty line, headers are done
		if line == "" {
			break
		}
		// Host: localhost:8080
		// User-Agent: curl/7.64.1
		headerParts := strings.Split(line, ":")
		// key : value
		headers[headerParts[0]] = headerParts[1]
		fmt.Printf("%s : %s\n", headerParts[0], headerParts[1])
	}

	// body example: {"name": "John Doe", "email": "john.doe@example.com"}
	body := ""
	if len(lines) > len(headers)+2 { // Check if there's a body, considering headers and blank line
		body = strings.Join(lines[len(headers)+2:], "\r\n") // Join lines after the headers if they exist
		fmt.Println("Body:", body)
	}

	parsedRequest := &IncomingRequest{
		Method:  parts[0],
		URI:     parts[1],
		Version: parts[2],
		Headers: headers,
		Body:    body,
	}

	return parsedRequest

}
