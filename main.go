package main

// we use net to make a tcp connection but not http

// 1. handle GET request

import (
	// "bufio"
	"fmt"
	"net"
	"strings"
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

		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	buffer := make([]byte, 1024)
	conn.Read(buffer)
	req := string(buffer)
	if req == "" {
		fmt.Println("Empty Request")
		return
	}

	method := getRequestMethod(req)
	fmt.Println("Method:", method)

	var response *Response
	switch method {
	case "GET":
		request := parseGetRequest(req)
		response = createResponse(200, map[string]string{"Content-Type": "text/html; charset=utf-8"}, request.Body)
	case "POST":
		request := parsePostRequest(req)
		response = createResponse(200, map[string]string{"Content-Type": "application/json"}, request.Body)
	default:
		response = createResponse(405, map[string]string{"Content-Type": "text/plain"}, "Method Not Allowed")
	}

	sendResponse(conn, response)
	fmt.Println("Response:", response)
}

func parseJsonFromString(s string) {
}

type Request struct {
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

func createResponse(statusCode int, headers map[string]string, body string) *Response {
	return &Response{
		StatusCode: statusCode,
		Headers:    headers,
		Body:       body,
	}
}

func createStatusLine(statusCode int) string {
	switch statusCode {
	case 200:
		return "HTTP/1.1 200 OK"
	case 404:
		return "HTTP/1.1 404 Not Found"
	case 500:
		return "HTTP/1.1 500 Internal Server Error"
	case 405:
		return "HTTP/1.1 405 Method Not Allowed"
	default:
		return "HTTP/1.1 200 OK"
	}
}

func sendResponse(conn net.Conn, responseObject *Response) {

	statusline := createStatusLine(responseObject.StatusCode)
	headers := responseObject.Headers
	headers["Content-Length"] = fmt.Sprint(len(sampleBody))
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
	conn.Write([]byte(responseObject.Body))
}

func getRequestMethod(req string) string {
	return strings.Split(req, " ")[0]
}

func parseHeaders(lines []string) map[string]string {

	headers := make(map[string]string)
	for _, line := range lines[1:] {
		// if we hit an empty line, headers are done
		if line == "" {
			break
		}
		// Host: localhost:8080
		// User-Agent: curl/7.64.1
		headersParts := strings.Split(line, ":")
		// key : value
		headers[headersParts[0]] = headersParts[1]
	}

	return headers
}

func parseGetRequest(req string) *Request {
	lines := strings.Split(req, "\r\n")
	requestLine := lines[0]
	parts := strings.Fields(requestLine)
	headers := parseHeaders(lines[1:])

	request := &Request{
		Method:  parts[0],
		URI:     parts[1],
		Version: parts[2],
		Headers: headers,
		Body:    sampleBody,
	}
	return request
}

func parseBody(lines []string) string {
	// body example: {"name": "John Doe", "email": "john.doe@example.com"}
	// if body is empty
	if len(lines) == 0 {
		return ""
	}
	return strings.Join(lines, "\r\n")
}
func parsePostRequest(req string) *Request {
	lines := strings.Split(req, "\r\n")
	requestLine := lines[0]
	parts := strings.Fields(requestLine)

	headers := parseHeaders(lines[1:])
	body := parseBody(lines[len(headers)+2:])

	parsedRequest := &Request{
		Method:  parts[0],
		URI:     parts[1],
		Version: parts[2],
		Headers: headers,
		Body:    body,
	}

	return parsedRequest
}
