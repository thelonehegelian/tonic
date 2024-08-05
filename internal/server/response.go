package server

import (
	"fmt"
	"net"
)

type Response struct {
	StatusCode int
	Headers    map[string]string
	Body       string
}

func SendResponse(conn net.Conn, responseObject *Response) {

	statusline := createStatusLine(responseObject.StatusCode)
	headers := responseObject.Headers
	headers["Content-Length"] = fmt.Sprint(len(responseObject.Body))
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
func CreateResponse(statusCode int, headers map[string]string, body string) *Response {
	return &Response{
		StatusCode: statusCode,
		Headers:    headers,
		Body:       body,
	}
}
